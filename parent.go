package notion

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

// Database parent
type DatabaseParent struct{}

// Page parent
type PageParent struct{}

// Workspace parent
type WorkspaceParent struct{}

// Block parent
type BlockParent struct{}
