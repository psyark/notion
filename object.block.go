package notion

import (
	"encoding/json"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/block

/*
A block object represents a piece of content within Notion. The API translates the headings, toggles, paragraphs, lists, media, and more that you can interact with in the Notion UI as different block type objects.

 For example, the following block object represents a Heading 2 in the Notion UI:

Use the Retrieve block children endpoint to list all of the blocks on a page.

Block types that support child blocks
Some block types contain nested blocks. The following block types support child blocks:

- Bulleted list item
- Callout
- Child database
- Child page
- Column
- Heading 1, when the is_toggleable property is true
- Heading 2, when the is_toggleable property is true
- Heading 3, when the is_toggleable property is true
- Numbered list item
- Paragraph
- Quote
- Synced block
- Table
- Template
- To do
- Toggle

The API does not support all block types.
Only the block type objects listed in the reference below are supported. Any unsupported block types appear in the structure, but contain a `type` set to `"unsupported"`.
*/
type Block struct {
	Type             string                 `json:"type"`
	Object           alwaysBlock            `json:"object"`             // Always "block".
	Id               uuid.UUID              `json:"id"`                 // Identifier for the block.
	Parent           Parent                 `json:"parent"`             // Information about the block's parent. See Parent object.
	CreatedTime      ISO8601String          `json:"created_time"`       // Date and time when this block was created. Formatted as an ISO 8601 date time string.
	CreatedBy        User                   `json:"created_by"`         // User who created the block.
	LastEditedTime   ISO8601String          `json:"last_edited_time"`   // Date and time when this block was last updated. Formatted as an ISO 8601 date time string.
	LastEditedBy     User                   `json:"last_edited_by"`     // User who last edited the block.
	Archived         bool                   `json:"archived"`           // The archived status of the block.
	HasChildren      bool                   `json:"has_children"`       // Whether or not the block has children blocks nested within it.
	Bookmark         *BlockBookmark         `json:"bookmark"`           // Bookmark
	Breadcrumb       struct{}               `json:"breadcrumb"`         //  Breadcrumb block objects do not contain any information within the breadcrumb property.
	BulletedListItem *BlockBulletedListItem `json:"bulleted_list_item"` // Bulleted list item
	Callout          *BlockCallout          `json:"callout"`            // Callout
	ChildDatabase    *BlockChildDatabase    `json:"child_database"`     // Child database
	ChildPage        *BlockChildPage        `json:"child_page"`         // Child page
	Code             *BlockCode             `json:"code"`               // Code
	ColumnList       struct{}               `json:"column_list"`        //  Column lists are parent blocks for columns. They do not contain any information within the column_list property.
	Column           struct{}               `json:"column"`             // Columns are parent blocks for any block types listed in this reference except for other columns. They do not contain any information within the column property. They can only be appended to column_lists.
	Divider          struct{}               `json:"divider"`            //  Divider block objects do not contain any information within the divider property.
	Embed            *BlockEmbed            `json:"embed"`              // Embed
	Equation         *BlockEquation         `json:"equation"`           // Equation
	File             FileWithCaption        `json:"file"`               // File
	Heading1         *BlockHeading          `json:"heading_1"`
	Heading2         *BlockHeading          `json:"heading_2"`
	Heading3         *BlockHeading          `json:"heading_3"`
	Image            File                   `json:"image"`        //  Image block objects contain a file object detailing information about the image.
	LinkPreview      *BlockLinkPreview      `json:"link_preview"` //  Link Preview block objects contain the originally pasted url:
	Paragraph        *BlockParagraph        `json:"paragraph"`    // Paragraph
}

func (o Block) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		// TODO
	}
	type Alias Block
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"bookmark":           o.Type == "bookmark",
		"breadcrumb":         o.Type == "breadcrumb",
		"bulleted_list_item": o.Type == "bulleted_list_item",
		"callout":            o.Type == "callout",
		"child_database":     o.Type == "child_database",
		"child_page":         o.Type == "child_page",
		"code":               o.Type == "code",
		"column":             o.Type == "column",
		"column_list":        o.Type == "column_list",
		"divider":            o.Type == "divider",
		"embed":              o.Type == "embed",
		"equation":           o.Type == "equation",
		"file":               o.Type == "file",
		"heading_1":          o.Type == "heading_1",
		"heading_2":          o.Type == "heading_2",
		"heading_3":          o.Type == "heading_3",
		"image":              o.Type == "image",
		"link_preview":       o.Type == "link_preview",
		"paragraph":          o.Type == "paragraph",
	}
	return omitFields(data, visibility)
}

/*
Bookmark
Bookmark block objects contain the following information within the bookmark property:
*/
type BlockBookmark struct {
	Caption []RichText `json:"caption"` // The caption for the bookmark.
	Url     string     `json:"url"`     // The link for the bookmark.
}

/*
Bulleted list item
Bulleted list item block objects contain the following information within the bulleted_list_item property:
*/
type BlockBulletedListItem struct {
	RichText []RichText `json:"rich_text"` // The rich text in the bulleted_list_item block.
	Color    string     `json:"color"`     // The color of the block. Possible values are:   - "blue" - "blue_background" - "brown" -  "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" - "orange_background" - "yellow" - "green" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background" - "yellow_background"
	Children []Block    `json:"children"`  // The nested child blocks (if any) of the bulleted_list_item block.
}

/*
Callout
Callout block objects contain the following information within the callout property:
*/
type BlockCallout struct {
	RichText []RichText  `json:"rich_text"` // The rich text in the callout block.
	Icon     FileOrEmoji `json:"icon"`      // An emoji or file object that represents the callout's icon. If the callout does not have an icon.
	Color    string      `json:"color"`     // The color of the block. Possible values are:   - "blue" - "blue_background" - "brown" -  "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" - "orange_background" - "yellow" - "green" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background" - "yellow_background"
}

/*

Child database block objects contain the following information within the child_database property:
*/
type BlockChildDatabase struct {
	Title string `json:"title"` // The plain text title of the database.
}

/*

Child page block objects contain the following information within the child_page property:
*/
type BlockChildPage struct {
	Title string `json:"title"` // The plain text title of the page.
}

/*

Code block objects contain the following information within the code property:
*/
type BlockCode struct {
	Caption  []RichText `json:"caption"`   // The rich text in the caption of the code block.
	RichText []RichText `json:"rich_text"` // The rich text in the code block.
	Language string     `json:"language"`  // The language of the code contained in the code block.
}

/*

Embed block objects include information about another website displayed within the Notion UI. The embed property contains the following information:
*/
type BlockEmbed struct {
	Url string `json:"url"` // The link to the website that the embed block displays.
}

/*

Equation block objects are represented as children of paragraph blocks. They are nested within a rich text object and contain the following information within the equation property:
*/
type BlockEquation struct {
	Expression string `json:"expression"` // A KaTeX compatible string.
}

/*

All heading block objects, heading_1, heading_2, and heading_3, contain the following information within their corresponding objects:
*/
type BlockHeading struct {
	RichText     []RichText `json:"rich_text"`     // The rich text of the heading.
	Color        string     `json:"color"`         // The color of the block. Possible values are:   - "blue" - "blue_background" - "brown" -  "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" - "orange_background" - "yellow" - "green" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background" - "yellow_background"
	IsToggleable bool       `json:"is_toggleable"` // Whether or not the heading block is a toggle heading or not. If true, then the heading block toggles and can support children. If false, then the heading block is a static heading block.
}

/*

Link Preview block objects contain the originally pasted url:
*/
type BlockLinkPreview struct {
	Url string `json:"url"`
}

/*

Paragraph block objects contain the following information within the paragraph property:
*/
type BlockParagraph struct {
	RichText []RichText `json:"rich_text"`          // The rich text displayed in the paragraph block.
	Color    string     `json:"color"`              // The color of the block. Possible values are:   - "blue" - "blue_background" - "brown" -  "brown_background" - "default" - "gray" - "gray_background" - "green" - "green_background" - "orange" - "orange_background" - "yellow" - "green" - "pink" - "pink_background" - "purple" - "purple_background" - "red" - "red_background" - "yellow_background"
	Children []Block    `json:"children,omitempty"` // The nested child blocks (if any) of the paragraph block.
}
