package notion

import (
	"encoding/json"
	uuid "github.com/google/uuid"
	nullv4 "gopkg.in/guregu/null.v4"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/user

/*
The User object represents a user in a Notion workspace. Users include full workspace members, and integrations. Guests are not included. You can find more information about members and guests in this guide.

Where user objects appear in the API

User objects appear in the API in nearly all objects returned by the API, including:

Block object under created_by and last_edited_by.Page object under created_by and last_edited_by and in people property items.Database object under created_by and last_edited_by.Rich text object, as user mentions.Property object when the property is a people property.

User objects will always contain object and id keys, as described below. The remaining properties may appear if the user is being rendered in a rich text or page property context, and the bot has the correct capabilities to access those properties. For more about capabilities, see the Capabilities guide and the Authorization guide.
*/
type User struct {
	Type      string        `json:"type,omitempty"`
	Object    alwaysUser    `json:"object"`     // Always "user"
	Id        uuid.UUID     `json:"id"`         // Unique identifier for this user.
	Name      string        `json:"name"`       // User's name, as displayed in Notion.
	AvatarUrl nullv4.String `json:"avatar_url"` // Chosen avatar image.
	Person    *UserPerson   `json:"person"`     // User objects that represent people have the type property set to "person". These objects also have the following properties:
	Bot       *UserBot      `json:"bot"`        // A user object's type property is"bot" when the user object represents a bot. A bot user object has the following properties:
}

func (o User) MarshalJSON() ([]byte, error) {
	t := o.Type
	type Alias User
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"avatar_url": t != "",
		"bot":        t == "bot",
		"name":       t != "",
		"person":     t == "person",
	}
	return omitFields(data, visibility)
}

// User objects that represent people have the type property set to "person". These objects also have the following properties:
type UserPerson struct {
	Email string `json:"email"` // Email address of person. This is only present if an integration has user capabilities that allow access to email addresses.
}

// A user object's type property is"bot" when the user object represents a bot. A bot user object has the following properties:
type UserBot struct {
	Owner         *BotUserDataOwner `json:"owner,omitempty"`          // Information about who owns this bot.
	WorkspaceName string            `json:"workspace_name,omitempty"` // If the owner.type is "workspace", then workspace.name identifies the name of the workspace that owns the bot. If the owner.type is "user", then workspace.name is null.
}

// Information about who owns this bot.
type BotUserDataOwner struct {
	Type      string `json:"type"`                // The type of owner, either "workspace" or "user".
	Workspace bool   `json:"workspace,omitempty"` // undocumented
	User      bool   `json:"user,omitempty"`      // undocumented
}
