package objects

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
}

// TODO この関数は文字列を渡すだけで良いのでは？
func (o *ConcreteObject) AddToUnion(union *UnionObject) {
	o.unions = append(o.unions, union)
	union.members = append(union.members, o)
}

func (o *ConcreteObject) AddFields(fields ...fieldRenderer) *ConcreteObject {
	o.fields = append(o.fields, fields...)
	return o
}
