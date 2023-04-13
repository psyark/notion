package notion

import (
	"context"
	"fmt"
	"net/http"

	uuid "github.com/google/uuid"
)

// Retrieve a page property item
// https://developers.notion.com/reference/retrieve-a-page-property
func (c *Client) RetrievePagePropertyItem(ctx context.Context, page_id uuid.UUID, property_id string, options ...callOption) (PropertyItemOrPropertyItemPagination, error) {
	result := &propertyItemOrPropertyItemPaginationUnmarshaler{}
	co := &callOptions{
		method: http.MethodGet,
		params: nil,
		path:   fmt.Sprintf("/v1/pages/%v/properties/%v", page_id, property_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result.value, c.call(ctx, co)
}
