package doc2api

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

type coder interface {
	code() jen.Code
}

var _ = []coder{
	&class{},
}

type builder []coder

func (b *builder) add(c coder) {
	*b = append(*b, c)
}

func (b *builder) statement() jen.Statement {
	s := jen.Statement{}
	for _, item := range *b {
		s = append(s, item.code())
	}
	return s
}

func (b *builder) getClass(name string) *class {
	for _, item := range *b {
		if item, ok := item.(*class); ok && item.name == name {
			return item
		}
	}
	return nil
}

type class struct {
	name    string
	comment string
	fields  []*field
}

func (c *class) addField(f *field) {
	c.fields = append(c.fields, f)
}

func (c *class) code() jen.Code {
	fields := []jen.Code{}
	for _, field := range c.fields {
		fields = append(fields, field.code())
	}
	return jen.Line().Comment(c.comment).Line().Type().Id(c.name).Struct(fields...)
}

type field struct {
	name     string
	typeCode jen.Code
	comment  string
}

func (f *field) code() jen.Code {
	goName := strings.ToUpper(f.name[0:1]) + f.name[1:]
	return jen.Id(goName).Add(f.typeCode).Tag(map[string]string{"json": f.name}).Comment(f.comment)
}
