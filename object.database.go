package notion

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/database

/*
Database objects describe the property schema of a database in Notion. Pages are the items (or children) in a database. Page property values must conform to the property objects laid out in the parent database object.
*/
type Database struct {
	Object         alwaysDatabase `json:"object"`           // Always "database".
	Id             uuid.UUID      `json:"id"`               // Unique identifier for the database.
	CreatedTime    ISO8601String  `json:"created_time"`     // Date and time when this database was created. Formatted as an ISO 8601 date time string.
	CreatedBy      PartialUser    `json:"created_by"`       // User who created the database.
	LastEditedTime ISO8601String  `json:"last_edited_time"` // Date and time when this database was updated. Formatted as an ISO 8601 date time string.
	LastEditedBy   PartialUser    `json:"last_edited_by"`   // User who last edited the database.
	Title          RichTextArray  `json:"title"`            // Name of the database as it appears in Notion. See rich text object) for a breakdown of the properties.
	Description    RichTextArray  `json:"description"`      // Description of the database as it appears in Notion. See rich text object) for a breakdown of the properties.
	Icon           FileOrEmoji    `json:"icon"`             // Page icon.
	Cover          *ExternalFile  `json:"cover"`            // Page cover image.
	Properties     PropertyMap    `json:"properties"`       // Schema of properties for the database as they appear in Notion.  key string The name of the property as it appears in Notion.  value object A Property object.
	Parent         Parent         `json:"parent"`           // Information about the database's parent. See Parent object.
	Url            string         `json:"url"`              // The URL of the Notion database.
	Archived       bool           `json:"archived"`         // The archived status of the  database.
	IsInline       bool           `json:"is_inline"`        // Has the value true if the database appears in the page as an inline block. Otherwise has the value false if the database appears as a child page.
}

func (o *Database) UnmarshalJSON(data []byte) error {
	type Alias Database
	t := &struct {
		*Alias
		Parent parentUnmarshaler `json:"parent"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling Database: %w", err)
	}
	o.Parent = t.Parent.value
	return nil
}
