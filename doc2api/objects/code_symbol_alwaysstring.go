package objects

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// TODO 名前を考える。 FixedString？ Discriminator？
type AlwaysString string

func (c AlwaysString) code(_ *Converter) jen.Code {
	code := jen.Type().Id(c.name()).String().Line()
	code.Func().Params(jen.Id("s").Id(c.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().List(jen.Index().Byte().Call(jen.Lit(fmt.Sprintf("%q", string(c)))), jen.Nil()),
	)
	return code
}

func (c AlwaysString) name() string {
	return "always" + strcase.UpperCamelCase(string(c))
}
