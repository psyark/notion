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
	symbols       *SyncMap[string, CodeSymbol]
	globalBuilder *CodeBuilder
	builders      []*CodeBuilder
}

func NewConverter() *Converter {
	global := &CodeBuilder{fileName: "objects_global_generated.go"}
	return &Converter{
		symbols:       &SyncMap[string, CodeSymbol]{},
		globalBuilder: global,
		builders:      []*CodeBuilder{global},
	}
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

	builder := &CodeBuilder{
		url:       url,
		fileName:  fmt.Sprintf("objects_%s_generated.go", strings.TrimPrefix(url, "https://developers.notion.com/reference/")),
		converter: c,
	}

	c.builders = append(c.builders, builder)

	return &DocumentComparator{elements: elements, builder: builder}
}

func (c *Converter) OutputAllBuilders() {
	for _, b := range c.builders {
		b.output()
	}
}
