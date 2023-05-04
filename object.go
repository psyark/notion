package notion

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type FileWithCaption any    // TODO
type CommentList any        // TODO
type PageOrDatabaseList any // TODO

type PropertyMap map[string]Property

func (m *PropertyMap) UnmarshalJSON(data []byte) error {
	*m = PropertyMap{}
	type Alias PropertyMap
	return json.Unmarshal(data, (*Alias)(m))
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

func (pvm PropertyValueMap) Get(id string) PropertyValue {
	for _, v := range pvm {
		if v.GetId() == id {
			return v
		}
	}
	return nil
}
