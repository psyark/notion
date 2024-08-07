package objects

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/yuin/goldmark/ast"
	xast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
)

// オブジェクトのドキュメントの内容はMarkdownとして取得できます。
// ここではそのMarkdownからデータ構造を取り出すための手段を提供します。

var (
	KindSpecialBlock = ast.NewNodeKind("SpecialBlock")

	_ ast.Node           = &specialBlock{}
	_ parser.BlockParser = &specialBlockParser{}
	_ renderer.Renderer  = &docRenderer{}
)

// specialBlock は [block][/block] 記法です
type specialBlock struct{ ast.BaseBlock }

// Kind returns a kind of this node.
func (b *specialBlock) Kind() ast.NodeKind { return KindSpecialBlock }

// Dump dumps an AST tree structure to stdout.
// This function completely aimed for debugging.
// level is a indent level. Implementer should indent informations with
// 2 * level spaces.
func (b *specialBlock) Dump(source []byte, level int) {}

// specialBlockParser は [block][/block] 記法のパーサです
type specialBlockParser struct{}

// Trigger returns a list of characters that triggers Parse method of
// this parser.
// If Trigger returns a nil, Open will be called with any lines.
func (t *specialBlockParser) Trigger() []byte { return []byte("[") }

// Open parses the current line and returns a result of parsing.
//
// Open must not parse beyond the current line.
// If Open has been able to parse the current line, Open must advance a reader
// position by consumed byte length.
//
// If Open has not been able to parse the current line, Open should returns
// (nil, NoChildren). If Open has been able to parse the current line, Open
// should returns a new Block node and returns HasChildren or NoChildren.
func (t *specialBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, seg := reader.PeekLine()
	if strings.HasPrefix(string(line), "[block:") {
		b := &specialBlock{}
		b.Lines().Append(seg)
		return b, parser.NoChildren
	}
	return nil, parser.NoChildren
}

// Continue parses the current line and returns a result of parsing.
//
// Continue must not parse beyond the current line.
// If Continue has been able to parse the current line, Continue must advance
// a reader position by consumed byte length.
//
// If Continue has not been able to parse the current line, Continue should
// returns Close. If Continue has been able to parse the current line,
// Continue should returns (Continue | NoChildren) or
// (Continue | HasChildren)
func (t *specialBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, seg := reader.PeekLine()
	node.Lines().Append(seg)

	if string(line) == "[/block]\n" {
		reader.AdvanceLine()
		return parser.Close
	}

	return parser.Continue | parser.NoChildren
}

// Close will be called when the parser returns Close.
func (t *specialBlockParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {}

// CanInterruptParagraph returns true if the parser can interrupt paragraphs,
// otherwise false.
func (t *specialBlockParser) CanInterruptParagraph() bool { return true }

// CanAcceptIndentedLine returns true if the parser can open new node when
// the given line is being indented more than 3 spaces.
func (t *specialBlockParser) CanAcceptIndentedLine() bool { return false }

// docRenderer は ドキュメント1ページのASTからトークンを取り出すだけの疑似レンダラーです
type docRenderer struct {
	OnElement func(DocumentElement)
}

func (r *docRenderer) Render(_ io.Writer, source []byte, n ast.Node) error {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if err := r.renderToplevelChild(source, c); err != nil {
			return err
		}
	}
	return nil
}
func (r *docRenderer) renderToplevelChild(source []byte, n ast.Node) error {
	if n.Kind() == xast.KindTable {
		return r.renderMarkdownTable(source, n)
	}

	text := ""
	if n.Kind() == ast.KindFencedCodeBlock {
		for i := 0; i < n.Lines().Len(); i++ {
			line := n.Lines().At(i)
			text += string(line.Value(source))
		}
	} else {
		text = string(n.Text(source))
	}
	if n.Kind() == KindSpecialBlock {
		switch {
		case strings.HasPrefix(text, "[block:parameters]"):
			text = strings.TrimSuffix(strings.TrimPrefix(text, "[block:parameters]"), "[/block]")

			table := struct {
				Data map[string]string `json:"data"`
				Rows int               `json:"rows"`
				Cols int               `json:"cols"`
			}{}
			if err := json.Unmarshal([]byte(text), &table); err != nil {
				return err
			}
			for row := 0; row < table.Rows; row++ {
				param := Parameter{}
				for col := 0; col < table.Cols; col++ {
					h := table.Data[fmt.Sprintf("h-%d", col)]
					v := table.Data[fmt.Sprintf("%d-%d", row, col)]
					if err := param.setCell(h, v); err != nil {
						return err
					}
				}
				r.OnElement(&param)
			}
		case strings.HasPrefix(text, "[block:code]"):
			text = strings.TrimSuffix(strings.TrimPrefix(text, "[block:code]"), "[/block]")

			codes := struct {
				Codes []struct {
					Code     string `json:"code"`
					Language string `json:"language"`
					Name     string `json:"name"`
				} `json:"codes"`
			}{}
			if err := json.Unmarshal([]byte(text), &codes); err != nil {
				return err
			}
			for _, code := range codes.Codes {
				r.OnElement(&Block{Kind: ast.KindFencedCodeBlock.String(), Text: code.Code})
			}
		case strings.HasPrefix(text, "[block:callout]"):
			text = strings.TrimSuffix(strings.TrimPrefix(text, "[block:callout]"), "[/block]")
			callout := struct {
				Body  string `json:"body"`
				Type  string `json:"type"`
				Title string `json:"title"`
			}{}
			if err := json.Unmarshal([]byte(text), &callout); err != nil {
				return err
			}
			r.OnElement(&Block{Kind: ast.KindBlockquote.String(), Text: callout.Body})
		case strings.HasPrefix(text, "[block:image]"):
			// ignore
		default:
			return fmt.Errorf("unsupported block: %v", text)
		}
	} else {
		r.OnElement(&Block{Kind: n.Kind().String(), Text: text})
	}
	return nil
}
func (r *docRenderer) renderMarkdownTable(source []byte, n ast.Node) error {
	headers := []string{}
	for row := n.FirstChild(); row != nil; row = row.NextSibling() {
		switch row.Kind() {
		case xast.KindTableHeader:
			for col := row.FirstChild(); col != nil; col = col.NextSibling() {
				headers = append(headers, string(col.Text(source)))
			}
		case xast.KindTableRow:
			param := Parameter{}
			i := 0
			for col := row.FirstChild(); col != nil; col = col.NextSibling() {
				v := string(col.Text(source))
				if err := param.setCell(headers[i], v); err != nil {
					return err
				}
				i++
			}
			r.OnElement(&param)
		default:
			return fmt.Errorf("unsupported: %v", row.Kind())
		}
	}
	return nil
}
func (r *docRenderer) AddOptions(...renderer.Option) {}
