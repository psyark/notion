package objects_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestPropertyItem(t *testing.T) {
	t.Parallel()

	addTest := func(e *Block, b *CodeBuilder) {
		b.AddUnmarshalTest("PropertyItemOrPropertyItemPaginationMap", e.Text)
	}

	c := converter.FetchDocument("https://developers.notion.com/reference/property-item-object")

	var propertyItem *AdaptiveObject

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A property_item object describes the identifier, type, and value of a page property. It's returned from the Retrieve a page property item",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem = b.AddAdaptiveObject("PropertyItem", "type", e.Text)

		union := converter.RegisterUnionInterface("PropertyItemOrPropertyItemPagination", "object")
		converter.RegisterUnionMember(union, propertyItem, "")
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "All property items"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Each page property item object contains the following keys. In addition, it will contain a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below."})

	c.ExpectParameter(&Parameter{
		Property:     "object",
		Type:         `"property_item"`,
		Description:  `Always "property_item".`,
		ExampleValue: `"property_item"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		propertyItem.AddFields(b.NewDiscriminatorField(e))
	})

	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string",
		Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
		ExampleValue: `"f%5C%5C%3Ap"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		propertyItem.AddFields(b.NewField(e, jen.String()))
	})

	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "string (enum)",
		Description:  "Type of the property. Possible values are \"rich_text\", \"number\", \"select\", \"multi_select\", \"date\", \"formula\", \"relation\", \"rollup\", \"title\", \"people\", \"files\", \"checkbox\", \"url\", \"email\", \"phone_number\", \"created_time\", \"created_by\", \"last_edited_time\", and \"last_edited_by\".",
		ExampleValue: `"rich_text"`,
	})

	var paginatedPropertyInfo *AdaptiveObject

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Paginated property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information.",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO 良い名前
		paginatedPropertyInfo = b.AddAdaptiveObject("PaginatedPropertyInfo", "type", e.Text)
		paginatedPropertyInfo.AddFields(b.NewField(&Parameter{Property: "id", Description: UNDOCUMENTED}, jen.String()))
		for _, derived := range []string{"title", "rich_text", "relation", "people"} {
			paginatedPropertyInfo.AddAdaptiveFieldWithEmptyStruct(derived, "")
		}
		paginatedPropertyInfo.AddAdaptiveFieldWithType("rollup", UNDOCUMENTED, jen.Id("Rollup"))
	})
	c.ExpectParameter(&Parameter{
		Property:     "object",
		Type:         `"list"`,
		Description:  `Always "list".`,
		ExampleValue: `"list"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         `"property_item"`,
		Description:  `Always "property_item".`,
		ExampleValue: `"property_item"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "results",
		Type:         "list",
		Description:  "List of property_item objects.",
		ExampleValue: `[{"object": "property_item", "id": "vYdV", "type": "relation", "relation": { "id": "535c3fb2-95e6-4b37-a696-036e5eac5cf6"}}... ]`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "property_item",
		Type:         "object",
		Description:  "A property_item object that describes the property.",
		ExampleValue: `{"id": "title", "next_url": null, "type": "title", "title": {}}`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "next_url",
		Type:         "string or null",
		Description:  "The URL the user can request to get the next page of results.",
		ExampleValue: `"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/vYdV?start_cursor=LYxaUO&page_size=25"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		paginatedPropertyInfo.AddFields(b.NewField(e, jen.Id("*").String()))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Title property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Title property value objects contain an array of rich text objects within the title property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		// ドキュメントには "array of rich text" と書いてあるが間違い
		propertyItem.AddAdaptiveFieldWithType("title", e.Text, jen.Id("RichText"))
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Name\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"title\",\n        \"type\": \"title\",\n        \"title\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"The title\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"The title\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"title\",\n      \"next_url\": null,\n      \"type\": \"title\",\n      \"title\": {}\n    }\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Rich Text property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Rich Text property value objects contain an array of rich text objects within the rich_text property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		// ドキュメントには "array of rich text" と書いてあるが間違い
		propertyItem.AddAdaptiveFieldWithType("rich_text", e.Text, jen.Id("RichText"))
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Details\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"Some more text with \",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"Some more text with \",\n          \"href\": null\n        }\n      },\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"fun formatting\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": true,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"fun formatting\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"NVv^\",\n      \"next_url\": null,\n      \"type\": \"rich_text\",\n      \"rich_text\": {}\n    }\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Number property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Number property value objects contain a number within the number property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("number", e.Text, jen.Op("*").Float64())
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Quantity\": {\n    \"object\": \"property_item\",\n    \"id\": \"XpXf\",\n    \"type\": \"number\",\n    \"number\": 1234\n  }\n}",
	}).Output(addTest)
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Select property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Select property value objects contain the following data within the select property:",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("select", e.Text, jen.Op("*").Id("Option"))
		propertyItem.AddAdaptiveFieldWithType("status", "undocumented", jen.Op("*").Id("Option"))
	})
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string (UUIDv4)",
		Description:  "ID of the option.\n\nWhen updating a select property, you can use either name or id.",
		ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "name",
		Type:         "string",
		Description:  "Name of the option as it appears in Notion.\n\nIf the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
		ExampleValue: `"Fruit"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "color",
		Type:         "string (enum)",
		Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
		ExampleValue: `"red"`,
	}) // Optionで共通化
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Option\": {\n    \"object\": \"property_item\",\n    \"id\": \"%7CtzR\",\n    \"type\": \"select\",\n    \"select\": {\n      \"id\": \"64190ec9-e963-47cb-bc37-6a71d6b71206\",\n      \"name\": \"Option 1\",\n      \"color\": \"orange\"\n    }\n  }\n}",
	}).Output(addTest)
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Multi-select property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Multi-select property value objects contain an array of multi-select option values within the multi_select property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("multi_select", e.Text, jen.Index().Id("Option"))
	})
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Multi-select option values",
	})
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string (UUIDv4)",
		Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
		ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "name",
		Type:         "string",
		Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
		ExampleValue: `"Fruit"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     "color",
		Type:         "string (enum)",
		Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
		ExampleValue: `"red"`,
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Tags\": {\n    \"object\": \"property_item\",\n    \"id\": \"z%7D%5C%3C\",\n    \"type\": \"multi_select\",\n    \"multi_select\": [\n      {\n        \"id\": \"91e6959e-7690-4f55-b8dd-d3da9debac45\",\n        \"name\": \"A\",\n        \"color\": \"orange\"\n      },\n      {\n        \"id\": \"2f998e2d-7b1c-485b-ba6b-5e6a815ec8f5\",\n        \"name\": \"B\",\n        \"color\": \"purple\"\n      }\n    ]\n  }\n}",
	}).Output(addTest)

	var propertyItemDate *SimpleObject

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Date property values",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItemDate = propertyItem.AddAdaptiveFieldWithSpecificObject("date", e.Text, b)
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Date property value objects contain the following data within the date property:"})
	c.ExpectParameter(&Parameter{
		Property:     "start",
		Type:         "string (ISO 8601 date and time)",
		Description:  "An ISO 8601 format date, with optional time.",
		ExampleValue: `"2020-12-08T12:00:00Z"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		propertyItemDate.AddFields(b.NewField(e, jen.Id("ISO8601String")))
	})
	c.ExpectParameter(&Parameter{
		Property:     "end",
		Type:         "string (optional, ISO 8601 date and time)",
		Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
		ExampleValue: `"2020-12-08T12:00:00Z"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		propertyItemDate.AddFields(b.NewField(e, jen.Op("*").Id("ISO8601String")))
	})
	c.ExpectParameter(&Parameter{
		Property:     "time_zone",
		Type:         "string (optional, enum)",
		Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
		ExampleValue: `"America/Los_Angeles"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		propertyItemDate.AddFields(b.NewField(e, jen.Op("*").String()))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Shipment Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"i%3Ahj\",\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\",\n      \"end\": null,\n      \"time_zone\": null\n    }\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Formula property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Formula property value objects represent the result of evaluating a formula described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("formula", e.Text, jen.Op("*").Id("Formula"))
	})
	c.ExpectParameter(&Parameter{
		Property:    "type",
		Type:        "string (enum)",
		Description: "The type of the formula result. Possible values are \"string\", \"number\", \"boolean\", and \"date\".",
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "String formula property values"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "String formula property values contain an optional string within the string property."})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Number formula property values"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Number formula property values contain an optional number within the number property."})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Boolean formula property values"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Boolean formula property values contain a boolean within the boolean property."})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Date formula property values"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Date formula property values contain an optional date property value within the date property."})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Formula\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"formula\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Relation property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Relation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("relation", e.Text, jen.Op("*").Id("PageReference")) // 単一
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Project\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"vYdV\",\n      \"next_url\": null,\n      \"type\": \"relation\",\n      \"relation\": {}\n    }\n  }\n}",
	}).Output(addTest)

	var rollup *AdaptiveObject
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Rollup property values",
	}).Output(func(e *Block, b *CodeBuilder) {
		// PropertyValueとPropertyItemで共通の Rollup を使う
		// PropertyValue のArray Rollupにはfunction が無いなど不正確であるため、
		// 比較的ドキュメントが充実しているこちらで作成を行う
		// https://developers.notion.com/reference/property-value-object#rollup-property-values
		// https://developers.notion.com/reference/property-item-object#rollup-property-values
		rollup = b.AddAdaptiveObject("Rollup", "type", "")
		propertyItem.AddAdaptiveFieldWithType("rollup", e.Text, jen.Op("*").Id("Rollup"))
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Rollup property value objects represent the result of evaluating a rollup described in thedatabase's properties. The property is returned as a list object of type property_item with a list of relation items used to computed the rollup under results."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A rollup property item is also returned under the property_type key that describes the rollup aggregation and computed result."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "In order to avoid timeouts, if the rollup has a with a large number of aggregations or properties the endpoint returns a next_cursor value that is used to determinate the aggregation value so far for the subset of relations that have been paginated through."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Once has_more is false, then the final rollup value is returned.  See the Pagination documentation for more information on pagination in the Notion API."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Computing the values of following aggregations are not supported. Instead the endpoint returns a list of property_item objects for the rollup:"})
	c.ExpectBlock(&Block{Kind: "List", Text: "show_unique (Show unique values)unique (Count unique values)median(Median)"})
	c.ExpectParameter(&Parameter{
		Property:    "type",
		Type:        "string (enum)",
		Description: "The type of rollup. Possible values are \"number\", \"date\", \"array\", \"unsupported\" and \"incomplete\".",
	})
	c.ExpectParameter(&Parameter{
		Property:    "function",
		Type:        "string (enum)",
		Description: "Describes the aggregation used. \nPossible values include: count,  count_values,  empty,  not_empty,  unique,  show_unique,  percent_empty,  percent_not_empty,  sum,  average,  median,  min,  max,  range,  earliest_date,  latest_date,  date_range,  checked,  unchecked,  percent_checked,  percent_unchecked,  count_per_group,  percent_per_group,  show_original",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		rollup.AddFields(b.NewField(e, jen.String()))
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Number rollup property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Number rollup property values contain a number within the number property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		rollup.AddAdaptiveFieldWithType("number", e.Text, jen.Op("*").Float64())
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Date rollup property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Date rollup property values contain a date property value within the date property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		rollup.AddAdaptiveFieldWithType("date", e.Text, jen.Op("*").Id("PropertyItemDate"))
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Array rollup property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Array rollup property values contain an array of property_item objects within the results property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		// ドキュメントには array of property_item とあるが、
		// type="rich_text"の場合に入る値などから
		// array of property_value が正しいと判断している
		rollup.AddAdaptiveFieldWithType("array", e.Text, jen.Index().Id("PropertyValue"))
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Incomplete rollup property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Rollups with an aggregation with more than one page of aggregated results will return a rollup object of type \"incomplete\". To obtain the final value paginate through the next values in the rollup using the next_cursor or next_url property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		rollup.AddAdaptiveFieldWithType("incomplete", e.Text, jen.Struct())
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Rollup\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }...\n    ],\t\n    \"next_cursor\": \"1QaTunT5\",\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"y}~p\",\n      \"next_url\": \"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/y%7D~p?start_cursor=1QaTunT5&page_size=25\",\n      \"type\": \"rollup\",\n      \"rollup\": {\n        \"function\": \"sum\",\n        \"type\": \"incomplete\",\n        \"incomplete\": {}\n      }\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		e.Text = strings.Replace(e.Text, "...", "", 1)
		b.AddUnmarshalTest("PropertyItemOrPropertyItemPaginationMap", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "People property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "People property value objects contain an array of user objects within the people property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("people", e.Text, jen.Id("User")) // 単一では？
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Owners\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"people\",\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"285e5768-3fdc-4742-ab9e-125f9050f3b8\",\n        \"name\": \"Example Avo\",\n        \"avatar_url\": null,\n        \"type\": \"person\",\n        \"person\": {\n          \"email\": \"avo@example.com\"\n        }\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO peopleに配列が入っているが、実際はPaginated property valuesになるはずなのでドキュメントの不具合では？確認する
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Files property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("files", e.Text, jen.Index().Id("File"))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Files\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"files\",\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": \"https://website.domain/images/space.png\"\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO サンプルコードの間違いを訂正してテストを通す
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Checkbox property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Checkbox property value objects contain a boolean within the checkbox property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("checkbox", e.Text, jen.Bool())
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Done?\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"checkbox\",\n    \"checkbox\": true\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "URL property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("url", e.Text, jen.Op("*").String())
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Website\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"url\",\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Email property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Email property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("email", e.Text, jen.Op("*").String())
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Shipper's Contact\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"email\",\n    \"email\": \"hello@test.com\"\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Phone number property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Phone number property value objects contain a string within the phone_number property. No structure is enforced.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("phone_number", e.Text, jen.Op("*").String())
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Shipper's No.\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"phone_number\",\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
	}).Output(addTest)

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Created time property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("created_time", e.Text, jen.Id("ISO8601String"))
	})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Created Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"create_time\",\n  \t\"created_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		e.Text = strings.Replace(e.Text, "create_time", "created_time", 1) // ドキュメントの不具合
		b.AddUnmarshalTest("PropertyItemOrPropertyItemPaginationMap", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Created by property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("created_by", e.Text, jen.Op("*").Id("User"))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Created By\": {\n    \"created_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"dsEa\": {\n    \"created_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Last edited time property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("last_edited_time", e.Text, jen.Id("ISO8601String"))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Last Edited Time\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"as0w\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Last edited by property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyItem.AddAdaptiveFieldWithType("last_edited_by", e.Text, jen.Op("*").Id("User"))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Last Edited By\": {\n    \"last_edited_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"as12\": {\n    \"last_edited_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})

	c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
		propertyItem.AddFields(UndocumentedRequestID(b))
	})
}
