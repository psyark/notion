package notion

import (
	"encoding/json"
	"fmt"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/rich-text

/*
Rich text objects contain the data that Notion uses to display formatted text, mentions, and inline equations. Arrays of rich text objects within database property objects and page property value objects are used to create what a user experiences as a single text value in Notion.

{
  "type": "text",
  "text": {
    "content": "Some words ",
    "link": null
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "Some words ",
  "href": null
}

Rich text object limits
Refer to the request limits documentation page for information about [limits on the size of rich text objects](https://developers.notion.com/reference/request-limits#limits-for-property-values).
*/
type RichText interface {
	isRichText()
}
type richTextCommon struct {
	Annotations Annotations `json:"annotations"` // The information used to style the rich text object. Refer to the annotation object section below for details.
	PlainText   string      `json:"plain_text"`  // The plain text without annotations.
	Href        *string     `json:"href"`        // The URL of any link or Notion mention in this text, if any.
}

type richTextUnmarshaler struct {
	value RichText
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *richTextUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "type")) {
	case "\"equation\"":
		u.value = &EquationRichText{}
	case "\"mention\"":
		u.value = &MentionRichText{}
	case "\"text\"":
		u.value = &TextRichText{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

type RichTextArray []RichText

func (a *RichTextArray) UnmarshalJSON(data []byte) error {
	t := []richTextUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*a = make([]RichText, len(t))
	for i, u := range t {
		(*a)[i] = u.value
	}
	return nil
}

/*
The annotation object

All rich text objects contain an annotations object that sets the styling for the rich text. annotations includes the following fields:
*/
type Annotations struct {
	Bold          bool   `json:"bold"`          // Whether the text is bolded.
	Italic        bool   `json:"italic"`        // Whether the text is italicized.
	Strikethrough bool   `json:"strikethrough"` // Whether the text is struck through.
	Underline     bool   `json:"underline"`     // Whether the text is underlined.
	Code          bool   `json:"code"`          // Whether the text is code style.
	Color         string `json:"color"`         // Color of the text. Possible values include:   - "blue" - "blue_background" - "brown" - "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" -"orange_background" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background” - "yellow" - "yellow_background"
}

// Rich text type objects

/*
Equation
Notion supports inline LaTeX equations as rich text object’s with a type value of "equation". The corresponding equation type object contains the following:

Example rich text equation object
{
  "type": "equation",
  "equation": {
    "expression": "E = mc^2"
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "E = mc^2",
  "href": null
}
*/
type EquationRichText struct {
	richTextCommon
	Type       string `always:"equation" json:"type"`
	Expression string `json:"expression"` // The LaTeX string representing the inline equation.
}

func (_ *EquationRichText) isRichText() {}

/*
Mention
Mention objects represent an inline mention of a database, date, link preview mention, page, template mention, or user. A mention is created in the Notion UI when a user types @ followed by the name of the reference.

If a rich text object’s type value is "mention", then the corresponding mention object contains the following:
*/
type Mention interface {
	isMention()
}

type mentionUnmarshaler struct {
	value Mention
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *mentionUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "type")) {
	case "\"database\"":
		u.value = &DatabaseMention{}
	case "\"date\"":
		u.value = &DateMention{}
	case "\"link_preview\"":
		u.value = &LinkPreviewMention{}
	case "\"page\"":
		u.value = &PageMention{}
	case "\"template_mention\"":
		u.value = &TemplateMention{}
	case "\"user\"":
		u.value = &UserMention{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

// Mention
type MentionRichText struct {
	richTextCommon
	Type    string  `always:"mention" json:"type"`
	Mention Mention `json:"mention"`
}

func (_ *MentionRichText) isRichText() {}

/*
Database mention type object
Database mentions contain a database reference within the corresponding database field. A database reference is an object with an id key and a string value (UUIDv4) corresponding to a database ID.

If an integration doesn’t have access to the mentioned database, then the mention is returned with just the ID. The plain_text value that would be a title appears as "Untitled" and the annotation object’s values are defaults.

Example rich text mention object for a database mention
{
  "type": "mention",
  "mention": {
    "type": "database",
    "database": {
      "id": "a1d8501e-1ac1-43e9-a6bd-ea9fe6c8822b"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "Database with test things",
  "href": "https://www.notion.so/a1d8501e1ac143e9a6bdea9fe6c8822b"
}
*/
type DatabaseMention struct {
	Type     string        `always:"database" json:"type"`
	Database PageReference `json:"database"`
}

func (_ *DatabaseMention) isMention() {}

/*
Date mention type object
Date mentions contain a date property value object within the corresponding date field.

Example rich text mention object for a date mention
{
  "type": "mention",
  "mention": {
    "type": "date",
    "date": {
      "start": "2022-12-16",
      "end": null
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "2022-12-16",
  "href": null
}
*/
type DateMention struct {
	Type string            `always:"date" json:"type"`
	Date DatePropertyValue `json:"date"`
}

func (_ *DateMention) isMention() {}

/*
Link Preview mention type object
If a user opts to share a Link Preview as a mention, then the API handles the Link Preview mention as a rich text object with a type value of link_preview. Link preview rich text mentions contain a corresponding link_preview object that includes the url that is used to create the Link Preview mention.

Example rich text mention object for a link_preview mention
{
  "type": "mention",
  "mention": {
    "type": "link_preview",
    "link_preview": {
      "url": "https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD",
  "href": "https://workspace.slack.com/archives/C04PF0F9QSD/z1671139297838409?thread_ts=1671139274.065079&cid=C03PF0F9QSD"
}
*/
type LinkPreviewMention struct {
	Type        string       `always:"link_preview" json:"type"`
	LinkPreview URLReference `json:"link_preview"`
}

func (_ *LinkPreviewMention) isMention() {}

/*
Page mention type object
Page mentions contain a page reference within the corresponding page field. A page reference is an object with an id property and a string value (UUIDv4) corresponding to a page ID.

If an integration doesn’t have access to the mentioned page, then the mention is returned with just the ID. The plain_text value that would be a title appears as "Untitled" and the annotation object’s values are defaults.

Example rich text mention object for a page mention
{
  "type": "mention",
  "mention": {
    "type": "page",
    "page": {
      "id": "3c612f56-fdd0-4a30-a4d6-bda7d7426309"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "This is a test page",
  "href": "https://www.notion.so/3c612f56fdd04a30a4d6bda7d7426309"
}
*/
type PageMention struct {
	Type string        `always:"page" json:"type"`
	Page PageReference `json:"page"`
}

func (_ *PageMention) isMention() {}

/*
Template mention type object
The content inside a template button in the Notion UI can include placeholder date and user mentions that populate when a template is duplicated. Template mention type objects contain these populated values.

Template mention rich text objects contain a template_mention object with a nested type key that is either "template_mention_date" or "template_mention_user".

If the type key is "template_mention_date", then the rich text object contains the following template_mention_date field:
*/
type TemplateMention struct {
	Type            string              `always:"template_mention" json:"type"`
	TemplateMention TemplateMentionData `json:"template_mention"`
}

func (_ *TemplateMention) isMention() {}

type TemplateMentionData interface {
	isTemplateMentionData()
}

type templateMentionDataUnmarshaler struct {
	value TemplateMentionData
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *templateMentionDataUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "type")) {
	case "\"template_mention_date\"":
		u.value = &TemplateMentionDate{}
	case "\"template_mention_user\"":
		u.value = &TemplateMentionUser{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

/*
Example rich text mention object for a template_mention_date mention
{
  "type": "mention",
  "mention": {
    "type": "template_mention",
    "template_mention": {
      "type": "template_mention_date",
      "template_mention_date": "today"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "@Today",
  "href": null
}
*/
type TemplateMentionDate struct {
	Type                string `always:"template_mention_date" json:"type"`
	TemplateMentionDate string `json:"template_mention_date"` // The type of the date mention. Possible values include: "today" and "now".
}

func (_ *TemplateMentionDate) isTemplateMentionData() {}

/*
Example rich text mention object for a template_mention_user mention
{
  "type": "mention",
  "mention": {
    "type": "template_mention",
    "template_mention": {
      "type": "template_mention_user",
      "template_mention_user": "me"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "@Me",
  "href": null
}
*/
type TemplateMentionUser struct {
	Type                string `always:"template_mention_user" json:"type"`
	TemplateMentionUser string `always:"me" json:"template_mention_user"` // The type of the user mention. The only possible value is "me".
}

func (_ *TemplateMentionUser) isTemplateMentionData() {}

/*
User mention type object
If a rich text object’s type value is "user", then the corresponding user field contains a user object.
If your integration doesn’t yet have access to the mentioned user, then the `plain_text` that would include a user’s name reads as `"@Anonymous"`. To update the integration to get access to the user, update the integration capabilities on the integration settings page.

Example rich text mention object for a user mention
{
  "type": "mention",
  "mention": {
    "type": "user",
    "user": {
      "object": "user",
      "id": "b2e19928-b427-4aad-9a9d-fde65479b1d9"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "@Anonymous",
  "href": null
}
*/
type UserMention struct {
	Type string      `always:"user" json:"type"`
	User PartialUser `json:"user"`
}

func (_ *UserMention) isMention() {}

// Text
type TextRichText struct {
	richTextCommon
	Type string `always:"text" json:"type"`
	Text Text   `json:"text"`
}

func (_ *TextRichText) isRichText() {}

/*

If a rich text object’s type value is "text", then the corresponding text field contains an object including the following:

Example rich text text object without link
{
  "type": "text",
  "text": {
    "content": "This is an ",
    "link": null
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "This is an ",
  "href": null
}

Example rich text text object with link
{
  "type": "text",
  "text": {
    "content": "inline link",
    "link": {
      "url": "https://developers.notion.com/"
    }
  },
  "annotations": {
    "bold": false,
    "italic": false,
    "strikethrough": false,
    "underline": false,
    "code": false,
    "color": "default"
  },
  "plain_text": "inline link",
  "href": "https://developers.notion.com/"
}
*/
type Text struct {
	Content string        `json:"content"` // The actual text content of the text.
	Link    *URLReference `json:"link"`    // An object with information about any inline link in this text, if included.   If the text contains an inline link, then the object key is url and the value is the URL’s string web address.   If the text doesn’t have any inline links, then the value is null.
}
