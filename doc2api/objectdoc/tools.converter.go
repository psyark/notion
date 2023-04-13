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

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
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

	defer res.Body.Close()

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
	b := &builder{global: global, fileName: "../../" + fileName, url: c.url}

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

	switch remote := remote.(type) {
	case *objectDocHeadingElement:
		return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
			jen.Id("Text"):   jen.Lit(remote.Text),
			jen.Id("output"): outputCode,
		})
	case *objectDocParagraphElement:
		return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
			jen.Id("Text"):   jen.Lit(remote.Text),
			jen.Id("output"): outputCode,
		})
	case *objectDocCalloutElement:
		return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
			jen.Id("Type"):   jen.Lit(remote.Type),
			jen.Id("Title"):  jen.Lit(remote.Title),
			jen.Id("Body"):   jen.Lit(remote.Body),
			jen.Id("output"): outputCode,
		})
	case *objectDocCodeElement:
		outputCode = jen.Func().Params(jen.Id("e").Op("*").Id("objectDocCodeElementCode"), jen.Id("b").Op("*").Id("builder")).Error().Block(jen.Return().Nil().Comment("TODO"))
		codes := jen.Statement{}
		for _, c := range remote.Codes {
			codes = append(codes, jen.Values(jen.Dict{
				jen.Id("Name"):     jen.Lit(c.Name),
				jen.Id("Language"): jen.Lit(c.Language),
				jen.Id("Code"):     jen.Lit(c.Code),
				jen.Id("output"):   outputCode,
			}))
		}
		return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
			jen.Id("Codes"): jen.Index().Op("*").Id("objectDocCodeElementCode").Values(codes...),
		})
	case *objectDocParametersElement:
		outputCode = jen.Func().Params(jen.Id("e").Op("*").Id("objectDocParameter"), jen.Id("b").Op("*").Id("builder")).Error().Block(jen.Return().Nil().Comment("TODO"))
		params := jen.Statement{}
		for _, p := range *remote {
			values := jen.Dict{
				jen.Id("Type"):        jen.Lit(p.Type),
				jen.Id("Description"): jen.Lit(p.Description),
				jen.Id("output"):      outputCode,
			}
			if p.Property != "" {
				values[jen.Id("Property")] = jen.Lit(p.Property)
			}
			if p.Field != "" {
				values[jen.Id("Field")] = jen.Lit(p.Field)
			}
			if p.ExampleValue != "" {
				values[jen.Id("ExampleValue")] = jen.Lit(p.ExampleValue)
			}
			if p.ExampleValues != "" {
				values[jen.Id("ExampleValues")] = jen.Lit(p.ExampleValues)
			}
			params = append(params, jen.Values(values))
		}
		return jen.Line().Op("&").Id(typeName).Values(params...)
	default:
		panic(remote)
	}
}

var registeredConverters []converter

// registerConverter は後で実行するためにconverterを登録します
func registerConverter(c converter) {
	registeredConverters = append(registeredConverters, c)
}

// convertAll は登録された全てのconverterで変換を実行します
func convertAll() error {
	global := &builder{fileName: "../../object.global.go"}
	builders := []*builder{global}

	eg := errgroup.Group{}
	for _, c := range registeredConverters {
		c := c
		eg.Go(func() error {
			if b, err := c.fetchAndBuild(global); err != nil {
				return fmt.Errorf("fetchAndBuild: %s: %w", c.url, err)
			} else {
				builders = append(builders, b)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	// グローバルビルダーをソートし、冪等性を保ちます
	sort.Slice(global.coders, func(i, j int) bool {
		o1, ok1 := global.coders[i].(*abstractObject)
		o2, ok2 := global.coders[j].(*abstractObject)
		if ok1 && ok2 {
			return o1.name < o2.name
		}
		return false
	})
	for _, c := range global.coders {
		if c, ok := c.(*abstractObject); ok {
			sort.Slice(c.variants, func(i, j int) bool {
				return c.variants[i].getName() < c.variants[j].getName()
			})
		}
	}

	for _, b := range builders {
		if len(b.coders) != 0 {
			file := jen.NewFile("notion")
			file.Comment("Code generated by notion.doc2api; DO NOT EDIT.")
			if b.url != "" {
				file.Comment(b.url)
			}
			file.Add(b.statement()...)
			if err := file.Save(b.fileName); err != nil {
				return err
			}
		} else {
			_ = os.Remove(b.fileName)
		}
	}

	return nil
}
