package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	registerTranslator(
		"https://developers.notion.com/reference/database",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Database objects describe the property schema of a database in Notion. Pages are the items (or children) in a database. Page property values must conform to the property objects laid out in the parent database object.",
			}, func(e blockElement) {
				b.addConcreteObject("Database", e.Text)
			})
		},
		func(c *comparator, b *builder) /* All databases */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "All databases",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘Properties marked with an \\* are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. For more information on integration capabilities, see the capabilities guide.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     `object\`,
				Type:         "string",
				Description:  `Always "database".`,
				ExampleValue: `"database"`,
			}, func(e parameterElement) {
				e.Property = strings.TrimSuffix(e.Property, `\`)
				getSymbol[concreteObject](b, "Database").addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Property:     `id\`,
				Type:         "string (UUID)",
				Description:  "Unique identifier for the database.",
				ExampleValue: `"2f26ee68-df30-4251-aad4-8ddc420cba3d"`,
			}, func(e parameterElement) {
				e.Property = strings.TrimSuffix(e.Property, `\`)
				getSymbol[concreteObject](b, "Database").addFields(e.asField(UUID))
			})
			c.nextMustParameter(parameterElement{
				Property:     "created_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this database was created. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "created_by",
				Type:         "Partial User",
				Description:  "User who created the database.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("User")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "last_edited_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this database was updated. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T21:49:37.913Z"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "last_edited_by",
				Type:         "Partial User",
				Description:  "User who last edited the database.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("User")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "title",
				Type:         "array of rich text objects",
				Description:  "Name of the database as it appears in Notion.  \nSee rich text object) for a breakdown of the properties.",
				ExampleValue: "\"title\": [\n        {\n            \"type\": \"text\",\n            \"text\": {\n                \"content\": \"Can I create a URL property\",\n                \"link\": null\n            },\n            \"annotations\": {\n                \"bold\": false,\n                \"italic\": false,\n                \"strikethrough\": false,\n                \"underline\": false,\n                \"code\": false,\n                \"color\": \"default\"\n            },\n            \"plain_text\": \"Can I create a URL property\",\n            \"href\": null\n        }\n    ]",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "description",
				Type:        "array of rich text objects",
				Description: "Description of the database as it appears in Notion.  \nSee rich text object) for a breakdown of the properties.",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "icon",
				Type:        "File Object or Emoji object",
				Description: "Page icon.",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asUnionField("FileOrEmoji"))
			})
			c.nextMustParameter(parameterElement{
				Property:    "cover",
				Type:        `File object (only type of "external" is supported currently)`,
				Description: "Page cover image.",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Op("*").Id("File")))
			})
			c.nextMustParameter(parameterElement{
				Property:    `properties\*`,
				Type:        "object",
				Description: "Schema of properties for the database as they appear in Notion.  \n  \nkey string  \nThe name of the property as it appears in Notion.  \n  \nvalue object  \nA Property object.",
			}, func(e parameterElement) {
				e.Property = strings.TrimSuffix(e.Property, `\*`)
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Map(jen.String()).Id("Property")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "parent",
				Type:         "object",
				Description:  "Information about the database's parent. See Parent object.",
				ExampleValue: `{ "type": "page_id", "page_id": "af5f89b5-a8ff-4c56-a5e8-69797d11b9f8" }`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Id("Parent")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "url",
				Type:         "string",
				Description:  "The URL of the Notion database.",
				ExampleValue: `"https://www.notion.so/668d797c76fa49349b05ad288df2d136"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "archived",
				Type:         "boolean",
				Description:  "The archived status of the  database.",
				ExampleValue: "false",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_inline",
				Type:         "boolean",
				Description:  "Has the value true if the database appears in the page as an inline block. Otherwise has the value false if the database appears as a child page.",
				ExampleValue: "false",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Database").addFields(e.asField(jen.Bool()))
			})
		},
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧 Maximum schema size recommendationNotion recommends a maximum schema size of 50KB. Updates to database schemas that are too large will be blocked to help maintain database performance.",
			}, func(e blockElement) {})
		},
	)
}
