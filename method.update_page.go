package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

// Update page
// https://developers.notion.com/reference/patch-page
func (c *Client) UpdatePage(ctx context.Context, page_id uuid.UUID, params *UpdatePageParams, options ...callOption) (*Page, error) {
	result := &Page{}
	co := &callOptions{
		method: http.MethodPatch,
		params: params,
		path:   fmt.Sprintf("/v1/pages/%v", page_id),
		result: result,
	}
	for _, o := range options {
		o(co)
	}
	return result, c.call(ctx, co)
}

type UpdatePageParams struct {
	Properties PropertyValueMap `json:"properties,omitempty"` // The property values to update for the page. The keys are the names or IDs of the property and the values are property values. If a page property ID is not included, then it is not changed.
	Archived   *bool            `json:"archived,omitempty"`   // Whether the page is archived (deleted). Set to true to archive a page. Set to false to un-archive (restore) a page.
	Icon       FileOrEmoji      `json:"icon,omitempty"`       // A page icon for the page. Supported types are [external file object](https://developers.notion.com/reference/file-object) or [emoji object](https://developers.notion.com/reference/emoji-object).
	Cover      *ExternalFile    `json:"cover,omitempty"`      // A cover image for the page. Only [external file objects](https://developers.notion.com/reference/file-object) are supported.
}
