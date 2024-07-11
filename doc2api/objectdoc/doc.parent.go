package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	var parent *adaptiveObject

	registerTranslator(
		"https://developers.notion.com/reference/parent-object",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the \"parent\". Parent information is represented by a consistent parent object throughout the API.",
			}, func(e blockElement) {
				parent = b.addAdaptiveObject("Parent", "type", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "General parenting rules:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Pages can be parented by other pages, databases, blocks, or by the whole workspace.Blocks can be parented by pages, databases, or blocks.Databases can be parented by pages, blocks, or by the whole workspace.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧These parenting rules reflect the possible response you may receive when retrieving information about pages, databases, and blocks via Notion’s REST API. If you are creating new pages, databases, or blocks via Notion’s public REST API, the parenting rules may vary. For example, the parent of a database currently must be a page if it is created via the REST API.Refer to the API reference documentation for creating pages, databases, and blocks for more information on current parenting rules.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Database parent */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Database parent",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Always \"database_id\".",
				ExampleValue: `"database_id"`,
				Property:     "type",
				Type:         "string",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The ID of the database that this page belongs to.",
				ExampleValue: `"b8595b75-abd1-4cad-8dfe-f935a8ef57cb"`,
				Property:     "database_id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {
				parent.addAdaptiveFieldWithType("database_id", e.Description, UUID)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"database_id\",\n  \"database_id\": \"d9824bdc-8445-4327-be8b-5b47500af6ce\"\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Parent", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Page parent */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Page parent",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Always \"page_id\".",
				ExampleValue: `"page_id"`,
				Property:     "type",
				Type:         "string",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The ID of the page that this page belongs to.",
				ExampleValue: `"59833787-2cf9-4fdf-8782-e53db20768a5"`,
				Property:     "page_id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {
				parent.addAdaptiveFieldWithType("page_id", e.Description, UUID)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"page_id\",\n\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Parent", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Workspace parent */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Workspace parent",
			}, func(e blockElement) {
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A page with a workspace parent is a top-level page within a Notion workspace. The parent property is an object containing the following keys:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Always \"workspace\".",
				ExampleValue: `"workspace"`,
				Property:     "type",
				Type:         "type",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Always true.",
				ExampleValue: "true",
				Property:     "workspace",
				Type:         "boolean",
			}, func(e parameterElement) {
				parent.addAdaptiveFieldWithType("workspace", e.Description, jen.Bool())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n\t\"type\": \"workspace\",\n\t\"workspace\": true\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Parent", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Block parent */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Block parent",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A page may have a block parent if it is created inline in a chunk of text, or is located beneath another block like a toggle or bullet block. The parent property is an object containing the following keys:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Always \"block_id\".",
				ExampleValue: `"block_id"`,
				Property:     "type",
				Type:         "type",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The ID of the page that this page belongs to.",
				ExampleValue: `"ea29285f-7282-4b00-b80c-32bdbab50261"`,
				Property:     "block_id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {
				parent.addAdaptiveFieldWithType("block_id", e.Description, UUID)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n\t\"type\": \"block_id\",\n\t\"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\"\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Parent", e.Text)
			})
		},
	)
}
