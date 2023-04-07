package notion

import "github.com/google/uuid"

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
