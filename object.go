package notion

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type UndefNullString any

type FileWithCaption any    // TODO
type CommentList any        // TODO
type PageOrDatabaseList any // TODO

type PropertyMap map[string]Property

func (m *PropertyMap) UnmarshalJSON(data []byte) error {
	*m = PropertyMap{}
	type Alias PropertyMap
	return json.Unmarshal(data, (*Alias)(m))
}

type PropertyValueMap map[string]PropertyValue

type PropertyItemOrPropertyItemPaginationMap map[string]PropertyItemOrPropertyItemPagination

func (m *PropertyItemOrPropertyItemPaginationMap) UnmarshalJSON(data []byte) error {
	t := map[string]propertyItemOrPropertyItemPaginationUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("unmarshaling PropertyValueMap: %w", err)
	}
	*m = PropertyItemOrPropertyItemPaginationMap{}
	for k, u := range t {
		(*m)[k] = u.value
	}
	return nil
}

type ISO8601String = string

type URLReference struct {
	URL string `json:"url"`
}

type PageReference struct {
	Id uuid.UUID `json:"id"`
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

func String(rta []RichText) string {
	str := ""
	for _, rt := range rta {
		str += rt.PlainText
	}
	return str
}
