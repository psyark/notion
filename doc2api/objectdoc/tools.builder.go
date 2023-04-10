package objectdoc

import (
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

func (b *builder) add(c coder) {
	b.coders = append(b.coders, c)
}

// addGlobalIfNotExists はグローバルビルダーに（そのコーダーが登録されていなければ）登録します
func (b *builder) addGlobalIfNotExists(c coder) {
	switch c := c.(type) {
	case *specificObject:
		if b.global.getSpecificObject(c.name) != nil {
			return
		}
	case *abstractObject:
		if b.global.getAbstractObject(c.name) != nil {
			return
		}
	}
	b.global.add(c)
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
