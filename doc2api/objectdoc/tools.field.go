package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// fieldCoderはstructフィールドのコード生成器です
type fieldCoder interface {
	fieldCode() jen.Code
	goName() string // jsonではなくgolang側の名前
	getTypeCode() jen.Code
}

var _ = []fieldCoder{
	&field{},
	&interfaceField{},
	&fixedStringField{},
}

// 一般的なフィールド
type field struct {
	name      string
	typeCode  jen.Code
	comment   string
	omitEmpty bool
}

func (f *field) fieldCode() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Add(f.typeCode)

	tag := f.name
	if f.omitEmpty {
		tag += ",omitempty"
	}
	if tag != "" {
		code.Tag(map[string]string{"json": tag})
	}

	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}

func (f *field) goName() string {
	return strcase.UpperCamelCase(f.name)
}
func (f *field) getTypeCode() jen.Code {
	return f.typeCode
}

// インターフェイスが入るフィールド
// TODO 自動判別するようにしたい
type interfaceField struct {
	name     string
	typeName string
	comment  string
}

func (f *interfaceField) fieldCode() jen.Code {
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

func (f *interfaceField) goName() string {
	return strcase.UpperCamelCase(f.name)
}
func (f *interfaceField) getTypeCode() jen.Code {
	return jen.Id(f.typeName)
}

// 固定文字列が入るフィールド
type fixedStringField struct {
	name    string
	value   string
	comment string
}

func (f *fixedStringField) fieldCode() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Id("always" + strcase.UpperCamelCase(f.value)).Tag(map[string]string{"json": f.name})
	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}

func (f *fixedStringField) goName() string {
	return strcase.UpperCamelCase(f.name)
}
func (f *fixedStringField) getTypeCode() jen.Code {
	return jen.Id("always" + strcase.UpperCamelCase(f.value))
}
