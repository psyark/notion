package objectdoc

import (
	"encoding/json"

	"github.com/dave/jennifer/jen"
)

func init() {
	registerTranslator(
		"https://developers.notion.com/reference/rich-text",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Rich text objects contain the data that Notion uses to display formatted text, mentions, and inline equations. Arrays of rich text objects within database property objects and page property value objects are used to create what a user experiences as a single text value in Notion.",
			}, func(e blockElement) {
				b.addAdaptiveObject("RichText", "type", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"Some words \",\n    \"link\": null\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"Some words \",\n  \"href\": null\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Each rich text object contains the following fields.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         "string (enum)",
				Description:  "The type of this rich text object. Possible type values are: \"text\", \"mention\", \"equation\".",
				ExampleValue: `"text"`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "An object containing type-specific configuration. \n\nRefer to the rich text type objects section below for details on type-specific values.",
				ExampleValue: "Refer to the rich text type objects section below for examples.",
				Property:     "text | mention | equation",
				Type:         "object",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The information used to style the rich text object. Refer to the annotation object section below for details.",
				ExampleValue: "Refer to the annotation object section below for examples.",
				Property:     "annotations",
				Type:         "object",
			}, func(e parameterElement) {
				getSymbol[adaptiveObject](b, "RichText").addFields(e.asField(jen.Id("Annotations")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "plain_text",
				Type:         "string",
				Description:  "The plain text without annotations.",
				ExampleValue: `"Some words "`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject](b, "RichText").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "href",
				Type:         "string (optional)",
				Description:  "The URL of any link or Notion mention in this text, if any.",
				ExampleValue: `"https://www.notion.so/Avocado-d093f1d200464ce78b36e58a3f0d8043"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject](b, "RichText").addFields(e.asField(NullString)) // RetrivePageでnullを確認
			})
		},
		func(c *comparator, b *builder) /* The annotation object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "The annotation object",
			}, func(e blockElement) {
				b.addConcreteObject("Annotations", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "All rich text objects contain an annotations object that sets the styling for the rich text. annotations includes the following fields:",
			}, func(e blockElement) {
				getSymbol[concreteObject](b, "Annotations").addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:     "bold",
				Type:         "boolean",
				Description:  "Whether the text is bolded.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Annotations").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "italic",
				Type:         "boolean",
				Description:  "Whether the text is italicized.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Annotations").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the text is struck through.",
				ExampleValue: "false",
				Property:     "strikethrough",
				Type:         "boolean",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Annotations").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the text is underlined.",
				ExampleValue: "false",
				Property:     "underline",
				Type:         "boolean",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Annotations").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "code",
				Type:         "boolean",
				Description:  "Whether the text is code style.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Annotations").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the text. Possible values include: \n\n- \"blue\"\n- \"blue_background\"\n- \"brown\"\n- \"brown_background\"\n- \"default\"\n- \"gray\"\n- \"gray_background\"\n- \"green\"\n- \"green_background\"\n- \"orange\"\n-\"orange_background\"\n- \"pink\"\n- \"pink_background\"\n- \"purple\"\n- \"purple_background\"\n- \"red\"\n- \"red_background”\n- \"yellow\"\n- \"yellow_background\"",
				ExampleValue: `"green"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "Annotations").addFields(e.asField(jen.String(), omitEmpty))
			})
		},
		func(c *comparator, b *builder) /* Rich text type objects */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rich text type objects",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Equation */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Equation",
			}, func(e blockElement) {
				b.addAdaptiveField("RichText", "equation", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Notion supports inline LaTeX equations as rich text object’s with a type value of \"equation\". The corresponding equation type object contains the following:",
			}, func(e blockElement) {
				getSymbol[concreteObject](b, "RichTextEquation").addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:     "expression",
				Type:         "string",
				Description:  "The LaTeX string representing the inline equation.",
				ExampleValue: `"\frac{{ - b \pm \sqrt {b^2 - 4ac} }}{{2a}}"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "RichTextEquation").addFields(e.asField(jen.String()))
			})
		},
		func(c *comparator, b *builder) /* Example rich text equation object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example rich text equation object",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"equation\",\n  \"equation\": {\n    \"expression\": \"E = mc^2\"\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"E = mc^2\",\n  \"href\": null\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Mention */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Mention",
			}, func(e blockElement) {
				b.addAdaptiveObject("Mention", "type", e.Text)
				b.addAdaptiveFieldWithType("RichText", "mention", e.Text, jen.Op("*").Id("Mention"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Mention objects represent an inline mention of a database, date, link preview mention, page, template mention, or user. A mention is created in the Notion UI when a user types\u00a0@\u00a0followed by the name of the reference.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject](b, "Mention").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If a rich text object’s type value is \"mention\", then the corresponding mention object contains the following:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         "string\u00a0(enum)",
				Description:  "The type of the inline mention. Possible values include:\n\n- \"database\"\n- \"date\"\n- \"link_preview\"\n- \"page\"\n- \"template_mention\"\n- \"user\"",
				ExampleValue: `"user"`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "database | date | link_preview | page | template_mention | user",
				Type:         "object",
				Description:  "An object containing type-specific configuration. Refer to the mention type object sections below for details.",
				ExampleValue: "Refer to the mention type object sections below for example values.",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* Database mention type object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Database mention type object",
			}, func(e blockElement) {
				b.addAdaptiveFieldWithType("Mention", "database", e.Text, jen.Op("*").Id("PageReference"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Database mentions contain a database reference within the corresponding\u00a0database\u00a0field. A database reference is an object with an\u00a0id\u00a0key and a string value (UUIDv4) corresponding to a database ID.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If an integration doesn’t have access to the mentioned database, then the mention is returned with just the ID. The plain_text value that would be a title appears as \"Untitled\" and the annotation object’s values are defaults.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Example rich text mention object for a database mention",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"database\",\n    \"database\": {\n      \"id\": \"a1d8501e-1ac1-43e9-a6bd-ea9fe6c8822b\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"Database with test things\",\n  \"href\": \"https://www.notion.so/a1d8501e1ac143e9a6bdea9fe6c8822b\"\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Date mention type object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date mention type object",
			}, func(e blockElement) {
				b.addAdaptiveFieldWithType("Mention", "date", e.Text, jen.Op("*").Id("PropertyValueDate"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date mentions contain a\u00a0date property value object\u00a0within the corresponding\u00a0date\u00a0field.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Example rich text mention object for a date mention",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2022-12-16\",\n      \"end\": null\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"2022-12-16\",\n  \"href\": null\n}",
			}, func(e blockElement) {
				// TODO これでいいか確認
				// おそらくドキュメントの不具合。time_zoneはomitemptyではなさそう
				var tmp map[string]any
				json.Unmarshal([]byte(e.Text), &tmp)
				tmp["mention"].(map[string]any)["date"].(map[string]any)["time_zone"] = nil
				data, _ := json.Marshal(tmp)
				b.addUnmarshalTest("RichText", string(data))
			})
		},
		func(c *comparator, b *builder) /* Link Preview mention type object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Link Preview mention type object",
			}, func(e blockElement) {
				b.addAdaptiveField("Mention", "link_preview", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If a user opts to share a Link Preview as a mention, then the API handles the Link Preview mention as a rich text object with a type value of link_preview. Link preview rich text mentions contain a corresponding link_preview object that includes the url that is used to create the Link Preview mention.",
			}, func(e blockElement) {
				getSymbol[concreteObject](b, "MentionLinkPreview").addFields(
					&field{name: "url", typeCode: jen.String()},
				).addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Example rich text mention object for a link_preview mention",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"link_preview\",\n    \"link_preview\": {\n      \"url\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\",\n  \"href\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\"\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Page mention type object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Page mention type object",
			}, func(e blockElement) {
				b.addAdaptiveFieldWithType("Mention", "page", e.Text, jen.Op("*").Id("PageReference"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Page mentions contain a page reference within the corresponding\u00a0page\u00a0field. A page reference is an object with an\u00a0id\u00a0property and a string value (UUIDv4) corresponding to a page ID.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If an integration doesn’t have access to the mentioned page, then the mention is returned with just the ID. The plain_text value that would be a title appears as \"Untitled\" and the annotation object’s values are defaults.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Example rich text mention object for a page mention",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"page\",\n    \"page\": {\n      \"id\": \"3c612f56-fdd0-4a30-a4d6-bda7d7426309\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"This is a test page\",\n  \"href\": \"https://www.notion.so/3c612f56fdd04a30a4d6bda7d7426309\"\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Template mention type object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Template mention type object",
			}, func(e blockElement) {
				b.addAdaptiveObject("TemplateMention", "type", e.Text)
				b.addAdaptiveFieldWithType("Mention", "template_mention", e.Text, jen.Op("*").Id("TemplateMention"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The content inside a template button in the Notion UI can include placeholder date and user mentions that populate when a template is duplicated. Template mention type objects contain these populated values.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject](b, "TemplateMention").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Template mention rich text objects contain a\u00a0template_mention\u00a0object with a nested\u00a0type\u00a0key that is either\u00a0\"template_mention_date\"\u00a0or\u00a0\"template_mention_user\".",
			}, func(e blockElement) {
				getSymbol[adaptiveObject](b, "TemplateMention").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If the\u00a0type\u00a0key is\u00a0\"template_mention_date\", then the rich text object contains the following template_mention_date field:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "template_mention_date",
				Type:         "string\u00a0(enum)",
				Description:  "The type of the date mention. Possible values include:\u00a0\"today\"\u00a0and\u00a0\"now\".",
				ExampleValue: `"today"`,
			}, func(e parameterElement) {
				b.addAdaptiveFieldWithType("TemplateMention", e.Property, e.Description, jen.String())
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "_Example rich text mention object for a template_mention_date mention _",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"template_mention\",\n    \"template_mention\": {\n      \"type\": \"template_mention_date\",\n      \"template_mention_date\": \"today\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Today\",\n  \"href\": null\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If the\u00a0type\u00a0key is\u00a0\"template_mention_user\", then the rich text object contains the following template_mention_user field:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "template_mention_user",
				Type:         "string\u00a0(enum)",
				Description:  "The type of the user mention. The only possible value is\u00a0\"me\".",
				ExampleValue: `"me"`,
			}, func(e parameterElement) {
				b.addAdaptiveFieldWithType("TemplateMention", e.Property, e.Description, jen.String())
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "_Example rich text mention object for a template_mention_user mention _",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"template_mention\",\n    \"template_mention\": {\n      \"type\": \"template_mention_user\",\n      \"template_mention_user\": \"me\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Me\",\n  \"href\": null\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* User mention type object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "User mention type object",
			}, func(e blockElement) {
				b.addAdaptiveFieldWithType("Mention", "user", e.Text, jen.Op("*").Id("User"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If a rich text object’s type value is \"user\", then the corresponding user field contains a\u00a0user object.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "If your integration doesn’t yet have access to the mentioned user, then the plain_text that would include a user’s name reads as \"@Anonymous\". To update the integration to get access to the user, update the integration capabilities on the integration settings page.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Example rich text mention object for a user mention",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"user\",\n    \"user\": {\n      \"object\": \"user\",\n      \"id\": \"b2e19928-b427-4aad-9a9d-fde65479b1d9\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Anonymous\",\n  \"href\": null\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Text */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Text",
			}, func(e blockElement) {
				b.addAdaptiveField("RichText", "text", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If a rich text object’s type value is \"text\", then the corresponding text field contains an object including the following:",
			}, func(e blockElement) {
				getSymbol[concreteObject](b, "RichTextText").addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:     "content",
				Type:         "string",
				Description:  "The actual text content of the text.",
				ExampleValue: `"Some words "`,
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "RichTextText").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "link",
				Type:         "object\u00a0(optional)",
				Description:  "An object with information about any inline link in this text, if included. \n\nIf the text contains an inline link, then the object key is url and the value is the URL’s string web address. \n\nIf the text doesn’t have any inline links, then the value is null.",
				ExampleValue: "{\n  \"url\": \"https://developers.notion.com/\"\n}",
			}, func(e parameterElement) {
				getSymbol[concreteObject](b, "RichTextText").addFields(e.asField(jen.Op("*").Id("URLReference"))) // RetrivePageでnullを確認
			})
		},
		func(c *comparator, b *builder) /* Example rich text text object without link */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example rich text text object without link",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"This is an \",\n    \"link\": null\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"This is an \",\n  \"href\": null\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Example rich text text object with link */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example rich text text object with link",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"inline link\",\n    \"link\": {\n      \"url\": \"https://developers.notion.com/\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"inline link\",\n  \"href\": \"https://developers.notion.com/\"\n}",
			}, func(e blockElement) {
				b.addUnmarshalTest("RichText", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Refer to the request limits documentation page for information about limits on the size of rich text objects.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject](b, "RichText").addComment(e.Text)
			})
		},
	)
}
