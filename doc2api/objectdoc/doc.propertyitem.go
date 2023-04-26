package objectdoc

import (
	"github.com/dave/jennifer/jen"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/property-item-object",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "A property_item object describes the identifier, type, and value of a page property. It's returned from the Retrieve a page property item \n",
				output: func(e *objectDocParagraphElement, b *builder) {
					union := b.addUnionToGlobalIfNotExists("PropertyItemOrPropertyItemPagination", "object")
					b.addAbstractObject("PropertyItem", "type", e.Text, addList(), addMap()).addToUnion(union)
				},
			},
			&objectDocHeadingElement{
				Text:   "All property items",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text: "\nEach page property item object contains the following keys. In addition, it will contain a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[abstractObject](b, "PropertyItem").fieldsComment = e.Text
				},
			},
			&objectDocParametersElement{{
				Property:     "object",
				Type:         `"property_item"`,
				Description:  `Always "property_item".`,
				ExampleValue: `"property_item"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[abstractObject](b, "PropertyItem").addFields(e.asFixedStringField())
				},
			}, {
				Property:     "id",
				Type:         "string",
				Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
				ExampleValue: "\"f%5C%5C%3Ap\"",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[abstractObject](b, "PropertyItem").addFields(e.asField(jen.String()))
				},
			}, {
				Property:     "type",
				Type:         "string (enum)",
				Description:  `Type of the property. Possible values are "rich_text", "number", "select", "multi_select", "date", "formula", "relation", "rollup", "title", "people", "files", "checkbox", "url", "email", "phone_number", "created_time", "created_by", "last_edited_time", and "last_edited_by".`,
				ExampleValue: `"rich_text"`,
				output: func(e *objectDocParameter, b *builder) {
					// 各バリアントで定義
				},
			}},
			&objectDocHeadingElement{
				Text: "Paginated property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					// doc.pagination.go で作成
				},
			},
			&objectDocParagraphElement{
				Text: "\nThe title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information. ",
				output: func(e *objectDocParagraphElement, b *builder) {
					// TODO 良い名前
					b.addAbstractObject("PaginatedPropertyInfo", "type", e.Text).addFields(
						&field{name: "id", typeCode: jen.String()},
					)
					for _, derived := range []string{"title", "rich_text", "relation", "people"} {
						b.addDerived(derived, "PaginatedPropertyInfo", "").addFields(
							&field{name: derived, typeCode: jen.Struct()},
						)
					}
					b.addDerived("rollup", "PaginatedPropertyInfo", "undocumented").addFields(
						&interfaceField{name: "rollup", typeName: "Rollup"},
					)
				},
			},
			&objectDocParametersElement{{
				Property:     "object",
				Type:         `"list"`,
				Description:  `Always "list".`,
				ExampleValue: `"list"`,
				output:       func(e *objectDocParameter, b *builder) {},
			}, {
				Property:     "type",
				Type:         `"property_item"`,
				Description:  `Always "property_item".`,
				ExampleValue: `"property_item"`,
				output:       func(e *objectDocParameter, b *builder) {},
			}, {
				Property:     "results",
				Type:         "list",
				Description:  "List of property_item objects.",
				ExampleValue: "[{\"object\": \"property_item\", \"id\": \"vYdV\", \"type\": \"relation\", \"relation\": { \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"}}... ]",
				output:       func(e *objectDocParameter, b *builder) {},
			}, {
				Property:     "property_item",
				Type:         "object",
				Description:  "A property_item object that describes the property.",
				ExampleValue: `{"id": "title", "next_url": null, "type": "title", "title": {}}`,
				output:       func(e *objectDocParameter, b *builder) {},
			}, {
				Property:     "next_url",
				Type:         "string or null",
				Description:  "The URL the user can request to get the next page of results.",
				ExampleValue: `"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/vYdV?start_cursor=LYxaUO&page_size=25"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[abstractObject](b, "PaginatedPropertyInfo").addFields(e.asField(jen.Id("*").String()))
				},
			}},
			&objectDocHeadingElement{
				Text: "Title property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("title", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nTitle property value objects contain an array of rich text objects within the title property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					// ドキュメントには "array of rich text" と書いてあるが間違い
					getSymbol[concreteObject](b, "TitlePropertyItem").addFields(
						&interfaceField{name: "title", typeName: "RichText", comment: e.Text},
					)
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Name\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"title\",\n        \"type\": \"title\",\n        \"title\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"The title\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"The title\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"title\",\n      \"next_url\": null,\n      \"type\": \"title\",\n      \"title\": {}\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Rich Text property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("rich_text", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nRich Text property value objects contain an array of rich text objects within the rich_text property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					// ドキュメントには "array of rich text" と書いてあるが間違い
					getSymbol[concreteObject](b, "RichTextPropertyItem").addFields(
						&interfaceField{name: "rich_text", typeName: "RichText", comment: e.Text},
					)
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Details\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"Some more text with \",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"Some more text with \",\n          \"href\": null\n        }\n      },\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"fun formatting\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": true,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"fun formatting\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"NVv^\",\n      \"next_url\": null,\n      \"type\": \"rich_text\",\n      \"rich_text\": {}\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Number property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("number", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber property value objects contain a number within the number property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "NumberPropertyItem").addFields(&field{name: "number", typeCode: NullFloat, comment: e.Text})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Quantity\": {\n    \"object\": \"property_item\",\n    \"id\": \"XpXf\",\n    \"type\": \"number\",\n    \"number\": 1234\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Select property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("select", "PropertyItem", e.Text).addFields(
						&field{name: "select", typeCode: jen.Id("Option")},
					)
					b.addDerived("status", "PropertyItem", "undocumented").addFields(
						&field{name: "status", typeCode: jen.Id("Option")},
					)
				},
			},
			&objectDocParagraphElement{
				Text:   "\nSelect property value objects contain the following data within the select property:",
				output: func(e *objectDocParagraphElement, b *builder) {}, // Optionで共通化
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
				Code:     "{\n  \"Option\": {\n    \"object\": \"property_item\",\n    \"id\": \"%7CtzR\",\n    \"type\": \"select\",\n    \"select\": {\n      \"id\": \"64190ec9-e963-47cb-bc37-6a71d6b71206\",\n      \"name\": \"Option 1\",\n      \"color\": \"orange\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Multi-select property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("multi_select", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nMulti-select property value objects contain an array of multi-select option values within the multi_select property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "MultiSelectPropertyItem").addFields(&field{
						name:     "multi_select",
						typeCode: jen.Index().Id("Option"),
						comment:  e.Text,
					})
				},
			},
			&objectDocHeadingElement{
				Text:   "Multi-select option values",
				output: func(e *objectDocHeadingElement, b *builder) {},
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
				Code:     "{\n  \"Tags\": {\n    \"object\": \"property_item\",\n    \"id\": \"z%7D%5C%3C\",\n    \"type\": \"multi_select\",\n    \"multi_select\": [\n      {\n        \"id\": \"91e6959e-7690-4f55-b8dd-d3da9debac45\",\n        \"name\": \"A\",\n        \"color\": \"orange\"\n      },\n      {\n        \"id\": \"2f998e2d-7b1c-485b-ba6b-5e6a815ec8f5\",\n        \"name\": \"B\",\n        \"color\": \"purple\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Date property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("date", "PropertyItem", e.Text, addSpecificField(true))
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate property value objects contain the following data within the date property:",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "DatePropertyItem").comment += e.Text
				},
			},
			&objectDocParametersElement{{
				Property:     "start",
				Type:         "string (ISO 8601 date and time)",
				Description:  "An ISO 8601 format date, with optional time.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DatePropertyItem").typeSpecificObject.addFields(e.asField(jen.Id("ISO8601String")))
				},
			}, {
				Property:     "end",
				Type:         "string (optional, ISO 8601 date and time)",
				Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DatePropertyItem").typeSpecificObject.addFields(e.asField(jen.Op("*").Id("ISO8601String")))
				},
			}, {
				Property:     "time_zone",
				Type:         "string (optional, enum)",
				Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
				ExampleValue: `"America/Los_Angeles"`,
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[concreteObject](b, "DatePropertyItem").typeSpecificObject.addFields(e.asField(NullString))
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipment Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"i%3Ahj\",\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\",\n      \"end\": null,\n      \"time_zone\": null\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("formula", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nFormula property value objects represent the result of evaluating a formula described in the \ndatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "FormulaPropertyItem").addFields(&interfaceField{
						name:     "formula",
						typeName: "Formula",
						comment:  e.Text,
					})
				},
			},
			&objectDocParametersElement{{
				Description: "The type of the formula result. Possible values are \"string\", \"number\", \"boolean\", and \"date\".",
				Property:    "type",
				Type:        "string (enum)",
				output:      func(e *objectDocParameter, b *builder) {},
			}},
			&objectDocHeadingElement{
				Text:   "String formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nString formula property values contain an optional string within the string property.\n\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Number formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nNumber formula property values contain an optional number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Boolean formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nBoolean formula property values contain a boolean within the boolean property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocHeadingElement{
				Text:   "Date formula property values",
				output: func(e *objectDocHeadingElement, b *builder) {},
			},
			&objectDocParagraphElement{
				Text:   "\nDate formula property values contain an optional date property value within the date property.",
				output: func(e *objectDocParagraphElement, b *builder) {},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Formula\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"formula\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					b.addUnmarshalTest("PropertyItemMap", e.Code)
				},
			}}},
			&objectDocHeadingElement{
				Text: "Relation property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("relation", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nRelation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "RelationPropertyItem").addFields(&field{
						name:     "relation",
						typeCode: jen.Id("PageReference"), // 1つ
						comment:  e.Text,
					})
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Project\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"vYdV\",\n      \"next_url\": null,\n      \"type\": \"relation\",\n      \"relation\": {}\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					// PropertyValueとPropertyItemで共通の Rollup を使う
					// PropertyValue のArray Rollupにはfunction が無いなど不正確であるため、
					// 比較的ドキュメントが充実しているこちらで作成を行う
					// https://developers.notion.com/reference/property-value-object#rollup-property-values
					// https://developers.notion.com/reference/property-item-object#rollup-property-values
					b.addDerived("rollup", "PropertyItem", e.Text)
					b.addAbstractObject("Rollup", "type", "")
				},
			},
			&objectDocParagraphElement{
				Text: "\nRollup property value objects represent the result of evaluating a rollup described in the \ndatabase's properties. The property is returned as a list object of type property_item with a list of relation items used to computed the rollup under results. \n\nA rollup property item is also returned under the property_type key that describes the rollup aggregation and computed result. \n\nIn order to avoid timeouts, if the rollup has a with a large number of aggregations or properties the endpoint returns a next_cursor value that is used to determinate the aggregation value so far for the subset of relations that have been paginated through. \n\nOnce has_more is false, then the final rollup value is returned.  See the Pagination documentation for more information on pagination in the Notion API. \n\nComputing the values of following aggregations are not supported. Instead the endpoint returns a list of property_item objects for the rollup:\n* show_unique (Show unique values)\n* unique (Count unique values)\n* median(Median)",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "RollupPropertyItem").addFields(&interfaceField{
						name:     "rollup",
						typeName: "Rollup",
						comment:  e.Text,
					})
				},
			},
			&objectDocParametersElement{{
				Description: "The type of rollup. Possible values are \"number\", \"date\", \"array\", \"unsupported\" and \"incomplete\".",
				Property:    "type",
				Type:        "string (enum)",
				output:      func(e *objectDocParameter, b *builder) {},
			}, {
				Description: "Describes the aggregation used. \nPossible values include: count,  count_values,  empty,  not_empty,  unique,  show_unique,  percent_empty,  percent_not_empty,  sum,  average,  median,  min,  max,  range,  earliest_date,  latest_date,  date_range,  checked,  unchecked,  percent_checked,  percent_unchecked,  count_per_group,  percent_per_group,  show_original",
				Property:    "function",
				Type:        "string (enum)",
				output: func(e *objectDocParameter, b *builder) {
					getSymbol[abstractObject](b, "Rollup").addFields(e.asField(jen.String()))
				},
			}},
			&objectDocHeadingElement{
				Text: "Number rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("number", "Rollup", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber rollup property values contain a number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "NumberRollup").addFields(
						&field{name: "number", typeCode: NullFloat},
					).comment += e.Text
				},
			},
			&objectDocHeadingElement{
				Text: "Date rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("date", "Rollup", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate rollup property values contain a date property value within the date property.\n",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "DateRollup").addFields(
						&field{name: "date", typeCode: jen.Id("DatePropertyItem")},
					).comment += e.Text
				},
			},
			&objectDocHeadingElement{
				Text: "Array rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("array", "Rollup", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nArray rollup property values contain an array of property_item objects within the results property. \n",
				output: func(e *objectDocParagraphElement, b *builder) {
					// ドキュメントには array of property_item とあるが、
					// type="rich_text"の場合に入る値などから
					// array of property_value が正しいと判断している
					getSymbol[concreteObject](b, "ArrayRollup").addFields(
						&field{name: "array", typeCode: jen.Id("PropertyValueList")},
					).comment += e.Text
				},
			},
			&objectDocHeadingElement{
				Text: "Incomplete rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					// TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nRollups with an aggregation with more than one page of aggregated results will return a rollup object of type \"incomplete\". To obtain the final value paginate through the next values in the rollup using the next_cursor or next_url property. ",
				output: func(e *objectDocParagraphElement, b *builder) {
					// TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Rollup\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }...\n    ],\t\n    \"next_cursor\": \"1QaTunT5\",\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"y}~p\",\n      \"next_url\": \"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/y%7D~p?start_cursor=1QaTunT5&page_size=25\",\n      \"type\": \"rollup\",\n      \"rollup\": {\n        \"function\": \"sum\",\n        \"type\": \"incomplete\",\n        \"incomplete\": {}\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "People property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("people", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nPeople property value objects contain an array of user objects within the people property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "PeoplePropertyItem").addFields(
						&interfaceField{name: "people", typeName: "User"},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Owners\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"people\",\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"285e5768-3fdc-4742-ab9e-125f9050f3b8\",\n        \"name\": \"Example Avo\",\n        \"avatar_url\": null,\n        \"type\": \"person\",\n        \"person\": {\n          \"email\": \"avo@example.com\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Files property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("files", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nFile property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "FilesPropertyItem").addFields(
						&field{name: "files", typeCode: jen.Id("FileList")},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Files\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"files\",\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": \"https://website.domain/images/space.png\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Checkbox property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("checkbox", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nCheckbox property value objects contain a boolean within the checkbox property.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "CheckboxPropertyItem").addFields(
						&field{name: "checkbox", typeCode: jen.Bool()},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Done?\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"checkbox\",\n    \"checkbox\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "URL property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("url", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nURL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "UrlPropertyItem").addFields(
						&field{name: "url", typeCode: NullString}, // null
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Website\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"url\",\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Email property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("email", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nEmail property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "EmailPropertyItem").addFields(
						&field{name: "email", typeCode: NullString}, // null
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's Contact\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"email\",\n    \"email\": \"hello@test.com\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Phone number property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("phone_number", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nPhone number property value objects contain a string within the phone_number property. No structure is enforced.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "PhoneNumberPropertyItem").addFields(
						&field{name: "phone_number", typeCode: NullString}, // null
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's No.\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"phone_number\",\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Created time property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("created_time", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "CreatedTimePropertyItem").addFields(
						&field{name: "created_time", typeCode: jen.Id("ISO8601String")},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Created Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"create_time\",\n  \t\"created_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Created by property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("created_by", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated by property value objects contain a user object within the created_by property. The user object describes the user who created this page.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "CreatedByPropertyItem").addFields(
						&interfaceField{name: "created_by", typeName: "User"},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Created By\": {\n    \"created_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"dsEa\": {\n    \"created_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
			&objectDocHeadingElement{
				Text: "Last edited time property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("last_edited_time", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "LastEditedTimePropertyItem").addFields(
						&field{name: "last_edited_time", typeCode: jen.Id("ISO8601String")},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Last Edited Time\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					// TODO typeを持たなくてもアンマーシャルできる仕組み
				},
			}, {
				Code:     "{\n  \"as0w\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) {
					// TODO typeを持たなくてもアンマーシャルできる仕組み
				},
			}}},
			&objectDocHeadingElement{
				Text: "Last edited by property values",
				output: func(e *objectDocHeadingElement, b *builder) {
					b.addDerived("last_edited_by", "PropertyItem", e.Text)
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page.",
				output: func(e *objectDocParagraphElement, b *builder) {
					getSymbol[concreteObject](b, "LastEditedByPropertyItem").addFields(
						&interfaceField{name: "last_edited_by", typeName: "User"},
					).comment += e.Text
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Last Edited By\": {\n    \"last_edited_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}, {
				Code:     "{\n  \"as12\": {\n    \"last_edited_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) {},
			}}},
		},
	})
}
