package objectdoc

import (
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// adaptiveObject は、Block や User に代表される抽象クラスのための
// シンプルで効率的な表現を、従来の abstractObject に代わって提供します
type adaptiveObject struct {
	objectCommon
	discriminatorKey string // "type", "object" など、派生を識別するためのプロパティ名
}

func (o *adaptiveObject) addFields(fields ...fieldCoder) *adaptiveObject {
	o.fields = append(o.fields, fields...)
	return o
}

func (c *adaptiveObject) addToUnion(union *unionObject) {
	c.unions = append(c.unions, union)
	union.members = append(union.members, c)
}

func (o *adaptiveObject) symbolCode(b *builder) jen.Code {
	code := &jen.Statement{
		o.objectCommon.symbolCode(b),
	}

	for _, u := range o.unions {
		code.Line().Func().Params(jen.Id("o").Id(o.name())).Id("is" + u.name()).Params().Block()
	}

	// TODO type 未設定の場合の自動推定
	code.Line().Func().Params(jen.Id("o").Id(o.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.If(jen.Id("o").Dot(strcase.UpperCamelCase(o.discriminatorKey)).Op("==").Lit("")).Block(
			jen.Comment("TODO"),
		),
		jen.Type().Id("Alias").Id(o.name()),
		jen.List(jen.Id("data"), jen.Err()).Op(":=").Qual("encoding/json", "Marshal").Call(jen.Id("Alias").Call(jen.Id("o"))),
		jen.If(jen.Err().Op("!=").Nil()).Block(jen.Return().List(jen.Nil(), jen.Err())),
		jen.Id("visibility").Op(":=").Map(jen.String()).Bool().Values(jen.DictFunc(func(d jen.Dict) {
			for _, f := range o.fields {
				if f, ok := f.(*field); ok {
					if f.discriminatorNotEmpty {
						d[jen.Lit(f.name)] = jen.Id("o").Dot(strcase.UpperCamelCase(o.discriminatorKey)).Op("!=").Lit("")
					} else if f.discriminatorValue != "" {
						d[jen.Lit(f.name)] = jen.Id("o").Dot(strcase.UpperCamelCase(o.discriminatorKey)).Op("==").Lit(f.discriminatorValue)
					}
				}
			}
		})),
		jen.Return().Id("omitFields").Call(jen.Id("data"), jen.Id("visibility")),
	)

	return code
}
