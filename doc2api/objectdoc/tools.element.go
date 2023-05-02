package objectdoc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type errUnmatch struct {
	remote any
	local  any
}

func (err errUnmatch) Error() string {
	r, _ := json.MarshalIndent(err.remote, "", "  ")
	l, _ := json.MarshalIndent(err.local, "", "  ")
	return fmt.Sprintf("unmatch:\nremote: %s\n local: %s", r, l)
}

// objectDocElement はNotion API Reference のObjects以下 (e.g. https://developers.notion.com/reference/block) で
// 出現するエレメントのインターフェイスです
//
// このインターフェイスは以下の2通りの使われ方をします
// 1. Notion API Referenceの最新のデータ（リモート。変換ロジックは常にnilです）
// 2. ある時点の(1)がGoコードとして保存されたもの（ローカルコピー。多くの場合変換ロジックを伴います）
type objectDocElement interface {
	localCopy(typeName string, outputCode jen.Code) jen.Code
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
	&objectDocImageElement{},
}

// objectDocHeadingElement は見出しを表すobjectDocElementです
type objectDocHeadingElement struct {
	Text   string
	output func(*objectDocHeadingElement, *builder)
}

func (e *objectDocHeadingElement) localCopy(typeName string, outputCode jen.Code) jen.Code {
	return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
		jen.Id("Text"):   jen.Lit(e.Text),
		jen.Id("output"): outputCode,
	})
}

func (e *objectDocHeadingElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocHeadingElement); ok {
		if e.Text == remote.Text {
			e.output(e, b)
			return nil
		}
	}
	return errUnmatch{remote: remote, local: e}
}

// objectDocParagraphElement は段落を表すobjectDocElementです
type objectDocParagraphElement struct {
	Text   string
	output func(*objectDocParagraphElement, *builder)
}

func (e *objectDocParagraphElement) localCopy(typeName string, outputCode jen.Code) jen.Code {
	return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
		jen.Id("Text"):   jen.Lit(e.Text),
		jen.Id("output"): outputCode,
	})
}

func (e *objectDocParagraphElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocParagraphElement); ok {
		if e.Text == remote.Text {
			e.output(e, b)
			return nil
		}
	}
	return errUnmatch{remote: remote, local: e}
}

// objectDocAPIHeaderElement はAPI Headerを表すobjectDocElementです
type objectDocAPIHeaderElement struct {
	Title  string `json:"title"`
	output func(*objectDocAPIHeaderElement, *builder)
}

func (e *objectDocAPIHeaderElement) localCopy(typeName string, outputCode jen.Code) jen.Code {
	panic("TODO")
}

func (e *objectDocAPIHeaderElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocAPIHeaderElement); ok {
		if e.Title == remote.Title {
			e.output(e, b)
			return nil
		}
	}
	return errUnmatch{remote: remote, local: e}
}

// objectDocCodeElement はコードを表すobjectDocElementです
type objectDocCodeElement struct {
	Codes []*objectDocCodeElementCode `json:"codes"`
}

func (e *objectDocCodeElement) localCopy(typeName string, _ jen.Code) jen.Code {
	outputCode := jen.Func().Params(jen.Id("e").Op("*").Id("objectDocCodeElementCode"), jen.Id("b").Op("*").Id("builder")).Error().Block(jen.Return().Nil().Comment("TODO"))
	codes := jen.Statement{}
	for _, c := range e.Codes {
		codes = append(codes, jen.Values(jen.Dict{
			jen.Id("Name"):     jen.Lit(c.Name),
			jen.Id("Language"): jen.Lit(c.Language),
			jen.Id("Code"):     jen.Lit(c.Code),
			jen.Id("output"):   outputCode,
		}))
	}
	return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
		jen.Id("Codes"): jen.Index().Op("*").Id("objectDocCodeElementCode").Values(codes...),
	})
}

func (e *objectDocCodeElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocCodeElement); ok {
		if len(e.Codes) == len(remote.Codes) {
			for i, code := range e.Codes {
				rc := remote.Codes[i]
				if code.Name != rc.Name || code.Language != rc.Language || code.Code != rc.Code {
					return errUnmatch{remote: rc, local: code}
				}
				code.output(code, b)
			}
			return nil
		}
	}
	return errUnmatch{remote: remote, local: e}
}

type objectDocCodeElementCode struct {
	Name     string `json:"string"`
	Language string `json:"language"`
	Code     string `json:"code"`
	output   func(*objectDocCodeElementCode, *builder)
}

// objectDocCodeElement はコールアウトを表すobjectDocElementです
type objectDocCalloutElement struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	output func(*objectDocCalloutElement, *builder)
}

func (e *objectDocCalloutElement) localCopy(typeName string, outputCode jen.Code) jen.Code {
	return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
		jen.Id("Type"):   jen.Lit(e.Type),
		jen.Id("Title"):  jen.Lit(e.Title),
		jen.Id("Body"):   jen.Lit(e.Body),
		jen.Id("output"): outputCode,
	})
}

func (e *objectDocCalloutElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocCalloutElement); ok {
		if e.Type == remote.Type && e.Title == remote.Title && e.Body == remote.Body {
			e.output(e, b)
			return nil
		}
	}
	return errUnmatch{remote: remote, local: e}
}

type objectDocParametersElement []*objectDocParameter

func (e *objectDocParametersElement) localCopy(typeName string, _ jen.Code) jen.Code {
	outputCode := jen.Func().Params(jen.Id("e").Op("*").Id("objectDocParameter"), jen.Id("b").Op("*").Id("builder")).Error().Block(jen.Return().Nil().Comment("TODO"))
	params := jen.Statement{}
	for _, p := range *e {
		values := jen.Dict{
			jen.Id("Type"):        jen.Lit(p.Type),
			jen.Id("Description"): jen.Lit(p.Description),
			jen.Id("output"):      outputCode,
		}
		if p.Property != "" {
			values[jen.Id("Property")] = jen.Lit(p.Property)
		}
		if p.Field != "" {
			values[jen.Id("Field")] = jen.Lit(p.Field)
		}
		if p.ExampleValue != "" {
			values[jen.Id("ExampleValue")] = jen.Lit(p.ExampleValue)
		}
		if p.ExampleValues != "" {
			values[jen.Id("ExampleValues")] = jen.Lit(p.ExampleValues)
		}
		params = append(params, jen.Values(values))
	}
	return jen.Line().Op("&").Id(typeName).Values(params...)
}

func (e *objectDocParametersElement) checkAndOutput(remote objectDocElement, b *builder) error {
	if remote, ok := remote.(*objectDocParametersElement); ok {
		if len(*e) == len(*remote) {
			for i, param := range *e {
				if err := param.checkAndOutput((*remote)[i], b); err != nil {
					return err
				}
			}
			return nil
		}
	}
	return errUnmatch{remote: remote, local: e}
}

func (t *objectDocParametersElement) UnmarshalJSON(data []byte) error {
	raw := objectDocParametersElementRaw{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mapSlice := make([]map[string]string, raw.Rows)
	for r := range mapSlice {
		m := map[string]string{}
		for c := 0; c < raw.Cols; c++ {
			key := raw.Data[fmt.Sprintf("h-%d", c)]
			if key == "" {
				key = "Property" // https://developers.notion.com/reference/emoji-object
			}
			m[key] = stripMarkdown(raw.Data[fmt.Sprintf("%d-%d", r, c)])
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

type objectDocParametersElementRaw struct {
	Data map[string]string `json:"data"`
	Cols int               `json:"cols"`
	Rows int               `json:"rows"`
}

// objectDocParameter はドキュメントに書かれているパラメータです
type objectDocParameter struct {
	Property      string `json:",omitempty"`
	Field         string `json:",omitempty"`
	Type          string `json:",omitempty"`
	Description   string `json:",omitempty"`
	ExampleValue  string `json:"Example value,omitempty"`
	ExampleValues string `json:"Example values,omitempty"`
	output        func(*objectDocParameter, *builder)
}

func (e *objectDocParameter) checkAndOutput(remote *objectDocParameter, b *builder) error {
	if e.Property != remote.Property || e.Field != remote.Field || e.Type != remote.Type || e.Description != remote.Description || e.ExampleValue != remote.ExampleValue || e.ExampleValues != remote.ExampleValues {
		return errUnmatch{remote: remote, local: e}
	}
	e.output(e, b)
	return nil
}

type fieldOption func(f *field)

var omitEmpty fieldOption = func(f *field) {
	f.omitEmpty = true
}

var discriminatorNotEmpty fieldOption = func(f *field) {
	f.discriminatorNotEmpty = true
}

// asField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *objectDocParameter) asField(typeCode jen.Code, options ...fieldOption) *field {
	f := &field{
		name:     e.Property + e.Field,
		typeCode: typeCode,
		comment:  e.Description,
	}
	for _, o := range options {
		o(f)
	}
	return f
}

// asInterfaceField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *objectDocParameter) asInterfaceField(typeName string, options ...fieldOption) *interfaceField {
	return &interfaceField{
		name:     e.Property + e.Field,
		typeName: typeName,
		comment:  e.Description,
	}
}

// asFixedStringField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *objectDocParameter) asFixedStringField() *fixedStringField {
	for _, value := range []string{e.ExampleValue, e.ExampleValues, e.Type} {
		if value != "" {
			if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
				return &fixedStringField{
					name:    e.Property + e.Field,
					value:   strings.TrimPrefix(strings.TrimSuffix(value, `"`), `"`),
					comment: e.Description,
				}
			}
			panic(value)
		}
	}
	panic("asFixedStringField")
}

type objectDocImageElement struct {
	Images []*objectDocImageElementImage `json:"images"`
}

type objectDocImageElementImage struct {
	Url     string
	Name    string
	Width   int
	Height  int
	Color   string
	Caption string `json:"caption"`
	output  func(*objectDocImageElementImage, *builder)
}

func (e *objectDocImageElementImage) UnmarshalJSON(data []byte) error {
	type Alias objectDocImageElementImage
	t := &struct {
		*Alias
		Image [5]any `json:"image"`
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	e.Url = t.Image[0].(string)
	e.Name = t.Image[1].(string)
	e.Width = int(t.Image[2].(float64))
	e.Height = int(t.Image[3].(float64))
	e.Color = t.Image[4].(string)
	return nil
}

func (e *objectDocImageElement) localCopy(typeName string, _ jen.Code) jen.Code {
	outputCode := jen.Func().Params(jen.Id("e").Op("*").Id("objectDocImageElementImage"), jen.Id("b").Op("*").Id("builder")).Error().Block(jen.Return().Nil().Comment("TODO"))
	codes := jen.Statement{}
	for _, c := range e.Images {
		codes = append(codes, jen.Values(jen.Dict{
			jen.Id("Url"):     jen.Lit(c.Url),
			jen.Id("Name"):    jen.Lit(c.Name),
			jen.Id("Width"):   jen.Lit(c.Width),
			jen.Id("Height"):  jen.Lit(c.Height),
			jen.Id("Color"):   jen.Lit(c.Color),
			jen.Id("Caption"): jen.Lit(c.Caption),
			jen.Id("output"):  outputCode,
		}))
	}
	return jen.Line().Op("&").Id(typeName).Values(jen.Dict{
		jen.Id("Images"): jen.Index().Op("*").Id("objectDocImageElementImage").Values(codes...),
	})
}

func (e *objectDocImageElement) checkAndOutput(remote objectDocElement, b *builder) error {
	return nil
}
