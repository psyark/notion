package methoddoc

import "github.com/dave/jennifer/jen"

type MethodCoderType interface {
	New() jen.Code
	Returns() jen.Code
	Access(name string) jen.Code
}

type structRefReturningMethodCoder string
type interfaceReturningMethodCoder string

// Deprecated: Use structRefReturningMethodCoder
type ReturnsStructRef struct {
	Name string
}

func (r *ReturnsStructRef) New() jen.Code {
	return jen.Op("&").Id(r.Name)
}
func (r *ReturnsStructRef) Returns() jen.Code {
	return jen.Op("*").Id(r.Name)
}
func (r *ReturnsStructRef) Access(name string) jen.Code {
	return jen.Id(name)
}

// Deprecated: interfaceReturningMethodCoder
type ReturnsInterface struct {
	Name string
}
