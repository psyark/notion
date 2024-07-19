package objects

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// TODO Union関係だと分かる名前にする　UnionMember？
// memberCoder はunionObjectのメンバーとして存在できるオブジェクトを作成するためのCoderです
type memberCoder interface {
	Symbol
	getDiscriminatorValues(identifierKey string) []string
	isGeneric() bool
}

var _ = []memberCoder{
	&AdaptiveObject{},
	&SimpleObject{},
}

// TODO コメント書く
// TODO 名前を考える FlexibleUnion？

// UnionObject は、共通の変数に格納される可能性のあるオブジェクトの集合です
//
// interfaceが使われる点やUnmarshalerが生成される点で abstractObjectと似ていますが、
// 以下のような違いがあり、目的に応じて使い分けがされます
// - UnionObject は必ずしも共通のフィールドを必要としない
// - とあるオブジェクトが直接所属する UnionObject の数に制限はない
// - （ドキュメントのページを跨ぐため）常に objects.global.go に書き込まれる
//
// 例えば FileOrEmoji や PropertyItemOrPropertyItemPagination がUnionObjectです
type UnionObject struct {
	namedSymbol
	discriminator string // "type" や "object" など
}

func (u *UnionObject) code(c *Converter) jen.Code {
	// インターフェイス本体
	code := jen.Type().Id(u.name()).Interface(jen.Id("is" + u.name()).Params()).Line().Line()
	// Unmarshaler
	code.Type().Id(u.memberUnmarshalerName()).Struct(
		jen.Id("value").Id(u.name()),
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
							panic(fmt.Errorf("メンバー %v を識別するための %s の値がありません", entry.member.name(), u.discriminator))
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

func (u *UnionObject) memberUnmarshalerName() string {
	return strcase.LowerCamelCase(u.name()) + "Unmarshaler"
}
