// Code generated by notion.doc2api; DO NOT EDIT.

package notion

import "encoding/json"

// https://developers.notion.com/reference/property-value-object

/*
A property value defines the identifier, type, and value of a page property in a page object. It's used when retrieving and updating pages, ex: Create and Update pages.

Any property value that has other pages in its value will only use the first 25 page references. Use the Retrieve a page property endpoint to paginate through the full value.

Each page property value object contains the following keys. In addition, it contains a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
*/
type PropertyValue struct {
	Type           string             `json:"type,omitempty"`
	Id             string             `json:"id,omitempty"`     // Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string. The id may be used in place of name when creating or updating pages.
	Title          []RichText         `json:"title"`            // Title property value objects contain an array of rich text objects within the title property.
	RichText       []RichText         `json:"rich_text"`        // Rich Text property value objects contain an array of rich text objects within the rich_text property.
	Number         *float64           `json:"number"`           // Number property value objects contain a number within the number property.
	Select         *Option            `json:"select"`           // Select property value objects contain the following data within the select property:
	Status         *Option            `json:"status"`           // Status property value objects contain the following data within the status property:
	MultiSelect    []Option           `json:"multi_select"`     // Multi-select property value objects contain an array of multi-select option values within the multi_select property.
	Date           *PropertyValueDate `json:"date"`             // Date property value objects contain the following data within the date property:
	Formula        *Formula           `json:"formula"`          // Formula property value objects represent the result of evaluating a formula described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a formula cannot be updated directly.
	Relation       []PageReference    `json:"relation"`         // Relation property value objects contain an array of page references within the relation property. A page reference is an object with an id key and a string value (UUIDv4) corresponding to a page ID in another database.
	HasMore        bool               `json:"has_more"`         // A relation includes a has_more property in the Retrieve a page endpoint response object. The endpoint returns a maximum of 25 page references for a relation. If a relation has more than 25 references, then the has_more value for the relation in the response object is true. If a relation doesn’t exceed the limit, then has_more is false.
	Rollup         *Rollup            `json:"rollup"`           // Rollup property value objects represent the result of evaluating a rollup described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value of a rollup cannot be updated directly.
	People         []User             `json:"people"`           // People property value objects contain an array of user objects within the people property.
	Files          []File             `json:"files"`            // File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. "Whole_Earth_Catalog.jpg").
	Checkbox       bool               `json:"checkbox"`         // Checkbox property value objects contain a boolean within the checkbox property.
	Url            *string            `json:"url"`              // URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. "http://worrydream.com/EarlyHistoryOfSmalltalk/").
	Email          *string            `json:"email"`            // Email property value objects contain a string within the email property. The string describes an email address (i.e. "hello@example.org").
	PhoneNumber    *string            `json:"phone_number"`     // Phone number property value objects contain a string within the phone_number property. No structure is enforced.
	CreatedTime    ISO8601String      `json:"created_time"`     // Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z"). The value of created_time cannot be updated. See the Property Item Object to see how these values are returned.
	CreatedBy      User               `json:"created_by"`       // Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page. The value of created_by cannot be updated. See the Property Item Object to see how these values are returned.
	LastEditedTime ISO8601String      `json:"last_edited_time"` // Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z"). The value of last_edited_time cannot be updated. See the Property Item Object to see how these values are returned.
	LastEditedBy   User               `json:"last_edited_by"`   // Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page. The value of last_edited_by cannot be updated. See the Property Item Object to see how these values are returned.
}

func (o PropertyValue) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.Title):
			o.Type = "title"
		case defined(o.RichText):
			o.Type = "rich_text"
		case defined(o.Number):
			o.Type = "number"
		case defined(o.Select):
			o.Type = "select"
		case defined(o.Status):
			o.Type = "status"
		case defined(o.MultiSelect):
			o.Type = "multi_select"
		case defined(o.Date):
			o.Type = "date"
		case defined(o.Formula):
			o.Type = "formula"
		case defined(o.Relation):
			o.Type = "relation"
		case defined(o.HasMore):
			o.Type = "relation"
		case defined(o.Rollup):
			o.Type = "rollup"
		case defined(o.People):
			o.Type = "people"
		case defined(o.Files):
			o.Type = "files"
		case defined(o.Checkbox):
			o.Type = "checkbox"
		case defined(o.Url):
			o.Type = "url"
		case defined(o.Email):
			o.Type = "email"
		case defined(o.PhoneNumber):
			o.Type = "phone_number"
		case defined(o.CreatedTime):
			o.Type = "created_time"
		case defined(o.CreatedBy):
			o.Type = "created_by"
		case defined(o.LastEditedTime):
			o.Type = "last_edited_time"
		case defined(o.LastEditedBy):
			o.Type = "last_edited_by"
		}
	}
	type Alias PropertyValue
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
		"has_more":         o.Type == "relation",
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

// Date property value objects contain the following data within the date property:
type PropertyValueDate struct {
	Start    ISO8601String  `json:"start"`     // An ISO 8601 format date, with optional time.
	End      *ISO8601String `json:"end"`       // An ISO 8601 formatted date, with optional time. Represents the end of a date range. If null, this property's date value is not a range.
	TimeZone *string        `json:"time_zone"` // Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js. When time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information. If null, time zone information will be contained in UTC offsets in start and end.
}

// Formula property values
type Formula struct {
	Type    string            `json:"type"`
	String  *string           `json:"string"`  // String formula property values contain an optional string within the string property.
	Number  float64           `json:"number"`  // Number formula property values contain an optional number within the number property.
	Boolean bool              `json:"boolean"` // Boolean formula property values contain a boolean within the boolean property.
	Date    PropertyValueDate `json:"date"`    // Date formula property values contain an optional date property value within the date property.
}

func (o Formula) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.String):
			o.Type = "string"
		case defined(o.Number):
			o.Type = "number"
		case defined(o.Boolean):
			o.Type = "boolean"
		case defined(o.Date):
			o.Type = "date"
		}
	}
	type Alias Formula
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"boolean": o.Type == "boolean",
		"date":    o.Type == "date",
		"number":  o.Type == "number",
		"string":  o.Type == "string",
	}
	return omitFields(data, visibility)
}