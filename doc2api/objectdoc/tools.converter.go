// Package doc2api は Notion API Reference の更新を検知し、Goコードへの適切な変換を
// 継続的に行うための一連の仕組みを提供します。
//
// Goコードへの変換ルールは、命令形のコードではなくデータ構造として objects.*.go に格納されます。
// このデータ構造には Notion API Referenceのローカルコピーも含まれるため、
// ドキュメントの更新に対してGoコードへの変換ルールが古いままになることを防ぎます。
package objectdoc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
	"golang.org/x/sync/errgroup"
)

// converter はNotion API ReferenceからGoコードへの変換ルールです。
type converter struct {
	url       string // ドキュメントのURL
	localCopy []objectDocElement
}

// fetchAndValidate はリモートドキュメントの取得・ローカルコピーとの比較・クラスの組立てを行います
func (c converter) fetchAndBuild(global *builder) (*builder, error) {
	// URLの取得
	res, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}

	defer func() {
		_, _ = io.ReadAll(res.Body)
		_ = res.Body.Close()
	}()

	// goqueryでのパース
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	ssrProps := struct {
		Doc struct {
			Body string `json:"body"`
		} `json:"doc"`
	}{}
	if err := json.Unmarshal(ssrPropsBytes, &ssrProps); err != nil {
		return nil, err
	}

	lines := strings.Split(ssrProps.Doc.Body, "\n")
	odt := &objectDocTokenizer{lines, 0}

	requiredCopy := jen.Statement{}
	fileName := "object." + strings.TrimPrefix(c.url, "https://developers.notion.com/reference/") + ".go"

	b := &builder{
		global:        global,
		globalSymbols: global.globalSymbols,
		fileName:      "../../" + fileName,
		url:           c.url,
	}

	for i := 0; ; i++ {
		remote, err := odt.next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		if len(c.localCopy) < i+1 {
			requiredCopy = append(requiredCopy, createLocalCopy(remote))
		} else if err := c.localCopy[i].checkAndOutput(remote, b); err != nil {
			if errors.Is(err, errUnmatch) {
				ld, _ := json.MarshalIndent(c.localCopy[i], "", "  ")
				rd, _ := json.MarshalIndent(remote, "", "  ")
				return nil, fmt.Errorf("%w\nlocal:\n%#v\nremote:\n%#v", err, string(ld), string(rd))
			} else {
				return nil, err
			}
		}
	}

	if len(requiredCopy) != 0 {
		gostr := jen.Var().Id("LOCAL_COPY").Op("=").Index().Id("objectDocElement").Values(jen.List(requiredCopy...), jen.Line()).GoString()
		if err := os.WriteFile("tmp/"+fileName, []byte(gostr), 0666); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("localCopyが足りません (see tmp/%s)", fileName)
	} else {
		_ = os.Remove("tmp/" + fileName)
	}

	return b, nil
}

func createLocalCopy(remote objectDocElement) jen.Code {
	typeName := strings.TrimPrefix(fmt.Sprintf("%T", remote), "*objectdoc.")
	outputCode := jen.Func().Params(jen.Id("e").Op("*").Id(typeName), jen.Id("b").Op("*").Id("builder")).Error().Block(jen.Return().Nil().Comment("TODO"))
	return remote.localCopy(typeName, outputCode)
}

var registeredConverters []converter

// registerConverter は後で実行するためにconverterを登録します
func registerConverter(c converter) {
	registeredConverters = append(registeredConverters, c)
}

// convertAll は登録された全てのconverterで変換を実行します
func convertAll() error {
	global := &builder{
		globalSymbols: &sync.Map{},
		fileName:      "../../object.global.go",
	}
	builders := []*builder{global}

	copyAlways := func(b *builder) {
		for _, c := range b.localSymbols {
			switch c := c.(type) {
			case *abstractObject:
				for _, f := range c.fields {
					if f, ok := f.(*fixedStringField); ok {
						global.addAlwaysStringIfNotExists(f.value)
					}
				}
			case *specificObject:
				for _, f := range c.fields {
					if f, ok := f.(*fixedStringField); ok {
						global.addAlwaysStringIfNotExists(f.value)
					}
				}
			}
		}
	}

	eg := errgroup.Group{}
	for _, c := range registeredConverters {
		c := c
		eg.Go(func() error {
			if b, err := c.fetchAndBuild(global); err != nil {
				return fmt.Errorf("fetchAndBuild: %s: %w", c.url, err)
			} else {
				builders = append(builders, b)
				copyAlways(b)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	{
		cases := []jen.Code{}
		for _, v := range global.getAbstractObject("PropertyValue").variants {
			cases = append(cases, jen.Case(jen.Op("*").Id(strings.TrimSuffix(v.name(), "Value"))))
			var field fieldCoder
			for _, f := range v.(*specificObject).fields {
				if f.goName() != "Type" && f.goName() != "HasMore" {
					field = f
					break
				}
			}

			typeCodeStr := (jen.Var().Id("_").Add(field.getTypeCode())).GoString()
			typeCodeStr = strings.TrimPrefix(typeCodeStr, "var _ ")
			cases = append(cases, jen.Return().Lit(typeCodeStr))
		}

		file := jen.NewFile("notion")
		file.Func().Id("getTypeForBinding").Params(jen.Id("p").Id("Property")).String().Block(
			jen.Switch(jen.Id("p").Assert(jen.Type())).Block(cases...),
			jen.Panic(jen.Id("p")),
		)
		if err := file.Save("../../binding.helper.go"); err != nil {
			return err
		}
	}

	copyAlways(global)

	// グローバルビルダーをソートし、冪等性を保ちます
	sort.Slice(global.localSymbols, func(i, j int) bool {
		return global.localSymbols[i].name() < global.localSymbols[j].name()
	})
	for _, c := range global.localSymbols {
		if c, ok := c.(*abstractObject); ok {
			sort.Slice(c.variants, func(i, j int) bool {
				return c.variants[i].getSpecifyingField(c.specifiedBy).value < c.variants[j].getSpecifyingField(c.specifiedBy).value
			})
		}
	}

	for _, b := range builders {
		if len(b.localSymbols) != 0 {
			file := jen.NewFile("notion")
			file.Comment("Code generated by notion.doc2api; DO NOT EDIT.")
			if b.url != "" {
				file.Comment(b.url)
			}
			for _, s := range b.localSymbols {
				file.Line().Line().Add(s.symbolCode())
			}
			if err := file.Save(b.fileName); err != nil {
				return err
			}
		} else {
			_ = os.Remove(b.fileName)
		}

		// テストコード
		if len(b.unmarshalTests) != 0 {
			file := jen.NewFile("notion")
			file.Comment("Code generated by notion.doc2api; DO NOT EDIT.")
			if b.url != "" {
				file.Comment(b.url)
			}

			for name, tests := range b.unmarshalTests {
				file.Line().Func().Id("Test"+name+"_unmarshal").Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
					jen.Id("u").Op(":=").Id(strcase.LowerCamelCase(name)+"Unmarshaler").Values(),
					jen.Id("tests").Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
						for _, t := range tests {
							g.Line().Lit(t)
						}
						g.Line()
					}),
					jen.For().List(jen.Id("_"), jen.Id("wantStr")).Op(":=").Range().Id("tests").Block(
						jen.Id("want").Op(":=").Index().Byte().Call(jen.Id("wantStr")),
						jen.If(jen.Err().Op(":=").Id("u").Dot("UnmarshalJSON").Call(jen.Id("want"))).Op(";").Err().Op("!=").Nil().Block(
							jen.Id("t").Dot("Fatal").Call(jen.Err()),
						),
						jen.List(jen.Id("got"), jen.Id("_")).Op(":=").Qual("encoding/json", "MarshalIndent").Call(jen.Id("u").Dot("value"), jen.Lit(""), jen.Lit("  ")),
						jen.If(jen.List(jen.Id("want"), jen.Id("got"), jen.Id("ok")).Op(":=").Id("compareJSON").Call(jen.Id("want"), jen.Id("got")).Op(";").Op("!").Id("ok")).Block(
							jen.Id("t").Dot("Fatal").Call(jen.Qual("fmt", "Errorf").Call(jen.Lit("mismatch:\nwant: %s\ngot : %s"), jen.Id("want"), jen.Id("got"))),
						),
					),
				)
			}

			if err := file.Save(strings.Replace(b.fileName, ".go", "_test.go", 1)); err != nil {
				return err
			}
		} else {
			_ = os.Remove(strings.Replace(b.fileName, ".go", "_test.go", 1))
		}
	}

	return nil
}
