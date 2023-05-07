package notion

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type FileWithCaption any    // TODO
type CommentList any        // TODO
type PageOrDatabaseList any // TODO

type alwaysTitle string

func (a alwaysTitle) MarshalJSON() ([]byte, error) { return []byte(`"title"`), nil }

type alwaysRichText string

func (a alwaysRichText) MarshalJSON() ([]byte, error) { return []byte(`"rich_text"`), nil }

type alwaysRelation string

func (a alwaysRelation) MarshalJSON() ([]byte, error) { return []byte(`"relation"`), nil }

type alwaysPeople string

func (a alwaysPeople) MarshalJSON() ([]byte, error) { return []byte(`"people"`), nil }

type alwaysRollup string

func (a alwaysRollup) MarshalJSON() ([]byte, error) { return []byte(`"rollup"`), nil }

type PropertyMap map[string]Property

func (m *PropertyMap) UnmarshalJSON(data []byte) error {
	*m = PropertyMap{}
	type Alias PropertyMap
	return json.Unmarshal(data, (*Alias)(m))
}

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
