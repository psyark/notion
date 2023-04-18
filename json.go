package notion

import (
	"encoding/json"
	"fmt"
)

// getType はobject型のJSONメッセージからtypeキーの文字列を取り出して返却します
func getType(msg json.RawMessage) string {
	t := struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(msg, &t); err != nil {
		panic(fmt.Errorf("%w: %v", err, string(msg)))
	}
	return t.Type
}

// getObject はobject型のJSONメッセージからobjectキーの文字列を取り出して返却します
func getObject(msg json.RawMessage) string {
	t := struct {
		Object string `json:"object"`
	}{}
	if err := json.Unmarshal(msg, &t); err != nil {
		panic(fmt.Errorf("%w: %v", err, string(msg)))
	}
	return t.Object
}
