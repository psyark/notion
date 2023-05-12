package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	// addTest := func(b *builder) func(blockElement) {
	// 	return func(e blockElement) {
	// 		b.addUnmarshalTest("PropertyValueMap", e.Text)
	// 	}
	// }

	var propertyValue, formula *adaptiveObject

	registerTranslator(
		"https://developers.notion.com/reference/property-value-object",

		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.",
			}, func(e blockElement) {
				propertyValue = b.addAdaptiveObject("PropertyValue", "type", e.Text)
				for _, f := range propertyValue.fields {
					if f, ok := f.(*field); ok && f.name == "type" {
						f.omitEmpty = true // TODO きれいに
					}
				}
				// pv.specialMethods = append(pv.specialMethods, &getMethodCoder{}, &setMethodCoder{})
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Any property value that has other pages in its value will only use the first 25 page references. Use the Retrieve a page property endpoint to paginate through the full value.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject]("PropertyValue").addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* All property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "All property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Each page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
			}, func(e blockElement) {
				getSymbol[adaptiveObject]("PropertyValue").addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
				ExampleValue: "\"f%5C%5C%3Ap\"",
				Property:     "id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("PropertyValue").addFields(e.asField(jen.String(), omitEmpty)) // Rollup内でIDが無い場合がある
			})
			c.nextMustParameter(parameterElement{
				Description:  "Type of the property. Possible values are \"rich_text\", \"number\", \"select\", \"multi_select\", \"status\", \"date\", \"formula\", \"relation\", \"rollup\", \"title\", \"people\", \"files\", \"checkbox\", \"url\", \"email\", \"phone_number\", \"created_time\", \"created_by\", \"last_edited_time\", and \"last_edited_by\".",
				ExampleValue: "\"rich_text\"",
				Property:     "type (optional)",
				Type:         "string (enum)",
			}, func(e parameterElement) {})
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
				propertyValue.addAdaptiveFieldWithType("title", e.Text, jen.Index().Id("RichText"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Name\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"title\": {\n    \"title\": [\n      {\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "The Retrieve a page endpoint returns a maximum of 25 inline page or person references for a title property. If a title property includes more than 25 references, then you can use the\u00a0Retrieve a page property endpoint for the specific title property to get its complete list of references.",
			}, func(e blockElement) {})
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
				propertyValue.addAdaptiveFieldWithType("rich_text", e.Text, jen.Index().Id("RichText"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Details\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"D[X|\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "The Retrieve a page endpoint returns a maximum of 25 populated inline page or person references for a rich_text property. If a rich_text property includes more than 25 references, then you can use the Retrieve a page property endpoint for the specific rich_text property to get its complete list of references.",
			}, func(e blockElement) {})
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
				propertyValue.addAdaptiveFieldWithType("number", e.Text, NullFloat)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Quantity\": {\n    \"number\": 1234\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"pg@s\": {\n    \"number\": 1234\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
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
				propertyValue.addAdaptiveFieldWithType("select", e.Text, jen.Op("*").Id("Option")) // may null
			})
			c.nextMustParameter(parameterElement{
				Description:  "ID of the option.\n\nWhen updating a select property, you can use either name or id.",
				ExampleValue: "\"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb\"",
				Property:     "id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Name of the option as it appears in Notion.\n\nIf the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: "\"Fruit\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: "\"red\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Option\": {\n    \"select\": {\n      \"name\": \"Option 1\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"XMqQ\": {\n    \"select\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Status property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Status property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Status property value objects contain the following data within the status property:",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("status", e.Text, jen.Op("*").Id("Option"))
			})
			c.nextMustParameter(parameterElement{
				Description:  "ID of the option.",
				ExampleValue: "\"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb\"",
				Property:     "id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Name of the option as it appears in Notion.",
				ExampleValue: "\"In progress\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: "\"red\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Status\": {\n    \"status\": {\n      \"name\": \"In progress\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"XMqQ\": {\n    \"status\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO
			})
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
				propertyValue.addAdaptiveFieldWithType("multi_select", e.Text, jen.Index().Id("Option"))
			})
		},
		func(c *comparator, b *builder) /* Multi-select option values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Multi-select option values",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
				ExampleValue: "\"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb\"",
				Property:     "id",
				Type:         "string (UUIDv4)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: "\"Fruit\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: "\"red\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Tags\": {\n    \"multi_select\": [\n      {\n        \"name\": \"B\"\n      },\n      {\n        \"name\": \"C\"\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"uyn@\": {\n    \"multi_select\": [\n      {\n        \"id\": \"3d3ca089-f964-4831-a8a2-0c6d746f4162\"\n      },\n      {\n        \"id\": \"1919ba02-1bf3-4e73-8832-8c0020f17363\"\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO
			})
		},
		func(c *comparator, b *builder) /* Date property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date property value objects contain the following data within the date property:",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithSpecificObject("date", e.Text, b)
			})
			c.nextMustParameter(parameterElement{
				Description:  "An ISO 8601 format date, with optional time.",
				ExampleValue: "\"2020-12-08T12:00:00Z\"",
				Property:     "start",
				Type:         "string (ISO 8601 date and time)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyValueDate").addFields(e.asField(jen.Id("ISO8601String")))
			})
			c.nextMustParameter(parameterElement{
				Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
				ExampleValue: "\"2020-12-08T12:00:00Z\"",
				Property:     "end",
				Type:         "string (optional, ISO 8601 date and time)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyValueDate").addFields(e.asField(jen.Op("*").Id("ISO8601String"))) // APIでnullがあるのでomitemptyしない
			})
			c.nextMustParameter(parameterElement{
				Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
				ExampleValue: "\"America/Los_Angeles\"",
				Property:     "time_zone",
				Type:         "string (optional, enum)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyValueDate").addFields(e.asField(jen.Op("*").String())) // APIでnullがあるのでomitemptyしない
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Shipment Time\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"CbFP\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Preparation Range\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"\\\\rm}\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Delivery Time\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"DgRt\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
		},
		func(c *comparator, b *builder) /* Formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Formula property values",
			}, func(e blockElement) {
				formula = b.addAdaptiveObject("Formula", "type", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Formula property value objects represent the result of evaluating a formula described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a formula cannot be updated directly.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("formula", e.Text, jen.Op("*").Id("Formula"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Formulas returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Formula\": {\n    \"id\": \"1lab\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustParameter(parameterElement{
				Description:  "",
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
			}, func(e blockElement) {
				formula.addAdaptiveFieldWithType("string", e.Text, NullString)
			})
		},
		func(c *comparator, b *builder) /* Number formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Number formula property values contain an optional number within the number property.",
			}, func(e blockElement) {
				formula.addAdaptiveFieldWithType("number", e.Text, jen.Float64())
			})
		},
		func(c *comparator, b *builder) /* Boolean formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Boolean formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Boolean formula property values contain a boolean within the boolean property.",
			}, func(e blockElement) {
				formula.addAdaptiveFieldWithType("boolean", e.Text, jen.Bool())
			})
		},
		func(c *comparator, b *builder) /* Date formula property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date formula property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date formula property values contain an optional date property value within the date property.",
			}, func(e blockElement) {
				formula.addAdaptiveFieldWithType("date", e.Text, jen.Id("PropertyValueDate"))
			})
		},
		func(c *comparator, b *builder) /* Relation property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Relation property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Relation property value objects contain an array of page references within the\u00a0relation property. A page reference is an object with an id key and a string value (UUIDv4) corresponding to a page ID in another database.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("relation", e.Text, jen.Index().Id("PageReference"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A relation includes a has_more property in the Retrieve a page endpoint response object. The endpoint returns a maximum of 25 page references for a relation. If a relation has more than 25 references, then the has_more value for the relation in the response object is true. If a relation doesn’t exceed the limit, then has_more is false.",
			}, func(e blockElement) {
				// TODO 他と似たようにする
				getSymbol[adaptiveObject]("PropertyValue").addFields(&field{
					name:               "has_more",
					typeCode:           jen.Bool(),
					comment:            e.Text,
					discriminatorValue: "relation",
				})
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Note that updating a relation property value with an empty array will clear the list.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Project\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"mODt\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
		},
		func(c *comparator, b *builder) /* Rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rollup property values",
			}, func(e blockElement) {
				// TODO RollupはPropertyItemで、FormulaはPropertyValueでtranslateしているのを統一
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Rollup property value objects represent the result of evaluating a rollup described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a rollup cannot be updated directly.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("rollup", e.Text, jen.Op("*").Id("Rollup"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Rollup\": {\n    \"id\": \"aJ3l\",\n    \"rollup\": {\n      \"type\": \"number\",\n      \"number\": 1234,\n      \"function\": \"sum\"\n    }\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Rollups returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* String rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "String rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "String rollup property values contain an optional string within the string property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Number rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Number rollup property values contain a number within the number property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Date rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Date rollup property values contain a date property value within the date property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Array rollup property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Array rollup property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Array rollup property values contain an array of number, date, or string objects within the results property.",
			}, func(e blockElement) {})
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
				propertyValue.addAdaptiveFieldWithType("people", e.Text, jen.Index().Id("User"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "The Retrieve a page endpoint can’t be guaranteed to return more than 25 people per people page property. If a people page property includes more than 25 people, then you can use the\u00a0Retrieve a page property endpoint for the specific people property to get a complete list of people.",
			}, func(e blockElement) {})
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
				propertyValue.addAdaptiveFieldWithType("files", e.Text, jen.Index().Id("File"))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Files\": {\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": {\n          \t\"url\": \"https://website.domain/images/space.png\"\n        }\n      }\n    ]\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "Although we do not support uploading files, if you pass a file object containing a file hosted by Notion, it will remain one of the files. To remove any file, just do not pass it in the update response.",
			}, func(e blockElement) {})
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
				propertyValue.addAdaptiveFieldWithType("checkbox", e.Text, jen.Bool()) // never null | undefined
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Done?\": {\n    \"checkbox\": true\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"RirO\": {\n    \"checkbox\": true\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
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
				propertyValue.addAdaptiveFieldWithType("url", e.Text, NullString)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Website\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"<tdn\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
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
				propertyValue.addAdaptiveFieldWithType("email", e.Text, NullString)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Shipper's Contact\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"}=RV\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
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
				propertyValue.addAdaptiveFieldWithType("phone_number", e.Text, NullString)
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"Shipper's No.\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "{\n  \"_A<p\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
			}, func(e blockElement) {
				// TODO テスト通す
			})
		},
		func(c *comparator, b *builder) /* Created time property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Created time property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of created_time cannot be updated. See the Property Item Object to see how these values are returned.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("created_time", e.Text, jen.Id("ISO8601String"))
			})
		},
		func(c *comparator, b *builder) /* Created by property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Created by property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page. The value of created_by cannot be updated. See the Property Item Object to see how these values are returned.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("created_by", e.Text, jen.Id("User"))
			})
		},
		func(c *comparator, b *builder) /* Last edited time property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Last edited time property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of last_edited_time cannot be updated. See the Property Item Object to see how these values are returned.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("last_edited_time", e.Text, jen.Id("ISO8601String"))
			})
		},
		func(c *comparator, b *builder) /* Last edited by property values */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Last edited by property values",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page. The value of last_edited_by cannot be updated. See the Property Item Object to see how these values are returned.",
			}, func(e blockElement) {
				propertyValue.addAdaptiveFieldWithType("last_edited_by", e.Text, jen.Id("User"))
			})
		},
	)
}
