package endpoints_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/endpoints"
)

func TestCreateDatabase(t *testing.T) {
	t.Parallel()

	Fetch(
		"https://developers.notion.com/reference/create-a-database",
	).Generate(GenericStructRef{Name: "Database"}, ParamAnnotations{
		"parent":     jen.Id("Parent"),
		"title":      jen.Index().Id("RichText"),
		"properties": jen.Id("PropertySchema"),
	})
}

func TestQueryDatabase(t *testing.T) {
	t.Parallel()

	Fetch(
		"https://developers.notion.com/reference/post-database-query",
	).Generate(GenericStructRef{Name: "Pagination", GenericTypeArg: "Page"}, ParamAnnotations{
		"database_id":  UUID,
		"filter":       jen.Id("Filter"),
		"sorts":        jen.Id("Sort"),
		"start_cursor": UUID,
		"page_size":    jen.Int(),
	})
}

func TestRetrieveDatabase(t *testing.T) {
	t.Parallel()

	Fetch(
		"https://developers.notion.com/reference/retrieve-a-database",
	).Generate(StructRef("Database"), ParamAnnotations{
		"database_id": jen.Qual("github.com/google/uuid", "UUID"),
	})
}
