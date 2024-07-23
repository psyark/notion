// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-schema-object

package notion

import (
	uuid "github.com/google/uuid"
	json "github.com/psyark/notion/json"
)

/*
Metadata that controls how a database property behaves.

Database properties

Each database property schema object has at least one key which is the property type. This type contains behavior of this property. Possible values of this key are "title", "rich_text", "number", "select", "multi_select", "date", "people", "files", "checkbox", "url", "email", "phone_number", "formula", "relation", "rollup", "created_time", "created_by", "last_edited_time", "last_edited_by".
*/
type PropertySchema struct {
	Title          *struct{}                  `json:"title,omitempty"`            // Each database must have exactly one database property schema object of type "title". This database property controls the title that appears at the top of the page when the page is opened. Title database property objects have no additional configuration within the title property.
	RichText       *struct{}                  `json:"rich_text,omitempty"`        // Text database property schema objects have no additional configuration within the rich_text property.
	Number         *PropertySchemaNumber      `json:"number,omitempty"`           // Number database property schema objects optionally contain the following configuration within the number property.
	Select         *PropertySchemaSelect      `json:"select,omitempty"`           // Select database property schema objects optionally contain the following configuration within the select property:
	MultiSelect    *PropertySchemaMultiSelect `json:"multi_select,omitempty"`     // Multi-select database property schema objects optionally contain the following configuration within the multi_select property:
	Date           *struct{}                  `json:"date,omitempty"`             // Date database property schema objects have no additional configuration within the date property.
	People         *struct{}                  `json:"people,omitempty"`           // People database property schema objects have no additional configuration within the people property.
	Files          *struct{}                  `json:"files,omitempty"`            // File database property schema objects have no additional configuration within the file property.
	Checkbox       *struct{}                  `json:"checkbox,omitempty"`         // Checkbox database property schema objects have no additional configuration within the checkbox property.
	Url            *struct{}                  `json:"url,omitempty"`              // URL database property schema objects have no additional configuration within the url property.
	Email          *struct{}                  `json:"email,omitempty"`            // Email database property schema objects have no additional configuration within the email property.
	PhoneNumber    *struct{}                  `json:"phone_number,omitempty"`     // Phone number database property schema objects have no additional configuration within the phone_number property.
	Formula        *PropertySchemaFormula     `json:"formula,omitempty"`          // Formula database property schema objects contain the following configuration within the formula property:
	Relation       *PropertySchemaRelation    `json:"relation,omitempty"`         // Relation database property objects contain the following configuration within the relation property. In addition, they must contain a key corresponding with the value of type. The value is an object containing type-specific configuration. The type-specific configurations are defined below.
	Rollup         *PropertySchemaRollup      `json:"rollup,omitempty"`           // Rollup database property objects contain the following configuration within the rollup property:
	CreatedTime    *struct{}                  `json:"created_time,omitempty"`     // Created time database property schema objects have no additional configuration within the created_time property.
	CreatedBy      *struct{}                  `json:"created_by,omitempty"`       // Created by database property schema objects have no additional configuration within the created_by property.
	LastEditedTime *struct{}                  `json:"last_edited_time,omitempty"` // Last edited time database property schema objects have no additional configuration within the last_edited_time property.
	LastEditedBy   *struct{}                  `json:"last_edited_by,omitempty"`   // Last edited by database property schema objects have no additional configuration within the last_edited_by property.
	Button         *struct{}                  `json:"button,omitempty"`           // UNDOCUMENTED
	UniqueId       *PropertySchemaUniqueId    `json:"unique_id,omitempty"`        // UNDOCUMENTED
}

// Number database property schema objects optionally contain the following configuration within the number property.
type PropertySchemaNumber struct {
	Format string `json:"format"` // How the number is displayed in Notion. Potential values include: number, number_with_commas, percent, dollar, canadian_dollar, euro, pound, yen, ruble, rupee, won, yuan, real, lira, rupiah, franc, hong_kong_dollar, new_zealand_dollar, krona, norwegian_krone, mexican_peso, rand, new_taiwan_dollar, danish_krone, zloty, baht, forint, koruna, shekel, chilean_peso, philippine_peso, dirham, colombian_peso, riyal, ringgit, leu, argentine_peso, uruguayan_peso, singapore_dollar.
}

// Select database property schema objects optionally contain the following configuration within the select property:
type PropertySchemaSelect struct {
	Options []PropertySchemaOption `json:"options"` // Sorted list of options available for this property.
}

// Select options
type PropertySchemaOption struct {
	Name  string `json:"name"`  // Name of the option as it appears in Notion.
	Color string `json:"color"` // Color of the option. Possible values include: default, gray, brown, orange, yellow, green, blue, purple, pink, red.
}

// Multi-select database property schema objects optionally contain the following configuration within the multi_select property:
type PropertySchemaMultiSelect struct {
	Options []PropertySchemaOption `json:"options"` // Settings for multi select properties.
}

// Formula database property schema objects contain the following configuration within the formula property:
type PropertySchemaFormula struct {
	Expression string `json:"expression"` // Formula to evaluate for this property. You can read more about the syntax for formulas in the help center.
}

// Relation database property objects contain the following configuration within the relation property. In addition, they must contain a key corresponding with the value of type. The value is an object containing type-specific configuration. The type-specific configurations are defined below.
type PropertySchemaRelation struct {
	Type           string    `json:"type"`
	DatabaseId     uuid.UUID `json:"database_id"`     // The database this relation refers to. This database must be shared with the integration.
	SingleProperty struct{}  `json:"single_property"` // Single property relation objects have no additional configuration within the single_property property.
	DualProperty   struct{}  `json:"dual_property"`   // Dual property relation objects have no additional configuration within the dual_property property.
}

func (o PropertySchemaRelation) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.SingleProperty):
			o.Type = "single_property"
		case defined(o.DualProperty):
			o.Type = "dual_property"
		}
	}
	type Alias PropertySchemaRelation
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"dual_property":   o.Type == "dual_property",
		"single_property": o.Type == "single_property",
	}
	return omitFields(data, visibility)
}

// Rollup database property objects contain the following configuration within the rollup property:
type PropertySchemaRollup struct {
	RelationPropertyName string `json:"relation_property_name,omitempty"` // The name of the relation property this property is responsible for rolling up. This relation is in the same database where the new rollup property is being created. One of relation_property_name or relation_property_id must be provided.
	RelationPropertyId   string `json:"relation_property_id,omitempty"`   // The id of the relation property this property is responsible for rolling up. This relation is in the same database where the new rollup property is being created. One of relation_property_name or relation_property_id must be provided.
	RollupPropertyName   string `json:"rollup_property_name,omitempty"`   // The name of the property in the related database that is used as an input to function. The related database must be shared with the integration. One of rollup_property_name or rollup_property_id must be provided.
	RollupPropertyId     string `json:"rollup_property_id,omitempty"`     // The id of the property  in the related database that is used as an input to function. The related database must be shared with the integration. One of rollup_property_name or rollup_property_id must be provided.
	Function             string `json:"function,omitempty"`               // The function that is evaluated for every page in the relation of the rollup. Possible values include: count_all, count_values, count_unique_values, count_empty, count_not_empty, percent_empty, percent_not_empty, sum, average, median, min, max, range, show_original
}

// UNDOCUMENTED
type PropertySchemaUniqueId struct {
	Prefix *string `json:"prefix"` // UNDOCUMENTED
}
