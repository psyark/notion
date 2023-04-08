package notion

import (
	"encoding/json"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/page

/*
The Page object contains the page property values of a single Notion page.
All pages have a Parent. If the parent is a database, the property values conform to the schema laid out database's properties. Otherwise, the only property value is the title.

Page content is available as blocks. The content can be read using retrieve block children and appended using append block children.
*/
type Page struct {
	Object         string           `always:"page" json:"object"` // Always "page".
	Id             uuid.UUID        `json:"id"`                   // Unique identifier of the page.
	CreatedTime    ISO8601String    `json:"created_time"`         // Date and time when this page was created. Formatted as an ISO 8601 date time string.
	CreatedBy      PartialUser      `json:"created_by"`           // User who created the page.
	LastEditedTime ISO8601String    `json:"last_edited_time"`     // Date and time when this page was updated. Formatted as an ISO 8601 date time string.
	LastEditedBy   PartialUser      `json:"last_edited_by"`       // User who last edited the page.
	Archived       bool             `json:"archived"`             // The archived status of the page.
	Icon           FileOrEmoji      `json:"icon"`                 // Page icon.
	Cover          File             `json:"cover"`                // Page cover image.
	Properties     PropertyValueMap `json:"properties"`           // Property values of this page. As of version 2022-06-28, properties only contains the ID of the property; in prior versions properties contained the values as well.  If parent.type is "page_id" or "workspace", then the only valid key is title.  If parent.type is "database_id", then the keys and values of this field are determined by the properties  of the database this page belongs to.  key string Name of a property as it appears in Notion.  value object See Property value object.
	Parent         Parent           `json:"parent"`               // Information about the page's parent. See Parent object.
	Url            string           `json:"url"`                  // The URL of the Notion page.
}

func (o *Page) UnmarshalJSON(data []byte) error {
	o.Icon = newFileOrEmoji(getRawProperty(data, "icon"))
	o.Cover = newFile(getRawProperty(data, "cover"))
	o.Parent = newParent(getRawProperty(data, "parent"))
	type Alias Page
	return json.Unmarshal(data, (*Alias)(o))
}
