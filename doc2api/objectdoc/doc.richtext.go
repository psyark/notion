package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/rich-text",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "Rich text objects contain the data that Notion uses to display formatted text, mentions, and inline equations. Arrays of rich text objects within database property objects and page property value objects are used to create what a user experiences as a single text value in Notion.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addAbstractObject("RichText", "type", e.Text).listName = "RichTextArray"
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"Some words \",\n    \"link\": null\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"Some words \",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getAbstractObject("RichText").comment += "\n\n" + e.Code
					return nil
				},
			}}},
			&objectDocParagraphElement{
				Text:   "Each rich text object contains the following fields.",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocParametersElement{{
				Field:        "type",
				Type:         "string (enum)",
				Description:  "The type of this rich text object. Possible type values are: \"text\", \"mention\", \"equation\".",
				ExampleValue: `"text"`,
				output: func(e *objectDocParameter, b *builder) error {
					return nil // 各structで定義
				},
			}, {
				Field:        "text | mention | equation",
				Type:         "object",
				Description:  "An object containing type-specific configuration. \n\nRefer to the rich text type objects section below for details on type-specific values.",
				ExampleValue: "Refer to the rich text type objects section below for examples.",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // 各structで定義
				},
			}, {
				Field:        "annotations",
				Type:         "object",
				Description:  "The information used to style the rich text object. Refer to the annotation object section below for details.",
				ExampleValue: "Refer to the annotation object section below for examples.",
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("RichText").addFields(e.asField(jen.Id("Annotations"), false))
					return nil
				},
			}, {
				Field:        "plain_text",
				Type:         "string",
				Description:  "The plain text without annotations.",
				ExampleValue: `"Some words "`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("RichText").addFields(e.asField(jen.String(), false))
					return nil
				},
			}, {
				Field:        "href",
				Type:         "string (optional)",
				Description:  "The URL of any link or Notion mention in this text, if any.",
				ExampleValue: `"https://www.notion.so/Avocado-d093f1d200464ce78b36e58a3f0d8043"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("RichText").addFields(e.asField(jen.Op("*").String(), false)) // RetrivePageでnullを確認
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "The annotation object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.addSpecificObject("Annotations", e.Text)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nAll rich text objects contain an annotations object that sets the styling for the rich text. annotations includes the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("Annotations").comment += "\n" + e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "bold",
				Type:         "boolean",
				Description:  "Whether the text is bolded.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Annotations").addFields(e.asField(jen.Bool(), false))
					return nil
				},
			}, {
				Property:     "italic",
				Type:         "boolean",
				Description:  "Whether the text is italicized.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Annotations").addFields(e.asField(jen.Bool(), false))
					return nil
				},
			}, {
				Property:     "strikethrough",
				Type:         "boolean",
				Description:  "Whether the text is struck through.",
				ExampleValue: "false",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Annotations").addFields(e.asField(jen.Bool(), false))
					return nil
				},
			}, {
				Property:     "underline",
				Type:         "boolean",
				Description:  "Whether the text is underlined.",
				ExampleValue: "false",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Annotations").addFields(e.asField(jen.Bool(), false))
					return nil
				},
			}, {
				Property:     "code",
				Type:         "boolean",
				Description:  "Whether the text is code style.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Annotations").addFields(e.asField(jen.Bool(), false))
					return nil
				},
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the text. Possible values include: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n- \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n-\"orange_background\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background”\n- \"yellow\"\n- \"yellow_background\"",
				ExampleValue: `"green"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Annotations").addFields(e.asField(jen.String(), false))
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "Rich text type objects",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.addComment(e.Text)
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Equation",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("RichText").addVariant(
						b.addSpecificObject("EquationRichText", e.Text).addFields(
							&fixedStringField{name: "type", value: "equation"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nNotion supports inline LaTeX equations as rich text object’s with a type value of \"equation\". The corresponding equation type object contains the following: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("EquationRichText").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "expression",
				Type:         "string",
				Description:  "The LaTeX string representing the inline equation.",
				ExampleValue: `"\frac{{ - b \pm \sqrt {b^2 - 4ac} }}{{2a}}"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("EquationRichText").addFields(e.asField(jen.String(), false))
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "Example rich text equation object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("EquationRichText").comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"equation\",\n  \"equation\": {\n    \"expression\": \"E = mc^2\"\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"E = mc^2\",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("EquationRichText").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Mention",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.addAbstractObject("Mention", "type", e.Text)
					b.getAbstractObject("RichText").addVariant(
						b.addSpecificObject("MentionRichText", e.Text).addFields(
							&fixedStringField{name: "type", value: "mention"},
							&field{name: "mention", typeCode: jen.Id("Mention")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nMention objects represent an inline mention of a database, date, link preview mention, page, template mention, or user. A mention is created in the Notion UI when a user types\u00a0@\u00a0followed by the name of the reference.\n\nIf a rich text object’s type value is \"mention\", then the corresponding mention object contains the following:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("Mention").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "type",
				Type:         "string\u00a0(enum)",
				Description:  "The type of the inline mention. Possible values include:\n\n- \"database\"\n- \"date\"\n- \"link_preview\"\n- \"page\"\n- \"template_mention\"\n- \"user\"",
				ExampleValue: `"user"`,
				output:       func(e *objectDocParameter, b *builder) error { return nil },
			}, {
				Field:        "database | date | link_preview | page | template_mention | user",
				Type:         "object",
				Description:  "An object containing type-specific configuration. Refer to the mention type object sections below for details.",
				ExampleValue: "Refer to the mention type object sections below for example values.",
				output:       func(e *objectDocParameter, b *builder) error { return nil },
			}},
			&objectDocHeadingElement{
				Text: "Database mention type object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Mention").addVariant(
						b.addSpecificObject("DatabaseMention", e.Text).addFields(
							&fixedStringField{name: "type", value: "database"},
							&field{name: "database", typeCode: jen.Id("PageReference")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nDatabase mentions contain a database reference within the corresponding\u00a0database\u00a0field. A database reference is an object with an\u00a0id\u00a0key and a string value (UUIDv4) corresponding to a database ID.\n\nIf an integration doesn’t have access to the mentioned database, then the mention is returned with just the ID. The plain_text value that would be a title appears as \"Untitled\" and the annotation object’s values are defaults.\n\nExample rich text mention object for a database mention ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DatabaseMention").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"database\",\n    \"database\": {\n      \"id\": \"a1d8501e-1ac1-43e9-a6bd-ea9fe6c8822b\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"Database with test things\",\n  \"href\": \"https://www.notion.so/a1d8501e1ac143e9a6bdea9fe6c8822b\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("DatabaseMention").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Date mention type object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Mention").addVariant(
						b.addSpecificObject("DateMention", e.Text).addFields(
							&fixedStringField{name: "type", value: "date"},
							&field{name: "date", typeCode: jen.Id("DatePropertyValue")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate mentions contain a\u00a0date property value object\u00a0within the corresponding\u00a0date\u00a0field.\n\nExample rich text mention object for a date mention",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DateMention").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2022-12-16\",\n      \"end\": null\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"2022-12-16\",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("DateMention").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Link Preview mention type object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Mention").addVariant(
						b.addSpecificObject("LinkPreviewMention", e.Text).addFields(
							&fixedStringField{name: "type", value: "link_preview"},
							&field{name: "link_preview", typeCode: jen.Id("URLReference")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nIf a user opts to share a Link Preview as a mention, then the API handles the Link Preview mention as a rich text object with a type value of link_preview. Link preview rich text mentions contain a corresponding link_preview object that includes the url that is used to create the Link Preview mention.\n\nExample rich text mention object for a link_preview mention ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("LinkPreviewMention").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"link_preview\",\n    \"link_preview\": {\n      \"url\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\",\n  \"href\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("LinkPreviewMention").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Page mention type object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Mention").addVariant(
						b.addSpecificObject("PageMention", e.Text).addFields(
							&fixedStringField{name: "type", value: "page"},
							&field{name: "page", typeCode: jen.Id("PageReference")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nPage mentions contain a page reference within the corresponding\u00a0page\u00a0field. A page reference is an object with an\u00a0id\u00a0property and a string value (UUIDv4) corresponding to a page ID.\n\nIf an integration doesn’t have access to the mentioned page, then the mention is returned with just the ID. The plain_text value that would be a title appears as \"Untitled\" and the annotation object’s values are defaults.\n\nExample rich text mention object for a page mention ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PageMention").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"page\",\n    \"page\": {\n      \"id\": \"3c612f56-fdd0-4a30-a4d6-bda7d7426309\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"This is a test page\",\n  \"href\": \"https://www.notion.so/3c612f56fdd04a30a4d6bda7d7426309\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("PageMention").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Template mention type object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Mention").addVariant(
						b.addSpecificObject("TemplateMention", e.Text).addFields(
							&fixedStringField{name: "type", value: "template_mention"},
							&field{name: "template_mention", typeCode: jen.Id("TemplateMentionData")},
						),
					)
					b.addAbstractObject("TemplateMentionData", "type", "")
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nThe content inside a template button in the Notion UI can include placeholder date and user mentions that populate when a template is duplicated. Template mention type objects contain these populated values. \n\nTemplate mention rich text objects contain a\u00a0template_mention\u00a0object with a nested\u00a0type\u00a0key that is either\u00a0\"template_mention_date\"\u00a0or\u00a0\"template_mention_user\".\n\nIf the\u00a0type\u00a0key is\u00a0\"template_mention_date\", then the rich text object contains the following template_mention_date field:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("TemplateMention").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "template_mention_date",
				Type:         "string\u00a0(enum)",
				Description:  "The type of the date mention. Possible values include:\u00a0\"today\"\u00a0and\u00a0\"now\".",
				ExampleValue: `"today"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("TemplateMentionData").addVariant(
						b.addSpecificObject("TemplateMentionDate", "").addFields(
							&fixedStringField{name: "type", value: e.Field},
							&field{name: e.Field, typeCode: jen.String(), comment: e.Description},
						),
					)
					return nil
				},
			}},
			&objectDocParagraphElement{
				Text: "Example rich text mention object for a template_mention_date mention ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("TemplateMentionDate").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"template_mention\",\n    \"template_mention\": {\n      \"type\": \"template_mention_date\",\n      \"template_mention_date\": \"today\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Today\",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("TemplateMentionDate").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocParagraphElement{
				Text: "If the\u00a0type\u00a0key is\u00a0\"template_mention_user\", then the rich text object contains the following template_mention_user field: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "template_mention_user",
				Type:         "string\u00a0(enum)",
				Description:  "The type of the user mention. The only possible value is\u00a0\"me\".",
				ExampleValue: `"me"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("TemplateMentionData").addVariant(
						b.addSpecificObject("TemplateMentionUser", "").addFields(
							&fixedStringField{name: "type", value: e.Field},
							e.asFixedStringField(),
						),
					)
					return nil
				},
			}},
			&objectDocParagraphElement{
				Text: "Example rich text mention object for a template_mention_user mention ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("TemplateMentionUser").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"template_mention\",\n    \"template_mention\": {\n      \"type\": \"template_mention_user\",\n      \"template_mention_user\": \"me\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Me\",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("TemplateMentionUser").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "User mention type object",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Mention").addVariant(
						b.addSpecificObject("UserMention", e.Text).addFields(
							&fixedStringField{name: "type", value: "user"},
							&field{name: "user", typeCode: jen.Id("PartialUser")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nIf a rich text object’s type value is \"user\", then the corresponding user field contains a\u00a0user object. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("UserMention").comment += e.Text
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:  "If your integration doesn’t yet have access to the mentioned user, then the `plain_text` that would include a user’s name reads as `\"@Anonymous\"`. To update the integration to get access to the user, update the integration capabilities on the integration settings page.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("UserMention").comment += "\n" + e.Body
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "Example rich text mention object for a user mention",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("UserMention").comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"user\",\n    \"user\": {\n      \"object\": \"user\",\n      \"id\": \"b2e19928-b427-4aad-9a9d-fde65479b1d9\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Anonymous\",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("UserMention").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Text ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("RichText").addVariant(
						b.addSpecificObject("TextRichText", e.Text).addFields(
							&fixedStringField{name: "type", value: "text"},
							&field{name: "text", typeCode: jen.Id("Text")},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nIf a rich text object’s type value is \"text\", then the corresponding text field contains an object including the following:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addSpecificObject("Text", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "content",
				Type:         "string",
				Description:  "The actual text content of the text.",
				ExampleValue: `"Some words "`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Text").addFields(e.asField(jen.String(), false))
					return nil
				},
			}, {
				Field:        "link",
				Type:         "object\u00a0(optional)",
				Description:  "An object with information about any inline link in this text, if included. \n\nIf the text contains an inline link, then the object key is url and the value is the URL’s string web address. \n\nIf the text doesn’t have any inline links, then the value is null.",
				ExampleValue: "{\n  \"url\": \"https://developers.notion.com/\"\n}",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("Text").addFields(e.asField(jen.Op("*").Id("URLReference"), false)) // RetrivePageでnullを確認
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "Example rich text text object without link ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("Text").comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"This is an \",\n    \"link\": null\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"This is an \",\n  \"href\": null\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("Text").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Example rich text text object with link ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("Text").comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"inline link\",\n    \"link\": {\n      \"url\": \"https://developers.notion.com/\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"inline link\",\n  \"href\": \"https://developers.notion.com/\"\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					b.getSpecificObject("Text").comment += "\n" + e.Code
					return nil
				},
			}}},
			&objectDocCalloutElement{
				Body:  "Refer to the request limits documentation page for information about [limits on the size of rich text objects](https://developers.notion.com/reference/request-limits#limits-for-property-values).",
				Title: "Rich text object limits",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getAbstractObject("RichText").comment += "\n\n" + e.Title + "\n" + e.Body
					return nil
				},
			},
		},
	})
}
