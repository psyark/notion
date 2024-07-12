package notion

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func omitFields(data []byte, visibility map[string]bool) ([]byte, error) {
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
