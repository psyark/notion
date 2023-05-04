package objectdoc

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
)

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
func (t *specialBlockParser) CanInterruptParagraph() bool { return false }

// CanAcceptIndentedLine returns true if the parser can open new node when
// the given line is being indented more than 3 spaces.
func (t *specialBlockParser) CanAcceptIndentedLine() bool { return false }

// docRenderer は ドキュメント1ページのASTからトークンを取り出すだけの疑似レンダラーです
type docRenderer struct {
	elements []docElement
}

func (r *docRenderer) Render(w io.Writer, source []byte, n ast.Node) error {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		text := ""
		if c.Kind() == ast.KindFencedCodeBlock {
			for i := 0; i < c.Lines().Len(); i++ {
				line := c.Lines().At(i)
				text += string(line.Value(source))
			}
		} else {
			text = string(c.Text(source))
		}
		if c.Kind() == KindSpecialBlock {
			switch {
			case strings.HasPrefix(text, "[block:parameters]"):
				text = strings.TrimPrefix(text, "[block:parameters]")
				text = strings.TrimSuffix(text, "[/block]")

				table := struct {
					Data map[string]string `json:"data"`
					Rows int               `json:"rows"`
					Cols int               `json:"cols"`
				}{}
				if err := json.Unmarshal([]byte(text), &table); err != nil {
					return err
				}
				for row := 0; row < table.Rows; row++ {
					param := parameterElement{}
					for col := 0; col < table.Cols; col++ {
						h := table.Data[fmt.Sprintf("h-%d", col)]
						v := table.Data[fmt.Sprintf("%d-%d", row, col)]

						switch h {
						case "Property", "Field":
							param.Property = v
						case "Type":
							param.Type = v
						case "Description":
							param.Description = v
						case "Example value", "Example values":
							param.ExampleValue = v
						default:
							panic(h)
						}
					}
					r.elements = append(r.elements, &param)
				}
			default:
				return fmt.Errorf("%v", text)
			}
		} else {
			r.elements = append(r.elements, &blockElement{Kind: c.Kind().String(), Text: text})
		}
	}
	return nil
}
func (r *docRenderer) AddOptions(...renderer.Option) {}
