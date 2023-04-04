package doc2api

func init() {
	registerConverter(converter{
		url:      "https://developers.notion.com/reference/block",
		fileName: "block.go",
		matchers: []elementMatcher{
			paragraphElementMatcher{
				local:  &objectDocParagraphElement{"A block object represents a piece of content within Notion. The API translates the headings, toggles, paragraphs, lists, media, and more that you can interact with in the Notion UI as different block type objects. \n\n For example, the following block object represents a Heading 2 in the Notion UI:"},
				output: func(e *objectDocParagraphElement) error { return nil },
			},
			codeElementMatcher{
				local: &objectDocCodeElement{Codes: []objectDocCodeElementCode{
					{
						Name:     "",
						Language: "json",
						Code:     "{\n\t\"object\": \"block\",\n\t\"id\": \"c02fc1d3-db8b-45c5-a222-27595b15aea7\",\n\t\"parent\": {\n\t\t\"type\": \"page_id\",\n\t\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n\t},\n\t\"created_time\": \"2022-03-01T19:05:00.000Z\",\n\t\"last_edited_time\": \"2022-07-06T19:41:00.000Z\",\n\t\"created_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"last_edited_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"has_children\": false,\n\t\"archived\": false,\n\t\"type\": \"heading_2\",\n\t\"heading_2\": {\n\t\t\"rich_text\": [\n\t\t\t{\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"content\": \"Lacinato kale\",\n\t\t\t\t\t\"link\": null\n\t\t\t\t},\n\t\t\t\t\"annotations\": {\n\t\t\t\t\t\"bold\": false,\n\t\t\t\t\t\"italic\": false,\n\t\t\t\t\t\"strikethrough\": false,\n\t\t\t\t\t\"underline\": false,\n\t\t\t\t\t\"code\": false,\n\t\t\t\t\t\"color\": \"green\"\n\t\t\t\t},\n\t\t\t\t\"plain_text\": \"Lacinato kale\",\n\t\t\t\t\"href\": null\n\t\t\t}\n\t\t],\n\t\t\"color\": \"default\",\n    \"is_toggleable\": false\n\t}\n}",
					},
				}},
				output: func(e *objectDocCodeElement) error { return nil },
			},
			paragraphElementMatcher{
				local:  &objectDocParagraphElement{Text: "Use the Retrieve block children endpoint to list all of the blocks on a page. \n"},
				output: func(odpe *objectDocParagraphElement) error { return nil },
			},
			headingElementMatcher{
				local:  &objectDocHeadingElement{Text: "Keys"},
				output: func(odhe *objectDocHeadingElement) error { return nil },
			},
			calloutElementMatcher{
				local:  &objectDocCalloutElement{Type: "info", Title: "", Body: "Fields marked with an * are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. Consult the [integration capabilities reference](https://developers.notion.com/reference/capabilities) for details."},
				output: func(odce *objectDocCalloutElement) error { return nil },
			},
			paragraphElementMatcher{
				local:  &objectDocParagraphElement{Text: ""},
				output: func(odpe *objectDocParagraphElement) error { return nil },
			},
			parametersElementMatcher{
				local: &objectDocParametersElement{
					{Name: "object*", Type: "string", Description: "Always \"block\".", ExampleValue: "\"block\""},
					{Name: "id*", Type: "string (UUIDv4)", Description: "Identifier for the block.", ExampleValue: "\"7af38973-3787-41b3-bd75-0ed3a1edfac9\""},
					{Name: "parent", Type: "object", Description: "Information about the block's parent. See Parent object.", ExampleValue: "{ \"type\": \"block_id\", \"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\" }"},
					{Name: "type", Type: "string (enum)", Description: "Type of block. Possible values are:\n\n- \"bookmark\"\n- \"breadcrumb\"\n- \"bulleted_list_item\"\n- \"callout\"\n- \"child_database\"\n- \"child_page\"\n- \"column\"\n- \"column_list\"\n- \"divider\"\n- \"embed\"\n- \"equation\"\n- \"file\"\n-  \"heading_1\"\n- \"heading_2\"\n- \"heading_3\"\n- \"image\"\n- \"link_preview\"\n- \"link_to_page\"\n-  \"numbered_list_item\"\n- \"paragraph\"\n- \"pdf\"\n- \"quote\"\n- \"synced_block\"\n- \"table\"\n- \"table_of_contents\"\n- \"table_row\"\n- \"template\"\n- \"to_do\"\n- \"toggle\"\n- \"unsupported\"\n- \"video\"", ExampleValue: "\"paragraph\""},
					{Name: "created_time", Type: "string (ISO 8601 date time)", Description: "Date and time when this block was created. Formatted as an ISO 8601 date time string.", ExampleValue: "\"2020-03-17T19:10:04.968Z\""},
					{Name: "created_by", Type: "Partial User", Description: "User who created the block.", ExampleValue: "{\"object\": \"user\",\"id\": \"45ee8d13-687b-47ce-a5ca-6e2e45548c4b\"}"},
					{Name: "last_edited_time", Type: "string (ISO 8601 date time)", Description: "Date and time when this block was last updated. Formatted as an ISO 8601 date time string.", ExampleValue: "\"2020-03-17T19:10:04.968Z\""},
					{Name: "last_edited_by", Type: "Partial User", Description: "User who last edited the block.", ExampleValue: "{\"object\": \"user\",\"id\": \"45ee8d13-687b-47ce-a5ca-6e2e45548c4b\"}"},
					{Name: "archived", Type: "boolean", Description: "The archived status of the block.", ExampleValue: "false"},
					{Name: "has_children", Type: "boolean", Description: "Whether or not the block has children blocks nested within it.", ExampleValue: "true"},
					{Name: "{type}", Type: "block type object", Description: "An object containing type-specific block information.", ExampleValue: "Refer to the block type object section for examples of each block type."},
				},
				output: func(odpe *objectDocParametersElement) error { return nil },
			},
			headingElementMatcher{
				local:  &objectDocHeadingElement{Text: "Block types that support child blocks"},
				output: func(odhe *objectDocHeadingElement) error { return nil },
			},
			paragraphElementMatcher{
				local:  &objectDocParagraphElement{Text: "\nSome block types contain nested blocks. The following block types support child blocks: \n\n- Bulleted list item\n- Callout \n- Child database\n- Child page\n- Column\n- Heading 1, when the is_toggleable property is true\n- Heading 2, when the is_toggleable property is true\n- Heading 3, when the is_toggleable property is true\n- Numbered list item\n- Paragraph \n- Quote\n- Synced block\n- Table\n- Template\n- To do\n- Toggle "},
				output: func(odpe *objectDocParagraphElement) error { return nil },
			},
			calloutElementMatcher{
				local:  &objectDocCalloutElement{Type: "info", Title: "The API does not support all block types.", Body: "Only the block type objects listed in the reference below are supported. Any unsupported block types appear in the structure, but contain a `type` set to `\"unsupported\"`."},
				output: func(odce *objectDocCalloutElement) error { return nil },
			},
		},
	})
}
