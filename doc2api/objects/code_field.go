package objects

import (
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

var lineBreak = regexp.MustCompile(`\s*\n\s*`)

// fieldRenderer はstructフィールドのコード生成器です
type fieldRenderer interface {
	renderField() jen.Code
}

var _ = []fieldRenderer{
	&VariableField{},
	&DiscriminatorField{},
}

// 一般的なフィールド
type VariableField struct {
	name                  string
	typeCode              jen.Code
	comment               string
	omitEmpty             bool
	discriminatorValue    string
	discriminatorNotEmpty bool // Userに使う
}

func (f *VariableField) renderField() jen.Code {
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

// このフィールドがUnionInterface型である場合、それを返します
func (f *VariableField) getUnionInterface(c *Converter) *UnionInterface {
	code := jen.Var().Id("_").Add(f.typeCode).GoString()
	name := strings.TrimPrefix(code, "var _ ")
	return c.getUnionInterface(name)
}

// 識別子が入るフィールド
type DiscriminatorField struct {
	name    string
	value   string
	comment string
}

func (f *DiscriminatorField) renderField() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Id("always" + strcase.UpperCamelCase(f.value)).Tag(map[string]string{"json": f.name})
	if f.comment != "" {
		code.Comment(lineBreak.ReplaceAllString(f.comment, " "))
	}
	return code
}
