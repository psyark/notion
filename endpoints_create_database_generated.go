// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/create-a-database

package notion

import (
	"context"
	"net/http"
)

/*
Creates a database as a subpage in the specified parent page, with the specified properties schema. Currently, the parent of a new database must be a Notion page or a [wiki database](https://www.notion.so/help/wikis-and-verified-pages).

> 📘 Integration capabilities
>
> This endpoint requires an integration to have insert content capabilities. Attempting to call this API without insert content capabilities will return an HTTP response with a 403 status code. For more information on integration capabilities, see the [capabilities guide](ref:capabilities).

> 🚧 Limitations
>
> - Only empty `status` database properties can be created

### Errors

Returns a 404 if the specified parent page does not exist, or if the integration does not have access to the parent page.

Returns a 400 if the request is incorrectly formatted, or a 429 HTTP response if the request exceeds the [request limits](ref:request-limits).

_Note: Each Public API endpoint can return several possible error codes. See the [Error codes section](https://developers.notion.com/reference/status-codes#error-codes) of the Status codes documentation for more information._
*/
func (c *Client) CreateDatabase(ctx context.Context, params CreateDatabaseParams, options ...callOption) (*Database, error) {
	return call(
		ctx,
		c.accessToken,
		http.MethodPost,
		"/v1/databases",
		params,
		func(u *Database) *Database {
			return u
		},
		options...,
	)
}

type CreateDatabaseParams map[string]any

// A [page parent](/reference/database#page-parent)
func (p CreateDatabaseParams) Parent(parent Parent) CreateDatabaseParams {
	p["parent"] = parent
	return p
}

// Title of database as it appears in Notion. An array of [rich text objects](ref:rich-text).
func (p CreateDatabaseParams) Title(title []RichText) CreateDatabaseParams {
	p["title"] = title
	return p
}

// Property schema of database. The keys are the names of properties as they appear in Notion and the values are [property schema objects](https://developers.notion.com/reference/property-schema-object).
func (p CreateDatabaseParams) Properties(properties map[string]PropertySchema) CreateDatabaseParams {
	p["properties"] = properties
	return p
}
