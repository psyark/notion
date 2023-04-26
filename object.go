package notion

import (
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

func (rta RichTextList) String() string {
	str := ""
	for _, rt := range rta {
		str += rt.GetPlainText()
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
