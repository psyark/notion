package notion

import (
	"encoding/json"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/parent-object

// Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the "parent". Parent information is represented by a consistent parent object throughout the API.
type Parent struct {
	Type       string    `json:"type"`
	DatabaseId uuid.UUID `json:"database_id"` // The ID of the database that this page belongs to.
	PageId     uuid.UUID `json:"page_id"`     // The ID of the page that this page belongs to.
	Workspace  bool      `json:"workspace"`   // Always true.
	BlockId    uuid.UUID `json:"block_id"`    // The ID of the page that this page belongs to.
}

func (o Parent) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		// TODO
	}
	type Alias Parent
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"block_id":    o.Type == "block_id",
		"database_id": o.Type == "database_id",
		"page_id":     o.Type == "page_id",
		"workspace":   o.Type == "workspace",
	}
	return omitFields(data, visibility)
}
