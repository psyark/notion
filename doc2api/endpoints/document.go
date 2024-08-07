package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
	"github.com/samber/lo"
	"github.com/stoewer/go-strcase"
)

type endpointDocument struct {
	url      string
	ssrProps ssrProps
}

// ParamAnnotations はパラメータの型アノテーションです
type ParamAnnotations map[string]jen.Code

// Generate は、このドキュメントからコードを出力します
func (r endpointDocument) Generate(returnType ReturnType, paramAnnots ParamAnnotations) {
	file := jen.NewFile("notion")

	hasBodyParams := lo.SomeBy(r.ssrProps.Doc.API.Params, func(p ssrPropsParam) bool { return p.In == "body" })

	// ヘッダーコメント
	file.HeaderComment("Code generated by notion.doc2api; DO NOT EDIT.")
	file.HeaderComment(r.url)

	// メソッド本体の出力
	file.Comment(r.ssrProps.Doc.Body)
	file.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(r.methodName()).ParamsFunc(func(g *jen.Group) {
		g.Id("ctx").Qual("context", "Context")
		for _, param := range r.ssrProps.Doc.API.Params {
			if param.In == "path" {
				annot, ok := paramAnnots[param.Name]
				if !ok {
					panic(fmt.Errorf("パラメータ %q のアノテーションが存在しません", param.Name))
				}
				g.Id(param.Name).Add(annot)
			}
		}
		if hasBodyParams {
			g.Id("params").Id(r.paramsName())
		}

		g.Id("options").Op("...").Id("CallOption")
	}).Params(returnType.Type(), jen.Error()).BlockFunc(func(g *jen.Group) {
		g.Return().Id("call").Call(
			jen.Line().Id("ctx"),
			jen.Line().Id("c").Dot("accessToken"),
			jen.Line().Add(jen.Qual("net/http", fmt.Sprintf("Method%s", strcase.UpperCamelCase(r.ssrProps.Doc.API.Method)))),
			jen.Line().Add(r.pathCode()),
			jen.Line().Add(lo.Ternary(hasBodyParams, jen.Id("params"), jen.Nil())),
			jen.Line().Add(returnType.Accessor()),
			jen.Line().Id("options").Op("..."),
			jen.Line(),
		)
	})

	// パラメータの出力
	if hasBodyParams {
		file.Type().Id(r.paramsName()).Map(jen.String()).Any()

		for _, param := range r.ssrProps.Doc.API.Params {
			if param.In == "body" {
				annot, ok := paramAnnots[param.Name]
				if !ok {
					panic(fmt.Errorf("パラメータ %q のアノテーションが存在しません", param.Name))
				}

				publicName := strcase.UpperCamelCase(param.Name)
				file.Comment(param.Desc)
				file.Func().Params(jen.Id("p").Id(r.paramsName())).Id(publicName).Params(jen.Id(param.Name).Add(annot)).Id(r.paramsName()).Block(
					jen.Id("p").Index(jen.Lit(param.Name)).Op("=").Id(param.Name),
					jen.Return().Id("p"),
				)
			}
		}
	}

	lo.Must0(file.Save(r.fileName()))
}

func (r endpointDocument) baseName() string {
	return strings.ReplaceAll(r.ssrProps.Doc.Title, " a ", " ")
}
func (r endpointDocument) fileName() string {
	return fmt.Sprintf("../../endpoints_%s_generated.go", strcase.SnakeCase(r.baseName()))
}
func (r endpointDocument) methodName() string {
	return strcase.UpperCamelCase(r.baseName())
}
func (r endpointDocument) paramsName() string {
	return r.methodName() + "Params"
}
func (r *endpointDocument) pathCode() jen.Code {
	pathParams := []jen.Code{}
	pathFormat := regexp.MustCompile(`\{\w+\}`).ReplaceAllStringFunc(r.ssrProps.Doc.API.URL, func(s string) string {
		pathParams = append(pathParams, jen.Id(s[1:len(s)-1]))
		return "%v"
	})

	// パスパラメータが無いのでURLをそのままリテラルにする
	if len(pathParams) == 0 {
		return jen.Lit(r.ssrProps.Doc.API.URL)
	}

	// パラメータの先頭にフォーマットを追加
	return jen.Qual("fmt", "Sprintf").CallFunc(func(g *jen.Group) {
		g.Lit(pathFormat)
		for _, p := range pathParams {
			g.Add(p)
		}
	})
}

// Fetch は エンドポイントのドキュメントを取得します
func Fetch(url string) endpointDocument {
	// URLの取得
	res := lo.Must(http.Get(url))
	defer res.Body.Close()

	// goqueryでのパース
	doc := lo.Must(goquery.NewDocumentFromReader(res.Body))

	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	props := ssrProps{}

	d := json.NewDecoder(bytes.NewReader(ssrPropsBytes))
	d.DisallowUnknownFields()
	lo.Must0(d.Decode(&props))

	return endpointDocument{url: url, ssrProps: props}
}
