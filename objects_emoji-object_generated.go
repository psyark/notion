// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/emoji-object

package notion

// An emoji object contains information about an emoji character. It is most often used to represent an emoji that is rendered as a page icon in the Notion UI.
type Emoji struct {
	Type  alwaysEmoji `json:"type"`  // The constant string "emoji" that represents the object type.
	Emoji string      `json:"emoji"` // The emoji character.
}

func (Emoji) isFileOrEmoji() {}
