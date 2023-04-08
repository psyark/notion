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

func (rta *RichTextArray) UnmarshalJSON(data []byte) error {
	x := []json.RawMessage{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	*rta = make([]RichText, len(x))
	for i, m := range x {
		(*rta)[i] = newRichText(m)
	}
	return nil
}

type PropertyValueMap map[string]PropertyValue

func (pvm *PropertyValueMap) UnmarshalJSON(data []byte) error {
	x := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	*pvm = PropertyValueMap{}
	for k, v := range x {
		(*pvm)[k] = newPropertyValue(v)
	}
	return nil
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
