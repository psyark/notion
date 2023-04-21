package notion

import (
	"encoding/json"
	"fmt"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/file-object

/*
The Notion API does not yet support uploading files to Notion.

File objects contain data about a file that is uploaded to Notion, or data about an external file that is linked to in Notion.

Page, embed, image, video, file, pdf, and bookmark block types all contain file objects. Icon and cover page object values also contain file objects.

Each file object includes the following fields:

To modify page or database property values that are made from file objects, like `icon`, `cover`, or `files` page property values, use the [update page](https://developers.notion.com/reference/patch-page) or [update database](https://developers.notion.com/reference/update-a-database) endpoints.
*/
type File interface {
	isFile()
}

type fileUnmarshaler struct {
	value File
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *fileUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch getType(data) {
	case "file":
		u.value = &NotionHostedFile{}
	case "external":
		u.value = &ExternalFile{}
	default:
		return fmt.Errorf("unmarshaling File: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *fileUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type Files []File

func (a *Files) UnmarshalJSON(data []byte) error {
	t := []fileUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("unmarshaling Files: %w", err)
	}
	*a = make([]File, len(t))
	for i, u := range t {
		(*a)[i] = u.value
	}
	return nil
}

// Notion-hosted files
type NotionHostedFile struct {
	Type alwaysFile           `json:"type"`
	Name string               `json:"name,omitempty"` // undocumented
	File NotionHostedFileData `json:"file"`
}

func (_ *NotionHostedFile) isFile()        {}
func (_ *NotionHostedFile) isFileOrEmoji() {}

/*

All Notion-hosted files have a type of "file". The corresponding file specific object contains the following fields:
You can retrieve links to Notion-hosted files via the Retrieve block children endpoint.
*/
type NotionHostedFileData struct {
	Url        string        `json:"url"`         // An authenticated S3 URL to the file.   The URL is valid for one hour. If the link expires, then you can send an API request to get an updated URL.
	ExpiryTime ISO8601String `json:"expiry_time"` // The date and time when the link expires, formatted as an ISO 8601 date time string.
}

// External files
type ExternalFile struct {
	Type     alwaysExternal   `json:"type"`
	External ExternalFileData `json:"external"`
}

func (_ *ExternalFile) isFile()        {}
func (_ *ExternalFile) isFileOrEmoji() {}

/*

An external file is any URL linked to in Notion that isn’t hosted by Notion. All external files have a type of "external". The corresponding file specific object contains the following fields:

The Notion API supports adding, retrieving, and updating links to external files.
*/
type ExternalFileData struct {
	Url string `json:"url"` // A link to the externally hosted content.
}
