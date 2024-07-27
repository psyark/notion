// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/post-page

package notion

import (
	"context"
	"net/http"
)

/*
Creates a new page that is a child of an existing page or database.

If the new page is a child of an existing page,`title` is the only valid property in the `properties` body param.

If the new page is a child of an existing database, the keys of the `properties` object body param must match the parent [database's properties](https://developers.notion.com/reference/property-object).

This endpoint can be used to create a new page with or without content using the `children` option. To add content to a page after creating it, use the [Append block children](https://developers.notion.com/reference/patch-block-children) endpoint.

Returns a new [page object](https://developers.notion.com/reference/page).

> 🚧 Some page `properties` are not supported via the API.
>
> A request body that includes `rollup`, `created_by`, `created_time`, `last_edited_by`, or `last_edited_time` values in the properties object returns an error. These Notion-generated values cannot be created or updated via the API. If the `parent` contains any of these properties, then the new page’s corresponding values are automatically created.

> 📘 Requirements
>
> Your integration must have [Insert Content capabilities](https://developers.notion.com/reference/capabilities#content-capabilities) on the target parent page or database in order to call this endpoint. To update your integrations capabilities, navigation to the [My integrations](https://www.notion.so/my-integrations) dashboard, select your integration, go to the **Capabilities** tab, and update your settings as needed.
>
> Attempting a query without update content capabilities returns an HTTP response with a 403 status code.

### Errors

Each Public API endpoint can return several possible error codes. See the [Error codes section](https://developers.notion.com/reference/status-codes#error-codes) of the Status codes documentation for more information.
*/
func (c *Client) CreatePage(ctx context.Context, params CreatePageParams, options ...CallOption) (*Page, error) {
	return call(
		ctx,
		c.accessToken,
		http.MethodPost,
		"/v1/pages",
		params,
		accessValue[*Page],
		options...,
	)
}

type CreatePageParams map[string]any

// The parent page or database where the new page is inserted, represented as a JSON object with a `page_id` or `database_id` key, and the corresponding ID.
func (p CreatePageParams) Parent(parent Parent) CreatePageParams {
	p["parent"] = parent
	return p
}

// The values of the page’s properties. If the `parent` is a database, then the schema must match the parent database’s properties. If the `parent` is a page, then the only valid object key is `title`.
func (p CreatePageParams) Properties(properties PropertyValueMap) CreatePageParams {
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
