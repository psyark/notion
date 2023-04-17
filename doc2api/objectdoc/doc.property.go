package objectdoc

import (
	"regexp"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/property-object",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text: "All database objects include a child properties object. This properties object is composed of individual database property objects. These property objects define the database schema and are rendered in the Notion UI as database columns. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addAbstractObject("Property", "type", e.Text).strMapName = "PropertyMap"
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:   "If you’re looking for information about how to use the API to work with database rows, then refer to the [page property values](https://developers.notion.com/reference/property-value-object) documentation. The API treats database rows as pages.",
				Title:  "Database rows",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text: "Every database property object contains the following keys: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("Property").fieldsComment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "id",
				Type:         "string",
				Description:  "An identifier for the property, usually a short string of random letters and symbols.\n\nSome automatically generated property types have special human-readable IDs. For example, all Title properties have an id of \"title\".",
				ExampleValue: `"fy:{"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("Property").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:       "name",
				Type:        "string",
				Description: "The name of the property as it appears in Notion.",
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("Property").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "type",
				Type:         "string (enum)",
				Description:  "The type that controls the behavior of the property. Possible values are: \n\n- \"checkbox\"\n- \"created_by\"\n- \"created_time\"\n- \"date\"\n- \"email\"\n- \"files\"\n- \"formula\"\n- \"last_edited_by\"\n- \"last_edited_time\"\n- \"multi_select\"\n- \"number\"\n- \"people\"\n- \"phone_number\"\n- \"relation\"\n- \"rich_text\"\n- \"rollup\"\n- \"select\"\n- \"status\"\n- \"title\"\n- \"url\"",
				ExampleValue: `"rich_text"`,
				output: func(e *objectDocParameter, b *builder) error {
					for _, m := range regexp.MustCompile(`- "(\w+)"`).FindAllStringSubmatch(e.Description, -1) {
						b.getAbstractObject("Property").addVariant(
							b.addSpecificObject(strcase.UpperCamelCase(m[1])+"Property", "").addFields(
								&fixedStringField{name: "type", value: m[1]},
							),
						)
					}
					return nil
				},
			}},
			&objectDocParagraphElement{
				Text:   "Each database property object also contains a type object. The key of the object is the type of the object, and the value is an object containing type-specific configuration. The following sections detail these type-specific objects along with example property objects for each type. \n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text: "Checkbox",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("CheckboxProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA checkbox database property is rendered in the Notion UI as a column that contains checkboxes. The checkbox type object is empty; there is no additional property configuration. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("CheckboxProperty").addFields(
						&field{name: "checkbox", typeCode: jen.Struct(), comment: e.Text},
					)
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Task complete\": {\n  \"id\": \"BBla\",\n  \"name\": \"Task complete\",\n  \"type\": \"checkbox\",\n  \"checkbox\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Created by ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("CreatedByProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA created by database property is rendered in the Notion UI as a column that contains people mentions of each row's author as values. \n\nThe created_by type object is empty. There is no additional property configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("CreatedByProperty").addFields(
						&field{name: "created_by", typeCode: jen.Struct(), comment: e.Text},
					)
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Created by\": {\n  \"id\": \"%5BJCR\",\n  \"name\": \"Created by\",\n  \"type\": \"created_by\",\n  \"created_by\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Created time",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("CreatedTimeProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA created time database property is rendered in the Notion UI as a column that contains timestamps of when each row was created as values. \n\nThe created_time type object is empty. There is no additional property configuration. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("CreatedTimeProperty").addFields(
						&field{name: "created_time", typeCode: jen.Struct(), comment: e.Text},
					)
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Created time\": {\n  \"id\": \"XcAf\",\n  \"name\": \"Created time\",\n  \"type\": \"created_time\",\n  \"created_time\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Date ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("DateProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA date database property is rendered in the Notion UI as a column that contains date values. \n\nThe date type object is empty; there is no additional configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("DateProperty").addFields(
						&field{name: "date", typeCode: jen.Struct(), comment: e.Text},
					)
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Task due date\" {\n  \"id\": \"AJP%7D\",\n  \"name\": \"Task due date\",\n  \"type\": \"date\",\n  \"date\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Email",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("EmailProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nAn email database property is represented in the Notion UI as a column that contains email values. \n\nThe email type object is empty. There is no additional property configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("EmailProperty").addFields(&field{name: "email", typeCode: jen.Struct()}).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Contact email\": {\n  \"id\": \"oZbC\",\n  \"name\": \"Contact email\",\n  \"type\": \"email\",\n  \"email\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Files",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("FilesProperty").comment = e.Text
					return nil
				},
			},
			&objectDocCalloutElement{
				Body:  "The Notion API does not yet support uploading files to Notion.",
				Title: "",
				Type:  "info",
				output: func(e *objectDocCalloutElement, b *builder) error {
					b.getSpecificObject("FilesProperty").comment += "\n" + e.Body
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "A files database property is rendered in the Notion UI as a column that has values that are either files uploaded directly to Notion or external links to files. The files type object is empty; there is no additional configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("FilesProperty").addFields(&field{name: "files", typeCode: jen.Struct()}).comment += "\n" + e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Product image\": {\n  \"id\": \"pb%3E%5B\",\n  \"name\": \"Product image\",\n  \"type\": \"files\",\n  \"files\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Formula",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("FormulaProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA formula database property is rendered in the Notion UI as a column that contains values derived from a provided expression. \n\nThe formula type object defines the expression in the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("FormulaProperty").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "expression",
				Type:         "string",
				Description:  "The formula that is used to compute the values for this property. \n\nRefer to the Notion help center for information about formula syntax.",
				ExampleValue: `"prop(\"Price\") * 2"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("FormulaProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Updated price\": {\n  \"id\": \"YU%7C%40\",\n  \"name\": \"Updated price\",\n  \"type\": \"formula\",\n  \"formula\": {\n    \"expression\": \"prop(\\\"Price\\\") * 2\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Last edited by ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("LastEditedByProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA last edited by database property is rendered in the Notion UI as a column that contains people mentions of the person who last edited each row as values. \n\nThe last_edited_by type object is empty. There is no additional property configuration.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("LastEditedByProperty").addFields(
						&field{name: "last_edited_by", typeCode: jen.Struct()},
					).comment += e.Text
					return nil
				},
			},
			&objectDocHeadingElement{
				Text: "Last edited time",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("LastEditedTimeProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA last edited time database property is rendered in the Notion UI as a column that contains timestamps of when each row was last edited as values. \n\nThe last_edited_time type object is empty. There is no additional property configuration. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("LastEditedTimeProperty").addFields(
						&field{name: "last_edited_time", typeCode: jen.Struct()},
					).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Last edited time\": {\n  \"id\": \"jGdo\",\n  \"name\": \"Last edited time\",\n  \"type\": \"last_edited_time\",\n  \"last_edited_time\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Multi-select ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("MultiSelectProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA multi-select database property is rendered in the Notion UI as a column that contains values from a range of options. Each row can contain one or multiple options. \n\nThe multi_select type object includes an array of options objects. Each option object details settings for the option, indicating the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("MultiSelectProperty").comment += e.Text
					// TODO 似たようなオブジェクトの共通化
					b.getSpecificObject("MultiSelectProperty").typeObject.addFields(
						&field{name: "options", typeCode: jen.Index().Id("MultiSelectPropertyOption")},
					)
					b.addSpecificObject("MultiSelectPropertyOption", "")
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "color",
				Type:         "string (enum)",
				Description:  "The color of the option as rendered in the Notion UI. Possible values include: \n\n- blue\n- brown\n- default\n- gray\n- green\n- orange\n- pink\n- purple\n- red\n- yellow",
				ExampleValue: `"blue"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyOption").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "id",
				Type:         "string",
				Description:  "An identifier for the option, which does not change if the name is changed. An id is sometimes, but not always, a UUID.",
				ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyOption").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Field:        "name",
				Type:         "string",
				Description:  "The name of the option as it appears in Notion.\n\nNote: Commas (\",\") are not valid for multi-select properties.",
				ExampleValue: `"Fruit"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("MultiSelectPropertyOption").addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Store availability\": {\n  \"id\": \"flsb\",\n  \"name\": \"Store availability\",\n  \"type\": \"multi_select\",\n  \"multi_select\": {\n    \"options\": [\n      {\n        \"id\": \"5de29601-9c24-4b04-8629-0bca891c5120\",\n        \"name\": \"Duc Loi Market\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"385890b8-fe15-421b-b214-b02959b0f8d9\",\n        \"name\": \"Rainbow Grocery\",\n        \"color\": \"gray\"\n      },\n      {\n        \"id\": \"72ac0a6c-9e00-4e8c-80c5-720e4373e0b9\",\n        \"name\": \"Nijiya Market\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"9556a8f7-f4b0-4e11-b277-f0af1f8c9490\",\n        \"name\": \"Gus's Community Market\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Number",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("NumberProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA number database property is rendered in the Notion UI as a column that contains numeric values. The number type object contains the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("NumberProperty").comment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "format",
				Type:         "string (enum)",
				Description:  "The way that the number is displayed in Notion. Potential values include: \n\n- argentine_peso\n- baht\n- canadian_dollar\n- chilean_peso\n- colombian_peso\n- danish_krone\n- dirham\n- dollar\n- euro\n- forint\n- franc\n- hong_kong_dollar\n- koruna\n- krona\n- leu\n- lira\n-  mexican_peso\n- new_taiwan_dollar\n- new_zealand_dollar\n- norwegian_krone\n- number\n- number_with_commas\n- percent\n- philippine_peso\n- pound \n- rand\n- real\n- ringgit\n- riyal\n- ruble\n- rupee\n- rupiah\n- shekel\n- singapore_dollar\n- uruguayan_peso\n- yen,\n- yuan\n- won\n- zloty",
				ExampleValue: `"percent"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("NumberProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Price\"{\n  \"id\": \"%7B%5D_P\",\n  \"name\": \"Price\",\n  \"type\": \"number\",\n  \"number\": {\n    \"format\": \"dollar\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "People ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("PeopleProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA people database property is rendered in the Notion UI as a column that contains people mentions.  The people type object is empty; there is no additional configuration. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PeopleProperty").addFields(
						&field{name: "people", typeCode: jen.Struct()},
					).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Project owner\": {\n  \"id\": \"FlgQ\",\n  \"name\": \"Project owner\",\n  \"type\": \"people\",\n  \"people\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Phone number",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("PhoneNumberProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA phone number database property is rendered in the Notion UI as a column that contains phone number values. \n\nThe phone_number type object is empty. There is no additional property configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("PhoneNumberProperty").addFields(
						&field{name: "phone_number", typeCode: jen.Struct()},
					).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Contact phone number\": {\n  \"id\": \"ULHa\",\n  \"name\": \"Contact phone number\",\n  \"type\": \"phone_number\",\n  \"phone_number\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Relation",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("RelationProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA relation database property is rendered in the Notion UI as column that contains relations, references to pages in another database, as values. \n\nThe relation type object contains the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RelationProperty").addFields(
						&interfaceField{name: "relation", typeName: "Relation"},
					).comment += e.Text
					b.addAbstractObject("Relation", "type", e.Text).addVariant(
						b.addSpecificObject("SinglePropertyRelation", "undocumented").addFields(
							&fixedStringField{name: "type", value: "single_property"},
							&field{name: "single_property", typeCode: jen.Struct()},
						),
					).addVariant(
						b.addSpecificObject("DualPropertyRelation", "undocumented").addFields(
							&fixedStringField{name: "type", value: "dual_property"},
						),
					)
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "database_id",
				Type:         "string (UUID)",
				Description:  "The database that the relation property refers to. \n\nThe corresponding linked page values must belong to the database in order to be valid.",
				ExampleValue: `"668d797c-76fa-4934-9b05-ad288df2d136"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("Relation").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Field:        "synced_property_id",
				Type:         "string",
				Description:  "The id of the corresponding property that is updated in the related database when this property is changed.",
				ExampleValue: `"fy:{"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DualPropertyRelation").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "synced_property_name",
				Type:         "string",
				Description:  "The name of the corresponding property that is updated in the related database when this property is changed.",
				ExampleValue: `"Ingredients"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("DualPropertyRelation").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Projects\": {\n  \"id\": \"~pex\",\n  \"name\": \"Projects\",\n  \"type\": \"relation\",\n  \"relation\": {\n    \"database_id\": \"6c4240a9-a3ce-413e-9fd0-8a51a4d0a49b\",\n    \"synced_property_name\": \"Tasks\",\n    \"synced_property_id\": \"JU]K\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCalloutElement{
				Body:   "",
				Title:  "To update a database relation property via the API, share the related parent database with the integration.",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text: "Rich text",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("RichTextProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA rich text database property is rendered in the Notion UI as a column that contains text values. The rich_text type object is empty; there is no additional configuration. ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RichTextProperty").addFields(
						&field{name: "rich_text", typeCode: jen.Struct()},
					).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Project description\": {\n  \"id\": \"NZZ%3B\",\n  \"name\": \"Project description\",\n  \"type\": \"rich_text\",\n  \"rich_text\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Rollup ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("RollupProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA rollup database property is rendered in the Notion UI as a column with values that are rollups, specific properties that are pulled from a related database. \n\nThe rollup type object contains the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("RollupProperty").typeObject.comment = e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "function",
				Type:         "string (enum)",
				Description:  "The function that computes the rollup value from the related pages.\n\nPossible values include: \n\n- average\n- checked\n- count_per_group\n- count\n- count_values \n- date_range\n- earliest_date \n- empty\n- latest_date\n- max\n- median\n- min\n- not_empty\n- percent_checked\n- percent_empty\n- percent_not_empty\n- percent_per_group\n- percent_unchecked\n- range\n- unchecked\n- unique\n- show_original\n- show_unique\n- sum",
				ExampleValue: `"sum"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("RollupProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "relation_property_id",
				Type:         "string",
				Description:  "The id of the related database property that is rolled up.",
				ExampleValue: `"fy:{"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("RollupProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "relation_property_name",
				Type:         "string",
				Description:  "The name of the related database property that is rolled up.",
				ExampleValue: `Tasks"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("RollupProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "rollup_property_id",
				Type:         "string",
				Description:  "The id of the rollup property.",
				ExampleValue: `"fy:{"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("RollupProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "rollup_property_name",
				Type:         "string",
				Description:  "The name of the rollup property.",
				ExampleValue: `"Days to complete"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("RollupProperty").typeObject.addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Estimated total project time\": {\n  \"id\": \"%5E%7Cy%3C\",\n  \"name\": \"Estimated total project time\",\n  \"type\": \"rollup\",\n  \"rollup\": {\n    \"rollup_property_name\": \"Days to complete\",\n    \"relation_property_name\": \"Tasks\",\n    \"rollup_property_id\": \"\\\\nyY\",\n    \"relation_property_id\": \"Y]<y\",\n    \"function\": \"sum\"\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocHeadingElement{
				Text: "Select",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("SelectProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA select database property is rendered in the Notion UI as a column that contains values from a selection of options. Only one option is allowed per row. \n\nThe select type object contains an array of objects representing the available options. Each option object includes the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("SelectProperty").comment = e.Text
					b.getSpecificObject("SelectProperty").typeObject.addFields(
						// TODO 共通化
						&field{name: "options", typeCode: jen.Index().Id("SelectPropertyOption")},
					)
					b.addSpecificObject("SelectPropertyOption", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "color",
				Type:         "string (enum)",
				Description:  "The color of the option as rendered in the Notion UI. Possible values include: \n\n- blue\n- brown\n- default\n- gray\n- green\n- orange\n- pink\n- purple\n- red\n- yellow",
				ExampleValue: `- "red"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyOption").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "id",
				Type:         "string",
				Description:  "An identifier for the option. It doesn't change if the name is changed. These are sometimes, but not always, UUIDs.",
				ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyOption").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Field:        "name",
				Type:         "string",
				Description:  "The name of the option as it appears in the Notion UI.\n\nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: `"Fruit"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("SelectPropertyOption").addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Food group\": {\n  \"id\": \"%40Q%5BM\",\n  \"name\": \"Food group\",\n  \"type\": \"select\",\n  \"select\": {\n    \"options\": [\n      {\n        \"id\": \"e28f74fc-83a7-4469-8435-27eb18f9f9de\",\n        \"name\": \"🥦Vegetable\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"6132d771-b283-4cd9-ba44-b1ed30477c7f\",\n        \"name\": \"🍎Fruit\",\n        \"color\": \"red\"\n      },\n      {\n        \"id\": \"fc9ea861-820b-4f2b-bc32-44ed9eca873c\",\n        \"name\": \"💪Protein\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output: func(e *objectDocCodeElementCode, b *builder) error {
					return nil // TODO
				},
			}}},
			&objectDocHeadingElement{
				Text: "Status",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("StatusProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA status database property is rendered in the Notion UI as a column that contains values from a list of status options. The status type object includes an array of options objects and an array of groups objects. \n\nThe options array is a sorted list of list of the available status options for the property. Each option object in the array has the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("StatusProperty").typeObject.addFields(
						&field{name: "options", typeCode: jen.Index().Id("StatusPropertyDataOption")},
						&field{name: "groups", typeCode: jen.Index().Id("StatusPropertyDataGroup")},
					)
					b.addSpecificObject("StatusPropertyDataOption", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "color",
				Type:         "string (enum)",
				Description:  "The color of the option as rendered in the Notion UI. Possible values include: \n\n- blue\n- brown\n- default\n- gray\n- green\n- orange\n- pink\n- purple\n- red\n- yellow",
				ExampleValue: `"green"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataOption").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "id",
				Type:         "string",
				Description:  "An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.",
				ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataOption").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Field:        "name",
				Type:         "string",
				Description:  "The name of the option as it appears in the Notion UI.\n\nNote: Commas (\",\") are not valid for status values.",
				ExampleValue: `"In progress"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataOption").addFields(e.asField(jen.String()))
					return nil
				},
			}},
			&objectDocParagraphElement{
				Text: "A group is a collection of options. The groups array is a sorted list of the available groups for the property. Each group object in the array has the following fields: ",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.addSpecificObject("StatusPropertyDataGroup", e.Text)
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:        "color",
				Type:         "string (enum)",
				Description:  "The color of the option as rendered in the Notion UI. Possible values include: \n\n- blue\n- brown\n- default\n- gray\n- green\n- orange\n- pink\n- purple\n- red\n- yellow",
				ExampleValue: `"purple"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataGroup").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Field:        "id",
				Type:         "string",
				Description:  "An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.",
				ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataGroup").addFields(e.asField(UUID))
					return nil
				},
			}, {
				Field:        "name",
				Type:         "string",
				Description:  "The name of the option as it appears in the Notion UI.\n\nNote: Commas (\",\") are not valid for status values.",
				ExampleValue: `"To do"`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataGroup").addFields(e.asField(jen.String()))
					return nil
				},
			}, {
				Description:  "A sorted list of ids of all of the options that belong to a group.",
				ExampleValue: "Refer to the example status object below.",
				Field:        "option_ids",
				Type:         "an array of strings (UUID)",
				output: func(e *objectDocParameter, b *builder) error {
					b.getSpecificObject("StatusPropertyDataGroup").addFields(e.asField(jen.Index().Add(UUID)))
					return nil
				},
			}},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Status\": {\n  \"id\": \"biOx\",\n  \"name\": \"Status\",\n  \"type\": \"status\",\n  \"status\": {\n    \"options\": [\n      {\n        \"id\": \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\",\n        \"name\": \"Not started\",\n        \"color\": \"default\"\n      },\n      {\n        \"id\": \"330aeafb-598c-4e1c-bc13-1148aa5963d3\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\",\n        \"name\": \"Done\",\n        \"color\": \"green\"\n      }\n    ],\n    \"groups\": [\n      {\n        \"id\": \"b9d42483-e576-4858-a26f-ed940a5f678f\",\n        \"name\": \"To-do\",\n        \"color\": \"gray\",\n        \"option_ids\": [\n          \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\"\n        ]\n      },\n      {\n        \"id\": \"cf4952eb-1265-46ec-86ab-4bded4fa2e3b\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\",\n        \"option_ids\": [\n          \"330aeafb-598c-4e1c-bc13-1148aa5963d3\"\n        ]\n      },\n      {\n        \"id\": \"4fa7348e-ae74-46d9-9585-e773caca6f40\",\n        \"name\": \"Complete\",\n        \"color\": \"green\",\n        \"option_ids\": [\n          \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\"\n        ]\n      }\n    ]\n  }\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCalloutElement{
				Body:   "Update these values from the Notion UI, instead.",
				Title:  "It is not possible to update a status database property's `name` or `options` values via the API.",
				Type:   "warning",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text: "Title",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("TitleProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA title database property controls the title that appears at the top of a page when a database row is opened. The title type object itself is empty; there is no additional configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("TitleProperty").addFields(&field{
						name:     "title",
						typeCode: jen.Struct(),
					}).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Project name\": {\n  \"id\": \"title\",\n  \"name\": \"Project name\",\n  \"type\": \"title\",\n  \"title\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
			&objectDocCalloutElement{
				Body:   "The API throws errors if you send a request to [Create a database](https://developers.notion.com/reference/create-a-database) without a `title` property, or if you attempt to [Update a database](https://developers.notion.com/reference/update-a-database) to add or remove a `title` property.",
				Title:  "All databases require one, and only one, `title` property.",
				Type:   "warning",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocCalloutElement{
				Body:   "A `title` database property is a type of column in a database. \n\nA database `title` defines the title of the database and is found on the [database object](https://developers.notion.com/reference/database#all-databases). \n\nEvery database requires both a database `title` and a `title` database property.",
				Title:  "Title database property vs. database title",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text: "URL",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.getSpecificObject("UrlProperty").comment = e.Text
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nA URL database property is represented in the Notion UI as a column that contains URL values. \n\nThe url type object is empty. There is no additional property configuration.",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getSpecificObject("UrlProperty").addFields(
						&field{name: "url", typeCode: jen.Struct()},
					).comment += e.Text
					return nil
				},
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "\"Project URL\": {\n  \"id\": \"BZKU\",\n  \"name\": \"Project URL\",\n  \"type\": \"url\",\n  \"url\": {}\n}",
				Language: "json",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
		},
	})
}
