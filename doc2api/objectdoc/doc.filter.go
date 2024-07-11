package objectdoc

import (
	"encoding/json"
	"regexp"

	"github.com/dave/jennifer/jen"
	"github.com/google/uuid"
)

func extractFilter(code string) string {
	t := struct {
		Filter json.RawMessage `json:"filter"`
	}{}
	if err := json.Unmarshal([]byte(code), &t); err != nil {
		panic(err)
	}
	return string(t.Filter)
}

func init() {
	var filter, rollupFilter *adaptiveObject

	registerTranslator(
		"https://developers.notion.com/reference/post-database-query-filter",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "When you query a database, you can send a filter object in the body of the request that limits the returned entries based on the specified criteria.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "For example, the below query limits the response to entries where the \"Task completed\"  checkbox property value is true:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "curl -X POST 'https://api.notion.com/v1/databases/897e5a76ae524b489fdfe71f5945d1af/query' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H 'Notion-Version: 2022-06-28' \\\n  -H \"Content-Type: application/json\" \\\n--data '{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n        \"equals\": true\n   }\n  }\n}'\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Here is the same query using the Notion SDK for JavaScript:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "const { Client } = require('@notionhq/client');\n\nconst notion = new Client({ auth: process.env.NOTION_API_KEY });\n// replace with your own database ID\nconst databaseId = 'd9824bdc-8445-4327-be8b-5b47500af6ce';\n\nconst filteredRows = async () => {\n\tconst response = await notion.databases.query({\n\t  database_id: databaseId,\n\t  filter: {\n\t    property: \"Task completed\",\n\t    checkbox: {\n\t      equals: true\n\t    }\n\t  },\n\t});\n  return response;\n}\n\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Filters can be chained with the and and or keys so that multiple filters are applied at the same time. (See Query a database for additional examples.)",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"and\": [\n    {\n      \"property\": \"Done\",\n      \"checkbox\": {\n        \"equals\": true\n      }\n    }, \n    {\n      \"or\": [\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"A\"\n        },\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"B\"\n        }\n      ]\n    }\n  ]\n}\n",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If no filter is provided, all the pages in the database will be returned with pagination.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* The filter object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "The filter object",
			}, func(e blockElement) {
				filter = b.addAdaptiveObject("Filter", "", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Each filter object contains the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "property",
				Type:         "string",
				Description:  "The name of the property as it appears in the database, or the property ID.",
				ExampleValue: `"Task completed"`,
			}, func(e parameterElement) {
				filter.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The type-specific filter condition for the query. Only types listed in the Field column of this table are supported.  \n  \nRefer to type-specific filter conditions for details on corresponding object values.",
				ExampleValue: "\"checkbox\": {\n  \"equals\": true\n}",
				Property:     "checkbox  \ndate  \nfiles  \nformula  \nmulti_select  \nnumber  \npeople  \nphone_number  \nrelation  \nrich_text  \nselect  \nstatus  \ntimestamp  \nID",
				Type:         "object",
			}, func(e parameterElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"equals\": true\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "👍The filter object mimics the database filter option in the Notion UI.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Type-specific filter conditions */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Type-specific filter conditions",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Checkbox */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Checkbox",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("checkbox", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether a checkbox property value matches the provided value exactly.  \n  \nReturns or excludes all database entries with an exact value match.",
				ExampleValue: "false",
				Property:     "equals",
				Type:         "boolean",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether a checkbox property value differs from the provided value.  \n  \nReturns or excludes all database entries with a difference in values.",
				ExampleValue: "true",
				Property:     "does_not_equal",
				Type:         "boolean",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"does_not_equal\": true\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Date */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("date", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘For the after, before, equals, on_or_before, and on_or_after fields, if a date string with a time is provided, then the comparison is done with millisecond precision.If no timezone is provided, then the timezone defaults to UTC.",
			}, func(e blockElement) {
				specificObject.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A date filter condition can be used to limit date property value types and the timestamp property types created_time and last_edited_time.",
			}, func(e blockElement) {
				specificObject.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The condition contains the below fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "after",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is after the provided date.",
				ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "before",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is before the provided date.",
				ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "equals",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is the provided date.",
				ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_empty",
				Type:         "true",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value contains no data.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_not_empty",
				Type:         "true",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is not empty.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "next_month",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is within the next month.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "next_week",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is within the next week.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "next_year",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is within the next year.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "on_or_after",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is on or after the provided date.",
				ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "on_or_before",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.  \n  \nReturns database entries where the date property value is on or before the provided date.",
				ExampleValue: "\"2021-05-10\"  \n  \n\"2021-05-10T12:00:00\"  \n  \n\"2021-10-15T12:00:00-07:00\"",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "past_month",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is within the past month.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "past_week",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is within the past week.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "past_year",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is within the past year.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "this_week",
				Type:         "object (empty)",
				Description:  "A filter that limits the results to database entries where the date property value is this week.",
				ExampleValue: "{}",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Due date\",\n    \"date\": {\n      \"on_or_after\": \"2023-02-08\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Files */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Files",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("files", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_empty",
				Type:         "true",
				Description:  "Whether the files property value does not contain any data.  \n  \nReturns all database entries with an empty files property value.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_not_empty",
				Type:         "true",
				Description:  "Whether the files property value contains data.  \n  \nReturns all entries with a populated files property value.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Blueprint\",\n    \"files\": {\n      \"is_not_empty\": true\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Formula */ {
			var formulaFilter *adaptiveObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Formula",
			}, func(e blockElement) {
				formulaFilter = b.addAdaptiveObject("FormulaFilter", "", e.Text)
				filter.addAdaptiveFieldWithType("formula", e.Text, jen.Op("*").Id("FormulaFilter"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The primary field of the formula filter condition object matches the type of the formula’s result. For example, to filter a formula property that computes a checkbox, use a formula filter condition object with a checkbox field containing a checkbox filter condition as its value.",
			}, func(e blockElement) {
				formulaFilter.addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Description:  "A checkbox filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the checkbox filter condition.",
				Property:     "checkbox",
				Type:         "object",
			}, func(e parameterElement) {
				formulaFilter.addAdaptiveFieldWithType("checkbox", e.Description, jen.Op("*").Id("FilterCheckbox"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "A date filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the date filter condition.",
				Property:     "date",
				Type:         "object",
			}, func(e parameterElement) {
				formulaFilter.addAdaptiveFieldWithType("date", e.Description, jen.Op("*").Id("FilterDate"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "A number filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the number filter condition.",
				Property:     "number",
				Type:         "object",
			}, func(e parameterElement) {
				formulaFilter.addAdaptiveFieldWithType("number", e.Description, jen.Op("*").Id("FilterNumber"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "A rich text filter condition to compare the formula result against.  \n  \nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the rich text filter condition.",
				Property:     "string",
				Type:         "object",
			}, func(e parameterElement) {
				formulaFilter.addAdaptiveFieldWithType("string", e.Description, jen.Op("*").Id("FilterRichText"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"One month deadline\",\n    \"formula\": {\n      \"date\":{\n          \"after\": \"2021-05-10\"\n      }\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Multi-select */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Multi-select",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("multi_select", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:     "contains",
				Type:         "string",
				Description:  "The value to compare the multi-select property value against.  \n  \nReturns database entries where the multi-select value matches the provided string.",
				ExampleValue: `"Marketing"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_contain",
				Type:         "string",
				Description:  "The value to compare the multi-select property value against.  \n  \nReturns database entries where the multi-select value does not match the provided string.",
				ExampleValue: `"Engineering"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_empty",
				Type:         "true",
				Description:  "Whether the multi-select property value is empty.  \n  \nReturns database entries where the multi-select value does not contain any data.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_not_empty",
				Type:         "true",
				Description:  "Whether the multi-select property value is not empty.  \n  \nReturns database entries where the multi-select value does contains data.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Programming language\",\n    \"multi_select\": {\n      \"contains\": \"TypeScript\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Number */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("number", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value differs from the provided number.",
				ExampleValue: "42",
				Property:     "does_not_equal",
				Type:         "number",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is the same as the provided number.",
				ExampleValue: "42",
				Property:     "equals",
				Type:         "number",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value exceeds the provided number.",
				ExampleValue: "42",
				Property:     "greater_than",
				Type:         "number",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is equal to or exceeds the provided number.",
				ExampleValue: "42",
				Property:     "greater_than_or_equal_to",
				Type:         "number",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the number property value is empty.  \n  \nReturns database entries where the number property value does not contain any data.",
				ExampleValue: "true",
				Property:     "is_empty",
				Type:         "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the number property value is not empty.  \n  \nReturns database entries where the number property value contains data.",
				ExampleValue: "true",
				Property:     "is_not_empty",
				Type:         "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is less than the provided number.",
				ExampleValue: "42",
				Property:     "less_than",
				Type:         "number",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The number to compare the number property value against.  \n  \nReturns database entries where the number property value is equal to or is less than the provided number.",
				ExampleValue: "42",
				Property:     "less_than_or_equal_to",
				Type:         "number",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Estimated working days\",\n    \"number\": {\n      \"less_than_or_equal_to\": 5\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* People */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "People",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("people", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "You can apply a people filter condition to people, created_by, and last_edited_by database property types.",
			}, func(e blockElement) {
				specificObject.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The people filter condition contains the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "contains",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the people property value against.  \n  \nReturns database entries where the people property value contains the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(NullUUID, omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_contain",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the people property value against.  \n  \nReturns database entries where the people property value does not contain the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(NullUUID, omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the people property value does not contain any data.  \n  \nReturns database entries where the people property value does not contain any data.",
				ExampleValue: "true",
				Property:     "is_empty",
				Type:         "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the people property value contains data.  \n  \nReturns database entries where the people property value is not empty.",
				ExampleValue: "true",
				Property:     "is_not_empty",
				Type:         "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Last edited by\",\n    \"people\": {\n      \"contains\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Relation */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Relation",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("relation", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:     "contains",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the relation property value against.  \n  \nReturns database entries where the relation property value contains the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(NullUUID, omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_contain",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the relation property value against.  \n  \nReturns entries where the relation property value does not contain the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(NullUUID, omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the relation property value does not contain data.  \n  \nReturns database entries where the relation property value does not contain any data.",
				ExampleValue: "true",
				Property:     "is_empty",
				Type:         "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Description:  "Whether the relation property value contains data.  \n  \nReturns database entries where the property value is not empty.",
				ExampleValue: "true",
				Property:     "is_not_empty",
				Type:         "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"✔️ Task List\",\n    \"relation\": {\n      \"contains\": \"0c1f7cb280904f18924ed92965055e32\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				text := regexp.MustCompile(`\w{32}`).ReplaceAllStringFunc(e.Text, func(s string) string {
					return uuid.MustParse(s).String()
				})
				b.addUnmarshalTest("Filter", extractFilter(text))
			})
		},
		func(c *comparator, b *builder) /* Rich text */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rich text",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("rich_text", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:     "contains",
				Type:         "string",
				Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that includes the provided string.",
				ExampleValue: `"Moved to Q2"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_contain",
				Type:         "string",
				Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that does not include the provided string.",
				ExampleValue: `"Moved to Q2"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_equal",
				Type:         "string",
				Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that does not match the provided string.",
				ExampleValue: `"Moved to Q2"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "ends_with",
				Type:         "string",
				Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that ends with the provided string.",
				ExampleValue: `"Q2"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "equals",
				Type:         "string",
				Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that matches the provided string.",
				ExampleValue: `"Moved to Q2"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_empty",
				Type:         "true",
				Description:  "Whether the text property value does not contain any data.  \n  \nReturns database entries with a text property value that is empty.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_not_empty",
				Type:         "true",
				Description:  "Whether the text property value contains any data.  \n  \nReturns database entries with a text property value that contains data.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "starts_with",
				Type:         "string",
				Description:  "The string to compare the text property value against.  \n  \nReturns database entries with a text property value that starts with the provided string.",
				ExampleValue: `"Moved"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Description\",\n    \"rich_text\": {\n      \"contains\": \"cross-team\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Rollup */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rollup",
			}, func(e blockElement) {
				rollupFilter = b.addAdaptiveObject("RollupFilter", "", e.Text)
				filter.addAdaptiveFieldWithType("rollup", e.Text, jen.Op("*").Id("RollupFilter"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A rollup database property can evaluate to an array, date, or number value. The filter condition for the rollup property contains a rollup key and a corresponding object value that depends on the computed value type.",
			}, func(e blockElement) {
				rollupFilter.addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* Filter conditions for array rollup values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Filter conditions for array rollup values",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type.  \n  \nReturns database entries where the rollup property value matches the provided criteria.",
				ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
				Property:     "any",
				Type:         "object",
			}, func(e parameterElement) {
				rollupFilter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Op("*").Id("Filter"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type.  \n  \nReturns database entries where every rollup property value matches the provided criteria.",
				ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
				Property:     "every",
				Type:         "object",
			}, func(e parameterElement) {
				rollupFilter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Op("*").Id("Filter"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type.  \n  \nReturns database entries where no rollup property value matches the provided criteria.",
				ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
				Property:     "none",
				Type:         "object",
			}, func(e parameterElement) {
				rollupFilter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Op("*").Id("Filter"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Related tasks\",\n    \"rollup\": {\n      \"any\": {\n        \"rich_text\": {\n          \"contains\": \"Migrate database\"\n        }\n      }\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Filter conditions for date rollup values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Filter conditions for date rollup values",
			}, func(e blockElement) {
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A rollup value is stored as a date only if the \"Earliest date\", \"Latest date\", or \"Date range\" computation is selected for the property in the Notion UI.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "A date filter condition to compare the rollup value against.  \n  \nReturns database entries where the rollup value matches the provided condition.",
				ExampleValue: "Refer to the date filter condition.",
				Property:     "date",
				Type:         "object",
			}, func(e parameterElement) {
				rollupFilter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Op("*").Id("FilterDate"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Parent project due date\",\n    \"rollup\": {\n      \"date\": {\n        \"on_or_before\": \"2023-02-08\"\n      }\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Filter conditions for number rollup values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Filter conditions for number rollup values",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "A number filter condition to compare the rollup value against.  \n  \nReturns database entries where the rollup value matches the provided condition.",
				ExampleValue: "Refer to the number filter condition.",
				Property:     "number",
				Type:         "object",
			}, func(e parameterElement) {
				rollupFilter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Op("*").Id("FilterNumber"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Total estimated working days\",\n    \"rollup\": {\n      \"number\": {\n        \"does_not_equal\": 42\n      }\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Select */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Select",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("select", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:     "equals",
				Type:         "string",
				Description:  "The string to compare the select property value against.  \n  \nReturns database entries where the select property value matches the provided string.",
				ExampleValue: `"This week"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_equal",
				Type:         "string",
				Description:  "The string to compare the select property value against.  \n  \nReturns database entries where the select property value does not match the provided string.",
				ExampleValue: `"Backlog"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_empty",
				Type:         "true",
				Description:  "Whether the select property value does not contain data.  \n  \nReturns database entries where the select property value is empty.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_not_empty",
				Type:         "true",
				Description:  "Whether the select property value contains data.  \n  \nReturns database entries where the select property value is not empty.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Frontend framework\",\n    \"select\": {\n      \"equals\": \"React\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Status */ {
			var specificObject *concreteObject

			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Status",
			}, func(e blockElement) {
				specificObject = filter.addAdaptiveFieldWithSpecificObject("status", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Property:     "equals",
				Type:         "string",
				Description:  "The string to compare the status property value against.  \n  \nReturns database entries where the status property value matches the provided string.",
				ExampleValue: `"This week"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_equal",
				Type:         "string",
				Description:  "The string to compare the status property value against.  \n  \nReturns database entries where the status property value does not match the provided string.",
				ExampleValue: `"Backlog"`,
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_empty",
				Type:         "true",
				Description:  "Whether the status property value does not contain data.  \n  \nReturns database entries where the status property value is empty.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "is_not_empty",
				Type:         "true",
				Description:  "Whether the status property value contains data.  \n  \nReturns database entries where the status property value is not empty.",
				ExampleValue: "true",
			}, func(e parameterElement) {
				specificObject.addFields(e.asField(jen.Bool(), omitEmpty))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"property\": \"Project status\",\n    \"status\": {\n      \"equals\": \"Not started\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
		func(c *comparator, b *builder) /* Timestamp */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Timestamp",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Use a timestamp filter condition to filter results based on created_time or last_edited_time values.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "timestamp",
				Type:         "created_time last_edited_time",
				Description:  "A constant string representing the type of timestamp to use as a filter.",
				ExampleValue: `"created_time"`,
			}, func(e parameterElement) {
				filter.addAdaptiveFieldWithType("timestamp", e.Description, jen.String())
			})
			c.nextMustParameter(parameterElement{
				Description:  "A date filter condition used to filter the specified timestamp.",
				ExampleValue: "Refer to the date filter condition.",
				Property:     "created_time  \nlast_edited_time",
				Type:         "object",
			}, func(e parameterElement) {
				filter.addFields(
					&field{name: "created_time", typeCode: jen.Op("*").Id("FilterDate"), comment: e.Description, omitEmpty: true},
					&field{name: "last_edited_time", typeCode: jen.Op("*").Id("FilterDate"), comment: e.Description, omitEmpty: true},
				)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"timestamp\": \"created_time\",\n    \"created_time\": {\n      \"on_or_before\": \"2022-10-13\"\n    }\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧The timestamp filter condition does not require a property name. The API throws an error if you provide one.",
			}, func(e blockElement) {
				filter.addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* ID */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "ID",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Use a timestamp filter condition to filter results based on the unique_id value.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "does_not_equal",
				Type:         "number",
				Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value differs from the provided value.",
				ExampleValue: "42",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "equals",
				Type:         "number",
				Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is the same as the provided value.",
				ExampleValue: "42",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "greater_than",
				Type:         "number",
				Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value exceeds the provided value.",
				ExampleValue: "42",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "greater_than_or_equal_to",
				Type:         "number",
				Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is equal to or exceeds the provided value.",
				ExampleValue: "42",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "less_than",
				Type:         "number",
				Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is less than the provided value.",
				ExampleValue: "42",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "less_than_or_equal_to",
				Type:         "number",
				Description:  "The value to compare the unique_id property value against.  \n  \nReturns database entries where the unique_id property value is equal to or is less than the provided value.",
				ExampleValue: "42",
			}, func(e parameterElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"and\": [\n      {\n        \"property\": \"ID\",\n        \"unique_id\": {\n          \"greater_than\": 1\n        }\n      },\n      {\n        \"property\": \"ID\",\n        \"unique_id\": {\n          \"less_than\": 3\n        }\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Compound filter conditions */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Compound filter conditions",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "You can use a compound filter condition to limit the results of a database query based on multiple conditions. This mimics filter chaining in the Notion UI.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The above filters in the Notion UI are equivalent to the following compound filter condition via the API:",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"and\": [\n    {\n      \"property\": \"Done\",\n      \"checkbox\": {\n        \"equals\": true\n      }\n    }, \n    {\n      \"or\": [\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"A\"\n        },\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"B\"\n        }\n      ]\n    }\n  ]\n}\n",
			}, func(e blockElement) {
				// TODO サンプルコードがめちゃくちゃ
				// b.addUnmarshalTest("Filter", e.Code)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A compound filter condition contains an and or or key with a value that is an array of filter objects or nested compound filter objects. Nesting is supported up to two levels deep.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "An array of filter objects or compound filter conditions.  \n  \nReturns database entries that match all of the provided filter conditions.",
				ExampleValue: "Refer to the examples below.",
				Property:     "and",
				Type:         "array",
			}, func(e parameterElement) {
				filter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Index().Id("Filter"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "An array of filter objects or compound filter conditions.  \n  \nReturns database entries that match any of the provided filter conditions",
				ExampleValue: "Refer to the examples below.",
				Property:     "or",
				Type:         "array",
			}, func(e parameterElement) {
				filter.addAdaptiveFieldWithType(e.Property, e.Description, jen.Index().Id("Filter"))
			})
		},
		func(c *comparator, b *builder) /* Example compound filter conditions */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example compound filter conditions",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"and\": [\n      {\n        \"property\": \"Complete\",\n        \"checkbox\": {\n          \"equals\": true\n        }\n      },\n      {\n        \"property\": \"Working days\",\n        \"number\": {\n          \"greater_than\": 10\n        }\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"filter\": {\n    \"or\": [\n      {\n        \"property\": \"Description\",\n        \"rich_text\": {\n          \"contains\": \"2023\"\n        }\n      },\n      {\n        \"and\": [\n          {\n            \"property\": \"Department\",\n            \"select\": {\n              \"equals\": \"Engineering\"\n            }\n          },\n          {\n            \"property\": \"Priority goal\",\n            \"checkbox\": {\n              \"equals\": true\n            }\n          }\n        ]\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("Filter", extractFilter(e.Text))
			})
		},
	)
}
