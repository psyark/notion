package objects

import (
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// UnionStruct は、Block や User に代表される抽象クラスのための
// シンプルで効率的な表現を、従来の abstractObject に代わって提供します

// UnionStruct は、structで表現されるUnionです。
//
// 類似するUnionInterface とは異なり、Unionのメンバー定義を内包しています。
// 例えば Block(Bookmark/Breadcrumb/...), User(Person/Bot) のように、
// 互いに関連が高く、ドキュメントの同じページで完結するようなUnionを表現します。
type UnionStruct struct {
	SimpleObject
	discriminator string // "type", "object" など、派生を識別するためのフィールド名
}

// AddPayloadField は、このUnionStructにペイロードフィールドを追加します。
// ペイロードフィールドとは、以下のようなフィールドです。
// - 指定した discriminatorValue を名前に持ちます。
// - UnionStructのdiscriminatorの値が discriminatorValue のときに有効になります。
func (o *UnionStruct) AddPayloadField(discriminatorValue string, comment string, option addPayloadFieldOption) *SimpleObject {
	field := &VariableField{name: discriminatorValue, comment: comment, discriminatorValue: discriminatorValue}
	o.AddFields(field)
	return option(o, field)
}

type addPayloadFieldOption func(union *UnionStruct, field *VariableField) *SimpleObject

func WithType(code jen.Code) addPayloadFieldOption {
	return func(union *UnionStruct, field *VariableField) *SimpleObject {
		field.typeCode = code
		return nil
	}
}
func WithEmptyStructRef() addPayloadFieldOption {
	return WithType(jen.Op("*").Struct())
}
func WithPayloadObject(b *CodeBuilder) addPayloadFieldOption {
	return func(union *UnionStruct, field *VariableField) *SimpleObject {
		payloadName := union.name() + strcase.UpperCamelCase(field.discriminatorValue)
		payload := b.AddSimpleObject(payloadName, field.comment)
		field.typeCode = jen.Op("*").Qual("github.com/psyark/notion", payloadName)
		return payload
	}
}

func (o *UnionStruct) getDiscriminatorValues(discriminator string) []string {
	// このUnionStructがUnionInterfaceの一部である場合に呼ばれます
	if o.discriminator == discriminator {
		// このUnionStructが、所属するUnionInterfaceと同じdiscriminatorの場合に
		// このブロックが評価されます（現時点では FileOrEmojiのメンバーのFileがそれに該当します）
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

func (o *UnionStruct) code(c *Converter) jen.Code {
	code := &jen.Statement{o.SimpleObject.code(c)}

	// discriminatorに対応するGoのフィールド
	discriminatorProp := strcase.UpperCamelCase(o.discriminator)
	code.Line().Func().Params(jen.Id("o").Add(o.typeCode(false))).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
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

		jen.Type().Id("Alias").Add(o.typeCode(false)),
		jen.List(jen.Id("data"), jen.Err()).Op(":=").Qual("github.com/psyark/notion/json", "Marshal").Call(jen.Id("Alias").Call(jen.Id("o"))),
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
