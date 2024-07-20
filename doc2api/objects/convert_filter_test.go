package objects_test

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/google/uuid"
	. "github.com/psyark/notion/doc2api/objects"
	"github.com/stoewer/go-strcase"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	extractFilter := func(code string) string {
		t := struct {
			Filter json.RawMessage `json:"filter"`
		}{}
		if err := json.Unmarshal([]byte(code), &t); err != nil {
			panic(err)
		}
		return string(t.Filter)
	}

	c := converter.FetchDocument("https://developers.notion.com/reference/post-database-query-filter")

	var filter *SimpleObject

	// Type-specific filter conditions に従い、
	// フィルターの型ごとのフィールドとペイロードオブジェクトを作成します
	// （このような使い方はFilter特有のため、CodeBuilderに同様の機能を作りません）
	newSpecificObject := func(b *CodeBuilder, name string, comment string) *SimpleObject {
		objName := "Filter" + strcase.UpperCamelCase(name)
		filter.AddFields(b.NewField(&Parameter{Property: name, Description: comment}, jen.Op("*").Id(objName), OmitEmpty))
		return b.AddSimpleObject(objName, comment)
	}

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "When you query a database, you can send a filter object in the body of the request that limits the returned entries based on the specified criteria."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "For example, the below query limits the response to entries where the \"Task completed\"  checkbox property value is true:"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl -X POST 'https://api.notion.com/v1/databases/897e5a76ae524b489fdfe71f5945d1af/query' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H 'Notion-Version: 2022-06-28' \\\n  -H \"Content-Type: application/json\" \\\n--data '{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n        \"equals\": true\n   }\n  }\n}'\n"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Here is the same query using the Notion SDK for JavaScript:"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "const { Client } = require('@notionhq/client');\n\nconst notion = new Client({ auth: process.env.NOTION_API_KEY });\n// replace with your own database ID\nconst databaseId = 'd9824bdc-8445-4327-be8b-5b47500af6ce';\n\nconst filteredRows = async () => {\n\tconst response = await notion.databases.query({\n\t  database_id: databaseId,\n\t  filter: {\n\t    property: \"Task completed\",\n\t    checkbox: {\n\t      equals: true\n\t    }\n\t  },\n\t});\n  return response;\n}\n\n"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Filters can be chained with the and and or keys so that multiple filters are applied at the same time. (See Query a database for additional examples.)"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "{\n  \"and\": [\n    {\n      \"property\": \"Done\",\n      \"checkbox\": {\n        \"equals\": true\n      }\n    }, \n    {\n      \"or\": [\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"A\"\n        },\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"B\"\n        }\n      ]\n    }\n  ]\n}\n"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "If no filter is provided, all the pages in the database will be returned with pagination."})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "The filter object",
	}).Output(func(e *Block, b *CodeBuilder) {
		filter = b.AddSimpleObject("Filter", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Each filter object contains the following fields:"})

	c.ExpectParameter(&Parameter{
		Property:     "property",
		Type:         "string",
		Description:  "The name of the property as it appears in the database, or the property ID.",
		ExampleValue: `"Task completed"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		filter.AddFields(b.NewField(e, jen.String(), OmitEmpty))
	})
	c.ExpectParameter(&Parameter{
		Property:     "checkbox  \ndate  \nfiles  \nformula  \nmulti_select  \nnumber  \npeople  \nphone_number  \nrelation  \nrich_text  \nselect  \nstatus  \ntimestamp  \nID",
		Type:         "object",
		Description:  "The type-specific filter condition for the query. Only types listed in the Field column of this table are supported.  \n  \nRefer to type-specific filter conditions for details on corresponding object values.",
		ExampleValue: "\"checkbox\": {\n  \"equals\": true\n}",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"equals\": true\n    }\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "👍The filter object mimics the database filter option in the Notion UI.",
	})
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Type-specific filter conditions",
	})

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Checkbox",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "checkbox", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "equals",
			Type:         "boolean",
			Description:  "Whether a checkbox property value matches the provided value exactly.  \n  \nReturns or excludes all database entries with an exact value match.",
			ExampleValue: "false",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_equal",
			Type:         "boolean",
			Description:  "Whether a checkbox property value differs from the provided value.  \n  \nReturns or excludes all database entries with a difference in values.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"does_not_equal\": true\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Date",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "date", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Blockquote",
			Text: "📘For the after, before, equals, on_or_before, and on_or_after fields, if a date string with a time is provided, then the comparison is done with millisecond precision.If no timezone is provided, then the timezone defaults to UTC.",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A date filter condition can be used to limit date property value types and the timestamp property types created_time and last_edited_time.",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The condition contains the below fields:",
		})
		c.ExpectParameter(&Parameter{
			Property:     "after",
			Type:         "string (ISO 8601 date)",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is after the provided date.",
			ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Id("ISO8601String"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "before",
			Type:         "string (ISO 8601 date)",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is before the provided date.",
			ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Id("ISO8601String"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "equals",
			Type:         "string (ISO 8601 date)",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is the provided date.",
			ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Id("ISO8601String"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value contains no data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is not empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "next_month",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is within the next month.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "next_week",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is within the next week.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "next_year",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is within the next year.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "on_or_after",
			Type:         "string (ISO 8601 date)",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is on or after the provided date.",
			ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Id("ISO8601String"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "on_or_before",
			Type:         "string (ISO 8601 date)",
			Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is on or before the provided date.",
			ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Id("ISO8601String"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "past_month",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is within the past month.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "past_week",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is within the past week.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "past_year",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is within the past year.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "this_week",
			Type:         "object (empty)",
			Description:  "A filter that limits the results to database entries where the date property value is this week.",
			ExampleValue: "{}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Struct(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Due date\",\n    \"date\": {\n      \"on_or_after\": \"2023-02-08\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Files",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "files", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the files property value does not contain any data.  \n  \nReturns all database entries with an empty files property value.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the files property value contains data.  \n  \nReturns all entries with a populated files property value.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Blueprint\",\n    \"files\": {\n      \"is_not_empty\": true\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var filterFormula *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Formula",
		}).Output(func(e *Block, b *CodeBuilder) {
			filterFormula = newSpecificObject(b, "formula", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The primary field of the formula filter condition object matches the type of the formula’s result. For example, to filter a formula property that computes a checkbox, use a formula filter condition object with a checkbox field containing a checkbox filter condition as its value.",
		}).Output(func(e *Block, b *CodeBuilder) {
			filterFormula.AddComment(e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "checkbox",
			Type:         "object",
			Description:  "A checkbox filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
			ExampleValue: "Refer to the checkbox filter condition.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterFormula.AddFields(b.NewField(e, jen.Op("*").Id("FilterCheckbox"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "date",
			Type:         "object",
			Description:  "A date filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
			ExampleValue: "Refer to the date filter condition.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterFormula.AddFields(b.NewField(e, jen.Op("*").Id("FilterDate"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "number",
			Type:         "object",
			Description:  "A number filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
			ExampleValue: "Refer to the number filter condition.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterFormula.AddFields(b.NewField(e, jen.Op("*").Id("FilterNumber"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "string",
			Type:         "object",
			Description:  "A rich text filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
			ExampleValue: "Refer to the rich text filter condition.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterFormula.AddFields(b.NewField(e, jen.Op("*").Id("FilterRichText"), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"One month deadline\",\n    \"formula\": {\n      \"date\":{\n          \"after\": \"2021-05-10\"\n      }\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Multi-select",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "multi_select", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "contains",
			Type:         "string",
			Description:  "The value to compare the multi-select property value against.  \n  \nReturns database entries where the multi-select value matches the provided string.",
			ExampleValue: `"Marketing"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_contain",
			Type:         "string",
			Description:  "The value to compare the multi-select property value against.  \n  \nReturns database entries where the multi-select value does not match the provided string.",
			ExampleValue: `"Engineering"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the multi-select property value is empty.  \n  \nReturns database entries where the multi-select value does not contain any data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the multi-select property value is not empty.  \n  \nReturns database entries where the multi-select value does contains data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Programming language\",\n    \"multi_select\": {\n      \"contains\": \"TypeScript\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Number",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "number", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_equal",
			Type:         "number",
			Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value differs from the provided number.",
			ExampleValue: "42",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Float64(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "equals",
			Type:         "number",
			Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is the same as the provided number.",
			ExampleValue: "42",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Float64(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "greater_than",
			Type:         "number",
			Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value exceeds the provided number.",
			ExampleValue: "42",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Float64(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "greater_than_or_equal_to",
			Type:         "number",
			Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is equal to or exceeds the provided number.",
			ExampleValue: "42",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Float64(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the number property value is empty.  \n  \nReturns database entries where the number property value does not contain any data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the number property value is not empty.  \n  \nReturns database entries where the number property value contains data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "less_than",
			Type:         "number",
			Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is less than the provided number.",
			ExampleValue: "42",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Float64(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "less_than_or_equal_to",
			Type:         "number",
			Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is equal to or is less than the provided number.",
			ExampleValue: "42",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Op("*").Float64(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Estimated working days\",\n    \"number\": {\n      \"less_than_or_equal_to\": 5\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "People",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "people", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "You can apply a people filter condition to people, created_by, and last_edited_by database property types.",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The people filter condition contains the following fields:",
		})
		c.ExpectParameter(&Parameter{
			Property:     "contains",
			Type:         "string (UUIDv4)",
			Description:  "The value to compare the people property value against.  \n  \nReturns database entries where the people property value contains the provided string.",
			ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, NullUUID, OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_contain",
			Type:         "string (UUIDv4)",
			Description:  "The value to compare the people property value against.  \n  \nReturns database entries where the people property value does not contain the provided string.",
			ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, NullUUID, OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the people property value does not contain any data.  \n  \nReturns database entries where the people property value does not contain any data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the people property value contains data.  \n  \nReturns database entries where the people property value is not empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Last edited by\",\n    \"people\": {\n      \"contains\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Relation",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "relation", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "contains",
			Type:         "string (UUIDv4)",
			Description:  "The value to compare the relation property value against.  \n  \nReturns database entries where the relation property value contains the provided string.",
			ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, NullUUID, OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_contain",
			Type:         "string (UUIDv4)",
			Description:  "The value to compare the relation property value against.  \n  \nReturns entries where the relation property value does not contain the provided string.",
			ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, NullUUID, OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the relation property value does not contain data.  \n  \nReturns database entries where the relation property value does not contain any data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the relation property value contains data.  \n  \nReturns database entries where the property value is not empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"✔️ Task List\",\n    \"relation\": {\n      \"contains\": \"0c1f7cb280904f18924ed92965055e32\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			text := regexp.MustCompile(`\w{32}`).ReplaceAllStringFunc(e.Text, func(s string) string {
				return uuid.MustParse(s).String()
			})
			converter.AddUnmarshalTest("Filter", extractFilter(text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Rich text",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "rich_text", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "contains",
			Type:         "string",
			Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that includes the provided string.",
			ExampleValue: `"Moved to Q2"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_contain",
			Type:         "string",
			Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that does not include the provided string.",
			ExampleValue: `"Moved to Q2"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_equal",
			Type:         "string",
			Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that does not match the provided string.",
			ExampleValue: `"Moved to Q2"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "ends_with",
			Type:         "string",
			Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that ends with the provided string.",
			ExampleValue: `"Q2"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "equals",
			Type:         "string",
			Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that matches the provided string.",
			ExampleValue: `"Moved to Q2"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the text property value does not contain any data.  \n  \nReturns database entries with a text property value that is empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the text property value contains any data.  \n  \nReturns database entries with a text property value that contains data.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "starts_with",
			Type:         "string",
			Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that starts with the provided string.",
			ExampleValue: `"Moved"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Description\",\n    \"rich_text\": {\n      \"contains\": \"cross-team\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var filterRollup *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Rollup",
		}).Output(func(e *Block, b *CodeBuilder) {
			filterRollup = newSpecificObject(b, "rollup", e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A rollup database property can evaluate to an array, date, or number value. The filter condition for the rollup property contains a rollup key and a corresponding object value that depends on the computed value type.",
		}).Output(func(e *Block, b *CodeBuilder) {
			filterRollup.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Filter conditions for array rollup values",
		})
		c.ExpectParameter(&Parameter{
			Property:     "any",
			Type:         "object",
			Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type.  \n  \nReturns database entries where the rollup property value matches the provided criteria.",
			ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterRollup.AddFields(b.NewField(e, jen.Op("*").Id("Filter"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "every",
			Type:         "object",
			Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type.  \n  \nReturns database entries where every rollup property value matches the provided criteria.",
			ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterRollup.AddFields(b.NewField(e, jen.Op("*").Id("Filter"), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "none",
			Type:         "object",
			Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type.  \n  \nReturns database entries where no rollup property value matches the provided criteria.",
			ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterRollup.AddFields(b.NewField(e, jen.Op("*").Id("Filter"), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Related tasks\",\n    \"rollup\": {\n      \"any\": {\n        \"rich_text\": {\n          \"contains\": \"Migrate database\"\n        }\n      }\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Filter conditions for date rollup values"})
		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A rollup value is stored as a date only if the \"Earliest date\", \"Latest date\", or \"Date range\" computation is selected for the property in the Notion UI."})
		c.ExpectParameter(&Parameter{
			Property:     "date",
			Type:         "object",
			Description:  "A date filter condition to compare the rollup value against.  \n  \nReturns database entries where the rollup value matches the provided condition.",
			ExampleValue: "Refer to the date filter condition.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterRollup.AddFields(b.NewField(e, jen.Op("*").Id("FilterDate"), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Parent project due date\",\n    \"rollup\": {\n      \"date\": {\n        \"on_or_before\": \"2023-02-08\"\n      }\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Filter conditions for number rollup values",
		})
		c.ExpectParameter(&Parameter{
			Property:     "number",
			Type:         "object",
			Description:  "A number filter condition to compare the rollup value against.  \n  \nReturns database entries where the rollup value matches the provided condition.",
			ExampleValue: "Refer to the number filter condition.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			filterRollup.AddFields(b.NewField(e, jen.Op("*").Id("FilterNumber"), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Total estimated working days\",\n    \"rollup\": {\n      \"number\": {\n        \"does_not_equal\": 42\n      }\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Select",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "select", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "equals",
			Type:         "string",
			Description:  "The string to compare the select property value against.  \n  \nReturns database entries where the select property value matches the provided string.",
			ExampleValue: `"This week"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_equal",
			Type:         "string",
			Description:  "The string to compare the select property value against.  \n  \nReturns database entries where the select property value does not match the provided string.",
			ExampleValue: `"Backlog"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the select property value does not contain data.  \n  \nReturns database entries where the select property value is empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the select property value contains data.  \n  \nReturns database entries where the select property value is not empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Frontend framework\",\n    \"select\": {\n      \"equals\": \"React\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	{
		var specificObject *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Status",
		}).Output(func(e *Block, b *CodeBuilder) {
			specificObject = newSpecificObject(b, "status", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "equals",
			Type:         "string",
			Description:  "The string to compare the status property value against.  \n  \nReturns database entries where the status property value matches the provided string.",
			ExampleValue: `"This week"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "does_not_equal",
			Type:         "string",
			Description:  "The string to compare the status property value against.  \n  \nReturns database entries where the status property value does not match the provided string.",
			ExampleValue: `"Backlog"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_empty",
			Type:         "true",
			Description:  "Whether the status property value does not contain data.  \n  \nReturns database entries where the status property value is empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "is_not_empty",
			Type:         "true",
			Description:  "Whether the status property value contains data.  \n  \nReturns database entries where the status property value is not empty.",
			ExampleValue: "true",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			specificObject.AddFields(b.NewField(e, jen.Bool(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "{\n  \"filter\": {\n    \"property\": \"Project status\",\n    \"status\": {\n      \"equals\": \"Not started\"\n    }\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
		})
	}

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Timestamp",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Use a timestamp filter condition to filter results based on created_time or last_edited_time values.",
	})
	c.ExpectParameter(&Parameter{
		Property:     "timestamp",
		Type:         "created_time last_edited_time",
		Description:  "A constant string representing the type of timestamp to use as a filter.",
		ExampleValue: `"created_time"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		filter.AddFields(b.NewField(e, jen.String(), OmitEmpty))
	})
	c.ExpectParameter(&Parameter{
		Property:     "created_time  \nlast_edited_time",
		Type:         "object",
		Description:  "A date filter condition used to filter the specified timestamp.",
		ExampleValue: "Refer to the date filter condition.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		for _, prop := range strings.Split(e.Property, "\n") {
			e.Property = strings.TrimSpace(prop)
			filter.AddFields(b.NewField(e, jen.Op("*").Id("FilterDate"), OmitEmpty))
		}
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"filter\": {\n    \"timestamp\": \"created_time\",\n    \"created_time\": {\n      \"on_or_before\": \"2022-10-13\"\n    }\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "🚧The timestamp filter condition does not require a property name. The API throws an error if you provide one.",
	}).Output(func(e *Block, b *CodeBuilder) {
		filter.AddComment(e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "ID",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Use a timestamp filter condition to filter results based on the unique_id value.",
	})
	c.ExpectParameter(&Parameter{
		Property:     "does_not_equal",
		Type:         "number",
		Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value differs from the provided value.",
		ExampleValue: "42",
	})
	c.ExpectParameter(&Parameter{
		Property:     "equals",
		Type:         "number",
		Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is the same as the provided value.",
		ExampleValue: "42",
	})
	c.ExpectParameter(&Parameter{
		Property:     "greater_than",
		Type:         "number",
		Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value exceeds the provided value.",
		ExampleValue: "42",
	})
	c.ExpectParameter(&Parameter{
		Property:     "greater_than_or_equal_to",
		Type:         "number",
		Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is equal to or exceeds the provided value.",
		ExampleValue: "42",
	})
	c.ExpectParameter(&Parameter{
		Property:     "less_than",
		Type:         "number",
		Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is less than the provided value.",
		ExampleValue: "42",
	})
	c.ExpectParameter(&Parameter{
		Property:     "less_than_or_equal_to",
		Type:         "number",
		Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is equal to or is less than the provided value.",
		ExampleValue: "42",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"filter\": {\n    \"and\": [\n      {\n        \"property\": \"ID\",\n        \"unique_id\": {\n          \"greater_than\": 1\n        }\n      },\n      {\n        \"property\": \"ID\",\n        \"unique_id\": {\n          \"less_than\": 3\n        }\n      }\n    ]\n  }\n}\n",
	})
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Compound filter conditions",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "You can use a compound filter condition to limit the results of a database query based on multiple conditions. This mimics filter chaining in the Notion UI.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The above filters in the Notion UI are equivalent to the following compound filter condition via the API:",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"and\": [\n    {\n      \"property\": \"Done\",\n      \"checkbox\": {\n        \"equals\": true\n      }\n    }, \n    {\n      \"or\": [\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"A\"\n        },\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"B\"\n        }\n      ]\n    }\n  ]\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		// TODO サンプルコードがめちゃくちゃ
		// converter.AddUnmarshalTest("Filter", e.Code)
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A compound filter condition contains an and or or key with a value that is an array of filter objects or nested compound filter objects. Nesting is supported up to two levels deep.",
	})
	c.ExpectParameter(&Parameter{
		Property:     "and",
		Type:         "array",
		Description:  "An array of filter objects or compound filter conditions.  \n  \nReturns database entries that match all of the provided filter conditions.",
		ExampleValue: "Refer to the examples below.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		filter.AddFields(b.NewField(e, jen.Index().Id("Filter"), OmitEmpty))
	})
	c.ExpectParameter(&Parameter{
		Property:     "or",
		Type:         "array",
		Description:  "An array of filter objects or compound filter conditions.  \n  \nReturns database entries that match any of the provided filter conditions",
		ExampleValue: "Refer to the examples below.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		filter.AddFields(b.NewField(e, jen.Index().Id("Filter"), OmitEmpty))
	})
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Example compound filter conditions",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"filter\": {\n    \"and\": [\n      {\n        \"property\": \"Complete\",\n        \"checkbox\": {\n          \"equals\": true\n        }\n      },\n      {\n        \"property\": \"Working days\",\n        \"number\": {\n          \"greater_than\": 10\n        }\n      }\n    ]\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"filter\": {\n    \"or\": [\n      {\n        \"property\": \"Description\",\n        \"rich_text\": {\n          \"contains\": \"2023\"\n        }\n      },\n      {\n        \"and\": [\n          {\n            \"property\": \"Department\",\n            \"select\": {\n              \"equals\": \"Engineering\"\n            }\n          },\n          {\n            \"property\": \"Priority goal\",\n            \"checkbox\": {\n              \"equals\": true\n            }\n          }\n        ]\n      }\n    ]\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Filter", extractFilter(e.Text))
	})
}
