package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/parent-object",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the \"parent\". Parent information is represented by a consistent parent object throughout the API.\n\nParenting rules:\n* Pages can be parented by other pages, databases, blocks, or by the whole workspace.\n* Blocks can be parented by pages, databases, or blocks.\n* Databases can be parented by pages, blocks, or by the whole workspace.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addAbstractObject("Parent", "type", e.Text)
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Database parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Parent").addDerived(
						b.addSpecificObject("DatabaseParent", e.Text),
					)
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:      "type",
				Type:          "string",
				Description:   `Always "database_id".`,
				ExampleValues: `"database_id"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatabaseParent").addFields(e.asFixedStringField())
					return nil
				},
			}, {
				Property:      "database_id",
				Type:          "string (UUIDv4)",
				Description:   "The ID of the database that this page belongs to.",
				ExampleValues: `"b8595b75-abd1-4cad-8dfe-f935a8ef57cb"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatabaseParent").addFields(e.asField(UUID))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"database_id\",\n  \"database_id\": \"d9824bdc-8445-4327-be8b-5b47500af6ce\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("DatabaseParent").comment += "\n\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Page parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Parent").addDerived(
						b.addSpecificObject("PageParent", e.Text),
					)
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:      "type",
				Type:          "string",
				Description:   `Always "page_id".`,
				ExampleValues: `"page_id"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PageParent").addFields(e.asFixedStringField())
					return nil
				},
			}, {
				Property:      "page_id",
				Type:          "string (UUIDv4)",
				Description:   "The ID of the page that this page belongs to.",
				ExampleValues: `"59833787-2cf9-4fdf-8782-e53db20768a5"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PageParent").addFields(e.asField(UUID))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"page_id\",\n\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("PageParent").comment += "\n\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Workspace parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Parent").addDerived(
						b.addSpecificObject("WorkspaceParent", e.Text),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA page with a workspace parent is a top-level page within a Notion workspace. The parent property is an object containing the following keys:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("WorkspaceParent").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:      "type",
				Type:          "type",
				Description:   `Always "workspace".`,
				ExampleValues: `"workspace"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("WorkspaceParent").addFields(e.asFixedStringField())
					return nil
				},
			}, {
				Property:      "workspace",
				Type:          "boolean",
				Description:   "Always true.",
				ExampleValues: "true",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("WorkspaceParent").addFields(e.asField(jen.Bool()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n\t\"type\": \"workspace\",\n\t\"workspace\": true\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("WorkspaceParent").comment += "\n\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Block parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Parent").addDerived(
						b.addSpecificObject("BlockParent", e.Text),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA page may have a block parent if it is created inline in a chunk of text, or is located beneath another block like a toggle or bullet block. The parent property is an object containing the following keys:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("BlockParent").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:      "type",
				Type:          "type",
				Description:   `Always "block_id".`,
				ExampleValues: `"block_id"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("BlockParent").addFields(e.asFixedStringField())
					return nil
				},
			}, {
				Property:      "block_id",
				Type:          "string (UUIDv4)",
				Description:   "The ID of the page that this page belongs to.",
				ExampleValues: `"ea29285f-7282-4b00-b80c-32bdbab50261"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("BlockParent").addFields(e.asField(UUID))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n\t\"type\": \"block_id\",\n\t\"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("BlockParent").comment += "\n\n" + e.Code
					return nil
				},
			}}},
		},
	})
}
