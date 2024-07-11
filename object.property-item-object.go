package notion

import "encoding/json"

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-item-object

// A property_item object describes the identifier, type, and value of a page property. It's returned from the Retrieve a page property item
type PropertyItem struct {
	Type           string             `json:"type"`
	Object         alwaysPropertyItem `json:"object"`           // Always "property_item".
	Id             string             `json:"id"`               // Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string. The id may be used in place of name when creating or updating pages.
	Title          RichText           `json:"title"`            // Title property value objects contain an array of rich text objects within the title property.
	RichText       RichText           `json:"rich_text"`        // Rich Text property value objects contain an array of rich text objects within the rich_text property.
	Number         *float64           `json:"number"`           // Number property value objects contain a number within the number property.
	Select         *Option            `json:"select"`           // Select property value objects contain the following data within the select property:
	Status         *Option            `json:"status"`           // undocumented
	MultiSelect    []Option           `json:"multi_select"`     // Multi-select property value objects contain an array of multi-select option values within the multi_select property.
	Date           *PropertyItemDate  `json:"date"`             // Date property values
	Formula        *Formula           `json:"formula"`          // Formula property value objects represent the result of evaluating a formula described in thedatabase's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
	Relation       *PageReference     `json:"relation"`         // Relation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.
	Rollup         *Rollup            `json:"rollup"`           // Rollup property values
	People         User               `json:"people"`           // People property value objects contain an array of user objects within the people property.
	Files          []File             `json:"files"`            // File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. "Whole_Earth_Catalog.jpg").
	Checkbox       bool               `json:"checkbox"`         // Checkbox property value objects contain a boolean within the checkbox property.
	Url            *string            `json:"url"`              // URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. "http://worrydream.com/EarlyHistoryOfSmalltalk/").
	Email          *string            `json:"email"`            // Email property value objects contain a string within the email property. The string describes an email address (i.e. "hello@example.org").
	PhoneNumber    *string            `json:"phone_number"`     // Phone number property value objects contain a string within the phone_number property. No structure is enforced.
	CreatedTime    ISO8601String      `json:"created_time"`     // Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z").
	CreatedBy      *User              `json:"created_by"`       // Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page.
	LastEditedTime ISO8601String      `json:"last_edited_time"` // Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z").
	LastEditedBy   *User              `json:"last_edited_by"`   // Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page.
	RequestId      string             `json:"request_id"`       // undocumented
}

func (o PropertyItem) isPropertyItemOrPropertyItemPagination() {}
func (o PropertyItem) MarshalJSON() ([]byte, error) {
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
	type Alias PropertyItem
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

// The title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information.
type PaginatedPropertyInfo struct {
	Type     string   `json:"type"`
	Id       string   `json:"id"` // undocumented
	Title    struct{} `json:"title"`
	RichText struct{} `json:"rich_text"`
	Relation struct{} `json:"relation"`
	People   struct{} `json:"people"`
	Rollup   Rollup   `json:"rollup"`   // undocumented
	NextUrl  *string  `json:"next_url"` // The URL the user can request to get the next page of results.
}

func (o PaginatedPropertyInfo) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.Title):
			o.Type = "title"
		case defined(o.RichText):
			o.Type = "rich_text"
		case defined(o.Relation):
			o.Type = "relation"
		case defined(o.People):
			o.Type = "people"
		case defined(o.Rollup):
			o.Type = "rollup"
		}
	}
	type Alias PaginatedPropertyInfo
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"people":    o.Type == "people",
		"relation":  o.Type == "relation",
		"rich_text": o.Type == "rich_text",
		"rollup":    o.Type == "rollup",
		"title":     o.Type == "title",
	}
	return omitFields(data, visibility)
}

// Date property values
type PropertyItemDate struct {
	Start    ISO8601String  `json:"start"`     // An ISO 8601 format date, with optional time.
	End      *ISO8601String `json:"end"`       // An ISO 8601 formatted date, with optional time. Represents the end of a date range. If null, this property's date value is not a range.
	TimeZone *string        `json:"time_zone"` // Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js. When time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information. If null, time zone information will be contained in UTC offsets in start and end.
}

type Rollup struct {
	Type       string            `json:"type"`
	Function   string            `json:"function"`   // Describes the aggregation used. Possible values include: count,  count_values,  empty,  not_empty,  unique,  show_unique,  percent_empty,  percent_not_empty,  sum,  average,  median,  min,  max,  range,  earliest_date,  latest_date,  date_range,  checked,  unchecked,  percent_checked,  percent_unchecked,  count_per_group,  percent_per_group,  show_original
	Number     *float64          `json:"number"`     // Number rollup property values contain a number within the number property.
	Date       *PropertyItemDate `json:"date"`       // Date rollup property values contain a date property value within the date property.
	Array      []PropertyValue   `json:"array"`      // Array rollup property values contain an array of property_item objects within the results property.
	Incomplete struct{}          `json:"incomplete"` // Rollups with an aggregation with more than one page of aggregated results will return a rollup object of type "incomplete". To obtain the final value paginate through the next values in the rollup using the next_cursor or next_url property.
}

func (o Rollup) MarshalJSON() ([]byte, error) {
	if o.Type == "" {
		switch {
		case defined(o.Number):
			o.Type = "number"
		case defined(o.Date):
			o.Type = "date"
		case defined(o.Array):
			o.Type = "array"
		case defined(o.Incomplete):
			o.Type = "incomplete"
		}
	}
	type Alias Rollup
	data, err := json.Marshal(Alias(o))
	if err != nil {
		return nil, err
	}
	visibility := map[string]bool{
		"array":      o.Type == "array",
		"date":       o.Type == "date",
		"incomplete": o.Type == "incomplete",
		"number":     o.Type == "number",
	}
	return omitFields(data, visibility)
}
