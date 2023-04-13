package notion

import (
	"encoding/json"
	"fmt"
)

// Code generated by notion.doc2api; DO NOT EDIT.

type FileOrEmoji interface {
	isFileOrEmoji()
}

type fileOrEmojiUnmarshaler struct {
	value FileOrEmoji
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *fileOrEmojiUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "type")) {
	case "\"emoji\"":
		u.value = &Emoji{}
	case "\"external\"":
		u.value = &ExternalFile{}
	case "\"file\"":
		u.value = &NotionHostedFile{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

/*

Endpoints that return lists of objects support cursor-based pagination requests. By default, Notion returns ten items per API call. If the number of items in a response from a support endpoint exceeds the default, then an integration can use pagination to request a specific set of the results and/or to limit the number of returned items.
*/
type Pagination interface {
	isPagination()
}

/*

If an endpoint supports pagination, then the response object contains the below fields.
*/
type paginationCommon struct {
	HasMore    bool   `json:"has_more"`             // Whether the response includes the end of the list. false if there are no more results. Otherwise, true.
	NextCursor string `json:"next_cursor"`          // A string that can be used to retrieve the next page of results by passing the value as the start_cursor parameter to the same endpoint.  Only available when has_more is true.
	Object     string `always:"list" json:"object"` // The constant string "list".
}

type paginationUnmarshaler struct {
	value Pagination
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *paginationUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "type")) {
	case "\"property_item\"":
		u.value = &PropertyItemPagination{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

type PropertyItemOrPropertyItemPagination interface {
	isPropertyItemOrPropertyItemPagination()
}

type propertyItemOrPropertyItemPaginationUnmarshaler struct {
	value PropertyItemOrPropertyItemPagination
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "object" field of the message.
*/
func (u *propertyItemOrPropertyItemPaginationUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "object")) {
	case "\"property_item\"":
		t := &propertyItemUnmarshaler{}
		if err := t.UnmarshalJSON(data); err != nil {
			return err
		}
		u.value = t.value
		return nil
	case "\"list\"":
		u.value = &PropertyItemPagination{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}
