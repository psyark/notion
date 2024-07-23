package notion

import (
	"encoding/json"

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
