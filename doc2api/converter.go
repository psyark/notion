// Package doc2api は Notion API Reference の更新を検知し、Goコードへの適切な変換を
// 継続的に行うための一連の仕組みを提供します。
//
// Goコードへの変換ルールは、命令形のコードではなくデータ構造として objects.*.go に格納されます。
// このデータ構造には Notion API Referenceのローカルコピーも含まれるため、
// ドキュメントの更新に対してGoコードへの変換ルールが古いままになることを防ぎます。
package doc2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
)

// converter はNotion API ReferenceからGoコードへの変換ルールです。
type converter struct {
	url       string // ドキュメントのURL
	fileName  string // 出力するファイル名
	localCopy []objectDocElement
}

// convert は変換を実行します
func (c converter) convert() error {
	// URLの取得
	res, err := http.Get(c.url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// goqueryでのパース
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	ssrProps := struct {
		Doc struct {
			Body string `json:"body"`
		} `json:"doc"`
	}{}
	if err := json.Unmarshal(ssrPropsBytes, &ssrProps); err != nil {
		return err
	}

	fmt.Println(c.fileName, c.url)

	lines := strings.Split(ssrProps.Doc.Body, "\n")
	odt := &objectDocTokenizer{lines, 0}

	requiredCopy := jen.Statement{}

	for i := 0; ; i++ {
		remote, err := odt.next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}

		if len(c.localCopy) < i+1 {
			requiredCopy = append(requiredCopy, createLocalCopy(remote))
		} else if err := c.localCopy[i].checkAndOutput(remote); err != nil {
			if errors.Is(err, errUnmatch) {
				ld, _ := json.MarshalIndent(c.localCopy[i], "", "  ")
				rd, _ := json.MarshalIndent(remote, "", "  ")
				return fmt.Errorf("%w\nlocal:\n%#v\nremote:\n%#v", err, string(ld), string(rd))
			} else {
				return err
			}
		}
	}

	if len(requiredCopy) != 0 {
		gostr := jen.Var().Id("LOCAL_COPY").Op("=").Index().Id("objectDocElement").Values(jen.List(requiredCopy...), jen.Line()).GoString()
		if err := os.WriteFile("required_copy.txt", []byte(gostr), 0666); err != nil {
			return err
		}
		return fmt.Errorf("localCopyが足りません (see required_copy.txt)")
	}
	return nil
}

func createLocalCopy(remote objectDocElement) jen.Code {
	typeName := strings.TrimPrefix(fmt.Sprintf("%T", remote), "*doc2api.")
	outputCode := jen.Func().Params(jen.Id("e").Op("*").Id(typeName)).Error().Block(jen.Return().Nil().Comment("TODO"))

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
		outputCode = jen.Func().Params(jen.Id("e").Op("*").Id("objectDocCodeElementCode")).Error().Block(jen.Return().Nil().Comment("TODO"))
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
		outputCode = jen.Func().Params(jen.Id("e").Op("*").Id("objectDocParameter")).Error().Block(jen.Return().Nil().Comment("TODO"))
		params := jen.Statement{}
		for _, p := range *remote {
			params = append(params, jen.Values(jen.Dict{
				jen.Id("Property"):     jen.Lit(p.Property),
				jen.Id("Field"):        jen.Lit(p.Field),
				jen.Id("Type"):         jen.Lit(p.Type),
				jen.Id("Description"):  jen.Lit(p.Description),
				jen.Id("ExampleValue"): jen.Lit(p.ExampleValue),
				jen.Id("output"):       outputCode,
			}))
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
	for _, c := range registeredConverters {
		if err := c.convert(); err != nil {
			return fmt.Errorf("convert: %s: %w", c.fileName, err)
		}
	}
	return nil
}
