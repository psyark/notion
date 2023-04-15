package objectdoc

import "github.com/dave/jennifer/jen"

// doc.*.go はドキュメントに対応する変換ロジックです

var (
	UUID       = jen.Qual("github.com/google/uuid", "UUID")
	NullString = jen.Qual("gopkg.in/guregu/null.v4", "String")
)
