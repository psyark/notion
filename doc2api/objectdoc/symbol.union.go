package objectdoc

import (
	"fmt"
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// memberCoder はunionObjectのメンバーとして存在できるオブジェクトを作成するためのCoderです
type memberCoder interface {
	symbolCoder
	getIdentifierValue(identifierKey string) string
}

var _ = []memberCoder{
	&concreteObject{},
	&abstractObject{},
}

// unionObject は、共通の変数に格納される可能性のあるオブジェクトの集合です
//
// interfaceが使われる点やUnmarshalerが生成される点で abstractObjectと似ていますが、
// 以下のような違いがあり、目的に応じて使い分けがされます
// - unionObject は必ずしも共通のフィールドを必要としない
// - とあるオブジェクトが直接所属する unionObject の数に制限はない
// - （ドキュメントのページを跨ぐため）常に objects.global.go に書き込まれる
//
// 例えば FileOrEmoji や PropertyItemOrPropertyItemPagination がunionObjectです
type unionObject struct {
	objectCommon
	identifierKey string        // "type" や "object" など
	members       []memberCoder // このUnionのメンバー
}

func (u *unionObject) symbolCode(b *builder) jen.Code {
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
			// 識別可能なもの
			canIdentify := []memberCoder{}
			for _, member := range u.members {
				canIdentify = append(canIdentify, u.findCanIdentify(member)...)
			}

			sort.Slice(canIdentify, func(i, j int) bool {
				return canIdentify[i].getIdentifierValue(u.identifierKey) < canIdentify[j].getIdentifierValue(u.identifierKey)
			})
			for _, member := range canIdentify {
				g.Case(jen.Lit(member.getIdentifierValue(u.identifierKey)))
				switch member := member.(type) {
				case *concreteObject:
					g.Id("u").Dot("value").Op("=").Op("&").Id(member.name()).Values()
				case *abstractObject:
					g.Id("t").Op(":=").Op("&").Id(member.derivedUnmarshalerName()).Values()
					g.If(jen.Err().Op(":=").Id("t").Dot("UnmarshalJSON").Call(jen.Id("data"))).Op(";").Err().Op("!=").Nil().Block(jen.Return().Err())
					g.Id("u").Dot("value").Op("=").Id("t").Dot("value")
					g.Return().Nil()
				}
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

func (u *unionObject) findCanIdentify(member memberCoder) []memberCoder {
	result := []memberCoder{}
	switch member := member.(type) {
	case *concreteObject:
		result = append(result, member)
	case *abstractObject:
		if member.derivedIdentifierKey == u.identifierKey {
			for _, member2 := range member.derivedObjects {
				result = append(result, member2)
			}
		} else {
			result = append(result, member)
		}
	default:
		panic(member)
	}
	return result
}

func (u *unionObject) memberUnmarshalerName() string {
	return strcase.LowerCamelCase(u.name()) + "Unmarshaler"
}
