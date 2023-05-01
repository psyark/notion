package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// concreteObject はAPIのjson応答に実際に出現する具体的なオブジェクトです。
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
type concreteObject struct {
	objectCommon
	discriminatorValue string

	// parent はこのオブジェクトの派生元です。派生元とは共通のフィールドを提供しているオブジェクトであり、
	// 例えば ExternalFile に対する File を指します。一方、FileOrIcon は unionsとして表現します。
	parent *abstractObject

	// Deprecated:
	typeSpecificObject *concreteObject
	// Deprecated:
	typeSpecificAbstract *abstractObject
}

func (c *concreteObject) setParent(parent *abstractObject) {
	if c.parent != nil {
		panic(fmt.Errorf("👪 %s has two parents: %s vs %s", c.name(), c.parent.name(), parent.name()))
	}
	c.parent = parent
}

// 指定した discriminatorKey（"type" または "object"） に対してこのオブジェクトが持つ固有の値（"external" など）を返す
// abstractがderivedを見分ける際のロジックではこれを使わない戦略へ移行しているが
// unionがmemberを見分ける際には依然としてこの方法しかない
func (c *concreteObject) getDiscriminatorValue(identifierKey string) string {
	for _, f := range c.fields {
		if f, ok := f.(*fixedStringField); ok && f.name == identifierKey {
			return f.value
		}
	}
	if c.parent != nil {
		return c.parent.getDiscriminatorValue(identifierKey)
	}
	return ""
}

func (c *concreteObject) addToUnion(union *unionObject) {
	c.unions = append(c.unions, union)
	union.members = append(union.members, c)
}

func (c *concreteObject) addFields(fields ...fieldCoder) *concreteObject {
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

func (c *concreteObject) symbolCode(b *builder) jen.Code {
	// struct本体
	code := &jen.Statement{}
	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	code.Type().Id(c.name_).StructFunc(func(g *jen.Group) {
		if c.parent != nil && len(c.parent.fields) != 0 {
			g.Id(c.parent.commonObjectName())
		}
		for _, f := range c.fields {
			g.Add(f.fieldCode())
		}
	}).Line()

	// 先祖インターフェイスを実装
	for _, union := range c.unions {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name())).Id("is" + union.name()).Params().Block().Line()
	}
	if c.parent != nil {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name())).Id("is" + c.parent.name()).Params().Block().Line()
		for _, union := range c.parent.unions {
			code.Func().Params(jen.Id("_").Op("*").Id(c.name())).Id("is" + union.name()).Params().Block().Line()
		}
	}

	// 親のスペシャルメソッドを実装
	if c.parent != nil {
		for _, sm := range c.parent.specialMethods {
			code.Add(sm.implementationCode(c))
		}
	}
	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(c.fieldUnmarshalerCode(b))

	return code
}
