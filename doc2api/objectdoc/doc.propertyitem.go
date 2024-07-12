package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	var propertyItem, rollup *adaptiveObject

	addTest := func(b *builder) func(e blockElement) {
		return func(e blockElement) {
			b.addUnmarshalTest("PropertyItemOrPropertyItemPaginationMap", e.Text)
		}
	}

	registerTranslator(
		"https://developers.notion.com/reference/property-item-object",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A property_item object describes the identifier, type, and value of a page property. It's returned from the Retrieve a page property item",
			}, func(e blockElement) {
				propertyItem = b.addAdaptiveObject("PropertyItem", "type", e.Text)
				propertyItem.addToUnion(b.addUnionToGlobalIfNotExists("PropertyItemOrPropertyItemPagination", "object"))
			})
		},
		func(c *comparator, b *builder) /* All property items */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "All property items",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Each page property item object contains the following keys. In addition, it will contain a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Always \"property_item\".",
				ExampleValue: `"property_item"`,
				Property:     "object",
				Type:         `"property_item"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("PropertyItem").addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
				ExampleValue: `"f%5C%5C%3Ap"`,
				Property:     "id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("PropertyItem").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Type of the property. Possible values are \"rich_text\", \"number\", \"select\", \"multi_select\", \"date\", \"formula\", \"relation\", \"rollup\", \"title\", \"people\", \"files\", \"checkbox\", \"url\", \"email\", \"phone_number\", \"created_time\", \"created_by\", \"last_edited_time\", and \"last_edited_by\".",
				ExampleValue: `"rich_text"`,
				Property:     "type",
				Type:         "string (enum)",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* Paginated property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Paginated property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information.",
			}, func(e blockElement) {
				// TODO 良い名前
				ppi := b.addAdaptiveObject("PaginatedPropertyInfo", "type", e.Text).addFields(
					&field{name: "id", typeCode: jen.String(), comment: "undocumented"},
				)
				for _, derived := range []string{"title", "rich_text", "relation", "people"} {
					ppi.addAdaptiveFieldWithEmptyStruct(derived, "")
				}
				ppi.addAdaptiveFieldWithType("rollup", "undocumented", jen.Id("Rollup"))
			})
			c.nextMustParameter(parameterElement{
				Property:     "object",
				Type:         `"list"`,
				Description:  `Always "list".`,
				ExampleValue: `"list"`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         `"property_item"`,
				Description:  `Always "property_item".`,
				ExampleValue: `"property_item"`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "results",
				Type:         "list",
				Description:  "List of property_item objects.",
				ExampleValue: `[{"object": "property_item", "id": "vYdV", "type": "relation", "relation": { "id": "535c3fb2-95e6-4b37-a696-036e5eac5cf6"}}... ]`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "property_item",
				Type:         "object",
				Description:  "A property_item object that describes the property.",
				ExampleValue: `{"id": "title", "next_url": null, "type": "title", "title": {}}`,
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "next_url",
				Type:         "string or null",
				Description:  "The URL the user can request to get the next page of results.",
				ExampleValue: `"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/vYdV?start_cursor=LYxaUO&page_size=25"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("PaginatedPropertyInfo").addFields(e.asField(jen.Id("*").String()))
			})
		},
		func(c *comparator, b *builder) /* Title property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Title property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Title property value objects contain an array of rich text objects within the title property.",
			}, func(e blockElement) {
				// ドキュメントには "array of rich text" と書いてあるが間違い
				propertyItem.addAdaptiveFieldWithType("title", e.Text, jen.Id("RichText"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Name\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"title\",\n        \"type\": \"title\",\n        \"title\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"The title\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"The title\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"title\",\n      \"next_url\": null,\n      \"type\": \"title\",\n      \"title\": {}\n    }\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Rich Text property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rich Text property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Rich Text property value objects contain an array of rich text objects within the rich_text property.",
			}, func(e blockElement) {
				// ドキュメントには "array of rich text" と書いてあるが間違い
				propertyItem.addAdaptiveFieldWithType("rich_text", e.Text, jen.Id("RichText"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Details\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"Some more text with \",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"Some more text with \",\n          \"href\": null\n        }\n      },\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"fun formatting\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": true,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"fun formatting\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"NVv^\",\n      \"next_url\": null,\n      \"type\": \"rich_text\",\n      \"rich_text\": {}\n    }\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Number property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Number property value objects contain a number within the number property.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("number", e.Text, jen.Op("*").Float64())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Quantity\": {\n    \"object\": \"property_item\",\n    \"id\": \"XpXf\",\n    \"type\": \"number\",\n    \"number\": 1234\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Select property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Select property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Select property value objects contain the following data within the select property:",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("select", e.Text, jen.Op("*").Id("Option"))
				propertyItem.addAdaptiveFieldWithType("status", "undocumented", jen.Op("*").Id("Option"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "ID of the option.\n\nWhen updating a select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				Property:     "id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Name of the option as it appears in Notion.\n\nIf the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Option\": {\n    \"object\": \"property_item\",\n    \"id\": \"%7CtzR\",\n    \"type\": \"select\",\n    \"select\": {\n      \"id\": \"64190ec9-e963-47cb-bc37-6a71d6b71206\",\n      \"name\": \"Option 1\",\n      \"color\": \"orange\"\n    }\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Multi-select property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Multi-select property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Multi-select property value objects contain an array of multi-select option values within the multi_select property.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("multi_select", e.Text, jen.Index().Id("Option"))
			})
		},
		func(c *comparator, b *builder) /* Multi-select option values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Multi-select option values",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				Property:     "id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Tags\": {\n    \"object\": \"property_item\",\n    \"id\": \"z%7D%5C%3C\",\n    \"type\": \"multi_select\",\n    \"multi_select\": [\n      {\n        \"id\": \"91e6959e-7690-4f55-b8dd-d3da9debac45\",\n        \"name\": \"A\",\n        \"color\": \"orange\"\n      },\n      {\n        \"id\": \"2f998e2d-7b1c-485b-ba6b-5e6a815ec8f5\",\n        \"name\": \"B\",\n        \"color\": \"purple\"\n      }\n    ]\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Date property values */ {
			var propertyItemDate *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date property values",
			}, func(e blockElement) {
				propertyItemDate = propertyItem.addAdaptiveFieldWithSpecificObject("date", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date property value objects contain the following data within the date property:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "start",
				Type:         "string (ISO 8601 date and time)",
				Description:  "An ISO 8601 format date, with optional time.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
			}, func(e parameterElement) {
				propertyItemDate.addFields(e.asField(jen.Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "end",
				Type:         "string (optional, ISO 8601 date and time)",
				Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
			}, func(e parameterElement) {
				propertyItemDate.addFields(e.asField(jen.Op("*").Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Property:     "time_zone",
				Type:         "string (optional, enum)",
				Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
				ExampleValue: `"America/Los_Angeles"`,
			}, func(e parameterElement) {
				propertyItemDate.addFields(e.asField(jen.Op("*").String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Shipment Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"i%3Ahj\",\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\",\n      \"end\": null,\n      \"time_zone\": null\n    }\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Formula property value objects represent the result of evaluating a formula described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("formula", e.Text, jen.Op("*").Id("Formula"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The type of the formula result. Possible values are \"string\", \"number\", \"boolean\", and \"date\".",
				ExampleValue: "",
				Property:     "type",
				Type:         "string (enum)",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* String formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "String formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "String formula property values contain an optional string within the string property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Number formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Number formula property values contain an optional number within the number property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Boolean formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Boolean formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Boolean formula property values contain a boolean within the boolean property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Date formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date formula property values contain an optional date property value within the date property.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Formula\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"formula\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Relation property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Relation property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Relation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("relation", e.Text, jen.Op("*").Id("PageReference")) // 単一
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Project\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"vYdV\",\n      \"next_url\": null,\n      \"type\": \"relation\",\n      \"relation\": {}\n    }\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rollup property values",
			}, func(e blockElement) {
				// PropertyValueとPropertyItemで共通の Rollup を使う
				// PropertyValue のArray Rollupにはfunction が無いなど不正確であるため、
				// 比較的ドキュメントが充実しているこちらで作成を行う
				// https://developers.notion.com/reference/property-value-object#rollup-property-values
				// https://developers.notion.com/reference/property-item-object#rollup-property-values
				rollup = b.addAdaptiveObject("Rollup", "type", "")
				propertyItem.addAdaptiveFieldWithType("rollup", e.Text, jen.Op("*").Id("Rollup"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Rollup property value objects represent the result of evaluating a rollup described in thedatabase's properties. The property is returned as a list object of type property_item with a list of relation items used to computed the rollup under results.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A rollup property item is also returned under the property_type key that describes the rollup aggregation and computed result.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "In order to avoid timeouts, if the rollup has a with a large number of aggregations or properties the endpoint returns a next_cursor value that is used to determinate the aggregation value so far for the subset of relations that have been paginated through.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Once has_more is false, then the final rollup value is returned.  See the Pagination documentation for more information on pagination in the Notion API.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Computing the values of following aggregations are not supported. Instead the endpoint returns a list of property_item objects for the rollup:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "show_unique (Show unique values)unique (Count unique values)median(Median)",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The type of rollup. Possible values are \"number\", \"date\", \"array\", \"unsupported\" and \"incomplete\".",
				ExampleValue: "",
				Property:     "type",
				Type:         "string (enum)",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "Describes the aggregation used. \nPossible values include: count,  count_values,  empty,  not_empty,  unique,  show_unique,  percent_empty,  percent_not_empty,  sum,  average,  median,  min,  max,  range,  earliest_date,  latest_date,  date_range,  checked,  unchecked,  percent_checked,  percent_unchecked,  count_per_group,  percent_per_group,  show_original",
				ExampleValue: "",
				Property:     "function",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				rollup.addFields(e.asField(jen.String()))
			})
		},
		func(c *comparator, b *builder) /* Number rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Number rollup property values contain a number within the number property.",
			}, func(e blockElement) {
				rollup.addAdaptiveFieldWithType("number", e.Text, jen.Op("*").Float64())
			})
		},
		func(c *comparator, b *builder) /* Date rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date rollup property values contain a date property value within the date property.",
			}, func(e blockElement) {
				rollup.addAdaptiveFieldWithType("date", e.Text, jen.Op("*").Id("PropertyItemDate"))
			})
		},
		func(c *comparator, b *builder) /* Array rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Array rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Array rollup property values contain an array of property_item objects within the results property.",
			}, func(e blockElement) {
				// ドキュメントには array of property_item とあるが、
				// type="rich_text"の場合に入る値などから
				// array of property_value が正しいと判断している
				rollup.addAdaptiveFieldWithType("array", e.Text, jen.Index().Id("PropertyValue"))
			})
		},
		func(c *comparator, b *builder) /* Incomplete rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Incomplete rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Rollups with an aggregation with more than one page of aggregated results will return a rollup object of type \"incomplete\". To obtain the final value paginate through the next values in the rollup using the next_cursor or next_url property.",
			}, func(e blockElement) {
				rollup.addAdaptiveFieldWithType("incomplete", e.Text, jen.Struct())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Rollup\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }...\n    ],\t\n    \"next_cursor\": \"1QaTunT5\",\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"y}~p\",\n      \"next_url\": \"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/y%7D~p?start_cursor=1QaTunT5&page_size=25\",\n      \"type\": \"rollup\",\n      \"rollup\": {\n        \"function\": \"sum\",\n        \"type\": \"incomplete\",\n        \"incomplete\": {}\n      }\n    }\n  }\n}",
			}, func(e blockElement) {
				e.Text = strings.Replace(e.Text, "...", "", 1)
				b.addUnmarshalTest("PropertyItemOrPropertyItemPaginationMap", e.Text)
			})
		},
		func(c *comparator, b *builder) /* People property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "People property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "People property value objects contain an array of user objects within the people property.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("people", e.Text, jen.Id("User")) // 単一では？
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Owners\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"people\",\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"285e5768-3fdc-4742-ab9e-125f9050f3b8\",\n        \"name\": \"Example Avo\",\n        \"avatar_url\": null,\n        \"type\": \"person\",\n        \"person\": {\n          \"email\": \"avo@example.com\"\n        }\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO peopleに配列が入っているが、実際はPaginated property valuesになるはずなのでドキュメントの不具合では？確認する
			})
		},
		func(c *comparator, b *builder) /* Files property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Files property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("files", e.Text, jen.Index().Id("File"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Files\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"files\",\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": \"https://website.domain/images/space.png\"\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO サンプルコードの間違いを訂正してテストを通す
			})
		},
		func(c *comparator, b *builder) /* Checkbox property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Checkbox property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Checkbox property value objects contain a boolean within the checkbox property.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("checkbox", e.Text, jen.Bool())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Done?\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"checkbox\",\n    \"checkbox\": true\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* URL property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "URL property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("url", e.Text, jen.Op("*").String())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Website\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"url\",\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Email property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Email property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Email property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("email", e.Text, jen.Op("*").String())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Shipper's Contact\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"email\",\n    \"email\": \"hello@test.com\"\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Phone number property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Phone number property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Phone number property value objects contain a string within the phone_number property. No structure is enforced.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("phone_number", e.Text, jen.Op("*").String())
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Shipper's No.\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"phone_number\",\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
			}, addTest(b))
		},
		func(c *comparator, b *builder) /* Created time property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Created time property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("created_time", e.Text, jen.Id("ISO8601String"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Created Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"create_time\",\n  \t\"created_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
			}, func(e blockElement) {
				e.Text = strings.Replace(e.Text, "create_time", "created_time", 1) // ドキュメントの不具合
				b.addUnmarshalTest("PropertyItemOrPropertyItemPaginationMap", e.Text)
			})
		},
		func(c *comparator, b *builder) /* Created by property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Created by property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("created_by", e.Text, jen.Op("*").Id("User"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Created By\": {\n    \"created_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"dsEa\": {\n    \"created_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
		},
		func(c *comparator, b *builder) /* Last edited time property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Last edited time property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("last_edited_time", e.Text, jen.Id("ISO8601String"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Last Edited Time\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"as0w\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
		},
		func(c *comparator, b *builder) /* Last edited by property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Last edited by property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page.",
			}, func(e blockElement) {
				propertyItem.addAdaptiveFieldWithType("last_edited_by", e.Text, jen.Op("*").Id("User"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Last Edited By\": {\n    \"last_edited_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"as12\": {\n    \"last_edited_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
		},
		func(c *comparator, b *builder) {
			propertyItem.addFields(undocumentedRequestID)
		},
	)
}
