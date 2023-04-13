package notion

import (
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-item-object

/*
A property_item object describes the identifier, type, and value of a page property. It's returned from the Retrieve a page property item
*/
type PropertyItem interface {
	isPropertyItem()
	isPropertyItemOrPropertyItemPagination()
}

/*

Each page property item object contains the following keys. In addition, it will contain a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
*/
type propertyItemCommon struct {
	propertyItemOrPropertyItemPaginationCommon
	Object string `always:"property_item" json:"object"` // Always "property_item".
	Id     string `json:"id"`                            // Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.  The id may be used in place of name when creating or updating pages.
}

type propertyItemUnmarshaler struct {
	value PropertyItem
}

func (u *propertyItemUnmarshaler) UnmarshalJSON(data []byte) error {
	switch string(getRawProperty(data, "type")) {
	case "\"title\"":
		u.value = &TitlePropertyItem{}
	case "\"rich_text\"":
		u.value = &RichTextPropertyItem{}
	case "\"number\"":
		u.value = &NumberPropertyItem{}
	case "\"select\"":
		u.value = &SelectPropertyItem{}
	case "\"multi_select\"":
		u.value = &MultiSelectPropertyItem{}
	case "\"date\"":
		u.value = &DatePropertyItem{}
	case "\"formula\"":
		u.value = &FormulaPropertyItem{}
	case "\"relation\"":
		u.value = &RelationPropertyItem{}
	case "\"rollup\"":
		u.value = &RollupPropertyItem{}
	case "\"people\"":
		u.value = &PeoplePropertyItem{}
	case "\"files\"":
		u.value = &FilesPropertyItem{}
	case "\"checkbox\"":
		u.value = &CheckboxPropertyItem{}
	case "\"url\"":
		u.value = &UrlPropertyItem{}
	case "\"email\"":
		u.value = &EmailPropertyItem{}
	case "\"phone_number\"":
		u.value = &PhoneNumberPropertyItem{}
	case "\"created_time\"":
		u.value = &CreatedTimePropertyItem{}
	case "\"created_by\"":
		u.value = &CreatedByPropertyItem{}
	case "\"last_edited_time\"":
		u.value = &LastEditedTimePropertyItem{}
	case "\"last_edited_by\"":
		u.value = &LastEditedByPropertyItem{}
	default:
		return fmt.Errorf("unknown type: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

/*
Paginated property values
The title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information.
*/
type PropertyItemPagination struct {
	paginationCommon
	propertyItemOrPropertyItemPaginationCommon
	Object       string         `always:"list" json:"object"`        // Always "list".
	Type         string         `always:"property_item" json:"type"` // Always "property_item".
	Results      []PropertyItem `json:"results"`                     // List of property_item objects.
	PropertyItem PropertyItem   `json:"property_item"`               // A property_item object that describes the property.
	NextUrl      string         `json:"next_url"`                    // The URL the user can request to get the next page of results.
}

func (_ *PropertyItemPagination) isPagination()                           {}
func (_ *PropertyItemPagination) isPropertyItemOrPropertyItemPagination() {}

// Title property values
type TitlePropertyItem struct {
	propertyItemCommon
	Type  string        `always:"title" json:"type"`
	Title RichTextArray `json:"title"` //  Title property value objects contain an array of rich text objects within the title property.
}

func (_ *TitlePropertyItem) isPropertyItem()                         {}
func (_ *TitlePropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Rich Text property values
type RichTextPropertyItem struct {
	propertyItemCommon
	Type     string        `always:"rich_text" json:"type"`
	RichText RichTextArray `json:"rich_text"` //  Rich Text property value objects contain an array of rich text objects within the rich_text property.
}

func (_ *RichTextPropertyItem) isPropertyItem()                         {}
func (_ *RichTextPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Number property values
type NumberPropertyItem struct {
	propertyItemCommon
	Type   string  `always:"number" json:"type"`
	Number float64 `json:"number"` //  Number property value objects contain a number within the number property.
}

func (_ *NumberPropertyItem) isPropertyItem()                         {}
func (_ *NumberPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Select property values
type SelectPropertyItem struct {
	propertyItemCommon
	Type   string                 `always:"select" json:"type"`
	Select SelectPropertyItemData `json:"select"` //  Select property value objects contain the following data within the select property:
}

func (_ *SelectPropertyItem) isPropertyItem()                         {}
func (_ *SelectPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*

Select property value objects contain the following data within the select property:
*/
type SelectPropertyItemData struct {
	Id    uuid.UUID `json:"id"`    // ID of the option.  When updating a select property, you can use either name or id.
	Name  string    `json:"name"`  // Name of the option as it appears in Notion.  If the select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.  Note: Commas (",") are not valid for select values.
	Color string    `json:"color"` // Color of the option. Possible values are: "default", "gray", "brown", "red", "orange", "yellow", "green", "blue", "purple", "pink". Defaults to "default".  Not currently editable.
}

// Multi-select property values
type MultiSelectPropertyItem struct {
	propertyItemCommon
	Type        string                        `always:"multi_select" json:"type"`
	MultiSelect []MultiSelectPropertyItemData `json:"multi_select"` //  Multi-select property value objects contain an array of multi-select option values within the multi_select property.
}

func (_ *MultiSelectPropertyItem) isPropertyItem()                         {}
func (_ *MultiSelectPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*

Multi-select property value objects contain an array of multi-select option values within the multi_select property.
*/
type MultiSelectPropertyItemData struct {
	Id    uuid.UUID `json:"id"`    // ID of the option.  When updating a multi-select property, you can use either name or id.
	Name  string    `json:"name"`  // Name of the option as it appears in Notion.  If the multi-select database property does not yet have an option by that name, it will be added to the database schema if the integration also has write access to the parent database.  Note: Commas (",") are not valid for select values.
	Color string    `json:"color"` // Color of the option. Possible values are: "default", "gray", "brown", "red", "orange", "yellow", "green", "blue", "purple", "pink". Defaults to "default".  Not currently editable.
}

// Date property values
type DatePropertyItem struct {
	propertyItemCommon
	Type string               `always:"date" json:"type"`
	Date DatePropertyItemData `json:"date"` //  Date property value objects contain the following data within the date property:
}

func (_ *DatePropertyItem) isPropertyItem()                         {}
func (_ *DatePropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*

Date property value objects contain the following data within the date property:
*/
type DatePropertyItemData struct {
	Start    ISO8601String `json:"start"`     // An ISO 8601 format date, with optional time.
	End      ISO8601String `json:"end"`       // An ISO 8601 formatted date, with optional time. Represents the end of a date range.  If null, this property's date value is not a range.
	TimeZone string        `json:"time_zone"` // Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.  When time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.  If null, time zone information will be contained in UTC offsets in start and end.
}

// Formula property values
type FormulaPropertyItem struct {
	propertyItemCommon
	Type    string  `always:"formula" json:"type"`
	Formula Formula `json:"formula"` //  Formula property value objects represent the result of evaluating a formula described in the  database's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
}

func (_ *FormulaPropertyItem) isPropertyItem()                         {}
func (_ *FormulaPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Relation property values
type RelationPropertyItem struct {
	propertyItemCommon
	Type     string          `always:"relation" json:"type"`
	Relation []PageReference `json:"relation"` //  Relation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.
}

func (_ *RelationPropertyItem) isPropertyItem()                         {}
func (_ *RelationPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Rollup property values
type RollupPropertyItem struct {
	propertyItemCommon
	Type   string `always:"rollup" json:"type"`
	Rollup Rollup `json:"rollup"` //  Rollup property value objects represent the result of evaluating a rollup described in the  database's properties. The property is returned as a list object of type property_item with a list of relation items used to computed the rollup under results.   A rollup property item is also returned under the property_type key that describes the rollup aggregation and computed result.   In order to avoid timeouts, if the rollup has a with a large number of aggregations or properties the endpoint returns a next_cursor value that is used to determinate the aggregation value so far for the subset of relations that have been paginated through.   Once has_more is false, then the final rollup value is returned.  See the Pagination documentation for more information on pagination in the Notion API.   Computing the values of following aggregations are not supported. Instead the endpoint returns a list of property_item objects for the rollup: * show_unique (Show unique values) * unique (Count unique values) * median(Median)
}

func (_ *RollupPropertyItem) isPropertyItem()                         {}
func (_ *RollupPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// People property values
type PeoplePropertyItem struct {
	propertyItemCommon
	Type string `always:"people" json:"type"`
}

func (_ *PeoplePropertyItem) isPropertyItem()                         {}
func (_ *PeoplePropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Files property values
type FilesPropertyItem struct {
	propertyItemCommon
	Type string `always:"files" json:"type"`
}

func (_ *FilesPropertyItem) isPropertyItem()                         {}
func (_ *FilesPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Checkbox property values
type CheckboxPropertyItem struct {
	propertyItemCommon
	Type string `always:"checkbox" json:"type"`
}

func (_ *CheckboxPropertyItem) isPropertyItem()                         {}
func (_ *CheckboxPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// URL property values
type UrlPropertyItem struct {
	propertyItemCommon
	Type string `always:"url" json:"type"`
}

func (_ *UrlPropertyItem) isPropertyItem()                         {}
func (_ *UrlPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Email property values
type EmailPropertyItem struct {
	propertyItemCommon
	Type string `always:"email" json:"type"`
}

func (_ *EmailPropertyItem) isPropertyItem()                         {}
func (_ *EmailPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Phone number property values
type PhoneNumberPropertyItem struct {
	propertyItemCommon
	Type string `always:"phone_number" json:"type"`
}

func (_ *PhoneNumberPropertyItem) isPropertyItem()                         {}
func (_ *PhoneNumberPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Created time property values
type CreatedTimePropertyItem struct {
	propertyItemCommon
	Type string `always:"created_time" json:"type"`
}

func (_ *CreatedTimePropertyItem) isPropertyItem()                         {}
func (_ *CreatedTimePropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Created by property values
type CreatedByPropertyItem struct {
	propertyItemCommon
	Type string `always:"created_by" json:"type"`
}

func (_ *CreatedByPropertyItem) isPropertyItem()                         {}
func (_ *CreatedByPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Last edited time property values
type LastEditedTimePropertyItem struct {
	propertyItemCommon
	Type string `always:"last_edited_time" json:"type"`
}

func (_ *LastEditedTimePropertyItem) isPropertyItem()                         {}
func (_ *LastEditedTimePropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Last edited by property values
type LastEditedByPropertyItem struct {
	propertyItemCommon
	Type string `always:"last_edited_by" json:"type"`
}

func (_ *LastEditedByPropertyItem) isPropertyItem()                         {}
func (_ *LastEditedByPropertyItem) isPropertyItemOrPropertyItemPagination() {}
