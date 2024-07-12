package methoddoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// returnTypeCoder はメソッドの戻り型を取り扱うインターフェイスです
type returnTypeCoder interface {
	returnType() jen.Code             // returnType はメソッドの戻り型を表すコードです
	unmarshaller() jen.Code           // unmarshaller はアンマーシャラーを生成するためのコードです
	returnValue(name string) jen.Code // returnValue はアンマーシャラーから値を取り出すためのコードです
}

var _ = []returnTypeCoder{
	returnsStructRef(""),
	returnsInterface(""),
}

type returnsStructRef string

func (r returnsStructRef) unmarshaller() jen.Code {
	return jen.Id(string(r))
}
func (r returnsStructRef) returnType() jen.Code {
	return jen.Op("*").Id(string(r))
}
func (r returnsStructRef) returnValue(name string) jen.Code {
	return jen.Id(name)
}

type returnsInterface string

func (r returnsInterface) unmarshaller() jen.Code {
	return jen.Id(fmt.Sprintf("%vUnmarshaler", strcase.LowerCamelCase(string(r))))
}
func (r returnsInterface) returnType() jen.Code {
	return jen.Id(string(r))
}
func (r returnsInterface) returnValue(name string) jen.Code {
	return jen.Id(name).Dot("value")
}
