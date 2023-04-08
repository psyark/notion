package methoddoc

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	localCopyOfBodyParams []ssrPropsParam // ボディパラメータのローカルコピー
	returnType            MethodCoderType // リターンタイプ
}

func (c *converter) convert(file *jen.File) error {
	fmt.Println(c.url)

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

	// TODO
	// ローカルコピーとの比較
	remoteMap := map[string]ssrPropsParam{}
	for _, p := range sp.Doc.API.Params {
		if p.In == "body" {
			remoteMap[p.Name] = p
		}
	}
	if len(remoteMap) != len(c.localCopyOfBodyParams) {
		return fmt.Errorf("localCopyOfBodyParamsの数(%d)がリモート(%d)と一致しません", len(c.localCopyOfBodyParams), len(remoteMap))
	}

	c.output(file, sp.Doc)

	return nil
}

func (c *converter) output(file *jen.File, doc ssrPropsDoc) {
	code := jen.Comment(doc.Title).Line().Comment(c.url).Line()
	hasBodyParams := len(doc.API.filterParams("body")) != 0

	methodName := strcase.UpperCamelCase(strings.ReplaceAll(doc.Title, " a ", " "))
	methodParams := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}

	// APIのパスパラメータを引数化
	for _, param := range doc.API.filterParams("path") {
		if param.Type != "string" {
			panic(param.Type)
		}

		methodParams = append(methodParams, jen.Id(param.Name).String())
	}
	// オプション構造体引数
	if hasBodyParams {
		methodParams = append(methodParams, jen.Id("params").Op("*").Id(methodName+"Params"))
	}
	methodParams = append(methodParams, jen.Id("options").Op("...").Id("callOption"))

	pathParams := []jen.Code{}
	path := regexp.MustCompile(`\{\w+\}`).ReplaceAllStringFunc(doc.API.URL, func(s string) string {
		pathParams = append(pathParams, jen.Id(s[1:len(s)-1]))
		return "%v"
	})
	pathParams = append([]jen.Code{jen.Lit(path)}, pathParams...)

	params := jen.Nil()
	if hasBodyParams {
		params = jen.Id("params")
	}

	code.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(methodName).Params(methodParams...).Params(c.returnType.Returns(), jen.Error()).Block(
		jen.Id("result").Op(":=").Add(c.returnType.New()).Values(),
		jen.Id("co").Op(":=").Op("&").Id("callOptions").Values(jen.Dict{
			jen.Id("method"): jen.Qual("net/http", "Method"+strcase.UpperCamelCase(doc.API.Method)),
			jen.Id("path"):   jen.Qual("fmt", "Sprintf").Call(pathParams...),
			jen.Id("params"): params,
			jen.Id("result"): jen.Id("result"),
		}),
		jen.For(jen.List(jen.Id("_"), jen.Id("o")).Op(":=").Range().Id("options")).Block(
			jen.Id("o").Call(jen.Id("co")),
		),
		jen.List(jen.Return().Add(c.returnType.Access("result")), jen.Id("c").Dot("call").Call(
			jen.Id("ctx"),
			jen.Id("co"),
		)),
	).Line()

	if hasBodyParams {
		fields := []jen.Code{}
		for _, param := range c.localCopyOfBodyParams {
			fields = append(fields, jen.Id(param.Name).Add(param.typeCode).Tag(map[string]string{"json": param.Name}).Comment(param.Desc))
		}
		code.Type().Id(methodName + "Params").Struct(fields...).Line()
	}

	file.Add(*code...)
}

var registeredConverters []converter

// registerConverter は後で実行するためにconverterを登録します
func registerConverter(c converter) {
	registeredConverters = append(registeredConverters, c)
}

// convertAll は登録された全てのconverterで変換を実行します
func convertAll() error {
	file := jen.NewFile("notion")

	eg := errgroup.Group{}
	for _, c := range registeredConverters {
		c := c
		eg.Go(func() error {
			if err := c.convert(file); err != nil {
				return fmt.Errorf("convert: %s: %w", c.url, err)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	return file.Save("../../client.methods.go")
}
