package endpoints_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/endpoints"
)

func TestAppendBlockChildren(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/patch-block-children").Generate(GenericStructRef{Name: "Pagination", GenericTypeArg: "Block"}, ParamAnnotations{
		"block_id": UUID,
		"children": jen.Index().Id("Block"),
		"after":    UUID,
	})
}

func TestRetrieveBlockChildren(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/get-block-children").Generate(GenericStructRef{Name: "Pagination", GenericTypeArg: "Block"}, ParamAnnotations{
		"block_id": UUID,
	})
}

func TestDeleteBlock(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/delete-a-block").Generate(StructRef("Block"), ParamAnnotations{
		"block_id": UUID,
	})
}
