package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
	"github.com/samber/lo"
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
	converter.OutputBindingHelper()
}

func UndocumentedRequestID(b *CodeBuilder) *VariableField {
	return b.NewField(&Parameter{Property: "request_id", Description: UNDOCUMENTED}, jen.String(), OmitEmpty)
}

// rewriteInaccurateExampleJSON は不正確なドキュメントのJSONを書き換えるための関数です
// 上記の目的にのみ使ってください
func rewriteInaccurateExampleJSON(jsonStr string, fn func(data any) any) string {
	var data any
	lo.Must0(json.Unmarshal([]byte(jsonStr), &data))
	data = fn(data)
	return string(lo.Must(json.Marshal(data)))
}
