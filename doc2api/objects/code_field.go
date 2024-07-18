package objects

import (
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

var lineBreak = regexp.MustCompile(`\s*\n\s*`)

// fieldCoderはstructフィールドのコード生成器です
type fieldCoder interface {
	fieldCode() jen.Code
	goName() string // jsonではなくgolang側の名前
	getTypeCode() jen.Code
}

var (
	_ fieldCoder = &VariableField{}
	_ fieldCoder = &fixedStringField{}
)

// 一般的なフィールド
type VariableField struct {
	name                  string
	typeCode              jen.Code
	comment               string
	omitEmpty             bool
	discriminatorValue    string
	discriminatorNotEmpty bool // Userに使う
}

func (f *VariableField) fieldCode() jen.Code {
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
		code.Comment(lineBreak.ReplaceAllString(f.comment, " "))
	}
	return code
}

func (f *VariableField) goName() string {
	return strcase.UpperCamelCase(f.name)
}
func (f *VariableField) getTypeCode() jen.Code {
	return f.typeCode
}

func (f *VariableField) getUnion(c *Converter) *UnionObject {
	code := jen.Var().Id("_").Add(f.typeCode).GoString()
	name := strings.TrimPrefix(code, "var _ ")
	return c.getUnionObject(name)
}

// 固定文字列が入るフィールド
// TODO DiscriminatorField とかにする
type fixedStringField struct {
	name    string
	value   string
	comment string
}

func (f *fixedStringField) fieldCode() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Id("always" + strcase.UpperCamelCase(f.value)).Tag(map[string]string{"json": f.name})
	if f.comment != "" {
		code.Comment(lineBreak.ReplaceAllString(f.comment, " "))
	}
	return code
}

func (f *fixedStringField) goName() string {
	return strcase.UpperCamelCase(f.name)
}
func (f *fixedStringField) getTypeCode() jen.Code {
	return jen.Id("always" + strcase.UpperCamelCase(f.value))
}
