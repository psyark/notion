package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// abstractObject は、共通の性質を持つspecificObject又はabstractObjectの総称です
// （例：File）
// 生成されるGoコードではinterface（共通するフィールドがある場合はstructも）で表現されます
type abstractObject struct {
	objectCommon
	specifiedBy    string // "type", "object" など、バリアントを識別するためのプロパティ // TODO variantTypeKeyとかにする？
	fieldsComment  string
	variants       []objectCoder
	specialMethods []specialMethodCoder
}

// addVariant は指定したobjectCoderをこのインターフェイスのバリアントとして登録し、code()に以下のことを行わせます
// - バリアントに対してインターフェイスメソッドを実装
// - JSONメッセージからこのインターフェイスの適切なバリアントを作成するUnmarshalerを作成
func (c *abstractObject) addVariant(variant objectCoder) *abstractObject {
	variant.addParent(c)
	c.variants = append(c.variants, variant)
	return c
}

func (c *abstractObject) addFields(fields ...fieldCoder) *abstractObject {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *abstractObject) hasCommonField() bool {
	if len(c.fields) != 0 {
		return true
	}
	for _, p := range c.parents {
		if p.hasCommonField() {
			return true
		}
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
		for _, p := range c.parents {
			g.Id(p.name()) // 親インターフェイスの継承
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

	// variant Unmarshaler
	if c.specifiedBy != "" {
		code.Add(c.variantUnmarshaler())
	}

	return code
}

func (c *abstractObject) variantUnmarshalerName() string {
	return strcase.LowerCamelCase(c.name()) + "Unmarshaler"
}

func (c *abstractObject) variantUnmarshaler() jen.Code {
	code := &jen.Statement{}
	if len(c.variants) != 0 {
		code.Line().Type().Id(c.variantUnmarshalerName()).Struct(
			jen.Id("value").Id(c.name()),
		)
		code.Line().Comment(fmt.Sprintf("UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance\naccording to the %q field of the message.", c.specifiedBy))
		code.Line().Func().Params(jen.Id("u").Op("*").Id(c.variantUnmarshalerName())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
			jen.If(jen.String().Call(jen.Id("data"))).Op("==").Lit("null").Block(
				jen.Id("u").Dot("value").Op("=").Nil(),
				jen.Return().Nil(),
			),
			jen.Switch().Id("get"+strcase.UpperCamelCase(c.specifiedBy)).Call(jen.Id("data")).BlockFunc(func(g *jen.Group) {

				for _, variant := range c.variants {
					if sf := variant.getSpecifyingField(c.specifiedBy); sf != nil {
						g.Case(jen.Lit(sf.value))
					} else {
						g.Case(jen.Lit(""))
					}

					switch variant := variant.(type) {
					case *specificObject:
						g.Id("u").Dot("value").Op("=").Op("&").Id(variant.name()).Values()
					case *abstractObject:
						g.Id("t").Op(":=").Op("&").Id(variant.variantUnmarshalerName()).Values()
						g.If(jen.Err().Op(":=").Id("t").Dot("UnmarshalJSON").Call(jen.Id("data"))).Op(";").Err().Op("!=").Nil().Block(jen.Return().Err())
						g.Id("u").Dot("value").Op("=").Id("t").Dot("value")
						g.Return().Nil()
					default:
						panic(fmt.Sprintf("%#v", variant))
					}
				}

				g.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: data has unknown %s field: %%s", c.name(), c.specifiedBy)), jen.String().Call(jen.Id("data"))))
			}),
			jen.Return().Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("u").Dot("value")),
		).Line()
		code.Line().Func().Params(jen.Id("u").Op("*").Id(c.variantUnmarshalerName())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
			jen.Return().Qual("encoding/json", "Marshal").Call(jen.Id("u").Dot("value")),
		).Line()
	}
	return code
}

func (c *abstractObject) commonObjectName() string {
	return c.name() + "Common"
}

// specialMethodCoder はabstractObject（とその実装オブジェクト）に特別なメソッドを追加したい場合に使います
type specialMethodCoder interface {
	declarationCode() jen.Code
	implementationCode(*specificObject) jen.Code
}

type abstractList struct {
	name_      string // Deprecated: TODO 消す
	targetName string
}

func (c *abstractList) name() string { return c.name_ }
func (c *abstractList) symbolCode(b *builder) jen.Code {
	target := b.getAbstractObject(c.targetName)
	return jen.Line().Type().Id(c.name()).Index().Id(c.targetName).Line().Func().Params(jen.Id("a").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.Id("t").Op(":=").Index().Id(target.variantUnmarshalerName()).Values(),
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
	target := b.getAbstractObject(c.targetName)
	return jen.Line().Type().Id(c.name()).Map(jen.String()).Id(c.targetName).Line().Func().Params(jen.Id("m").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.Id("t").Op(":=").Map(jen.String()).Id(target.variantUnmarshalerName()).Values(),
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
