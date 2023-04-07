package notion

import uuid "github.com/google/uuid"

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-value-object

/*
A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.

Property values in the page object have a 25 page reference limit
Any property value that has other pages in its value will only use the first 25 page references. Use the [Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint to paginate through the full value.
*/
type PropertyValue interface {
	isPropertyValue()
}

/*
All property values
Each page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
*/
type propertyValueCommon struct {
	Id string `json:"id"` // Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.  The id may be used in place of name when creating or updating pages.
}

/*
Title property values

The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 inline page or person references for a `title` property. If a `title` property includes more than 25 references, then you can use the [Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint for the specific `title` property to get its complete list of references.
*/
type TitlePropertyValue struct {
	Type  string        `always:"title" json:"type"`
	Title RichTextArray `json:"title"` //  Title property value objects contain an array of rich text objects within the title property.
}

func (_ *TitlePropertyValue) isPropertyValue() {}

/*
Rich Text property values

The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 populated inline page or person references for a `rich_text` property. If a `rich_text` property includes more than 25 references, then you can use the [Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `rich_text` property to get its complete list of references.
*/
type RichTextPropertyValue struct {
	Type     string        `always:"rich_text" json:"type"`
	RichText RichTextArray `json:"rich_text"` //  Rich Text property value objects contain an array of rich text objects within the rich_text property.
}

func (_ *RichTextPropertyValue) isPropertyValue() {}

// Number property values
type NumberPropertyValue struct {
	Type   string  `always:"number" json:"type"`
	Number float64 `json:"number"` //  Number property value objects contain a number within the number property.
}

func (_ *NumberPropertyValue) isPropertyValue() {}

// Select property values
type SelectPropertyValue struct {
	Type   string                  `always:"select" json:"type"`
	Select SelectPropertyValueData `json:"select"` //  Select property value objects contain the following data within the select property:
}

func (_ *SelectPropertyValue) isPropertyValue() {}

/*

Select property value objects contain the following data within the select property:
*/
type SelectPropertyValueData struct {
	Id    uuid.UUID `json:"id"`    // ID of the option.  When updating a select property, you can use either name or id.
	Name  string    `json:"name"`  // Name of the option as it appears in Notion.  If the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.  Note: Commas (",") are not valid for select values.
	Color string    `json:"color"` // Color of the option. Possible values are: "default", "gray", "brown", "red", "orange", "yellow", "green", "blue", "purple", "pink". Defaults to "default".  Not currently editable.
}

// Status property values
type StatusPropertyValue struct {
	Type   string                  `always:"status" json:"type"`
	Status StatusPropertyValueData `json:"status"` //  Status property value objects contain the following data within the status property:
}

func (_ *StatusPropertyValue) isPropertyValue() {}

/*

Status property value objects contain the following data within the status property:
*/
type StatusPropertyValueData struct {
	Id    uuid.UUID `json:"id"`    // ID of the option.
	Name  string    `json:"name"`  // Name of the option as it appears in Notion.
	Color string    `json:"color"` // Color of the option. Possible values are: "default", "gray", "brown", "red", "orange", "yellow", "green", "blue", "purple", "pink". Defaults to "default".  Not currently editable.
}

// Multi-select property values
type MultiSelectPropertyValue struct {
	Type        string              `always:"multi_select" json:"type"`
	MultiSelect []MultiSelectOption `json:"multi_select"` //  Multi-select property value objects contain an array of multi-select option values within the multi_select property.
}

func (_ *MultiSelectPropertyValue) isPropertyValue() {}

// Multi-select option values
type MultiSelectOption struct {
	Id    uuid.UUID `json:"id"`    // ID of the option.  When updating a multi-select property, you can use either name or id.
	Name  string    `json:"name"`  // Name of the option as it appears in Notion.  If the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.  Note: Commas (",") are not valid for select values.
	Color string    `json:"color"` // Color of the option. Possible values are: "default", "gray", "brown", "red", "orange", "yellow", "green", "blue", "purple", "pink". Defaults to "default".  Not currently editable.
}

// Date property values
type DatePropertyValue struct {
	Type string                `always:"date" json:"type"`
	Date DatePropertyValueData `json:"date"` //  Date property value objects contain the following data within the date property:
}

func (_ *DatePropertyValue) isPropertyValue() {}

/*

Date property value objects contain the following data within the date property:
*/
type DatePropertyValueData struct {
	Start    ISO8601String `json:"start"`     // An ISO 8601 format date, with optional time.
	End      ISO8601String `json:"end"`       // An ISO 8601 formatted date, with optional time. Represents the end of a date range.  If null, this property's date value is not a range.
	TimeZone string        `json:"time_zone"` // Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.  When time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.  If null, time zone information will be contained in UTC offsets in start and end.
}

/*
Formula property values
Formula property value objects represent the result of evaluating a formula described in the
database's properties. These objects contain a type key and a key corresponding with the value of type. The value of a formula cannot be updated directly.

Formula values may not match the Notion UI.
Formulas returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.
*/
type FormulaPropertyValue struct {
	Type    string  `always:"formula" json:"type"`
	Formula Formula `json:"formula"`
}

func (_ *FormulaPropertyValue) isPropertyValue() {}

//
type Formula interface {
	isFormula()
}

// String formula property values
type StringFormula struct {
	Type   string `always:"string" json:"type"`
	String string `json:"string"` //  String formula property values contain an optional string within the string property.
}

func (_ *StringFormula) isFormula() {}

// Number formula property values
type NumberFormula struct {
	Type   string  `always:"number" json:"type"`
	Number float64 `json:"number"` //  Number formula property values contain an optional number within the number property.
}

func (_ *NumberFormula) isFormula() {}

// Boolean formula property values
type BooleanFormula struct {
	Type    string `always:"boolean" json:"type"`
	Boolean bool   `json:"boolean"` //  Boolean formula property values contain a boolean within the boolean property.
}

func (_ *BooleanFormula) isFormula() {}

// Date formula property values
type DateFormula struct {
	Type string            `always:"date" json:"type"`
	Date DatePropertyValue `json:"date"` //  Date formula property values contain an optional date property value within the date property.
}

func (_ *DateFormula) isFormula() {}

// Relation property values
type RelationPropertyValue struct {
	Type string `always:"relation" json:"type"`
}

func (_ *RelationPropertyValue) isPropertyValue() {}

// Rollup property values
type RollupPropertyValue struct {
	Type string `always:"rollup" json:"type"`
}

func (_ *RollupPropertyValue) isPropertyValue() {}

// People property values
type PeoplePropertyValue struct {
	Type string `always:"people" json:"type"`
}

func (_ *PeoplePropertyValue) isPropertyValue() {}

// Files property values
type FilesPropertyValue struct {
	Type string `always:"files" json:"type"`
}

func (_ *FilesPropertyValue) isPropertyValue() {}

// Checkbox property values
type CheckboxPropertyValue struct {
	Type string `always:"checkbox" json:"type"`
}

func (_ *CheckboxPropertyValue) isPropertyValue() {}

// URL property values
type UrlPropertyValue struct {
	Type string `always:"url" json:"type"`
}

func (_ *UrlPropertyValue) isPropertyValue() {}

// Email property values
type EmailPropertyValue struct {
	Type string `always:"email" json:"type"`
}

func (_ *EmailPropertyValue) isPropertyValue() {}

// Phone number property values
type PhoneNumberPropertyValue struct {
	Type string `always:"phone_number" json:"type"`
}

func (_ *PhoneNumberPropertyValue) isPropertyValue() {}

// Created time property values
type CreatedTimePropertyValue struct {
	Type string `always:"created_time" json:"type"`
}

func (_ *CreatedTimePropertyValue) isPropertyValue() {}

// Created by property values
type CreatedByPropertyValue struct {
	Type string `always:"created_by" json:"type"`
}

func (_ *CreatedByPropertyValue) isPropertyValue() {}

// Last edited time property values
type LastEditedTimePropertyValue struct {
	Type string `always:"last_edited_time" json:"type"`
}

func (_ *LastEditedTimePropertyValue) isPropertyValue() {}

// Last edited by property values
type LastEditedByPropertyValue struct {
	Type string `always:"last_edited_by" json:"type"`
}

func (_ *LastEditedByPropertyValue) isPropertyValue() {}
