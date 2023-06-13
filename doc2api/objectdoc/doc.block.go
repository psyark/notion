package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	var block *adaptiveObject
	registerTranslator(
		"https://developers.notion.com/reference/block",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A block object represents a piece of content within Notion. The API translates the headings, toggles, paragraphs, lists, media, and more that you can interact with in the Notion UI as different block type objects.",
			}, func(e blockElement) {
				block = b.addAdaptiveObject("Block", "type", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "For example, the following block object represents a Heading 2 in the Notion UI:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n\t\"object\": \"block\",\n\t\"id\": \"c02fc1d3-db8b-45c5-a222-27595b15aea7\",\n\t\"parent\": {\n\t\t\"type\": \"page_id\",\n\t\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n\t},\n\t\"created_time\": \"2022-03-01T19:05:00.000Z\",\n\t\"last_edited_time\": \"2022-07-06T19:41:00.000Z\",\n\t\"created_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"last_edited_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"has_children\": false,\n\t\"archived\": false,\n\t\"type\": \"heading_2\",\n\t\"heading_2\": {\n\t\t\"rich_text\": [\n\t\t\t{\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"content\": \"Lacinato kale\",\n\t\t\t\t\t\"link\": null\n\t\t\t\t},\n\t\t\t\t\"annotations\": {\n\t\t\t\t\t\"bold\": false,\n\t\t\t\t\t\"italic\": false,\n\t\t\t\t\t\"strikethrough\": false,\n\t\t\t\t\t\"underline\": false,\n\t\t\t\t\t\"code\": false,\n\t\t\t\t\t\"color\": \"green\"\n\t\t\t\t},\n\t\t\t\t\"plain_text\": \"Lacinato kale\",\n\t\t\t\t\"href\": null\n\t\t\t}\n\t\t],\n\t\t\"color\": \"default\",\n    \"is_toggleable\": false\n\t}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Block", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Use the Retrieve block children endpoint to list all of the blocks on a page.",
			}, func(e blockElement) {
				block.addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* Keys */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Keys",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘Fields marked with an \\* are available to integrations with any capabilities. Other properties require read content capabilities in order to be returned from the Notion API. Consult the integration capabilities reference for details.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "object\\",
				Type:         "string",
				Description:  `Always "block".`,
				ExampleValue: `"block"`,
			}, func(e parameterElement) {
				e.Property = strings.TrimSuffix(e.Property, "\\")
				block.addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Property:     "id\\",
				Type:         "string (UUIDv4)",
				Description:  "Identifier for the block.",
				ExampleValue: `"7af38973-3787-41b3-bd75-0ed3a1edfac9"`,
			}, func(e parameterElement) {
				e.Property = strings.TrimSuffix(e.Property, "\\")
				block.addFields(e.asField(UUID))
			})
			c.nextMustParameter(parameterElement{
				Property:     "parent",
				Type:         "object",
				Description:  "Information about the block's parent. See Parent object.",
				ExampleValue: `{ "type": "block_id", "block_id": "7d50a184-5bbe-4d90-8f29-6bec57ed817b" }`,
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Op("*").Id("Parent"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         "string (enum)",
				Description:  "Type of block. Possible values are:  \n  \n- \"bookmark\"  \n- \"breadcrumb\"  \n- \"bulleted_list_item\"  \n- \"callout\"  \n- \"child_database\"  \n- \"child_page\"  \n- \"column\"  \n- \"column_list\"  \n- \"divider\"  \n- \"embed\"  \n- \"equation\"  \n- \"file\"  \n-  \"heading_1\"  \n- \"heading_2\"  \n- \"heading_3\"  \n- \"image\"  \n- \"link_preview\"  \n- \"link_to_page\"  \n-  \"numbered_list_item\"  \n- \"paragraph\"  \n- \"pdf\"  \n- \"quote\"  \n- \"synced_block\"  \n- \"table\"  \n- \"table_of_contents\"  \n- \"table_row\"  \n- \"template\"  \n- \"to_do\"  \n- \"toggle\"  \n- \"unsupported\"  \n- \"video\"",
				ExampleValue: `"paragraph"`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "created_time",
				Type:         "string (ISO 8601 date time)",
				Description:  "Date and time when this block was created. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "created_by",
				Type:         "Partial User",
				Description:  "User who created the block.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Id("User")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "last_edited_time",
				Type:         "string (ISO 8601 date time)",
				Description:  "Date and time when this block was last updated. Formatted as an ISO 8601 date time string.",
				ExampleValue: `"2020-03-17T19:10:04.968Z"`,
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "last_edited_by",
				Type:         "Partial User",
				Description:  "User who last edited the block.",
				ExampleValue: `{"object": "user","id": "45ee8d13-687b-47ce-a5ca-6e2e45548c4b"}`,
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Id("User")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "archived",
				Type:         "boolean",
				Description:  "The archived status of the block.",
				ExampleValue: "false",
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "has_children",
				Type:         "boolean",
				Description:  "Whether or not the block has children blocks nested within it.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				block.addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "An object containing type-specific block information.",
				ExampleValue: "Refer to the block type object section for examples of each block type.",
				Property:     "{type}",
				Type:         "block type object",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* Block types that support child blocks */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Block types that support child blocks",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Some block types contain nested blocks. The following block types support child blocks:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Bulleted list itemCalloutChild databaseChild pageColumnHeading 1, when the is_toggleable property is trueHeading 2, when the is_toggleable property is trueHeading 3, when the is_toggleable property is trueNumbered list itemParagraphQuoteSynced blockTableTemplateTo doToggle",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘 The API does not support all block types.Only the block type objects listed in the reference below are supported. Any unsupported block types appear in the structure, but contain a type set to \"unsupported\".",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Block type objects */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Block type objects",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Every block object has a key corresponding to the value of type. Under the key is an object with type-specific block information.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Bookmark */ {
			var bookmark *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘Many block types support rich text. In cases where it is supported, a rich_text object will be included in the block type object. All rich_text objects will include a plain_text property, which provides a convenient way for developers to access unformatted text from the Notion block.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Bookmark",
			}, func(e blockElement) {
				bookmark = block.addAdaptiveFieldWithSpecificObject("bookmark", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Bookmark block objects contain the following information within the bookmark property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "caption",
				Type:        "array of rich text objects text",
				Description: "The caption for the bookmark.",
			}, func(e parameterElement) {
				bookmark.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "url",
				Type:        "string",
				Description: "The link for the bookmark.",
			}, func(e parameterElement) {
				bookmark.addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"bookmark\",\n  //...other keys excluded\n  \"bookmark\": {\n    \"caption\": [],\n    \"url\": \"https://companywebsite.com\"\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Breadcrumb */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Breadcrumb",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Breadcrumb block objects do not contain any information within the breadcrumb property.",
			}, func(e blockElement) {
				block.addAdaptiveFieldWithEmptyStruct("breadcrumb", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"breadcrumb\",\n  //...other keys excluded\n  \"breadcrumb\": {}\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Bulleted list item */ {
			var bulletedListItem *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Bulleted list item",
			}, func(e blockElement) {
				bulletedListItem = block.addAdaptiveFieldWithSpecificObject("bulleted_list_item", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Bulleted list item block objects contain the following information within the bulleted_list_item property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The rich text in the bulleted_list_item block.",
				ExampleValue: "",
				Property:     "rich_text",
				Type:         "array of rich text objects",
			}, func(e parameterElement) {
				bulletedListItem.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
				ExampleValue: "",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				bulletedListItem.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The nested child blocks (if any) of the bulleted_list_item block.",
				ExampleValue: "",
				Property:     "children",
				Type:         "array of block objects",
			}, func(e parameterElement) {
				bulletedListItem.addFields(e.asField(jen.Index().Id("Block")))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"bulleted_list_item\",\n  //...other keys excluded\n  \"bulleted_list_item\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n      // ..other keys excluded\n    }],\n    \"color\": \"default\",\n    \"children\":[{\n      \"type\": \"paragraph\"\n      // ..other keys excluded\n    }]\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Callout */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Callout",
			}, func(e blockElement) {
				specificObject = block.addAdaptiveFieldWithSpecificObject("callout", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Callout block objects contain the following information within the callout property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "rich_text",
				Type:        "array of rich text objects",
				Description: "The rich text in the callout block.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "icon",
				Type:        "object",
				Description: "An emoji or file object that represents the callout's icon. If the callout does not have an icon.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Id("FileOrEmoji"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:    "color",
				Type:        "string (enum)",
				Description: "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n\t\"type\": \"callout\",\n   // ..other keys excluded\n   \"callout\": {\n   \t\"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n      // ..other keys excluded\n    }],\n     \"icon\": {\n       \"emoji\": \"⭐\"\n     },\n     \"color\": \"default\"\n   }\n}\n",
			}, func(e blockElement) {})

			specificObject.addFields(&field{name: "children", comment: "undocumented", typeCode: jen.Index().Id("Block"), omitEmpty: true})
		},
		func(c *comparator, b *builder) /* Child database */ {
			var childDatabase *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Child database",
			}, func(e blockElement) {
				childDatabase = block.addAdaptiveFieldWithSpecificObject("child_database", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Child database block objects contain the following information within the child_database property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "title",
				Type:        "string",
				Description: "The plain text title of the database.",
			}, func(e parameterElement) {
				childDatabase.addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"child_database\",\n  //...other keys excluded\n  \"child_database\": {\n    \"title\": \"My database\"\n  }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘 Creating and updating child_database blocksTo create or update child_database type blocks, use the Create a database and the Update a database endpoints, specifying the ID of the parent page in the parent body param.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Child page */ {
			var childPage *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Child page",
			}, func(e blockElement) {
				childPage = block.addAdaptiveFieldWithSpecificObject("child_page", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Child page block objects contain the following information within the child_page property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "title",
				Type:        "string",
				Description: "The plain text title of the page.",
			}, func(e parameterElement) {
				childPage.addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"child_page\",\n  //...other keys excluded\n  \"child_page\": {\n    \"title\": \"Lacinato kale\"\n  }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘 Creating and updating child_page blocksTo create or update child_page type blocks, use the Create a page and the Update page endpoints, specifying the ID of the parent page in the parent body param.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Code */ {
			var code *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Code",
			}, func(e blockElement) {
				code = block.addAdaptiveFieldWithSpecificObject("code", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Code block objects contain the following information within the code property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "caption",
				Type:        "array of Rich text object text objects",
				Description: "The rich text in the caption of the code block.",
			}, func(e parameterElement) {
				code.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "rich_text",
				Type:        "array of Rich text object text objects",
				Description: "The rich text in the code block.",
			}, func(e parameterElement) {
				code.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "language",
				Type:        "- \"abap\"  \n- \"arduino\"  \n- \"bash\"  \n- \"basic\"  \n- \"c\"  \n- \"clojure\"  \n- \"coffeescript\"  \n- \"c++\"  \n- \"c#\"  \n- \"css\"  \n- \"dart\"  \n- \"diff\"  \n- \"docker\"  \n- \"elixir\"  \n- \"elm\"  \n- \"erlang\"  \n- \"flow\"  \n- \"fortran\"  \n- \"f#\"  \n- \"gherkin\"  \n- \"glsl\"  \n- \"go\"  \n- \"graphql\"  \n- \"groovy\"  \n- \"haskell\"  \n- \"html\"  \n- \"java\"  \n- \"javascript\"  \n- \"json\"  \n- \"julia\"  \n- \"kotlin\"  \n- \"latex\"  \n- \"less\"  \n- \"lisp\"  \n- \"livescript\"  \n- \"lua\"  \n- \"makefile\"  \n- \"markdown\"  \n- \"markup\"  \n- \"matlab\"  \n- \"mermaid\"  \n- \"nix\"  \n- \"objective-c\"  \n- \"ocaml\"  \n- \"pascal\"  \n- \"perl\"  \n- \"php\"  \n- \"plain text\"  \n- \"powershell\"  \n- \"prolog\"  \n- \"protobuf\"  \n- \"python\"  \n- \"r\"  \n- \"reason\"  \n- \"ruby\"  \n- \"rust\"  \n- \"sass\"  \n- \"scala\"  \n- \"scheme\"  \n- \"scss\"  \n- \"shell\"  \n- \"sql\"  \n- \"swift\"  \n- \"typescript\"  \n- \"vb.net\"  \n- \"verilog\"  \n- \"vhdl\"  \n- \"visual basic\"  \n- \"webassembly\"  \n- \"xml\"  \n- \"yaml\"  \n- \"java/c/c++/c#\"",
				Description: "The language of the code contained in the code block.",
			}, func(e parameterElement) {
				code.addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"code\",\n  //...other keys excluded\n  \"code\": {\n   \t\"caption\": [],\n \t\t\"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"const a = 3\"\n      }\n    }],\n    \"language\": \"javascript\"\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Column list and column */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Column list and column",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Column lists are parent blocks for columns. They do not contain any information within the column_list property.",
			}, func(e blockElement) {
				block.addAdaptiveFieldWithEmptyStruct("column_list", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"column_list\",\n  //...other keys excluded\n  \"column_list\": {}\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Columns are parent blocks for any block types listed in this reference except for other columns. They do not contain any information within the column property. They can only be appended to column_lists.",
			}, func(e blockElement) {
				block.addAdaptiveFieldWithEmptyStruct("column", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"column\",\n  //...other keys excluded\n  \"column\": {}\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "When creating a column_list block via Append block children, the column_list must have at least two columns, and each column must have at least one child.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Retrieve the content in a column list */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Retrieve the content in a column list",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Follow these steps to fetch the content in a column_list:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Get the column_list ID from a query to Retrieve block children for the parent page.Get the column children from a query to Retrieve block children for the column_list.Get the content in each individual column from a query to Retrieve block children for the unique column ID.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Divider */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Divider",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Divider block objects do not contain any information within the divider property.",
			}, func(e blockElement) {
				block.addAdaptiveFieldWithEmptyStruct("divider", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"divider\",\n  //...other keys excluded\n  \"divider\": {}\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Embed */ {
			var embed *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Embed",
			}, func(e blockElement) {
				embed = block.addAdaptiveFieldWithSpecificObject("embed", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Embed block objects include information about another website displayed within the Notion UI. The embed property contains the following information:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "url",
				Type:        "string",
				Description: "The link to the website that the embed block displays.",
			}, func(e parameterElement) {
				embed.addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"embed\",\n  //...other keys excluded\n  \"embed\": {\n    \"url\": \"https://companywebsite.com\"\n  }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧 Differences in embed blocks between the Notion app and the APIThe Notion app uses a 3rd-party service, iFramely, to validate and request metadata for embeds given a URL. This works well in a web app because Notion can kick off an asynchronous request for URL information, which might take seconds or longer to complete, and then update the block with the metadata in the UI after receiving a response from iFramely.We chose not to call iFramely when creating embed blocks in the API because the API needs to be able to return faster than the UI, and because the response from iFramely could actually cause us to change the block type. This would result in a slow and potentially confusing experience as the block in the response would not match the block sent in the request.The result is that embed blocks created via the API may not look exactly like their counterparts created in the Notion app.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Equation */ {
			var equation *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Equation",
			}, func(e blockElement) {
				equation = block.addAdaptiveFieldWithSpecificObject("equation", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Equation block objects are represented as children of paragraph blocks. They are nested within a rich text object and contain the following information within the equation property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "expression",
				Type:        "string",
				Description: "A KaTeX compatible string.",
			}, func(e parameterElement) {
				equation.addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"equation\",\n  //...other keys excluded\n  \"equation\": {\n    \"expression\": \"e=mc^2\"\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* File */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "File",
			}, func(e blockElement) {
				block.addAdaptiveFieldWithType("file", e.Text, jen.Op("*").Id("File")) // TODO FileWithCaption
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "File block objects contain the following information within the file property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The caption of the file block.",
				ExampleValue: "",
				Property:     "caption",
				Type:         "array of rich text objects",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "A constant string.",
				ExampleValue: "",
				Property:     "type",
				Type:         "\"file\"  \n  \n\"external\"",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "A file object that details information about the file contained in the block.",
				ExampleValue: "",
				Property:     "file",
				Type:         "file object",
			}, func(e parameterElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"file\",\n  //...other keys excluded\n  \"file\": {\n\t\t\"caption\": [],\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://companywebsite.com/files/doc.txt\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Headings */ {
			var heading *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Headings",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "All heading block objects, heading_1, heading_2, and heading_3, contain the following information within their corresponding objects:",
			}, func(e blockElement) {
				heading = b.addConcreteObject("BlockHeading", e.Text)
				block.addAdaptiveFieldWithType("heading_1", "", jen.Op("*").Id("BlockHeading"))
				block.addAdaptiveFieldWithType("heading_2", "", jen.Op("*").Id("BlockHeading"))
				block.addAdaptiveFieldWithType("heading_3", "", jen.Op("*").Id("BlockHeading"))
			})
			c.nextMustParameter(parameterElement{
				Property:    "rich_text",
				Type:        "array of rich text objects",
				Description: "The rich text of the heading.",
			}, func(e parameterElement) {
				heading.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "color",
				Type:        "string (enum)",
				Description: "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
			}, func(e parameterElement) {
				heading.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:    "is_toggleable",
				Type:        "boolean",
				Description: "Whether or not the heading block is a toggle heading or not. If true, then the heading block toggles and can support children. If false, then the heading block is a static heading block.",
			}, func(e parameterElement) {
				heading.addFields(e.asField(jen.Bool()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"heading_1\",\n  //...other keys excluded\n  \"heading_1\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\",\n    \"is_toggleable\": false\n  }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"heading_2\",\n  //...other keys excluded\n  \"heading_2\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\",\n    \"is_toggleable\": false\n  }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"heading_3\",\n  //...other keys excluded\n  \"heading_3\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\",\n    \"is_toggleable\": false\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Image */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Image",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Image block objects contain a file object detailing information about the image.",
			}, func(e blockElement) {
				block.addAdaptiveFieldWithType("image", e.Text, jen.Op("*").Id("File"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"image\",\n  //...other keys excluded\n  \"image\": {\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://website.domain/images/image.png\"\n    }\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Supported image types */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Supported image types",
			}, func(e blockElement) {
				// TODO 確認する
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The image must be directly hosted. In other words, the url cannot point to a service that retrieves the image. The following image types are supported:",
			}, func(e blockElement) {
				// TODO 確認する
			})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: ".bmp.gif.heic.jpeg.jpg.png.svg.tif.tiff",
			}, func(e blockElement) {
				// TODO 確認する
			})
		},
		func(c *comparator, b *builder) /* Link Preview */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Link Preview",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Link Preview block objects contain the originally pasted url:",
			}, func(e blockElement) {
				linkPreview := block.addAdaptiveFieldWithSpecificObject("link_preview", e.Text, b)
				linkPreview.addFields(&field{name: "url", typeCode: jen.String()})
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"link_preview\",\n  //...other keys excluded\n  \"link_preview\": {\n    \"url\": \"https://github.com/example/example-repo/pull/1234\"\n  }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧The link_preview block can only be returned as part of a response. The API does not support creating or appending link_preview blocks.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Mention */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Mention",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A mention block object is a child of a rich text object that is nested within a paragraph block object. This block type represents any @ tag in the Notion UI, for a user, date, Notion page, Notion database, or a miniaturized version of a Link Preview.",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A mention block object contains the following fields:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "A constant string representing the type of the mention.",
				ExampleValue: "",
				Property:     "type",
				Type:         "\"database\"  \n  \n\"date\"  \n  \n\"link_preview\"  \n  \n\"page\"  \n  \n\"user\"",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "An object with type-specific information about the mention.",
				ExampleValue: "",
				Property:     "\"database\"  \n  \n\"date\"  \n  \n\"link_preview\"  \n  \n\"page\"  \n  \n\"user\"",
				Type:         "object",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"page\",\n  \"page\": {\n    \"id\": \"3c612f56-fdd0-4a30-a4d6-bda7d7426309\"\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Numbered list item */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Numbered list item",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Numbered list item block objects contain the following information within the numbered_list_item property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The rich text displayed in the numbered_list_item block.",
				ExampleValue: "",
				Property:     "rich_text",
				Type:         "array of rich text objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
				ExampleValue: "",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The nested child blocks (if any) of the numbered_list_item block.",
				ExampleValue: "",
				Property:     "children",
				Type:         "array of block objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"numbered_list_item\",\n  \"numbered_list_item\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Finish reading the docs\",\n          \"link\": null\n        }\n      }\n    ],\n    \"color\": \"default\"\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Paragraph */ {
			var paragraph *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Paragraph",
			}, func(e blockElement) {
				paragraph = block.addAdaptiveFieldWithSpecificObject("paragraph", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Paragraph block objects contain the following information within the paragraph property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "rich_text",
				Type:        "array of rich text objects",
				Description: "The rich text displayed in the paragraph block.",
			}, func(e parameterElement) {
				paragraph.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "color",
				Type:        "string (enum)",
				Description: "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
			}, func(e parameterElement) {
				paragraph.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:    "children",
				Type:        "array of block objects",
				Description: "The nested child blocks (if any) of the paragraph block.",
			}, func(e parameterElement) {
				paragraph.addFields(e.asField(jen.Index().Id("Block"), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"paragraph\",\n  //...other keys excluded\n  \"paragraph\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Lacinato kale\",\n        \"link\": null\n      }\n    }],\n    \"color\": \"default\"\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n//...other keys excluded\n\t\"type\": \"paragraph\",\n  \t\"paragraph\":{\n  \t\t\"rich_text\": [\n    \t\t{\n      \t\t\"type\": \"mention\",\n      \t\t\"mention\": {\n        \t\t\"type\": \"date\",\n        \t\t\"date\": {\n          \t\t\"start\": \"2023-03-01\",\n          \t\t\"end\": null,\n          \t\t\"time_zone\": null\n        \t\t}\n      \t\t},\n      \t\t\"annotations\": {\n        \t\t\"bold\": false,\n        \t\t\"italic\": false,\n        \t\t\"strikethrough\": false,\n        \t\t\"underline\": false,\n        \t\t\"code\": false,\n        \t\t\"color\": \"default\"\n      \t\t},\n      \t\t\"plain_text\": \"2023-03-01\",\n      \t\t\"href\": null\n    \t\t},\n    \t\t{\n          \"type\": \"text\",\n      \t\t\"text\": {\n        \t\t\"content\": \" \",\n        \t\t\"link\": null\n      \t\t},\n      \t\t\"annotations\": {\n        \t\t\"bold\": false,\n        \t\t\"italic\": false,\n        \t\t\"strikethrough\": false,\n        \t\t\"underline\": false,\n        \t\t\"code\": false,\n        \t\t\"color\": \"default\"\n      \t\t},\n      \t\t\"plain_text\": \" \",\n      \t\t\"href\": null\n    \t\t}\n  \t\t],\n  \t\t\"color\": \"default\"\n  \t}\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* PDF */ {
			var pdf *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "PDF",
			}, func(e blockElement) {
				pdf = block.addAdaptiveFieldWithSpecificObject("pdf", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A PDF block object represents a PDF that has been embedded within a Notion page. It contains the following fields:",
			}, func(e blockElement) {
				pdf.addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:     "caption",
				Type:         "array of rich text objects",
				Description:  "A caption, if provided, for the PDF block.",
				ExampleValue: "",
			}, func(e parameterElement) {
				pdf.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         "\"external\"  \n  \n\"file\"",
				Description:  "A constant string representing the type of PDF. file indicates a Notion-hosted file, and external represents a third-party link.",
				ExampleValue: "",
			}, func(e parameterElement) {
				pdf.addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Property:    "external  \n  \nfile",
				Type:        "file object",
				Description: "An object containing type-specific information about the PDF.",
			}, func(e parameterElement) {
				pdf.addFields(
					&field{name: "external", typeCode: jen.Op("*").Id("FileExternal"), comment: e.Description, omitEmpty: true},
					&field{name: "file", typeCode: jen.Op("*").Id("FileFile"), comment: e.Description, omitEmpty: true},
				)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n\t\"type\": \"pdf\",\n  //...other keys excluded\n  \"pdf\": {\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://website.domain/files/doc.pdf\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Quote */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Quote",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Quote block objects contain the following information within the quote property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The rich text displayed in the quote block.",
				ExampleValue: "",
				Property:     "rich_text",
				Type:         "array of rich text objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
				ExampleValue: "",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The nested child blocks, if any, of the quote block.",
				ExampleValue: "",
				Property:     "children",
				Type:         "array of block objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n\t//...other keys excluded\n\t\"type\": \"quote\",\n   //...other keys excluded\n   \"quote\": {\n   \t\"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"To be or not to be...\",\n        \"link\": null\n      },\n    \t//...other keys excluded\n    }],\n    //...other keys excluded\n    \"color\": \"default\"\n   }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Synced block */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Synced block",
			}, func(e blockElement) {
				specificObject = block.addAdaptiveFieldWithSpecificObject("synced_block", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Similar to the Notion UI, there are two versions of a synced_block object: the original block that was created first and doesn't yet sync with anything else, and the duplicate block or blocks synced to the original.",
			}, func(e blockElement) {
				specificObject.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘An original synced block must be created before corresponding duplicate block or blocks can be made.",
			}, func(e blockElement) {
				specificObject.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Original synced block",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Original synced block objects contain the following information within the synced_block property:",
			}, func(e blockElement) {
				specificObject.addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:    "synced_from",
				Type:        "null",
				Description: "The value is always null to signify that this is an original synced block that does not refer to another block.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Id("SyncedFrom")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "children",
				Type:        "array of block objects",
				Description: "The nested child blocks, if any, of the synced_block block. These blocks will be mirrored in the duplicate synced_block.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Index().Id("Block"), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n    //...other keys excluded\n  \t\"type\": \"synced_block\",\n    \"synced_block\": {\n        \"synced_from\": null,\n        \"children\": [\n            {\n                \"callout\": {\n                    \"rich_text\": [\n                        {\n                            \"type\": \"text\",\n                            \"text\": {\n                                \"content\": \"Callout in synced block\"\n                            }\n                        }\n                    ]\n                }\n            }\n        ]\n    }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Duplicate synced block",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Duplicate synced block objects contain the following information within the synced_from object:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The type of the synced from object.  \n  \nPossible values are:  \n- \"block_id\"",
				ExampleValue: "",
				Property:     "type",
				Type:         "string (enum)",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "An identifier for the original synced_block.",
				ExampleValue: "",
				Property:     "block_id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {
				b.addConcreteObject("SyncedFrom", "").addFields(e.asField(UUID))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n    //...other keys excluded\n  \t\"type\": \"synced_block\",\n    \"synced_block\": {\n        \"synced_from\": {\n            \"block_id\": \"original_synced_block_id\"\n        }\n    }\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧The API does not supported updating synced block content.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Table */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Table",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Table block objects are parent blocks for table row children. Table block objects contain the following fields within the table property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number of columns in the table.  \n  \nNote that this cannot be changed via the public API once a table is created.",
				ExampleValue: "",
				Property:     "table_width",
				Type:         "integer",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the table has a column header. If true, then the first row in the table appears visually distinct from the other rows.",
				ExampleValue: "",
				Property:     "has_column_header",
				Type:         "boolean",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the table has a header row. If true, then the first column in the table appears visually distinct from the other columns.",
				ExampleValue: "",
				Property:     "has_row_header",
				Type:         "boolean",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"table\",\n  \"table\": {\n    \"table_width\": 2,\n    \"has_column_header\": false,\n    \"has_row_header\": false\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧 table_width can only be set when the table is first created.Note that the number of columns in a table can only be set when the table is first created. Calls to the Update block endpoint to update table_width fail.",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Table rows */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Table rows",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Follow these steps to fetch the table_rows of a table:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Get the table ID from a query to Retrieve block children for the parent page.Get the table_rows from a query to Retrieve block children for the table.",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A table_row block object contains the following fields within the table_row property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "An array of cell contents in horizontal display order. Each cell is an array of rich text objects.",
				ExampleValue: "",
				Property:     "cells",
				Type:         "array of array of rich text objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"table_row\",\n  \"table_row\": {\n    \"cells\": [\n      [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"column 1 content\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"column 1 content\",\n          \"href\": null\n        }\n      ],\n      [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"column 2 content\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"column 2 content\",\n          \"href\": null\n        }\n      ],\n      [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"column 3 content\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"column 3 content\",\n          \"href\": null\n        }\n      ]\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘When creating a table block via the Append block children endpoint, the table must have at least one table_row whose cells array has the same length as the table_width.",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Table of contents */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Table of contents",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Table of contents block objects contain the following information within the table_of_contents property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
				ExampleValue: "",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n\t\"type\": \"table_of_contents\",\n  \"table_of_contents\": {\n  \t\"color\": \"default\"\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Template */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Template",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "❗️ Deprecation NoticeAs of March 27, 2023 creation of template blocks will no longer be supported.",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Template blocks represent template buttons in the Notion UI.",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Template block objects contain the following information within the template property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The rich text displayed in the title of the template.",
				ExampleValue: "",
				Property:     "rich_text",
				Type:         "array of rich text objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The nested child blocks, if any, of the template block. These blocks are duplicated when the template block is used in the UI.",
				ExampleValue: "",
				Property:     "children",
				Type:         "array of block objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"template\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Add a new to-do\",\n          \"link\": null\n        },\n        \"annotations\": {\n          //...other keys excluded\n        },\n        \"plain_text\": \"Add a new to-do\",\n        \"href\": null\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* To do */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "To do",
			}, func(e blockElement) {
				specificObject = block.addAdaptiveFieldWithSpecificObject("to_do", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "To do block objects contain the following information within the to_do property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:    "rich_text",
				Type:        "array of rich text objects",
				Description: "The rich text displayed in the To do block.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Index().Id("RichText")))
			})
			c.nextMustParameter(parameterElement{
				Property:    "checked",
				Type:        "boolean (optional)",
				Description: "Whether the To do is checked.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:    "color",
				Type:        "string (enum)",
				Description: "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:    "children",
				Type:        "array of block objects",
				Description: "The nested child blocks, if any, of the To do block.",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Index().Id("Block"), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"to_do\",\n  \"to_do\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Finish Q3 goals\",\n        \"link\": null\n      }\n    }],\n    \"checked\": false,\n    \"color\": \"default\",\n    \"children\":[{\n      \"type\": \"paragraph\"\n      // ..other keys excluded\n    }]\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Toggle blocks */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Toggle blocks",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Toggle block objects contain the following information within the toggle property:",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The rich text displayed in the Toggle block.",
				ExampleValue: "",
				Property:     "rich_text",
				Type:         "array of rich text objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the block. Possible values are:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n-  \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n- \"orange_background\"  \n- \"yellow\"  \n- \"green\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background\"  \n- \"yellow_background\"",
				ExampleValue: "",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustParameter(parameterElement{
				Description:  "The nested child blocks, if any, of the Toggle block.",
				ExampleValue: "",
				Property:     "children",
				Type:         "array of block objects",
			}, func(e parameterElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  //...other keys excluded\n  \"type\": \"toggle\",\n  \"toggle\": {\n    \"rich_text\": [{\n      \"type\": \"text\",\n      \"text\": {\n        \"content\": \"Additional project details\",\n        \"link\": null\n      }\n      //...other keys excluded\n    }],\n    \"color\": \"default\",\n    \"children\":[{\n      \"type\": \"paragraph\"\n      // ..other keys excluded\n    }]\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Video */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Video",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Video block objects contain a file object detailing information about the image.",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"video\",\n  //...other keys excluded\n  \"video\": {\n    \"type\": \"external\",\n    \"external\": {\n \t  \t\"url\": \"https://companywebsite.com/files/video.mp4\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Supported video types */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Supported video types",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: ".amv.asf.avi.f4v.flv.gifv.mkv.mov.mpg.mpeg.mpv.mp4.m4v.qt.wmvYouTube video links that include embed or watch.E.g. https://www.youtube.com/watch?v=[id], https://www.youtube.com/embed/[id]",
			}, func(e blockElement) {
				// TODO
			})
		},
	)
}
