package endpoints_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/endpoints"
)

func TestSearchByTitle(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/post-search").Generate(GenericStructRef{"Pagination", "PageOrDatabase"}, ParamAnnotations{
		"query":        jen.String(),
		"sort":         jen.Id("Sort"),
		"filter":       jen.Id("SearchFilter"),
		"start_cursor": jen.String(),
		"page_size":    jen.Int(),
	})
}
