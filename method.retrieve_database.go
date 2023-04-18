package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Retrieve a database
// https://developers.notion.com/reference/retrieve-a-database
func (c *Client) RetrieveDatabase(ctx context.Context, database_id uuid.UUID, options ...callOption) (*Database, error) {
	result := &Database{}
	co := &callOptions{
		method: http.MethodGet,
		params: nil,
		path:   fmt.Sprintf("/v1/databases/%v", database_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}
