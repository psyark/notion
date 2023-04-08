package notion

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type ISO8601String = string

type URLReference struct {
	URL string `json:"url"`
}

type PageReference struct {
	Id uuid.UUID `json:"id"`
}

type RichTextArray []RichText

type PropertyValueMap map[string]PropertyValue

func (pvm *PropertyValueMap) UnmarshalJSON(data []byte) error {
	x := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	*pvm = PropertyValueMap{}
	for k, v := range x {
		(*pvm)[k] = newPropertyValue(v)
		fmt.Println(string(v))
		fmt.Printf("%#v\n", (*pvm)[k])
		json.Unmarshal(v, (*pvm)[k])
		fmt.Printf("%#v\n", (*pvm)[k])
	}

	return nil
	// type Alias PropertyValueMap
	// return json.Unmarshal(data, (*Alias)(pvm))
}

// https://developers.notion.com/reference/errors
type Error struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("(%v) %v", e.Code, e.Message)
}
