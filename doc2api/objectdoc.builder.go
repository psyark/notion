package doc2api

import (
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

type coder interface {
	code() jen.Code
}

var _ = []coder{
	&classStruct{},
}

type builder []coder

func (b *builder) add(c coder) {
	*b = append(*b, c)
}

func (b *builder) statement() jen.Statement {
	s := jen.Statement{}
	for _, item := range *b {
		s = append(s, jen.Line().Line(), item.code())
	}
	return s
}

func (b *builder) getClassStruct(name string) *classStruct {
	for _, item := range *b {
		if item, ok := item.(*classStruct); ok && item.name == name {
			return item
		}
	}
	return nil
}

type classStruct struct {
	name    string
	comment string
	fields  []coder
}

func (c *classStruct) addField(f coder) {
	c.fields = append(c.fields, f)
}

// func (c *classStruct) getField(name string) *field {
// 	for _, f := range c.fields {
// 		if f, ok := f.(*field); ok && f.name == name {
// 			return f
// 		}
// 	}
// 	return nil
// }

func (c *classStruct) code() jen.Code {
	fields := []jen.Code{}
	for _, field := range c.fields {
		fields = append(fields, field.code())
	}
	return jen.Comment(c.comment).Line().Type().Id(c.name).Struct(fields...)
}

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
		code.Comment(f.comment)
	}
	return code
}

type fixedStringField struct {
	name    string
	value   string
	comment string
}

func (f *fixedStringField) code() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).String().Tag(map[string]string{"json": f.name, "always": f.value})
	if f.comment != "" {
		code.Comment(f.comment)
	}
	return code
}

type classInterface struct {
	name    string
	comment string
}

func (c *classInterface) code() jen.Code {
	return jen.Comment(c.comment).Line().Type().Id(c.name).Interface()
}

type independentComment string

func (c independentComment) code() jen.Code {
	return jen.Comment(string(c))
}
