package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Append block children
// https://developers.notion.com/reference/patch-block-children
func (c *Client) AppendBlockChildren(ctx context.Context, block_id uuid.UUID, params *AppendBlockChildrenParams, options ...callOption) (*Pagination, error) {
	result := &Pagination{}
	co := &callOptions{
		method: http.MethodPatch,
		params: params,
		path:   fmt.Sprintf("/v1/blocks/%v/children", block_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}

type AppendBlockChildrenParams struct {
	Children []Block `json:"children"` // Child content to append to a container block as an array of [block objects](ref:block)
}
