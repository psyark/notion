package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/emoji-object",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "An emoji object contains information about an emoji character. It is most often used to represent an emoji that is rendered as a page icon in the Notion UI. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.add(&classStruct{name: "Emoji", comment: e.Text})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"emoji\",\n  \"emoji\": \"😻\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil
				},
			}}},
			&objectDocParagraphElement{
				Text: "The object contains the following fields:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "type",
				Type:         `"emoji"`,
				Description:  `The constant string "emoji" that represents the object type.`,
				ExampleValue: `"emoji"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getClassStruct("Emoji").addField(&fixedStringField{
						name:    e.Property,
						value:   strings.ReplaceAll(e.Type, `"`, ""),
						comment: e.Description,
					})
					return nil
				},
			}, {
				Property:     "emoji",
				Type:         "string",
				Description:  "The emoji character.",
				ExampleValue: `"😻"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getClassStruct("Emoji").addField(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocParagraphElement{
				Text:   "To use the Notion API to render an emoji object as a page icon, set a page’s icon property field to an emoji object. \n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Example: set a page icon via the Create a page endpoint",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "curl 'https://api.notion.com/v1/pages' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  --data '{\n  \"parent\": {\n    \"page_id\": \"13d6da822f9343fa8ec14c89b8184d5a\"\n  },\n  \"properties\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"A page with an avocado icon\",\n          \"link\": null\n        }\n      }\n    ]\n  },\n  \"icon\": {\n    \"type\": \"emoji\",\n    \"emoji\": \"🥑\"\n  }\n}'",
				Language: "curl",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text:   "Example: set a page icon via the Update page endpoint",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "curl https://api.notion.com/v1/pages/60bdc8bd-3880-44b8-a9cd-8a145b3ffbd7 \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  -X PATCH \\\n\t--data '{\n  \"icon\": {\n\t  \"type\": \"emoji\", \n\t  \"emoji\": \"🥨\"\n    }\n}'",
				Language: "curl",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
		},
	})
}
