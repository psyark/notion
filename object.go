package notion

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type FileWithCaption any // TODO

type PropertyMap map[string]Property // for test

func (p *Page) GetProperty(id string) *PropertyValue {
	for _, pv := range p.Properties {
		if pv.Id == id {
			return &pv
		}
	}
	return nil
}

func (p *Pagination) Pages() ([]Page, error) {
	pages := []Page{}
	return pages, json.Unmarshal(p.Results, &pages)
}

func (p *Pagination) Blocks() ([]Block, error) {
	blocks := []Block{}
	return blocks, json.Unmarshal(p.Results, &blocks)
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

func defined(fieldValue any) bool {
	v := reflect.ValueOf(fieldValue)

	switch v.Kind() {
	case reflect.Ptr, reflect.Slice:
		return !v.IsNil()
	case reflect.Struct:
		return false
	case reflect.Array: // UUID
		return !v.IsZero()
	case reflect.Bool:
		return !v.IsZero()
	case reflect.String:
		return !v.IsZero()
	default:
		return !v.IsZero()
		// panic(v.Kind())
	}
}
