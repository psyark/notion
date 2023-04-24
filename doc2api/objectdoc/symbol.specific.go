package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// specificObject は本来"type"や"object"キーで区別される各オブジェクトです
// （例：type="external" である ExternalFile）
//
// ですが、現在はspecificObjectの入れ子として存在するデータ構造にも使われています
// （例：Annotations, PersonData）
// TODO 上記を解消し、derivedObjectみたいな名前にする？
//
// 生成されるGoコードではstructポインタで表現されます
type specificObject struct {
	objectCommon
	derivedIdentifierValue string

	// typeObject はこのspecificObjectが そのtype値と同名のフィールドに保持する固有データです
	// Every block object has a key corresponding to the value of type. Under the key is an object with type-specific block information.
	// TODO typeObjectがAbstractだった場合の対応（TemplateMentionData）
	typeObject        objectCommon
	typeObjectMayNull bool
}

func (c *specificObject) addToUnion(union *unionObject) {
	c.unions = append(c.unions, union)
	union.members = append(union.members, c)
}

func (c *specificObject) addFields(fields ...fieldCoder) *specificObject {
	if c.derivedIdentifierValue != "" {
		for _, f := range fields {
			if f, ok := f.(*fixedStringField); ok {
				if f.value == c.derivedIdentifierValue {
					panic(fmt.Errorf("%s に自明の fixedStringField %s がaddFieldされました", c.name(), f.value))
				}
			}
		}
	}
	c.fields = append(c.fields, fields...)
	return c
}

func (c *specificObject) symbolCode(b *builder) jen.Code {
	// typeObjectが使われているならtypeObjectへの参照を追加する
	if len(c.typeObject.fields) != 0 {
		if c.derivedIdentifierValue == "" {
			panic(fmt.Sprintf("タイプが不明です: %v", c.name()))
		}

		var valueOfTypeField *field
		for _, f := range c.fields {
			if f, ok := f.(*field); ok && f.name == c.derivedIdentifierValue {
				valueOfTypeField = f
				break
			}
		}
		if valueOfTypeField == nil {
			if c.typeObjectMayNull {
				c.addFields(&field{name: c.derivedIdentifierValue, typeCode: jen.Op("*").Id(c.name() + "Data")})
			} else {
				c.addFields(&field{name: c.derivedIdentifierValue, typeCode: jen.Id(c.name() + "Data")})
			}
		} else {
			fmt.Printf("👻 %s には %s が存在しますが、自動生成されるため消すことが望ましいです\n", c.name(), valueOfTypeField.name)
		}
	}

	// struct本体
	code := &jen.Statement{}
	code.Add(c.objectCommon.symbolCode(b))

	// インターフェイスを実装
	for _, a := range c.getAncestors() {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name())).Id("is" + a.name()).Params().Block().Line()
	}
	// 親のスペシャルメソッドを実装 TODO リカーシブ
	if c.parent != nil {
		for _, sm := range c.parent.specialMethods {
			code.Add(sm.implementationCode(c))
		}
	}

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(c.fieldUnmarshalerCode(b))

	// type object
	if len(c.typeObject.fields) != 0 {
		c.typeObject.name_ = c.name() + "Data"
		code.Add(c.typeObject.symbolCode(b))
	}

	return code
}
