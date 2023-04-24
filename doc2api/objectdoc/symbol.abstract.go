package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// abstractObject は、共通の性質を持つ concreteObject 又は abstractObject の総称です
//
// interfaceが使われる点やUnmarshalerが生成される点で unionObjectと似ていますが、
// 以下のような違いがあり、目的に応じて使い分けがされます
// - abstractObject は共通のフィールドを必要とする場合がある
// - とあるオブジェクトが直接所属する abstractObject は最大1つ
//
// 例えば File や Property が abstractObject です
// 生成されるGoコードではinterface（共通するフィールドがある場合はstructも）で表現されます
type abstractObject struct {
	objectCommon
	derivedIdentifierKey string // "type", "object" など、派生を識別するためのプロパティ名
	fieldsComment        string
	derivedObjects       []*concreteObject // derivedObjects は、この abstractObjectから派生した concreteObject です
	specialMethods       []specialMethodCoder
}

func (c *abstractObject) addToUnion(union *unionObject) {
	c.unions = append(c.unions, union)
	union.members = append(union.members, c)
}

func (c *abstractObject) addFields(fields ...fieldCoder) *abstractObject {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *abstractObject) hasCommonField() bool {
	if len(c.fields) != 0 {
		return true
	}
	if c.parent != nil {
		return c.parent.hasCommonField()
	}
	return false
}

func (c *abstractObject) symbolCode(b *builder) jen.Code {
	code := jen.Line()

	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	// インターフェイス本体とisメソッド
	code.Type().Id(c.name()).InterfaceFunc(func(g *jen.Group) {
		for _, iface := range c.allInterfaces() {
			g.Id(iface.name()) // 親インターフェイスの継承
		}
		g.Id("is" + c.name()).Params() // このインターフェイスのisメソッド
		for _, sm := range c.specialMethods {
			g.Add(sm.declarationCode()) // 特殊メソッド
		}
		// g.Id("Get"+strcase.UpperCamelCase(c.specifiedBy)).Params().String()
		// 共通フィールドのgetter宣言
		for _, f := range c.fields {
			g.Id("Get" + f.goName()).Params().Add(f.getTypeCode())
		}
	}).Line()

	// 共通フィールド
	if len(c.fieldCodes()) != 0 {
		copyOfC := *c
		copyOfC.name_ = c.commonObjectName()
		copyOfC.comment = c.fieldsComment
		code.Add(copyOfC.objectCommon.symbolCode(b))
		// 共通フィールドのgetter定義
		for _, f := range c.fields {
			code.Line().Func().Params(jen.Id("c").Op("*").Id(copyOfC.name_)).Id("Get" + f.goName()).Params().Add(f.getTypeCode()).Block(
				jen.Return().Id("c").Dot(f.goName()),
			)
		}
	}

	code.Add(c.derivedUnmarshaler())

	return code
}

func (c *abstractObject) derivedUnmarshalerName() string {
	return strcase.LowerCamelCase(c.name()) + "Unmarshaler"
}

func (c *abstractObject) derivedUnmarshaler() jen.Code {
	code := &jen.Statement{}
	code.Line().Type().Id(c.derivedUnmarshalerName()).Struct(
		jen.Id("value").Id(c.name()),
	)
	code.Line().Comment(fmt.Sprintf("UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance\naccording to the %q field of the message.", c.derivedIdentifierKey))
	code.Line().Func().Params(jen.Id("u").Op("*").Id(c.derivedUnmarshalerName())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().BlockFunc(func(g *jen.Group) {
		g.If(jen.String().Call(jen.Id("data"))).Op("==").Lit("null").Block(
			jen.Id("u").Dot("value").Op("=").Nil(),
			jen.Return().Nil(),
		)

		if c.derivedIdentifierKey != "" && false {
			// TODO 従来こちらのコードを使ってきたが、Filterのようなオブジェクトではelse以下を使い分ける必要があった
			// 試しに else 以下に統一したところ、特段問題が無いようだったので、いずれ後者に統一するかを検討する
			// その場合、getType やderivedIdentifierKeyも削れるか検討する
			g.Switch().Id("get" + strcase.UpperCamelCase(c.derivedIdentifierKey)).Call(jen.Id("data")).BlockFunc(func(g *jen.Group) {
				for _, derived := range c.derivedObjects {
					g.Case(jen.Lit(derived.getIdentifierValue(c.derivedIdentifierKey)))
					g.Id("u").Dot("value").Op("=").Op("&").Id(derived.name()).Values()
				}
				g.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: data has unknown %s field: %%s", c.name(), c.derivedIdentifierKey)), jen.String().Call(jen.Id("data"))))
			})
		} else {
			g.Id("t").Op(":=").StructFunc(func(g *jen.Group) {
				for _, derived := range c.derivedObjects {
					if derived.derivedIdentifierValue != "" {
						g.Id(strcase.UpperCamelCase(derived.derivedIdentifierValue)).Qual("encoding/json", "RawMessage").Tag(map[string]string{"json": derived.derivedIdentifierValue})
					}
				}
			}).Values()
			g.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("t")).Op(";").Id("err").Op("!=").Nil()).Block(
				jen.Return().Err(),
			)
			g.Switch().BlockFunc(func(g *jen.Group) {
				defaultCode := jen.Default().Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshal %s: %%s", c.name())), jen.String().Call(jen.Id("data")))
				for _, derived := range c.derivedObjects {
					if derived.derivedIdentifierValue != "" {
						g.Case(jen.Id("t").Dot(strcase.UpperCamelCase(derived.derivedIdentifierValue)).Op("!=").Nil())
						g.Id("u").Dot("value").Op("=").Op("&").Id(derived.name()).Values()
					} else {
						defaultCode = jen.Default().Id("u").Dot("value").Op("=").Op("&").Id(derived.name()).Values()
					}
				}
				g.Add(defaultCode)
			})
		}

		g.Return().Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("u").Dot("value"))
	}).Line()
	code.Line().Func().Params(jen.Id("u").Op("*").Id(c.derivedUnmarshalerName())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().Qual("encoding/json", "Marshal").Call(jen.Id("u").Dot("value")),
	).Line()
	return code
}

func (c *abstractObject) commonObjectName() string {
	return c.name() + "Common"
}

// specialMethodCoder はabstractObject（とその実装オブジェクト）に特別なメソッドを追加したい場合に使います
type specialMethodCoder interface {
	declarationCode() jen.Code
	implementationCode(*concreteObject) jen.Code
}

type abstractList struct {
	name_      string // Deprecated: TODO 消す
	targetName string
}

func (c *abstractList) name() string { return c.name_ }
func (c *abstractList) symbolCode(b *builder) jen.Code {
	target := getSymbol[abstractObject](b, c.targetName)
	return jen.Line().Type().Id(c.name()).Index().Id(c.targetName).Line().Func().Params(jen.Id("a").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.Id("t").Op(":=").Index().Id(target.derivedUnmarshalerName()).Values(),
		jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("t")).Op(";").Err().Op("!=").Nil()).Block(
			jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
		),
		jen.Op("*").Id("a").Op("=").Make(jen.Index().Id(c.targetName), jen.Len(jen.Id("t"))),
		jen.For(jen.List(jen.Id("i"), jen.Id("u")).Op(":=").Range().Id("t")).Block(
			jen.Parens(jen.Op("*").Id("a")).Index(jen.Id("i")).Op("=").Id("u").Dot("value"),
		),
		jen.Return().Nil(),
	)
}

type abstractMap struct {
	name_      string // Deprecated: TODO 消す
	targetName string
}

func (c *abstractMap) name() string { return c.name_ }
func (c *abstractMap) symbolCode(b *builder) jen.Code {
	target := getSymbol[abstractObject](b, c.targetName)
	return jen.Line().Type().Id(c.name()).Map(jen.String()).Id(c.targetName).Line().Func().Params(jen.Id("m").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.Id("t").Op(":=").Map(jen.String()).Id(target.derivedUnmarshalerName()).Values(),
		jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("t")).Op(";").Err().Op("!=").Nil()).Block(
			jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
		),
		jen.Op("*").Id("m").Op("=").Id(c.name()).Values(),
		jen.For(jen.List(jen.Id("k"), jen.Id("u")).Op(":=").Range().Id("t")).Block(
			jen.Parens(jen.Op("*").Id("m")).Index(jen.Id("k")).Op("=").Id("u").Dot("value"),
		),
		jen.Return().Nil(),
	)
}
