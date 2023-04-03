package doc2api

func init() {
	registerConverter(converter{
		url:      "https://developers.notion.com/reference/block",
		fileName: "block.go",
		matchers: []elementMatcher{
			paragraphElementMatcher{
				local:  &objectDocParagraphElement{"A block object represents a piece of content within Notion. The API translates the headings, toggles, paragraphs, lists, media, and more that you can interact with in the Notion UI as different block type objects. \n\n For example, the following block object represents a Heading 2 in the Notion UI:"},
				output: func(e objectDocParagraphElement) error { return nil },
			},
			codeElementMatcher{
				local: objectDocCodeElement{Codes: []objectDocCodeElementCode{
					{
						Name:     "",
						Language: "json",
						Code:     "{\n\t\"object\": \"block\",\n\t\"id\": \"c02fc1d3-db8b-45c5-a222-27595b15aea7\",\n\t\"parent\": {\n\t\t\"type\": \"page_id\",\n\t\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n\t},\n\t\"created_time\": \"2022-03-01T19:05:00.000Z\",\n\t\"last_edited_time\": \"2022-07-06T19:41:00.000Z\",\n\t\"created_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"last_edited_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"has_children\": false,\n\t\"archived\": false,\n\t\"type\": \"heading_2\",\n\t\"heading_2\": {\n\t\t\"rich_text\": [\n\t\t\t{\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"content\": \"Lacinato kale\",\n\t\t\t\t\t\"link\": null\n\t\t\t\t},\n\t\t\t\t\"annotations\": {\n\t\t\t\t\t\"bold\": false,\n\t\t\t\t\t\"italic\": false,\n\t\t\t\t\t\"strikethrough\": false,\n\t\t\t\t\t\"underline\": false,\n\t\t\t\t\t\"code\": false,\n\t\t\t\t\t\"color\": \"green\"\n\t\t\t\t},\n\t\t\t\t\"plain_text\": \"Lacinato kale\",\n\t\t\t\t\"href\": null\n\t\t\t}\n\t\t],\n\t\t\"color\": \"default\",\n    \"is_toggleable\": false\n\t}\n}",
					},
				}},
				output: func(e objectDocCodeElement) error { return nil },
			},
		},
	})
}
