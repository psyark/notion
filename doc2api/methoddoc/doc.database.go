package methoddoc

import "github.com/dave/jennifer/jen"

func init() {
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/post-database-query",
		returnType: returnsStructRef("Pagination"),
		localCopyOfPathParams: ssrPropsParams{
			{
				Desc:     "Identifier for a Notion database.",
				In:       "path",
				Name:     "database_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
		localCopyOfBodyParams: ssrPropsParams{
			{
				Desc:     "When supplied, limits which pages are returned based on the [filter conditions](ref:post-database-query-filter).",
				In:       "body",
				Name:     "filter",
				Type:     "json",
				typeCode: jen.Id("Filter"),
			},
			{
				Desc:     "When supplied, orders the results based on the provided [sort criteria](ref:post-database-query-sort).",
				In:       "body",
				Name:     "sorts",
				Type:     "array_object",
				typeCode: jen.Id("Sort"),
			},
			{
				Desc:     "When supplied, returns a page of results starting after the cursor provided. If not supplied, this endpoint will return the first page of results.",
				In:       "body",
				Name:     "start_cursor",
				Type:     "string",
				typeCode: jen.Qual("github.com/google/uuid", "UUID"), // body.start_cursor should be a valid uuid or `undefined`, instead was ""
			},
			{
				Default:  "100",
				Desc:     "The number of items from the full list desired in the response. Maximum: 100",
				In:       "body",
				Name:     "page_size",
				Type:     "int",
				typeCode: jen.Int(),
			},
		},
	})
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/retrieve-a-database",
		returnType: returnsStructRef("Database"),
		localCopyOfPathParams: ssrPropsParams{
			{
				Desc:     "An identifier for the Notion database.",
				In:       "path",
				Name:     "database_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
	})
}
