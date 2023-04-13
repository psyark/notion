package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

type coder interface {
	code() jen.Code
}

var _ = []coder{
	&specificObject{},
	&abstractObject{},
	&field{},
	&fixedStringField{},
	independentComment(""),
}

type builder struct {
	coders []coder
	global *builder
}

func (b *builder) addComment(comment string) {
	b.coders = append(b.coders, independentComment(comment))
}

func (b *builder) addSpecificObject(name string, comment string) *specificObject {
	o := &specificObject{}
	o.name = strings.TrimSpace(name)
	o.comment = comment
	b.coders = append(b.coders, o)
	return o
}

func (b *builder) addAbstractObject(name string, specifiedBy string, comment string) *abstractObject {
	o := &abstractObject{}
	o.name = strings.TrimSpace(name)
	o.specifiedBy = specifiedBy
	o.comment = comment
	b.coders = append(b.coders, o)
	return o
}

func (b *builder) addAbstractObjectToGlobalIfNotExists(name string, specifiedBy string) *abstractObject {
	if o := b.global.getAbstractObject(name); o != nil {
		return o
	}
	return b.global.addAbstractObject(name, specifiedBy, "")
}

func (b *builder) statement() jen.Statement {
	s := jen.Statement{}
	for _, item := range b.coders {
		s = append(s, jen.Line().Line(), item.code())
	}
	return s
}

func (b *builder) getAbstractObject(name string) *abstractObject {
	for _, item := range b.coders {
		if item, ok := item.(*abstractObject); ok && item.name == name {
			return item
		}
	}
	if b.global != nil {
		return b.global.getAbstractObject(name)
	}
	return nil
}

func (b *builder) getSpecificObject(name string) *specificObject {
	for _, item := range b.coders {
		if item, ok := item.(*specificObject); ok && item.name == name {
			return item
		}
	}
	if b.global != nil {
		return b.global.getSpecificObject(name)
	}
	return nil
}

type independentComment string

func (c independentComment) code() jen.Code {
	return jen.Comment(string(c))
}
