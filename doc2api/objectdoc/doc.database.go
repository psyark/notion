package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/database",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "Database objects describe the property schema of a database in Notion. Pages are the items (or children) in a database. Page property values must conform to the property objects laid out in the parent database object.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.addConcreteObject("Database", e.Text)
				},
			},
			&objectDocHeadingElement{
				Text:   "All databases",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocCalloutElement{
				Body:   "Properties marked with an * are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. For more information on integration capabilities, see the [capabilities guide](ref:capabilities).",
				Title:  "",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) {},
			},
			&objectDocParametersElement{{
				Field:        "object*",
				Type:         "string",
				Description:  `Always "database".`,
				ExampleValue: `"database"`,
				output: func(e *objectDocParameter, b *builder) {
					e.Field = strings.TrimSuffix(e.Field, "*")
					getSymbol[concreteObject](b, "Database").addFields(e.asFixedStringField())
				},
			}, {
				Field:        "id*",
				Type:         "string (UUID)",
				Description:  "Unique identifier for the database.",
				ExampleValue: `"2f26ee68-df30-4251-aad4-8ddc420cba3d"`,
				output: func(e *objectDocParameter, b *builder) {
					e.Field = strings.TrimSuffix(e.Field, "*")
					getSymbol[concreteObject](b, "Database").addFields(e.asField(UUID))
				},
			}, {
				Field:        "created_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this database was created. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("ISO8601String")))
				},
			}, {
				Field:        "created_by",
				Type:         "Partial User",
				Description:  "User who created the database.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("User")))
				},
			}, {
				Field:        "last_edited_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this database was updated. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T21:49:37.913Z"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("ISO8601String")))
				},
			}, {
				Field:        "last_edited_by",
				Type:         "Partial User",
				Description:  "User who last edited the database.",
				ExampleValue: "{\"object\": \"user\",\"id\": \"45ee8d13-687b-47ce-a5ca-6e2e45548c4b\"}",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("User")))
				},
			}, {
				Field:        "title",
				Type:         "array of rich text objects",
				Description:  "Name of the database as it appears in Notion.\nSee rich text object) for a breakdown of the properties.",
				ExampleValue: "\"title\": [\n        {\n            \"type\": \"text\",\n            \"text\": {\n                \"content\": \"Can I create a URL property\",\n                \"link\": null\n            },\n            \"annotations\": {\n                \"bold\": false,\n                \"italic\": false,\n                \"strikethrough\": false,\n                \"underline\": false,\n                \"code\": false,\n                \"color\": \"default\"\n            },\n            \"plain_text\": \"Can I create a URL property\",\n            \"href\": null\n        }\n    ]",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Index().Id("RichText")))
				},
			}, {
				Field:       "description",
				Type:        "array of rich text objects",
				Description: "Description of the database as it appears in Notion.\nSee rich text object) for a breakdown of the properties.",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Index().Id("RichText")))
				},
			}, {
				Description: "Page icon.",
				Field:       "icon",
				Type:        "File Object or Emoji object",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asInterfaceField("FileOrEmoji"))
				},
			}, {
				Description: "Page cover image.",
				Field:       "cover",
				Type:        "File object (only type of \"external\" is supported currently)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Op("*").Id("File")))
				},
			}, {
				Description: "Schema of properties for the database as they appear in Notion.\n\nkey string\nThe name of the property as it appears in Notion.\n\nvalue object\nA Property object.",
				Field:       "properties*",
				Type:        "object",
				output: func(e *objectDocParameter, b *builder) {
					e.Field = strings.TrimSuffix(e.Field, "*")
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("PropertyMap")))
				},
			}, {
				Description:  "Information about the database's parent. See Parent object.",
				ExampleValue: "{ \"type\": \"page_id\", \"page_id\": \"af5f89b5-a8ff-4c56-a5e8-69797d11b9f8\" }",
				Field:        "parent",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asInterfaceField("Parent"))
				},
			}, {
				Field:        "url",
				Type:         "string",
				Description:  "The URL of the Notion database.",
				ExampleValue: "\"https://www.notion.so/668d797c76fa49349b05ad288df2d136\"",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.String()))
				},
			}, {
				Field:        "archived",
				Type:         "boolean",
				Description:  "The archived status of the  database.",
				ExampleValue: "false",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Bool()))
				},
			}, {
				Field:        "is_inline",
				Type:         "boolean",
				Description:  "Has the value true if the database appears in the page as an inline block. Otherwise has the value false if the database appears as a child page.",
				ExampleValue: "false",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Bool()))
				},
			}},
		},
	})
}
