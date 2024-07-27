package objects_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestPropertyValue(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/property-value-object")

	var propertyValue, formula *UnionStruct

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue = b.AddUnionStruct("PropertyValue", "type", e.Text, DiscriminatorOmitEmpty())
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "Any property value that has other pages in its value will only use the first 25 page references. Use the Retrieve a page property endpoint to paginate through the full value.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddComment(e.Text)
	})
	/* All property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "All property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Each page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddComment(e.Text)
	})
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string",
		Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
		ExampleValue: `"f%5C%5C%3Ap"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		propertyValue.AddFields(b.NewField(e, jen.String(), OmitEmpty)) // Rollup内でIDが無い場合がある
	})
	c.ExpectParameter(&Parameter{
		Property:     "type (optional)",
		Type:         "string (enum)",
		Description:  "Type of the property. Possible values are \"rich_text\", \"number\", \"select\", \"multi_select\", \"status\", \"date\", \"formula\", \"relation\", \"rollup\", \"title\", \"people\", \"files\", \"checkbox\", \"url\", \"email\", \"phone_number\", \"created_time\", \"created_by\", \"last_edited_time\", and \"last_edited_by\".",
		ExampleValue: `"rich_text"`,
	})
	/* Title property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Title property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Title property value objects contain an array of rich text objects within the title property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("title", e.Text, WithType(jen.Id("RichTextArray")))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Name\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"title\": {\n    \"title\": [\n      {\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "The Retrieve a page endpoint returns a maximum of 25 inline page or person references for a title property. If a title property includes more than 25 references, then you can use the\u00a0Retrieve a page property endpoint for the specific title property to get its complete list of references.",
	})
	/* Rich Text property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Rich Text property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Rich Text property value objects contain an array of rich text objects within the rich_text property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("rich_text", e.Text, WithType(jen.Id("RichTextArray")))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Details\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"D[X|\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "The Retrieve a page endpoint returns a maximum of 25 populated inline page or person references for a rich_text property. If a rich_text property includes more than 25 references, then you can use the Retrieve a page property endpoint for the specific rich_text property to get its complete list of references.",
	})
	/* Number property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Number property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Number property value objects contain a number within the number property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("number", e.Text, WithType(jen.Op("*").Float64()))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Quantity\": {\n    \"number\": 1234\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"pg@s\": {\n    \"number\": 1234\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* Select property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Select property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Select property value objects contain the following data within the select property:",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("select", e.Text, WithType(jen.Op("*").Id("Option"))) // may null
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
		Text: "{\n  \"Option\": {\n    \"select\": {\n      \"name\": \"Option 1\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"XMqQ\": {\n    \"select\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO
	})
	/* Status property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Status property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Status property value objects contain the following data within the status property:",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("status", e.Text, WithType(jen.Op("*").Id("Option")))
	})
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string (UUIDv4)",
		Description:  "ID of the option.",
		ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "name",
		Type:         "string",
		Description:  "Name of the option as it appears in Notion.",
		ExampleValue: `"In progress"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "color",
		Type:         "string (enum)",
		Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
		ExampleValue: `"red"`,
	}) // Optionで共通化
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Status\": {\n    \"status\": {\n      \"name\": \"In progress\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"XMqQ\": {\n    \"status\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO
	})
	/* Multi-select property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Multi-select property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Multi-select property value objects contain an array of multi-select option values within the multi_select property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("multi_select", e.Text, WithType(jen.Index().Id("Option")))
	})
	/* Multi-select option values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Multi-select option values",
	})
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string (UUIDv4)",
		Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
		ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "name",
		Type:         "string",
		Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
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
		Text: "{\n  \"Tags\": {\n    \"multi_select\": [\n      {\n        \"name\": \"B\"\n      },\n      {\n        \"name\": \"C\"\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"uyn@\": {\n    \"multi_select\": [\n      {\n        \"id\": \"3d3ca089-f964-4831-a8a2-0c6d746f4162\"\n      },\n      {\n        \"id\": \"1919ba02-1bf3-4e73-8832-8c0020f17363\"\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO
	})
	/* Date property values */
	var date *SimpleObject

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Date property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Date property value objects contain the following data within the date property:",
	}).Output(func(e *Block, b *CodeBuilder) {
		date = propertyValue.AddPayloadField("date", e.Text, WithPayloadObject(b))
	})
	c.ExpectParameter(&Parameter{
		Property:     "start",
		Type:         "string (ISO 8601 date and time)",
		Description:  "An ISO 8601 format date, with optional time.",
		ExampleValue: `"2020-12-08T12:00:00Z"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		date.AddFields(b.NewField(e, jen.Id("ISO8601String")))
	})
	c.ExpectParameter(&Parameter{
		Property:     "end",
		Type:         "string (optional, ISO 8601 date and time)",
		Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
		ExampleValue: `"2020-12-08T12:00:00Z"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		date.AddFields(b.NewField(e, jen.Op("*").Id("ISO8601String"))) // APIでnullがあるのでomitemptyしない
	})
	c.ExpectParameter(&Parameter{
		Property:     "time_zone",
		Type:         "string (optional, enum)",
		Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
		ExampleValue: `"America/Los_Angeles"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		date.AddFields(b.NewField(e, jen.Op("*").String())) // APIでnullがあるのでomitemptyしない
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Shipment Time\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"CbFP\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Preparation Range\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"\\\\rm}\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Delivery Time\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"DgRt\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* Formula property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Formula property values",
	}).Output(func(e *Block, b *CodeBuilder) {
		formula = b.AddUnionStruct("Formula", "type", e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Formula property value objects represent the result of evaluating a formula described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a formula cannot be updated directly.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("formula", e.Text, WithType(jen.Op("*").Id("Formula")))
	})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "Formulas returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value."})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Formula\": {\n    \"id\": \"1lab\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectParameter(&Parameter{
		Property: "type",
		Type:     "string (enum)",
	})
	/* String formula property values */
	c.ExpectBlock(&Block{Kind: "Heading", Text: "String formula property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "String formula property values contain an optional string within the string property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		formula.AddPayloadField("string", e.Text, WithType(jen.Op("*").String()))
	})
	/* Number formula property values */
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Number formula property values"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Number formula property values contain an optional number within the number property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		formula.AddPayloadField("number", e.Text, WithType(jen.Float64()))
	})
	/* Boolean formula property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Boolean formula property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Boolean formula property values contain a boolean within the boolean property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		formula.AddPayloadField("boolean", e.Text, WithType(jen.Bool()))
	})
	/* Date formula property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Date formula property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Date formula property values contain an optional date property value within the date property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		formula.AddPayloadField("date", e.Text, WithType(jen.Id("PropertyValueDate")))
	})
	/* Relation property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Relation property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Relation property value objects contain an array of page references within the\u00a0relation property. A page reference is an object with an id key and a string value (UUIDv4) corresponding to a page ID in another database.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("relation", e.Text, WithType(jen.Index().Id("PageReference")))
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A relation includes a has_more property in the Retrieve a page endpoint response object. The endpoint returns a maximum of 25 page references for a relation. If a relation has more than 25 references, then the has_more value for the relation in the response object is true. If a relation doesn’t exceed the limit, then has_more is false.",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO 他と似たようにする
		propertyValue.AddFields(b.NewField(&Parameter{Property: "has_more", Description: e.Text}, jen.Bool(), DiscriminatorValue("relation")))
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Note that updating a relation property value with an empty array will clear the list.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Project\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"mODt\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* Rollup property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Rollup property values",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO RollupはPropertyItemで、FormulaはPropertyValueでtranslateしているのを統一
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Rollup property value objects represent the result of evaluating a rollup described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a rollup cannot be updated directly.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("rollup", e.Text, WithType(jen.Op("*").Id("Rollup")))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Rollup\": {\n    \"id\": \"aJ3l\",\n    \"rollup\": {\n      \"type\": \"number\",\n      \"number\": 1234,\n      \"function\": \"sum\"\n    }\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "Rollups returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
	})
	/* String rollup property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "String rollup property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "String rollup property values contain an optional string within the string property.",
	})
	/* Number rollup property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Number rollup property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Number rollup property values contain a number within the number property.",
	})
	/* Date rollup property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Date rollup property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Date rollup property values contain a date property value within the date property.",
	})
	/* Array rollup property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Array rollup property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Array rollup property values contain an array of number, date, or string objects within the results property.",
	})
	/* People property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "People property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "People property value objects contain an array of user objects within the people property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("people", e.Text, WithType(jen.Index().Id("User")))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "The Retrieve a page endpoint can’t be guaranteed to return more than 25 people per people page property. If a people page property includes more than 25 people, then you can use the\u00a0Retrieve a page property endpoint for the specific people property to get a complete list of people.",
	})
	/* Files property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Files property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("files", e.Text, WithType(jen.Index().Id("File")))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Files\": {\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": {\n          \t\"url\": \"https://website.domain/images/space.png\"\n        }\n      }\n    ]\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "Although we do not support uploading files, if you pass a file object containing a file hosted by Notion, it will remain one of the files. To remove any file, just do not pass it in the update response.",
	})
	/* Checkbox property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Checkbox property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Checkbox property value objects contain a boolean within the checkbox property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("checkbox", e.Text, WithType(jen.Bool())) // never null | undefined
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Done?\": {\n    \"checkbox\": true\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"RirO\": {\n    \"checkbox\": true\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* URL property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "URL property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("url", e.Text, WithType(jen.Op("*").String()))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Website\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"<tdn\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* Email property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Email property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Email property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("email", e.Text, WithType(jen.Op("*").String()))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Shipper's Contact\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"}=RV\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* Phone number property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Phone number property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Phone number property value objects contain a string within the phone_number property. No structure is enforced.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("phone_number", e.Text, WithType(jen.Op("*").String()))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"Shipper's No.\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"_A<p\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO テスト通す
	})
	/* Created time property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Created time property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of created_time cannot be updated. See the Property Item Object to see how these values are returned.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("created_time", e.Text, WithType(jen.Id("ISO8601String")))
	})
	/* Created by property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Created by property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page. The value of created_by cannot be updated. See the Property Item Object to see how these values are returned.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("created_by", e.Text, WithType(jen.Id("User")))
	})
	/* Last edited time property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Last edited time property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of last_edited_time cannot be updated. See the Property Item Object to see how these values are returned.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("last_edited_time", e.Text, WithType(jen.Id("ISO8601String")))
	})
	/* Last edited by property values */
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Last edited by property values",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page. The value of last_edited_by cannot be updated. See the Property Item Object to see how these values are returned.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertyValue.AddPayloadField("last_edited_by", e.Text, WithType(jen.Id("User")))
	})

	c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
		payload := propertyValue.AddPayloadField("unique_id", UNDOCUMENTED, WithPayloadObject(b))
		payload.AddFields(b.NewField(&Parameter{Property: "number", Description: UNDOCUMENTED}, jen.Int()))
		payload.AddFields(b.NewField(&Parameter{Property: "prefix", Description: UNDOCUMENTED}, jen.String()))

		propertyValue.AddPayloadField("button", UNDOCUMENTED, WithEmptyStructRef())
	})
}
