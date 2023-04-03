package doc2api

import (
	"encoding/json"
)

// var enumStringRegex = regexp.MustCompile(`(?m:- "(.+?)"$\n?)`)

// objectDocParameter はドキュメントに書かれているパラメータです
type objectDocParameter struct {
	Name         string
	Type         string
	Description  string
	ExampleValue string `json:"Example value"`
}

func (p *objectDocParameter) UnmarshalJSON(data []byte) error {
	type Alias objectDocParameter

	// 以前は "Property" とドキュメントに書いてあったが、 https://developers.notion.com/reference/parent-object
	// 現在は殆どのドキュメントでは "Field" となっている。 https://developers.notion.com/reference/page-property-values
	// この関数では上記のうちいずれかを Name に代入することで透過的に扱うようにする
	t := struct {
		*Alias
		Property string `json:"Property"`
		Field    string `json:"Field"`
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t.Property != "" {
		t.Name = t.Property
	} else {
		t.Name = t.Field
	}
	return nil
}
