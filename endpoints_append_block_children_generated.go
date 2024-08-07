// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/patch-block-children

package notion

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"net/http"
)

/*
Creates and appends new children blocks to the parent `block_id` specified. Blocks can be parented by other blocks, pages, or databases.

Returns a paginated list of newly created first level children [block objects](ref:block).

Existing blocks cannot be moved using this endpoint. Blocks are appended to the bottom of the parent block. To append a block in a specific place other than the bottom of the parent block, use the `"after"` parameter and set its value to the ID of the block that the new block should be appended after. Once a block is appended as a child, it can't be moved elsewhere via the API.

For blocks that allow children, we allow up to **two** levels of nesting in a single request.

There is a limit of **100 block children** that can be appended by a single API request. Arrays of block children longer than 100 will result in an error.

> 📘 Integration capabilities
>
> This endpoint requires an integration to have insert content capabilities. Attempting to call this API without insert content capabilities will return an HTTP response with a 403 status code. For more information on integration capabilities, see the [capabilities guide](ref:capabilities).

### Errors

Returns a 404 HTTP response if the block specified by `id` doesn't exist, or if the integration doesn't have access to the block.

Returns a 400 or 429 HTTP response if the request exceeds the [request limits](ref:request-limits).

_Note: Each Public API endpoint can return several possible error codes. To see a full description of each type of error code, see the [Error codes section](https://developers.notion.com/reference/status-codes#error-codes) of the Status codes documentation._
*/
func (c *Client) AppendBlockChildren(ctx context.Context, block_id uuid.UUID, params AppendBlockChildrenParams, options ...CallOption) (*Pagination[Block], error) {
	return call(
		ctx,
		c.accessToken,
		http.MethodPatch,
		fmt.Sprintf("/v1/blocks/%v/children", block_id),
		params,
		accessValue[*Pagination[Block]],
		options...,
	)
}

type AppendBlockChildrenParams map[string]any

// Child content to append to a container block as an array of [block objects](ref:block)
func (p AppendBlockChildrenParams) Children(children []Block) AppendBlockChildrenParams {
	p["children"] = children
	return p
}

// The ID of the existing block that the new block should be appended after.
func (p AppendBlockChildrenParams) After(after uuid.UUID) AppendBlockChildrenParams {
	p["after"] = after
	return p
}
