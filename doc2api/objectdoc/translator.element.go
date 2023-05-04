package objectdoc

import (
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

// asInterfaceField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *parameterElement) asInterfaceField(typeName string, options ...fieldOption) *interfaceField {
	return &interfaceField{
		name:     e.Property,
		typeName: typeName,
		comment:  e.Description,
	}
}

// asFixedStringField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (e *parameterElement) asFixedStringField() *fixedStringField {
	for _, value := range []string{e.ExampleValue, e.ExampleValue, e.Type} {
		if value != "" {
			if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
				return &fixedStringField{
					name:    e.Property,
					value:   strings.TrimPrefix(strings.TrimSuffix(value, `"`), `"`),
					comment: e.Description,
				}
			}
			panic(value)
		}
	}
	panic("asFixedStringField")
}
