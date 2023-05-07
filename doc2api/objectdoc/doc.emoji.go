package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	registerTranslator(
		"https://developers.notion.com/reference/emoji-object",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "An emoji object contains information about an emoji character. It is most often used to represent an emoji that is rendered as a page icon in the Notion UI.",
			}, func(e blockElement) {
				union := b.addUnionToGlobalIfNotExists("FileOrEmoji", "type")
				b.addConcreteObject("Emoji", e.Text).addToUnion(union)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"emoji\",\n  \"emoji\": \"😻\"\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("Emoji", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The object contains the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         `"emoji"`,
				Description:  `The constant string "emoji" that represents the object type.`,
				ExampleValue: `"emoji"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Emoji").addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Property:     "emoji",
				Type:         "string",
				Description:  "The emoji character.",
				ExampleValue: `"😻"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Emoji").addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "To use the Notion API to render an emoji object as a page icon, set a page’s icon property field to an emoji object.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Example: set a page icon via the Create a page endpoint */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example: set a page icon via the Create a page endpoint",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "curl 'https://api.notion.com/v1/pages' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  --data '{\n  \"parent\": {\n    \"page_id\": \"13d6da822f9343fa8ec14c89b8184d5a\"\n  },\n  \"properties\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"A page with an avocado icon\",\n          \"link\": null\n        }\n      }\n    ]\n  },\n  \"icon\": {\n    \"type\": \"emoji\",\n    \"emoji\": \"🥑\"\n  }\n}'",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Example: set a page icon via the Update page endpoint */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example: set a page icon via the Update page endpoint",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "curl https://api.notion.com/v1/pages/60bdc8bd-3880-44b8-a9cd-8a145b3ffbd7 \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  -X PATCH \\\n\t--data '{\n  \"icon\": {\n\t  \"type\": \"emoji\", \n\t  \"emoji\": \"🥨\"\n    }\n}'",
			}, func(e blockElement) {})
		},
	)
}
