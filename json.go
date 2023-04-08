package notion

import "encoding/json"

func getChild(data []byte, childName string) []byte {
	t := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &t); err != nil {
		panic(err)
	}

	return t[childName]
}
