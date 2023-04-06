package doc2api

func init() {
	registerConverter(converter{
		url:      "https://developers.notion.com/reference/parent-object",
		fileName: "parent.go",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "Pages, databases, and blocks are either located inside other pages, databases, and blocks, or are located at the top level of a workspace. This location is known as the \"parent\". Parent information is represented by a consistent parent object throughout the API.\n\nParenting rules:\n* Pages can be parented by other pages, databases, blocks, or by the whole workspace.\n* Blocks can be parented by pages, databases, or blocks.\n* Databases can be parented by pages, blocks, or by the whole workspace.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.add(&classInterface{name: "Parent", comment: e.Text})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Database parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(&classStruct{name: "DatabaseParent", comment: e.Text})
					return nil
				},
			},
			&objectDocParametersElement{{
				Description:  "Always \"database_id\".",
				ExampleValue: "",
				Field:        "",
				Property:     "type",
				Type:         "string",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The ID of the database that this page belongs to.",
				ExampleValue: "",
				Field:        "",
				Property:     "database_id",
				Type:         "string (UUIDv4)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"database_id\",\n  \"database_id\": \"d9824bdc-8445-4327-be8b-5b47500af6ce\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Page parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(&classStruct{name: "PageParent", comment: e.Text})
					return nil
				},
			},
			&objectDocParametersElement{{
				Description:  "Always \"page_id\".",
				ExampleValue: "",
				Field:        "",
				Property:     "type",
				Type:         "string",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The ID of the page that this page belongs to.",
				ExampleValue: "",
				Field:        "",
				Property:     "page_id",
				Type:         "string (UUIDv4)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"page_id\",\n\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Workspace parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(&classStruct{name: "WorkspaceParent", comment: e.Text})
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA page with a workspace parent is a top-level page within a Notion workspace. The parent property is an object containing the following keys:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "Always \"workspace\".",
				ExampleValue: "",
				Field:        "",
				Property:     "type",
				Type:         "type",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "Always true.",
				ExampleValue: "",
				Field:        "",
				Property:     "workspace",
				Type:         "boolean",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n\t\"type\": \"workspace\",\n\t\"workspace\": true\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Block parent",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(&classStruct{name: "BlockParent", comment: e.Text})
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA page may have a block parent if it is created inline in a chunk of text, or is located beneath another block like a toggle or bullet block. The parent property is an object containing the following keys:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "Always \"block_id\".",
				ExampleValue: "",
				Field:        "",
				Property:     "type",
				Type:         "type",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The ID of the page that this page belongs to.",
				ExampleValue: "",
				Field:        "",
				Property:     "block_id",
				Type:         "string (UUIDv4)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n\t\"type\": \"block_id\",\n\t\"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
		},
	})
}
