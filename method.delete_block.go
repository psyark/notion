package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Delete a block
// https://developers.notion.com/reference/delete-a-block
func (c *Client) DeleteBlock(ctx context.Context, block_id uuid.UUID, options ...callOption) (*Block, error) {
	result := &Block{}
	co := &callOptions{
		method: http.MethodDelete,
		params: nil,
		path:   fmt.Sprintf("/v1/blocks/%v", block_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}
