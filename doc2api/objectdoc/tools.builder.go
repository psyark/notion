package objectdoc

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

type coder interface {
	code() jen.Code
}

var _ = []coder{
	&specificObject{},
	&abstractObject{},
	&field{},
	&fixedStringField{},
	&interfaceField{},
	independentComment(""),
	alwaysString(""),
}

type builder struct {
	fileName string
	url      string
	coders   []coder
	global   *builder
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

func (b *builder) addAlwaysStringIfNotExists(value string) {
	for _, c := range b.coders {
		if c, ok := c.(alwaysString); ok && string(c) == value {
			return
		}
	}
	b.coders = append(b.coders, alwaysString(value))
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

type alwaysString string

func (c alwaysString) code() jen.Code {
	name := "always" + strcase.UpperCamelCase(string(c))
	code := jen.Type().Id(name).String().Line()
	code.Func().Params(jen.Id("s").Id(name)).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().List(jen.Index().Byte().Call(jen.Lit(fmt.Sprintf("%q", string(c)))), jen.Nil()),
	)
	return code
}
