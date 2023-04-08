package objectdoc

import (
	"encoding/json"
	"fmt"
)

var errUnmatch = fmt.Errorf("unmatch")

// objectDocElement はNotion API Reference のObjects以下 (e.g. https://developers.notion.com/reference/block) で
// 出現するエレメントのインターフェイスです
//
// このインターフェイスは以下の2通りの使われ方をします
// 1. Notion API Referenceの最新のデータ（リモート。変換ロジックは常にnilです）
// 2. ある時点の(1)がGoコードとして保存されたもの（ローカルコピー。多くの場合変換ロジックを伴います）
type objectDocElement interface {
	// checkAndOutput は2つのことを行います
	// 1. このエレメント（ローカルコピー）と渡されたエレメント（最新のリモート）を比べ、一致しなければ errUnmatch を返す
	// 2. このエレメント（ローカルコピー）に変換ロジックが設定されていればそれを実行し、エラーを返す
	checkAndOutput(remote objectDocElement, b *builder) error
}

var _ = []objectDocElement{
	&objectDocHeadingElement{},
	&objectDocParagraphElement{},
	&objectDocCodeElement{},
	&objectDocCalloutElement{},
	&objectDocParametersElement{},
	&objectDocAPIHeaderElement{},
}

// objectDocHeadingElement は見出しを表すobjectDocElementです
type objectDocHeadingElement struct {
	Text   string
	output func(*objectDocHeadingElement, *builder) error
}

func (e *objectDocHeadingElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocHeadingElement); !ok {
		return errUnmatch
	} else if e.Text != remote.Text {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e, b)
	} else {
		return nil
	}
}

// objectDocParagraphElement は段落を表すobjectDocElementです
type objectDocParagraphElement struct {
	Text   string
	output func(*objectDocParagraphElement, *builder) error
}

func (e *objectDocParagraphElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocParagraphElement); !ok {
		return errUnmatch
	} else if e.Text != remote.Text {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e, b)
	} else {
		return nil
	}
}

// objectDocAPIHeaderElement はAPI Headerを表すobjectDocElementです
type objectDocAPIHeaderElement struct {
	Title  string `json:"title"`
	output func(*objectDocAPIHeaderElement, *builder) error
}

func (e *objectDocAPIHeaderElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocAPIHeaderElement); !ok {
		return errUnmatch
	} else if e.Title != remote.Title {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e, b)
	} else {
		return nil
	}
}

// objectDocCodeElement はコードを表すobjectDocElementです
type objectDocCodeElement struct {
	Codes []*objectDocCodeElementCode `json:"codes"`
}

func (e *objectDocCodeElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocCodeElement); !ok {
		return errUnmatch
	} else if len(e.Codes) != len(remote.Codes) {
		return errUnmatch
	} else {
		for i, code := range e.Codes {
			rc := remote.Codes[i]
			if code.Name != rc.Name || code.Language != rc.Language || code.Code != rc.Code {
				return errUnmatch
			} else if code.output != nil {
				if err := code.output(code, b); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

type objectDocCodeElementCode struct {
	Name     string `json:"string"`
	Language string `json:"language"`
	Code     string `json:"code"`
	output   func(*objectDocCodeElementCode, *builder) error
}

// objectDocCodeElement はコールアウトを表すobjectDocElementです
type objectDocCalloutElement struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	output func(*objectDocCalloutElement, *builder) error
}

func (e *objectDocCalloutElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocCalloutElement); !ok {
		return errUnmatch
	} else if e.Type != remote.Type || e.Title != remote.Title || e.Body != remote.Body {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e, b)
	} else {
		return nil
	}
}

type objectDocParametersElement []*objectDocParameter

func (e *objectDocParametersElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocParametersElement); !ok {
		return errUnmatch
	} else if len(*e) != len(*remote) {
		return errUnmatch
	} else {
		for i, param := range *e {
			if err := param.checkAndOutput((*remote)[i], b); err != nil {
				return err
			}
		}
		return nil
	}
}

func (t *objectDocParametersElement) UnmarshalJSON(data []byte) error {
	raw := struct {
		Data map[string]string `json:"data"`
		Cols int               `json:"cols"`
		Rows int               `json:"rows"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mapSlice := make([]map[string]string, raw.Rows)
	for r := range mapSlice {
		m := map[string]string{}
		for c := 0; c < raw.Cols; c++ {
			m[raw.Data[fmt.Sprintf("h-%d", c)]] = stripMarkdown(raw.Data[fmt.Sprintf("%d-%d", r, c)])
		}
		mapSlice[r] = m
	}

	data, err := json.Marshal(mapSlice)
	if err != nil {
		return err
	}

	type Alias objectDocParametersElement
	return json.Unmarshal(data, (*Alias)(t))
}

// objectDocParameter はドキュメントに書かれているパラメータです
type objectDocParameter struct {
	Property      string `json:",omitempty"`
	Field         string `json:",omitempty"`
	Type          string `json:",omitempty"`
	Description   string `json:",omitempty"`
	ExampleValue  string `json:"Example value,omitempty"`
	ExampleValues string `json:"Example values,omitempty"`
	output        func(*objectDocParameter, *builder) error
}

func (e *objectDocParameter) checkAndOutput(remote *objectDocParameter, b *builder) error {
	if e.Property != remote.Property || e.Field != remote.Field || e.Type != remote.Type || e.Description != remote.Description || e.ExampleValue != remote.ExampleValue || e.ExampleValues != remote.ExampleValues {
		return errUnmatch
	} else if e.output != nil {
		if err := e.output(e, b); err != nil {
			return err
		}
	}
	return nil
}
