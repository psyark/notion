package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	var user *adaptiveObject

	registerTranslator(
		"https://developers.notion.com/reference/user",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The User object represents a user in a Notion workspace. Users include full workspace members, and integrations. Guests are not included. You can find more information about members and guests in this guide.",
			}, func(e blockElement) {
				user = b.addAdaptiveObject("User", "type", e.Text)
				for _, f := range user.fields { // TODO なんとかする？
					if f, ok := f.(*field); ok && f.name == "type" {
						f.omitEmpty = true
					}
				}
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "The SCIM API is available for workspaces in Notion's Enterprise Plan. Learn more about using SCIM with Notion.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Single sign-on (SSO) can be configured for workspaces in Notion's Enterprise Plan. Learn more about SSO with Notion.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Where user objects appear in the API */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Where user objects appear in the API",
			}, func(e blockElement) {
				getSymbol[adaptiveObject]("User").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "User objects appear in the API in nearly all objects returned by the API, including:",
			}, func(e blockElement) {
				getSymbol[adaptiveObject]("User").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Block object under created_by and last_edited_by.Page object under created_by and last_edited_by and in people property items.Database object under created_by and last_edited_by.Rich text object, as user mentions.Property object when the property is a people property.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject]("User").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "User objects will always contain object and id keys, as described below. The remaining properties may appear if the user is being rendered in a rich text or page property context, and the bot has the correct capabilities to access those properties. For more about capabilities, see the Capabilities guide and the Authorization guide.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject]("User").addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* All users */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "All users",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "These fields are shared by all users, including people and bots. Fields marked with * are always present.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "object",
				Type:         `"user"`,
				Description:  `Always "user"`,
				ExampleValue: `"user"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("User").addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Property:     "id",
				Type:         "string (UUID)",
				Description:  "Unique identifier for this user.",
				ExampleValue: `"e79a0b74-3aba-4149-9f74-0bb5791a6ee6"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("User").addFields(e.asField(UUID))
			})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         "string (optional, enum)",
				Description:  `Type of the user. Possible values are "person" and "bot".`,
				ExampleValue: `"person"`,
			}, func(e parameterElement) {}) // PersonとBotで定義
			c.nextMustParameter(parameterElement{
				Description:  "User's name, as displayed in Notion.",
				ExampleValue: "\"Avocado Lovelace\"",
				Property:     "name",
				Type:         "string (optional)",
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("User").addFields(e.asField(jen.String(), discriminatorNotEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "avatar_url",
				Type:         "string (optional)",
				Description:  "Chosen avatar image.",
				ExampleValue: `"https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("User").addFields(e.asField(jen.Op("*").String(), discriminatorNotEmpty))
			})
		},
		func(c *comparator, b *builder) /* People */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "People",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "User objects that represent people have the type property set to \"person\". These objects also have the following properties:",
			}, func(e blockElement) {
				user.addAdaptiveFieldWithSpecificObject("person", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:    "person",
				Type:        "object",
				Description: "Properties only present for non-bot users.",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "person.email",
				Type:         "string",
				Description:  "Email address of person. This is only present if an integration has user capabilities that allow access to email addresses.",
				ExampleValue: `"avo@example.org"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject]("UserPerson").addFields(&field{
					name:     "email",
					typeCode: jen.String(),
					comment:  e.Description,
				})
			})
		},
		func(c *comparator, b *builder) /* Bots */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Bots",
			}, func(e blockElement) {
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A user object's type property is\"bot\" when the user object represents a bot. A bot user object has the following properties:",
			}, func(e blockElement) {
				user.addAdaptiveFieldWithSpecificObject("bot", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Description:  "If you're using GET /v1/users/me or GET /v1/users/{{your_bot_id}}, then this field returns data about the bot, including owner, owner.type, and workspace_name. These properties are detailed below.",
				ExampleValue: "{\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\",\n    \"name\": \"Test Integration\",\n    \"avatar_url\": null,\n    \"type\": \"bot\",\n    \"bot\": {\n        \"owner\": {\n        \"type\": \"workspace\",\n        \"workspace\": true\n        },\n \"workspace_name\": \"Ada Lovelace’s Notion\"\n    }\n}",
				Property:     "bot",
				Type:         "object",
			}, func(e parameterElement) {
				b.addUnmarshalTest("User", e.ExampleValue)
			})
			c.nextMustParameter(parameterElement{
				Property:     "owner",
				Type:         "object",
				Description:  "Information about who owns this bot.",
				ExampleValue: "{\n    \"type\": \"workspace\",\n    \"workspace\": true\n}",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("UserBot").addFields(e.asField(jen.Op("*").Id("BotUserDataOwner"), omitEmpty))
				b.addConcreteObject("BotUserDataOwner", e.Description)
				b.addUnmarshalTest("BotUserDataOwner", e.ExampleValue)
			})
			c.nextMustParameter(parameterElement{
				Property:     "owner.type",
				Type:         "string enum",
				Description:  `The type of owner, either "workspace" or "user".`,
				ExampleValue: `"workspace"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject]("BotUserDataOwner").addFields(
					&field{name: "type", typeCode: jen.String(), comment: e.Description},
					&field{name: "workspace", typeCode: jen.Bool(), comment: "undocumented", omitEmpty: true},
					&field{name: "user", typeCode: jen.Bool(), comment: "undocumented", omitEmpty: true},
				)
			})
			c.nextMustParameter(parameterElement{
				Property:     "workspace_name",
				Type:         "string enum",
				Description:  `If the owner.type is "workspace", then workspace.name identifies the name of the workspace that owns the bot. If the owner.type is "user", then workspace.name is null.`,
				ExampleValue: `"Ada Lovelace’s Notion"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject]("UserBot").addFields(e.asField(jen.String(), omitEmpty))
			})
		},
	)
}
