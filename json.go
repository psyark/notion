package notion

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// TODO 書き直す
func omitFields(data []byte, visibility map[string]bool) ([]byte, error) {
	// ここでは encoding.json を使う！
	t := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, errors.Wrap(err, "omitFields")
	}
	for k, v := range visibility {
		if !v {
			delete(t, k)
		}
	}
	return json.Marshal(t)
}

// TODO 書き直す
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

// TODO 書き直す
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
