package notion

import (
	"context"
	"fmt"
	"net/http"
)

// Retrieve a page
// https://developers.notion.com/reference/retrieve-a-page
func (c *Client) RetrievePage(ctx context.Context, page_id string, options ...callOption) (*Page, error) {
	result := &Page{}
	co := &callOptions{
		method: http.MethodGet,
		params: nil,
		path:   fmt.Sprintf("/v1/pages/%v", page_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}
