package notion

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-object

// All database objects include a child properties object. This properties object is composed of individual database property objects. These property objects define the database schema and are rendered in the Notion UI as database columns.
type Property interface {
	isProperty()
	GetId() string
	GetName() string
}

// Every database property object contains the following keys:
type PropertyCommon struct {
	Id   string `json:"id"`   // An identifier for the property, usually a short string of random letters and symbols.  Some automatically generated property types have special human-readable IDs. For example, all Title properties have an id of "title".
	Name string `json:"name"` // The name of the property as it appears in Notion.
}

func (c *PropertyCommon) GetId() string {
	return c.Id
}
func (c *PropertyCommon) GetName() string {
	return c.Name
}

type propertyUnmarshaler struct {
	value Property
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *propertyUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch getType(data) {
	case "checkbox":
		u.value = &CheckboxProperty{}
	case "created_by":
		u.value = &CreatedByProperty{}
	case "created_time":
		u.value = &CreatedTimeProperty{}
	case "date":
		u.value = &DateProperty{}
	case "email":
		u.value = &EmailProperty{}
	case "files":
		u.value = &FilesProperty{}
	case "formula":
		u.value = &FormulaProperty{}
	case "last_edited_by":
		u.value = &LastEditedByProperty{}
	case "last_edited_time":
		u.value = &LastEditedTimeProperty{}
	case "multi_select":
		u.value = &MultiSelectProperty{}
	case "number":
		u.value = &NumberProperty{}
	case "people":
		u.value = &PeopleProperty{}
	case "phone_number":
		u.value = &PhoneNumberProperty{}
	case "relation":
		u.value = &RelationProperty{}
	case "rich_text":
		u.value = &RichTextProperty{}
	case "rollup":
		u.value = &RollupProperty{}
	case "select":
		u.value = &SelectProperty{}
	case "status":
		u.value = &StatusProperty{}
	case "title":
		u.value = &TitleProperty{}
	case "url":
		u.value = &UrlProperty{}
	default:
		return fmt.Errorf("unmarshaling Property: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *propertyUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type PropertyMap map[string]Property

func (m *PropertyMap) UnmarshalJSON(data []byte) error {
	t := map[string]propertyUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("unmarshaling PropertyMap: %w", err)
	}
	*m = PropertyMap{}
	for k, u := range t {
		(*m)[k] = u.value
	}
	return nil
}

// Checkbox
type CheckboxProperty struct {
	PropertyCommon
	Type     alwaysCheckbox `json:"type"`
	Checkbox struct{}       `json:"checkbox"` //  A checkbox database property is rendered in the Notion UI as a column that contains checkboxes. The checkbox type object is empty; there is no additional property configuration.
}

func (_ *CheckboxProperty) isProperty() {}

// Created by
type CreatedByProperty struct {
	PropertyCommon
	Type      alwaysCreatedBy `json:"type"`
	CreatedBy struct{}        `json:"created_by"` //  A created by database property is rendered in the Notion UI as a column that contains people mentions of each row's author as values.   The created_by type object is empty. There is no additional property configuration.
}

func (_ *CreatedByProperty) isProperty() {}

// Created time
type CreatedTimeProperty struct {
	PropertyCommon
	Type        alwaysCreatedTime `json:"type"`
	CreatedTime struct{}          `json:"created_time"` //  A created time database property is rendered in the Notion UI as a column that contains timestamps of when each row was created as values.   The created_time type object is empty. There is no additional property configuration.
}

func (_ *CreatedTimeProperty) isProperty() {}

// Date
type DateProperty struct {
	PropertyCommon
	Type alwaysDate `json:"type"`
	Date struct{}   `json:"date"` //  A date database property is rendered in the Notion UI as a column that contains date values.   The date type object is empty; there is no additional configuration.
}

func (_ *DateProperty) isProperty() {}

/*
Email
An email database property is represented in the Notion UI as a column that contains email values.

The email type object is empty. There is no additional property configuration.
*/
type EmailProperty struct {
	PropertyCommon
	Type  alwaysEmail `json:"type"`
	Email struct{}    `json:"email"`
}

func (_ *EmailProperty) isProperty() {}

/*
Files
The Notion API does not yet support uploading files to Notion.
A files database property is rendered in the Notion UI as a column that has values that are either files uploaded directly to Notion or external links to files. The files type object is empty; there is no additional configuration.
*/
type FilesProperty struct {
	PropertyCommon
	Type  alwaysFiles `json:"type"`
	Files struct{}    `json:"files"`
}

func (_ *FilesProperty) isProperty() {}

/*
Formula
A formula database property is rendered in the Notion UI as a column that contains values derived from a provided expression.

The formula type object defines the expression in the following fields:
*/
type FormulaProperty struct {
	PropertyCommon
	Type    alwaysFormula       `json:"type"`
	Formula FormulaPropertyData `json:"formula"`
}

func (_ *FormulaProperty) isProperty() {}

type FormulaPropertyData struct {
	Expression string `json:"expression"` // The formula that is used to compute the values for this property.   Refer to the Notion help center for information about formula syntax.
}

/*
Last edited by
A last edited by database property is rendered in the Notion UI as a column that contains people mentions of the person who last edited each row as values.

The last_edited_by type object is empty. There is no additional property configuration.
*/
type LastEditedByProperty struct {
	PropertyCommon
	Type         alwaysLastEditedBy `json:"type"`
	LastEditedBy struct{}           `json:"last_edited_by"`
}

func (_ *LastEditedByProperty) isProperty() {}

/*
Last edited time
A last edited time database property is rendered in the Notion UI as a column that contains timestamps of when each row was last edited as values.

The last_edited_time type object is empty. There is no additional property configuration.
*/
type LastEditedTimeProperty struct {
	PropertyCommon
	Type           alwaysLastEditedTime `json:"type"`
	LastEditedTime struct{}             `json:"last_edited_time"`
}

func (_ *LastEditedTimeProperty) isProperty() {}

/*
Multi-select
A multi-select database property is rendered in the Notion UI as a column that contains values from a range of options. Each row can contain one or multiple options.

The multi_select type object includes an array of options objects. Each option object details settings for the option, indicating the following fields:
*/
type MultiSelectProperty struct {
	PropertyCommon
	Type        alwaysMultiSelect       `json:"type"`
	MultiSelect MultiSelectPropertyData `json:"multi_select"`
}

func (_ *MultiSelectProperty) isProperty() {}

type MultiSelectPropertyData struct {
	Options []Option `json:"options"`
}

/*
Number
A number database property is rendered in the Notion UI as a column that contains numeric values. The number type object contains the following fields:
*/
type NumberProperty struct {
	PropertyCommon
	Type   alwaysNumber       `json:"type"`
	Number NumberPropertyData `json:"number"`
}

func (_ *NumberProperty) isProperty() {}

type NumberPropertyData struct {
	Format string `json:"format"` // The way that the number is displayed in Notion. Potential values include:   - argentine_peso - baht - canadian_dollar - chilean_peso - colombian_peso - danish_krone - dirham - dollar - euro - forint - franc - hong_kong_dollar - koruna - krona - leu - lira -  mexican_peso - new_taiwan_dollar - new_zealand_dollar - norwegian_krone - number - number_with_commas - percent - philippine_peso - pound  - rand - real - ringgit - riyal - ruble - rupee - rupiah - shekel - singapore_dollar - uruguayan_peso - yen, - yuan - won - zloty
}

/*
People
A people database property is rendered in the Notion UI as a column that contains people mentions.  The people type object is empty; there is no additional configuration.
*/
type PeopleProperty struct {
	PropertyCommon
	Type   alwaysPeople `json:"type"`
	People struct{}     `json:"people"`
}

func (_ *PeopleProperty) isProperty() {}

/*
Phone number
A phone number database property is rendered in the Notion UI as a column that contains phone number values.

The phone_number type object is empty. There is no additional property configuration.
*/
type PhoneNumberProperty struct {
	PropertyCommon
	Type        alwaysPhoneNumber `json:"type"`
	PhoneNumber struct{}          `json:"phone_number"`
}

func (_ *PhoneNumberProperty) isProperty() {}

/*
Relation
A relation database property is rendered in the Notion UI as column that contains relations, references to pages in another database, as values.

The relation type object contains the following fields:
*/
type RelationProperty struct {
	PropertyCommon
	Type     alwaysRelation `json:"type"`
	Relation Relation       `json:"relation"`
}

func (_ *RelationProperty) isProperty() {}
func (o *RelationProperty) UnmarshalJSON(data []byte) error {
	type Alias RelationProperty
	t := &struct {
		*Alias
		Relation relationUnmarshaler `json:"relation"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling RelationProperty: %w", err)
	}
	o.Relation = t.Relation.value
	return nil
}

/*
Rich text
A rich text database property is rendered in the Notion UI as a column that contains text values. The rich_text type object is empty; there is no additional configuration.
*/
type RichTextProperty struct {
	PropertyCommon
	Type     alwaysRichText `json:"type"`
	RichText struct{}       `json:"rich_text"`
}

func (_ *RichTextProperty) isProperty() {}

// Rollup
type RollupProperty struct {
	PropertyCommon
	Type   alwaysRollup       `json:"type"`
	Rollup RollupPropertyData `json:"rollup"`
}

func (_ *RollupProperty) isProperty() {}

/*

A rollup database property is rendered in the Notion UI as a column with values that are rollups, specific properties that are pulled from a related database.

The rollup type object contains the following fields:
*/
type RollupPropertyData struct {
	Function             string `json:"function"`               // The function that computes the rollup value from the related pages.  Possible values include:   - average - checked - count_per_group - count - count_values  - date_range - earliest_date  - empty - latest_date - max - median - min - not_empty - percent_checked - percent_empty - percent_not_empty - percent_per_group - percent_unchecked - range - unchecked - unique - show_original - show_unique - sum
	RelationPropertyId   string `json:"relation_property_id"`   // The id of the related database property that is rolled up.
	RelationPropertyName string `json:"relation_property_name"` // The name of the related database property that is rolled up.
	RollupPropertyId     string `json:"rollup_property_id"`     // The id of the rollup property.
	RollupPropertyName   string `json:"rollup_property_name"`   // The name of the rollup property.
}

/*

A select database property is rendered in the Notion UI as a column that contains values from a selection of options. Only one option is allowed per row.

The select type object contains an array of objects representing the available options. Each option object includes the following fields:
*/
type SelectProperty struct {
	PropertyCommon
	Type   alwaysSelect       `json:"type"`
	Select SelectPropertyData `json:"select"`
}

func (_ *SelectProperty) isProperty() {}

type SelectPropertyData struct {
	Options []Option `json:"options"`
}

// Status
type StatusProperty struct {
	PropertyCommon
	Type   alwaysStatus       `json:"type"`
	Status StatusPropertyData `json:"status"`
}

func (_ *StatusProperty) isProperty() {}

type StatusPropertyData struct {
	Options []Option      `json:"options"`
	Groups  []StatusGroup `json:"groups"`
}

/*
Title
A title database property controls the title that appears at the top of a page when a database row is opened. The title type object itself is empty; there is no additional configuration.
*/
type TitleProperty struct {
	PropertyCommon
	Type  alwaysTitle `json:"type"`
	Title struct{}    `json:"title"`
}

func (_ *TitleProperty) isProperty() {}

/*
URL
A URL database property is represented in the Notion UI as a column that contains URL values.

The url type object is empty. There is no additional property configuration.
*/
type UrlProperty struct {
	PropertyCommon
	Type alwaysUrl `json:"type"`
	Url  struct{}  `json:"url"`
}

func (_ *UrlProperty) isProperty() {}

/*

A relation database property is rendered in the Notion UI as column that contains relations, references to pages in another database, as values.

The relation type object contains the following fields:
*/
type Relation interface {
	isRelation()
	GetDatabaseId() uuid.UUID
}
type RelationCommon struct {
	DatabaseId uuid.UUID `json:"database_id"` // The database that the relation property refers to.   The corresponding linked page values must belong to the database in order to be valid.
}

func (c *RelationCommon) GetDatabaseId() uuid.UUID {
	return c.DatabaseId
}

type relationUnmarshaler struct {
	value Relation
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *relationUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch getType(data) {
	case "single_property":
		u.value = &SinglePropertyRelation{}
	case "dual_property":
		u.value = &DualPropertyRelation{}
	default:
		return fmt.Errorf("unmarshaling Relation: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *relationUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// undocumented
type SinglePropertyRelation struct {
	RelationCommon
	Type           alwaysSingleProperty `json:"type"`
	SingleProperty struct{}             `json:"single_property"`
}

func (_ *SinglePropertyRelation) isRelation() {}

// undocumented
type DualPropertyRelation struct {
	RelationCommon
	Type         alwaysDualProperty       `json:"type"`
	DualProperty DualPropertyRelationData `json:"dual_property"`
}

func (_ *DualPropertyRelation) isRelation() {}

type DualPropertyRelationData struct {
	SyncedPropertyId   string `json:"synced_property_id"`   // The id of the corresponding property that is updated in the related database when this property is changed.
	SyncedPropertyName string `json:"synced_property_name"` // The name of the corresponding property that is updated in the related database when this property is changed.
}

/*

A select database property is rendered in the Notion UI as a column that contains values from a selection of options. Only one option is allowed per row.

The select type object contains an array of objects representing the available options. Each option object includes the following fields:
*/
type Option struct {
	Color string `json:"color"` // The color of the option as rendered in the Notion UI. Possible values include:   - blue - brown - default - gray - green - orange - pink - purple - red - yellow
	Id    string `json:"id"`    // An identifier for the option. It doesn't change if the name is changed. These are sometimes, but not always, UUIDs.
	Name  string `json:"name"`  // The name of the option as it appears in the Notion UI.  Note: Commas (",") are not valid for select values.
}

// A group is a collection of options. The groups array is a sorted list of the available groups for the property. Each group object in the array has the following fields:
type StatusGroup struct {
	Color     string   `json:"color"`      // The color of the option as rendered in the Notion UI. Possible values include:   - blue - brown - default - gray - green - orange - pink - purple - red - yellow
	Id        string   `json:"id"`         // An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.
	Name      string   `json:"name"`       // The name of the option as it appears in the Notion UI.  Note: Commas (",") are not valid for status values.
	OptionIds []string `json:"option_ids"` // A sorted list of ids of all of the options that belong to a group.
}
