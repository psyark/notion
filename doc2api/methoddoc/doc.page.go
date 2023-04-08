package methoddoc

func init() {
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/retrieve-a-page",
		returnType: returnsStructRef("Page"),
	})
	// registerConverter(converter{
	// 	url: "https://developers.notion.com/reference/post-page",
	// 	localCopyOfBodyParams: []ssrPropsParam{
	// 		{
	// 			Desc:     "The parent page or database where the new page is inserted, represented as a JSON object with a `page_id` or `database_id` key, and the corresponding ID.",
	// 			In:       "body",
	// 			Name:     "parent",
	// 			Required: true,
	// 			Type:     "json",
	// 			typeCode: jen.Id("Parent"),
	// 		},
	// 		{
	// 			Desc:     "The values of the page’s properties. If the `parent` is a database, then the schema must match the parent database’s properties. If the `parent` is a page, then the only valid object key is `title`.",
	// 			In:       "body",
	// 			Name:     "properties",
	// 			Required: true,
	// 			Type:     "json",
	// 			typeCode: jen.Map(jen.String()).Id("PropertyValue"),
	// 		},
	// 		{
	// 			Desc:     "The content to be rendered on the new page, represented as an array of [block objects](https://developers.notion.com/reference/block).",
	// 			In:       "body",
	// 			Name:     "children",
	// 			Type:     "array_string",
	// 			typeCode: jen.Index().Id("Block"),
	// 		},
	// 		{
	// 			Desc:     "The icon of the new page. Either an [emoji object](https://developers.notion.com/reference/emoji-object) or an [external file object](https://developers.notion.com/reference/file-object)..",
	// 			In:       "body",
	// 			Name:     "icon",
	// 			Type:     "json",
	// 			typeCode: jen.Id("EmojiOrExternalFile"),
	// 		},
	// 		{
	// 			Desc:     "The cover image of the new page, represented as a [file object](https://developers.notion.com/reference/file-object).",
	// 			In:       "body",
	// 			Name:     "cover",
	// 			Type:     "json",
	// 			typeCode: jen.Id("File"),
	// 		},
	// 	},
	// 	returnType: returnsStructRef("Page"),
	// })
}
