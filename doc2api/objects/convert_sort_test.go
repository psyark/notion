package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
	"github.com/samber/lo"
)

func TestSort(t *testing.T) {
	t.Parallel()

	createTest := func(e *Block, b *CodeBuilder) {
		wrapper := struct {
			Sorts []json.RawMessage `json:"sorts"`
		}{}

		lo.Must0(json.Unmarshal([]byte(e.Text), &wrapper))

		for _, s := range wrapper.Sorts {
			b.AddUnmarshalTest("Sort", string(s))
		}
	}

	c := converter.FetchDocument("https://developers.notion.com/reference/post-database-query-sort")

	var sort *SimpleObject

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A sort is a condition used to order the entries returned from a database query.",
	}).Output(func(e *Block, b *CodeBuilder) {
		sort = b.AddSimpleObject("Sort", e.Text)
	})

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A database query can be sorted by a property and/or timestamp and in a given direction. For example, a library database can be sorted by the \"Name of a book\" (i.e. property) and in ascending (i.e. direction).",
	}).Output(func(e *Block, b *CodeBuilder) {
		sort.AddComment(e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Here is an example of a sort on a database property."})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n    \"sorts\": [\n        {\n            \"property\": \"Name\",\n            \"direction\": \"ascending\"\n        }\n    ]\n}\n",
	}).Output(createTest)

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "If you’re using the Notion SDK for JavaScript, you can apply this sorting property to your query like so:"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "const { Client } = require('@notionhq/client');\n\nconst notion = new Client({ auth: process.env.NOTION_API_KEY });\n// replace with your own database ID\nconst databaseId = 'd9824bdc-8445-4327-be8b-5b47500af6ce';\n\nconst sortedRows = async () => {\n\tconst response = await notion.databases.query({\n\t  database_id: databaseId,\n\t  sorts: [\n\t    {\n\t      property: \"Name\",\n\t      direction: \"ascending\"\n\t\t  }\n\t  ],\n\t});\n  return response;\n}\n"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Database queries can also be sorted by two or more properties, which is formally called a nested sort. The sort object listed first in the nested sort list takes precedence."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Here is an example of a nested sort."})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n    \"sorts\": [\n                {\n            \"property\": \"Food group\",\n            \"direction\": \"descending\"\n        },\n        {\n            \"property\": \"Name\",\n            \"direction\": \"ascending\"\n        }\n    ]\n}\n",
	}).Output(createTest)

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "In this example, the database query will first be sorted by \"Food group\" and the set with the same food group is then sorted by \"Name\"."})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Sort object"})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Property value sort"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "This sort orders the database query by a particular property."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The sort object must contain the following properties:"})

	c.ExpectParameter(&Parameter{
		Property:     "property",
		Type:         "string",
		Description:  "The name of the property to sort against.",
		ExampleValue: `"Ingredients"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		sort.AddFields(b.NewField(e, jen.String(), OmitEmpty))
	})

	c.ExpectParameter(&Parameter{
		Property:     "direction",
		Type:         "string (enum)",
		Description:  `The direction to sort. Possible values include "ascending" and "descending".`,
		ExampleValue: `"descending"`,
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Entry timestamp sort"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "This sort orders the database query by the timestamp associated with a database entry."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The sort object must contain the following properties:"})

	c.ExpectParameter(&Parameter{
		Property:     "timestamp",
		Type:         "string (enum)",
		Description:  "The name of the timestamp to sort against. Possible values include \"created_time\" and \"last_edited_time\".",
		ExampleValue: `"last_edited_time"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		sort.AddFields(b.NewField(e, jen.String(), OmitEmpty))
	})

	c.ExpectParameter(&Parameter{
		Property:     "direction",
		Type:         "string (enum)",
		Description:  `The direction to sort. Possible values include "ascending" and "descending".`,
		ExampleValue: `"descending"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		sort.AddFields(b.NewField(e, jen.String()))
	})
}
