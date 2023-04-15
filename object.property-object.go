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
}

// Every database property object contains the following keys:
type propertyCommon struct {
	Id   string `json:"id"`   // An identifier for the property, usually a short string of random letters and symbols.  Some automatically generated property types have special human-readable IDs. For example, all Title properties have an id of "title".
	Name string `json:"name"` // The name of the property as it appears in Notion.
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
	switch string(getRawProperty(data, "type")) {
	case "\"checkbox\"":
		u.value = &CheckboxProperty{}
	case "\"created_by\"":
		u.value = &CreatedByProperty{}
	case "\"created_time\"":
		u.value = &CreatedTimeProperty{}
	case "\"date\"":
		u.value = &DateProperty{}
	case "\"email\"":
		u.value = &EmailProperty{}
	case "\"files\"":
		u.value = &FilesProperty{}
	case "\"formula\"":
		u.value = &FormulaProperty{}
	case "\"last_edited_by\"":
		u.value = &LastEditedByProperty{}
	case "\"last_edited_time\"":
		u.value = &LastEditedTimeProperty{}
	case "\"multi_select\"":
		u.value = &MultiSelectProperty{}
	case "\"number\"":
		u.value = &NumberProperty{}
	case "\"people\"":
		u.value = &PeopleProperty{}
	case "\"phone_number\"":
		u.value = &PhoneNumberProperty{}
	case "\"relation\"":
		u.value = &RelationProperty{}
	case "\"rich_text\"":
		u.value = &RichTextProperty{}
	case "\"rollup\"":
		u.value = &RollupProperty{}
	case "\"select\"":
		u.value = &SelectProperty{}
	case "\"status\"":
		u.value = &StatusProperty{}
	case "\"title\"":
		u.value = &TitleProperty{}
	case "\"url\"":
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
	propertyCommon
	Type     alwaysCheckbox `json:"type"`
	Checkbox struct{}       `json:"checkbox"` //  A checkbox database property is rendered in the Notion UI as a column that contains checkboxes. The checkbox type object is empty; there is no additional property configuration.
}

func (_ *CheckboxProperty) isProperty() {}

// Created by
type CreatedByProperty struct {
	propertyCommon
	Type      alwaysCreatedBy `json:"type"`
	CreatedBy struct{}        `json:"created_by"` //  A created by database property is rendered in the Notion UI as a column that contains people mentions of each row's author as values.   The created_by type object is empty. There is no additional property configuration.
}

func (_ *CreatedByProperty) isProperty() {}

// Created time
type CreatedTimeProperty struct {
	propertyCommon
	Type        alwaysCreatedTime `json:"type"`
	CreatedTime struct{}          `json:"created_time"` //  A created time database property is rendered in the Notion UI as a column that contains timestamps of when each row was created as values.   The created_time type object is empty. There is no additional property configuration.
}

func (_ *CreatedTimeProperty) isProperty() {}

// Date
type DateProperty struct {
	propertyCommon
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
	propertyCommon
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
	propertyCommon
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
	propertyCommon
	Type    alwaysFormula       `json:"type"`
	Formula FormulaPropertyData `json:"formula"`
}

func (_ *FormulaProperty) isProperty() {}

/*
Last edited by
A last edited by database property is rendered in the Notion UI as a column that contains people mentions of the person who last edited each row as values.

The last_edited_by type object is empty. There is no additional property configuration.
*/
type LastEditedByProperty struct {
	propertyCommon
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
	propertyCommon
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
	propertyCommon
	Type        alwaysMultiSelect       `json:"type"`
	MultiSelect MultiSelectPropertyData `json:"multi_select"`
}

func (_ *MultiSelectProperty) isProperty() {}

/*
Number
A number database property is rendered in the Notion UI as a column that contains numeric values. The number type object contains the following fields:
*/
type NumberProperty struct {
	propertyCommon
	Type   alwaysNumber       `json:"type"`
	Number NumberPropertyData `json:"number"`
}

func (_ *NumberProperty) isProperty() {}

type PeopleProperty struct {
	propertyCommon
	Type alwaysPeople `json:"type"`
}

func (_ *PeopleProperty) isProperty() {}

type PhoneNumberProperty struct {
	propertyCommon
	Type alwaysPhoneNumber `json:"type"`
}

func (_ *PhoneNumberProperty) isProperty() {}

type RelationProperty struct {
	propertyCommon
	Type alwaysRelation `json:"type"`
}

func (_ *RelationProperty) isProperty() {}

type RichTextProperty struct {
	propertyCommon
	Type alwaysRichText `json:"type"`
}

func (_ *RichTextProperty) isProperty() {}

type RollupProperty struct {
	propertyCommon
	Type alwaysRollup `json:"type"`
}

func (_ *RollupProperty) isProperty() {}

type SelectProperty struct {
	propertyCommon
	Type alwaysSelect `json:"type"`
}

func (_ *SelectProperty) isProperty() {}

// Status
type StatusProperty struct {
	propertyCommon
	Type   alwaysStatus       `json:"type"`
	Status StatusPropertyData `json:"status"` //  A status database property is rendered in the Notion UI as a column that contains values from a list of status options. The status type object includes an array of options objects and an array of groups objects.   The options array is a sorted list of list of the available status options for the property. Each option object in the array has the following fields:
}

func (_ *StatusProperty) isProperty() {}

/*
Title
A title database property controls the title that appears at the top of a page when a database row is opened. The title type object itself is empty; there is no additional configuration.
*/
type TitleProperty struct {
	propertyCommon
	Type  alwaysTitle `json:"type"`
	Title struct{}    `json:"title"`
}

func (_ *TitleProperty) isProperty() {}

type UrlProperty struct {
	propertyCommon
	Type alwaysUrl `json:"type"`
}

func (_ *UrlProperty) isProperty() {}

/*

A formula database property is rendered in the Notion UI as a column that contains values derived from a provided expression.

The formula type object defines the expression in the following fields:
*/
type FormulaPropertyData struct {
	Expression string `json:"expression"` // The formula that is used to compute the values for this property.   Refer to the Notion help center for information about formula syntax.
}

type MultiSelectPropertyData struct {
	Options []MultiSelectPropertyOption `json:"options"`
}

type MultiSelectPropertyOption struct {
	Color string    `json:"color"` // The color of the option as rendered in the Notion UI. Possible values include:   - blue - brown - default - gray - green - orange - pink - purple - red - yellow
	Id    uuid.UUID `json:"id"`    // An identifier for the option, which does not change if the name is changed. An id is sometimes, but not always, a UUID.
	Name  string    `json:"name"`  // The name of the option as it appears in Notion.  Note: Commas (",") are not valid for multi-select properties.
}

type NumberPropertyData struct {
	Format string `json:"format"` // The way that the number is displayed in Notion. Potential values include:   - argentine_peso - baht - canadian_dollar - chilean_peso - colombian_peso - danish_krone - dirham - dollar - euro - forint - franc - hong_kong_dollar - koruna - krona - leu - lira -  mexican_peso - new_taiwan_dollar - new_zealand_dollar - norwegian_krone - number - number_with_commas - percent - philippine_peso - pound  - rand - real - ringgit - riyal - ruble - rupee - rupiah - shekel - singapore_dollar - uruguayan_peso - yen, - yuan - won - zloty
}

type StatusPropertyData struct {
	Options []StatusPropertyDataOption `json:"options"`
	Groups  []StatusPropertyDataGroup  `json:"groups"`
}

/*

A status database property is rendered in the Notion UI as a column that contains values from a list of status options. The status type object includes an array of options objects and an array of groups objects.

The options array is a sorted list of list of the available status options for the property. Each option object in the array has the following fields:
*/
type StatusPropertyDataOption struct {
	Color string    `json:"color"` // The color of the option as rendered in the Notion UI. Possible values include:   - blue - brown - default - gray - green - orange - pink - purple - red - yellow
	Id    uuid.UUID `json:"id"`    // An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.
	Name  string    `json:"name"`  // The name of the option as it appears in the Notion UI.  Note: Commas (",") are not valid for status values.
}

// A group is a collection of options. The groups array is a sorted list of the available groups for the property. Each group object in the array has the following fields:
type StatusPropertyDataGroup struct {
	Color     string      `json:"color"`      // The color of the option as rendered in the Notion UI. Possible values include:   - blue - brown - default - gray - green - orange - pink - purple - red - yellow
	Id        uuid.UUID   `json:"id"`         // An identifier for the option. The id does not change if the name is changed. It is sometimes, but not always, a UUID.
	Name      string      `json:"name"`       // The name of the option as it appears in the Notion UI.  Note: Commas (",") are not valid for status values.
	OptionIds []uuid.UUID `json:"option_ids"` // A sorted list of ids of all of the options that belong to a group.
}
