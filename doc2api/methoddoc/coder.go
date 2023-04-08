package methoddoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// returnTypeCoder はメソッドの戻り型を取り扱うインターフェイスです
type returnTypeCoder interface {
	New() jen.Code
	Returns() jen.Code
	Access(name string) jen.Code
}

var _ = []returnTypeCoder{
	ReturnsStructRef(""),
	ReturnsInterface(""),
}

type ReturnsStructRef string

func (r ReturnsStructRef) New() jen.Code {
	return jen.Op("&").Id(string(r))
}
func (r ReturnsStructRef) Returns() jen.Code {
	return jen.Op("*").Id(string(r))
}
func (r ReturnsStructRef) Access(name string) jen.Code {
	return jen.Id(name)
}

type ReturnsInterface string

func (r ReturnsInterface) New() jen.Code {
	return jen.Op("&").Id(fmt.Sprintf("_%vUnmarshaller", string(r)))
}
func (r ReturnsInterface) Returns() jen.Code {
	return jen.Id(string(r))
}
func (r ReturnsInterface) Access(name string) jen.Code {
	return jen.Id(name).Dot(string(r))
}
