package objects

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// UnionInterface は、interfaceで表現されるUnionです。
//
// 類似するUnionStruct とは異なり、Unionのメンバー定義を内包していません。
// 例えば PropertyItemOrPropertyItemPagination のように、
// 互いに関連が低く、ドキュメントのページを跨ぐようなUnionを表現します。
//
// 🚨 アンマーシャリングは透過的に行いません
//   - 出力する型に対して固有の Unmarshaler が生成されます。
//   - UnionInterface が生成する型をフィールドに持つオブジェクトには、
//     そのフィールドをアンマーシャルするための UnmarshalJSON が生成されます。
type UnionInterface struct {
	nameImpl
	discriminator string // "type" や "object" など
}

func (u *UnionInterface) code(c *Converter) jen.Code {
	// インターフェイス本体
	code := jen.Type().Id(u.name()).Interface(jen.Id("is" + u.name()).Params()).Line().Line()
	// Unmarshaler
	code.Type().Id(u.memberUnmarshalerName()).Struct(
		jen.Id("value").Id(u.name()),
	).Line()

	code.Func().Params(jen.Id("u").Id(u.memberUnmarshalerName())).Id("getValue").Params().Id(u.name()).Block(
		jen.Return().Id("u").Dot("value"),
	).Line()
	code.Comment(fmt.Sprintf("UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance\naccording to the %q field of the message.", u.discriminator)).Line()
	code.Func().Params(jen.Id("u").Op("*").Id(u.memberUnmarshalerName())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.If(jen.String().Call(jen.Id("data")).Op("==").Lit("null")).Block(
			jen.Id("u").Dot("value").Op("=").Nil(),
			jen.Return().Nil(),
		),
		jen.Switch(jen.Id("get"+strcase.UpperCamelCase(u.discriminator))).Call(jen.Id("data")).BlockFunc(func(g *jen.Group) {
			slices.SortFunc(c.unionMemberRegistry, func(a, b unionMemberEntry) int {
				return strings.Compare(a.member.name(), b.member.name())
			})
			for _, entry := range c.unionMemberRegistry {
				if entry.union == u {
					g.CaseFunc(func(g *jen.Group) {
						dvs := entry.member.getDiscriminatorValues(u.discriminator)
						if len(dvs) == 0 {
							panic(fmt.Errorf("メンバー %v を判別するための %s の値がありません", entry.member.name(), u.discriminator))
						}
						for _, v := range dvs {
							g.Lit(v)
						}
					})
					if entry.typeArg != "" {
						g.Id("u").Dot("value").Op("=").Op("&").Id(entry.member.name()).Index(jen.Id(entry.typeArg)).Values()
					} else {
						g.Id("u").Dot("value").Op("=").Op("&").Id(entry.member.name()).Values()
					}
				}
			}
			g.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: data has unknown %s field: %%s", u.name(), u.discriminator)), jen.String().Call(jen.Id("data"))))
		}),
		jen.Return().Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("u").Dot("value")),
	).Line().Line()
	code.Func().Params(jen.Id("u").Op("*").Id(u.memberUnmarshalerName())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().Qual("encoding/json", "Marshal").Call(jen.Id("u").Dot("value")),
	).Line()
	return code
}

func (u *UnionInterface) memberUnmarshalerName() string {
	return strcase.LowerCamelCase(u.name()) + "Unmarshaler"
}

// unionInterfaceMember は、UnionInterfaceのメンバーになれるオブジェクトです
type unionInterfaceMember interface {
	Symbol

	// UnionInterfaceのdiscriminatorに対して、そのメンバーが取りうる値のリストを返します。
	//
	// 現在のところ、これが2つ以上の値を返すのは FileOrEmojiに対して Fileが返す "file", "external" です。
	getDiscriminatorValues(discriminator string) []string
	isGeneric() bool
}

var _ = []unionInterfaceMember{
	&UnionStruct{},
	&SimpleObject{},
}
