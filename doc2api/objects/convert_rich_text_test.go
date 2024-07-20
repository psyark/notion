package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestRichText(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/rich-text")

	var richText *UnionStruct

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Notion uses rich text to allow users to customize their content. Rich text refers to a type of document where content can be styled and formatted in a variety of customizable ways. This includes styling decisions, such as the use of italics, font size, and font color, as well as formatting, such as the use of hyperlinks or code blocks.",
	}).Output(func(e *Block, b *CodeBuilder) {
		richText = b.AddUnionStruct("RichText", "type", e.Text)
	})

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Notion includes rich text objects in block objects to indicate how blocks in a page are represented. Blocks that support rich text will include a rich text object; however, not all block types offer rich text.",
	}).Output(func(e *Block, b *CodeBuilder) {
		richText.AddComment(e.Text)
	})

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "When blocks are retrieved from a page using the Retrieve a block or Retrieve block children endpoints, an array of rich text objects will be included in the block object (when available). Developers can use this array to retrieve the plain text (plain_text) for the block or get all the rich text styling and formatting options applied to the block.",
	}).Output(func(e *Block, b *CodeBuilder) {
		richText.AddComment(e.Text)
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"Some words \",\n    \"link\": null\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"Some words \",\n  \"href\": null\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		b.AddUnmarshalTest("RichText", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘Many block types support rich text. In cases where it is supported, a rich_text object will be included in the block type object. All rich_text objects will include a plain_text property, which provides a convenient way for developers to access unformatted text from the Notion block."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Each rich text object contains the following fields."})

	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "string (enum)",
		Description:  "The type of this rich text object. Possible type values are: \"text\", \"mention\", \"equation\".",
		ExampleValue: `"text"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     `text \| mention \| equation`,
		Type:         "object",
		Description:  "An object containing type-specific configuration.  \n  \nRefer to the rich text type objects section below for details on type-specific values.",
		ExampleValue: "Refer to the rich text type objects section below for examples.",
	})
	c.ExpectParameter(&Parameter{
		Property:     "annotations",
		Type:         "object",
		Description:  "The information used to style the rich text object. Refer to the annotation object section below for details.",
		ExampleValue: "Refer to the annotation object section below for examples.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		richText.AddFields(b.NewField(e, jen.Id("Annotations")))
	})
	c.ExpectParameter(&Parameter{
		Property:     "plain_text",
		Type:         "string",
		Description:  "The plain text without annotations.",
		ExampleValue: `"Some words "`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		richText.AddFields(b.NewField(e, jen.String()))
	})
	c.ExpectParameter(&Parameter{
		Property:     "href",
		Type:         "string (optional)",
		Description:  "The URL of any link or Notion mention in this text, if any.",
		ExampleValue: `"https://www.notion.so/Avocado-d093f1d200464ce78b36e58a3f0d8043"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		richText.AddFields(b.NewField(e, jen.Op("*").String())) // RetrivePageでnullを確認
	})

	{
		var annotations *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "The annotation object",
		}).Output(func(e *Block, b *CodeBuilder) {
			annotations = b.AddSimpleObject("Annotations", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "All rich text objects contain an annotations object that sets the styling for the rich text. annotations includes the following fields:",
		}).Output(func(e *Block, b *CodeBuilder) {
			annotations.AddComment(e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "bold",
			Type:         "boolean",
			Description:  "Whether the text is bolded.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			annotations.AddFields(b.NewField(e, jen.Bool()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "italic",
			Type:         "boolean",
			Description:  "Whether the text is italicized.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			annotations.AddFields(b.NewField(e, jen.Bool()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "strikethrough",
			Type:         "boolean",
			Description:  "Whether the text is struck through.",
			ExampleValue: "false",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			annotations.AddFields(b.NewField(e, jen.Bool()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "underline",
			Type:         "boolean",
			Description:  "Whether the text is underlined.",
			ExampleValue: "false",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			annotations.AddFields(b.NewField(e, jen.Bool()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "code",
			Type:         "boolean",
			Description:  "Whether the text is code style.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			annotations.AddFields(b.NewField(e, jen.Bool()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "color",
			Type:         "string (enum)",
			Description:  "Color of the text. Possible values include:  \n  \n- \"blue\"  \n- \"blue_background\"  \n- \"brown\"  \n- \"brown_background\"  \n- \"default\"  \n- \"gray\"  \n- \"gray_background\"  \n- \"green\"  \n- \"green_background\"  \n- \"orange\"  \n-\"orange_background\"  \n- \"pink\"  \n- \"pink_background\"  \n- \"purple\"  \n- \"purple_background\"  \n- \"red\"  \n- \"red_background”  \n- \"yellow\"  \n- \"yellow_background\"",
			ExampleValue: `"green"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			annotations.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
	}

	{
		var richTextEquation *SimpleObject

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Rich text type objects"})
		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Equation",
		}).Output(func(e *Block, b *CodeBuilder) {
			richTextEquation = richText.AddAdaptiveFieldWithSpecificObject("equation", e.Text, b)
		})

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Notion supports inline LaTeX equations as rich text object’s with a type value of \"equation\". The corresponding equation type object contains the following:",
		}).Output(func(e *Block, b *CodeBuilder) {
			richTextEquation.AddComment(e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "expression",
			Type:         "string",
			Description:  "The LaTeX string representing the inline equation.",
			ExampleValue: `"\frac{{ - b \pm \sqrt {b^2 - 4ac} }}{{2a}}"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			richTextEquation.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectBlock(&Block{Kind: "Heading", Text: "Example rich text equation object"})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"equation\",\n  \"equation\": {\n    \"expression\": \"E = mc^2\"\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"E = mc^2\",\n  \"href\": null\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			b.AddUnmarshalTest("RichText", e.Text)
		})
	}

	{
		var mention *UnionStruct

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Mention",
		}).Output(func(e *Block, b *CodeBuilder) {
			mention = b.AddUnionStruct("Mention", "type", e.Text)
			richText.AddAdaptiveFieldWithType("mention", e.Text, jen.Op("*").Id("Mention"))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Mention objects represent an inline mention of a database, date, link preview mention, page, template mention, or user. A mention is created in the Notion UI when a user types\u00a0@\u00a0followed by the name of the reference.",
		}).Output(func(e *Block, b *CodeBuilder) {
			mention.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: `If a rich text object’s type value is "mention", then the corresponding mention object contains the following:`})
		c.ExpectParameter(&Parameter{
			Property:     "type",
			Type:         `string (enum)`,
			Description:  "The type of the inline mention. Possible values include:  \n  \n- \"database\"  \n- \"date\"  \n- \"link_preview\"  \n- \"page\"  \n- \"template_mention\"  \n- \"user\"",
			ExampleValue: `"user"`,
		})
		c.ExpectParameter(&Parameter{
			Property:     `database \| date \| link_preview \| page \| template_mention \| user`,
			Type:         "object",
			Description:  "An object containing type-specific configuration. Refer to the mention type object sections below for details.",
			ExampleValue: "Refer to the mention type object sections below for example values.",
		})
		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Database mention type object",
		}).Output(func(e *Block, b *CodeBuilder) {
			mention.AddAdaptiveFieldWithType("database", e.Text, jen.Op("*").Id("PageReference"))
		})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: `Database mentions contain a database reference within the corresponding database field. A database reference is an object with an id key and a string value (UUIDv4) corresponding to a database ID.`})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "If an integration doesn’t have access to the mentioned database, then the mention is returned with just the ID. The plain_text value that would be a title appears as \"Untitled\" and the annotation object’s values are defaults."})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Example rich text mention object for a database mention"})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"database\",\n    \"database\": {\n      \"id\": \"a1d8501e-1ac1-43e9-a6bd-ea9fe6c8822b\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"Database with test things\",\n  \"href\": \"https://www.notion.so/a1d8501e1ac143e9a6bdea9fe6c8822b\"\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			b.AddUnmarshalTest("RichText", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Date mention type object",
		}).Output(func(e *Block, b *CodeBuilder) {
			mention.AddAdaptiveFieldWithType("date", e.Text, jen.Op("*").Id("PropertyValueDate"))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: `Date mentions contain a date property value object within the corresponding date field.`,
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Example rich text mention object for a date mention",
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2022-12-16\",\n      \"end\": null\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"2022-12-16\",\n  \"href\": null\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			// TODO これでいいか確認
			// おそらくドキュメントの不具合。time_zoneはomitemptyではなさそう
			var tmp map[string]any
			json.Unmarshal([]byte(e.Text), &tmp)
			tmp["mention"].(map[string]any)["date"].(map[string]any)["time_zone"] = nil
			data, _ := json.Marshal(tmp)
			b.AddUnmarshalTest("RichText", string(data))
		})

		var mentionLinkPreview *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Link Preview mention type object",
		}).Output(func(e *Block, b *CodeBuilder) {
			mentionLinkPreview = mention.AddAdaptiveFieldWithSpecificObject("link_preview", e.Text, b)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "If a user opts to share a Link Preview as a mention, then the API handles the Link Preview mention as a rich text object with a type value of link_preview. Link preview rich text mentions contain a corresponding link_preview object that includes the url that is used to create the Link Preview mention.",
		}).Output(func(e *Block, b *CodeBuilder) {
			mentionLinkPreview.AddFields(b.NewField(&Parameter{Property: "url"}, jen.String()))
			mentionLinkPreview.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Example rich text mention object for a link_preview mention"})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"link_preview\",\n    \"link_preview\": {\n      \"url\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\",\n  \"href\": \"https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD\"\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			b.AddUnmarshalTest("RichText", e.Text)
		})

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Page mention type object",
		}).Output(func(e *Block, b *CodeBuilder) {
			mention.AddAdaptiveFieldWithType("page", e.Text, jen.Op("*").Id("PageReference"))
		})

		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Page mentions contain a page reference within the corresponding\u00a0page\u00a0field. A page reference is an object with an\u00a0id\u00a0property and a string value (UUIDv4) corresponding to a page ID."})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "If an integration doesn’t have access to the mentioned page, then the mention is returned with just the ID. The plain_text value that would be a title appears as \"Untitled\" and the annotation object’s values are defaults."})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Example rich text mention object for a page mention"})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"page\",\n    \"page\": {\n      \"id\": \"3c612f56-fdd0-4a30-a4d6-bda7d7426309\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"This is a test page\",\n  \"href\": \"https://www.notion.so/3c612f56fdd04a30a4d6bda7d7426309\"\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			b.AddUnmarshalTest("RichText", e.Text)
		})

		var templateMention *UnionStruct

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Template mention type object",
		}).Output(func(e *Block, b *CodeBuilder) {
			templateMention = b.AddUnionStruct("TemplateMention", "type", e.Text)
			mention.AddAdaptiveFieldWithType("template_mention", e.Text, jen.Op("*").Id("TemplateMention"))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The content inside a template button in the Notion UI can include placeholder date and user mentions that populate when a template is duplicated. Template mention type objects contain these populated values.",
		}).Output(func(e *Block, b *CodeBuilder) {
			templateMention.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: `Template mention rich text objects contain a template_mention object with a nested type key that is either "template_mention_date" or "template_mention_user".`,
		}).Output(func(e *Block, b *CodeBuilder) {
			templateMention.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: `If the type key is "template_mention_date", then the rich text object contains the following template_mention_date field:`,
		})
		c.ExpectParameter(&Parameter{
			Property:     "template_mention_date",
			Type:         `string (enum)`,
			Description:  `The type of the date mention. Possible values include: "today" and "now".`,
			ExampleValue: `"today"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			templateMention.AddFields(b.NewField(e, jen.String(), DiscriminatorValue(e.Property)))
		})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "_Example rich text mention object for a template_mention_date mention _"})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"template_mention\",\n    \"template_mention\": {\n      \"type\": \"template_mention_date\",\n      \"template_mention_date\": \"today\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Today\",\n  \"href\": null\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			b.AddUnmarshalTest("RichText", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "If the\u00a0type\u00a0key is\u00a0\"template_mention_user\", then the rich text object contains the following template_mention_user field:",
		})
		c.ExpectParameter(&Parameter{
			Property:     "template_mention_user",
			Type:         `string (enum)`,
			Description:  `The type of the user mention. The only possible value is "me".`,
			ExampleValue: `"me"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			templateMention.AddFields(b.NewField(e, jen.String(), DiscriminatorValue(e.Property)))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "_Example rich text mention object for a template_mention_user mention _",
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"template_mention\",\n    \"template_mention\": {\n      \"type\": \"template_mention_user\",\n      \"template_mention_user\": \"me\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Me\",\n  \"href\": null\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			b.AddUnmarshalTest("RichText", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "User mention type object",
		}).Output(func(e *Block, b *CodeBuilder) {
			mention.AddAdaptiveFieldWithType("user", e.Text, jen.Op("*").Id("User"))
		})
	}

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "If a rich text object’s type value is \"user\", then the corresponding user field contains a\u00a0user object."})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘If your integration doesn’t yet have access to the mentioned user, then the plain_text that would include a user’s name reads as \"@Anonymous\". To update the integration to get access to the user, update the integration capabilities on the integration settings page."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Example rich text mention object for a user mention"})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"mention\",\n  \"mention\": {\n    \"type\": \"user\",\n    \"user\": {\n      \"object\": \"user\",\n      \"id\": \"b2e19928-b427-4aad-9a9d-fde65479b1d9\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"@Anonymous\",\n  \"href\": null\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		b.AddUnmarshalTest("RichText", e.Text)
	})

	{
		var richTextText *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Text",
		}).Output(func(e *Block, b *CodeBuilder) {
			richTextText = richText.AddAdaptiveFieldWithSpecificObject("text", e.Text, b)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "If a rich text object’s type value is \"text\", then the corresponding text field contains an object including the following:",
		}).Output(func(e *Block, b *CodeBuilder) {
			richTextText.AddComment(e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "content",
			Type:         "string",
			Description:  "The actual text content of the text.",
			ExampleValue: `"Some words "`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			richTextText.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "link",
			Type:         `object (optional)`,
			Description:  "An object with information about any inline link in this text, if included.  \n  \nIf the text contains an inline link, then the object key is url and the value is the URL’s string web address.  \n  \nIf the text doesn’t have any inline links, then the value is null.",
			ExampleValue: "{\n  \"url\": \"https://developers.notion.com/\"\n}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			richTextText.AddFields(b.NewField(e, jen.Op("*").Id("URLReference"))) // RetrivePageでnullを確認
		})
	}

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example rich text text object without link"})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"This is an \",\n    \"link\": null\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"This is an \",\n  \"href\": null\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		b.AddUnmarshalTest("RichText", e.Text)
	})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example rich text text object with link"})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"text\",\n  \"text\": {\n    \"content\": \"inline link\",\n    \"link\": {\n      \"url\": \"https://developers.notion.com/\"\n    }\n  },\n  \"annotations\": {\n    \"bold\": false,\n    \"italic\": false,\n    \"strikethrough\": false,\n    \"underline\": false,\n    \"code\": false,\n    \"color\": \"default\"\n  },\n  \"plain_text\": \"inline link\",\n  \"href\": \"https://developers.notion.com/\"\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		b.AddUnmarshalTest("RichText", e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "📘 Rich text object limitsRefer to the request limits documentation page for information about limits on the size of rich text objects.",
	}).Output(func(e *Block, b *CodeBuilder) {
		richText.AddComment(e.Text)
	})
}
