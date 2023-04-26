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

	registerConverter(converter{
		url: "https://developers.notion.com/reference/post-database-query-filter",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "When you query a database, you can send a filter object in the body of the request that limits the returned entries based on the specified criteria. \n\nFor example, the below query limits the response to entries where the \"Task completed\"  checkbox property value is true: ",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.addAbstractObject("Filter", "", e.Text)
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "curl -X POST 'https://api.notion.com/v1/databases/897e5a76ae524b489fdfe71f5945d1af/query' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H 'Notion-Version: 2022-06-28' \\\n  -H \"Content-Type: application/json\" \\\n--data '{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n        \"equals\": true\n   }\n  }\n}'",
				Language: "curl",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "The filter object",
				output: func(e *objectDocHeadingElement, b *builder) {
					getSymbol[abstractObject](b, "Filter").comment += "\n" + e.Text
				},
			},
			&objectDocParagraphElement{
				Text: "\nEach filter object contains the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[abstractObject](b, "Filter").fieldsComment = e.Text
				},
			},
			&objectDocParametersElement{{
				Field:        "property",
				Type:         "string",
				Description:  "The name of the property as it appears in the database, or the property ID.",
				ExampleValue: `"Task completed"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[abstractObject](b, "Filter").addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "checkbox  \ndate\nfiles \nformula \nmulti_select \nnumber \npeople \nphone_number  \nrelation \nrich_text \nselect \nstatus \ntimestamp",
				Type:         "object",
				Description:  "The type-specific filter condition for the query. Only types listed in the Field column of this table are supported. \n\nRefer to type-specific filter conditions for details on corresponding object values.",
				ExampleValue: "\"checkbox\": {\n  \"equals\": true\n}",
				output:       func(e *objectDocParameter, b *builder) {},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"equals\": true\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocCalloutElement{
				Body:   "The filter object mimics the database [filter option in the Notion UI](https://www.notion.so/help/views-filters-and-sorts).",
				Title:  "",
				Type:   "success",
				output: func(e *objectDocCalloutElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Type-specific filter conditions",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text: "Checkbox",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("checkbox", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Description:  "Whether a checkbox property value matches the provided value exactly.\n\nReturns or excludes all database entries with an exact value match.",
				ExampleValue: "false",
				Field:        "equals",
				Type:         "boolean",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "CheckboxFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Bool(), omitEmpty))
				},
			}, {
				Description:  "Whether a checkbox property value differs from the provided value. \n\nReturns or excludes all database entries with a difference in values.",
				ExampleValue: "true",
				Field:        "does_not_equal",
				Type:         "boolean",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "CheckboxFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Bool(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"does_not_equal\": true\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Date",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("date", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocCalloutElement{
				Body:  "For the `after`, `before`, `equals, on_or_before`, and `on_or_after` fields, if a date string with a time is provided, then the comparison is done with millisecond precision.\n\nIf no timezone is provided, then the timezone defaults to UTC.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").comment += "\n" + e.Body
				},
			},
			&objectDocParagraphElement{
				Text: "A date filter condition can be used to limit date property value types and the timestamp property types created_time and last_edited_time.\n\nThe condition contains the below fields:  ",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").comment += "\n" + e.Text
				},
			},
			&objectDocParametersElement{{
				Field:        "after",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against. \n\nReturns database entries where the date property value is after the provided date.",
				ExampleValue: "\"2021-05-10\"\n\n\"2021-05-10T12:00:00\"\n\n\"2021-10-15T12:00:00-07:00\"",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
				},
			}, {
				Field:        "before",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.\n\nReturns database entries where the date property value is before the provided date.",
				ExampleValue: "\"2021-05-10\" \n\n\"2021-05-10T12:00:00\"\n\n\"2021-10-15T12:00:00-07:00\"",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
				},
			}, {
				Field:        "equals",
				Type:         "string (ISO 8601 date)",
				Description:  "The value to compare the date property value against.\n\nReturns database entries where the date property value is the provided date.",
				ExampleValue: "\"2021-05-10\" \n\n\"2021-05-10T12:00:00\"\n\n\"2021-10-15T12:00:00-07:00\"",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "The value to compare the date property value against.\n\nReturns database entries where the date property value contains no data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Description:  "The value to compare the date property value against.\n\nReturns database entries where the date property value is not empty.",
				ExampleValue: "true",
				Field:        "is_not_empty",
				Type:         "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is within the next month.",
				ExampleValue: "{}",
				Field:        "next_month",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is within the next week.",
				ExampleValue: "{}",
				Field:        "next_week",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is within the next year.",
				ExampleValue: "{}",
				Field:        "next_year",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}, {
				Description:  "The value to compare the date property value against.\n\nReturns database entries where the date property value is on or after the provided date.",
				ExampleValue: "\"2021-05-10\" \n\n\"2021-05-10T12:00:00\"\n\n\"2021-10-15T12:00:00-07:00\"",
				Field:        "on_or_after",
				Type:         "string (ISO 8601 date)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
				},
			}, {
				Description:  "The value to compare the date property value against. \n\nReturns database entries where the date property value is on or before the provided date.",
				ExampleValue: "\"2021-05-10\" \n\n\"2021-05-10T12:00:00\"\n\n\"2021-10-15T12:00:00-07:00\"",
				Field:        "on_or_before",
				Type:         "string (ISO 8601 date)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Id("ISO8601String"), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is within the past month.",
				ExampleValue: "{}",
				Field:        "past_month",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is within the past week.",
				ExampleValue: "{}",
				Field:        "past_week",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is within the past year.",
				ExampleValue: "{}",
				Field:        "past_year",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}, {
				Description:  "A filter that limits the results to database entries where the date property value is this week.",
				ExampleValue: "{}",
				Field:        "this_week",
				Type:         "object (empty)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Struct(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Due date\",\n    \"date\": {\n      \"on_or_after\": \"2023-02-08\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Files",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("files", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Description:  "Whether the files property value does not contain any data.\n\nReturns all database entries with an empty files property value.",
				ExampleValue: "true",
				Field:        "is_empty",
				Type:         "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "FilesFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Description:  "Whether the files property value contains data. \n\nReturns all entries with a populated files property value.",
				ExampleValue: "true",
				Field:        "is_not_empty",
				Type:         "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "FilesFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Blueprint\",\n    \"files\": {\n      \"is_not_empty\": true\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Formula",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("formula", "Filter", e.Text, addAbstractSpecificField(""))
				},
			},
			&objectDocParagraphElement{
				Text: "\nThe primary field of the formula filter condition object matches the type of the formula’s result. For example, to filter a formula property that computes a checkbox, use a formula filter condition object with a checkbox field containing a checkbox filter condition as its value.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "FormulaFilter").comment += e.Text
				},
			},
			&objectDocParametersElement{{
				Field:        "checkbox",
				Type:         "object",
				Description:  "A checkbox filter condition to compare the formula result against. \n\nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the checkbox filter condition.",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("checkbox", "FormulaFilterData", e.Description).addFields(e.asField(jen.Id("CheckboxFilterData")))
				},
			}, {
				Description:  "A date filter condition to compare the formula result against. \n\nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the date filter condition.",
				Field:        "date",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("date", "FormulaFilterData", e.Description).addFields(e.asField(jen.Id("DateFilterData")))
				},
			}, {
				Description:  "A number filter condition to compare the formula result against. \n\nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the number filter condition.",
				Field:        "number",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("number", "FormulaFilterData", e.Description).addFields(e.asField(jen.Id("NumberFilterData")))
				},
			}, {
				Description:  "A rich text filter condition to compare the formula result against. \n\nReturns database entries where the formula result matches the provided condition.",
				ExampleValue: "Refer to the rich text filter condition.",
				Field:        "string",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("string", "FormulaFilterData", e.Description).addFields(e.asField(jen.Id("RichTextFilterData")))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"One month deadline\",\n    \"formula\": {\n      \"date\":{\n          \"after\": \"2021-05-10\"\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Multi-select",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("multi_select", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Field:        "contains",
				Type:         "string",
				Description:  "The value to compare the multi-select property value against. \n\nReturns database entries where the multi-select value contains the provided string.",
				ExampleValue: `"Marketing"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "MultiSelectFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "does_not_contain",
				Type:         "string",
				Description:  "The value to the multi-select property value against. \n\nReturns database entries where the multi-select value does not contain the provided string.",
				ExampleValue: `"Engineering"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "MultiSelectFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "Whether the multi-select property value is empty.\n\nReturns database entries where the multi-select value does not contain any data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "MultiSelectFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Field:        "is_not_empty",
				Type:         "true",
				Description:  "Whether the multi-select property value is not empty.\n\nReturns database entries where the multi-select value does contains data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "MultiSelectFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Programming language\",\n    \"multi_select\": {\n      \"contains\": \"TypeScript\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Number",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("number", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Description:  "The number to compare the number property value against. \n\nReturns database entries where the number property value differs from the provided number.",
				ExampleValue: "42",
				Field:        "does_not_equal",
				Type:         "number",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
				},
			}, {
				Description:  "The number to compare the number property value against. \n\nReturns database entries where the number property value is the same as the provided number.",
				ExampleValue: "42",
				Field:        "equals",
				Type:         "number",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
				},
			}, {
				Description:  "The number to compare the number property value against. \n\nReturns database entries where the number property value exceeds the provided number.",
				ExampleValue: "42",
				Field:        "greater_than",
				Type:         "number",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
				},
			}, {
				Description:  "The number to compare the number property value against. \n\nReturns database entries where the number property value is equal to or exceeds the provided number.",
				ExampleValue: "42",
				Field:        "greater_than_or_equal_to",
				Type:         "number",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "Whether the number property value is empty. \n\nReturns database entries where the number property value does not contain any data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Field:        "is_not_empty",
				Type:         "true",
				Description:  "Whether the number property value is not empty. \n\nReturns database entries where the number property value contains data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Field:        "less_than",
				Type:         "number",
				Description:  "The number to compare the number property value against. \n\nReturns database entries where the page property value is less than the provided number.",
				ExampleValue: "42",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
				},
			}, {
				Description:  "The number to compare the number property value against. \n\nReturns database entries where the page property value is equal to or is less than the provided number.",
				ExampleValue: "42",
				Field:        "less_than_or_equal_to",
				Type:         "number",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberFilter").typeSpecificObject.addFields(e.asField(jen.Op("*").Float64(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Estimated working days\",\n    \"number\": {\n      \"less_than_or_equal_to\": 5\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "People",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("people", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParagraphElement{
				Text: "\nYou can apply a people filter condition to people, created_by, and last_edited_by database property types. \n\nThe people filter condition contains the following fields:",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "PeopleFilter").comment += e.Text
				},
			},
			&objectDocParametersElement{{
				Description:  "The value to compare the people property value against. \n\nReturns database entries where the people property value contains the provided string.",
				ExampleValue: "\"6c574cee-ca68-41c8-86e0-1b9e992689fb\"",
				Field:        "contains",
				Type:         "string (UUIDv4)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "PeopleFilter").typeSpecificObject.addFields(e.asField(NullUUID, omitEmpty))
				},
			}, {
				Field:        "does_not_contain",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the people property value against.\n\nReturns database entries where the people property value does not contain the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "PeopleFilter").typeSpecificObject.addFields(e.asField(NullUUID, omitEmpty))
				},
			}, {
				Description:  "Whether the people property value does not contain any data. \n\nReturns database entries where the people property value does not contain any data.",
				ExampleValue: "true",
				Field:        "is_empty",
				Type:         "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "PeopleFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Description:  "Whether the people property value contains data. \n\nReturns database entries where the people property value is not empty.",
				ExampleValue: "true",
				Field:        "is_not_empty",
				Type:         "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "PeopleFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Last edited by\",\n    \"people\": {\n      \"contains\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Relation",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("relation", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Field:        "contains",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the relation property value against. \n\nReturns database entries where the relation property value contains the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RelationFilter").typeSpecificObject.addFields(e.asField(NullUUID, omitEmpty))
				},
			}, {
				Field:        "does_not_contain",
				Type:         "string (UUIDv4)",
				Description:  "The value to compare the relation property value against. \n\nReturns entries where the relation property value does not contain the provided string.",
				ExampleValue: `"6c574cee-ca68-41c8-86e0-1b9e992689fb"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RelationFilter").typeSpecificObject.addFields(e.asField(NullUUID, omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "Whether the relation property value does not contain data. \n\nReturns database entries where the relation property value does not contain any data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RelationFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Field:        "is_not_empty",
				Type:         "true",
				Description:  "Whether the relation property value contains data. \n\nReturns database entries where the property value is not empty.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RelationFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"✔️ Task List\",\n    \"relation\": {\n      \"contains\": \"0c1f7cb280904f18924ed92965055e32\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					code := regexp.MustCompile(`\w{32}`).ReplaceAllStringFunc(e.Code, func(s string) string {
						return uuid.MustParse(s).String()
					})
					b.addUnmarshalTest("Filter", extractFilter(code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Rich text ",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("rich_text", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Field:        "contains",
				Type:         "string",
				Description:  "The string to compare the text property value against.\n\nReturns database entries with a text property value that includes the provided string.",
				ExampleValue: `"Moved to Q2"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "does_not_contain",
				Type:         "string",
				Description:  "The string to compare the text property value against.\n\nReturns database entries with a text property value that does not include the provided string.",
				ExampleValue: `"Moved to Q2"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "does_not_equal",
				Type:         "string",
				Description:  "The string to compare the text property value against.\n\nReturns database entries with a text property value that does not match the provided string.",
				ExampleValue: `"Moved to Q2"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "ends_with",
				Type:         "string",
				Description:  "The string to compare the text property value against.\n\nReturns database entries with a text property value that ends with the provided string.",
				ExampleValue: `"Q2"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "equals",
				Type:         "string",
				Description:  "The string to compare the text property value against.\n\nReturns database entries with a text property value that matches the provided string.",
				ExampleValue: `"Moved to Q2"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "Whether the text property value does not contain any data. \n\nReturns database entries with a text property value that is empty.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Field:        "is_not_empty",
				Type:         "true",
				Description:  "Whether the text property value contains any data. \n\nReturns database entries with a text property value that contains data.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.Bool(), omitEmpty))
				},
			}, {
				Field:        "starts_with",
				Type:         "string",
				Description:  "The string to compare the text property value against.\n\nReturns database entries with a text property value that starts with the provided string.",
				ExampleValue: `"Moved"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "RichTextFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Description\",\n    \"rich_text\": {\n      \"contains\": \"cross-team\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Rollup",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("rollup", "Filter", e.Text, addAbstractSpecificField(""))
				},
			},
			&objectDocParagraphElement{
				Text: "\nA rollup database property can evaluate to an array, date, or number value. The filter condition for the rollup property contains a rollup key and a corresponding object value that depends on the computed value type. \n",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[abstractObject](b, "RollupFilterData").comment += e.Text
				},
			},
			&objectDocHeadingElement{
				Text:   "Filter conditions for array rollup values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParametersElement{{
				Field:        "any",
				Type:         "object",
				Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type. \n\nReturns database entries where the rollup property value matches the provided criteria.",
				ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("any", "RollupFilterData", e.Description).addFields(e.asInterfaceField("Filter"))
				},
			}, {
				Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type. \n\nReturns database entries where every rollup property value matches the provided criteria.",
				ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
				Field:        "every",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("every", "RollupFilterData", e.Description).addFields(e.asInterfaceField("Filter"))
				},
			}, {
				Description:  "The value to compare each rollup property value against. Can be a filter condition for any other type. \n\nReturns database entries where no rollup property value matches the provided criteria.",
				ExampleValue: "\"rich_text\": {\n\"contains\": \"Take Fig on a walk\"\n}",
				Field:        "none",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					b.addDerived("none", "RollupFilterData", e.Description).addFields(e.asInterfaceField("Filter"))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Related tasks\",\n    \"rollup\": {\n      \"any\": {\n        \"rich_text\": {\n          \"contains\": \"Migrate database\"\n        }\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Filter conditions for date rollup values ",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("date", "RollupFilterData", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nA rollup value is stored as a date only if the \"Earliest date\", \"Latest date\", or \"Date range\" computation is selected for the property in the Notion UI. ",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "DateRollupFilterData").comment += e.Text
				},
			},
			&objectDocParametersElement{{
				Description:  "A date filter condition to compare the rollup value against. \n\nReturns database entries where the rollup value matches the provided condition.",
				ExampleValue: "Refer to the date filter condition.",
				Field:        "date",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DateRollupFilterData").addFields(e.asField(jen.Id("DateFilterData")))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Parent project due date\",\n    \"rollup\": {\n      \"date\": {\n        \"on_or_before\": \"2023-02-08\"\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Filter conditions for number rollup values ",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("number", "RollupFilterData", e.Text)
				},
			},
			&objectDocParametersElement{{
				Description:  "A number filter condition to compare the rollup value against. \n\nReturns database entries where the rollup value matches the provided condition.",
				ExampleValue: "Refer to the number filter condition.",
				Field:        "number",
				Type:         "object",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "NumberRollupFilterData").addFields(e.asField(jen.Id("NumberFilterData")))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Total estimated working days\",\n    \"rollup\": {\n      \"number\": {\n        \"does_not_equal\": 42\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Select",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("select", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Field:        "equals",
				Type:         "string",
				Description:  "The string to compare the select property value against.\n\nReturns database entries where the select property value matches the provided string.",
				ExampleValue: `"This week"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "SelectFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "does_not_equal",
				Type:         "string",
				Description:  "The string to compare the select property value against.\n\nReturns database entries where the select property value does not match the provided string.",
				ExampleValue: `"Backlog"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "SelectFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "Whether the select property value does not contain data.\n\nReturns database entries where the select property value is empty.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "SelectFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "is_not_empty",
				Type:         "true",
				Description:  "Whether the select property value contains data.\n\nReturns database entries where the select property value is not empty.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "SelectFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Frontend framework\",\n    \"select\": {\n      \"equals\": \"React\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Status",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("status", "Filter", e.Text, addSpecificField())
				},
			},
			&objectDocParametersElement{{
				Field:        "equals",
				Type:         "string",
				Description:  "The string to compare the status property value against.\n\nReturns database entries where the status property value matches the provided string.",
				ExampleValue: `"This week"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "StatusFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "does_not_equal",
				Type:         "string",
				Description:  "The string to compare the status property value against.\n\nReturns database entries where the status property value does not match the provided string.",
				ExampleValue: `"Backlog"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "StatusFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "is_empty",
				Type:         "true",
				Description:  "Whether the status property value does not contain data.\n\nReturns database entries where the status property value is empty.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "StatusFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}, {
				Field:        "is_not_empty",
				Type:         "true",
				Description:  "Whether the status property value contains data.\n\nReturns database entries where the status property value is not empty.",
				ExampleValue: "true",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "StatusFilter").typeSpecificObject.addFields(e.asField(jen.String(), omitEmpty))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"property\": \"Project status\",\n    \"status\": {\n      \"equals\": \"Not started\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocHeadingElement{
				Text: "Timestamp",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("timestamp", "Filter", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nUse a timestamp filter condition to filter results based on created_time or last_edited_time values. ",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "TimestampFilter").comment += e.Text
				},
			},
			&objectDocParametersElement{{
				Field:        "timestamp",
				Type:         "created_time last_edited_time",
				Description:  "A constant string representing the type of timestamp to use as a filter.",
				ExampleValue: `"created_time"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "TimestampFilter").addFields(e.asField(jen.String()))
				},
			}, {
				Field:        "created_time\nlast_edited_time",
				Type:         "object",
				Description:  "A date filter condition used to filter the specified timestamp.",
				ExampleValue: "Refer to the date filter condition.",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "TimestampFilter").addFields(
						&field{name: "created_time", typeCode: jen.Op("*").Id("DateFilterData"), comment: e.Description, omitEmpty: true},
						&field{name: "last_edited_time", typeCode: jen.Op("*").Id("DateFilterData"), comment: e.Description, omitEmpty: true},
					)
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"timestamp\": \"created_time\",\n    \"created_time\": {\n      \"on_or_before\": \"2022-10-13\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The `timestamp` filter condition does not require a property name. The API throws an error if you provide one.",
				Title: "",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) {
					getSymbol[concreteObject](b, "TimestampFilter").comment += "\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text:   "Compound filter conditions",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nYou can use a compound filter condition to limit the results of a database query based on multiple conditions. This mimics filter chaining in the Notion UI. ",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocImageElement{Images: []*objectDocImageElementImage{{
				Caption: "An example filter chain in the Notion UI",
				Color:   "#000000",
				Height:  550,
				Name:    "Untitled.png",
				Url:     "https://files.readme.io/14ec7e8-Untitled.png",
				Width:   1340,
				output:  func(e *objectDocImageElementImage, b *builder) {},
			}}},
			&objectDocParagraphElement{
				Text:   "The above filters in the Notion UI are equivalent to the following compound filter condition via the API:",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"and\": [\n    {\n      \"property\": \"Done\",\n      \"checkbox\": {\n        \"equals\": true\n      }\n    }, \n    {\n      \"or\": [\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"A\"\n        },\n        {\n          \"property\": \"Tags\",\n          \"contains\": \"B\"\n        }\n      ]\n    }\n  ]\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", e.Code)
				},
			}}},
			&objectDocParagraphElement{
				Text: "A compound filter condition contains an and or or key with a value that is an array of filter objects or nested compound filter objects. Nesting is supported up to two levels deep. ",
				output: func(e *objectDocParagraphElement, b *builder) {
					// TODO
				},
			},
			&objectDocParametersElement{{
				Description:  "An array of filter objects or compound filter conditions.\n\nReturns database entries that match all of the provided filter conditions.",
				ExampleValue: "Refer to the examples below.",
				Field:        "and",
				Type:         "array",
				output: func(e *objectDocParameter, b *builder) {
					// TODO
				},
			}, {
				Description:  "An array of filter objects or compound filter conditions.\n\nReturns database entries that match any of the provided filter conditions",
				ExampleValue: "Refer to the examples below.",
				Field:        "or",
				Type:         "array",
				output: func(e *objectDocParameter, b *builder) {
					// TODO
				},
			}},
			&objectDocHeadingElement{
				Text: "Example compound filter conditions",
				output: func(e *objectDocHeadingElement, b *builder) {
					// TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"and\": [\n      {\n        \"property\": \"Complete\",\n        \"checkbox\": {\n          \"equals\": true\n        }\n      },\n      {\n        \"property\": \"Working days\",\n        \"number\": {\n          \"greater_than\": 10\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"filter\": {\n    \"or\": [\n      {\n        \"property\": \"Description\",\n        \"rich_text\": {\n          \"contains\": \"2023\"\n        }\n      },\n      {\n        \"and\": [\n          {\n            \"property\": \"Department\",\n            \"select\": {\n              \"equals\": \"Engineering\"\n            }\n          },\n          {\n            \"property\": \"Priority goal\",\n            \"checkbox\": {\n              \"equals\": true\n            }\n          }\n        ]\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("Filter", extractFilter(e.Code))
				},
			}}},
		},
	})
}
