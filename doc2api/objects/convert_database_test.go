package objects_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestDatabase(t *testing.T) {
	t.Parallel()

	converter.FetchDocument("https://developers.notion.com/reference/database").WithScope(
		func(c *DocumentComparator) {

			var database *ConcreteObject

			c.ExpectBlock(&Block{
				Kind: "Paragraph",
				Text: "Database objects describe the property schema of a database in Notion. Pages are the items (or children) in a database. Page property values must conform to the property objects laid out in the parent database object.",
			}).Output(func(e *Block, b *CodeBuilder) {
				database = b.AddConcreteObject("Database", e.Text)
			})

			c.ExpectBlock(&Block{Kind: "Heading", Text: "All databases"})
			c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘Properties marked with an \\* are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. For more information on integration capabilities, see the capabilities guide."})

			c.ExpectParameter(&Parameter{
				Property:     `object\`,
				Type:         "string",
				Description:  `Always "database".`,
				ExampleValue: `"database"`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				e.Property = strings.TrimSuffix(e.Property, `\`)
				database.AddFields(b.NewFieldStringField(e))
			})

			c.ExpectParameter(&Parameter{
				Property:     `id\`,
				Type:         "string (UUID)",
				Description:  "Unique identifier for the database.",
				ExampleValue: `"2f26ee68-df30-4251-aad4-8ddc420cba3d"`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				e.Property = strings.TrimSuffix(e.Property, `\`)
				database.AddFields(b.NewField(e, UUID))
			})

			c.ExpectParameter(&Parameter{
				Property:     "created_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this database was created. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Id("ISO8601String")))
			})

			c.ExpectParameter(&Parameter{
				Property:     "created_by",
				Type:         "Partial User",
				Description:  "User who created the database.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Id("User")))
			})

			c.ExpectParameter(&Parameter{
				Property:     "last_edited_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this database was updated. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T21:49:37.913Z"`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Id("ISO8601String")))
			})

			c.ExpectParameter(&Parameter{
				Property:     "last_edited_by",
				Type:         "Partial User",
				Description:  "User who last edited the database.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Id("User")))
			})

			c.ExpectParameter(&Parameter{
				Property:     "title",
				Type:         "array of rich text objects",
				Description:  "Name of the database as it appears in Notion.  \nSee rich text object) for a breakdown of the properties.",
				ExampleValue: "\"title\": [\n        {\n            \"type\": \"text\",\n            \"text\": {\n                \"content\": \"Can I create a URL property\",\n                \"link\": null\n            },\n            \"annotations\": {\n                \"bold\": false,\n                \"italic\": false,\n                \"strikethrough\": false,\n                \"underline\": false,\n                \"code\": false,\n                \"color\": \"default\"\n            },\n            \"plain_text\": \"Can I create a URL property\",\n            \"href\": null\n        }\n    ]",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Index().Id("RichText")))
			})

			c.ExpectParameter(&Parameter{
				Property:    "description",
				Type:        "array of rich text objects",
				Description: "Description of the database as it appears in Notion.  \nSee rich text object) for a breakdown of the properties.",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Index().Id("RichText")))
			})

			c.ExpectParameter(&Parameter{
				Property:    "icon",
				Type:        "File Object or Emoji object",
				Description: "Page icon.",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Id("FileOrEmoji")))
			})

			c.ExpectParameter(&Parameter{
				Property:    "cover",
				Type:        `File object (only type of "external" is supported currently)`,
				Description: "Page cover image.",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Op("*").Id("File")))
			})

			c.ExpectParameter(&Parameter{
				Property:    `properties\*`,
				Type:        "object",
				Description: "Schema of properties for the database as they appear in Notion.  \n  \nkey string  \nThe name of the property as it appears in Notion.  \n  \nvalue object  \nA Property object.",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				e.Property = strings.TrimSuffix(e.Property, `\*`)
				database.AddFields(b.NewField(e, jen.Map(jen.String()).Id("Property")))
			})

			c.ExpectParameter(&Parameter{
				Property:     "parent",
				Type:         "object",
				Description:  "Information about the database's parent. See Parent object.",
				ExampleValue: `{ "type": "page_id", "page_id": "af5f89b5-a8ff-4c56-a5e8-69797d11b9f8" }`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Id("Parent")))
			})

			c.ExpectParameter(&Parameter{
				Property:     "url",
				Type:         "string",
				Description:  "The URL of the Notion database.",
				ExampleValue: `"https://www.notion.so/668d797c76fa49349b05ad288df2d136"`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.String()))
			})

			c.ExpectParameter(&Parameter{
				Property:     "archived",
				Type:         "boolean",
				Description:  "The archived status of the  database.",
				ExampleValue: "false",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Bool()))
			})

			c.ExpectParameter(&Parameter{
				Property:     "in_trash",
				Type:         "boolean",
				Description:  "Whether the database has been deleted.",
				ExampleValue: "false",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Bool()))
			})

			c.ExpectParameter(&Parameter{
				Property:     "is_inline",
				Type:         "boolean",
				Description:  "Has the value true if the database appears in the page as an inline block. Otherwise has the value false if the database appears as a child page.",
				ExampleValue: "false",
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Bool()))
			})

			c.ExpectParameter(&Parameter{
				Property:     "public_url",
				Type:         "string",
				Description:  "The public page URL if the page has been published to the web. Otherwise, null.",
				ExampleValue: `"https://jm-testing.notion.site/p1-6df2c07bfc6b4c46815ad205d132e22d"1`,
			}).Output(func(e *Parameter, b *CodeBuilder) {
				database.AddFields(b.NewField(e, jen.Op("*").String()))
			})

			c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
				database.AddFields(UndocumentedRequestID(b))
			})
		},
	)
}
