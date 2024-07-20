package objects_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestParent(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/parent-object")

	var parent *UnionStruct

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the \"parent\". Parent information is represented by a consistent parent object throughout the API.",
	}).Output(func(e *Block, b *CodeBuilder) {
		parent = b.AddUnionStruct("Parent", "type", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "General parenting rules:"})
	c.ExpectBlock(&Block{Kind: "List", Text: "Pages can be parented by other pages, databases, blocks, or by the whole workspace.Blocks can be parented by pages, databases, or blocks.Databases can be parented by pages, blocks, or by the whole workspace."})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "🚧These parenting rules reflect the possible response you may receive when retrieving information about pages, databases, and blocks via Notion’s REST API. If you are creating new pages, databases, or blocks via Notion’s public REST API, the parenting rules may vary. For example, the parent of a database currently must be a page if it is created via the REST API.Refer to the API reference documentation for creating pages, databases, and blocks for more information on current parenting rules."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Database parent"})

	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "string",
		Description:  `Always "database_id".`,
		ExampleValue: `"database_id"`,
	})

	c.ExpectParameter(&Parameter{
		Property:     "database_id",
		Type:         "string (UUIDv4)",
		Description:  "The ID of the database that this page belongs to.",
		ExampleValue: `"b8595b75-abd1-4cad-8dfe-f935a8ef57cb"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		parent.AddFields(b.NewField(e, UUID, DiscriminatorValue(e.Property)))
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"database_id\",\n  \"database_id\": \"d9824bdc-8445-4327-be8b-5b47500af6ce\"\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Parent", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Page parent"})
	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "string",
		Description:  `Always "page_id".`,
		ExampleValue: `"page_id"`,
	})

	c.ExpectParameter(&Parameter{
		Property:     "page_id",
		Type:         "string (UUIDv4)",
		Description:  "The ID of the page that this page belongs to.",
		ExampleValue: `"59833787-2cf9-4fdf-8782-e53db20768a5"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		parent.AddFields(b.NewField(e, UUID, DiscriminatorValue(e.Property)))
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"page_id\",\n\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Parent", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Workspace parent"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A page with a workspace parent is a top-level page within a Notion workspace. The parent property is an object containing the following keys:"})

	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "type",
		Description:  `Always "workspace".`,
		ExampleValue: `"workspace"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "workspace",
		Type:         "boolean",
		Description:  "Always true.",
		ExampleValue: "true",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		parent.AddFields(b.NewField(e, jen.Bool(), DiscriminatorValue(e.Property)))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n\t\"type\": \"workspace\",\n\t\"workspace\": true\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Parent", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Block parent"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A page may have a block parent if it is created inline in a chunk of text, or is located beneath another block like a toggle or bullet block. The parent property is an object containing the following keys:"})

	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "type",
		Description:  `Always "block_id".`,
		ExampleValue: `"block_id"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "block_id",
		Type:         "string (UUIDv4)",
		Description:  "The ID of the page that this page belongs to.",
		ExampleValue: `"ea29285f-7282-4b00-b80c-32bdbab50261"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		parent.AddFields(b.NewField(e, UUID, DiscriminatorValue(e.Property)))
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n\t\"type\": \"block_id\",\n\t\"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\"\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Parent", e.Text)
	})
}
