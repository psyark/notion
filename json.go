package notion

import "encoding/json"

// getRawProperty はobject型のJSONメッセージから指定されたキーを取り出し、RawMessageとして返却します
func getRawProperty(msg json.RawMessage, key string) json.RawMessage {
	t := map[string]json.RawMessage{}
	if err := json.Unmarshal(msg, &t); err != nil {
		panic(err)
	}
	return t[key]
}
