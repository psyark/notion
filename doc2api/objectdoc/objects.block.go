package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	var blockCommon *classStruct

	registerConverter(converter{
		url:      "https://developers.notion.com/reference/block",
		fileName: "block.go",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "A block object represents a piece of content within Notion. The API translates the headings, toggles, paragraphs, lists, media, and more that you can interact with in the Notion UI as different block type objects. \n\n For example, the following block object represents a Heading 2 in the Notion UI:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					blockCommon = &classStruct{name: "blockCommon", comment: e.Text}
					b.add(blockCommon)
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n\t\"object\": \"block\",\n\t\"id\": \"c02fc1d3-db8b-45c5-a222-27595b15aea7\",\n\t\"parent\": {\n\t\t\"type\": \"page_id\",\n\t\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n\t},\n\t\"created_time\": \"2022-03-01T19:05:00.000Z\",\n\t\"last_edited_time\": \"2022-07-06T19:41:00.000Z\",\n\t\"created_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"last_edited_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"has_children\": false,\n\t\"archived\": false,\n\t\"type\": \"heading_2\",\n\t\"heading_2\": {\n\t\t\"rich_text\": [\n\t\t\t{\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"content\": \"Lacinato kale\",\n\t\t\t\t\t\"link\": null\n\t\t\t\t},\n\t\t\t\t\"annotations\": {\n\t\t\t\t\t\"bold\": false,\n\t\t\t\t\t\"italic\": false,\n\t\t\t\t\t\"strikethrough\": false,\n\t\t\t\t\t\"underline\": false,\n\t\t\t\t\t\"code\": false,\n\t\t\t\t\t\"color\": \"green\"\n\t\t\t\t},\n\t\t\t\t\"plain_text\": \"Lacinato kale\",\n\t\t\t\t\"href\": null\n\t\t\t}\n\t\t],\n\t\t\"color\": \"default\",\n    \"is_toggleable\": false\n\t}\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					blockCommon.comment += "\n\n" + e.Code
					return nil
				},
			}}},
			&objectDocParagraphElement{
				Text: "Use the Retrieve block children endpoint to list all of the blocks on a page. \n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					blockCommon.comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocHeadingElement{
				Text:   "Keys",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocCalloutElement{
				Body:   "Fields marked with an * are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. Consult the [integration capabilities reference](https://developers.notion.com/reference/capabilities) for details.",
				Title:  "",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocParametersElement{{
				Field:        "object*",
				Type:         "string",
				Description:  `Always "block".`,
				ExampleValue: `"block"`,
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&fixedStringField{
						name:    strings.TrimSuffix(e.Field, "*"),
						value:   strings.ReplaceAll(e.ExampleValue, `"`, ""),
						comment: e.Description,
					})
					return nil
				},
			}, {
				Field:        "id*",
				Type:         "string (UUIDv4)",
				Description:  "Identifier for the block.",
				ExampleValue: `"7af38973-3787-41b3-bd75-0ed3a1edfac9"`,
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     strings.TrimSuffix(e.Field, "*"),
						typeCode: jen.Qual("github.com/google/uuid", "UUID"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "parent",
				Type:         "object",
				Description:  "Information about the block's parent. See Parent object.",
				ExampleValue: "{ \"type\": \"block_id\", \"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\" }",
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     strings.TrimSuffix(e.Field, "*"),
						typeCode: jen.Id("Parent"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "type",
				Type:         "string (enum)",
				Description:  "Type of block. Possible values are:\n\n- \"bookmark\"\n- \"breadcrumb\"\n- \"bulleted_list_item\"\n- \"callout\"\n- \"child_database\"\n- \"child_page\"\n- \"column\"\n- \"column_list\"\n- \"divider\"\n- \"embed\"\n- \"equation\"\n- \"file\"\n-  \"heading_1\"\n- \"heading_2\"\n- \"heading_3\"\n- \"image\"\n- \"link_preview\"\n- \"link_to_page\"\n-  \"numbered_list_item\"\n- \"paragraph\"\n- \"pdf\"\n- \"quote\"\n- \"synced_block\"\n- \"table\"\n- \"table_of_contents\"\n- \"table_row\"\n- \"template\"\n- \"to_do\"\n- \"toggle\"\n- \"unsupported\"\n- \"video\"",
				ExampleValue: `"paragraph"`,
				output: func(e *objectDocParameter, b *builder) error {
					return nil // 各structで定義
				},
			}, {
				Field:        "created_time",
				Type:         "string (ISO 8601 date time)",
				Description:  "Date and time when this block was created. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     e.Field,
						typeCode: jen.Id("ISO8601String"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "created_by",
				Type:         "Partial User",
				Description:  "User who created the block.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     e.Field,
						typeCode: jen.Id("PartialUser"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "last_edited_time",
				Type:         "string (ISO 8601 date time)",
				Description:  "Date and time when this block was last updated. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     e.Field,
						typeCode: jen.Id("ISO8601String"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "last_edited_by",
				Type:         "Partial User",
				Description:  "User who last edited the block.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     e.Field,
						typeCode: jen.Id("PartialUser"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "archived",
				Type:         "boolean",
				Description:  "The archived status of the block.",
				ExampleValue: "false",
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     e.Field,
						typeCode: jen.Bool(),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "has_children",
				Type:         "boolean",
				Description:  "Whether or not the block has children blocks nested within it.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) error {
					blockCommon.addField(&field{
						name:     e.Field,
						typeCode: jen.Bool(),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:        "{type}",
				Type:         "block type object",
				Description:  "An object containing type-specific block information.",
				ExampleValue: "Refer to the block type object section for examples of each block type.",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // 各structで定義
				},
			}},
			&objectDocHeadingElement{
				Text: "Block types that support child blocks",
				output: func(e *objectDocHeadingElement, b *builder) error {
					blockCommon.comment += "\n" + e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nSome block types contain nested blocks. The following block types support child blocks: \n\n- Bulleted list item\n- Callout \n- Child database\n- Child page\n- Column\n- Heading 1, when the is_toggleable property is true\n- Heading 2, when the is_toggleable property is true\n- Heading 3, when the is_toggleable property is true\n- Numbered list item\n- Paragraph \n- Quote\n- Synced block\n- Table\n- Template\n- To do\n- Toggle ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					blockCommon.comment += e.Text
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:  "Only the block type objects listed in the reference below are supported. Any unsupported block types appear in the structure, but contain a `type` set to `\"unsupported\"`.",
				Title: "The API does not support all block types.",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					blockCommon.comment += "\n\n" + e.Title + "\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Block type objects",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(independentComment(e.Text))
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "Every block object has a key corresponding to the value of type. Under the key is an object with type-specific block information.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.add(independentComment(strings.TrimSpace(e.Text)))
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Bookmark",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(&classStruct{
						name: "BookmarkBlock",
						fields: []coder{
							&field{typeCode: jen.Id(blockCommon.name)},
							&fixedStringField{name: "type", value: "bookmark"},
						},
					})
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nBookmark block objects contain the following information within the bookmark property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getClassStruct("BookmarkBlock").comment = strings.TrimSpace(e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:       "caption",
				Type:        "array of rich text objects text",
				Description: "The caption for the bookmark.",
				output: func(e *objectDocParameter, b *builder) error {
					b.getClassStruct("BookmarkBlock").addField(&field{
						name:     e.Field,
						typeCode: jen.Index().Id("RichText"),
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Field:       "url",
				Type:        "string",
				Description: "The link for the bookmark.",
				output: func(e *objectDocParameter, b *builder) error {
					b.getClassStruct("BookmarkBlock").addField(&field{
						name:     e.Field,
						typeCode: jen.String(),
						comment:  e.Description,
					})
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"bookmark\",\n  //...other keys excluded\n  \"bookmark\": {\n    \"caption\": [],\n    \"url\": \"https://companywebsite.com\"\n  }\n} ",
				Language: "json",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getClassStruct("BookmarkBlock").comment += "\n" + strings.TrimSpace(e.Code)
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Breadcrumb",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nBreadcrumb block objects do not contain any information within the breadcrumb property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"breadcrumb\",\n  //...other keys excluded\n  \"breadcrumb\": {}\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Bulleted list item",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nBulleted list item block objects contain the following information within the bulleted_list_item property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text in the bulleted_list_item block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks (if any) of the bulleted_list_item block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"bulleted_list_item\",\n  //...other keys excluded\n  \"bulleted_list_item\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n      // ..other keys excluded\n    }],\n    \"color\": \"default\",\n    \"children\":[{\n      \"type\": \"paragraph\"\n      // ..other keys excluded\n    }]\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Callout",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nCallout block objects contain the following information within the callout property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text in the callout block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "An emoji or file object that represents the callout's icon. If the callout does not have an icon.",
				ExampleValue: "",
				Field:        "icon",
				Property:     "",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n\t\"type\": \"callout\",\n   // ..other keys excluded\n   \"callout\": {\n   \t\"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n      // ..other keys excluded\n    }],\n     \"icon\": {\n       \"emoji\": \"⭐\"\n     },\n     \"color\": \"default\"\n   }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Child database",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nChild database block objects contain the following information within the child_database property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The plain text title of the database.",
				ExampleValue: "",
				Field:        "title",
				Property:     "",
				Type:         "string",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"child_database\",\n  //...other keys excluded\n  \"child_database\": {\n    \"title\": \"My database\"\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "To create or update `child_database` type blocks, use the [Create a database](ref:create-a-database) and the [Update a database](ref:update-a-database) endpoints, specifying the ID of the parent page in the `parent` body param.",
				Title: "Creating and updating `child_database` blocks",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Child page",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nChild page block objects contain the following information within the child_page property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The plain text title of the page.",
				ExampleValue: "",
				Field:        "title",
				Property:     "",
				Type:         "string",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"child_page\",\n  //...other keys excluded\n  \"child_page\": {\n    \"title\": \"Lacinato kale\"\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "To create or update `child_page` type blocks, use the [Create a page](https://developers.notion.com/reference/post-page) and the [Update page](https://developers.notion.com/reference/patch-page) endpoints, specifying the ID of the parent page in the `parent` body param.",
				Title: "Creating and updating `child_page` blocks",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Code",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nCode block objects contain the following information within the code property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text in the caption of the code block.",
				ExampleValue: "",
				Field:        "caption",
				Property:     "",
				Type:         "array of Rich text object text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The rich text in the code block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of Rich text object text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The language of the code contained in the code block.",
				ExampleValue: "",
				Field:        "language",
				Property:     "",
				Type:         "- \"abap\"\n- \"arduino\"\n- \"bash\"\n- \"basic\"\n- \"c\"\n- \"clojure\"\n- \"coffeescript\"\n- \"c++\"\n- \"c#\"\n- \"css\"\n- \"dart\"\n- \"diff\"\n- \"docker\"\n- \"elixir\"\n- \"elm\"\n- \"erlang\"\n- \"flow\"\n- \"fortran\"\n- \"f#\"\n- \"gherkin\"\n- \"glsl\"\n- \"go\"\n- \"graphql\"\n- \"groovy\"\n- \"haskell\"\n- \"html\"\n- \"java\"\n- \"javascript\"\n- \"json\"\n- \"julia\"\n- \"kotlin\"\n- \"latex\"\n- \"less\"\n- \"lisp\"\n- \"livescript\"\n- \"lua\"\n- \"makefile\"\n- \"markdown\"\n- \"markup\"\n- \"matlab\"\n- \"mermaid\"\n- \"nix\"\n- \"objective-c\"\n- \"ocaml\"\n- \"pascal\"\n- \"perl\"\n- \"php\"\n- \"plain text\"\n- \"powershell\"\n- \"prolog\"\n- \"protobuf\"\n- \"python\"\n- \"r\"\n- \"reason\"\n- \"ruby\"\n- \"rust\"\n- \"sass\"\n- \"scala\"\n- \"scheme\"\n- \"scss\"\n- \"shell\"\n- \"sql\"\n- \"swift\"\n- \"typescript\"\n- \"vb.net\"\n- \"verilog\"\n- \"vhdl\"\n- \"visual basic\"\n- \"webassembly\"\n- \"xml\"\n- \"yaml\"\n- \"java/c/c++/c#\"",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"code\",\n  //...other keys excluded\n  \"code\": {\n   \t\"caption\": [],\n \t\t\"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"const a = 3\"\n      }\n    }],\n    \"language\": \"javascript\"\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Column list and column",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nColumn lists are parent blocks for columns. They do not contain any information within the column_list property. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"column_list\",\n  //...other keys excluded\n  \"column_list\": {}\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocParagraphElement{
				Text: "Columns are parent blocks for any block types listed in this reference except for other columns. They do not contain any information within the column property. They can only be appended to column_lists.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"column\",\n  //...other keys excluded\n  \"column\": {}\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocParagraphElement{
				Text: "When creating a column_list block via Append block children, the column_list must have at least two columns, and each column must have at least one child.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Retrieve the content in a column list ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nFollow these steps to fetch the content in a column_list: \n\n1. Get the column_list ID from a query to Retrieve block children for the parent page. \n\n2. Get the column children from a query to Retrieve block children for the column_list. \n\n3. Get the content in each individual column from a query to Retrieve block children for the unique column ID. \n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Divider",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nDivider block objects do not contain any information within the divider property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"divider\",\n  //...other keys excluded\n  \"divider\": {}\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Embed",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nEmbed block objects include information about another website displayed within the Notion UI. The embed property contains the following information:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The link to the website that the embed block displays.",
				ExampleValue: "",
				Field:        "url",
				Property:     "",
				Type:         "string",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"embed\",\n  //...other keys excluded\n  \"embed\": {\n    \"url\": \"https://companywebsite.com\"\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The Notion app uses a 3rd-party service, iFramely, to validate and request metadata for embeds given a URL. This works well in a web app because Notion can kick off an asynchronous request for URL information, which might take seconds or longer to complete, and then update the block with the metadata in the UI after receiving a response from iFramely.\n\nWe chose not to call iFramely when creating embed blocks in the API because the API needs to be able to return faster than the UI, and because the response from iFramely could actually cause us to change the block type. This would result in a slow and potentially confusing experience as the block in the response would not match the block sent in the request.\n\nThe result is that embed blocks created via the API may not look exactly like their counterparts created in the Notion app.",
				Title: "Differences in embed blocks between the Notion app and the API",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Equation",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nEquation block objects are represented as children of paragraph blocks. They are nested within a rich text object and contain the following information within the equation property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "A KaTeX compatible string.",
				ExampleValue: "",
				Field:        "expression",
				Property:     "",
				Type:         "string",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"equation\",\n  //...other keys excluded\n  \"equation\": {\n    \"expression\": \"e=mc^2\"\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "File",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nFile block objects contain the following information within the file property: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The caption of the file block.",
				ExampleValue: "",
				Field:        "caption",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "A constant string.",
				ExampleValue: "",
				Field:        "type",
				Property:     "",
				Type:         "\"file\"\n\n\"external\"",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "A file object that details information about the file contained in the block.",
				ExampleValue: "",
				Field:        "file",
				Property:     "",
				Type:         "file object",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"file\",\n  //...other keys excluded\n  \"file\": {\n\t\t\"caption\": [],\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://companywebsite.com/files/doc.txt\"\n    }\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Headings",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nAll heading block objects, heading_1, heading_2, and heading_3, contain the following information within their corresponding objects:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text of the heading.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "Whether or not the heading block is a toggle heading or not. If true, then the heading block toggles and can support children. If false, then the heading block is a static heading block.",
				ExampleValue: "",
				Field:        "is_toggleable",
				Property:     "",
				Type:         "boolean",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"heading_1\",\n  //...other keys excluded\n  \"heading_1\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\",\n    \"is_toggleable\": false\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"heading_2\",\n  //...other keys excluded\n  \"heading_2\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\",\n    \"is_toggleable\": false\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"heading_3\",\n  //...other keys excluded\n  \"heading_3\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\",\n    \"is_toggleable\": false\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Image",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nImage block objects contain a file object detailing information about the image. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"image\",\n  //...other keys excluded\n  \"image\": {\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://website.domain/images/image.png\"\n    }\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Supported image types ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nThe image must be directly hosted. In other words, the url cannot point to a service that retrieves the image. The following image types are supported: \n\n- .bmp\n- .gif\n- .heic\n- .jpeg\n- .jpg\n-  .png\n- .svg\n- .tif\n- .tiff\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Link Preview",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nLink Preview block objects contain the originally pasted url:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"link_preview\",\n  //...other keys excluded\n  \"link_preview\": {\n    \"url\": \"https://github.com/example/example-repo/pull/1234\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The `link_preview` block can only be returned as part of a response. The API does not support creating or appending `link_preview` blocks.",
				Title: "",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Mention",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nA mention block object is a child of a rich text object that is nested within a paragraph block object. This block type represents any @ tag in the Notion UI, for a user, date, Notion page, Notion database, or a miniaturized version of a Link Preview.  \n\nA mention block object contains the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "A constant string representing the type of the mention.",
				ExampleValue: "",
				Field:        "type",
				Property:     "",
				Type:         "\"database\"\n\n\"date\" \n\n\"link_preview\" \n\n\"page\" \n\n\"user\"",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "An object with type-specific information about the mention.",
				ExampleValue: "",
				Field:        "\"database\"\n\n\"date\" \n\n\"link_preview\" \n\n\"page\" \n\n\"user\"",
				Property:     "",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"page\",\n  \"page\": {\n    \"id\": \"3c612f56-fdd0-4a30-a4d6-bda7d7426309\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Numbered list item",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumbered list item block objects contain the following information within the numbered_list_item property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text displayed in the numbered_list_item block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks (if any) of the numbered_list_item block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"numbered_list_item\",\n  \"numbered_list_item\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Finish reading the docs\",\n          \"link\": null\n        }\n      }\n    ],\n    \"color\": \"default\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Paragraph",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nParagraph block objects contain the following information within the paragraph property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text displayed in the paragraph block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks (if any) of the paragraph block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"paragraph\",\n  //...other keys excluded\n  \"paragraph\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\"\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n//...other keys excluded\n\t\"type\": \"paragraph\",\n  \t\"paragraph\":{\n  \t\t\"rich_text\": [\n    \t\t{\n      \t\t\"type\": \"mention\",\n      \t\t\"mention\": {\n        \t\t\"type\": \"date\",\n        \t\t\"date\": {\n          \t\t\"start\": \"2023-03-01\",\n          \t\t\"end\": null,\n          \t\t\"time_zone\": null\n        \t\t}\n      \t\t},\n      \t\t\"annotations\": {\n        \t\t\"bold\": false,\n        \t\t\"italic\": false,\n        \t\t\"strikethrough\": false,\n        \t\t\"underline\": false,\n        \t\t\"code\": false,\n        \t\t\"color\": \"default\"\n      \t\t},\n      \t\t\"plain_text\": \"2023-03-01\",\n      \t\t\"href\": null\n    \t\t},\n    \t\t{\n          \"type\": \"text\",\n      \t\t\"text\": {\n        \t\t\"content\": \" \",\n        \t\t\"link\": null\n      \t\t},\n      \t\t\"annotations\": {\n        \t\t\"bold\": false,\n        \t\t\"italic\": false,\n        \t\t\"strikethrough\": false,\n        \t\t\"underline\": false,\n        \t\t\"code\": false,\n        \t\t\"color\": \"default\"\n      \t\t},\n      \t\t\"plain_text\": \" \",\n      \t\t\"href\": null\n    \t\t}\n  \t\t],\n  \t\t\"color\": \"default\"\n  \t}\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "PDF",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nA PDF block object represents a PDF that has been embedded within a Notion page. It contains the following fields:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "A caption, if provided, for the PDF block.",
				ExampleValue: "",
				Field:        "",
				Property:     "caption",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "A constant string representing the type of PDF. file indicates a Notion-hosted file, and external represents a third-party link.",
				ExampleValue: "",
				Field:        "",
				Property:     "type",
				Type:         "\"external\"\n\n\"file\"",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "An object containing type-specific information about the PDF.",
				ExampleValue: "",
				Field:        "",
				Property:     "external \n\nfile",
				Type:         "file object",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n\t\"type\": \"pdf\",\n  //...other keys excluded\n  \"pdf\": {\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://website.domain/files/doc.pdf\"\n    }\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Quote",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nQuote block objects contain the following information within the quote property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text displayed in the quote block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks, if any, of the quote block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n\t//...other keys excluded\n\t\"type\": \"quote\",\n   //...other keys excluded\n   \"quote\": {\n   \t\"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"To be or not to be...\",\n        \"link\": null\n      },\n    \t//...other keys excluded\n    }],\n    //...other keys excluded\n    \"color\": \"default\"\n   }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Synced block",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nSimilar to the Notion UI, there are two versions of a synced_block object: the original block that was created first and doesn't yet sync with anything else, and the duplicate block or blocks synced to the original.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCalloutElement{
				Body:  "An original synced block must be created before corresponding duplicate block or blocks can be made.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Original synced block",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nOriginal synced block objects contain the following information within the synced_block property: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The value is always null to signify that this is an original synced block that does not refer to another block.",
				ExampleValue: "",
				Field:        "synced_from",
				Property:     "",
				Type:         "null",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks, if any, of the synced_block block. These blocks will be mirrored in the duplicate synced_block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n    //...other keys excluded\n  \t\"type\": \"synced_block\",\n    \"synced_block\": {\n        \"synced_from\": null,\n        \"children\": [\n            {\n                \"callout\": {\n                    \"rich_text\": [\n                        {\n                            \"type\": \"text\",\n                            \"text\": {\n                                \"content\": \"Callout in synced block\"\n                            }\n                        }\n                    ]\n                }\n            }\n        ]\n    }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Duplicate synced block",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nDuplicate synced block objects contain the following information within the synced_from object: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The type of the synced from object. \n\nPossible values are: \n- \"block_id\"",
				ExampleValue: "",
				Field:        "type",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "An identifier for the original synced_block.",
				ExampleValue: "",
				Field:        "block_id",
				Property:     "",
				Type:         "string (UUIDv4)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n    //...other keys excluded\n  \t\"type\": \"synced_block\",\n    \"synced_block\": {\n        \"synced_from\": {\n            \"block_id\": \"original_synced_block_id\"\n        }\n    }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The API does not supported updating synced block content.",
				Title: "",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Table",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nTable block objects are parent blocks for table row children. Table block objects contain the following fields within the table property: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The number of columns in the table. \n\nNote that this cannot be changed via the public API once a table is created.",
				ExampleValue: "",
				Field:        "table_width",
				Property:     "",
				Type:         "integer",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "Whether the table has a column header. If true, then the first row in the table appears visually distinct from the other rows.",
				ExampleValue: "",
				Field:        "has_column_header",
				Property:     "",
				Type:         "boolean",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "Whether the table has a header row. If true, then the first column in the table appears visually distinct from the other columns.",
				ExampleValue: "",
				Field:        "has_row_header",
				Property:     "",
				Type:         "boolean",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"table\",\n  \"table\": {\n    \"table_width\": 2,\n    \"has_column_header\": false,\n    \"has_row_header\": false\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "Note that the number of columns in a table can only be set when the table is first created. Calls to the Update block endpoint to update `table_width` fail.",
				Title: "`table_width` can only be set when the table is first created.",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Table rows",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nFollow these steps to fetch the table_rows of a table: \n\n1. Get the table ID from a query to Retrieve block children for the parent page. \n\n2. Get the table_rows from a query to Retrieve block children for the table. \n\nA table_row block object contains the following fields within the table_row property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "An array of cell contents in horizontal display order. Each cell is an array of rich text objects.",
				ExampleValue: "",
				Field:        "",
				Property:     "cells",
				Type:         "array of array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"table_row\",\n  \"table_row\": {\n    \"cells\": [\n      [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"column 1 content\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"column 1 content\",\n          \"href\": null\n        }\n      ],\n      [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"column 2 content\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"column 2 content\",\n          \"href\": null\n        }\n      ],\n      [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"column 3 content\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"column 3 content\",\n          \"href\": null\n        }\n      ]\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocCalloutElement{
				Body:  "When creating a table block via the [Append block children](ref:patch-block-children) endpoint, the `table` must have at least one `table_row` whose `cells` array has the same length as the `table_width`.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Table of contents",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nTable of contents block objects contain the following information within the table_of_contents property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "",
				Property:     "color",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n\t\"type\": \"table_of_contents\",\n  \"table_of_contents\": {\n  \t\"color\": \"default\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Template",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCalloutElement{
				Body:  "As of March 27, 2023 creation of template blocks will no longer be supported.",
				Title: "Deprecation Notice",
				Type:  "danger",
				output: func(e *objectDocCalloutElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "Template blocks represent template buttons in the Notion UI.\n\nTemplate block objects contain the following information within the template property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text displayed in the title of the template.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks, if any, of the template block. These blocks are duplicated when the template block is used in the UI.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"template\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Add a new to-do\",\n          \"link\": null\n        },\n        \"annotations\": {\n          //...other keys excluded\n        },\n        \"plain_text\": \"Add a new to-do\",\n        \"href\": null\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "To do",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nTo do block objects contain the following information within the to_do property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text displayed in the To do block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "Whether the To do is checked.",
				ExampleValue: "",
				Field:        "checked",
				Property:     "",
				Type:         "boolean (optional)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks, if any, of the To do block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"to_do\",\n  \"to_do\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Finish Q3 goals\",\n        \"link\": null\n      }\n    }],\n    \"checked\": false,\n    \"color\": \"default\",\n    \"children\":[{\n      \"type\": \"paragraph\"\n      // ..other keys excluded\n    }]\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Toggle blocks",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nToggle block objects contain the following information within the toggle property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "The rich text displayed in the Toggle block.",
				ExampleValue: "",
				Field:        "rich_text",
				Property:     "",
				Type:         "array of rich text objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The color of the block. Possible values are: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n-  \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n- \"orange_background\"\n- \"yellow\"\n- \"green\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background\"\n- \"yellow_background\"",
				ExampleValue: "",
				Field:        "color",
				Property:     "",
				Type:         "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}, {
				Description:  "The nested child blocks, if any, of the Toggle block.",
				ExampleValue: "",
				Field:        "children",
				Property:     "",
				Type:         "array of block objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  //...other keys excluded\n  \"type\": \"toggle\",\n  \"toggle\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Additional project details\",\n        \"link\": null\n      }\n      //...other keys excluded\n    }],\n    \"color\": \"default\",\n    \"children\":[{\n      \"type\": \"paragraph\"\n      // ..other keys excluded\n    }]\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Video",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nVideo block objects contain a file object detailing information about the image. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"video\",\n  //...other keys excluded\n  \"video\": {\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://companywebsite.com/files/video.mp4\"\n    }\n  }\n} ",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Supported video types",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\n- .amv\n- .asf\n- .avi\n- .f4v\n- .flv\n- .gifv\n- .mkv\n- .mov\n- .mpg\n- .mpeg\n- .mpv\n- .mp4\n- .m4v\n- .qt\n- .wmv",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
		},
	})
}
