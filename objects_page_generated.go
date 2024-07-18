// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/page

package notion

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
)

/*
The Page object contains the page property values of a single Notion page.

All pages have a Parent. If the parent is a database, the property values conform to the schema laid out database's properties. Otherwise, the only property value is the title.

Page content is available as blocks. The content can be read using retrieve block children and appended using append block children.
*/
type Page struct {
	Object         alwaysPage               `json:"object"`               // Always "page".
	Id             uuid.UUID                `json:"id"`                   // Unique identifier of the page.
	CreatedTime    ISO8601String            `json:"created_time"`         // Date and time when this page was created. Formatted as an ISO 8601 date time string.
	CreatedBy      User                     `json:"created_by"`           // User who created the page.
	LastEditedTime ISO8601String            `json:"last_edited_time"`     // Date and time when this page was updated. Formatted as an ISO 8601 date time string.
	LastEditedBy   User                     `json:"last_edited_by"`       // User who last edited the page.
	Archived       bool                     `json:"archived"`             // The archived status of the page.
	InTrash        bool                     `json:"in_trash"`             // Whether the page is in Trash.
	Icon           FileOrEmoji              `json:"icon"`                 // Page icon.
	Cover          *File                    `json:"cover"`                // Page cover image.
	Properties     map[string]PropertyValue `json:"properties"`           // Property values of this page. As of version 2022-06-28, properties only contains the ID of the property; in prior versions properties contained the values as well. If parent.type is "page_id" or "workspace", then the only valid key is title. If parent.type is "database_id", then the keys and values of this field are determined by the properties  of the database this page belongs to. key string Name of a property as it appears in Notion. value object See Property value object.
	Parent         Parent                   `json:"parent"`               // Information about the page's parent. See Parent object.
	Url            string                   `json:"url"`                  // The URL of the Notion page.
	PublicUrl      *string                  `json:"public_url"`           // The public page URL if the page has been published to the web. Otherwise, null.
	RequestId      string                   `json:"request_id,omitempty"` // UNDOCUMENTED
}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *Page) UnmarshalJSON(data []byte) error {
	type Alias Page
	t := &struct {
		*Alias
		Icon fileOrEmojiUnmarshaler `json:"icon"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling Page: %w", err)
	}
	o.Icon = t.Icon.value
	return nil
}
