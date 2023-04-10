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
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.add(&abstractObject{name: "PropertyItem", comment: e.Text})
					b.add(&abstractObject{name: "PropertyItemOrPropertyItemPagination", specifiedBy: "object"})
					// TODO: PropertyItemOrPropertyItemPaginationの派生としてPropertyItemを登録
					// b.getAbstractObject("PropertyItemOrPropertyItemPagination").addVariant(b.getAbstractObject("PropertyItem"))
					return nil
				},
			},
			&objectDocHeadingElement{
				Text:   "All property items",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text: "\nEach page property item object contains the following keys. In addition, it will contain a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("PropertyItem").fieldsComment = e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "object",
				Type:         `"property_item"`,
				Description:  `Always "property_item".`,
				ExampleValue: `"property_item"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("PropertyItem").addField(e.asFixedStringField())
					return nil
				},
			}, {
				Property:     "id",
				Type:         "string",
				Description:  "Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.\n\nThe id may be used in place of name when creating or updating pages.",
				ExampleValue: "\"f%5C%5C%3Ap\"",
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("PropertyItem").addField(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "type",
				Type:         "string (enum)",
				Description:  "Type of the property. Possible values are \"rich_text\", \"number\", \"select\", \"multi_select\", \"date\", \"formula\", \"relation\", \"rollup\", \"title\", \"people\", \"files\", \"checkbox\", \"url\", \"email\", \"phone_number\", \"created_time\", \"created_by\", \"last_edited_time\", and \"last_edited_by\".",
				ExampleValue: `"rich_text"`,
				output: func(e *objectDocParameter, b *builder) error {
					return nil // 各バリアントで定義
				},
			}},
			&objectDocHeadingElement{
				Text: "Paginated property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.add(&specificObject{name: "PaginatedPropertyValue", comment: e.Text})
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nThe title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PaginatedPropertyValue").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "object",
				Type:         `"list"`,
				Description:  `Always "list".`,
				ExampleValue: `"list"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PaginatedPropertyValue").addField(e.asFixedStringField())
					return nil
				},
			}, {
				Property:     "type",
				Type:         `"property_item"`,
				Description:  `Always "property_item".`,
				ExampleValue: `"property_item"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PaginatedPropertyValue").addField(e.asFixedStringField())
					return nil
				},
			}, {
				Property:     "results",
				Type:         "list",
				Description:  "List of property_item objects.",
				ExampleValue: "[{\"object\": \"property_item\", \"id\": \"vYdV\", \"type\": \"relation\", \"relation\": { \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"}}... ]",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PaginatedPropertyValue").addField(e.asField(jen.Index().Id("PropertyItem")))
					return nil
				},
			}, {
				Property:     "property_item",
				Type:         "object",
				Description:  "A property_item object that describes the property.",
				ExampleValue: `{"id": "title", "next_url": null, "type": "title", "title": {}}`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PaginatedPropertyValue").addField(e.asField(jen.Id("PropertyItem")))
					return nil
				},
			}, {
				Property:     "next_url",
				Type:         "string or null",
				Description:  "The URL the user can request to get the next page of results.",
				ExampleValue: `"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/vYdV?start_cursor=LYxaUO&page_size=25"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("PaginatedPropertyValue").addField(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text: "Title property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "TitlePropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "title"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nTitle property value objects contain an array of rich text objects within the title property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("TitlePropertyItem").addField(&field{name: "title", typeCode: jen.Id("RichTextArray"), comment: e.Text})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Name\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"title\",\n        \"type\": \"title\",\n        \"title\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"The title\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"The title\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"title\",\n      \"next_url\": null,\n      \"type\": \"title\",\n      \"title\": {}\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Rich Text property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "RichTextPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "rich_text"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nRich Text property value objects contain an array of rich text objects within the rich_text property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RichTextPropertyItem").addField(&field{name: "rich_text", typeCode: jen.Id("RichTextArray"), comment: e.Text})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Details\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"Some more text with \",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": false,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"Some more text with \",\n          \"href\": null\n        }\n      },\n      {\n        \"object\": \"property_item\",\n        \"id\": \"NVv%5E\",\n        \"type\": \"rich_text\",\n        \"rich_text\": {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"fun formatting\",\n            \"link\": null\n          },\n          \"annotations\": {\n            \"bold\": false,\n            \"italic\": true,\n            \"strikethrough\": false,\n            \"underline\": false,\n            \"code\": false,\n            \"color\": \"default\"\n          },\n          \"plain_text\": \"fun formatting\",\n          \"href\": null\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"NVv^\",\n      \"next_url\": null,\n      \"type\": \"rich_text\",\n      \"rich_text\": {}\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Number property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "NumberPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "number"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber property value objects contain a number within the number property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("NumberPropertyItem").addField(&field{name: "number", typeCode: jen.Float64(), comment: e.Text})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Quantity\": {\n    \"object\": \"property_item\",\n    \"id\": \"XpXf\",\n    \"type\": \"number\",\n    \"number\": 1234\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Select property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "SelectPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "select"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nSelect property value objects contain the following data within the select property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("SelectPropertyItem").addField(&field{name: "select", typeCode: jen.Id("SelectPropertyItemData"), comment: e.Text})
					b.add(&specificObject{name: "SelectPropertyItemData", comment: e.Text})
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.\n\nWhen updating a select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyItemData").addField(e.asField(jen.Qual("github.com/google/uuid", "UUID")))
					return nil
				},
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.\n\nIf the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyItemData").addField(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyItemData").addField(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Option\": {\n    \"object\": \"property_item\",\n    \"id\": \"%7CtzR\",\n    \"type\": \"select\",\n    \"select\": {\n      \"id\": \"64190ec9-e963-47cb-bc37-6a71d6b71206\",\n      \"name\": \"Option 1\",\n      \"color\": \"orange\"\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Multi-select property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "MultiSelectPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "multi_select"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nMulti-select property value objects contain an array of multi-select option values within the multi_select property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyItem").addField(&field{
						name:     "multi_select",
						typeCode: jen.Index().Id("MultiSelectPropertyItemData"),
						comment:  e.Text,
					})
					b.add(&specificObject{name: "MultiSelectPropertyItemData", comment: e.Text})
					return nil
				},
			},
			&objectDocHeadingElement{
				Text:   "Multi-select option values",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParametersElement{{
				Property:     "id",
				Type:         "string (UUIDv4)",
				Description:  "ID of the option.\n\nWhen updating a multi-select property, you can use either name or id.",
				ExampleValue: `"b3d773ca-b2c9-47d8-ae98-3c2ce3b2bffb"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyItemData").addField(e.asField(jen.Qual("github.com/google/uuid", "UUID")))
					return nil
				},
			}, {
				Property:     "name",
				Type:         "string",
				Description:  "Name of the option as it appears in Notion.\n\nIf the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyItemData").addField(e.asField(jen.String()))
					return nil
				},
			}, {
				Property:     "color",
				Type:         "string (enum)",
				Description:  "Color of the option. Possible values are: \"default\", \"gray\", \"brown\", \"red\", \"orange\", \"yellow\", \"green\", \"blue\", \"purple\", \"pink\". Defaults to \"default\".\n\nNot currently editable.",
				ExampleValue: `"red"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyItemData").addField(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Tags\": {\n    \"object\": \"property_item\",\n    \"id\": \"z%7D%5C%3C\",\n    \"type\": \"multi_select\",\n    \"multi_select\": [\n      {\n        \"id\": \"91e6959e-7690-4f55-b8dd-d3da9debac45\",\n        \"name\": \"A\",\n        \"color\": \"orange\"\n      },\n      {\n        \"id\": \"2f998e2d-7b1c-485b-ba6b-5e6a815ec8f5\",\n        \"name\": \"B\",\n        \"color\": \"purple\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Date property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "DatePropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "date"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate property value objects contain the following data within the date property:",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DatePropertyItem").addField(&field{name: "date", typeCode: jen.Id("DatePropertyItemData"), comment: e.Text})
					b.add(&specificObject{name: "DatePropertyItemData", comment: e.Text})
					return nil
				},
			},
			&objectDocParametersElement{{
				Property:     "start",
				Type:         "string (ISO 8601 date and time)",
				Description:  "An ISO 8601 format date, with optional time.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatePropertyItemData").addField(e.asField(jen.Id("ISO8601String")))
					return nil
				},
			}, {
				Property:     "end",
				Type:         "string (optional, ISO 8601 date and time)",
				Description:  "An ISO 8601 formatted date, with optional time. Represents the end of a date range.\n\nIf null, this property's date value is not a range.",
				ExampleValue: `"2020-12-08T12:00:00Z"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatePropertyItemData").addField(e.asField(jen.Id("ISO8601String")))
					return nil
				},
			}, {
				Property:     "time_zone",
				Type:         "string (optional, enum)",
				Description:  "Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.\n\nWhen time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.\n\nIf null, time zone information will be contained in UTC offsets in start and end.",
				ExampleValue: `"America/Los_Angeles"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DatePropertyItemData").addField(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipment Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"i%3Ahj\",\n    \"type\": \"date\",\n    \"date\": {\n      \"start\": \"2021-05-11T11:00:00.000-04:00\",\n      \"end\": null,\n      \"time_zone\": null\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "FormulaPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "formula"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nFormula property value objects represent the result of evaluating a formula described in the \ndatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("FormulaPropertyItem").addField(&field{
						name:     "formula",
						typeCode: jen.Id("Formula"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocParametersElement{{
				Description: "The type of the formula result. Possible values are \"string\", \"number\", \"boolean\", and \"date\".",
				Property:    "type",
				Type:        "string (enum)",
				output:      func(e *objectDocParameter, b *builder) error { return nil },
			}},
			&objectDocHeadingElement{
				Text:   "String formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\nString formula property values contain an optional string within the string property.\n\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Number formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\nNumber formula property values contain an optional number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Boolean formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\nBoolean formula property values contain a boolean within the boolean property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Date formula property values",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\nDate formula property values contain an optional date property value within the date property.",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Formula\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"formula\",\n    \"formula\": {\n      \"type\": \"number\",\n      \"number\": 1234\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Relation property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "RelationPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "relation"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nRelation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RelationPropertyItem").addField(&field{
						name:     "relation",
						typeCode: jen.Index().Id("PageReference"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Project\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }\n    ],\n    \"next_cursor\": null,\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"vYdV\",\n      \"next_url\": null,\n      \"type\": \"relation\",\n      \"relation\": {}\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "RollupPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "rollup"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nRollup property value objects represent the result of evaluating a rollup described in the \ndatabase's properties. The property is returned as a list object of type property_item with a list of relation items used to computed the rollup under results. \n\nA rollup property item is also returned under the property_type key that describes the rollup aggregation and computed result. \n\nIn order to avoid timeouts, if the rollup has a with a large number of aggregations or properties the endpoint returns a next_cursor value that is used to determinate the aggregation value so far for the subset of relations that have been paginated through. \n\nOnce has_more is false, then the final rollup value is returned.  See the Pagination documentation for more information on pagination in the Notion API. \n\nComputing the values of following aggregations are not supported. Instead the endpoint returns a list of property_item objects for the rollup:\n* show_unique (Show unique values)\n* unique (Count unique values)\n* median(Median)",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RollupPropertyItem").addField(&field{
						name:     "rollup",
						typeCode: jen.Id("Rollup"),
						comment:  e.Text,
					})
					return nil
				},
			},
			&objectDocParametersElement{{
				Description: "The type of rollup. Possible values are \"number\", \"date\", \"array\", \"unsupported\" and \"incomplete\".",
				Property:    "type",
				Type:        "string (enum)",
				output:      func(e *objectDocParameter, b *builder) error { return nil },
			}, {
				Description: "Describes the aggregation used. \nPossible values include: count,  count_values,  empty,  not_empty,  unique,  show_unique,  percent_empty,  percent_not_empty,  sum,  average,  median,  min,  max,  range,  earliest_date,  latest_date,  date_range,  checked,  unchecked,  percent_checked,  percent_unchecked,  count_per_group,  percent_per_group,  show_original",
				Property:    "function",
				Type:        "string (enum)",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // TODO
				},
			}},
			&objectDocHeadingElement{
				Text: "Number rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nNumber rollup property values contain a number within the number property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Date rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nDate rollup property values contain a date property value within the date property.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Array rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nArray rollup property values contain an array of property_item objects within the results property. \n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocHeadingElement{
				Text: "Incomplete rollup property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocParagraphElement{
				Text: "\nRollups with an aggregation with more than one page of aggregated results will return a rollup object of type \"incomplete\". To obtain the final value paginate through the next values in the rollup using the next_cursor or next_url property. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Rollup\": {\n    \"object\": \"list\",\n    \"results\": [\n      {\n        \"object\": \"property_item\",\n        \"id\": \"vYdV\",\n        \"type\": \"relation\",\n        \"relation\": {\n          \"id\": \"535c3fb2-95e6-4b37-a696-036e5eac5cf6\"\n        }\n      }...\n    ],\t\n    \"next_cursor\": \"1QaTunT5\",\n    \"has_more\": true,\n    \"type\": \"property_item\",\n    \"property_item\": {\n      \"id\": \"y}~p\",\n      \"next_url\": \"http://api.notion.com/v1/pages/0e5235bf86aa4efb93aa772cce7eab71/properties/y%7D~p?start_cursor=1QaTunT5&page_size=25\",\n      \"type\": \"rollup\",\n      \"rollup\": {\n        \"function\": \"sum\",\n        \"type\": \"incomplete\",\n        \"incomplete\": {}\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "People property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "PeoplePropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "people"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nPeople property value objects contain an array of user objects within the people property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Owners\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"people\",\n    \"people\": [\n      {\n        \"object\": \"user\",\n        \"id\": \"285e5768-3fdc-4742-ab9e-125f9050f3b8\",\n        \"name\": \"Example Avo\",\n        \"avatar_url\": null,\n        \"type\": \"person\",\n        \"person\": {\n          \"email\": \"avo@example.com\"\n        }\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Files property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "FilesPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "files"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nFile property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. \"Whole_Earth_Catalog.jpg\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Files\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"files\",\n    \"files\": [\n      {\n        \"type\": \"external\",\n        \"name\": \"Space Wallpaper\",\n        \"external\": \"https://website.domain/images/space.png\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Checkbox property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "CheckboxPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "checkbox"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nCheckbox property value objects contain a boolean within the checkbox property.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Done?\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"checkbox\",\n    \"checkbox\": true\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "URL property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "UrlPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "url"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nURL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. \"http://worrydream.com/EarlyHistoryOfSmalltalk/\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Website\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"url\",\n    \"url\": \"https://notion.so/notiondevs\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Email property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "EmailPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "email"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nEmail property value objects contain a string within the email property. The string describes an email address (i.e. \"hello@example.org\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's Contact\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"email\",\n    \"email\": \"hello@test.com\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Phone number property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "PhoneNumberPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "phone_number"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nPhone number property value objects contain a string within the phone_number property. No structure is enforced.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Shipper's No.\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"phone_number\",\n    \"phone_number\": \"415-000-1111\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Created time property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "CreatedTimePropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "created_time"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Created Time\": {\n    \"object\": \"property_item\",\n    \"id\": \"KpQq\",\n    \"type\": \"create_time\",\n  \t\"created_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Created by property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "CreatedByPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "created_by"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nCreated by property value objects contain a user object within the created_by property. The user object describes the user who created this page.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Created By\": {\n    \"created_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}, {
				Code:     "{\n  \"dsEa\": {\n    \"created_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Last edited time property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "LastEditedTimePropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "last_edited_time"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. \"2020-03-17T19:10:04.968Z\").",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Last Edited Time\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}, {
				Code:     "{\n  \"as0w\": {\n  \t\"last_edited_time\": \"2020-03-17T19:10:04.968Z\"\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Last edited by property values",
				output: func(e *objectDocHeadingElement, b *builder) error {
					cs := &specificObject{
						name:    "LastEditedByPropertyItem",
						comment: e.Text,
						fields:  []coder{&fixedStringField{name: "type", value: "last_edited_by"}},
					}
					b.add(cs)
					b.getAbstractObject("PropertyItem").addVariant(cs)
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nLast edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					return nil // TODO
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "{\n  \"Last Edited By\": {\n    \"last_edited_by\": {\n      \"object\": \"user\",\n      \"id\": \"23345d4f-cf71-4a70-89a5-226c95a6eaae\",\n      \"name\": \"Test User\",\n      \"type\": \"person\",\n      \"person\": {\n        \"email\": \"avo@example.org\"\n      }\n    }\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}, {
				Code:     "{\n  \"as12\": {\n    \"last_edited_by\": {\n\t\t\t\"object\": \"user\",\n\t\t\t\"id\": \"71e95936-2737-4e11-b03d-f174f6f13087\"\n  \t}\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
		},
	})
}
