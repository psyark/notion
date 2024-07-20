package objects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/samber/lo"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// Converter はObjectsの多数のドキュメントから一連のコードを生成する機能を提供します
type Converter struct {
	globalBuilder       *CodeBuilder
	globalTestBuilder   *CodeBuilder
	comparators         []*DocumentComparator
	unionMemberRegistry []unionMemberEntry
}

func NewConverter() *Converter {
	c := &Converter{}
	c.globalBuilder = &CodeBuilder{converter: c, fileName: "objects_global_generated.go"}
	c.globalTestBuilder = &CodeBuilder{converter: c, fileName: "objects_global_generated_test.go"}
	return c
}

// FetchDocument は オブジェクトのドキュメントを取得し、
// そのドキュメントの更新を検知するための DocumentComparator インスタンスと、
// それを通じてコードを生成するための CodeBuilder インスタンスを提供します
func (c *Converter) FetchDocument(url string) *DocumentComparator {
	res := lo.Must(http.Get(url))
	defer res.Body.Close()

	doc := lo.Must(goquery.NewDocumentFromReader(res.Body))

	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	ssrProps := struct {
		Doc struct {
			Body string `json:"body"`
		} `json:"doc"`
	}{}

	lo.Must0(json.Unmarshal(ssrPropsBytes, &ssrProps))
	elements := []DocumentElement{}

	// goldmarkを使ってMarkdownのパース
	ren := &docRenderer{OnElement: func(element DocumentElement) {
		elements = append(elements, element)
	}}
	md := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(parser.WithBlockParsers(util.Prioritized(&specialBlockParser{}, 0))),
		goldmark.WithRenderer(ren),
	)

	lo.Must0(md.Convert([]byte(ssrProps.Doc.Body), io.Discard))

	comparator := &DocumentComparator{
		elements: elements,
		builder: &CodeBuilder{
			url:       url,
			fileName:  fmt.Sprintf("objects_%s_generated.go", strings.TrimPrefix(url, "https://developers.notion.com/reference/")),
			converter: c,
		},
	}
	c.comparators = append(c.comparators, comparator)
	return comparator
}

func (c *Converter) AddUnmarshalTest(targetName string, jsonCode string, typeArg ...string) {
	ut := &UnmarshalTest{targetName: targetName} // UnmarshalTestを作る
	if len(typeArg) != 0 {
		ut.typeArg = typeArg[0]
	}

	if exists := c.getUnmarshalTest(ut.name()); exists != nil { // 同名のものが既にあるなら
		exists.jsonCodes = append(exists.jsonCodes, jsonCode) // JSONコードだけ追加
	} else { // 無ければ追加
		ut.jsonCodes = append(ut.jsonCodes, jsonCode)
		c.globalTestBuilder.symbols = append(c.globalTestBuilder.symbols, ut)
	}
}

// NewDiscriminatorField は、ドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (c *Converter) NewDiscriminatorField(p *Parameter) *DiscriminatorField {
	for _, value := range []string{p.ExampleValue, p.Type} {
		if match := regexp.MustCompile(`^"(\w+)"$`).FindStringSubmatch(value); match != nil {
			value = match[1]

			// 未登録なら DiscriminatorString を登録
			if ds := DiscriminatorString(value); c.getDiscriminatorString(ds.name()) == nil {
				c.globalBuilder.symbols = append(c.globalBuilder.symbols, ds)
			}

			return &DiscriminatorField{name: p.Property, value: value, comment: p.Description}
		}
	}
	panic(fmt.Errorf("パラメータ %v には文字列リテラルが含まれません", p))
}

func (c *Converter) OutputAllBuilders() {
	for _, c := range c.comparators {
		c.finish()
		c.builder.output(false)
	}
	c.globalBuilder.output(true)
	c.globalTestBuilder.output(true)
}

func findSymbol[T Symbol](builder *CodeBuilder, name string) T {
	for _, symbol := range builder.symbols {
		switch symbol := symbol.(type) {
		case T:
			if symbol.name() == name {
				return symbol
			}
		}
	}
	var zero T
	return zero
}

func (c *Converter) getUnionInterface(name string) *UnionInterface {
	return findSymbol[*UnionInterface](c.globalBuilder, name)
}
func (c *Converter) getDiscriminatorString(name string) *DiscriminatorString {
	return findSymbol[*DiscriminatorString](c.globalBuilder, name)
}
func (c *Converter) getUnmarshalTest(name string) *UnmarshalTest {
	return findSymbol[*UnmarshalTest](c.globalTestBuilder, name)
}

// RegisterUnionInterface は、指定された名前と判別子を持つ UnionInterfaceを定義し、返します。
// 二回目以降の呼び出しでは定義をスキップし、初回に定義されたものを返します。
func (c *Converter) RegisterUnionInterface(name string, discriminator string) *UnionInterface {
	if union := c.getUnionInterface(name); union != nil {
		return union
	}

	union := &UnionInterface{discriminator: discriminator}
	union.name_ = strings.TrimSpace(name)

	c.globalBuilder.symbols = append(c.globalBuilder.symbols, union)

	return union
}

// RegisterUnionMember は、UnionInterface と unionInterfaceMember を互いの親子として登録します。
func (c *Converter) RegisterUnionMember(union *UnionInterface, member unionInterfaceMember, typeArg string) {
	if member.isGeneric() && typeArg == "" {
		panic(fmt.Errorf("🚨 ジェネリック型 %sに対する型引数がありません", member.name()))
	}
	c.unionMemberRegistry = append(c.unionMemberRegistry, unionMemberEntry{union, member, typeArg})
}

type unionMemberEntry struct {
	union   *UnionInterface
	member  unionInterfaceMember
	typeArg string
}
