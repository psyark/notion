package objects_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

const (
	UNDOCUMENTED = "UNDOCUMENTED"
)

var (
	converter *Converter
	UUID      = jen.Qual("github.com/google/uuid", "UUID")
	NullUUID  = jen.Op("*").Qual("github.com/google/uuid", "UUID")
)

func TestMain(m *testing.M) {
	converter = NewConverter()
	m.Run()
	converter.OutputAllBuilders()
}

func UndocumentedRequestID(b *CodeBuilder) *VariableField {
	return b.NewField(&Parameter{Property: "request_id", Description: UNDOCUMENTED}, jen.String(), OmitEmpty)
}