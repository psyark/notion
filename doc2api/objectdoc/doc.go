package objectdoc

import "github.com/dave/jennifer/jen"

// doc.*.go はドキュメントに対応する変換ロジックです

var (
	UUID     = jen.Qual("github.com/google/uuid", "UUID")
	NullUUID = jen.Op("*").Qual("github.com/google/uuid", "UUID")
	// TODO: omitemptyと相性悪いので代替するか検討
	NullFloat = jen.Qual("gopkg.in/guregu/null.v4", "Float")
)
