package methoddoc

func init() {
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/retrieve-a-page",
		returnType: returnsStructRef("Page"),
		localCopy: []ssrPropsParam{
			{
				Desc: "Identifier for a Notion page",
				In:   "path",
				Name: "page_id",
				Type: "string",
			},
			{
				Desc: "A list of page property value IDs associated with the page. Use this param to limit the response to a specific page property value or values. To retrieve multiple properties, specify each page property ID. For example: `?filter_properties=iAk8&filter_properties=b7dh`.",
				In:   "query",
				Name: "filter_properties",
				Type: "string",
			},
			{
				Default:  "2022-06-28",
				In:       "header",
				Name:     "Notion-Version",
				Required: true,
				Type:     "string",
			},
		},
	})
	registerConverter(converter{
		url: "https://developers.notion.com/reference/post-page",
		localCopy: []ssrPropsParam{
			{
				Desc:     "The parent page or database where the new page is inserted, represented as a JSON object with a `page_id` or `database_id` key, and the corresponding ID.",
				In:       "body",
				Name:     "parent",
				Required: true,
				Type:     "json",
			},
			{
				Desc:     "The values of the page’s properties. If the `parent` is a database, then the schema must match the parent database’s properties. If the `parent` is a page, then the only valid object key is `title`.",
				In:       "body",
				Name:     "properties",
				Required: true,
				Type:     "json",
			},
			{
				Desc: "The content to be rendered on the new page, represented as an array of [block objects](https://developers.notion.com/reference/block).",
				In:   "body",
				Name: "children",
				Type: "array_string",
			},
			{
				Default:  "2022-06-28",
				In:       "header",
				Name:     "Notion-Version",
				Required: true,
				Type:     "string",
			},
			{
				Desc: "The icon of the new page. Either an [emoji object](https://developers.notion.com/reference/emoji-object) or an [external file object](https://developers.notion.com/reference/file-object)..",
				In:   "body",
				Name: "icon",
				Type: "json",
			},
			{
				Desc: "The cover image of the new page, represented as a [file object](https://developers.notion.com/reference/file-object).",
				In:   "body",
				Name: "cover",
				Type: "json",
			},
		},
		returnType: returnsStructRef("Page"),
	})
}
