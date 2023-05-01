package methoddoc

func init() {
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/get-block-children",
		returnType: returnsStructRef("BlockPagination"),
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
}
