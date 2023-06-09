package methoddoc

import "github.com/dave/jennifer/jen"

func init() {
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/patch-block-children",
		returnType: returnsStructRef("Pagination"),
		localCopyOfPathParams: ssrPropsParams{
			ssrPropsParam{
				Desc:     "Identifier for a [block](ref:block)",
				In:       "path",
				Name:     "block_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
		localCopyOfBodyParams: ssrPropsParams{
			ssrPropsParam{
				Desc:     "Child content to append to a container block as an array of [block objects](ref:block)",
				In:       "body",
				Name:     "children",
				Required: true,
				Type:     "array_object",
				typeCode: jen.Index().Id("Block"),
			},
		},
	})
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/get-block-children",
		returnType: returnsStructRef("Pagination"),
		localCopyOfPathParams: ssrPropsParams{
			ssrPropsParam{
				Desc:     "Identifier for a [block](ref:block)",
				In:       "path",
				Name:     "block_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
	})
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/delete-a-block",
		returnType: returnsStructRef("Block"),
		localCopyOfPathParams: ssrPropsParams{
			ssrPropsParam{
				Desc:     "Identifier for a Notion block",
				In:       "path",
				Name:     "block_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
	})
}
