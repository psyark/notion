package notion

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
	nullv4 "gopkg.in/guregu/null.v4"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-value-object

/*
A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.

Property values in the page object have a 25 page reference limit
Any property value that has other pages in its value will only use the first 25 page references. Use the [Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint to paginate through the full value.
Each page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
*/
type PropertyValue interface {
	isPropertyValue()
}
type propertyValueCommon struct {
	Id string `json:"id"` // Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.  The id may be used in place of name when creating or updating pages.
}

type propertyValueUnmarshaler struct {
	value PropertyValue
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *propertyValueUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch string(getRawProperty(data, "type")) {
	case "\"title\"":
		u.value = &TitlePropertyValue{}
	case "\"rich_text\"":
		u.value = &RichTextPropertyValue{}
	case "\"number\"":
		u.value = &NumberPropertyValue{}
	case "\"select\"":
		u.value = &SelectPropertyValue{}
	case "\"status\"":
		u.value = &StatusPropertyValue{}
	case "\"multi_select\"":
		u.value = &MultiSelectPropertyValue{}
	case "\"date\"":
		u.value = &DatePropertyValue{}
	case "\"formula\"":
		u.value = &FormulaPropertyValue{}
	case "\"relation\"":
		u.value = &RelationPropertyValue{}
	case "\"rollup\"":
		u.value = &RollupPropertyValue{}
	case "\"people\"":
		u.value = &PeoplePropertyValue{}
	case "\"files\"":
		u.value = &FilesPropertyValue{}
	case "\"checkbox\"":
		u.value = &CheckboxPropertyValue{}
	case "\"url\"":
		u.value = &UrlPropertyValue{}
	case "\"email\"":
		u.value = &EmailPropertyValue{}
	case "\"phone_number\"":
		u.value = &PhoneNumberPropertyValue{}
	case "\"created_time\"":
		u.value = &CreatedTimePropertyValue{}
	case "\"created_by\"":
		u.value = &CreatedByPropertyValue{}
	case "\"last_edited_time\"":
		u.value = &LastEditedTimePropertyValue{}
	case "\"last_edited_by\"":
		u.value = &LastEditedByPropertyValue{}
	default:
		return fmt.Errorf("unmarshaling PropertyValue: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *propertyValueUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type PropertyValueMap map[string]PropertyValue

func (m *PropertyValueMap) UnmarshalJSON(data []byte) error {
	t := map[string]propertyValueUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*m = PropertyValueMap{}
	for k, u := range t {
		(*m)[k] = u.value
	}
	return nil
}

/*
Title property values

The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 inline page or person references for a `title` property. If a `title` property includes more than 25 references, then you can use the [Retrieve a page property](https://developers.notion.com/reference/retrieve-a-page-property) endpoint for the specific `title` property to get its complete list of references.
*/
type TitlePropertyValue struct {
	propertyValueCommon
	Type  string        `always:"title" json:"type"`
	Title RichTextArray `json:"title"` //  Title property value objects contain an array of rich text objects within the title property.
}

func (_ *TitlePropertyValue) isPropertyValue() {}

/*
Rich Text property values

The [Retrieve a page endpoint](https://developers.notion.com/reference/retrieve-a-page) returns a maximum of 25 populated inline page or person references for a `rich_text` property. If a `rich_text` property includes more than 25 references, then you can use the [Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `rich_text` property to get its complete list of references.
*/
type RichTextPropertyValue struct {
	propertyValueCommon
	Type     string        `always:"rich_text" json:"type"`
	RichText RichTextArray `json:"rich_text"` //  Rich Text property value objects contain an array of rich text objects within the rich_text property.
}

func (_ *RichTextPropertyValue) isPropertyValue() {}

// Number property values
type NumberPropertyValue struct {
	propertyValueCommon
	Type   string       `always:"number" json:"type"`
	Number nullv4.Float `json:"number"` //  Number property value objects contain a number within the number property.
}

func (_ *NumberPropertyValue) isPropertyValue() {}

// Select property values
type SelectPropertyValue struct {
	propertyValueCommon
	Type   string                   `always:"select" json:"type"`
	Select *SelectPropertyValueData `json:"select"` //  Select property value objects contain the following data within the select property:
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
	propertyValueCommon
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
	propertyValueCommon
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
	propertyValueCommon
	Type string                 `always:"date" json:"type"`
	Date *DatePropertyValueData `json:"date"` //  Date property value objects contain the following data within the date property:
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
	propertyValueCommon
	Type    string  `always:"formula" json:"type"`
	Formula Formula `json:"formula"`
}

func (_ *FormulaPropertyValue) isPropertyValue() {}

type Formula interface {
	isFormula()
}

type formulaUnmarshaler struct {
	value Formula
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *formulaUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch string(getRawProperty(data, "type")) {
	case "\"string\"":
		u.value = &StringFormula{}
	case "\"number\"":
		u.value = &NumberFormula{}
	case "\"boolean\"":
		u.value = &BooleanFormula{}
	case "\"date\"":
		u.value = &DateFormula{}
	default:
		return fmt.Errorf("unmarshaling Formula: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *formulaUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
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

/*
Relation property values
Relation property value objects contain an array of page references within the relation property. A page reference is an object with an id key and a string value (UUIDv4) corresponding to a page ID in another database.

A relation includes a has_more property in the Retrieve a page endpoint response object. The endpoint returns a maximum of 25 page references for a relation. If a relation has more than 25 references, then the has_more value for the relation in the response object is true. If a relation doesn’t exceed the limit, then has_more is false.

Note that updating a relation property value with an empty array will clear the list.
*/
type RelationPropertyValue struct {
	propertyValueCommon
	Type     string          `always:"relation" json:"type"`
	Relation []PageReference `json:"relation"`
	HasMore  bool            `json:"has_more"`
}

func (_ *RelationPropertyValue) isPropertyValue() {}

/*
Rollup property values
Rollup property value objects represent the result of evaluating a rollup described in the
database's properties. These objects contain a type key and a key corresponding with the value of type. The value of a rollup cannot be updated directly.

Rollup values may not match the Notion UI.
Rollups returned in page objects are subject to a 25 page reference limitation. The Retrieve a page property endpoint should be used to get an accurate formula value.
*/
type RollupPropertyValue struct {
	propertyValueCommon
	Type string `always:"rollup" json:"type"`
}

func (_ *RollupPropertyValue) isPropertyValue() {}

type Rollup interface {
	isRollup()
}

type rollupUnmarshaler struct {
	value Rollup
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *rollupUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	switch string(getRawProperty(data, "type")) {
	case "\"string\"":
		u.value = &StringRollup{}
	case "\"number\"":
		u.value = &NumberRollup{}
	case "\"date\"":
		u.value = &DateRollup{}
	case "\"array\"":
		u.value = &ArrayRollup{}
	default:
		return fmt.Errorf("unmarshaling Rollup: data has unknown type field: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *rollupUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// String rollup property values
type StringRollup struct {
	Type   string `always:"string" json:"type"`
	String string `json:"string"` //  String rollup property values contain an optional string within the string property.
}

func (_ *StringRollup) isRollup() {}

// Number rollup property values
type NumberRollup struct {
	Type   string  `always:"number" json:"type"`
	Number float64 `json:"number"` //  Number rollup property values contain a number within the number property.
}

func (_ *NumberRollup) isRollup() {}

// Date rollup property values
type DateRollup struct {
	Type string            `always:"date" json:"type"`
	Date DatePropertyValue `json:"date"` //  Date rollup property values contain a date property value within the date property.
}

func (_ *DateRollup) isRollup() {}

// Array rollup property values
type ArrayRollup struct {
	Type  string   `always:"array" json:"type"`
	Array []Rollup `json:"array"` //  Array rollup property values contain an array of number, date, or string objects within the results property.
}

func (_ *ArrayRollup) isRollup() {}

/*
People property values

The [Retrieve a page](https://developers.notion.com/reference/retrieve-a-page) endpoint can’t be guaranteed to return more than 25 people per `people` page property. If a `people` page property includes more than 25 people, then you can use the [Retrieve a page property endpoint](https://developers.notion.com/reference/retrieve-a-page-property) for the specific `people` property to get a complete list of people.
*/
type PeoplePropertyValue struct {
	propertyValueCommon
	Type   string `always:"people" json:"type"`
	People Users  `json:"people"` //  People property value objects contain an array of user objects within the people property.
}

func (_ *PeoplePropertyValue) isPropertyValue() {}

/*
Files property values

When updating a file property, the value will be overwritten by the array of files passed.
Although we do not support uploading files, if you pass a `file` object containing a file hosted by Notion, it will remain one of the files. To remove any file, just do not pass it in the update response.
*/
type FilesPropertyValue struct {
	propertyValueCommon
	Type  string `always:"files" json:"type"`
	Files []File `json:"files"` //  File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. "Whole_Earth_Catalog.jpg").
}

func (_ *FilesPropertyValue) isPropertyValue() {}

// Checkbox property values
type CheckboxPropertyValue struct {
	propertyValueCommon
	Type     string `always:"checkbox" json:"type"`
	Checkbox bool   `json:"checkbox"` //  Checkbox property value objects contain a boolean within the checkbox property.
}

func (_ *CheckboxPropertyValue) isPropertyValue() {}

// URL property values
type UrlPropertyValue struct {
	propertyValueCommon
	Type string        `always:"url" json:"type"`
	Url  nullv4.String `json:"url"` //  URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. "http://worrydream.com/EarlyHistoryOfSmalltalk/").
}

func (_ *UrlPropertyValue) isPropertyValue() {}

// Email property values
type EmailPropertyValue struct {
	propertyValueCommon
	Type  string        `always:"email" json:"type"`
	Email nullv4.String `json:"email"` //  Email property value objects contain a string within the email property. The string describes an email address (i.e. "hello@example.org").
}

func (_ *EmailPropertyValue) isPropertyValue() {}

// Phone number property values
type PhoneNumberPropertyValue struct {
	propertyValueCommon
	Type        string `always:"phone_number" json:"type"`
	PhoneNumber string `json:"phone_number"` //  Phone number property value objects contain a string within the phone_number property. No structure is enforced.
}

func (_ *PhoneNumberPropertyValue) isPropertyValue() {}

// Created time property values
type CreatedTimePropertyValue struct {
	propertyValueCommon
	Type        string        `always:"created_time" json:"type"`
	CreatedTime ISO8601String `json:"created_time"` //  Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z"). The value of created_time cannot be updated. See the Property Item Object to see how these values are returned.
}

func (_ *CreatedTimePropertyValue) isPropertyValue() {}

// Created by property values
type CreatedByPropertyValue struct {
	propertyValueCommon
	Type      string `always:"created_by" json:"type"`
	CreatedBy User   `json:"created_by"` //  Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page. The value of created_by cannot be updated. See the Property Item Object to see how these values are returned.
}

func (_ *CreatedByPropertyValue) isPropertyValue() {}

// Last edited time property values
type LastEditedTimePropertyValue struct {
	propertyValueCommon
	Type           string        `always:"last_edited_time" json:"type"`
	LastEditedTime ISO8601String `json:"last_edited_time"` //  Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z"). The value of last_edited_time cannot be updated. See the Property Item Object to see how these values are returned.
}

func (_ *LastEditedTimePropertyValue) isPropertyValue() {}

// Last edited by property values
type LastEditedByPropertyValue struct {
	propertyValueCommon
	Type         string `always:"last_edited_by" json:"type"`
	LastEditedBy User   `json:"last_edited_by"` //  Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page. The value of last_edited_by cannot be updated. See the Property Item Object to see how these values are returned.
}

func (_ *LastEditedByPropertyValue) isPropertyValue() {}
