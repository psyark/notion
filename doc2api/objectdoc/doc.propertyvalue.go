package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/property-value-object",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addAbstractObject("PropertyValue", "type", e.Text).strMapName = "PropertyValueMap"
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:  "Any property value that has other pages in its value will only use the first 25 page references. Use the [Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint to paginate through the full value.",
				Title: "Property values in the page object have a 25 page reference limit",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getAbstractObject("PropertyValue").comment += "\n\n" + e.Title + "\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text:   "All property values",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text: "\nEach page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("PropertyValue").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string",
				Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
				ExampleValue: `"f%5C%5C%3Ap"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("PropertyValue").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "type (optional)",
				Type:         "string (enum)",
				Description:  `Type of the property. Possible values are "rich_text", "number", "select", "multi_select", "status", "date", "formula", "relation", "rollup", "title", "people", "files", "checkbox", "url", "email", "phone_number", "created_time", "created_by", "last_edited_time", and "last_edited_by".`,
				ExampleValue: `"rich_text"`,
				output:       func(e *objectDocParameter, b *builder) error { return nil },
			}},
			&objectDocHeadingElement{
				Text: "Title property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("TitlePropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "title"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nTitle property value objects contain an array of rich text objects within the title property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("TitlePropertyValue").addFields(&field{
						name:     "title",
						typeCode: jen.Id("RichTextArray"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Name\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					// b.getClassStruct("TitlePropertyValue").comment += "\n\n" + e.Code
					return nil
				},
			}, {
				Code:     "{\n  \"title\": {\n    \"title\": [\n      {\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					// b.getClassStruct("TitlePropertyValue").comment += "\n\n" + e.Code
					return nil
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 inline page or person references for a `title` property. If a `title` property includes more than 25 references, then you can use the\u00a0[Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint for the specific `title` property to get its complete list of references.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("TitlePropertyValue").comment += "\n\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Rich Text property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("RichTextPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "rich_text"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nRich Text property value objects contain an array of rich text objects within the rich_text property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RichTextPropertyValue").addFields(&field{
						name:     "rich_text",
						typeCode: jen.Id("RichTextArray"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Details\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil
				},
			}, {
				Code:     "{\n  \"D[X|\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 populated inline page or person references for a `rich_text` property. If a `rich_text` property includes more than 25 references, then you can use the [Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `rich_text` property to get its complete list of references.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("RichTextPropertyValue").comment += "\n\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Number property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("NumberPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "number"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber property value objects contain a number within the number property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("NumberPropertyValue").addFields(&field{
						name:     "number",
						typeCode: jen.Qual("gopkg.in/guregu/null.v4", "Float"), // null
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Quantity\": {\n    \"number\": 1234\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil
				},
			}, {
				Code:     "{\n  \"pg@s\": {\n    \"number\": 1234\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil
				},
			}}},
			&objectDocHeadingElement{
				Text: "Select property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("SelectPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "select"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nSelect property value objects contain the following data within the select property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("SelectPropertyValue").addFields(&field{
						name:     "select",
						typeCode: jen.Op("*").Id("SelectPropertyValueData"), // null
						comment:  e.Text,
					})
					b.addSpecificObject("SelectPropertyValueData", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.\n\nWhen updating a select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyValueData").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.\n\nIf the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyValueData").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyValueData").addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Option\": {\n    \"select\": {\n      \"name\": \"Option 1\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"XMqQ\": {\n    \"select\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Status property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("StatusPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "status"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nStatus property value objects contain the following data within the status property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("StatusPropertyValue").addFields(&field{
						name:     "status",
						typeCode: jen.Id("StatusPropertyValueData"),
						comment:  e.Text,
					})
					b.addSpecificObject("StatusPropertyValueData", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyValueData").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.",
				ExampleValue: `"In progress"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyValueData").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyValueData").addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Status\": {\n    \"status\": {\n      \"name\": \"In progress\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"XMqQ\": {\n    \"status\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Multi-select property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("MultiSelectPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "multi_select"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nMulti-select property value objects contain an array of multi-select option values within the multi_select property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyValue").addFields(&field{
						name:     "multi_select",
						typeCode: jen.Index().Id("MultiSelectOption"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Multi-select option values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.addSpecificObject("MultiSelectOption", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectOption").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectOption").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectOption").addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Tags\": {\n    \"multi_select\": [\n      {\n        \"name\": \"B\"\n      },\n      {\n        \"name\": \"C\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"uyn@\": {\n    \"multi_select\": [\n      {\n        \"id\": \"3d3ca089-f964-4831-a8a2-0c6d746f4162\"\n      },\n      {\n        \"id\": \"1919ba02-1bf3-4e73-8832-8c0020f17363\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Date property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("DatePropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "date"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate property value objects contain the following data within the date property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DatePropertyValue").addFields(&field{
						name:     "date",
						typeCode: jen.Op("*").Id("DatePropertyValueData"), // null
						comment:  e.Text,
					})
					b.addSpecificObject("DatePropertyValueData", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "start",
				Type:         "string (ISO 8601 date and time)",
				Description:  "An ISO 8601 format date, with optional time.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatePropertyValueData").addFields(e.asField(jen.Id("ISO8601String")))
					return nil
				},
			}, {
				Property:     "end",
				Type:         "string (optional, ISO 8601 date and time)",
				Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatePropertyValueData").addFields(e.asField(jen.Op("*").Id("ISO8601String")))
					return nil
				},
			}, {
				Property:     "time_zone",
				Type:         "string (optional, enum)",
				Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
				ExampleValue: `"America/Los_Angeles"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatePropertyValueData").addFields(e.asField(jen.Op("*").String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipment Time\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"CbFP\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Preparation Range\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"\\\\rm}\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
				Language: "text",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Delivery Time\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"DgRt\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("FormulaPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "formula"},
							&field{name: "formula", typeCode: jen.Id("Formula")},
						),
					)
					b.addAbstractObject("Formula", "type", "")
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nFormula property value objects represent the result of evaluating a formula described in the \ndatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a formula cannot be updated directly.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("FormulaPropertyValue").comment += e.Text
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:  "Formulas returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
				Title: "Formula values may not match the Notion UI.",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("FormulaPropertyValue").comment += "\n" + e.Title + "\n" + e.Body
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Formula\": {\n    \"id\": \"1lab\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocParametersElement{{
				Description: "",
				Property:    "type",
				Type:        "string (enum)",
				output:      func(e *objectDocParameter, b *builder) error { return nil },
			}},
			&objectDocHeadingElement{
				Text: "String formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Formula").addVariant(
						b.addSpecificObject("StringFormula", e.Text).addFields(
							&fixedStringField{name: "type", value: "string"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nString formula property values contain an optional string within the string property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("StringFormula").addFields(&field{
						name:     "string",
						typeCode: jen.String(),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Number formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Formula").addVariant(
						b.addSpecificObject("NumberFormula", e.Text).addFields(
							&fixedStringField{name: "type", value: "number"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber formula property values contain an optional number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("NumberFormula").addFields(&field{
						name:     "number",
						typeCode: jen.Float64(),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Boolean formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Formula").addVariant(
						b.addSpecificObject("BooleanFormula", e.Text).addFields(
							&fixedStringField{name: "type", value: "boolean"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nBoolean formula property values contain a boolean within the boolean property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("BooleanFormula").addFields(&field{
						name:     "boolean",
						typeCode: jen.Bool(),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Date formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Formula").addVariant(
						b.addSpecificObject("DateFormula", e.Text).addFields(
							&fixedStringField{name: "type", value: "date"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate formula property values contain an optional date property value within the date property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DateFormula").addFields(&field{
						name:     "date",
						typeCode: jen.Id("DatePropertyValue"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Relation property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("RelationPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "relation"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nRelation property value objects contain an array of page references within the\u00a0relation property. A page reference is an object with an id key and a string value (UUIDv4) corresponding to a page ID in another database.\n\nA relation includes a has_more property in the Retrieve a page endpoint response object. The endpoint returns a maximum of 25 page references for a relation. If a relation has more than 25 references, then the has_more value for the relation in the response object is true. If a relation doesn’t exceed the limit, then has_more is false.\n\nNote that updating a relation property value with an empty array will clear the list.  ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RelationPropertyValue").comment += e.Text
					b.getSpecificObject("RelationPropertyValue").addFields(&field{
						name:     "relation",
						typeCode: jen.Index().Id("PageReference"),
					})
					b.getSpecificObject("RelationPropertyValue").addFields(&field{
						name:     "has_more",
						typeCode: jen.Bool(),
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Project\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"mODt\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("RollupPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "rollup"},
						),
					)
					b.addAbstractObject("Rollup", "type", "")
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nRollup property value objects represent the result of evaluating a rollup described in the \ndatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a rollup cannot be updated directly.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RollupPropertyValue").comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Rollup\": {\n    \"id\": \"aJ3l\",\n    \"rollup\": {\n      \"type\": \"number\",\n      \"number\": 1234,\n      \"function\": \"sum\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCalloutElement{
				Body:  "Rollups returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
				Title: "Rollup values may not match the Notion UI.",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("RollupPropertyValue").comment += "\n\n" + e.Title + "\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "String rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Rollup").addVariant(
						b.addSpecificObject("StringRollup", e.Text).addFields(
							&fixedStringField{name: "type", value: "string"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nString rollup property values contain an optional string within the string property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("StringRollup").addFields(&field{
						name:     "string",
						typeCode: jen.String(),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Number rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Rollup").addVariant(
						b.addSpecificObject("NumberRollup", e.Text).addFields(
							&fixedStringField{name: "type", value: "number"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber rollup property values contain a number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("NumberRollup").addFields(&field{
						name:     "number",
						typeCode: jen.Float64(),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Date rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Rollup").addVariant(
						b.addSpecificObject("DateRollup", e.Text).addFields(
							&fixedStringField{name: "type", value: "date"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate rollup property values contain a date property value within the date property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DateRollup").addFields(&field{
						name:     "date",
						typeCode: jen.Id("DatePropertyValue"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Array rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("Rollup").addVariant(
						b.addSpecificObject("ArrayRollup", e.Text).addFields(
							&fixedStringField{name: "type", value: "array"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nArray rollup property values contain an array of number, date, or string objects within the results property. \n\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("ArrayRollup").addFields(&field{
						name:     "array",
						typeCode: jen.Index().Id("Rollup"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "People property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("PeoplePropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "people"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nPeople property value objects contain an array of user objects within the people property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PeoplePropertyValue").addFields(&field{
						name:     "people",
						typeCode: jen.Id("Users"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCalloutElement{
				Body:  "The [Retrieve a page](https://developers.notion.com/reference/retrieve-a-page) endpoint can’t be guaranteed to return more than 25 people per `people` page property. If a `people` page property includes more than 25 people, then you can use the\u00a0[Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `people` property to get a complete list of people.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("PeoplePropertyValue").comment += "\n\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Files property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("FilesPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "files"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nFile property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("FilesPropertyValue").addFields(&field{
						name:     "files",
						typeCode: jen.Index().Id("File"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Files\": {\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": {\n          \t\"url\": \"https://website.domain/images/space.png\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCalloutElement{
				Body:  "Although we do not support uploading files, if you pass a `file` object containing a file hosted by Notion, it will remain one of the files. To remove any file, just do not pass it in the update response.",
				Title: "When updating a file property, the value will be overwritten by the array of files passed.",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("FilesPropertyValue").comment += "\n\n" + e.Title + "\n" + e.Body
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Checkbox property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("CheckboxPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "checkbox"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nCheckbox property value objects contain a boolean within the checkbox property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("CheckboxPropertyValue").addFields(&field{
						name:     "checkbox",
						typeCode: jen.Bool(), // null や undefined は通らない
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Done?\": {\n    \"checkbox\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"RirO\": {\n    \"checkbox\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "URL property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("UrlPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "url"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nURL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("UrlPropertyValue").addFields(&field{
						name:     "url",
						typeCode: jen.Qual("gopkg.in/guregu/null.v4", "String"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Website\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"<tdn\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Email property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("EmailPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "email"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nEmail property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("EmailPropertyValue").addFields(&field{
						name:     "email",
						typeCode: jen.Qual("gopkg.in/guregu/null.v4", "String"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's Contact\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"}=RV\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Phone number property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("PhoneNumberPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "phone_number"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nPhone number property value objects contain a string within the phone_number property. No structure is enforced.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PhoneNumberPropertyValue").addFields(&field{
						name:     "phone_number",
						typeCode: jen.String(),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's No.\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}, {
				Code:     "{\n  \"_A<p\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Created time property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("CreatedTimePropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "created_time"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of created_time cannot be updated. See the Property Item Object to see how these values are returned. \n\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("CreatedTimePropertyValue").addFields(&field{
						name:     "created_time",
						typeCode: jen.Id("ISO8601String"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Created by property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("CreatedByPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "created_by"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated by property value objects contain a user object within the created_by property. The user object describes the user who created this page. The value of created_by cannot be updated. See the Property Item Object to see how these values are returned. \n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("CreatedByPropertyValue").addFields(&field{
						name:     "created_by",
						typeCode: jen.Id("User"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Last edited time property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("LastEditedTimePropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "last_edited_time"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of last_edited_time cannot be updated. See the Property Item Object to see how these values are returned. \n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("LastEditedTimePropertyValue").addFields(&field{
						name:     "last_edited_time",
						typeCode: jen.Id("ISO8601String"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Last edited by property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getAbstractObject("PropertyValue").addVariant(
						b.addSpecificObject("LastEditedByPropertyValue", e.Text).addFields(
							&fixedStringField{name: "type", value: "last_edited_by"},
						),
					)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page. The value of last_edited_by cannot be updated. See the Property Item Object to see how these values are returned.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("LastEditedByPropertyValue").addFields(&field{
						name:     "last_edited_by",
						typeCode: jen.Id("User"),
						comment:  e.Text,
					})
					return nil
				},
			},
		},
	})
}
