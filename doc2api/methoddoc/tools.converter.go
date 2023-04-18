package methoddoc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
	"golang.org/x/sync/errgroup"
)

// converter はNotion API ReferenceからGoコードへの変換ルールです。
type converter struct {
	url                   string          // ドキュメントのURL
	localCopyOfPathParams []ssrPropsParam // pathパラメータのローカルコピー
	localCopyOfBodyParams []ssrPropsParam // bodyパラメータのローカルコピー
	returnType            returnTypeCoder // リターンタイプ
}

func (c *converter) convert() error {
	file := jen.NewFile("notion")

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
	sp := ssrProps{}
	if err := json.Unmarshal(ssrPropsBytes, &sp); err != nil {
		return err
	}

	fileName := "tmp/" + strings.TrimPrefix(c.url, "https://developers.notion.com/reference/") + ".go"
	os.Remove(fileName)

	updateTmpFile := func() {
		if len(c.localCopyOfPathParams) == 0 || len(c.localCopyOfBodyParams) == 0 {
			gostr := jen.NewFile("tmp")
			if len(c.localCopyOfPathParams) == 0 {
				gostr.Var().Id("PATH_PARAMS").Op("=").Add(sp.Doc.API.Params.filter("path").code())
			}
			if len(c.localCopyOfBodyParams) == 0 {
				gostr.Var().Id("BODY_PARAMS").Op("=").Add(sp.Doc.API.Params.filter("body").code())
			}
			if err := gostr.Save(fileName); err != nil {
				panic(err)
			}
		}
	}

	// ローカルコピーとの比較
	if err := sp.Doc.API.Params.filter("path").compare(c.localCopyOfPathParams); err != nil {
		updateTmpFile()
		return fmt.Errorf("path params mismatch: %w", err)
	}
	if err := sp.Doc.API.Params.filter("body").compare(c.localCopyOfBodyParams); err != nil {
		updateTmpFile()
		return fmt.Errorf("body params mismatch: %w", err)
	}

	if err := c.output(file, sp.Doc); err != nil {
		return err
	}

	return file.Save(fmt.Sprintf("../../method.%s.go", strcase.SnakeCase(strings.ReplaceAll(sp.Doc.Title, " a ", " "))))
}

func (c *converter) output(file *jen.File, doc ssrPropsDoc) error {
	file.Comment(doc.Title).Line().Comment(c.url)
	hasBodyParams := len(c.localCopyOfBodyParams) != 0

	methodName := strcase.UpperCamelCase(strings.ReplaceAll(doc.Title, " a ", " "))
	methodParams := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}

	// APIのパスパラメータを引数化
	for _, param := range c.localCopyOfPathParams {
		if param.typeCode == nil {
			return fmt.Errorf("no typeCode for %s", param.Name)
		}

		methodParams = append(methodParams, jen.Id(param.Name).Add(param.typeCode))
	}

	// オプション構造体引数
	params := jen.Nil()
	if hasBodyParams {
		params = jen.Id("params")
		methodParams = append(methodParams, jen.Id("params").Op("*").Id(methodName+"Params"))
	}

	methodParams = append(methodParams, jen.Id("options").Op("...").Id("callOption"))

	file.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(methodName).Params(methodParams...).Params(c.returnType.returnType(), jen.Error()).Block(
		jen.Id("result").Op(":=").Add(c.returnType.unmarshaller()).Values(),
		jen.Id("co").Op(":=").Op("&").Id("callOptions").Values(jen.Dict{
			jen.Id("method"): jen.Qual("net/http", "Method"+strcase.UpperCamelCase(doc.API.Method)),
			jen.Id("path"):   c.pathCode(doc),
			jen.Id("params"): params,
			jen.Id("result"): jen.Id("result"),
		}),
		jen.For(jen.List(jen.Id("_"), jen.Id("o")).Op(":=").Range().Id("options")).Block(
			jen.Id("o").Call(jen.Id("co")),
		),
		jen.List(jen.Return().Add(c.returnType.returnValue("result")), jen.Id("c").Dot("call").Call(
			jen.Id("ctx"),
			jen.Id("co"),
		)),
	).Line()

	if hasBodyParams {
		fields := []jen.Code{}
		for _, param := range c.localCopyOfBodyParams {
			if param.typeCode == nil {
				return fmt.Errorf("no typeCode for %s", param.Name)
			}
			jsonTag := param.Name
			if param.omitEmpty {
				jsonTag += ",omitempty"
			}
			fields = append(fields, jen.Id(strcase.UpperCamelCase(param.Name)).Add(param.typeCode).Tag(map[string]string{"json": jsonTag}).Comment(param.Desc))
		}
		file.Type().Id(methodName + "Params").Struct(fields...).Line()
	}

	return nil
}

func (c *converter) pathCode(doc ssrPropsDoc) jen.Code {
	pathParams := []jen.Code{}
	pathFormat := regexp.MustCompile(`\{\w+\}`).ReplaceAllStringFunc(doc.API.URL, func(s string) string {
		pathParams = append(pathParams, jen.Id(s[1:len(s)-1]))
		return "%v"
	})

	// パスパラメータが無いのでURLをそのままリテラルにする
	if len(pathParams) == 0 {
		return jen.Lit(doc.API.URL)
	}

	// パラメータの先頭にフォーマットを追加
	pathParams = append([]jen.Code{jen.Lit(pathFormat)}, pathParams...)
	return jen.Qual("fmt", "Sprintf").Call(pathParams...)
}

var registeredConverters []converter

// registerConverter は後で実行するためにconverterを登録します
func registerConverter(c converter) {
	registeredConverters = append(registeredConverters, c)
}

// convertAll は登録された全てのconverterで変換を実行します
func convertAll() error {
	eg := errgroup.Group{}
	for _, c := range registeredConverters {
		c := c
		eg.Go(func() error {
			if err := c.convert(); err != nil {
				return fmt.Errorf("convert: %s: %w", c.url, err)
			}
			return nil
		})
	}
	return eg.Wait()
}
