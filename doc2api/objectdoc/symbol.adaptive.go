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

// 指定した discriminatorKey（"type" または "object"） に対してこのオブジェクトが持つ固有の値（"external" など）を返す
// abstractがderivedを見分ける際のロジックではこれを使わない戦略へ移行しているが
// unionがmemberを見分ける際には依然としてこの方法しかない
func (o *adaptiveObject) getDiscriminatorValues(discriminatorKey string) []string {
	if o.discriminatorKey == discriminatorKey {
		values := []string{}
		for _, f := range o.fields {
			if f, ok := f.(*field); ok && f.name == f.discriminatorValue {
				values = append(values, f.discriminatorValue)
			}
		}
		return values
	}
	return o.objectCommon.getDiscriminatorValues(discriminatorKey)
}

func (o *adaptiveObject) addFields(fields ...fieldCoder) *adaptiveObject {
	o.fields = append(o.fields, fields...)
	return o
}

// addAdaptiveFieldWithType は任意の型でAdaptiveFieldを追加します
func (o *adaptiveObject) addAdaptiveFieldWithType(discriminatorValue string, comment string, typeCode jen.Code) {
	o.addFields(&field{
		name:               discriminatorValue,
		typeCode:           typeCode,
		comment:            comment,
		discriminatorValue: discriminatorValue,
		omitEmpty:          o.discriminatorKey == "", // Filterなど
	})
}

// addAdaptiveFieldWithEmptyStruct は空のStructでAdaptiveFieldを追加します
func (o *adaptiveObject) addAdaptiveFieldWithEmptyStruct(discriminatorValue string, comment string) {
	o.addAdaptiveFieldWithType(discriminatorValue, comment, jen.Struct())
}

// addAdaptiveFieldWithSpecificObject は専用のConcreteObjectを作成し、その型のAdaptiveFieldを追加します
func (o *adaptiveObject) addAdaptiveFieldWithSpecificObject(discriminatorValue string, comment string, b *builder) *concreteObject {
	dataName := o.name() + strcase.UpperCamelCase(discriminatorValue)
	co := b.addConcreteObject(dataName, comment)
	o.addAdaptiveFieldWithType(discriminatorValue, comment, jen.Op("*").Id(dataName))
	return co
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

	if o.discriminatorKey != "" {
		discriminatorProp := strcase.UpperCamelCase(o.discriminatorKey)
		code.Line().Func().Params(jen.Id("o").Id(o.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
			// type 未設定の場合の自動推定
			jen.If(jen.Id("o").Dot(discriminatorProp).Op("==").Lit("")).Block(
				jen.Switch().BlockFunc(func(g *jen.Group) {
					for _, f := range o.fields {
						if f, ok := f.(*field); ok {
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
					if f, ok := f.(*field); ok {
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
	}

	return code
}
