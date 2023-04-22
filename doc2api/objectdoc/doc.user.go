package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/user",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "The User object represents a user in a Notion workspace. Users include full workspace members, and integrations. Guests are not included. You can find more information about members and guests in this guide. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addAbstractObject("DetailedUser", "type", e.Text)
					b.addAbstractObject("User", "type", e.Text)
					b.addAbstractList("User", "Users")
					b.getAbstractObject("User").addVariant(
						b.addSpecificObject("PartialUser", e.Text),
					)
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:   "The SCIM API is available for workspaces in Notion's Enterprise Plan. Learn more about [using SCIM with Notion](https://www.notion.so/help/provision-users-and-groups-with-scim).",
				Title:  "Provisioning users and groups using SCIM",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocCalloutElement{
				Body:   "Single sign-on (SSO) can be configured for workspaces in Notion's Enterprise Plan. [Learn more about SSO with Notion](https://www.notion.so/help/saml-sso-configuration).",
				Title:  "Setting up single sign-on (SSO) with Notion",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text: "Where user objects appear in the API",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("User").comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nUser objects appear in the API in nearly all objects returned by the API, including:\n* Block object under created_by and last_edited_by.\n* Page object under created_by and last_edited_by and in people property items.\n* Database object under created_by and last_edited_by.\n* Rich text object, as user mentions.\n* Property object when the property is a people property.\n\nUser objects will always contain object and id keys, as described below. The remaining properties may appear if the user is being rendered in a rich text or page property context, and the bot has the correct capabilities to access those properties. For more about capabilities, see the Capabilities guide and the Authorization guide.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("User").comment += "\n\n" + e.Text
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "All users",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("User").fieldsComment += "\n" + e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nThese fields are shared by all users, including people and bots. Fields marked with * are always present.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("User").fieldsComment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "object*",
				Type:         `"user"`,
				Description:  `Always "user"`,
				ExampleValue: `"user"`,
				output: func(e *objectDocParameter, b *builder) error {
					e.Property = strings.TrimSuffix(e.Property, "*")
					b.getAbstractObject("User").addFields(e.asFixedStringField())
					return nil
				},
			}, {
				Property:     "id*",
				Type:         "string (UUID)",
				Description:  "Unique identifier for this user.",
				ExampleValue: `"e79a0b74-3aba-4149-9f74-0bb5791a6ee6"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("User").addFields(&field{
						name:     strings.TrimSuffix(e.Property, "*"),
						typeCode: UUID,
						comment:  e.Description,
					})
					return nil
				},
			}, {
				Property:     "type",
				Type:         "string (optional, enum)",
				Description:  `Type of the user. Possible values are "person" and "bot".`,
				ExampleValue: `"person"`,
				output: func(e *objectDocParameter, b *builder) error {
					return nil // PersonとBotで定義
				},
			}, {
				Property:     "name",
				Type:         "string (optional)",
				Description:  "User's name, as displayed in Notion.",
				ExampleValue: `"Avocado Lovelace"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("DetailedUser").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "avatar_url",
				Type:         "string (optional)",
				Description:  "Chosen avatar image.",
				ExampleValue: `"https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("DetailedUser").addFields(e.asField(NullString))
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "People",
				output: func(e *objectDocHeadingElement, b *builder) error {
					so := b.addSpecificObject("PersonUser", e.Text).addFields(
						&fixedStringField{name: "type", value: "person"},
					)
					so.addParent(b.getAbstractObject("DetailedUser"))
					b.getAbstractObject("User").addVariant(so)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nUser objects that represent people have the type property set to \"person\". These objects also have the following properties:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PersonUser").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:    "person",
				Type:        "object",
				Description: "Properties only present for non-bot users.",
				output: func(e *objectDocParameter, b *builder) error {
					b.addSpecificObject("PersonData", e.Description)
					b.getSpecificObject("PersonUser").addFields(e.asField(jen.Id("PersonData")))
					return nil
				},
			}, {
				Property:     "person.email",
				Type:         "string",
				Description:  "Email address of person. This is only present if an integration has user capabilities that allow access to email addresses.",
				ExampleValue: `"avo@example.org"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PersonData").addFields(&field{
						name:     "email",
						typeCode: jen.String(),
						comment:  e.Description,
					})
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "Bots",
				output: func(e *objectDocHeadingElement, b *builder) error {
					so := b.addSpecificObject("BotUser", e.Text).addFields(
						&fixedStringField{name: "type", value: "bot"},
					)
					so.addParent(b.getAbstractObject("DetailedUser"))
					b.getAbstractObject("User").addVariant(so)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA user object's type property is\"bot\" when the user object represents a bot. A bot user object has the following properties:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("BotUser").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "bot",
				Type:         "object",
				Description:  "If you're using GET /v1/users/me or GET /v1/users/{{your_bot_id}}, then this field returns data about the bot, including owner, owner.type, and workspace_name. These properties are detailed below.",
				ExampleValue: "{\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\",\n    \"name\": \"Test Integration\",\n    \"avatar_url\": null,\n    \"type\": \"bot\",\n    \"bot\": {\n        \"owner\": {\n        \"type\": \"workspace\",\n        \"workspace\": true\n        },\n \"workspace_name\": \"Ada Lovelace’s Notion\"\n    }\n}",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("BotUser").addFields(e.asField(jen.Id("BotData")))
					b.addSpecificObject("BotData", e.Description)
					b.addUnmarshalTest("User", e.ExampleValue)
					return nil
				},
			}, {
				Property:     "owner",
				Type:         "object",
				Description:  "Information about who owns this bot.",
				ExampleValue: "{\n    \"type\": \"workspace\",\n    \"workspace\": true\n}",
				output: func(e *objectDocParameter, b *builder) error {
					field := e.asField(jen.Op("*").Id("BotDataOwner"))
					field.omitEmpty = true
					b.getSpecificObject("BotData").addFields(field)
					b.addSpecificObject("BotDataOwner", e.Description)
					b.addUnmarshalTest("BotDataOwner", e.ExampleValue)
					return nil
				},
			}, {
				Property:     "owner.type",
				Type:         "string enum",
				Description:  `The type of owner, either "workspace" or "user".`,
				ExampleValue: `"workspace"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("BotDataOwner").addFields(
						&field{name: "type", typeCode: jen.String(), comment: e.Description},
						&field{name: "workspace", typeCode: jen.Bool(), comment: "undocumented", omitEmpty: true},
						&field{name: "user", typeCode: jen.Bool(), comment: "undocumented", omitEmpty: true},
					)
					return nil
				},
			}, {
				Property:     "workspace_name",
				Type:         "string enum",
				Description:  `If the owner.type is "workspace", then workspace.name identifies the name of the workspace that owns the bot. If the owner.type is "user", then workspace.name is null.`,
				ExampleValue: `"Ada Lovelace’s Notion"`,
				output: func(e *objectDocParameter, b *builder) error {
					field := e.asField(jen.String())
					field.omitEmpty = true
					b.getSpecificObject("BotData").addFields(field)
					return nil
				},
			}},
		},
	})
}
