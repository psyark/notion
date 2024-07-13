package endpoints_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/endpoints"
)

func TestCreatePage(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/post-page").Generate(StructRef("Page"), ParamAnnotations{
		"parent":     jen.Id("Parent"),
		"properties": jen.Map(jen.String()).Id("PropertyValue"),
		"children":   jen.Index().Id("Block"),
		"icon":       jen.Id("FileOrEmoji"),
		"cover":      jen.Id("File"),
	})
}

func TestRetrievePage(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/retrieve-a-page").Generate(StructRef("Page"), ParamAnnotations{
		"page_id": UUID,
	})
}

func TestRetrievePagePropertyItem(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/retrieve-a-page-property").Generate(Interface("PropertyItemOrPropertyItemPagination"), ParamAnnotations{
		"page_id":     UUID,
		"property_id": jen.String(),
	})
}

func TestUpdatePageProperties(t *testing.T) {
	t.Parallel()

	Fetch("https://developers.notion.com/reference/patch-page").Generate(StructRef("Page"), ParamAnnotations{
		"page_id":    UUID,
		"properties": jen.Map(jen.String()).Id("PropertyValue"),
		"in_trash":   jen.Bool(),
		"icon":       jen.Id("FileOrEmoji"),
		"cover":      jen.Id("File"),
	})
}
