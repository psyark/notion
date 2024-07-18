package objects

import (
	"fmt"
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// TODO Union関係だと分かる名前にする　UnionMember？
// memberCoder はunionObjectのメンバーとして存在できるオブジェクトを作成するためのCoderです
type memberCoder interface {
	CodeSymbol
	getDiscriminatorValues(identifierKey string) []string
}

var _ = []memberCoder{
	&AdaptiveObject{},
	&ConcreteObject{},
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
	ObjectCommon
	// TODO discriminatorKey にする
	identifierKey string        // "type" や "object" など
	members       []memberCoder // このUnionのメンバー
}

func (u *UnionObject) code() jen.Code {
	// インターフェイス本体
	code := jen.Type().Id(u.name()).Interface(jen.Id("is" + u.name()).Params()).Line().Line()
	// Unmarshaler
	code.Type().Id(u.memberUnmarshalerName()).Struct(
		jen.Id("value").Id(u.name()),
	).Line()

	code.Comment(fmt.Sprintf("UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance\naccording to the %q field of the message.", u.identifierKey)).Line()
	code.Func().Params(jen.Id("u").Op("*").Id(u.memberUnmarshalerName())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.If(jen.String().Call(jen.Id("data")).Op("==").Lit("null")).Block(
			jen.Id("u").Dot("value").Op("=").Nil(),
			jen.Return().Nil(),
		),
		jen.Switch(jen.Id("get"+strcase.UpperCamelCase(u.identifierKey))).Call(jen.Id("data")).BlockFunc(func(g *jen.Group) {
			sort.Slice(u.members, func(i, j int) bool {
				return u.members[i].name() < u.members[j].name()
			})
			for _, member := range u.members {
				g.CaseFunc(func(g *jen.Group) {
					dvs := member.getDiscriminatorValues(u.identifierKey)
					if len(dvs) == 0 {
						panic(fmt.Errorf("メンバー %v を識別するための %s の値がありません", member.name(), u.identifierKey))
					}
					for _, v := range dvs {
						g.Lit(v)
					}
				})
				g.Id("u").Dot("value").Op("=").Op("&").Id(member.name()).Values()
			}
			g.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: data has unknown %s field: %%s", u.name(), u.identifierKey)), jen.String().Call(jen.Id("data"))))
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
