package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Query a database
// https://developers.notion.com/reference/post-database-query
func (c *Client) QueryDatabase(ctx context.Context, database_id uuid.UUID, params QueryDatabaseParams, options ...callOption) (*Pagination, error) {
	result := &Pagination{}
	co := &callOptions{
		method: http.MethodPost,
		params: params,
		path:   fmt.Sprintf("/v1/databases/%v/query", database_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}

type QueryDatabaseParams map[string]any

// When supplied, limits which pages are returned based on the [filter conditions](ref:post-database-query-filter).
func (p QueryDatabaseParams) Filter(filter Filter) QueryDatabaseParams {
	p["filter"] = filter
	return p
}

// When supplied, orders the results based on the provided [sort criteria](ref:post-database-query-sort).
func (p QueryDatabaseParams) Sorts(sorts Sort) QueryDatabaseParams {
	p["sorts"] = sorts
	return p
}

// When supplied, returns a page of results starting after the cursor provided. If not supplied, this endpoint will return the first page of results.
func (p QueryDatabaseParams) StartCursor(start_cursor uuid.UUID) QueryDatabaseParams {
	p["start_cursor"] = start_cursor
	return p
}

// The number of items from the full list desired in the response. Maximum: 100
func (p QueryDatabaseParams) PageSize(page_size int) QueryDatabaseParams {
	p["page_size"] = page_size
	return p
}
