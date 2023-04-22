package notion

import (
	"encoding/json"
	"fmt"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/intro

type BlockPagination struct {
	PaginationCommon
	Type alwaysBlock `json:"type"`
}

func (_ *BlockPagination) isPagination() {}

type CommentPagination struct {
	PaginationCommon
	Type alwaysComment `json:"type"`
}

func (_ *CommentPagination) isPagination() {}

type DatabasePagination struct {
	PaginationCommon
	Type alwaysDatabase `json:"type"`
}

func (_ *DatabasePagination) isPagination() {}

type PagePagination struct {
	PaginationCommon
	Type    alwaysPage `json:"type"`
	Page    struct{}   `json:"page"` // An object containing type-specific pagination information. For property_items, the value corresponds to the paginated page property type. For all other types, the value is an empty object.
	Results []Page     `json:"results"`
}

func (_ *PagePagination) isPagination() {}

type PageOrDatabasePagination struct {
	PaginationCommon
	Type alwaysPageOrDatabase `json:"type"`
}

func (_ *PageOrDatabasePagination) isPagination() {}

type PropertyItemPagination struct {
	PaginationCommon
	Type         alwaysPropertyItem    `json:"type"`
	PropertyItem PaginatedPropertyInfo `json:"property_item"` // An object containing type-specific pagination information. For property_items, the value corresponds to the paginated page property type. For all other types, the value is an empty object.
	Results      PropertyItemArray     `json:"results"`
}

func (_ *PropertyItemPagination) isPagination()                           {}
func (_ *PropertyItemPagination) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *PropertyItemPagination) UnmarshalJSON(data []byte) error {
	type Alias PropertyItemPagination
	t := &struct {
		*Alias
		PropertyItem paginatedPropertyInfoUnmarshaler `json:"property_item"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling PropertyItemPagination: %w", err)
	}
	o.PropertyItem = t.PropertyItem.value
	return nil
}

type UserPagination struct {
	PaginationCommon
	Type alwaysUser `json:"type"`
}

func (_ *UserPagination) isPagination() {}
