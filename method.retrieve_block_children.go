package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Retrieve block children
// https://developers.notion.com/reference/get-block-children
func (c *Client) RetrieveBlockChildren(ctx context.Context, block_id uuid.UUID, options ...callOption) (*BlockPagination, error) {
	result := &BlockPagination{}
	co := &callOptions{
		method: http.MethodGet,
		params: nil,
		path:   fmt.Sprintf("/v1/blocks/%v/children", block_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}
