package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

type coder interface {
	code() jen.Code
}

var _ = []coder{
	&classStruct{},
	&classInterface{},
	&field{},
	&fixedStringField{},
	independentComment(""),
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

func (b *builder) getClassInterface(name string) *classInterface {
	for _, item := range *b {
		if item, ok := item.(*classInterface); ok && item.name == name {
			return item
		}
	}
	return nil
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
	name       string
	comment    string
	fields     []coder
	implements []string // TODO: 廃止
}

func (c *classStruct) addField(f coder) {
	c.fields = append(c.fields, f)
}

func (c *classStruct) code() jen.Code {
	fields := []jen.Code{}
	hasInterface := false
	for _, f := range c.fields {
		fields = append(fields, f.code())
		if f, ok := f.(*field); ok && f.isInterface {
			hasInterface = true
		}
	}
	code := jen.Comment(c.comment).Line().Type().Id(c.name).Struct(fields...).Line()
	for _, ifName := range c.implements {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name)).Id("is" + ifName).Params().Block().Line()
	}

	// フィールドにインターフェイスを含むため、UnmarshalJSONで前処理を行う
	if hasInterface {
		bodyCodes := []jen.Code{}
		for _, f := range c.fields {
			fields = append(fields, f.code())
			if f, ok := f.(*field); ok && f.isInterface {
				interfaceName := (&jen.Statement{f.typeCode}).GoString()
				bodyCodes = append(bodyCodes, jen.Id("o").Dot(strcase.UpperCamelCase(f.name)).Op("=").Id("new"+interfaceName).Call(
					jen.Id("getRawProperty").Call(jen.Id("data"), jen.Lit(f.name)),
				))
			}
		}
		bodyCodes = append(bodyCodes, jen.Type().Id("Alias").Id(c.name))
		bodyCodes = append(bodyCodes, jen.Return().Qual("encoding/json", "Unmarshal").Call(
			jen.Id("data"),
			jen.Parens(jen.Op("*").Id("Alias")).Parens(jen.Id("o")),
		))
		code.Func().Params(jen.Id("o").Op("*").Id(c.name)).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(bodyCodes...).Line()
	}
	return code
}

type field struct {
	name        string
	typeCode    jen.Code
	comment     string
	isInterface bool // このフィールドはinterfaceのため、UnmarshalJSONの作成を促します
}

func (f *field) code() jen.Code {
	goName := strcase.UpperCamelCase(f.name)
	code := jen.Id(goName).Add(f.typeCode)
	if f.name != "" {
		code.Tag(map[string]string{"json": f.name})
	}
	if f.comment != "" {
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
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
		code.Comment(strings.ReplaceAll(f.comment, "\n", " "))
	}
	return code
}

type classInterface struct {
	name     string
	comment  string
	variants []*classStruct
}

// addVariant は指定したclassStructをこのインターフェイスのバリアントとして登録し、code()に以下のことを行わせます
// - バリアントに対してインターフェイスメソッドを実装
// - このインターフェイスの未判別のJSONメッセージから適切なバリアントのポインタを返すnew関数を作成
func (c *classInterface) addVariant(variant *classStruct) {
	c.variants = append(c.variants, variant)
}

func (c *classInterface) code() jen.Code {
	// インターフェイス本体とisメソッド
	code := jen.Comment(c.comment).Line().Type().Id(c.name).Interface(jen.Id("is" + c.name).Params()).Line()

	// バリアントにisメソッドを実装
	cases := []jen.Code{}
	for _, v := range c.variants {
		code.Func().Params(jen.Id("_").Op("*").Id(v.name)).Id("is" + c.name).Params().Block().Line()
		for _, f := range v.fields {
			if f, ok := f.(*fixedStringField); ok {
				cases = append(cases, jen.Case(jen.Lit(`"`+f.value+`"`)).Return().Op("&").Id(v.name).Values())
			}
		}
	}

	// new関数
	if len(c.variants) != 0 {
		code.Line().Func().Id("new"+c.name).Call(jen.Id("msg").Qual("encoding/json", "RawMessage")).Id(c.name).Block(
			jen.Switch().String().Call(jen.Id("getRawProperty").Call(jen.Id("msg"), jen.Lit("type"))).Block(cases...),
			jen.Panic(jen.String().Call(jen.Id("msg"))),
		).Line()
	}
	return code
}

type independentComment string

func (c independentComment) code() jen.Code {
	return jen.Comment(string(c))
}
