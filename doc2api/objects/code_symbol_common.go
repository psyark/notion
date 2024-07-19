package objects

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// TODO abstractObject は廃止されたのでは？コメントを修正する
// TODO 名前を考える SimpleObject？

// ConcreteObject はAPIのjson応答に実際に出現する具体的なオブジェクトです。
// これには以下の2パターンがあり、それぞれ次のような性質を持ちます
//
// (1) abstractObject の一種として出現するもの (derived object / specific object)
// - parent が存在します
// - discriminatorValue が設定されています （例：type="external" である ExternalFile）
//   - ただし、設定されていない一部の例外（PartialUser）があります
//
// (2) 他のオブジェクト固有のデータ
// （例：Annotations, PersonData）
//
// 生成されるGoコードではstructポインタで表現されます
type ObjectCommon struct {
	name_   string
	comment string
	fields  []fieldRenderer

	// unions は自分が所属するunionObjectです。
	// objectCommonを継承する各クラスは、symbolCode メソッド中でこのunionのisメソッドを実装する必要があります
	unions []*UnionObject
}

func (o *ObjectCommon) name() string {
	return o.name_
}
func (o *ObjectCommon) AddComment(comment string) {
	if o.comment != "" {
		o.comment += "\n\n"
	}
	o.comment += strings.TrimSuffix(strings.TrimPrefix(comment, "\n"), "\n")
}

// TODO この関数は文字列を渡すだけで良いのでは？
func (o *ObjectCommon) AddToUnion(union *UnionObject) {
	o.unions = append(o.unions, union)
	union.members = append(union.members, o)
}

func (o *ObjectCommon) AddFields(fields ...fieldRenderer) *ObjectCommon {
	o.fields = append(o.fields, fields...)
	return o
}

// 指定した discriminatorKey（"type" または "object"） に対してこのオブジェクトが持つ固有の値（"external" など）を返す
// abstractがderivedを見分ける際のロジックではこれを使わない戦略へ移行しているが
// unionがmemberを見分ける際には依然としてこの方法しかない
func (o *ObjectCommon) getDiscriminatorValues(discriminatorKey string) []string {
	for _, f := range o.fields {
		if f, ok := f.(*DiscriminatorField); ok && f.name == discriminatorKey {
			return []string{f.value}
		}
	}
	return nil
}

func (o *ObjectCommon) code(c *Converter) jen.Code {
	code := &jen.Statement{}
	if o.comment != "" {
		code.Comment(o.comment).Line()
	}

	code.Type().Id(o.name_).StructFunc(func(g *jen.Group) {
		for _, f := range o.fields {
			g.Add(f.renderField())
		}
	}).Line()

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(o.fieldUnmarshalerCode(c))

	for _, union := range o.unions {
		code.Func().Params(jen.Id(o.name())).Id("is" + union.name()).Params().Block().Line()
	}

	return code
}

func (o *ObjectCommon) fieldUnmarshalerCode(c *Converter) jen.Code {
	code := &jen.Statement{}

	unionFields := []*VariableField{}
	for _, f := range o.fields {
		if f, ok := f.(*VariableField); ok && f.getUnion(c) != nil {
			unionFields = append(unionFields, f)
		}
	}

	if len(unionFields) != 0 {
		code.Comment("UnmarshalJSON assigns the appropriate implementation to interface field(s)").Line()
		code.Func().Params(jen.Id("o").Op("*").Id(o.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().BlockFunc(func(g *jen.Group) {
			g.Type().Id("Alias").Id(o.name())
			g.Id("t").Op(":=").Op("&").StructFunc(func(g *jen.Group) {
				g.Op("*").Id("Alias")
				for _, f := range unionFields {
					g.Id(strcase.UpperCamelCase(f.name)).Id(f.getUnion(c).memberUnmarshalerName()).Tag(map[string]string{"json": f.name})
				}
			}).Values(jen.Dict{
				jen.Id("Alias"): jen.Parens(jen.Op("*").Id("Alias")).Call(jen.Id("o")),
			})
			g.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", o.name())), jen.Err()),
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
