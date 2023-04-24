package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

// PropertyValueに value()reflect.Value メソッドを追加し、bindingを行えるようにする
type getMethodCoder struct{}

func (c *getMethodCoder) declarationCode() jen.Code {
	return jen.Id("get").Call().Qual("reflect", "Value").Comment("for binding")
}

func (c *getMethodCoder) implementationCode(so *specificObject) jen.Code {
	var field fieldCoder
	for _, f := range so.fields {
		if f.goName() != "Type" && f.goName() != "HasMore" {
			field = f
			break
		}
	}

	return jen.Comment("for binding").Line().Func().Params(jen.Id("o").Op("*").Id(so.name())).Id("get").Call().Qual("reflect", "Value").Block(
		jen.Return().Qual("reflect", "ValueOf").Call(jen.Id("o").Dot(field.goName())),
	).Line()
}

type setMethodCoder struct{}

func (c *setMethodCoder) declarationCode() jen.Code {
	return jen.Id("set").Call(jen.Qual("reflect", "Value")).Comment("for binding")
}

func (c *setMethodCoder) implementationCode(so *specificObject) jen.Code {
	var field fieldCoder
	for _, f := range so.fields {
		if f.goName() != "Type" && f.goName() != "HasMore" {
			field = f
			break
		}
	}

	return jen.Comment("for binding").Line().Func().Params(jen.Id("o").Op("*").Id(so.name())).Id("set").Call(jen.Id("v").Qual("reflect", "Value")).Block(
		jen.Id("o").Dot(field.goName()).Op("=").Id("v").Dot("Interface").Call().Assert(field.getTypeCode()),
	).Line()
}

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/property-value-object",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.",
				output: func(e *objectDocParagraphElement, b *builder) {
					pv := b.addAbstractObject("PropertyValue", "type", e.Text)
					b.addAbstractList("PropertyValue", "PropertyValueArray")
					b.addAbstractMap("PropertyValue", "PropertyValueMap")
					pv.specialMethods = append(pv.specialMethods, &getMethodCoder{}, &setMethodCoder{})
				},
			},
			&objectDocCalloutElement{
				Body:  "Any property value that has other pages in its value will only use the first 25 page references. Use the [Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint to paginate through the full value.",
				Title: "Property values in the page object have a 25 page reference limit",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getAbstractObject("PropertyValue").comment += "\n\n" + e.Title + "\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text:   "All property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text: "\nEach page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getAbstractObject("PropertyValue").comment += e.Text
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string",
				Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
				ExampleValue: `"f%5C%5C%3Ap"`,
				output: func(e *objectDocParameter, b *builder) {
					field := e.asField(jen.String())
					field.omitEmpty = true // Rollup内でIDが無い場合がある
					b.getAbstractObject("PropertyValue").addFields(field)
				},
			}, {
				Property:     "type (optional)",
				Type:         "string (enum)",
				Description:  `Type of the property. Possible values are "rich_text", "number", "select", "multi_select", "status", "date", "formula", "relation", "rollup", "title", "people", "files", "checkbox", "url", "email", "phone_number", "created_time", "created_by", "last_edited_time", and "last_edited_by".`,
				ExampleValue: `"rich_text"`,
				output:       func(e *objectDocParameter, b *builder) {},
			}},
			&objectDocHeadingElement{
				Text: "Title property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("title", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nTitle property value objects contain an array of rich text objects within the title property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("TitlePropertyValue").addFields(&field{
						name:     "title",
						typeCode: jen.Id("RichTextArray"),
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Name\": {\n    \"title\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					// b.getClassStruct("TitlePropertyValue").comment += "\n\n" + e.Code
				},
			}, {
				Code:     "{\n  \"title\": {\n    \"title\": [\n      {\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"content\": \"The title\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					// b.getClassStruct("TitlePropertyValue").comment += "\n\n" + e.Code
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 inline page or person references for a `title` property. If a `title` property includes more than 25 references, then you can use the\u00a0[Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint for the specific `title` property to get its complete list of references.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getSpecificObject("TitlePropertyValue").comment += "\n\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text: "Rich Text property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("rich_text", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nRich Text property value objects contain an array of rich text objects within the rich_text property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("RichTextPropertyValue").addFields(&field{
						name:     "rich_text",
						typeCode: jen.Id("RichTextArray"),
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Details\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
				},
			}, {
				Code:     "{\n  \"D[X|\": {\n    \"rich_text\": [\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"Some more text with \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"some\"\n        },\n        \"annotations\": {\n          \"italic\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"fun\"\n        },\n        \"annotations\": {\n          \"bold\": true\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \" \"\n        }\n      },\n      {\n        \"type\": \"text\",\n        \"text\": {\n          \"content\": \"formatting\"\n        },\n        \"annotations\": {\n          \"color\": \"pink\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
				},
			}}},
			&objectDocCalloutElement{
				Body:  "The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 populated inline page or person references for a `rich_text` property. If a `rich_text` property includes more than 25 references, then you can use the [Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `rich_text` property to get its complete list of references.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getSpecificObject("RichTextPropertyValue").comment += "\n\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text: "Number property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("number", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber property value objects contain a number within the number property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("NumberPropertyValue").addFields(&field{
						name:     "number",
						typeCode: NullFloat,
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Quantity\": {\n    \"number\": 1234\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
				},
			}, {
				Code:     "{\n  \"pg@s\": {\n    \"number\": 1234\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
				},
			}}},
			&objectDocHeadingElement{
				Text: "Select property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("select", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nSelect property value objects contain the following data within the select property:",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("SelectPropertyValue").addFields(&field{
						name:     "select",
						typeCode: jen.Op("*").Id("Option"), // null
						comment:  e.Text,
					})
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.\n\nWhen updating a select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.\n\nIf the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Option\": {\n    \"select\": {\n      \"name\": \"Option 1\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"XMqQ\": {\n    \"select\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Status property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("status", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nStatus property value objects contain the following data within the status property:",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("StatusPropertyValue").addFields(&field{
						name:     "status",
						typeCode: jen.Id("Option"),
						comment:  e.Text,
					})
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.",
				ExampleValue: `"In progress"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Status\": {\n    \"status\": {\n      \"name\": \"In progress\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"XMqQ\": {\n    \"status\": {\n      \"id\": \"c3406b80-bda4-45e0-add2-2748ac1527b\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Multi-select property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("multi_select", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nMulti-select property value objects contain an array of multi-select option values within the multi_select property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("MultiSelectPropertyValue").addFields(&field{
						name:     "multi_select",
						typeCode: jen.Index().Id("Option"),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text:   "Multi-select option values",
				output: func(e *objectDocHeadingElement, b *builder) {}, // Optionで共通化
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output:       func(e *objectDocParameter, b *builder) {}, // Optionで共通化
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Tags\": {\n    \"multi_select\": [\n      {\n        \"name\": \"B\"\n      },\n      {\n        \"name\": \"C\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"uyn@\": {\n    \"multi_select\": [\n      {\n        \"id\": \"3d3ca089-f964-4831-a8a2-0c6d746f4162\"\n      },\n      {\n        \"id\": \"1919ba02-1bf3-4e73-8832-8c0020f17363\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Date property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("date", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate property value objects contain the following data within the date property:",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("DatePropertyValue").addFields(&field{
						name:     "date",
						typeCode: jen.Op("*").Id("DatePropertyValueData"), // null
						comment:  e.Text,
					})
					// 本来typeObjectを使いたいが、PropertyValueDataのバインディング利用の特殊性からこうしている
					b.addSpecificObject("DatePropertyValueData", e.Text)
				},
			},
			&objectDocParametersElement{{
				Property:     "start",
				Type:         "string (ISO 8601 date and time)",
				Description:  "An ISO 8601 format date, with optional time.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) {
					b.getSpecificObject("DatePropertyValueData").addFields(e.asField(jen.Id("ISO8601String")))
				},
			}, {
				Property:     "end",
				Type:         "string (optional, ISO 8601 date and time)",
				Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) {
					b.getSpecificObject("DatePropertyValueData").addFields(e.asField(jen.Op("*").Id("ISO8601String")))
				},
			}, {
				Property:     "time_zone",
				Type:         "string (optional, enum)",
				Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
				ExampleValue: `"America/Los_Angeles"`,
				output: func(e *objectDocParameter, b *builder) {
					b.getSpecificObject("DatePropertyValueData").addFields(e.asField(jen.Op("*").String()))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipment Time\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"CbFP\": {\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Preparation Range\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"\\\\rm}\": {\n    \"date\": {\n      \"start\": \"2021-04-26\",\n      \"end\": \"2021-05-07\"\n    }\n  }\n}",
				Language: "text",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Delivery Time\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"DgRt\": {\n    \"date\": {\n      \"start\": \"2020-12-08T12:00:00Z\",\n      \"time_zone\": \"America/New_York\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("formula", "PropertyValue", e.Text).addFields(
						&interfaceField{name: "formula", typeName: "Formula"},
					)
					b.addAbstractObject("Formula", "type", "")
				},
			},
			&objectDocParagraphElement{
				Text: "\nFormula property value objects represent the result of evaluating a formula described in the \ndatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a formula cannot be updated directly.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("FormulaPropertyValue").comment += e.Text
				},
			},
			&objectDocCalloutElement{
				Body:  "Formulas returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
				Title: "Formula values may not match the Notion UI.",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getSpecificObject("FormulaPropertyValue").comment += "\n" + e.Title + "\n" + e.Body
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Formula\": {\n    \"id\": \"1lab\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocParametersElement{{
				Description: "",
				Property:    "type",
				Type:        "string (enum)",
				output:      func(e *objectDocParameter, b *builder) {},
			}},
			&objectDocHeadingElement{
				Text: "String formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("string", "Formula", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nString formula property values contain an optional string within the string property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("StringFormula").addFields(&field{
						name:     "string",
						typeCode: NullString,
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Number formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("number", "Formula", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber formula property values contain an optional number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("NumberFormula").addFields(&field{
						name:     "number",
						typeCode: jen.Float64(),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Boolean formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("boolean", "Formula", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nBoolean formula property values contain a boolean within the boolean property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("BooleanFormula").addFields(&field{
						name:     "boolean",
						typeCode: jen.Bool(),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Date formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("date", "Formula", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate formula property values contain an optional date property value within the date property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("DateFormula").addFields(&field{
						name:     "date",
						typeCode: jen.Id("DatePropertyValue"),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Relation property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("relation", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nRelation property value objects contain an array of page references within the\u00a0relation property. A page reference is an object with an id key and a string value (UUIDv4) corresponding to a page ID in another database.\n\nA relation includes a has_more property in the Retrieve a page endpoint response object. The endpoint returns a maximum of 25 page references for a relation. If a relation has more than 25 references, then the has_more value for the relation in the response object is true. If a relation doesn’t exceed the limit, then has_more is false.\n\nNote that updating a relation property value with an empty array will clear the list.  ",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("RelationPropertyValue").comment += e.Text
					b.getSpecificObject("RelationPropertyValue").addFields(&field{
						name:     "relation",
						typeCode: jen.Index().Id("PageReference"),
					})
					b.getSpecificObject("RelationPropertyValue").addFields(&field{
						name:     "has_more",
						typeCode: jen.Bool(),
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Project\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"mODt\": {\n    \"relation\": [\n      {\n        \"id\": \"1d148a9e-783d-47a7-b3e8-2d9c34210355\"\n      }\n    ],\n      \"has_more\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					// PropertyValueとPropertyItemで共通の Rollup を使う
					// PropertyValue のArray Rollupにはfunction が無いなど不正確であり、
					// 比較的ドキュメントが充実しているPropertyItemに任せるため、こちらでは作成は行わない
					// https://developers.notion.com/reference/property-value-object#rollup-property-values
					// https://developers.notion.com/reference/property-item-object#rollup-property-values
					b.addDerived("rollup", "PropertyValue", e.Text).addFields(
						&interfaceField{name: "rollup", typeName: "Rollup"},
					)
				},
			},
			&objectDocParagraphElement{
				Text: "\nRollup property value objects represent the result of evaluating a rollup described in the \ndatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a rollup cannot be updated directly.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("RollupPropertyValue").comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Rollup\": {\n    \"id\": \"aJ3l\",\n    \"rollup\": {\n      \"type\": \"number\",\n      \"number\": 1234,\n      \"function\": \"sum\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocCalloutElement{
				Body:  "Rollups returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.",
				Title: "Rollup values may not match the Notion UI.",
				Type:  "warning",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getSpecificObject("RollupPropertyValue").comment += "\n\n" + e.Title + "\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text:   "String rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nString rollup property values contain an optional string within the string property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Number rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nNumber rollup property values contain a number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Date rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nDate rollup property values contain a date property value within the date property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Array rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nArray rollup property values contain an array of number, date, or string objects within the results property. \n\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text: "People property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("people", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nPeople property value objects contain an array of user objects within the people property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("PeoplePropertyValue").addFields(&field{
						name:     "people",
						typeCode: jen.Id("Users"),
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"Owners\": {\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"3e01cdb8-6131-4a85-8d83-67102c0fb98c\"\n      },\n      {\n        \"object\": \"user\",\n        \"id\": \"b32c006a-2898-45bb-abd2-de095f354592\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocCalloutElement{
				Body:  "The [Retrieve a page](https://developers.notion.com/reference/retrieve-a-page) endpoint can’t be guaranteed to return more than 25 people per `people` page property. If a `people` page property includes more than 25 people, then you can use the\u00a0[Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `people` property to get a complete list of people.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getSpecificObject("PeoplePropertyValue").comment += "\n\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text: "Files property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("files", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nFile property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("FilesPropertyValue").addFields(&field{
						name:     "files",
						typeCode: jen.Id("Files"),
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Files\": {\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": {\n          \t\"url\": \"https://website.domain/images/space.png\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocCalloutElement{
				Body:  "Although we do not support uploading files, if you pass a `file` object containing a file hosted by Notion, it will remain one of the files. To remove any file, just do not pass it in the update response.",
				Title: "When updating a file property, the value will be overwritten by the array of files passed.",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) {
					b.getSpecificObject("FilesPropertyValue").comment += "\n\n" + e.Title + "\n" + e.Body
				},
			},
			&objectDocHeadingElement{
				Text: "Checkbox property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("checkbox", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nCheckbox property value objects contain a boolean within the checkbox property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("CheckboxPropertyValue").addFields(&field{
						name:     "checkbox",
						typeCode: jen.Bool(), // null や undefined は通らない
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Done?\": {\n    \"checkbox\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"RirO\": {\n    \"checkbox\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "URL property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("url", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nURL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("UrlPropertyValue").addFields(&field{
						name:     "url",
						typeCode: NullString,
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Website\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"<tdn\": {\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Email property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("email", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nEmail property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("EmailPropertyValue").addFields(&field{
						name:     "email",
						typeCode: NullString,
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's Contact\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"}=RV\": {\n    \"email\": \"hello@test.com\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Phone number property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("phone_number", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nPhone number property value objects contain a string within the phone_number property. No structure is enforced.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("PhoneNumberPropertyValue").addFields(&field{
						name:     "phone_number",
						typeCode: NullString,
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's No.\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"_A<p\": {\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Created time property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("created_time", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of created_time cannot be updated. See the Property Item Object to see how these values are returned. \n\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("CreatedTimePropertyValue").addFields(&field{
						name:     "created_time",
						typeCode: jen.Id("ISO8601String"),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Created by property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("created_by", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated by property value objects contain a user object within the created_by property. The user object describes the user who created this page. The value of created_by cannot be updated. See the Property Item Object to see how these values are returned. \n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("CreatedByPropertyValue").addFields(&field{
						name:     "created_by",
						typeCode: jen.Id("PartialUser"),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Last edited time property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("last_edited_time", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\"). The value of last_edited_time cannot be updated. See the Property Item Object to see how these values are returned. \n",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("LastEditedTimePropertyValue").addFields(&field{
						name:     "last_edited_time",
						typeCode: jen.Id("ISO8601String"),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text: "Last edited by property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("last_edited_by", "PropertyValue", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page. The value of last_edited_by cannot be updated. See the Property Item Object to see how these values are returned.",
				output: func(e *objectDocParagraphElement, b *builder) {
					b.getSpecificObject("LastEditedByPropertyValue").addFields(&interfaceField{
						name:     "last_edited_by",
						typeName: "User",
						comment:  e.Text,
					})
				},
			},
		},
	})
}
