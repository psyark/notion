package objectdoc

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
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
	} else {
		startIndex := t.index
		for t.index < len(t.lines) && t.isParagraph(t.lines[t.index]) {
			t.index++
		}
		token := &objectDocParagraphElement{Text: stripMarkdown(strings.Join(t.lines[startIndex:t.index], "\n"))}
		if token.Text != "" {
			return token, nil
		}
		return t.next()
	}
}

func (t *objectDocTokenizer) isParagraph(line string) bool {
	return !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "[block:")
}
