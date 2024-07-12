package methoddoc

import "github.com/dave/jennifer/jen"

func init() {
	registerConverter(converter{
		url:        "https://developers.notion.com/reference/post-page",
		returnType: returnsStructRef("Page"),
		localCopyOfBodyParams: ssrPropsParams{
			{
				Desc:     "The parent page or database where the new page is inserted, represented as a JSON object with a `page_id` or `database_id` key, and the corresponding ID.",
				In:       "body",
				Name:     "parent",
				Required: true,
				Type:     "json",
				typeCode: jen.Id("Parent"),
			},
			{
				Desc:     "The values of the page’s properties. If the `parent` is a database, then the schema must match the parent database’s properties. If the `parent` is a page, then the only valid object key is `title`.",
				In:       "body",
				Name:     "properties",
				Required: true,
				Type:     "json",
				typeCode: jen.Map(jen.String()).Id("PropertyValue"),
			},
			{
				Desc:     "The content to be rendered on the new page, represented as an array of [block objects](https://developers.notion.com/reference/block).",
				In:       "body",
				Name:     "children",
				Type:     "array_string",
				typeCode: jen.Index().Id("Block"),
			},
			{
				Desc:     "The icon of the new page. Either an [emoji object](https://developers.notion.com/reference/emoji-object) or an [external file object](https://developers.notion.com/reference/file-object)..",
				In:       "body",
				Name:     "icon",
				Type:     "json",
				typeCode: jen.Id("FileOrEmoji"),
			},
			{
				Desc:     "The cover image of the new page, represented as a [file object](https://developers.notion.com/reference/file-object).",
				In:       "body",
				Name:     "cover",
				Type:     "json",
				typeCode: jen.Id("File"),
			},
		},
	})

	registerConverter(converter{
		url:        "https://developers.notion.com/reference/retrieve-a-page",
		returnType: returnsStructRef("Page"),
		localCopyOfPathParams: ssrPropsParams{
			{
				Desc:     "Identifier for a Notion page",
				In:       "path",
				Name:     "page_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
	})

	registerConverter(converter{
		url:        "https://developers.notion.com/reference/retrieve-a-page-property",
		returnType: returnsInterface("PropertyItemOrPropertyItemPagination"),
		localCopyOfPathParams: ssrPropsParams{
			{
				Desc:     "Identifier for a Notion page",
				In:       "path",
				Name:     "page_id",
				Type:     "string",
				typeCode: UUID,
			},
			{
				Desc:     "Identifier for a page [property](https://developers.notion.com/reference/page#all-property-values)",
				In:       "path",
				Name:     "property_id",
				Type:     "string",
				typeCode: jen.String(),
			},
		},
	})

	registerConverter(converter{
		url:        "https://developers.notion.com/reference/patch-page",
		returnType: returnsStructRef("Page"),
		localCopyOfPathParams: ssrPropsParams{
			{
				Desc:     "The identifier for the Notion page to be updated.",
				In:       "path",
				Name:     "page_id",
				Type:     "string",
				typeCode: UUID,
			},
		},
		localCopyOfBodyParams: ssrPropsParams{
			{
				Desc:     "The property values to update for the page. The keys are the names or IDs of the property and the values are property values. If a page property ID is not included, then it is not changed.",
				In:       "body",
				Name:     "properties",
				Type:     "json",
				typeCode: jen.Map(jen.String()).Id("PropertyValue"),
			},
			{
				Default:  "false",
				Desc:     "Set to true to delete a block. Set to false to restore a block.",
				In:       "body",
				Name:     "in_trash",
				Type:     "boolean",
				typeCode: jen.Bool(),
			},
			{
				Desc:     "A page icon for the page. Supported types are [external file object](https://developers.notion.com/reference/file-object) or [emoji object](https://developers.notion.com/reference/emoji-object).",
				In:       "body",
				Name:     "icon",
				Type:     "json",
				typeCode: jen.Id("FileOrEmoji"),
			},
			{
				Desc:     "A cover image for the page. Only [external file objects](https://developers.notion.com/reference/file-object) are supported.",
				In:       "body",
				Name:     "cover",
				Type:     "json",
				typeCode: jen.Id("File"),
			},
		},
	})
}
