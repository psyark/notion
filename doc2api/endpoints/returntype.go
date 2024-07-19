package endpoints

import (
	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// エンドポイントの返却値の型を出力するためのインターフェイスです

type ReturnType interface {
	Type() jen.Code     // Type はメソッドの戻り型のコードです
	Accessor() jen.Code // Accessor はアンマーシャラーから値を取り出す関数のコードです
}

var _ = []ReturnType{
	StructRef(""),
	Interface(""),
	GenericStructRef{},
}

// StructRef は、エンドポイントが構造体の参照を返す際に使います
type StructRef string

func (r StructRef) Type() jen.Code {
	return jen.Op("*").Id(string(r))
}
func (r StructRef) Accessor() jen.Code {
	return jen.Func().Params(jen.Id("u").Add(r.Type())).Add(r.Type()).Block(
		jen.Return().Id("u"),
	)
}

// Interface は、エンドポイントがインターフェイスを返す際に使います
type Interface string

func (r Interface) Type() jen.Code {
	return jen.Id(string(r))
}
func (r Interface) Accessor() jen.Code {
	return jen.Func().Params(jen.Id("u").Op("*").Id(strcase.LowerCamelCase(string(r) + "Unmarshaler"))).Add(r.Type()).Block(
		jen.Return().Id("u").Dot("value"),
	)
}

type GenericStructRef struct {
	Name           string
	GenericTypeArg string
}

func (g GenericStructRef) Type() jen.Code {
	code := jen.Op("*").Id(string(g.Name))
	if g.GenericTypeArg != "" {
		code.Index(jen.Id(g.GenericTypeArg))
	}
	return code
}

func (g GenericStructRef) Accessor() jen.Code {
	return jen.Func().Params(jen.Id("u").Add(g.Type())).Add(g.Type()).Block(
		jen.Return().Id("u"),
	)
}
