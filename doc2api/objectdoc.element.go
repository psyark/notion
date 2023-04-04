package doc2api

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
	checkAndOutput(remote objectDocElement) error
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
	output func(*objectDocHeadingElement) error
}

func (e *objectDocHeadingElement) checkAndOutput(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocHeadingElement); !ok {
		return errUnmatch
	} else if e.Text != remote.Text {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e)
	} else {
		return nil
	}
}

// objectDocParagraphElement は段落を表すobjectDocElementです
type objectDocParagraphElement struct {
	Text   string
	output func(*objectDocParagraphElement) error
}

func (e *objectDocParagraphElement) checkAndOutput(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocParagraphElement); !ok {
		return errUnmatch
	} else if e.Text != remote.Text {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e)
	} else {
		return nil
	}
}

// objectDocAPIHeaderElement はAPI Headerを表すobjectDocElementです
type objectDocAPIHeaderElement struct {
	Title  string `json:"title"`
	output func(*objectDocAPIHeaderElement) error
}

func (e *objectDocAPIHeaderElement) checkAndOutput(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocAPIHeaderElement); !ok {
		return errUnmatch
	} else if e.Title != remote.Title {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e)
	} else {
		return nil
	}
}

// objectDocCodeElement はコードを表すobjectDocElementです
type objectDocCodeElement struct {
	Codes []*objectDocCodeElementCode `json:"codes"`
}

func (e *objectDocCodeElement) checkAndOutput(remote objectDocElement) error {
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
				if err := code.output(code); err != nil {
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
	output   func(*objectDocCodeElementCode) error
}

// objectDocCodeElement はコールアウトを表すobjectDocElementです
type objectDocCalloutElement struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	output func(*objectDocCalloutElement) error
}

func (e *objectDocCalloutElement) checkAndOutput(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocCalloutElement); !ok {
		return errUnmatch
	} else if e.Type != remote.Type || e.Title != remote.Title || e.Body != remote.Body {
		return errUnmatch
	} else if e.output != nil {
		return e.output(e)
	} else {
		return nil
	}
}

type objectDocParametersElement []*objectDocParameter

func (e *objectDocParametersElement) checkAndOutput(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocParametersElement); !ok {
		return errUnmatch
	} else if len(*e) != len(*remote) {
		return errUnmatch
	} else {
		for i, param := range *e {
			rp := (*remote)[i]
			if param.Property != rp.Property || param.Field != rp.Field || param.Type != rp.Type || param.Description != rp.Description || param.ExampleValue != rp.ExampleValue {
				return errUnmatch
			} else if param.output != nil {
				if err := param.output(param); err != nil {
					return err
				}
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
	Property     string
	Field        string
	Type         string
	Description  string
	ExampleValue string `json:"Example value"`
	output       func(*objectDocParameter) error
}
