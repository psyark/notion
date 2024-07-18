package objects_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestEmoji(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/emoji-object")

	var emoji *ConcreteObject

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "An emoji object contains information about an emoji character. It is most often used to represent an emoji that is rendered as a page icon in the Notion UI.",
	}).Output(func(e *Block, b *CodeBuilder) {
		union := b.AddUnionToGlobalIfNotExists("FileOrEmoji", "type")
		emoji = b.AddConcreteObject("Emoji", e.Text)
		emoji.AddToUnion(union)
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"emoji\",\n  \"emoji\": \"😻\"\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		b.AddUnmarshalTest("Emoji", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The object contains the following fields:"})

	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         `"emoji"`,
		Description:  `The constant string "emoji" that represents the object type.`,
		ExampleValue: `"emoji"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		emoji.AddFields(b.NewFixedStringField(e))
	})

	c.ExpectParameter(&Parameter{
		Property:     "emoji",
		Type:         "string",
		Description:  "The emoji character.",
		ExampleValue: `"😻"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		emoji.AddFields(b.NewField(e, jen.String()))
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "To use the Notion API to render an emoji object as a page icon, set a page’s icon property field to an emoji object."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: set a page icon via the Create a page endpoint"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl 'https://api.notion.com/v1/pages' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  --data '{\n  \"parent\": {\n    \"page_id\": \"13d6da822f9343fa8ec14c89b8184d5a\"\n  },\n  \"properties\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"A page with an avocado icon\",\n          \"link\": null\n        }\n      }\n    ]\n  },\n  \"icon\": {\n    \"type\": \"emoji\",\n    \"emoji\": \"🥑\"\n  }\n}'"})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: set a page icon via the Update page endpoint"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl https://api.notion.com/v1/pages/60bdc8bd-3880-44b8-a9cd-8a145b3ffbd7 \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  -X PATCH \\\n\t--data '{\n  \"icon\": {\n\t  \"type\": \"emoji\", \n\t  \"emoji\": \"🥨\"\n    }\n}'"})
}
