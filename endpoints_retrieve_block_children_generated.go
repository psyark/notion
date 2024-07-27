// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/get-block-children

package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

/*
Returns a paginated array of child [block objects](ref:block) contained in the block using the ID specified. In order to receive a complete representation of a block, you may need to recursively retrieve the block children of child blocks.

> 👍
>
> Page content is represented by block children. See the [Working with page content guide](https://developers.notion.com/docs/working-with-page-content#modeling-content-as-blocks) for more information.

Returns only the first level of children for the specified block. See [block objects](ref:block) for more detail on determining if that block has nested children.

The response may contain fewer than `page_size` of results.

See [Pagination](https://developers.notion.com/reference/intro#pagination) for details about how to use a cursor to iterate through the list.

> 📘 Integration capabilities
>
> This endpoint requires an integration to have read content capabilities. Attempting to call this API without read content capabilities will return an HTTP response with a 403 status code. For more information on integration capabilities, see the [capabilities guide](ref:capabilities).

### Errors

Returns a 404 HTTP response if the block specified by `id` doesn't exist, or if the integration doesn't have access to the block.

Returns a 400 or 429 HTTP response if the request exceeds the [request limits](ref:request-limits).

_Note: Each Public API endpoint can return several possible error codes. See the [Error codes section](https://developers.notion.com/reference/status-codes#error-codes) of the Status codes documentation for more information._
*/
func (c *Client) RetrieveBlockChildren(ctx context.Context, block_id uuid.UUID, options ...CallOption) (*Pagination[Block], error) {
	return call(
		ctx,
		c.accessToken,
		http.MethodGet,
		fmt.Sprintf("/v1/blocks/%v/children", block_id),
		nil,
		accessValue[*Pagination[Block]],
		options...,
	)
}
