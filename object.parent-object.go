package notion

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/parent-object

/*
Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the "parent". Parent information is represented by a consistent parent object throughout the API.

Parenting rules:
* Pages can be parented by other pages, databases, blocks, or by the whole workspace.
* Blocks can be parented by pages, databases, or blocks.
* Databases can be parented by pages, blocks, or by the whole workspace.
*/
type Parent interface {
	isParent()
}

type parentUnmarshaler struct {
	value Parent
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *parentUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	t := struct {
		DatabaseId json.RawMessage `json:"database_id"`
		PageId     json.RawMessage `json:"page_id"`
		Workspace  json.RawMessage `json:"workspace"`
		BlockId    json.RawMessage `json:"block_id"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	switch {
	case t.DatabaseId != nil:
		u.value = &DatabaseParent{}
	case t.PageId != nil:
		u.value = &PageParent{}
	case t.Workspace != nil:
		u.value = &WorkspaceParent{}
	case t.BlockId != nil:
		u.value = &BlockParent{}
	default:
		return fmt.Errorf("unmarshal Parent: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *parentUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// Database parent
type DatabaseParent struct {
	Type       alwaysDatabaseId `json:"type"`
	DatabaseId uuid.UUID        `json:"database_id"` // The ID of the database that this page belongs to.
}

func (_ *DatabaseParent) isParent() {}

// Page parent
type PageParent struct {
	Type   alwaysPageId `json:"type"`
	PageId uuid.UUID    `json:"page_id"` // The ID of the page that this page belongs to.
}

func (_ *PageParent) isParent() {}

/*
Workspace parent
A page with a workspace parent is a top-level page within a Notion workspace. The parent property is an object containing the following keys:
*/
type WorkspaceParent struct {
	Type      alwaysWorkspace `json:"type"`
	Workspace bool            `json:"workspace"` // Always true.
}

func (_ *WorkspaceParent) isParent() {}

/*
Block parent
A page may have a block parent if it is created inline in a chunk of text, or is located beneath another block like a toggle or bullet block. The parent property is an object containing the following keys:
*/
type BlockParent struct {
	Type    alwaysBlockId `json:"type"`
	BlockId uuid.UUID     `json:"block_id"` // The ID of the page that this page belongs to.
}

func (_ *BlockParent) isParent() {}
