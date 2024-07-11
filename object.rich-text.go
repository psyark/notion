// Code generated by notion.doc2api; DO NOT EDIT.

package notion

import "encoding/json"

// https://developers.notion.com/reference/rich-text

/*
Notion uses rich text to allow users to customize their content. Rich text refers to a type of document where content can be styled and formatted in a variety of customizable ways. This includes styling decisions, such as the use of italics, font size, and font color, as well as formatting, such as the use of hyperlinks or code blocks.

Notion includes rich text objects in block objects to indicate how blocks in a page are represented. Blocks that support rich text will include a rich text object; however, not all block types offer rich text.

When blocks are retrieved from a page using the Retrieve a block or Retrieve block children endpoints, an array of rich text objects will be included in the block object (when available). Developers can use this array to retrieve the plain text (plain_text) for the block or get all the rich text styling and formatting options applied to the block.

📘 Rich text object limitsRefer to the request limits documentation page for information about limits on the size of rich text objects.
*/
type RichText struct {
	Type        string            `json:"type"`
	Annotations Annotations       `json:"annotations"` // The information used to style the rich text object. Refer to the annotation object section below for details.
	PlainText   string            `json:"plain_text"`  // The plain text without annotations.
	Href        *string           `json:"href"`        // The URL of any link or Notion mention in this text, if any.
	Equation    *RichTextEquation `json:"equation"`    // Equation
	Mention     *Mention          `json:"mention"`     // Mention
	Text        *RichTextText     `json:"text"`        // Text
}

func (o RichText) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.Equation):
			o.Type = "equation"
		case defined(o.Mention):
			o.Type = "mention"
		case defined(o.Text):
			o.Type = "text"
		}
	}
	type Alias RichText
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"equation": o.Type == "equation",
		"mention":  o.Type == "mention",
		"text":     o.Type == "text",
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
	Color         string `json:"color,omitempty"` // Color of the text. Possible values include: - "blue" - "blue_background" - "brown" - "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" -"orange_background" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background” - "yellow" - "yellow_background"
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
	if o.Type == "" {
		switch {
		case defined(o.Database):
			o.Type = "database"
		case defined(o.Date):
			o.Type = "date"
		case defined(o.LinkPreview):
			o.Type = "link_preview"
		case defined(o.Page):
			o.Type = "page"
		case defined(o.TemplateMention):
			o.Type = "template_mention"
		case defined(o.User):
			o.Type = "user"
		}
	}
	type Alias Mention
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"database":         o.Type == "database",
		"date":             o.Type == "date",
		"link_preview":     o.Type == "link_preview",
		"page":             o.Type == "page",
		"template_mention": o.Type == "template_mention",
		"user":             o.Type == "user",
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
	if o.Type == "" {
		switch {
		case defined(o.TemplateMentionDate):
			o.Type = "template_mention_date"
		case defined(o.TemplateMentionUser):
			o.Type = "template_mention_user"
		}
	}
	type Alias TemplateMention
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"template_mention_date": o.Type == "template_mention_date",
		"template_mention_user": o.Type == "template_mention_user",
	}
	return omitFields(data, visibility)
}

/*
Text

If a rich text object’s type value is "text", then the corresponding text field contains an object including the following:
*/
type RichTextText struct {
	Content string        `json:"content"` // The actual text content of the text.
	Link    *URLReference `json:"link"`    // An object with information about any inline link in this text, if included. If the text contains an inline link, then the object key is url and the value is the URL’s string web address. If the text doesn’t have any inline links, then the value is null.
}
