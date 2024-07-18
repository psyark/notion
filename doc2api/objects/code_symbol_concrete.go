package objects

import (
	"fmt"

	"github.com/dave/jennifer/jen"
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
type ConcreteObject struct {
	ObjectCommon
	discriminatorValue string
}

// TODO この関数は文字列を渡すだけで良いのでは？
func (o *ConcreteObject) AddToUnion(union *UnionObject) {
	o.unions = append(o.unions, union)
	union.members = append(union.members, o)
}

func (o *ConcreteObject) AddFields(fields ...fieldRenderer) *ConcreteObject {
	if o.discriminatorValue != "" {
		for _, f := range fields {
			if f, ok := f.(*DiscriminatorField); ok {
				if f.value == o.discriminatorValue {
					panic(fmt.Errorf("%s に自明の fixedStringField %s がaddFieldされました", o.name(), f.value))
				}
			}
		}
	}
	o.fields = append(o.fields, fields...)
	return o
}

func (o *ConcreteObject) code(c *Converter) jen.Code {
	// struct本体
	code := &jen.Statement{}
	if o.comment != "" {
		code.Comment(o.comment).Line()
	}

	code.Type().Id(o.name_).StructFunc(func(g *jen.Group) {
		for _, f := range o.fields {
			g.Add(f.renderField())
		}
	}).Line()

	// 先祖インターフェイスを実装
	for _, union := range o.unions {
		code.Func().Params(jen.Id("_").Op("*").Id(o.name())).Id("is" + union.name()).Params().Block().Line()
	}

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(o.fieldUnmarshalerCode(c))

	return code
}
