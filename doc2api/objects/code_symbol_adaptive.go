package objects

import (
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// TODO 名前を DiscriminatedUnionにする https://typescriptbook.jp/reference/values-types-variables/discriminated-union

// AdaptiveObject は、Block や User に代表される抽象クラスのための
// シンプルで効率的な表現を、従来の abstractObject に代わって提供します
type AdaptiveObject struct {
	SimpleObject
	discriminator string // "type", "object" など、派生を識別するためのフィールド名
}

// 指定した discriminatorKey（"type" または "object"） に対してこのオブジェクトが持つ固有の値（"external" など）を返す
// abstractがderivedを見分ける際のロジックではこれを使わない戦略へ移行しているが
// unionがmemberを見分ける際には依然としてこの方法しかない
func (o *AdaptiveObject) getDiscriminatorValues(discriminator string) []string {
	if o.discriminator == discriminator {
		values := []string{}
		for _, f := range o.fields {
			if f, ok := f.(*VariableField); ok && f.name == f.discriminatorValue {
				values = append(values, f.discriminatorValue)
			}
		}
		return values
	}
	return o.SimpleObject.getDiscriminatorValues(discriminator)
}

// TODO メソッドチェーンやめる
func (o *AdaptiveObject) AddFields(fields ...fieldRenderer) *AdaptiveObject {
	o.fields = append(o.fields, fields...)
	return o
}

// TODO 以下の3つのメソッドは共通化できるのでは？

// AddAdaptiveFieldWithType は任意の型でAdaptiveFieldを追加します
func (o *AdaptiveObject) AddAdaptiveFieldWithType(discriminatorValue string, comment string, typeCode jen.Code) {
	o.AddFields(&VariableField{
		name:               discriminatorValue,
		typeCode:           typeCode,
		comment:            comment,
		discriminatorValue: discriminatorValue,
		omitEmpty:          o.discriminator == "", // Filterなど
	})
}

// AddAdaptiveFieldWithEmptyStruct は空のStructでAdaptiveFieldを追加します
func (o *AdaptiveObject) AddAdaptiveFieldWithEmptyStruct(discriminatorValue string, comment string) {
	o.AddAdaptiveFieldWithType(discriminatorValue, comment, jen.Struct())
}

// AddAdaptiveFieldWithSpecificObject は専用の SimpleObject を作成し、その型のAdaptiveFieldを追加します
func (o *AdaptiveObject) AddAdaptiveFieldWithSpecificObject(discriminatorValue string, comment string, b *CodeBuilder) *SimpleObject {
	dataName := o.name() + strcase.UpperCamelCase(discriminatorValue)
	co := b.AddSimpleObject(dataName, comment)
	o.AddAdaptiveFieldWithType(discriminatorValue, comment, jen.Op("*").Id(dataName))
	return co
}

// TODO この関数は文字列を渡すだけで良いのでは？
func (c *AdaptiveObject) AddToUnion(union *UnionObject) {
	c.unions = append(c.unions, union)
	union.members = append(union.members, c)
}

func (o *AdaptiveObject) code(c *Converter) jen.Code {
	code := &jen.Statement{o.SimpleObject.code(c)}

	// discriminatorに対応するGoのフィールド
	discriminatorProp := strcase.UpperCamelCase(o.discriminator)
	code.Line().Func().Params(jen.Id("o").Id(o.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		// type 未設定の場合の自動推定
		jen.If(jen.Id("o").Dot(discriminatorProp).Op("==").Lit("")).Block(
			jen.Switch().BlockFunc(func(g *jen.Group) {
				for _, f := range o.fields {
					if f, ok := f.(*VariableField); ok {
						if f.discriminatorValue != "" {
							g.Case(jen.Id("defined").Call(jen.Id("o").Dot(strcase.UpperCamelCase(f.name)))).Id("o").Dot(discriminatorProp).Op("=").Lit(f.discriminatorValue)
						}
					}
				}
			}),
		),

		jen.Type().Id("Alias").Id(o.name()),
		jen.List(jen.Id("data"), jen.Err()).Op(":=").Qual("encoding/json", "Marshal").Call(jen.Id("Alias").Call(jen.Id("o"))),
		jen.If(jen.Err().Op("!=").Nil()).Block(jen.Return().List(jen.Nil(), jen.Err())),
		jen.Id("visibility").Op(":=").Map(jen.String()).Bool().Values(jen.DictFunc(func(d jen.Dict) {
			for _, f := range o.fields {
				if f, ok := f.(*VariableField); ok {
					if f.discriminatorNotEmpty {
						d[jen.Lit(f.name)] = jen.Id("o").Dot(discriminatorProp).Op("!=").Lit("")
					} else if f.discriminatorValue != "" {
						d[jen.Lit(f.name)] = jen.Id("o").Dot(discriminatorProp).Op("==").Lit(f.discriminatorValue)
					}
				}
			}
		})),
		jen.Return().Id("omitFields").Call(jen.Id("data"), jen.Id("visibility")),
	)

	return code
}
