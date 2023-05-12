package objectdoc

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	var property *adaptiveObject

	registerTranslator(
		"https://developers.notion.com/reference/property-object",
		func(c *comparator, b *builder) {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "All database objects include a child properties object. This properties object is composed of individual database property objects. These property objects define the database schema and are rendered in the Notion UI as database columns.",
			}, func(e blockElement) {
				property = b.addAdaptiveObject("Property", "type", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘 Database rowsIf you’re looking for information about how to use the API to work with database rows, then refer to the page property values documentation. The API treats database rows as pages.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Every database property object contains the following keys:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Property:     "id",
				Type:         "string",
				Description:  "An identifier for the property, usually a short string of random letters and symbols.  \n  \nSome automatically generated property types have special human-readable IDs. For example, all Title properties have an id of \"title\".",
				ExampleValue: `"fy:{"`,
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("Property").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Property:    "name",
				Type:        "string",
				Description: "The name of the property as it appears in Notion.",
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("Property").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Property:     "type",
				Type:         "string (enum)",
				Description:  "The type that controls the behavior of the property. Possible values are:  \n  \n- \"checkbox\"  \n- \"created_by\"  \n- \"created_time\"  \n- \"date\"  \n- \"email\"  \n- \"files\"  \n- \"formula\"  \n- \"last_edited_by\"  \n- \"last_edited_time\"  \n- \"multi_select\"  \n- \"number\"  \n- \"people\"  \n- \"phone_number\"  \n- \"relation\"  \n- \"rich_text\"  \n- \"rollup\"  \n- \"select\"  \n- \"status\"  \n- \"title\"  \n- \"url\"",
				ExampleValue: `"rich_text"`,
			}, func(e parameterElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Each database property object also contains a type object. The key of the object is the type of the object, and the value is an object containing type-specific configuration. The following sections detail these type-specific objects along with example property objects for each type.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) { // Checkbox
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Checkbox",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("checkbox", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A checkbox database property is rendered in the Notion UI as a column that contains checkboxes. The checkbox type object is empty; there is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Task complete\": {\n  \"id\": \"BBla\",\n  \"name\": \"Task complete\",\n  \"type\": \"checkbox\",\n  \"checkbox\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Created by
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Created by",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("created_by", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A created by database property is rendered in the Notion UI as a column that contains people mentions of each row's author as values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The created_by type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Created by\": {\n  \"id\": \"%5BJCR\",\n  \"name\": \"Created by\",\n  \"type\": \"created_by\",\n  \"created_by\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Created time
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Created time",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("created_time", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A created time database property is rendered in the Notion UI as a column that contains timestamps of when each row was created as values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The created_time type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Created time\": {\n  \"id\": \"XcAf\",\n  \"name\": \"Created time\",\n  \"type\": \"created_time\",\n  \"created_time\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Date
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Date",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("date", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A date database property is rendered in the Notion UI as a column that contains date values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The date type object is empty; there is no additional configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Task due date\" {\n  \"id\": \"AJP%7D\",\n  \"name\": \"Task due date\",\n  \"type\": \"date\",\n  \"date\": {}\n}\n",
			}, func(e blockElement) {
				e.Text = strings.Replace(e.Text, `"Task due date"`, `"Task due date":`, 1) // ドキュメントの不具合
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Email
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Email",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("email", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "An email database property is represented in the Notion UI as a column that contains email values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The email type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Contact email\": {\n  \"id\": \"oZbC\",\n  \"name\": \"Contact email\",\n  \"type\": \"email\",\n  \"email\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Files
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Files",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("files", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘The Notion API does not yet support uploading files to Notion.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A files database property is rendered in the Notion UI as a column that has values that are either files uploaded directly to Notion or external links to files. The files type object is empty; there is no additional configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Product image\": {\n  \"id\": \"pb%3E%5B\",\n  \"name\": \"Product image\",\n  \"type\": \"files\",\n  \"files\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Formula
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Formula",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithSpecificObject("formula", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A formula database property is rendered in the Notion UI as a column that contains values derived from a provided expression.",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertyFormula").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The formula type object defines the expression in the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The formula that is used to compute the values for this property.  \n  \nRefer to the Notion help center for information about formula syntax.",
				ExampleValue: "\"prop(\\\"Price\\\") * 2\"",
				Property:     "expression",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyFormula").addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Updated price\": {\n  \"id\": \"YU%7C%40\",\n  \"name\": \"Updated price\",\n  \"type\": \"formula\",\n  \"formula\": {\n    \"expression\": \"prop(\\\"Price\\\") * 2\"\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Last edited by
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Last edited by",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("last_edited_by", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A last edited by database property is rendered in the Notion UI as a column that contains people mentions of the person who last edited each row as values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The last_edited_by type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) { // Last edited time
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Last edited time",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("last_edited_time", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A last edited time database property is rendered in the Notion UI as a column that contains timestamps of when each row was last edited as values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The last_edited_time type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Last edited time\": {\n  \"id\": \"jGdo\",\n  \"name\": \"Last edited time\",\n  \"type\": \"last_edited_time\",\n  \"last_edited_time\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Multi-select
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Multi-select",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithSpecificObject("multi_select", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A multi-select database property is rendered in the Notion UI as a column that contains values from a range of options. Each row can contain one or multiple options.",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertyMultiSelect").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The multi_select type object includes an array of options objects. Each option object details settings for the option, indicating the following fields:",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertyMultiSelect").addFields(
					&field{name: "options", typeCode: jen.Index().Id("Option")},
				)
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
				ExampleValue: "\"blue\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "An identifier for the option, which does not change if the name is changed. An id is sometimes, but not always, a UUID.",
				ExampleValue: "\"ff8e9269-9579-47f7-8f6e-83a84716863c\"",
				Property:     "id",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "The name of the option as it appears in Notion.  \n  \nNote: Commas (\",\") are not valid for multi-select properties.",
				ExampleValue: "\"Fruit\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Store availability\": {\n  \"id\": \"flsb\",\n  \"name\": \"Store availability\",\n  \"type\": \"multi_select\",\n  \"multi_select\": {\n    \"options\": [\n      {\n        \"id\": \"5de29601-9c24-4b04-8629-0bca891c5120\",\n        \"name\": \"Duc Loi Market\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"385890b8-fe15-421b-b214-b02959b0f8d9\",\n        \"name\": \"Rainbow Grocery\",\n        \"color\": \"gray\"\n      },\n      {\n        \"id\": \"72ac0a6c-9e00-4e8c-80c5-720e4373e0b9\",\n        \"name\": \"Nijiya Market\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"9556a8f7-f4b0-4e11-b277-f0af1f8c9490\",\n        \"name\": \"Gus's Community Market\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Number
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Number",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithSpecificObject("number", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A number database property is rendered in the Notion UI as a column that contains numeric values. The number type object contains the following fields:",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertyNumber").addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:     "format",
				Type:         "string (enum)",
				Description:  "The way that the number is displayed in Notion. Potential values include:  \n  \n- argentine_peso  \n- baht  \n- canadian_dollar  \n- chilean_peso  \n- colombian_peso  \n- danish_krone  \n- dirham  \n- dollar  \n- euro  \n- forint  \n- franc  \n- hong_kong_dollar  \n- koruna  \n- krona  \n- leu  \n- lira  \n-  mexican_peso  \n- new_taiwan_dollar  \n- new_zealand_dollar  \n- norwegian_krone  \n- number  \n- number_with_commas  \n- percent  \n- philippine_peso  \n- pound  \n- peruvian_sol  \n- rand  \n- real  \n- ringgit  \n- riyal  \n- ruble  \n- rupee  \n- rupiah  \n- shekel  \n- singapore_dollar  \n- uruguayan_peso  \n- yen,  \n- yuan  \n- won  \n- zloty",
				ExampleValue: `"percent"`,
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyNumber").addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Price\"{\n  \"id\": \"%7B%5D_P\",\n  \"name\": \"Price\",\n  \"type\": \"number\",\n  \"number\": {\n    \"format\": \"dollar\"\n  }\n}\n",
			}, func(e blockElement) {
				e.Text = strings.Replace(e.Text, `"Price"`, `"Price":`, 1) // ドキュメントの不具合
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // People
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "People",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("people", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A people database property is rendered in the Notion UI as a column that contains people mentions.  The people type object is empty; there is no additional configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Project owner\": {\n  \"id\": \"FlgQ\",\n  \"name\": \"Project owner\",\n  \"type\": \"people\",\n  \"people\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Phone number
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Phone number",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("phone_number", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A phone number database property is rendered in the Notion UI as a column that contains phone number values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The phone_number type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Contact phone number\": {\n  \"id\": \"ULHa\",\n  \"name\": \"Contact phone number\",\n  \"type\": \"phone_number\",\n  \"phone_number\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Relation
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Relation",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithType("relation", e.Text, jen.Op("*").Id("PropertyRelation"))
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A relation database property is rendered in the Notion UI as column that contains relations, references to pages in another database, as values.",
			}, func(e blockElement) {
				pr := b.addAdaptiveObject("PropertyRelation", "type", e.Text)
				// getSymbol[concreteObject]("PropertyRelation").addFields(
				// 	&interfaceField{name: "relation", typeName: "Relation"},
				// ).addComment(e.Text)
				// b.addAbstractObject("Relation", "type", e.Text)
				pr.addAdaptiveFieldWithEmptyStruct("single_property", "undocumented")
				pr.addAdaptiveFieldWithSpecificObject("dual_property", "undocumented", b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The relation type object contains the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The database that the relation property refers to.  \n  \nThe corresponding linked page values must belong to the database in order to be valid.",
				ExampleValue: "\"668d797c-76fa-4934-9b05-ad288df2d136\"",
				Property:     "database_id",
				Type:         "string (UUID)",
			}, func(e parameterElement) {
				getSymbol[adaptiveObject]("PropertyRelation").addFields(e.asField(UUID))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The id of the corresponding property that is updated in the related database when this property is changed.",
				ExampleValue: "\"fy:{\"",
				Property:     "synced_property_id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRelationDualProperty").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The name of the corresponding property that is updated in the related database when this property is changed.",
				ExampleValue: "\"Ingredients\"",
				Property:     "synced_property_name",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRelationDualProperty").addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Projects\": {\n  \"id\": \"~pex\",\n  \"name\": \"Projects\",\n  \"type\": \"relation\",\n  \"relation\": {\n    \"database_id\": \"6c4240a9-a3ce-413e-9fd0-8a51a4d0a49b\",\n    \"synced_property_name\": \"Tasks\",\n    \"synced_property_id\": \"JU]K\"\n  }\n}\n",
			}, func(e blockElement) {
				// これを含めるとProjectのアンマーシャルができないのでスキップ
				// TODO 対応する
				// b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Code))
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘 To update a database relation property via the API, share the related parent database with the integration.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) { // Rich text
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rich text",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("rich_text", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A rich text database property is rendered in the Notion UI as a column that contains text values. The rich_text type object is empty; there is no additional configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Project description\": {\n  \"id\": \"NZZ%3B\",\n  \"name\": \"Project description\",\n  \"type\": \"rich_text\",\n  \"rich_text\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Rollup
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Rollup",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithSpecificObject("rollup", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A rollup database property is rendered in the Notion UI as a column with values that are rollups, specific properties that are pulled from a related database.",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertyRollup").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The rollup type object contains the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The function that computes the rollup value from the related pages.  \n  \nPossible values include:  \n  \n- average  \n- checked  \n- count_per_group  \n- count  \n- count_values  \n- date_range  \n- earliest_date  \n- empty  \n- latest_date  \n- max  \n- median  \n- min  \n- not_empty  \n- percent_checked  \n- percent_empty  \n- percent_not_empty  \n- percent_per_group  \n- percent_unchecked  \n- range  \n- unchecked  \n- unique  \n- show_original  \n- show_unique  \n- sum",
				ExampleValue: "\"sum\"",
				Property:     "function",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRollup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The id of the related database property that is rolled up.",
				ExampleValue: "\"fy:{\"",
				Property:     "relation_property_id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRollup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The name of the related database property that is rolled up.",
				ExampleValue: "Tasks\"",
				Property:     "relation_property_name",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRollup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The id of the rollup property.",
				ExampleValue: "\"fy:{\"",
				Property:     "rollup_property_id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRollup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The name of the rollup property.",
				ExampleValue: "\"Days to complete\"",
				Property:     "rollup_property_name",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("PropertyRollup").addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Estimated total project time\": {\n  \"id\": \"%5E%7Cy%3C\",\n  \"name\": \"Estimated total project time\",\n  \"type\": \"rollup\",\n  \"rollup\": {\n    \"rollup_property_name\": \"Days to complete\",\n    \"relation_property_name\": \"Tasks\",\n    \"rollup_property_id\": \"\\\\nyY\",\n    \"relation_property_id\": \"Y]<y\",\n    \"function\": \"sum\"\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Select
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Select",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithSpecificObject("select", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A select database property is rendered in the Notion UI as a column that contains values from a selection of options. Only one option is allowed per row.",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertySelect").addComment(e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The select type object contains an array of objects representing the available options. Each option object includes the following fields:",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertySelect").addFields(
					&field{name: "options", typeCode: jen.Index().Id("Option")},
				)
				// (Select, MultiSelect, Status) * (Property, PropertyItem, PropertyValue) の9箇所で
				// 以下の共通の構造体を使います
				b.addConcreteObject("Option", e.Text)
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
				ExampleValue: "- \"red\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("Option").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "An identifier for the option. It doesn't change if the name is changed. These are sometimes, but not always, UUIDs.",
				ExampleValue: "\"ff8e9269-9579-47f7-8f6e-83a84716863c\"",
				Property:     "id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("Option").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The name of the option as it appears in the Notion UI.  \n  \nNote: Commas (\",\") are not valid for select values.",
				ExampleValue: "\"Fruit\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("Option").addFields(e.asField(jen.String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Food group\": {\n  \"id\": \"%40Q%5BM\",\n  \"name\": \"Food group\",\n  \"type\": \"select\",\n  \"select\": {\n    \"options\": [\n      {\n        \"id\": \"e28f74fc-83a7-4469-8435-27eb18f9f9de\",\n        \"name\": \"🥦Vegetable\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"6132d771-b283-4cd9-ba44-b1ed30477c7f\",\n        \"name\": \"🍎Fruit\",\n        \"color\": \"red\"\n      },\n      {\n        \"id\": \"fc9ea861-820b-4f2b-bc32-44ed9eca873c\",\n        \"name\": \"💪Protein\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
		func(c *comparator, b *builder) { // Status
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Status",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithSpecificObject("status", e.Text, b)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A status database property is rendered in the Notion UI as a column that contains values from a list of status options. The status type object includes an array of options objects and an array of groups objects.",
			}, func(e blockElement) {
				getSymbol[concreteObject]("PropertyStatus").addFields(
					&field{name: "options", typeCode: jen.Index().Id("Option")},
					&field{name: "groups", typeCode: jen.Index().Id("StatusGroup")},
				)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The options array is a sorted list of list of the available status options for the property. Each option object in the array has the following fields:",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
				ExampleValue: "\"green\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.",
				ExampleValue: "\"ff8e9269-9579-47f7-8f6e-83a84716863c\"",
				Property:     "id",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustParameter(parameterElement{
				Description:  "The name of the option as it appears in the Notion UI.  \n  \nNote: Commas (\",\") are not valid for status values.",
				ExampleValue: "\"In progress\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {}) // Optionで共通化
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A group is a collection of options. The groups array is a sorted list of the available groups for the property. Each group object in the array has the following fields:",
			}, func(e blockElement) {
				b.addConcreteObject("StatusGroup", e.Text)
			})
			c.nextMustParameter(parameterElement{
				Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
				ExampleValue: "\"purple\"",
				Property:     "color",
				Type:         "string (enum)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("StatusGroup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.",
				ExampleValue: "\"ff8e9269-9579-47f7-8f6e-83a84716863c\"",
				Property:     "id",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("StatusGroup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The name of the option as it appears in the Notion UI.  \n  \nNote: Commas (\",\") are not valid for status values.",
				ExampleValue: "\"To do\"",
				Property:     "name",
				Type:         "string",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("StatusGroup").addFields(e.asField(jen.String()))
			})
			c.nextMustParameter(parameterElement{
				Description:  "A sorted list of ids of all of the options that belong to a group.",
				ExampleValue: "Refer to the example status object below.",
				Property:     "option_ids",
				Type:         "an array of strings (UUID)",
			}, func(e parameterElement) {
				getSymbol[concreteObject]("StatusGroup").addFields(e.asField(jen.Index().String()))
			})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Status\": {\n  \"id\": \"biOx\",\n  \"name\": \"Status\",\n  \"type\": \"status\",\n  \"status\": {\n    \"options\": [\n      {\n        \"id\": \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\",\n        \"name\": \"Not started\",\n        \"color\": \"default\"\n      },\n      {\n        \"id\": \"330aeafb-598c-4e1c-bc13-1148aa5963d3\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\",\n        \"name\": \"Done\",\n        \"color\": \"green\"\n      }\n    ],\n    \"groups\": [\n      {\n        \"id\": \"b9d42483-e576-4858-a26f-ed940a5f678f\",\n        \"name\": \"To-do\",\n        \"color\": \"gray\",\n        \"option_ids\": [\n          \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\"\n        ]\n      },\n      {\n        \"id\": \"cf4952eb-1265-46ec-86ab-4bded4fa2e3b\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\",\n        \"option_ids\": [\n          \"330aeafb-598c-4e1c-bc13-1148aa5963d3\"\n        ]\n      },\n      {\n        \"id\": \"4fa7348e-ae74-46d9-9585-e773caca6f40\",\n        \"name\": \"Complete\",\n        \"color\": \"green\",\n        \"option_ids\": [\n          \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\"\n        ]\n      }\n    ]\n  }\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧 It is not possible to update a status database property's name or options values via the API.Update these values from the Notion UI, instead.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) { // Title
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Title",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("title", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A title database property controls the title that appears at the top of a page when a database row is opened. The title type object itself is empty; there is no additional configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Project name\": {\n  \"id\": \"title\",\n  \"name\": \"Project name\",\n  \"type\": \"title\",\n  \"title\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "🚧 All databases require one, and only one, title property.The API throws errors if you send a request to Create a database without a title property, or if you attempt to Update a database to add or remove a title property.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "📘 Title database property vs. database titleA title database property is a type of column in a database.A database title defines the title of the database and is found on the database object.Every database requires both a database title and a title database property.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) { // URL
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "URL",
			}, func(e blockElement) {
				property.addAdaptiveFieldWithEmptyStruct("url", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "A URL database property is represented in the Notion UI as a column that contains URL values.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The url type object is empty. There is no additional property configuration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "\"Project URL\": {\n  \"id\": \"BZKU\",\n  \"name\": \"Project URL\",\n  \"type\": \"url\",\n  \"url\": {}\n}\n",
			}, func(e blockElement) {
				b.addUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
			})
		},
	)
}
