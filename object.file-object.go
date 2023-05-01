package notion

import "encoding/json"

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/file-object

/*
The Notion API does not yet support uploading files to Notion.

File objects contain data about a file that is uploaded to Notion, or data about an external file that is linked to in Notion.

Page, embed, image, video, file, pdf, and bookmark block types all contain file objects. Icon and cover page object values also contain file objects.

Each file object includes the following fields:

To modify page or database property values that are made from file objects, like `icon`, `cover`, or `files` page property values, use the [update page](https://developers.notion.com/reference/patch-page) or [update database](https://developers.notion.com/reference/update-a-database) endpoints.
*/
type File struct {
	Type     string            `json:"type"`
	Name     string            `json:"name,omitempty"` // undocumented
	File     *NotionHostedFile `json:"file"`           // Notion-hosted files
	External *ExternalFile     `json:"external"`       // External files
}

func (o File) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		// TODO
	}
	type Alias File
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"external": o.Type == "external",
		"file":     o.Type == "file",
	}
	return omitFields(data, visibility)
}

/*

All Notion-hosted files have a type of "file". The corresponding file specific object contains the following fields:
You can retrieve links to Notion-hosted files via the Retrieve block children endpoint.
*/
type NotionHostedFile struct {
	Url        string        `json:"url"`         // An authenticated S3 URL to the file.   The URL is valid for one hour. If the link expires, then you can send an API request to get an updated URL.
	ExpiryTime ISO8601String `json:"expiry_time"` // The date and time when the link expires, formatted as an ISO 8601 date time string.
}

/*
External files
An external file is any URL linked to in Notion that isn’t hosted by Notion. All external files have a type of "external". The corresponding file specific object contains the following fields:

The Notion API supports adding, retrieving, and updating links to external files.
*/
type ExternalFile struct {
	Url string `json:"url"` // A link to the externally hosted content.
}
