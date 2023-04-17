package notion

import (
	"encoding/json"
	"fmt"
)

// getType はobject型のJSONメッセージからtypeキーを取り出し、RawMessageとして返却します
func getType(msg json.RawMessage) json.RawMessage {
	t := struct {
		Type json.RawMessage `json:"type"`
	}{}
	if err := json.Unmarshal(msg, &t); err != nil {
		panic(fmt.Errorf("%w: %v", err, msg))
	}
	return t.Type
}

// getObject はobject型のJSONメッセージからobjectキーを取り出し、RawMessageとして返却します
func getObject(msg json.RawMessage) json.RawMessage {
	t := struct {
		Object json.RawMessage `json:"object"`
	}{}
	if err := json.Unmarshal(msg, &t); err != nil {
		panic(fmt.Errorf("%w: %v", err, msg))
	}
	return t.Object
}
