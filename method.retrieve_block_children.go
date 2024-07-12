// Code generated by notion.doc2api; DO NOT EDIT.

package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Retrieve block children
// https://developers.notion.com/reference/get-block-children
func (c *Client) RetrieveBlockChildren(ctx context.Context, block_id uuid.UUID, options ...callOption) (*Pagination, error) {
	return call(
		ctx,
		c.accessToken,
		http.MethodGet,
		fmt.Sprintf("/v1/blocks/%v/children", block_id),
		nil,
		func(u *Pagination) *Pagination {
			return u
		},
		options...,
	)
}
