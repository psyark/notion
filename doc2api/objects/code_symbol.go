package objects

import "github.com/dave/jennifer/jen"

// TODO 名前（Symbolで良くない？）
// CodeSymbol はソースコードのトップレベルに置かれる、名前を持つシンボルを表します。
type CodeSymbol interface {
	name() string   // 名前を返します
	code() jen.Code // コードを返します
}

var _ = []CodeSymbol{
	&ObjectCommon{},
	&ConcreteObject{},
	&AdaptiveObject{},
	&UnionObject{},
	&UnmarshalTest{},
	AlwaysString(""),
}
