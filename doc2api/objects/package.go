package objects

import "github.com/dave/jennifer/jen"

const packagePath = "github.com/psyark/notion"

func Qual(name string) *jen.Statement {
	return jen.Qual(packagePath, name)
}
