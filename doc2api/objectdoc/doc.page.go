package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	var page *concreteObject

	registerTranslator(
		"https://developers.notion.com/reference/page",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The Page object contains the page property values of a single Notion page.",
			}, func(e blockElement) {
				page = b.addConcreteObject("Page", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n    \"object\": \"page\",\n    \"id\": \"be633bf1-dfa0-436d-b259-571129a590e5\",\n    \"created_time\": \"2022-10-24T22:54:00.000Z\",\n    \"last_edited_time\": \"2023-03-08T18:25:00.000Z\",\n    \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n    },\n    \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n    },\n    \"cover\": null,\n    \"icon\": {\n        \"type\": \"emoji\",\n        \"emoji\": \"🐞\"\n    },\n    \"parent\": {\n        \"type\": \"database_id\",\n        \"database_id\": \"a1d8501e-1ac1-43e9-a6bd-ea9fe6c8822b\"\n    },\n    \"archived\": true,\n    \"properties\": {\n        \"Due date\": {\n            \"id\": \"M%3BBw\",\n            \"type\": \"date\",\n            \"date\": {\n                \"start\": \"2023-02-23\",\n                \"end\": null,\n                \"time_zone\": null\n            }\n        },\n        \"Status\": {\n            \"id\": \"Z%3ClH\",\n            \"type\": \"status\",\n            \"status\": {\n                \"id\": \"86ddb6ec-0627-47f8-800d-b65afd28be13\",\n                \"name\": \"Not started\",\n                \"color\": \"default\"\n            }\n        },\n        \"Title\": {\n            \"id\": \"title\",\n            \"type\": \"title\",\n            \"title\": [\n                {\n                    \"type\": \"text\",\n                    \"text\": {\n                        \"content\": \"Bug bash\",\n                        \"link\": null\n                    },\n                    \"annotations\": {\n                        \"bold\": false,\n                        \"italic\": false,\n                        \"strikethrough\": false,\n                        \"underline\": false,\n                        \"code\": false,\n                        \"color\": \"default\"\n                    },\n                    \"plain_text\": \"Bug bash\",\n                    \"href\": null\n                }\n            ]\n        }\n    },\n    \"url\": \"https://www.notion.so/Bug-bash-be633bf1dfa0436db259571129a590e5\"\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("Page", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "All pages have a Parent. If the parent is a database, the property values conform to the schema laid out database's properties. Otherwise, the only property value is the title.",
			}, func(e blockElement) {
				page.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Page content is available as blocks. The content can be read using retrieve block children and appended using append block children.",
			}, func(e blockElement) {
				page.addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* Page object properties */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Page object properties",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Properties marked with an * are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. For more information on integration capabilities, see the capabilities guide.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "object",
				Type:         "string",
				Description:  `Always "page".`,
				ExampleValue: `"page"`,
			}, func(e parameterElement) {
				page.addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "Unique identifier of the page.",
				ExampleValue: `"45ee8d13-687b-47ce-a5ca-6e2e45548c4b"`,
			}, func(e parameterElement) {
				page.addFields(e.asField(UUID))
			})
			c.nextMustParameter(parameterElement{
				Property:     "created_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this page was created. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "created_by",
				Type:         "Partial User",
				Description:  "User who created the page.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Id("User")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "last_edited_time",
				Type:         "string (ISO 8601 date and time)",
				Description:  "Date and time when this page was updated. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "last_edited_by",
				Type:         "Partial User",
				Description:  "User who last edited the page.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Id("User")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "archived",
				Type:         "boolean",
				Description:  "The archived status of the page.",
				ExampleValue: "false",
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "icon",
				Type:         `File Object (only type of "external" is supported currently) or Emoji object`,
				Description:  "Page icon.",
				ExampleValue: "",
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Id("FileOrEmoji")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "cover",
				Type:         `File object (only type of "external" is supported currently)`,
				Description:  "Page cover image.",
				ExampleValue: "",
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Op("*").Id("File")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "properties",
				Type:         "object",
				Description:  "Property values of this page. As of version 2022-06-28, properties only contains the ID of the property; in prior versions properties contained the values as well.\n\nIf parent.type is \"page_id\" or \"workspace\", then the only valid key is title.\n\nIf parent.type is \"database_id\", then the keys and values of this field are determined by the properties  of the database this page belongs to.\n\nkey string\nName of a property as it appears in Notion.\n\nvalue object\nSee Property value object.",
				ExampleValue: `{ "id": "A%40Hk" }`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Map(jen.String()).Id("PropertyValue")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "parent",
				Type:         "object",
				Description:  "Information about the page's parent. See Parent object.",
				ExampleValue: `{ "type": "database_id", "database_id": "d9824bdc-8445-4327-be8b-5b47500af6ce" }`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.Id("Parent")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "url",
				Type:         "string",
				Description:  "The URL of the Notion page.",
				ExampleValue: `"https://www.notion.so/Avocado-d093f1d200464ce78b36e58a3f0d8043"`,
			}, func(e parameterElement) {
				page.addFields(e.asField(jen.String()))
			})
		},
	)
}
