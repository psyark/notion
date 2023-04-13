package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// objectCoder はオブジェクトを作成するためのCoderです
type objectCoder interface {
	coder
	getName() string
	getFields() []coder
	addFields(...coder) objectCoder
	getSpecifyingField(specifiedBy string) *fixedStringField
	addParent(*abstractObject)
}

var _ = []objectCoder{
	&specificObject{},
	&abstractObject{},
}

type objectCommon struct {
	name    string
	comment string
	fields  []coder
	parents []*abstractObject
}

func (c *objectCommon) getName() string {
	return c.name
}
func (c *objectCommon) getFields() []coder {
	return c.fields
}

func (c *objectCommon) getSpecifyingField(specifiedBy string) *fixedStringField {
	for _, f := range c.fields {
		if f, ok := f.(*fixedStringField); ok && f.name == specifiedBy {
			return f
		}
	}
	panic(fmt.Errorf("%s not found for %v", specifiedBy, c.name))
}
func (c *objectCommon) addParent(parent *abstractObject) {
	c.parents = append(c.parents, parent)
}

func (c *objectCommon) getAncestors() []*abstractObject {
	ancestors := []*abstractObject{}
	for _, a := range c.parents {
		ancestors = append(ancestors, a)
		ancestors = append(ancestors, a.getAncestors()...)
	}
	return ancestors
}

// addFields はこのオブジェクトにフィールドを追加します
func (c *objectCommon) addFields(fields ...coder) objectCoder {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *objectCommon) code() jen.Code {
	panic("!")
}

// specificObject は"type"や"object"キーで区別される各オブジェクトです
// （例：type="external" である ExternalFile）
// 生成されるGoコードではstructポインタで表現されます
type specificObject struct {
	objectCommon
}

func (c *specificObject) addFields(fields ...coder) objectCoder {
	c.objectCommon.addFields(fields...)
	return c
}

func (c *specificObject) code() jen.Code {
	code := jen.Line()
	hasInterface := false

	// struct本体
	{
		fields := []jen.Code{}
		for _, p := range c.parents {
			fields = append(fields, jen.Id(strcase.LowerCamelCase(p.name)+"Common"))
		}
		for _, f := range c.fields {
			fields = append(fields, f.code())
			if f, ok := f.(*field); ok && f.isInterface {
				hasInterface = true
			}
		}
		code.Comment(c.comment).Line().Type().Id(c.name).Struct(fields...).Line()
	}

	// インターフェイスを実装
	for _, a := range c.getAncestors() {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name)).Id("is" + a.name).Params().Block().Line()
	}

	// フィールドにインターフェイスを含むため、UnmarshalJSONで前処理を行う
	if hasInterface {
		tmpFields := []jen.Code{jen.Op("*").Id("Alias")}
		bodyCodes := []jen.Code{}
		for _, f := range c.fields {
			if f, ok := f.(*field); ok && f.isInterface {
				interfaceName := (&jen.Statement{f.typeCode}).GoString()
				tmpFields = append(tmpFields, jen.Id(strcase.UpperCamelCase(f.name)).Id(strcase.LowerCamelCase(interfaceName)+"Unmarshaler").Tag(map[string]string{"json": f.name}))
				bodyCodes = append(bodyCodes, jen.Id("o").Dot(strcase.UpperCamelCase(f.name)).Op("=").Id("t").Dot(strcase.UpperCamelCase(f.name)).Dot("value").Line())
			}
		}

		bodyCodes = append(bodyCodes, jen.Return().Nil())

		// プロパティ用UnmarshalJSON
		code.Func().Params(jen.Id("o").Op("*").Id(c.name)).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
			jen.Type().Id("Alias").Id(c.name),
			jen.Id("t").Op(":=").Op("&").Struct(
				tmpFields...,
			).Values(jen.Dict{
				jen.Id("Alias"): jen.Parens(jen.Op("*").Id("Alias")).Call(jen.Id("o")),
			}),
			jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Return().Err(),
			),
			(&jen.Statement{}).Add(bodyCodes...),
		).Line()
	}
	return code
}

// abstractObject は、共通の性質を持つspecificObject又はabstractObjectの総称です
// （例：File）
// 生成されるGoコードではinterface（共通するフィールドがある場合はstructも）で表現されます
type abstractObject struct {
	objectCommon
	specifiedBy   string // "type", "object" など、バリアントを識別するためのプロパティ
	fieldsComment string
	variants      []objectCoder
}

// addVariant は指定したobjectCoderをこのインターフェイスのバリアントとして登録し、code()に以下のことを行わせます
// - バリアントに対してインターフェイスメソッドを実装
// - JSONメッセージからこのインターフェイスの適切なバリアントを作成するUnmarshalerを作成
func (c *abstractObject) addVariant(variant objectCoder) {
	variant.addParent(c)
	c.variants = append(c.variants, variant)
}

func (c *abstractObject) addFields(fields ...coder) objectCoder {
	c.objectCommon.addFields(fields...)
	return c
}

func (c *abstractObject) code() jen.Code {
	code := jen.Line()

	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	// インターフェイス本体とisメソッド
	{
		methods := []jen.Code{
			jen.Id("is" + c.name).Params(),
		}
		for _, a := range c.parents {
			methods = append(methods, jen.Id("is"+a.name).Params())
		}
		code.Type().Id(c.name).Interface(methods...).Line()
	}

	// 共通フィールド
	// TODO 不要な共通フィールドを省略
	// if len(c.fields) != 0 {
	// }
	{
		if c.fieldsComment != "" {
			code.Comment(c.fieldsComment).Line()
		}
		fields := []jen.Code{}
		for _, p := range c.parents {
			fields = append(fields, jen.Id(strcase.LowerCamelCase(p.name)+"Common"))
		}
		for _, f := range c.fields {
			fields = append(fields, f.code())
		}
		code.Type().Id(strcase.LowerCamelCase(c.name) + "Common").Struct(fields...).Line()
	}

	// バリアントにisメソッドを実装
	cases := []jen.Code{}

	for _, variant := range c.variants {
		sf := variant.getSpecifyingField(c.specifiedBy)
		switch variant := variant.(type) {
		case *specificObject:
			cases = append(cases, jen.Case(jen.Lit(`"`+sf.value+`"`)).Id("u").Dot("value").Op("=").Op("&").Id(variant.getName()).Values())
		case *abstractObject:
			cases = append(cases,
				jen.Case(jen.Lit(`"`+sf.value+`"`)).Id("t").Op(":=").Op("&").Id(strcase.LowerCamelCase(variant.getName())+"Unmarshaler").Values(),
				jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(jen.Return().Err()),
				jen.Id("u").Dot("value").Op("=").Id("t").Dot("value"),
				jen.Return().Nil(),
			)
		default:
			panic(fmt.Sprintf("%#v", variant))
		}
	}

	cases = append(cases, jen.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("unknown type: %s"), jen.String().Call(jen.Id("data")))))

	// variant Unmarshaler
	if len(c.variants) != 0 {
		name := strcase.LowerCamelCase(c.name) + "Unmarshaler"
		code.Line().Type().Id(name).Struct(
			jen.Id("value").Id(c.name),
		).Line().Func().Params(jen.Id("u").Op("*").Id(name)).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
			jen.Switch().String().Call(jen.Id("getRawProperty").Call(jen.Id("data"), jen.Lit(c.specifiedBy))).Block(cases...),
			jen.Return().Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("u").Dot("value")),
		).Line()
	}

	return code
}
