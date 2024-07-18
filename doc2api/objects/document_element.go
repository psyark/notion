package objects

import (
	"fmt"
)

var _ = []DocumentElement{
	&Block{},
	&Parameter{},
}

// DocumentElement はドキュメントに出現する要素です
type DocumentElement interface {
	isElement()
}

// Block は、KindとTextを持つ単純な要素です
type Block struct {
	Kind string // "FencedCodeBlock", "BlockQuote" etc...
	Text string
}

func (e *Block) isElement() {}

// Parameter はパラメータを表す要素です
type Parameter struct {
	Property     string // パラメータ名
	Type         string // パラメータの型を表す人間向けの（厳密性のない）説明
	Description  string // パラメータの説明
	ExampleValue string // 値の例
}

func (e *Parameter) setCell(header string, v string) error {
	switch header {
	case "Property", "Field", "Parameter", "": // "" for https://developers.notion.com/reference/emoji-object
		e.Property = v
	case "Type", "HTTP method":
		e.Type = v
	case "Description", "Endpoint":
		e.Description = v
	case "Example value", "Example values":
		e.ExampleValue = v
	case "Updatable": // https://developers.notion.com/reference/user
	default:
		return fmt.Errorf("setCell: %q", header)
	}
	return nil
}

func (e *Parameter) isElement() {}
