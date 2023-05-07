package notion

import (
	"encoding/json"
	"fmt"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/intro

/*
Pagination
Endpoints that return lists of objects support cursor-based pagination requests. By default, Notion returns ten items per API call. If the number of items in a response from a support endpoint exceeds the default, then an integration can use pagination to request a specific set of the results and/or to limit the number of returned items.
*/
type Pagination interface {
	isPagination()
	GetHasMore() bool
	GetNextCursor() *string
	GetObject() alwaysList
}

/*

If an endpoint supports pagination, then the response object contains the below fields.
*/
type PaginationCommon struct {
	HasMore    bool       `json:"has_more"`    // Whether the response includes the end of the list. false if there are no more results. Otherwise, true.
	NextCursor *string    `json:"next_cursor"` // A string that can be used to retrieve the next page of results by passing the value as the start_cursor parameter to the same endpoint.  Only available when has_more is true.
	Object     alwaysList `json:"object"`      // The constant string "list".
}

func (c *PaginationCommon) GetHasMore() bool {
	return c.HasMore
}
func (c *PaginationCommon) GetNextCursor() *string {
	return c.NextCursor
}
func (c *PaginationCommon) GetObject() alwaysList {
	return c.Object
}

type paginationUnmarshaler struct {
	value Pagination
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *paginationUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	t := struct {
		Block          json.RawMessage `json:"block"`
		Comment        json.RawMessage `json:"comment"`
		Database       json.RawMessage `json:"database"`
		Page           json.RawMessage `json:"page"`
		PageOrDatabase json.RawMessage `json:"page_or_database"`
		PropertyItem   json.RawMessage `json:"property_item"`
		User           json.RawMessage `json:"user"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	switch {
	case t.Block != nil:
		u.value = &BlockPagination{}
	case t.Comment != nil:
		u.value = &CommentPagination{}
	case t.Database != nil:
		u.value = &DatabasePagination{}
	case t.Page != nil:
		u.value = &PagePagination{}
	case t.PageOrDatabase != nil:
		u.value = &PageOrDatabasePagination{}
	case t.PropertyItem != nil:
		u.value = &PropertyItemPagination{}
	case t.User != nil:
		u.value = &UserPagination{}
	default:
		return fmt.Errorf("unmarshal Pagination: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *paginationUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type BlockPagination struct {
	PaginationCommon
	Type    alwaysBlock `json:"type"`
	Block   struct{}    `json:"block"`
	Results []Block     `json:"results"`
}

func (_ *BlockPagination) isPagination() {}

type CommentPagination struct {
	PaginationCommon
	Type    alwaysComment `json:"type"`
	Comment struct{}      `json:"comment"`
	Results CommentList   `json:"results"`
}

func (_ *CommentPagination) isPagination() {}

type DatabasePagination struct {
	PaginationCommon
	Type     alwaysDatabase `json:"type"`
	Database struct{}       `json:"database"`
	Results  []Database     `json:"results"`
}

func (_ *DatabasePagination) isPagination() {}

type PagePagination struct {
	PaginationCommon
	Type    alwaysPage `json:"type"`
	Page    struct{}   `json:"page"`
	Results []Page     `json:"results"`
}

func (_ *PagePagination) isPagination() {}

type PageOrDatabasePagination struct {
	PaginationCommon
	Type           alwaysPageOrDatabase `json:"type"`
	PageOrDatabase struct{}             `json:"page_or_database"`
	Results        PageOrDatabaseList   `json:"results"`
}

func (_ *PageOrDatabasePagination) isPagination() {}

type PropertyItemPagination struct {
	PaginationCommon
	Type         alwaysPropertyItem    `json:"type"`
	PropertyItem PaginatedPropertyInfo `json:"property_item"` // A constant string that represents the type of the objects in results.
	Results      []PropertyItem        `json:"results"`
}

func (_ *PropertyItemPagination) isPropertyItemOrPropertyItemPagination() {}
func (_ *PropertyItemPagination) isPagination()                           {}

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
	Type    alwaysUser `json:"type"`
	User    struct{}   `json:"user"`
	Results []User     `json:"results"`
}

func (_ *UserPagination) isPagination() {}
