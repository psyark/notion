package objectdoc

import (
	"encoding/json"

	"github.com/dave/jennifer/jen"
)

func init() {
	var sort *concreteObject

	createTest := func(b *builder) func(blockElement) {
		type exampleFormat struct {
			Sorts []json.RawMessage `json:"sorts"`
		}

		return func(e blockElement) {
			ef := exampleFormat{}
			if err := json.Unmarshal([]byte(e.Text), &ef); err != nil {
				panic(err)
			}
			for _, s := range ef.Sorts {
				b.addUnmarshalTest("Sort", string(s))
			}
		}
	}

	registerTranslator(
		"https://developers.notion.com/reference/post-database-query-sort",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A sort is a condition used to order the entries returned from a database query.",
			}, func(e blockElement) {
				sort = b.addConcreteObject("Sort", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A database query can be sorted by a property and/or timestamp and in a given direction. For example, a library database can be sorted by the \"Name of a book\" (i.e. property) and in ascending (i.e. direction).",
			}, func(e blockElement) {
				sort.addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Here is an example of a sort on a database property.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n    \"sorts\": [\n        {\n            \"property\": \"Name\",\n            \"direction\": \"ascending\"\n        }\n    ]\n}",
			}, createTest(b))
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Database queries can also be sorted by two or more properties, which is formally called a nested sort. The sort object listed first in the nested sort list takes precedence.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Here is an example of a nested sort.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n    \"sorts\": [\n                {\n            \"property\": \"Food group\",\n            \"direction\": \"descending\"\n        },\n        {\n            \"property\": \"Name\",\n            \"direction\": \"ascending\"\n        }\n    ]\n}",
			}, createTest(b))
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "In this example, the database query will first be sorted by \"Food group\" and the set with the same food group is then sorted by \"Name\".",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Sort object */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Sort object",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Property value sort */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Property value sort",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "This sort orders the database query by a particular property.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The sort object must contain the following properties:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "property",
				Type:         "string",
				Description:  "The name of the property to sort against.",
				ExampleValue: `"Ingredients"`,
			}, func(e parameterElement) {
				sort.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "direction",
				Type:         "string (enum)",
				Description:  `The direction to sort. Possible values include "ascending" and "descending".`,
				ExampleValue: `"descending"`,
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* Entry timestamp sort */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Entry timestamp sort",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "This sort orders the database query by the timestamp associated with a database entry.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The sort object must contain the following properties:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "timestamp",
				Type:         "string (enum)",
				Description:  "The name of the timestamp to sort against. Possible values include \"created_time\" and \"last_edited_time\".",
				ExampleValue: `"last_edited_time"`,
			}, func(e parameterElement) {
				sort.addFields(e.asField(jen.String(), omitEmpty))
			})
			c.nextMustParameter(parameterElement{
				Property:     "direction",
				Type:         "string (enum)",
				Description:  `The direction to sort. Possible values include "ascending" and "descending".`,
				ExampleValue: `"descending"`,
			}, func(e parameterElement) {
				sort.addFields(e.asField(jen.String()))
			})
		},
	)
}
