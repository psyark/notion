package objectdoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
)

var headingRegex = regexp.MustCompile("^#+")

type objectDocTokenizer struct {
	lines []string
	index int
}

func (t *objectDocTokenizer) next() (objectDocElement, error) {
	if t.index >= len(t.lines) {
		return nil, io.EOF
	}

	if strings.HasPrefix(t.lines[t.index], "[block:") {
		var block objectDocElement
		switch t.lines[t.index] {
		case "[block:code]":
			block = &objectDocCodeElement{}
		case "[block:callout]":
			block = &objectDocCalloutElement{}
		case "[block:parameters]":
			block = &objectDocParametersElement{}
		case "[block:api-header]":
			block = &objectDocAPIHeaderElement{}
		case "[block:image]":
			block = &objectDocImageElement{}
		default:
			return nil, fmt.Errorf("unknown block: %v", t.lines[t.index])
		}

		startIndex := t.index
		for t.index < len(t.lines) {
			if t.lines[t.index] == "[/block]" {
				t.index++

				content := []byte(strings.Join(t.lines[startIndex+1:t.index-1], "\n"))
				if err := json.Unmarshal(content, block); err != nil {
					return nil, err
				}
				return block, nil
			}
			t.index++
		}
		return nil, fmt.Errorf("[/block] not exists")
	} else if m := headingRegex.FindString(t.lines[t.index]); m != "" {
		token := &objectDocHeadingElement{Text: stripMarkdown(t.lines[t.index])}
		t.index++
		return token, nil
	} else if strings.HasPrefix(t.lines[t.index], "|") {
		startIndex := t.index
		for t.index < len(t.lines) && strings.HasPrefix(t.lines[t.index], "|") {
			t.index++
		}
		table := strings.Join(t.lines[startIndex:t.index], "\n")

		buf := bytes.NewBuffer(nil)
		goldmark.New(
			goldmark.WithExtensions(extension.Table),
			goldmark.WithRenderer(&tableRenderer{}),
		).Convert([]byte(table), buf)

		elem := objectDocParametersElement{}
		json.Unmarshal(buf.Bytes(), &elem)

		return &elem, nil
	} else if strings.HasPrefix(t.lines[t.index], "```") {
		startIndex := t.index
		for t.index < len(t.lines) {
			t.index++
			if t.lines[t.index] == "```" {
				t.index++
				break
			}
		}

		langName := strings.SplitN(strings.TrimPrefix(t.lines[startIndex], "```"), " ", 2)
		elem := &objectDocCodeElement{
			Codes: []*objectDocCodeElementCode{{
				Language: langName[0],
				Name:     langName[1],
				Code:     strings.Join(t.lines[startIndex+1:t.index-1], "\n"),
			}},
		}

		return elem, nil
	} else {
		startIndex := t.index
		for t.index < len(t.lines) && t.isParagraph(t.lines[t.index]) {
			t.index++
		}
		token := &objectDocParagraphElement{Text: stripMarkdown(strings.Join(t.lines[startIndex:t.index], "\n"))}
		if strings.TrimSpace(token.Text) != "" {
			return token, nil
		}
		return t.next()
	}
}

type tableRenderer struct {
	raw objectDocParametersElementRaw
}

func (r *tableRenderer) AddOptions(...renderer.Option) {}
func (r *tableRenderer) Render(w io.Writer, source []byte, n ast.Node) error {
	switch n.Kind().String() {
	case "TableHeader":
		r.raw.Rows = 0
		r.raw.Cols = 0
	case "TableRow":
		r.raw.Rows++
		r.raw.Cols = 0
	case "TableCell":
		r.raw.Cols++
	case "Text":
		if r.raw.Data == nil {
			r.raw.Data = map[string]string{}
		}
		if r.raw.Rows == 0 {
			r.raw.Data[fmt.Sprintf("h-%d", r.raw.Cols-1)] += string(n.Text(source))
		} else {
			r.raw.Data[fmt.Sprintf("%d-%d", r.raw.Rows-1, r.raw.Cols-1)] += string(n.Text(source))
		}
	}

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		r.Render(w, source, c)
	}

	if n.Type() == ast.TypeDocument {
		return json.NewEncoder(w).Encode(r.raw)
	}

	return nil
}

func (t *objectDocTokenizer) isParagraph(line string) bool {
	if strings.HasPrefix(line, "[block:") {
		return false // block
	}
	if strings.HasPrefix(line, "#") {
		return false // heading
	}
	if strings.HasPrefix(line, "```") {
		return false // code
	}
	if strings.HasPrefix(line, "|") {
		return false // table
	}
	return true // paragraph
}
