package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// fieldCoderはstructフィールドのコード生成器です
type fieldCoder interface {
	fieldCode() jen.Code
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
