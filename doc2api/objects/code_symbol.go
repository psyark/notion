package objects

import "github.com/dave/jennifer/jen"

// Symbol はソースコードのトップレベルに置かれる、名前を持つシンボルを表します。
type Symbol interface {
	name() string               // 名前を返します
	code(c *Converter) jen.Code // コードを返します
}

var _ = []Symbol{
	&SimpleObject{},
	&AdaptiveObject{},
	&UnionObject{},
	&UnmarshalTest{},
	DiscriminatorString(""),
}

type namedSymbol struct {
	name_ string
}

func (n *namedSymbol) name() string {
	return n.name_
}
