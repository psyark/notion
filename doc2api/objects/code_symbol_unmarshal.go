package objects

import (
	"fmt"
	"slices"

	"github.com/dave/jennifer/jen"
)

type UnmarshalTest struct {
	targetName string
	typeArg    string
	jsonCodes  []string
}

func (c *UnmarshalTest) name() string {
	return fmt.Sprintf("Test%s%s_unmarshal", c.targetName, c.typeArg)
}

func (c *UnmarshalTest) code(_ *Converter) jen.Code {
	typeCode := jen.Id(c.targetName)
	if c.typeArg != "" {
		typeCode.Index(jen.Id(c.typeArg))
	}

	return jen.Line().Func().Id(c.name()).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.Id("tests").Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
			slices.Sort(c.jsonCodes) // 並列実行で出力が変わるのを防ぐため
			for _, t := range c.jsonCodes {
				g.Line().Lit(t)
			}
			g.Line()
		}),
		jen.For().List(jen.Id("_"), jen.Id("wantStr")).Op(":=").Range().Id("tests").Block(
			jen.Id("checkUnmarshal").Types(typeCode).Call(jen.Id("t"), jen.Id("wantStr")),
		),
	)
}
