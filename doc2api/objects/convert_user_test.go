package objects_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestUser(t *testing.T) {
	t.Parallel()

	converter.FetchDocument("https://developers.notion.com/reference/user").WithScope(func(c *DocumentComparator) {
		var user *AdaptiveObject
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The User object represents a user in a Notion workspace. Users include full workspace members, guests, and integrations. You can find more information about members and guests in this guide.",
		}).Output(func(e *Block, b *CodeBuilder) {
			user = b.AddAdaptiveObject("User", "type", e.Text, DiscriminatorOmitEmpty())
		})
		c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘 Provisioning users and groups using SCIMThe SCIM API is available for workspaces in Notion's Enterprise Plan. Learn more about using SCIM with Notion."})
		c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘 Setting up single sign-on (SSO) with NotionSingle sign-on (SSO) can be configured for workspaces in Notion's Enterprise Plan. Learn more about SSO with Notion."})

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Where user objects appear in the API",
		}).Output(func(e *Block, b *CodeBuilder) {
			user.AddComment(e.Text)
		})

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "User objects appear in the API in nearly all objects returned by the API, including:",
		}).Output(func(e *Block, b *CodeBuilder) {
			user.AddComment(e.Text)
		})

		c.ExpectBlock(&Block{
			Kind: "List",
			Text: "Block object under created_by and last_edited_by.Page object under created_by and last_edited_by and in people property items.Database object under created_by and last_edited_by.Rich text object, as user mentions.Property object when the property is a people property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			user.AddComment(e.Text)
		})

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "User objects will always contain object and id keys, as described below. The remaining properties may appear if the user is being rendered in a rich text or page property context, and the bot has the correct capabilities to access those properties. For more about capabilities, see the Capabilities guide and the Authorization guide.",
		}).Output(func(e *Block, b *CodeBuilder) {
			user.AddComment(e.Text)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "All users"})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "These fields are shared by all users, including people and bots. Fields marked with \\* are always present."})

		c.ExpectParameter(&Parameter{
			Property:     `object\*`,
			Type:         `"user"`,
			Description:  `Always "user"`,
			ExampleValue: `"user"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			e.Property = strings.TrimSuffix(e.Property, `\*`)
			user.AddFields(b.NewFixedStringField(e))
		})
		c.ExpectParameter(&Parameter{
			Property:     `id\*`,
			Type:         "string (UUID)",
			Description:  "Unique identifier for this user.",
			ExampleValue: `"e79a0b74-3aba-4149-9f74-0bb5791a6ee6"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			e.Property = strings.TrimSuffix(e.Property, `\*`)
			user.AddFields(b.NewField(e, UUID))
		})
		c.ExpectParameter(&Parameter{
			Property:     "type",
			Type:         "string (optional, enum)",
			Description:  `Type of the user. Possible values are "person" and "bot".`,
			ExampleValue: `"person"`,
		}) // PersonとBotで定義

		c.ExpectParameter(&Parameter{
			Description:  "User's name, as displayed in Notion.",
			ExampleValue: `"Avocado Lovelace"`,
			Property:     "name",
			Type:         "string (optional)",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			user.AddFields(b.NewField(e, jen.String(), DiscriminatorNotEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "avatar_url",
			Type:         "string (optional)",
			Description:  "Chosen avatar image.",
			ExampleValue: `"https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			user.AddFields(b.NewField(e, jen.Op("*").String(), DiscriminatorNotEmpty))
		})

		var person *ConcreteObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "People",
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: `User objects that represent people have the type property set to "person". These objects also have the following properties:`,
		}).Output(func(e *Block, b *CodeBuilder) {
			person = user.AddAdaptiveFieldWithSpecificObject("person", e.Text, b)
		})
		c.ExpectParameter(&Parameter{
			Property:    "person",
			Type:        "object",
			Description: "Properties only present for non-bot users.",
		}).Output(func(e *Parameter, b *CodeBuilder) {})
		c.ExpectParameter(&Parameter{
			Property:     "person.email",
			Type:         "string",
			Description:  "Email address of person. This is only present if an integration has user capabilities that allow access to email addresses.",
			ExampleValue: `"avo@example.org"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			e.Property = "email"
			person.AddFields(b.NewField(e, jen.String()))
		})

		var bot *ConcreteObject
		var botOwner *ConcreteObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Bots",
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A user object's type property is\"bot\" when the user object represents a bot. A bot user object has the following properties:",
		}).Output(func(e *Block, b *CodeBuilder) {
			bot = user.AddAdaptiveFieldWithSpecificObject("bot", e.Text, b)
		})
		c.ExpectParameter(&Parameter{
			Description:  "If you're using GET /v1/users/me or GET /v1/users/{{your_bot_id}}, then this field returns data about the bot, including owner, owner.type, and workspace_name. These properties are detailed below.",
			ExampleValue: "{     \"object\": \"user\",     \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\",     \"name\": \"Test Integration\",     \"avatar_url\": null,     \"type\": \"bot\",     \"bot\": {         \"owner\": {         \"type\": \"workspace\",         \"workspace\": true         },  \"workspace_name\": \"Ada Lovelace’s Notion\"     } }",
			Property:     "bot",
			Type:         "object",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			b.AddUnmarshalTest("User", e.ExampleValue)
		})
		c.ExpectParameter(&Parameter{
			Property:     "owner",
			Type:         "object",
			Description:  "Information about who owns this bot.",
			ExampleValue: "{     \"type\": \"workspace\",     \"workspace\": true }",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			bot.AddFields(b.NewField(e, jen.Op("*").Id("BotUserDataOwner"), OmitEmpty))
			botOwner = b.AddConcreteObject("BotUserDataOwner", e.Description)
			b.AddUnmarshalTest("BotUserDataOwner", e.ExampleValue)
		})
		c.ExpectParameter(&Parameter{
			Property:     "owner.type",
			Type:         "string enum",
			Description:  `The type of owner, either "workspace" or "user".`,
			ExampleValue: `"workspace"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			e.Property = "type"
			botOwner.AddFields(
				b.NewField(e, jen.String()),
				b.NewField(&Parameter{Property: "workspace", Description: UNDOCUMENTED}, jen.Bool(), OmitEmpty),
				b.NewField(&Parameter{Property: "user", Description: UNDOCUMENTED}, jen.Bool(), OmitEmpty),
			)
		})
		c.ExpectParameter(&Parameter{
			Property:     "workspace_name",
			Type:         "string enum",
			Description:  `If the owner.type is "workspace", then workspace.name identifies the name of the workspace that owns the bot. If the owner.type is "user", then workspace.name is null.`,
			ExampleValue: `"Ada Lovelace’s Notion"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			bot.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
	})
}
