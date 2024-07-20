package objects_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestProperty(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/property-object")

	var property *UnionStruct

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "All database objects include a child properties object. This properties object is composed of individual database property objects. These property objects define the database schema and are rendered in the Notion UI as database columns.",
	}).Output(func(e *Block, b *CodeBuilder) {
		property = b.AddUnionStruct("Property", "type", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘 Database rowsIf you’re looking for information about how to use the API to work with database rows, then refer to the page property values documentation. The API treats database rows as pages."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Every database property object contains the following keys:"})

	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string",
		Description:  "An identifier for the property, usually a short string of random letters and symbols.  \n  \nSome automatically generated property types have special human-readable IDs. For example, all Title properties have an id of \"title\".",
		ExampleValue: `"fy:{"`,
	}).Output(func(e *Parameter, b *CodeBuilder) {
		property.AddFields(b.NewField(e, jen.String()))
	})
	c.ExpectParameter(&Parameter{
		Property:    "name",
		Type:        "string",
		Description: "The name of the property as it appears in Notion.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		property.AddFields(b.NewField(e, jen.String()))
	})
	c.ExpectParameter(&Parameter{
		Property:    "description",
		Type:        "string",
		Description: "The description of a property as it appear in Notion. ",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		property.AddFields(b.NewField(e, jen.Op("*").String(), OmitEmpty))
	})
	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         "string (enum)",
		Description:  "The type that controls the behavior of the property. Possible values are:  \n  \n\\- \"checkbox\"  \n- \"created_by\"  \n- \"created_time\"  \n- \"date\"  \n- \"email\"  \n- \"files\"  \n- \"formula\"  \n- \"last_edited_by\"  \n- \"last_edited_time\"  \n- \"multi_select\"  \n- \"number\"  \n- \"people\"  \n- \"phone_number\"  \n- \"relation\"  \n- \"rich_text\"  \n- \"rollup\"  \n- \"select\"  \n- \"status\"  \n- \"title\"  \n- \"url\"",
		ExampleValue: `"rich_text"`,
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Each database property object also contains a type object. The key of the object is the type of the object, and the value is an object containing type-specific configuration. The following sections detail these type-specific objects along with example property objects for each type.",
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Checkbox",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("checkbox", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A checkbox database property is rendered in the Notion UI as a column that contains checkboxes. The checkbox type object is empty; there is no additional property configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Task complete\": {\n  \"id\": \"BBla\",\n  \"name\": \"Task complete\",\n  \"type\": \"checkbox\",\n  \"checkbox\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Created by",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("created_by", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A created by database property is rendered in the Notion UI as a column that contains people mentions of each row's author as values.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The created_by type object is empty. There is no additional property configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Created by\": {\n  \"id\": \"%5BJCR\",\n  \"name\": \"Created by\",\n  \"type\": \"created_by\",\n  \"created_by\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Created time",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("created_time", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A created time database property is rendered in the Notion UI as a column that contains timestamps of when each row was created as values.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The created_time type object is empty. There is no additional property configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Created time\": {\n  \"id\": \"XcAf\",\n  \"name\": \"Created time\",\n  \"type\": \"created_time\",\n  \"created_time\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Date",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("date", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A date database property is rendered in the Notion UI as a column that contains date values.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The date type object is empty; there is no additional configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Task due date\" {\n  \"id\": \"AJP%7D\",\n  \"name\": \"Task due date\",\n  \"type\": \"date\",\n  \"date\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		e.Text = strings.Replace(e.Text, `"Task due date"`, `"Task due date":`, 1) // ドキュメントの不具合
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Email",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("email", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "An email database property is represented in the Notion UI as a column that contains email values.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The email type object is empty. There is no additional property configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Contact email\": {\n  \"id\": \"oZbC\",\n  \"name\": \"Contact email\",\n  \"type\": \"email\",\n  \"email\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Files",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("files", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "📘The Notion API does not yet support uploading files to Notion.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A files database property is rendered in the Notion UI as a column that has values that are either files uploaded directly to Notion or external links to files. The files type object is empty; there is no additional configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Product image\": {\n  \"id\": \"pb%3E%5B\",\n  \"name\": \"Product image\",\n  \"type\": \"files\",\n  \"files\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	{
		var propertyFormula *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Formula",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyFormula = property.AddPayloadField("formula", e.Text, WithPayloadObject(b))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A formula database property is rendered in the Notion UI as a column that contains values derived from a provided expression.",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyFormula.AddComment(e.Text)
		})

		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The formula type object defines the expression in the following fields:"})

		c.ExpectParameter(&Parameter{
			Property:     "expression",
			Type:         "string",
			Description:  "The formula that is used to compute the values for this property.  \n  \nRefer to the Notion help center for information about formula syntax.",
			ExampleValue: "{{notion:block_property:BtVS:00000000-0000-0000-0000-000000000000:8994905a-074a-415f-9bcf-d1f8b4fa38e4}}/2",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyFormula.AddFields(b.NewField(e, jen.String()))
		})
	}

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Updated price\": {\n  \"id\": \"YU%7C%40\",\n  \"name\": \"Updated price\",\n  \"type\": \"formula\",\n  \"formula\": {\n    \"expression\": \"{{notion:block_property:BtVS:00000000-0000-0000-0000-000000000000:8994905a-074a-415f-9bcf-d1f8b4fa38e4}}/2\"\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Last edited by",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("last_edited_by", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A last edited by database property is rendered in the Notion UI as a column that contains people mentions of the person who last edited each row as values."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The last_edited_by type object is empty. There is no additional property configuration."})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Last edited time",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("last_edited_time", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A last edited time database property is rendered in the Notion UI as a column that contains timestamps of when each row was last edited as values."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The last_edited_time type object is empty. There is no additional property configuration."})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Last edited time\": {\n  \"id\": \"jGdo\",\n  \"name\": \"Last edited time\",\n  \"type\": \"last_edited_time\",\n  \"last_edited_time\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	{
		var propertyMultiSelect *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Multi-select",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyMultiSelect = property.AddPayloadField("multi_select", e.Text, WithPayloadObject(b))
		})

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A multi-select database property is rendered in the Notion UI as a column that contains values from a range of options. Each row can contain one or multiple options.",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyMultiSelect.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The multi_select type object includes an array of options objects. Each option object details settings for the option, indicating the following fields:",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyMultiSelect.AddFields(b.NewField(&Parameter{Property: "options"}, jen.Index().Id("OptionDescription")))
		})
	}

	c.ExpectParameter(&Parameter{
		Property:     "color",
		Type:         "string (enum)",
		Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n\\- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
		ExampleValue: `"blue"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string",
		Description:  "An identifier for the option, which does not change if the name is changed. An id is sometimes, but not always, a UUID.",
		ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "name",
		Type:         "string",
		Description:  "The name of the option as it appears in Notion.  \n  \nNote: Commas (\",\") are not valid for multi-select properties.",
		ExampleValue: `"Fruit"`,
	}) // Optionで共通化
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Store availability\": {\n  \"id\": \"flsb\",\n  \"name\": \"Store availability\",\n  \"type\": \"multi_select\",\n  \"multi_select\": {\n    \"options\": [\n      {\n        \"id\": \"5de29601-9c24-4b04-8629-0bca891c5120\",\n        \"name\": \"Duc Loi Market\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"385890b8-fe15-421b-b214-b02959b0f8d9\",\n        \"name\": \"Rainbow Grocery\",\n        \"color\": \"gray\"\n      },\n      {\n        \"id\": \"72ac0a6c-9e00-4e8c-80c5-720e4373e0b9\",\n        \"name\": \"Nijiya Market\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"9556a8f7-f4b0-4e11-b277-f0af1f8c9490\",\n        \"name\": \"Gus's Community Market\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	{
		var propertyNumber *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Number",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyNumber = property.AddPayloadField("number", e.Text, WithPayloadObject(b))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A number database property is rendered in the Notion UI as a column that contains numeric values. The number type object contains the following fields:",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyNumber.AddComment(e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "format",
			Type:         "string (enum)",
			Description:  "The way that the number is displayed in Notion. Potential values include:  \n  \n\\- argentine_peso  \n- baht  \n- australian_dollar  \n- canadian_dollar  \n- chilean_peso  \n- colombian_peso  \n- danish_krone  \n- dirham  \n- dollar  \n- euro  \n- forint  \n- franc  \n- hong_kong_dollar  \n- koruna  \n- krona  \n- leu  \n- lira  \n-  mexican_peso  \n- new_taiwan_dollar  \n- new_zealand_dollar  \n- norwegian_krone  \n- number  \n- number_with_commas  \n- percent  \n- philippine_peso  \n- pound  \n- peruvian_sol  \n- rand  \n- real  \n- ringgit  \n- riyal  \n- ruble  \n- rupee  \n- rupiah  \n- shekel  \n- singapore_dollar  \n- uruguayan_peso  \n- yen,  \n- yuan  \n- won  \n- zloty",
			ExampleValue: `"percent"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyNumber.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "\"Price\"{\n  \"id\": \"%7B%5D_P\",\n  \"name\": \"Price\",\n  \"type\": \"number\",\n  \"number\": {\n    \"format\": \"dollar\"\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			e.Text = strings.Replace(e.Text, `"Price"`, `"Price":`, 1) // ドキュメントの不具合
			converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
		})
	}

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "People",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("people", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A people database property is rendered in the Notion UI as a column that contains people mentions.  The people type object is empty; there is no additional configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Project owner\": {\n  \"id\": \"FlgQ\",\n  \"name\": \"Project owner\",\n  \"type\": \"people\",\n  \"people\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Phone number",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("phone_number", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A phone number database property is rendered in the Notion UI as a column that contains phone number values.",
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "The phone_number type object is empty. There is no additional property configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Contact phone number\": {\n  \"id\": \"ULHa\",\n  \"name\": \"Contact phone number\",\n  \"type\": \"phone_number\",\n  \"phone_number\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	{
		var propertyRelation *UnionStruct
		var propertyRelationDualProperty *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Relation",
		}).Output(func(e *Block, b *CodeBuilder) {
			property.AddPayloadField("relation", e.Text, WithType(jen.Op("*").Id("PropertyRelation")))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A relation database property is rendered in the Notion UI as column that contains relations, references to pages in another database, as values.",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyRelation = b.AddUnionStruct("PropertyRelation", "type", e.Text)
			propertyRelation.AddPayloadField("single_property", "undocumented", WithEmptyStruct())
			propertyRelationDualProperty = propertyRelation.AddPayloadField("dual_property", "undocumented", WithPayloadObject(b))
		})

		c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The relation type object contains the following fields:"})

		c.ExpectParameter(&Parameter{
			Property:     "database_id",
			Type:         "string (UUID)",
			Description:  "The database that the relation property refers to.  \n  \nThe corresponding linked page values must belong to the database in order to be valid.",
			ExampleValue: `"668d797c-76fa-4934-9b05-ad288df2d136"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRelation.AddFields(b.NewField(e, UUID))
		})
		c.ExpectParameter(&Parameter{
			Property:     "synced_property_id",
			Type:         "string",
			Description:  "The id of the corresponding property that is updated in the related database when this property is changed.",
			ExampleValue: `"fy:{"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRelationDualProperty.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "synced_property_name",
			Type:         "string",
			Description:  "The name of the corresponding property that is updated in the related database when this property is changed.",
			ExampleValue: `"Ingredients"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRelationDualProperty.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "\"Projects\": {\n  \"id\": \"~pex\",\n  \"name\": \"Projects\",\n  \"type\": \"relation\",\n  \"relation\": {\n    \"database_id\": \"6c4240a9-a3ce-413e-9fd0-8a51a4d0a49b\",\n    \"synced_property_name\": \"Tasks\",\n    \"synced_property_id\": \"JU]K\"\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			// これを含めるとProjectのアンマーシャルができないのでスキップ
			// TODO 対応する
			// converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Code))
		})
		c.ExpectBlock(&Block{
			Kind: "Blockquote",
			Text: "📘 Database relations must be shared with your integrationTo retrieve properties from database relations, the related database must be shared with your integration in addition to the database being retrieved. If the related database is not shared, properties based on relations will not be included in the API response.Similarly, to update a database relation property via the API, share the related database with the integration.",
		})
	}

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Rich text",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("rich_text", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "A rich text database property is rendered in the Notion UI as a column that contains text values. The rich_text type object is empty; there is no additional configuration.",
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Project description\": {\n  \"id\": \"NZZ%3B\",\n  \"name\": \"Project description\",\n  \"type\": \"rich_text\",\n  \"rich_text\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})

	{
		var propertyRollup *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Rollup",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyRollup = property.AddPayloadField("rollup", e.Text, WithPayloadObject(b))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A rollup database property is rendered in the Notion UI as a column with values that are rollups, specific properties that are pulled from a related database.",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyRollup.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The rollup type object contains the following fields:",
		})
		c.ExpectParameter(&Parameter{
			Property:     "function",
			Type:         "string (enum)",
			Description:  "The function that computes the rollup value from the related pages.  \n  \nPossible values include:  \n  \n\\- average  \n- checked  \n- count_per_group  \n- count  \n- count_values  \n- date_range  \n- earliest_date  \n- empty  \n- latest_date  \n- max  \n- median  \n- min  \n- not_empty  \n- percent_checked  \n- percent_empty  \n- percent_not_empty  \n- percent_per_group  \n- percent_unchecked  \n- range  \n- unchecked  \n- unique  \n- show_original  \n- show_unique  \n- sum",
			ExampleValue: `"sum"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRollup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "relation_property_id",
			Type:         "string",
			Description:  "The id of the related database property that is rolled up.",
			ExampleValue: `"fy:{"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRollup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "relation_property_name",
			Type:         "string",
			Description:  "The name of the related database property that is rolled up.",
			ExampleValue: `Tasks"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRollup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "rollup_property_id",
			Type:         "string",
			Description:  "The id of the rollup property.",
			ExampleValue: `"fy:{"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRollup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "rollup_property_name",
			Type:         "string",
			Description:  "The name of the rollup property.",
			ExampleValue: `"Days to complete"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertyRollup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "\"Estimated total project time\": {\n  \"id\": \"%5E%7Cy%3C\",\n  \"name\": \"Estimated total project time\",\n  \"type\": \"rollup\",\n  \"rollup\": {\n    \"rollup_property_name\": \"Days to complete\",\n    \"relation_property_name\": \"Tasks\",\n    \"rollup_property_id\": \"\\\\nyY\",\n    \"relation_property_id\": \"Y]<y\",\n    \"function\": \"sum\"\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
		})
	}

	{
		var propertySelect *SimpleObject
		var option *SimpleObject
		var optionDescription *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Select",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertySelect = property.AddPayloadField("select", e.Text, WithPayloadObject(b))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A select database property is rendered in the Notion UI as a column that contains values from a selection of options. Only one option is allowed per row.",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertySelect.AddComment(e.Text)
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The select type object contains an array of objects representing the available options. Each option object includes the following fields:",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertySelect.AddFields(b.NewField(&Parameter{Property: "options"}, jen.Index().Id("OptionDescription")))
			// (Select, MultiSelect, Status) * (Property, PropertyItem, PropertyValue) の9箇所で
			// 以下の共通の構造体を使います
			option = b.AddSimpleObject("Option", e.Text)
			optionDescription = b.AddSimpleObject("OptionDescription", e.Text) // TODO Optionにdescriptionが入る場合と入らない場合があるので切り分ける
		})
		c.ExpectParameter(&Parameter{
			Property:     "color",
			Type:         "string (enum)",
			Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n\\- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
			ExampleValue: `\- "red"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			option.AddFields(b.NewField(e, jen.String(), OmitEmpty))
			optionDescription.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "id",
			Type:         "string",
			Description:  "An identifier for the option. It doesn't change if the name is changed. These are sometimes, but not always, UUIDs.",
			ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			option.AddFields(b.NewField(e, jen.String(), OmitEmpty))
			optionDescription.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectParameter(&Parameter{
			Property:     "name",
			Type:         "string",
			Description:  "The name of the option as it appears in the Notion UI.  \n  \nNote: Commas (\",\") are not valid for select values.",
			ExampleValue: `"Fruit"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			option.AddFields(b.NewField(e, jen.String(), OmitEmpty))
			optionDescription.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "\"Food group\": {\n  \"id\": \"%40Q%5BM\",\n  \"name\": \"Food group\",\n  \"type\": \"select\",\n  \"select\": {\n    \"options\": [\n      {\n        \"id\": \"e28f74fc-83a7-4469-8435-27eb18f9f9de\",\n        \"name\": \"🥦Vegetable\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"6132d771-b283-4cd9-ba44-b1ed30477c7f\",\n        \"name\": \"🍎Fruit\",\n        \"color\": \"red\"\n      },\n      {\n        \"id\": \"fc9ea861-820b-4f2b-bc32-44ed9eca873c\",\n        \"name\": \"💪Protein\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
		})

		c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
			optionDescription.AddFields(b.NewField(&Parameter{Property: "description", Description: UNDOCUMENTED}, jen.Op("*").String()))
		})
	}

	{
		var propertyStatus *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Status",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyStatus = property.AddPayloadField("status", e.Text, WithPayloadObject(b))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A status database property is rendered in the Notion UI as a column that contains values from a list of status options. The status type object includes an array of options objects and an array of groups objects.",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertyStatus.AddFields(
				b.NewField(&Parameter{Property: "options"}, jen.Index().Id("OptionDescription")),
				b.NewField(&Parameter{Property: "groups"}, jen.Index().Id("StatusGroup")),
			)
		})
	}

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The options array is a sorted list of list of the available status options for the property. Each option object in the array has the following fields:"})
	c.ExpectParameter(&Parameter{
		Property:     "color",
		Type:         "string (enum)",
		Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n\\- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
		ExampleValue: `"green"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "id",
		Type:         "string",
		Description:  "An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.",
		ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
	}) // Optionで共通化
	c.ExpectParameter(&Parameter{
		Property:     "name",
		Type:         "string",
		Description:  "The name of the option as it appears in the Notion UI.  \n  \nNote: Commas (\",\") are not valid for status values.",
		ExampleValue: `"In progress"`,
	}) // Optionで共通化

	{
		var statusGroup *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "A group is a collection of options. The groups array is a sorted list of the available groups for the property. Each group object in the array has the following fields:",
		}).Output(func(e *Block, b *CodeBuilder) {
			statusGroup = b.AddSimpleObject("StatusGroup", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "color",
			Type:         "string (enum)",
			Description:  "The color of the option as rendered in the Notion UI. Possible values include:  \n  \n\\- blue  \n- brown  \n- default  \n- gray  \n- green  \n- orange  \n- pink  \n- purple  \n- red  \n- yellow",
			ExampleValue: `"purple"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			statusGroup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "id",
			Type:         "string",
			Description:  "An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.",
			ExampleValue: `"ff8e9269-9579-47f7-8f6e-83a84716863c"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			statusGroup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "name",
			Type:         "string",
			Description:  "The name of the option as it appears in the Notion UI.  \n  \nNote: Commas (\",\") are not valid for status values.",
			ExampleValue: `"To do"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			statusGroup.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "option_ids",
			Type:         "an array of strings (UUID)",
			Description:  "A sorted list of ids of all of the options that belong to a group.",
			ExampleValue: "Refer to the example status object below.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			statusGroup.AddFields(b.NewField(e, jen.Index().String()))
		})
		c.ExpectBlock(&Block{
			Kind: "FencedCodeBlock",
			Text: "\"Status\": {\n  \"id\": \"biOx\",\n  \"name\": \"Status\",\n  \"type\": \"status\",\n  \"status\": {\n    \"options\": [\n      {\n        \"id\": \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\",\n        \"name\": \"Not started\",\n        \"color\": \"default\"\n      },\n      {\n        \"id\": \"330aeafb-598c-4e1c-bc13-1148aa5963d3\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\",\n        \"name\": \"Done\",\n        \"color\": \"green\"\n      }\n    ],\n    \"groups\": [\n      {\n        \"id\": \"b9d42483-e576-4858-a26f-ed940a5f678f\",\n        \"name\": \"To-do\",\n        \"color\": \"gray\",\n        \"option_ids\": [\n          \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\"\n        ]\n      },\n      {\n        \"id\": \"cf4952eb-1265-46ec-86ab-4bded4fa2e3b\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\",\n        \"option_ids\": [\n          \"330aeafb-598c-4e1c-bc13-1148aa5963d3\"\n        ]\n      },\n      {\n        \"id\": \"4fa7348e-ae74-46d9-9585-e773caca6f40\",\n        \"name\": \"Complete\",\n        \"color\": \"green\",\n        \"option_ids\": [\n          \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\"\n        ]\n      }\n    ]\n  }\n}\n",
		}).Output(func(e *Block, b *CodeBuilder) {
			converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
		})
		c.ExpectBlock(&Block{
			Kind: "Blockquote",
			Text: "🚧 It is not possible to update a status database property's name or options values via the API.Update these values from the Notion UI, instead.",
		})
	}

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Title",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("title", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A title database property controls the title that appears at the top of a page when a database row is opened. The title type object itself is empty; there is no additional configuration."})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Project name\": {\n  \"id\": \"title\",\n  \"name\": \"Project name\",\n  \"type\": \"title\",\n  \"title\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "🚧 All databases require one, and only one, title property.The API throws errors if you send a request to Create a database without a title property, or if you attempt to Update a database to add or remove a title property."})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘 Title database property vs. database titleA title database property is a type of column in a database.A database title defines the title of the database and is found on the database object.Every database requires both a database title and a title database property."})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "URL",
	}).Output(func(e *Block, b *CodeBuilder) {
		property.AddPayloadField("url", e.Text, WithEmptyStruct())
	})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "A URL database property is represented in the Notion UI as a column that contains URL values."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The url type object is empty. There is no additional property configuration."})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "\"Project URL\": {\n  \"id\": \"BZKU\",\n  \"name\": \"Project URL\",\n  \"type\": \"url\",\n  \"url\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("PropertyMap", fmt.Sprintf(`{%s}`, e.Text))
	})
}
