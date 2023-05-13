package notion

import (
	"context"
	"net/http"
)

// Create a page
// https://developers.notion.com/reference/post-page
func (c *Client) CreatePage(ctx context.Context, params CreatePageParams, options ...callOption) (*Page, error) {
	result := &Page{}
	co := &callOptions{
		method: http.MethodPost,
		params: params,
		path:   "/v1/pages",
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}

type CreatePageParams map[string]any

// The parent page or database where the new page is inserted, represented as a JSON object with a `page_id` or `database_id` key, and the corresponding ID.
func (p CreatePageParams) Parent(parent Parent) CreatePageParams {
	p["parent"] = parent
	return p
}

// The values of the page’s properties. If the `parent` is a database, then the schema must match the parent database’s properties. If the `parent` is a page, then the only valid object key is `title`.
func (p CreatePageParams) Properties(properties map[string]PropertyValue) CreatePageParams {
	p["properties"] = properties
	return p
}

// The content to be rendered on the new page, represented as an array of [block objects](https://developers.notion.com/reference/block).
func (p CreatePageParams) Children(children []Block) CreatePageParams {
	p["children"] = children
	return p
}

// The icon of the new page. Either an [emoji object](https://developers.notion.com/reference/emoji-object) or an [external file object](https://developers.notion.com/reference/file-object)..
func (p CreatePageParams) Icon(icon FileOrEmoji) CreatePageParams {
	p["icon"] = icon
	return p
}

// The cover image of the new page, represented as a [file object](https://developers.notion.com/reference/file-object).
func (p CreatePageParams) Cover(cover File) CreatePageParams {
	p["cover"] = cover
	return p
}
