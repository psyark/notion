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

// TODO 自動化
type PropertyValueMap map[string]PropertyValue

func (pvm *PropertyValueMap) UnmarshalJSON(data []byte) error {
	t := map[string]propertyValueUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*pvm = PropertyValueMap{}
	for k, v := range t {
		(*pvm)[k] = v.value
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
