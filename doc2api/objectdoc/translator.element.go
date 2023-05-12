package objectdoc

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

var _ = []docElement{
	&blockElement{},
	&parameterElement{},
}

type docElement interface {
	template() jen.Code
}

type blockElement struct {
	Kind string
	Text string
}

func (e *blockElement) template() jen.Code {
	return jen.Id("c").Dot("nextMustBlock").Call(
		jen.Id("blockElement").Values(jen.Dict{
			jen.Id("Kind"): jen.Lit(e.Kind),
			jen.Id("Text"): jen.Lit(e.Text),
		}),
		jen.Func().Params(jen.Id("e").Id("blockElement")).Block(jen.Comment("TODO")),
	)
}

type parameterElement struct {
	Property     string
	Type         string
	Description  string
	ExampleValue string
}

func (e *parameterElement) setCell(header string, v string) error {
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

func (e *parameterElement) template() jen.Code {
	return jen.Id("c").Dot("nextMustParameter").Call(
		jen.Id("parameterElement").Values(jen.Dict{
			jen.Id("Property"):     jen.Lit(e.Property),
			jen.Id("Type"):         jen.Lit(e.Type),
			jen.Id("Description"):  jen.Lit(e.Description),
			jen.Id("ExampleValue"): jen.Lit(e.ExampleValue),
		}),
		jen.Func().Params(jen.Id("e").Id("parameterElement")).Block(jen.Comment("TODO")),
	)
}

// asField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *parameterElement) asField(typeCode jen.Code, options ...fieldOption) *field {
	f := &field{
		name:     e.Property,
		typeCode: typeCode,
		comment:  e.Description,
	}
	for _, o := range options {
		o(f)
	}
	return f
}

// asUnionField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *parameterElement) asUnionField(unionName string) *unionField {
	return &unionField{
		name:      e.Property,
		unionName: unionName,
		comment:   e.Description,
	}
}

// asFixedStringField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *parameterElement) asFixedStringField(b *builder) *fixedStringField {
	for _, value := range []string{e.ExampleValue, e.Type} {
		if value != "" {
			if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
				value = strings.TrimPrefix(strings.TrimSuffix(value, `"`), `"`)
				b.global.addAlwaysStringIfNotExists(value)
				return &fixedStringField{
					name:    e.Property,
					value:   value,
					comment: e.Description,
				}
			}
			panic(value)
		}
	}
	panic("asFixedStringField")
}
