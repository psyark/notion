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

// 一般的なフィールド
type field struct {
	name                  string
	typeCode              jen.Code
	comment               string
	omitEmpty             bool
	discriminatorValue    string
	discriminatorNotEmpty bool // Userに使う
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

// Unionが入るフィールド
type unionField struct {
	name      string
	unionName string
	comment   string
}

func (f *unionField) fieldCode() jen.Code {
	code := jen.Id(f.goName()).Id(f.unionName)
	if f.name != "" {
		code.Tag(map[string]string{"json": f.name})
	}
	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}

func (f *unionField) goName() string {
	return strcase.UpperCamelCase(f.name)
}
func (f *unionField) getTypeCode() jen.Code {
	return jen.Id(f.unionName)
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
