package objectdoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

var registeredTranslators []translator

type translator struct {
	url    string
	scopes []translationScope
}

func (t *translator) translate() error {
	// URLの取得
	res, err := http.Get(t.url)
	if err != nil {
		return errors.Wrap(err, "fetch")
	}

	defer func() {
		_, _ = io.ReadAll(res.Body)
		_ = res.Body.Close()
	}()

	// goqueryでのパース
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return errors.Wrap(err, "parseHTML")
	}

	// JSONパース
	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	ssrProps := struct {
		Doc struct {
			Body string `json:"body"`
		} `json:"doc"`
	}{}
	if err := json.Unmarshal(ssrPropsBytes, &ssrProps); err != nil {
		return errors.Wrap(err, "unmarshalSSR")
	}

	// エレメント取り出し
	ren := &docRenderer{}
	md := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(parser.WithBlockParsers(util.Prioritized(&specialBlockParser{}, 0))),
		goldmark.WithRenderer(ren),
	)
	if err := md.Convert([]byte(ssrProps.Doc.Body), bytes.NewBuffer(nil)); err != nil {
		return errors.Wrap(err, "markdown")
	}

	// スコープ実行
	c := &comparator{elements: ren.elements}
	fileName := "object." + strings.TrimPrefix(t.url, "https://developers.notion.com/reference/") + ".go"
	b := &builder{
		global:        global,
		globalSymbols: global.globalSymbols,
		fileName:      "../../" + fileName,
		url:           t.url,
	}

	// スコープの実行とパニック時の脱出
	var scopeErr error
	for i, scope := range t.scopes {
		func() {
			defer func() {
				if err := recover(); err != nil {
					scopeErr = fmt.Errorf("scope[%d]: %s", i, err)
				}
			}()
			scope(c, b)
		}()
		if scopeErr != nil {
			return scopeErr
		}
	}

	// ドキュメントが不足していた場合の仮コード
	if c.index < len(c.elements) {
		groups := [][]docElement{}
		for _, elem := range c.elements[c.index:] {
			if be, ok := elem.(*blockElement); (ok && be.Kind == "Heading") || len(groups) == 0 {
				groups = append(groups, []docElement{})
			}
			g := &groups[len(groups)-1]
			*g = append(*g, elem)
		}

		f := jen.NewFile("template")
		f.Var().Id("ADD_THIS").Op("=").Index().Id("translationScope").ValuesFunc(func(g *jen.Group) {
			for _, group := range groups {
				comment := ""
				if be, ok := group[0].(*blockElement); ok && be.Kind == "Heading" {
					comment = be.Text
				}
				g.Line().Func().Params(jen.Id("c").Op("*").Id("comparator"), jen.Id("b").Op("*").Id("builder")).Comment(fmt.Sprintf("/* %s */", comment)).BlockFunc(func(g *jen.Group) {
					for _, elem := range group {
						g.Add(elem.template())
					}
				})
			}
			g.Line()
		})
		if err := f.Save("tmp/" + fileName); err != nil {
			return err
		}
		return fmt.Errorf("%d element(s) remains", len(c.elements)-c.index)
	}

	// ビルダー本体出力
	if len(b.localSymbols) != 0 {
		file := jen.NewFile("notion")
		file.Comment("Code generated by notion.doc2api; DO NOT EDIT.")
		if b.url != "" {
			file.Comment(b.url)
		}
		for _, s := range b.localSymbols {
			file.Line().Line().Add(s.symbolCode(b))
		}
		if err := file.Save(b.fileName); err != nil {
			return err
		}
	} else {
		_ = os.Remove(b.fileName)
	}

	// テストコード出力
	if len(b.testSymbols) != 0 {
		file := jen.NewFile("notion")
		file.Comment("Code generated by notion.doc2api; DO NOT EDIT.")
		if b.url != "" {
			file.Comment(b.url)
		}
		for _, s := range b.testSymbols {
			file.Line().Line().Add(s.symbolCode(b))
		}
		if err := file.Save(strings.Replace(b.fileName, ".go", "_test.go", 1)); err != nil {
			return err
		}
	} else {
		_ = os.Remove(strings.Replace(b.fileName, ".go", "_test.go", 1))
	}

	return nil // OK
}

type translationScope func(c *comparator, b *builder)

func registerTranslator(url string, scopes ...translationScope) {
	registeredTranslators = append(registeredTranslators, translator{url: url, scopes: scopes})
}

func translateAll() error {
	for _, t := range registeredTranslators {
		if err := t.translate(); err != nil {
			return errors.Wrap(err, t.url)
		}
	}
	return nil
}
