package objectdoc

import (
	"fmt"
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// objectCoder はオブジェクトを作成するためのCoderです
type objectCoder interface {
	coder
	getName() string
	getFields() []coder
	addFields(...coder) objectCoder
}

var _ = []objectCoder{
	&specificObject{},
	&abstractObject{},
}

type objectCommon struct {
	name    string
	comment string
	fields  []coder
}

func (c *objectCommon) getName() string {
	return c.name
}
func (c *objectCommon) getFields() []coder {
	return c.fields
}

// addFields はこのオブジェクトにフィールドを追加します
// ただし、無名フィールドは有名フィールドより前に追加されます
func (c *objectCommon) addFields(fields ...coder) objectCoder {
	c.fields = append(c.fields, fields...)
	sort.SliceStable(c.fields, func(i, j int) bool {
		isAnon := func(c coder) bool {
			if c, ok := c.(*field); ok && c.name == "" {
				return true
			}
			return false
		}
		return isAnon(c.fields[i]) && !isAnon(c.fields[j])
	})
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

func (c *specificObject) code() jen.Code {
	fields := []jen.Code{}
	hasInterface := false
	for _, f := range c.fields {
		fields = append(fields, f.code())
		if f, ok := f.(*field); ok && f.isInterface {
			hasInterface = true
		}
	}
	code := jen.Comment(c.comment).Line().Type().Id(c.name).Struct(fields...).Line()

	// フィールドにインターフェイスを含むため、UnmarshalJSONで前処理を行う
	if hasInterface {
		tmpFields := []jen.Code{jen.Op("*").Id("Alias")}
		bodyCodes := []jen.Code{}
		for _, f := range c.fields {
			fields = append(fields, f.code())
			if f, ok := f.(*field); ok && f.isInterface {
				interfaceName := (&jen.Statement{f.typeCode}).GoString()
				tmpFields = append(tmpFields, jen.Id(strcase.UpperCamelCase(f.name)).Id(strcase.LowerCamelCase(interfaceName)+"Unmarshaler").Tag(map[string]string{"json": f.name}))
				bodyCodes = append(bodyCodes, jen.Id("o").Dot(strcase.UpperCamelCase(f.name)).Op("=").Id("t").Dot(strcase.UpperCamelCase(f.name)).Dot("value").Line())
			}
		}

		bodyCodes = append(bodyCodes, jen.Return().Nil())

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
	variant.addFields(&field{typeCode: jen.Id(strcase.LowerCamelCase(c.name) + "Common")})
	c.variants = append(c.variants, variant)
}

func (c *abstractObject) code() jen.Code {
	// インターフェイス本体とisメソッド
	code := jen.Comment(c.comment).Line().Type().Id(c.name).Interface(jen.Id("is" + c.name).Params()).Line()

	// 共通フィールド
	// if len(c.fields) != 0 {
	// }
	if c.fieldsComment != "" {
		code.Comment(c.fieldsComment).Line()
	}
	fields := []jen.Code{}
	for _, f := range c.fields {
		fields = append(fields, f.code())
	}
	code.Type().Id(strcase.LowerCamelCase(c.name) + "Common").Struct(fields...).Line()

	// バリアントにisメソッドを実装
	cases := []jen.Code{}
	specifiedBy := c.specifiedBy
	if specifiedBy == "" {
		specifiedBy = "type"
		_, _ = fmt.Printf("%v にspecifiedByが指定されていません\n", c.name)
	}
	for _, v := range c.variants {
		code.Func().Params(jen.Id("_").Op("*").Id(v.getName())).Id("is" + c.name).Params().Block().Line()
		found := false
		for _, f := range v.getFields() {
			if f, ok := f.(*fixedStringField); ok && f.name == specifiedBy {
				cases = append(cases, jen.Case(jen.Lit(`"`+f.value+`"`)).Id("u").Dot("value").Op("=").Op("&").Id(v.getName()).Values())
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Errorf("%s not found for %v", specifiedBy, v.getName()))
		}
	}

	cases = append(cases, jen.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("unknown type: %s"), jen.String().Call(jen.Id("data")))))

	// Unmarshaler
	if len(c.variants) != 0 {
		name := strcase.LowerCamelCase(c.name) + "Unmarshaler"
		code.Line().Type().Id(name).Struct(
			jen.Id("value").Id(c.name),
		).Line().Func().Params(jen.Id("u").Op("*").Id(name)).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
			jen.Switch().String().Call(jen.Id("getRawProperty").Call(jen.Id("data"), jen.Lit(specifiedBy))).Block(cases...),
			jen.Return().Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("u").Dot("value")),
		).Line()
	}

	return code
}
