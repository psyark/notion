package notion

import (
	"encoding/json"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-object

// All database objects include a child properties object. This properties object is composed of individual database property objects. These property objects define the database schema and are rendered in the Notion UI as database columns.
type Property struct {
	Type           string               `json:"type"`
	Id             string               `json:"id"`               // An identifier for the property, usually a short string of random letters and symbols.      Some automatically generated property types have special human-readable IDs. For example, all Title properties have an id of "title".
	Name           string               `json:"name"`             // The name of the property as it appears in Notion.
	Checkbox       struct{}             `json:"checkbox"`         // Checkbox
	CreatedBy      struct{}             `json:"created_by"`       // Created by
	CreatedTime    struct{}             `json:"created_time"`     // Created time
	Date           struct{}             `json:"date"`             // Date
	Email          struct{}             `json:"email"`            // Email
	Files          struct{}             `json:"files"`            // Files
	Formula        *PropertyFormula     `json:"formula"`          // Formula
	LastEditedBy   struct{}             `json:"last_edited_by"`   // Last edited by
	LastEditedTime struct{}             `json:"last_edited_time"` // Last edited time
	MultiSelect    *PropertyMultiSelect `json:"multi_select"`     // Multi-select
	Number         *PropertyNumber      `json:"number"`           // Number
	People         struct{}             `json:"people"`           // People
	PhoneNumber    struct{}             `json:"phone_number"`     // Phone number
	Relation       *PropertyRelation    `json:"relation"`         // Relation
	RichText       struct{}             `json:"rich_text"`        // Rich text
	Rollup         *PropertyRollup      `json:"rollup"`           // Rollup
	Select         *PropertySelect      `json:"select"`           // Select
	Status         *PropertyStatus      `json:"status"`           // Status
	Title          struct{}             `json:"title"`            // Title
	Url            struct{}             `json:"url"`              // URL
}

func (o Property) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.Checkbox):
			o.Type = "checkbox"
		case defined(o.CreatedBy):
			o.Type = "created_by"
		case defined(o.CreatedTime):
			o.Type = "created_time"
		case defined(o.Date):
			o.Type = "date"
		case defined(o.Email):
			o.Type = "email"
		case defined(o.Files):
			o.Type = "files"
		case defined(o.Formula):
			o.Type = "formula"
		case defined(o.LastEditedBy):
			o.Type = "last_edited_by"
		case defined(o.LastEditedTime):
			o.Type = "last_edited_time"
		case defined(o.MultiSelect):
			o.Type = "multi_select"
		case defined(o.Number):
			o.Type = "number"
		case defined(o.People):
			o.Type = "people"
		case defined(o.PhoneNumber):
			o.Type = "phone_number"
		case defined(o.Relation):
			o.Type = "relation"
		case defined(o.RichText):
			o.Type = "rich_text"
		case defined(o.Rollup):
			o.Type = "rollup"
		case defined(o.Select):
			o.Type = "select"
		case defined(o.Status):
			o.Type = "status"
		case defined(o.Title):
			o.Type = "title"
		case defined(o.Url):
			o.Type = "url"
		}
	}
	type Alias Property
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"checkbox":         o.Type == "checkbox",
		"created_by":       o.Type == "created_by",
		"created_time":     o.Type == "created_time",
		"date":             o.Type == "date",
		"email":            o.Type == "email",
		"files":            o.Type == "files",
		"formula":          o.Type == "formula",
		"last_edited_by":   o.Type == "last_edited_by",
		"last_edited_time": o.Type == "last_edited_time",
		"multi_select":     o.Type == "multi_select",
		"number":           o.Type == "number",
		"people":           o.Type == "people",
		"phone_number":     o.Type == "phone_number",
		"relation":         o.Type == "relation",
		"rich_text":        o.Type == "rich_text",
		"rollup":           o.Type == "rollup",
		"select":           o.Type == "select",
		"status":           o.Type == "status",
		"title":            o.Type == "title",
		"url":              o.Type == "url",
	}
	return omitFields(data, visibility)
}

/*
Formula

A formula database property is rendered in the Notion UI as a column that contains values derived from a provided expression.
*/
type PropertyFormula struct {
	Expression string `json:"expression"` // The formula that is used to compute the values for this property.      Refer to the Notion help center for information about formula syntax.
}

/*
Multi-select

A multi-select database property is rendered in the Notion UI as a column that contains values from a range of options. Each row can contain one or multiple options.
*/
type PropertyMultiSelect struct {
	Options []Option `json:"options"`
}

/*
Number

A number database property is rendered in the Notion UI as a column that contains numeric values. The number type object contains the following fields:
*/
type PropertyNumber struct {
	Format string `json:"format"` // The way that the number is displayed in Notion. Potential values include:      - argentine_peso   - baht   - canadian_dollar   - chilean_peso   - colombian_peso   - danish_krone   - dirham   - dollar   - euro   - forint   - franc   - hong_kong_dollar   - koruna   - krona   - leu   - lira   -  mexican_peso   - new_taiwan_dollar   - new_zealand_dollar   - norwegian_krone   - number   - number_with_commas   - percent   - philippine_peso   - pound   - peruvian_sol   - rand   - real   - ringgit   - riyal   - ruble   - rupee   - rupiah   - shekel   - singapore_dollar   - uruguayan_peso   - yen,   - yuan   - won   - zloty
}

// A relation database property is rendered in the Notion UI as column that contains relations, references to pages in another database, as values.
type PropertyRelation struct {
	Type           string                        `json:"type"`
	SingleProperty struct{}                      `json:"single_property"` // undocumented
	DualProperty   *PropertyRelationDualProperty `json:"dual_property"`   // undocumented
	DatabaseId     uuid.UUID                     `json:"database_id"`     // The database that the relation property refers to.      The corresponding linked page values must belong to the database in order to be valid.
}

func (o PropertyRelation) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.SingleProperty):
			o.Type = "single_property"
		case defined(o.DualProperty):
			o.Type = "dual_property"
		}
	}
	type Alias PropertyRelation
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

// undocumented
type PropertyRelationDualProperty struct {
	SyncedPropertyId   string `json:"synced_property_id"`   // The id of the corresponding property that is updated in the related database when this property is changed.
	SyncedPropertyName string `json:"synced_property_name"` // The name of the corresponding property that is updated in the related database when this property is changed.
}

/*
Rollup

A rollup database property is rendered in the Notion UI as a column with values that are rollups, specific properties that are pulled from a related database.
*/
type PropertyRollup struct {
	Function             string `json:"function"`               // The function that computes the rollup value from the related pages.      Possible values include:      - average   - checked   - count_per_group   - count   - count_values   - date_range   - earliest_date   - empty   - latest_date   - max   - median   - min   - not_empty   - percent_checked   - percent_empty   - percent_not_empty   - percent_per_group   - percent_unchecked   - range   - unchecked   - unique   - show_original   - show_unique   - sum
	RelationPropertyId   string `json:"relation_property_id"`   // The id of the related database property that is rolled up.
	RelationPropertyName string `json:"relation_property_name"` // The name of the related database property that is rolled up.
	RollupPropertyId     string `json:"rollup_property_id"`     // The id of the rollup property.
	RollupPropertyName   string `json:"rollup_property_name"`   // The name of the rollup property.
}

/*
Select

A select database property is rendered in the Notion UI as a column that contains values from a selection of options. Only one option is allowed per row.
*/
type PropertySelect struct {
	Options []Option `json:"options"`
}

// The select type object contains an array of objects representing the available options. Each option object includes the following fields:
type Option struct {
	Color string `json:"color"` // The color of the option as rendered in the Notion UI. Possible values include:      - blue   - brown   - default   - gray   - green   - orange   - pink   - purple   - red   - yellow
	Id    string `json:"id"`    // An identifier for the option. It doesn't change if the name is changed. These are sometimes, but not always, UUIDs.
	Name  string `json:"name"`  // The name of the option as it appears in the Notion UI.      Note: Commas (",") are not valid for select values.
}

// Status
type PropertyStatus struct {
	Options []Option      `json:"options"`
	Groups  []StatusGroup `json:"groups"`
}

// A group is a collection of options. The groups array is a sorted list of the available groups for the property. Each group object in the array has the following fields:
type StatusGroup struct {
	Color     string   `json:"color"`      // The color of the option as rendered in the Notion UI. Possible values include:      - blue   - brown   - default   - gray   - green   - orange   - pink   - purple   - red   - yellow
	Id        string   `json:"id"`         // An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.
	Name      string   `json:"name"`       // The name of the option as it appears in the Notion UI.      Note: Commas (",") are not valid for status values.
	OptionIds []string `json:"option_ids"` // A sorted list of ids of all of the options that belong to a group.
}
