package notion

import (
	"encoding/json"
	nullv4 "gopkg.in/guregu/null.v4"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/rich-text

/*
Rich text objects contain the data that Notion uses to display formatted text, mentions, and inline equations. Arrays of rich text objects within database property objects and page property value objects are used to create what a user experiences as a single text value in Notion.

Refer to the request limits documentation page for information about limits on the size of rich text objects.
*/
type RichText struct {
	Type        string            `json:"type"`
	Annotations Annotations       `json:"annotations"` // The information used to style the rich text object. Refer to the annotation object section below for details.
	PlainText   string            `json:"plain_text"`  // The plain text without annotations.
	Href        nullv4.String     `json:"href"`        // The URL of any link or Notion mention in this text, if any.
	Equation    *RichTextEquation `json:"equation"`    // Equation
	Mention     *Mention          `json:"mention"`     // Mention
	Text        *RichTextText     `json:"text"`        // Text
}

func (o RichText) MarshalJSON() ([]byte, error) {
	t := o.Type
	type Alias RichText
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"equation": t == "equation",
		"mention":  t == "mention",
		"text":     t == "text",
	}
	return omitFields(data, visibility)
}

/*
The annotation object

All rich text objects contain an annotations object that sets the styling for the rich text. annotations includes the following fields:
*/
type Annotations struct {
	Bold          bool   `json:"bold"`            // Whether the text is bolded.
	Italic        bool   `json:"italic"`          // Whether the text is italicized.
	Strikethrough bool   `json:"strikethrough"`   // Whether the text is struck through.
	Underline     bool   `json:"underline"`       // Whether the text is underlined.
	Code          bool   `json:"code"`            // Whether the text is code style.
	Color         string `json:"color,omitempty"` // Color of the text. Possible values include:   - "blue" - "blue_background" - "brown" - "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" -"orange_background" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background” - "yellow" - "yellow_background"
}

/*
Equation

Notion supports inline LaTeX equations as rich text object’s with a type value of "equation". The corresponding equation type object contains the following:
*/
type RichTextEquation struct {
	Expression string `json:"expression"` // The LaTeX string representing the inline equation.
}

/*
Mention

Mention objects represent an inline mention of a database, date, link preview mention, page, template mention, or user. A mention is created in the Notion UI when a user types @ followed by the name of the reference.
*/
type Mention struct {
	Type            string              `json:"type"`
	Database        *PageReference      `json:"database"`         // Database mention type object
	Date            *PropertyValueDate  `json:"date"`             // Date mention type object
	LinkPreview     *MentionLinkPreview `json:"link_preview"`     // Link Preview mention type object
	Page            *PageReference      `json:"page"`             // Page mention type object
	TemplateMention *TemplateMention    `json:"template_mention"` // Template mention type object
	User            *User               `json:"user"`             // User mention type object
}

func (o Mention) MarshalJSON() ([]byte, error) {
	t := o.Type
	type Alias Mention
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"database":         t == "database",
		"date":             t == "date",
		"link_preview":     t == "link_preview",
		"page":             t == "page",
		"template_mention": t == "template_mention",
		"user":             t == "user",
	}
	return omitFields(data, visibility)
}

/*
Link Preview mention type object

If a user opts to share a Link Preview as a mention, then the API handles the Link Preview mention as a rich text object with a type value of link_preview. Link preview rich text mentions contain a corresponding link_preview object that includes the url that is used to create the Link Preview mention.
*/
type MentionLinkPreview struct {
	Url string `json:"url"`
}

/*
Template mention type object

The content inside a template button in the Notion UI can include placeholder date and user mentions that populate when a template is duplicated. Template mention type objects contain these populated values.

Template mention rich text objects contain a template_mention object with a nested type key that is either "template_mention_date" or "template_mention_user".
*/
type TemplateMention struct {
	Type                string `json:"type"`
	TemplateMentionDate string `json:"template_mention_date"` // The type of the date mention. Possible values include: "today" and "now".
	TemplateMentionUser string `json:"template_mention_user"` // The type of the user mention. The only possible value is "me".
}

func (o TemplateMention) MarshalJSON() ([]byte, error) {
	t := o.Type
	type Alias TemplateMention
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"template_mention_date": t == "template_mention_date",
		"template_mention_user": t == "template_mention_user",
	}
	return omitFields(data, visibility)
}

/*
Text

If a rich text object’s type value is "text", then the corresponding text field contains an object including the following:
*/
type RichTextText struct {
	Content string        `json:"content"` // The actual text content of the text.
	Link    *URLReference `json:"link"`    // An object with information about any inline link in this text, if included.   If the text contains an inline link, then the object key is url and the value is the URL’s string web address.   If the text doesn’t have any inline links, then the value is null.
}
