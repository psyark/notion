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

type RichTextArray []RichText

type FileOrEmoji interface {
	isFileOrEmoji()
}

func newFileOrEmoji(data []byte) FileOrEmoji {
	switch string(getChild(data, "type")) {
	case `"file"`:
		return &NotionHostedFile{}
	case `"external"`:
		return &ExternalFile{}
	case `"emoji"`:
		return &Emoji{}
	}
	panic(string(data))
}

func newFile(data []byte) File {
	switch string(getChild(data, "type")) {
	case `"file"`:
		return &NotionHostedFile{}
	case `"external"`:
		return &ExternalFile{}
	}
	panic(string(data))
}

// TODO 自動化
func newParent(data []byte) Parent {
	switch string(getChild(data, "type")) {
	case `"page_id"`:
		return &PageParent{}
	case `"database_id"`:
		return &DatabaseParent{}
	case `"workspace"`:
		return &WorkspaceParent{}
	case `"block_id"`:
		return &BlockParent{}
	}
	panic(string(data))
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
