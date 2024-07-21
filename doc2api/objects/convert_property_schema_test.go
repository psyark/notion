package objects_test

import (
	"regexp"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
	"github.com/stoewer/go-strcase"
)

func TestPropertySchema(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/property-schema-object")

	var propertySchema *SimpleObject

	addPayload := func(b *CodeBuilder, name string, comment string) *SimpleObject {
		payloadName := "PropertySchema" + strcase.UpperCamelCase(name)
		propertySchema.AddFields(b.NewField(&Parameter{Property: name, Description: comment}, jen.Op("*").Id(payloadName), OmitEmpty))
		return b.AddSimpleObject(payloadName, comment)
	}
	addPayloadAsUnion := func(b *CodeBuilder, name string, comment string) *UnionStruct {
		payloadName := "PropertySchema" + strcase.UpperCamelCase(name)
		propertySchema.AddFields(b.NewField(&Parameter{Property: name, Description: comment}, jen.Op("*").Id(payloadName), OmitEmpty))
		return b.AddUnionStruct(payloadName, "type", comment)
	}

	addEmptyPayload := func(b *CodeBuilder, name, comment string) {
		propertySchema.AddFields(b.NewField(&Parameter{Property: name, Description: comment}, jen.Op("*").Struct(), OmitEmpty))
	}
	addEmptyPayloadByBlock := func(b *CodeBuilder, e *Block) {
		match := regexp.MustCompile(`have no additional configuration within the (\w+) property.`).FindStringSubmatch(e.Text)
		addEmptyPayload(b, match[1], e.Text)
	}

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Metadata that controls how a database property behaves.",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertySchema = b.AddSimpleObject("PropertySchema", e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Database properties",
	}).Output(func(e *Block, b *CodeBuilder) {
		propertySchema.AddComment(e.Text)
	})

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: `Each database property schema object has at least one key which is the property type. This type contains behavior of this property. Possible values of this key are "title", "rich_text", "number", "select", "multi_select", "date", "people", "files", "checkbox", "url", "email", "phone_number", "formula", "relation", "rollup", "created_time", "created_by", "last_edited_time", "last_edited_by".`,
	}).Output(func(e *Block, b *CodeBuilder) {
		propertySchema.AddComment(e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Title configuration"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: `Each database must have exactly one database property schema object of type "title". This database property controls the title that appears at the top of the page when the page is opened. Title database property objects have no additional configuration within the title property.`,
	}).Output(func(e *Block, b *CodeBuilder) {
		addEmptyPayloadByBlock(b, e)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Text configuration"})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Text database property schema objects have no additional configuration within the rich_text property.",
	}).Output(func(e *Block, b *CodeBuilder) {
		addEmptyPayloadByBlock(b, e)
	})

	{
		var payload *SimpleObject

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Number configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Number database property schema objects optionally contain the following configuration within the number property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload = addPayload(b, "number", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "format",
			Type:         "optional string (enum)",
			Description:  "How the number is displayed in Notion. Potential values include: number, number_with_commas, percent, dollar, canadian_dollar, euro, pound, yen, ruble, rupee, won, yuan, real, lira, rupiah, franc, hong_kong_dollar, new_zealand_dollar, krona, norwegian_krone, mexican_peso, rand, new_taiwan_dollar, danish_krone, zloty, baht, forint, koruna, shekel, chilean_peso, philippine_peso, dirham, colombian_peso, riyal, ringgit, leu, argentine_peso, uruguayan_peso, singapore_dollar.",
			ExampleValue: `"percent"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String()))
		})
	}

	{
		var payload *SimpleObject

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Select configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Select database property schema objects optionally contain the following configuration within the select property:",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload = addPayload(b, "select", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:    "options",
			Type:        "optional array of select option objects.",
			Description: "Sorted list of options available for this property.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.Index().Id("PropertySchemaOption")))
		})
	}

	{
		var propertySchemaOption *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Select options",
		}).Output(func(e *Block, b *CodeBuilder) {
			propertySchemaOption = b.AddSimpleObject("PropertySchemaOption", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "name",
			Type:         "string",
			Description:  "Name of the option as it appears in Notion.",
			ExampleValue: `"Fruit"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertySchemaOption.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "color",
			Type:         "optional string (enum)",
			Description:  "Color of the option. Possible values include: default, gray, brown, orange, yellow, green, blue, purple, pink, red.",
			ExampleValue: `"red"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			propertySchemaOption.AddFields(b.NewField(e, jen.String()))
		})
	}

	{
		var payload *SimpleObject

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Multi-select configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Multi-select database property schema objects optionally contain the following configuration within the multi_select property:",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload = addPayload(b, "multi_select", e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:    "options",
			Type:        "optional array of multi-select option objects.",
			Description: "Settings for multi select properties.",
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.Index().Id("PropertySchemaOption")))
		})
	}

	{
		c.ExpectBlock(&Block{Kind: "Heading", Text: "Multi-select options"})
		c.ExpectParameter(&Parameter{
			Property:     "name",
			Type:         "string",
			Description:  "Name of the option as it appears in Notion.",
			ExampleValue: `"Fruit"`,
		})
		c.ExpectParameter(&Parameter{
			Property:     "color",
			Type:         "optional string (enum)",
			Description:  "Color of the option. Possible values include: default, gray, brown, orange, yellow, green, blue, purple, pink, red.",
			ExampleValue: `"red"`,
		})
	}

	{
		c.ExpectBlock(&Block{Kind: "Heading", Text: "Date configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Date database property schema objects have no additional configuration within the date property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "People configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "People database property schema objects have no additional configuration within the people property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "File configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "File database property schema objects have no additional configuration within the file property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Checkbox configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Checkbox database property schema objects have no additional configuration within the checkbox property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "URL configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "URL database property schema objects have no additional configuration within the url property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Email configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Email database property schema objects have no additional configuration within the email property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Phone number configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Phone number database property schema objects have no additional configuration within the phone_number property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})
	}

	{
		var payload *SimpleObject

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Formula configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Formula database property schema objects contain the following configuration within the formula property:",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload = addPayload(b, "formula", e.Text)
		})

		c.ExpectParameter(&Parameter{
			Property:     "expression",
			Type:         "string",
			Description:  "Formula to evaluate for this property. You can read more about the syntax for formulas in the help center.",
			ExampleValue: `"if(prop(\"In stock\"), 0, prop(\"Price\"))"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String()))
		})
	}

	{
		var payload *UnionStruct

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Relation configuration",
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: `Relation database property objects contain the following configuration within the relation property. In addition, they must contain a key corresponding with the value of type. The value is an object containing type-specific configuration. The type-specific configurations are defined below.`,
		}).Output(func(e *Block, b *CodeBuilder) {
			payload = addPayloadAsUnion(b, "relation", e.Text)
		})

		c.ExpectParameter(&Parameter{
			Property:     "database_id",
			Type:         "string (UUID)",
			Description:  "The database this relation refers to. This database must be shared with the integration.",
			ExampleValue: `"668d797c-76fa-4934-9b05-ad288df2d136"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "type",
			Type:         "string (optional enum)",
			Description:  `The type of the relation. Can be "single_property" or "dual_property".`,
			ExampleValue: `"single_property"`,
		})
		c.ExpectBlock(&Block{Kind: "Heading", Text: "Single property relation configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Single property relation objects have no additional configuration within the single_property property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload.AddPayloadField("single_property", e.Text, WithEmptyStruct())
		})
		c.ExpectBlock(&Block{Kind: "Heading", Text: "Dual property relation configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Dual property relation objects have no additional configuration within the dual_property property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload.AddPayloadField("dual_property", e.Text, WithEmptyStruct())
		})
	}

	{
		var payload *SimpleObject

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Rollup configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Rollup database property objects contain the following configuration within the rollup property:",
		}).Output(func(e *Block, b *CodeBuilder) {
			payload = addPayload(b, "rollup", e.Text)
		})

		c.ExpectParameter(&Parameter{
			Property:     "relation_property_name",
			Type:         "optional string",
			Description:  "The name of the relation property this property is responsible for rolling up. This relation is in the same database where the new rollup property is being created. One of relation_property_name or relation_property_id must be provided.",
			ExampleValue: `"Meals"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})

		c.ExpectParameter(&Parameter{
			Property:     "relation_property_id",
			Type:         "optional string",
			Description:  "The id of the relation property this property is responsible for rolling up. This relation is in the same database where the new rollup property is being created. One of relation_property_name or relation_property_id must be provided.",
			ExampleValue: `"fy:{"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})

		c.ExpectParameter(&Parameter{
			Property:     "rollup_property_name",
			Type:         "optional string",
			Description:  "The name of the property in the related database that is used as an input to function. The related database must be shared with the integration. One of rollup_property_name or rollup_property_id must be provided.",
			ExampleValue: `"Name"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})

		c.ExpectParameter(&Parameter{
			Property:     "rollup_property_id",
			Type:         "optional string",
			Description:  "The id of the property  in the related database that is used as an input to function. The related database must be shared with the integration. One of rollup_property_name or rollup_property_id must be provided.",
			ExampleValue: `"fy:{"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})

		c.ExpectParameter(&Parameter{
			Property:     "function",
			Type:         "string (enum)",
			Description:  "The function that is evaluated for every page in the relation of the rollup.\nPossible values include: count_all, count_values, count_unique_values, count_empty, count_not_empty, percent_empty, percent_not_empty, sum, average, median, min, max, range, show_original",
			ExampleValue: `"count"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			payload.AddFields(b.NewField(e, jen.String(), OmitEmpty))
		})
	}

	{
		c.ExpectBlock(&Block{Kind: "Heading", Text: "Created time configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Created time database property schema objects have no additional configuration within the created_time property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Created by configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Created by database property schema objects have no additional configuration within the created_by property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Last edited time configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Last edited time database property schema objects have no additional configuration within the last_edited_time property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})

		c.ExpectBlock(&Block{Kind: "Heading", Text: "Last edited by configuration"})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "Last edited by database property schema objects have no additional configuration within the last_edited_by property.",
		}).Output(func(e *Block, b *CodeBuilder) {
			addEmptyPayloadByBlock(b, e)
		})
	}

	{
		c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
			addEmptyPayload(b, "button", UNDOCUMENTED)
			payload := addPayload(b, "unique_id", UNDOCUMENTED)
			payload.AddFields(b.NewField(&Parameter{Property: "prefix", Description: UNDOCUMENTED}, jen.Op("*").String()))
		})
	}
}
