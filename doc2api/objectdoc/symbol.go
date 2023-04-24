package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// symbolCoder はソースコードのトップレベルに置かれる、名前を持つシンボルの生成器です。
type symbolCoder interface {
	symbolCode(*builder) jen.Code
	name() string
}

var _ = []symbolCoder{
	&objectCommon{},
	&specificObject{},
	&abstractObject{},
	&abstractList{},
	&abstractMap{},
	&unionObject{},
	&unmarshalTest{},
	alwaysString(""),
}

// derivedCoder はabstractObjectの派生として存在できるオブジェクトを作成するためのCoderです
// TODO これを unionのメンバー用に作り直す
type derivedCoder interface {
	symbolCoder
	getIdentifierValue(identifierKey string) string
	setParent(*abstractObject)
}

var _ = []derivedCoder{
	&specificObject{},
	&abstractObject{},
}

type objectCommon struct {
	name_   string
	comment string
	fields  []fieldCoder

	// parent はこのオブジェクトの派生元です。派生元とは共通のフィールドを提供しているオブジェクトであり、
	// 例えば ExternalFile に対する File を指します。一方、FileOrIcon は unionsとして表現します。
	parent *abstractObject

	// unions は自分が所属するunionObjectです。
	// objectCommonを継承する各クラスは、symbolCode メソッド中でこのunionのisメソッドを実装する必要があります
	unions []*unionObject
}

func (c *objectCommon) name() string {
	return c.name_
}

// TODO: derivedIdentifierValueみたいなプロパティにする？
func (c *objectCommon) getIdentifierValue(specifiedBy string) string {
	for _, f := range c.fields {
		if f, ok := f.(*fixedStringField); ok && f.name == specifiedBy {
			return f.value
		}
	}
	if c.parent != nil {
		if v := c.parent.getIdentifierValue(specifiedBy); v != "" {
			return v
		}
	}
	return ""
}

func (c *objectCommon) setParent(parent *abstractObject) {
	if c.parent != nil {
		panic(fmt.Errorf("👪 %s has two parents: %s vs %s", c.name(), c.parent.name(), parent.name()))
	}
	c.parent = parent
}

// TODO 良い名前に
func (c *objectCommon) getAncestors() []symbolCoder {
	ancestors := []symbolCoder{}
	for _, u := range c.unions {
		ancestors = append(ancestors, u)
	}
	if c.parent != nil {
		ancestors = append(ancestors, c.parent)
		ancestors = append(ancestors, c.parent.getAncestors()...)
	}
	return ancestors
}

func (c *objectCommon) fieldCodes() []jen.Code {
	fields := []jen.Code{}
	if c.parent != nil && c.parent.hasCommonField() {
		fields = append(fields, jen.Id(c.parent.commonObjectName()))
	}
	for _, f := range c.fields {
		fields = append(fields, f.fieldCode())
	}
	return fields
}

func (c *objectCommon) addFields(fields ...fieldCoder) *objectCommon {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *objectCommon) symbolCode(b *builder) jen.Code {
	code := &jen.Statement{}
	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	return code.Type().Id(c.name_).Struct(c.fieldCodes()...).Line()
}

func (c *objectCommon) fieldUnmarshalerCode(b *builder) jen.Code {
	code := &jen.Statement{}

	interfaceFields := []*interfaceField{}
	for _, f := range c.fields {
		if f, ok := f.(*interfaceField); ok {
			interfaceFields = append(interfaceFields, f)
		}
	}

	if len(interfaceFields) != 0 {
		code.Comment("UnmarshalJSON assigns the appropriate implementation to interface field(s)").Line()
		code.Func().Params(jen.Id("o").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().BlockFunc(func(g *jen.Group) {
			g.Type().Id("Alias").Id(c.name())
			g.Id("t").Op(":=").Op("&").StructFunc(func(g *jen.Group) {
				g.Op("*").Id("Alias")
				for _, f := range interfaceFields {
					if a := getSymbol[abstractObject](b, f.typeName); a != nil {
						g.Id(strcase.UpperCamelCase(f.name)).Id(a.derivedUnmarshalerName()).Tag(map[string]string{"json": f.name})
					} else if u := getSymbol[unionObject](b, f.typeName); u != nil {
						g.Id(strcase.UpperCamelCase(f.name)).Id(u.memberUnmarshalerName()).Tag(map[string]string{"json": f.name})
					} else {
						panic(fmt.Errorf("unknown symbol: %s", f.typeName))
					}
				}
			}).Values(jen.Dict{
				jen.Id("Alias"): jen.Parens(jen.Op("*").Id("Alias")).Call(jen.Id("o")),
			})
			g.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
			)
			for _, f := range interfaceFields {
				fieldName := strcase.UpperCamelCase(f.name)
				g.Id("o").Dot(fieldName).Op("=").Id("t").Dot(fieldName).Dot("value")
			}
			g.Return().Nil()
		}).Line()
	}

	return code
}

type unmarshalTest struct {
	targetName string
	jsonCodes  []string
}

func (c *unmarshalTest) name() string {
	return fmt.Sprintf("Test%s_unmarshal", c.targetName)
}
func (c *unmarshalTest) symbolCode(b *builder) jen.Code {
	var initCode, referCode jen.Code
	switch symbol := b.getSymbol(c.targetName).(type) {
	case *abstractObject:
		initCode = jen.Id("target").Op(":=").Op("&").Id(symbol.derivedUnmarshalerName()).Values()
		referCode = jen.Id("target").Dot("value")
	case *specificObject, *abstractMap:
		initCode = jen.Id("target").Op(":=").Op("&").Id(c.targetName).Values()
		referCode = jen.Id("target")
	default:
		panic(symbol)
	}

	return jen.Line().Func().Id(c.name()).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		initCode,
		jen.Id("tests").Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
			for _, t := range c.jsonCodes {
				g.Line().Lit(t)
			}
			g.Line()
		}),
		jen.For().List(jen.Id("_"), jen.Id("wantStr")).Op(":=").Range().Id("tests").Block(
			jen.Id("want").Op(":=").Index().Byte().Call(jen.Id("wantStr")),
			jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("want"), jen.Id("target"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Id("t").Dot("Fatal").Call(jen.Err()),
			),
			jen.List(jen.Id("got"), jen.Id("_")).Op(":=").Qual("encoding/json", "Marshal").Call(referCode),
			jen.If(jen.List(jen.Id("want"), jen.Id("got"), jen.Id("ok")).Op(":=").Id("compareJSON").Call(jen.Id("want"), jen.Id("got")).Op(";").Op("!").Id("ok")).Block(
				jen.Id("t").Dot("Fatal").Call(jen.Qual("fmt", "Errorf").Call(jen.Lit("mismatch:\nwant: %s\ngot : %s"), jen.Id("want"), jen.Id("got"))),
			),
		),
	)
}

type alwaysString string

func (c alwaysString) symbolCode(b *builder) jen.Code {
	code := jen.Type().Id(c.name()).String().Line()
	code.Func().Params(jen.Id("s").Id(c.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().List(jen.Index().Byte().Call(jen.Lit(fmt.Sprintf("%q", string(c)))), jen.Nil()),
	)
	return code
}

func (c alwaysString) name() string {
	return "always" + strcase.UpperCamelCase(string(c))
}
