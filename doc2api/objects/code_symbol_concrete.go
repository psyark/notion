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

func (c *ConcreteObject) AddToUnion(union *UnionObject) {
	c.unions = append(c.unions, union)
	union.members = append(union.members, c)
}

func (c *ConcreteObject) AddFields(fields ...fieldCoder) *ConcreteObject {
	if c.discriminatorValue != "" {
		for _, f := range fields {
			if f, ok := f.(*fixedStringField); ok {
				if f.value == c.discriminatorValue {
					panic(fmt.Errorf("%s に自明の fixedStringField %s がaddFieldされました", c.name(), f.value))
				}
			}
		}
	}
	c.fields = append(c.fields, fields...)
	return c
}

func (c *ConcreteObject) code() jen.Code {
	// struct本体
	code := &jen.Statement{}
	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	code.Type().Id(c.name_).StructFunc(func(g *jen.Group) {
		for _, f := range c.fields {
			g.Add(f.fieldCode())
		}
	}).Line()

	// 先祖インターフェイスを実装
	for _, union := range c.unions {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name())).Id("is" + union.name()).Params().Block().Line()
	}

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(c.fieldUnmarshalerCode())

	return code
}
