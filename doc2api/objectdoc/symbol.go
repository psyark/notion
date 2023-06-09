package objectdoc

import (
	"fmt"
	"strings"

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
	&concreteObject{},
	&adaptiveObject{},
	&unionObject{},
	&unmarshalTest{},
	alwaysString(""),
}

type objectCommon struct {
	name_   string
	comment string
	fields  []fieldCoder

	// unions は自分が所属するunionObjectです。
	// objectCommonを継承する各クラスは、symbolCode メソッド中でこのunionのisメソッドを実装する必要があります
	unions []*unionObject
}

func (c *objectCommon) name() string {
	return c.name_
}
func (c *objectCommon) addComment(comment string) {
	if c.comment != "" {
		c.comment += "\n\n"
	}
	c.comment += strings.TrimSuffix(strings.TrimPrefix(comment, "\n"), "\n")
}

// 指定した discriminatorKey（"type" または "object"） に対してこのオブジェクトが持つ固有の値（"external" など）を返す
// abstractがderivedを見分ける際のロジックではこれを使わない戦略へ移行しているが
// unionがmemberを見分ける際には依然としてこの方法しかない
func (c *objectCommon) getDiscriminatorValues(discriminatorKey string) []string {
	for _, f := range c.fields {
		if f, ok := f.(*fixedStringField); ok && f.name == discriminatorKey {
			return []string{f.value}
		}
	}
	return nil
}

func (c *objectCommon) symbolCode(b *builder) jen.Code {
	code := &jen.Statement{}
	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	code.Type().Id(c.name_).StructFunc(func(g *jen.Group) {
		for _, f := range c.fields {
			g.Add(f.fieldCode())
		}
	}).Line()

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(c.fieldUnmarshalerCode(b))

	return code
}

func (c *objectCommon) fieldUnmarshalerCode(b *builder) jen.Code {
	code := &jen.Statement{}

	unionFields := []*field{}
	for _, f := range c.fields {
		if f, ok := f.(*field); ok && f.getUnion() != nil {
			unionFields = append(unionFields, f)
		}
	}

	if len(unionFields) != 0 {
		code.Comment("UnmarshalJSON assigns the appropriate implementation to interface field(s)").Line()
		code.Func().Params(jen.Id("o").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().BlockFunc(func(g *jen.Group) {
			g.Type().Id("Alias").Id(c.name())
			g.Id("t").Op(":=").Op("&").StructFunc(func(g *jen.Group) {
				g.Op("*").Id("Alias")
				for _, f := range unionFields {
					g.Id(strcase.UpperCamelCase(f.name)).Id(f.getUnion().memberUnmarshalerName()).Tag(map[string]string{"json": f.name})
				}
			}).Values(jen.Dict{
				jen.Id("Alias"): jen.Parens(jen.Op("*").Id("Alias")).Call(jen.Id("o")),
			})
			g.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
			)
			for _, f := range unionFields {
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
	return jen.Line().Func().Id(c.name()).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.Id("tests").Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
			for _, t := range c.jsonCodes {
				g.Line().Lit(t)
			}
			g.Line()
		}),
		jen.For().List(jen.Id("_"), jen.Id("wantStr")).Op(":=").Range().Id("tests").Block(
			jen.If(jen.Err().Op(":=").Id("checkUnmarshal").Index(jen.Id(c.targetName)).Call(jen.Id("wantStr")).Op(";").Id("err").Op("!=").Nil()).Block(
				jen.Id("t").Dot("Error").Call(jen.Err()),
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
