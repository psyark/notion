package notion

import uuid "github.com/google/uuid"

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/parent-object

/*
Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the "parent". Parent information is represented by a consistent parent object throughout the API.

Parenting rules:
* Pages can be parented by other pages, databases, blocks, or by the whole workspace.
* Blocks can be parented by pages, databases, or blocks.
* Databases can be parented by pages, blocks, or by the whole workspace.
*/
type Parent interface{}

/*
Database parent

{
  "type": "database_id",
  "database_id": "d9824bdc-8445-4327-be8b-5b47500af6ce"
}
*/
type DatabaseParent struct {
	Type        string    `always:"database_id" json:"type"` // Always "database_id".
	Database_id uuid.UUID `json:"database_id"`               // The ID of the database that this page belongs to.
}

/*
Page parent

{
  "type": "page_id",
	"page_id": "59833787-2cf9-4fdf-8782-e53db20768a5"
}
*/
type PageParent struct {
	Type    string    `always:"page_id" json:"type"` // Always "page_id".
	Page_id uuid.UUID `json:"page_id"`               // The ID of the page that this page belongs to.
}

/*
Workspace parent
A page with a workspace parent is a top-level page within a Notion workspace. The parent property is an object containing the following keys:

{
	"type": "workspace",
	"workspace": true
}
*/
type WorkspaceParent struct {
	Type      string `always:"workspace" json:"type"` // Always "workspace".
	Workspace bool   `json:"workspace"`               // Always true.
}

/*
Block parent
A page may have a block parent if it is created inline in a chunk of text, or is located beneath another block like a toggle or bullet block. The parent property is an object containing the following keys:

{
	"type": "block_id",
	"block_id": "7d50a184-5bbe-4d90-8f29-6bec57ed817b"
}
*/
type BlockParent struct {
	Type     string    `always:"block_id" json:"type"` // Always "block_id".
	Block_id uuid.UUID `json:"block_id"`               // The ID of the page that this page belongs to.
}
