// Code generated by notion.doc2api; DO NOT EDIT.

package notion

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/psyark/notion/json"
)

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
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch jsoniter.Get(data, "type").ToString() {
	case "emoji":
		u.value = &Emoji{}
	case "file", "external":
		u.value = &File{}
	default:
		return fmt.Errorf("unmarshaling FileOrEmoji: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *fileOrEmojiUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type PageOrDatabase interface {
	isPageOrDatabase()
}

type pageOrDatabaseUnmarshaler struct {
	value PageOrDatabase
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "object" field of the message.
*/
func (u *pageOrDatabaseUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch jsoniter.Get(data, "object").ToString() {
	case "database":
		u.value = &Database{}
	case "page":
		u.value = &Page{}
	default:
		return fmt.Errorf("unmarshaling PageOrDatabase: data has unknown object field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *pageOrDatabaseUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
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
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch jsoniter.Get(data, "object").ToString() {
	case "list":
		u.value = &Pagination[PropertyItem]{}
	case "property_item":
		u.value = &PropertyItem{}
	default:
		return fmt.Errorf("unmarshaling PropertyItemOrPropertyItemPagination: data has unknown object field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *propertyItemOrPropertyItemPaginationUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type alwaysBlock string

func (s alwaysBlock) MarshalJSON() ([]byte, error) {
	return []byte("\"block\""), nil
}

type alwaysDatabase string

func (s alwaysDatabase) MarshalJSON() ([]byte, error) {
	return []byte("\"database\""), nil
}

type alwaysEmoji string

func (s alwaysEmoji) MarshalJSON() ([]byte, error) {
	return []byte("\"emoji\""), nil
}

type alwaysList string

func (s alwaysList) MarshalJSON() ([]byte, error) {
	return []byte("\"list\""), nil
}

type alwaysPage string

func (s alwaysPage) MarshalJSON() ([]byte, error) {
	return []byte("\"page\""), nil
}

type alwaysPropertyItem string

func (s alwaysPropertyItem) MarshalJSON() ([]byte, error) {
	return []byte("\"property_item\""), nil
}

type alwaysUser string

func (s alwaysUser) MarshalJSON() ([]byte, error) {
	return []byte("\"user\""), nil
}
