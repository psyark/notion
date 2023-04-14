package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Query a database
// https://developers.notion.com/reference/post-database-query
func (c *Client) QueryDatabase(ctx context.Context, database_id uuid.UUID, params *QueryDatabaseParams, options ...callOption) (*PagePagination, error) {
	result := &PagePagination{}
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

type QueryDatabaseParams struct {
	Filter      any    `json:"filter"`       // When supplied, limits which pages are returned based on the [filter conditions](ref:post-database-query-filter).
	Sorts       any    `json:"sorts"`        // When supplied, orders the results based on the provided [sort criteria](ref:post-database-query-sort).
	StartCursor string `json:"start_cursor"` // When supplied, returns a page of results starting after the cursor provided. If not supplied, this endpoint will return the first page of results.
	PageSize    int    `json:"page_size"`    // The number of items from the full list desired in the response. Maximum: 100
}
