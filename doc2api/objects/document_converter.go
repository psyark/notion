package objects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Converter) OutputAllBuilders() {
	for _, c := range c.comparators {
		c.finish()
		c.builder.output(false)
	}
	c.globalBuilder.output(true)
	c.globalTestBuilder.output(true)
}

func (c *Converter) getUnionInterface(name string) *UnionInterface {
	for _, symbol := range c.globalBuilder.symbols {
		switch symbol := symbol.(type) {
		case *UnionInterface:
			if symbol.name() == name {
				return symbol
			}
		}
	}
	return nil
}
func (c *Converter) getDiscriminatorString(name string) *DiscriminatorString {
	for _, symbol := range c.globalBuilder.symbols {
		switch symbol := symbol.(type) {
		case *DiscriminatorString:
			if symbol.name() == name {
				return symbol
			}
		}
	}
	return nil
}
func (c *Converter) getUnmarshalTest(name string) *UnmarshalTest {
	for _, symbol := range c.globalTestBuilder.symbols {
		switch symbol := symbol.(type) {
		case *UnmarshalTest:
			if symbol.name() == name {
				return symbol
			}
		}
	}
	return nil
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
