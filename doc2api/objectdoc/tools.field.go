package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// 一般的なフィールド
type field struct {
	name     string
	typeCode jen.Code
	comment  string
}

func (f *field) code() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Add(f.typeCode)
	if f.name != "" {
		code.Tag(map[string]string{"json": f.name})
	}
	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}

// インターフェイスが入るフィールド
type interfaceField struct {
	name     string
	typeName string
	comment  string
}

func (f *interfaceField) code() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Id(f.typeName)
	if f.name != "" {
		code.Tag(map[string]string{"json": f.name})
	}
	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}

// 固定文字列が入るフィールド
type fixedStringField struct {
	name    string
	value   string
	comment string
}

func (f *fixedStringField) code() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Id("always" + strcase.UpperCamelCase(f.value)).Tag(map[string]string{"json": f.name})
	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}
