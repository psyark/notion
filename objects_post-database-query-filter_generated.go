// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/post-database-query-filter

package notion

import uuid "github.com/google/uuid"

/*
The filter object

🚧The timestamp filter condition does not require a property name. The API throws an error if you provide one.
*/
type Filter struct {
	Property       string             `json:"property,omitempty"`         // The name of the property as it appears in the database, or the property ID.
	Checkbox       *FilterCheckbox    `json:"checkbox,omitempty"`         // Checkbox
	Date           *FilterDate        `json:"date,omitempty"`             // Date
	Files          *FilterFiles       `json:"files,omitempty"`            // Files
	Formula        *FilterFormula     `json:"formula,omitempty"`          // Formula
	MultiSelect    *FilterMultiSelect `json:"multi_select,omitempty"`     // Multi-select
	Number         *FilterNumber      `json:"number,omitempty"`           // Number
	People         *FilterPeople      `json:"people,omitempty"`           // People
	Relation       *FilterRelation    `json:"relation,omitempty"`         // Relation
	RichText       *FilterRichText    `json:"rich_text,omitempty"`        // Rich text
	Rollup         *FilterRollup      `json:"rollup,omitempty"`           // Rollup
	Select         *FilterSelect      `json:"select,omitempty"`           // Select
	Status         *FilterStatus      `json:"status,omitempty"`           // Status
	Timestamp      string             `json:"timestamp,omitempty"`        // A constant string representing the type of timestamp to use as a filter.
	CreatedTime    *FilterDate        `json:"created_time,omitempty"`     // A date filter condition used to filter the specified timestamp.
	LastEditedTime *FilterDate        `json:"last_edited_time,omitempty"` // A date filter condition used to filter the specified timestamp.
	And            []Filter           `json:"and,omitempty"`              // An array of filter objects or compound filter conditions. Returns database entries that match all of the provided filter conditions.
	Or             []Filter           `json:"or,omitempty"`               // An array of filter objects or compound filter conditions. Returns database entries that match any of the provided filter conditions
}

// Checkbox
type FilterCheckbox struct {
	Equals       *bool `json:"equals,omitempty"`         // Whether a checkbox property value matches the provided value exactly. Returns or excludes all database entries with an exact value match.
	DoesNotEqual *bool `json:"does_not_equal,omitempty"` // Whether a checkbox property value differs from the provided value. Returns or excludes all database entries with a difference in values.
}

/*
Date

📘For the after, before, equals, on_or_before, and on_or_after fields, if a date string with a time is provided, then the comparison is done with millisecond precision.If no timezone is provided, then the timezone defaults to UTC.

A date filter condition can be used to limit date property value types and the timestamp property types created_time and last_edited_time.
*/
type FilterDate struct {
	After      ISO8601String `json:"after,omitempty"`        // The value to compare the date property value against. Returns database entries where the date property value is after the provided date.
	Before     ISO8601String `json:"before,omitempty"`       // The value to compare the date property value against. Returns database entries where the date property value is before the provided date.
	Equals     ISO8601String `json:"equals,omitempty"`       // The value to compare the date property value against. Returns database entries where the date property value is the provided date.
	IsEmpty    bool          `json:"is_empty,omitempty"`     // The value to compare the date property value against. Returns database entries where the date property value contains no data.
	IsNotEmpty bool          `json:"is_not_empty,omitempty"` // The value to compare the date property value against. Returns database entries where the date property value is not empty.
	NextMonth  *struct{}     `json:"next_month,omitempty"`   // A filter that limits the results to database entries where the date property value is within the next month.
	NextWeek   *struct{}     `json:"next_week,omitempty"`    // A filter that limits the results to database entries where the date property value is within the next week.
	NextYear   *struct{}     `json:"next_year,omitempty"`    // A filter that limits the results to database entries where the date property value is within the next year.
	OnOrAfter  ISO8601String `json:"on_or_after,omitempty"`  // The value to compare the date property value against. Returns database entries where the date property value is on or after the provided date.
	OnOrBefore ISO8601String `json:"on_or_before,omitempty"` // The value to compare the date property value against. Returns database entries where the date property value is on or before the provided date.
	PastMonth  *struct{}     `json:"past_month,omitempty"`   // A filter that limits the results to database entries where the date property value is within the past month.
	PastWeek   *struct{}     `json:"past_week,omitempty"`    // A filter that limits the results to database entries where the date property value is within the past week.
	PastYear   *struct{}     `json:"past_year,omitempty"`    // A filter that limits the results to database entries where the date property value is within the past year.
	ThisWeek   *struct{}     `json:"this_week,omitempty"`    // A filter that limits the results to database entries where the date property value is this week.
}

// Files
type FilterFiles struct {
	IsEmpty    bool `json:"is_empty,omitempty"`     // Whether the files property value does not contain any data. Returns all database entries with an empty files property value.
	IsNotEmpty bool `json:"is_not_empty,omitempty"` // Whether the files property value contains data. Returns all entries with a populated files property value.
}

/*
Formula

The primary field of the formula filter condition object matches the type of the formula’s result. For example, to filter a formula property that computes a checkbox, use a formula filter condition object with a checkbox field containing a checkbox filter condition as its value.
*/
type FilterFormula struct {
	Checkbox *FilterCheckbox `json:"checkbox,omitempty"` // A checkbox filter condition to compare the formula result against. Returns database entries where the formula result matches the provided condition.
	Date     *FilterDate     `json:"date,omitempty"`     // A date filter condition to compare the formula result against. Returns database entries where the formula result matches the provided condition.
	Number   *FilterNumber   `json:"number,omitempty"`   // A number filter condition to compare the formula result against. Returns database entries where the formula result matches the provided condition.
	String   *FilterRichText `json:"string,omitempty"`   // A rich text filter condition to compare the formula result against. Returns database entries where the formula result matches the provided condition.
}

// Multi-select
type FilterMultiSelect struct {
	Contains       string `json:"contains,omitempty"`         // The value to compare the multi-select property value against. Returns database entries where the multi-select value matches the provided string.
	DoesNotContain string `json:"does_not_contain,omitempty"` // The value to compare the multi-select property value against. Returns database entries where the multi-select value does not match the provided string.
	IsEmpty        bool   `json:"is_empty,omitempty"`         // Whether the multi-select property value is empty. Returns database entries where the multi-select value does not contain any data.
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`     // Whether the multi-select property value is not empty. Returns database entries where the multi-select value does contains data.
}

// Number
type FilterNumber struct {
	DoesNotEqual         *float64 `json:"does_not_equal,omitempty"`           // The number to compare the number property value against. Returns database entries where the number property value differs from the provided number.
	Equals               *float64 `json:"equals,omitempty"`                   // The number to compare the number property value against. Returns database entries where the number property value is the same as the provided number.
	GreaterThan          *float64 `json:"greater_than,omitempty"`             // The number to compare the number property value against. Returns database entries where the number property value exceeds the provided number.
	GreaterThanOrEqualTo *float64 `json:"greater_than_or_equal_to,omitempty"` // The number to compare the number property value against. Returns database entries where the number property value is equal to or exceeds the provided number.
	IsEmpty              bool     `json:"is_empty,omitempty"`                 // Whether the number property value is empty. Returns database entries where the number property value does not contain any data.
	IsNotEmpty           bool     `json:"is_not_empty,omitempty"`             // Whether the number property value is not empty. Returns database entries where the number property value contains data.
	LessThan             *float64 `json:"less_than,omitempty"`                // The number to compare the number property value against. Returns database entries where the number property value is less than the provided number.
	LessThanOrEqualTo    *float64 `json:"less_than_or_equal_to,omitempty"`    // The number to compare the number property value against. Returns database entries where the number property value is equal to or is less than the provided number.
}

/*
People

You can apply a people filter condition to people, created_by, and last_edited_by database property types.
*/
type FilterPeople struct {
	Contains       *uuid.UUID `json:"contains,omitempty"`         // The value to compare the people property value against. Returns database entries where the people property value contains the provided string.
	DoesNotContain *uuid.UUID `json:"does_not_contain,omitempty"` // The value to compare the people property value against. Returns database entries where the people property value does not contain the provided string.
	IsEmpty        bool       `json:"is_empty,omitempty"`         // Whether the people property value does not contain any data. Returns database entries where the people property value does not contain any data.
	IsNotEmpty     bool       `json:"is_not_empty,omitempty"`     // Whether the people property value contains data. Returns database entries where the people property value is not empty.
}

// Relation
type FilterRelation struct {
	Contains       *uuid.UUID `json:"contains,omitempty"`         // The value to compare the relation property value against. Returns database entries where the relation property value contains the provided string.
	DoesNotContain *uuid.UUID `json:"does_not_contain,omitempty"` // The value to compare the relation property value against. Returns entries where the relation property value does not contain the provided string.
	IsEmpty        bool       `json:"is_empty,omitempty"`         // Whether the relation property value does not contain data. Returns database entries where the relation property value does not contain any data.
	IsNotEmpty     bool       `json:"is_not_empty,omitempty"`     // Whether the relation property value contains data. Returns database entries where the property value is not empty.
}

// Rich text
type FilterRichText struct {
	Contains       string `json:"contains,omitempty"`         // The string to compare the text property value against. Returns database entries with a text property value that includes the provided string.
	DoesNotContain string `json:"does_not_contain,omitempty"` // The string to compare the text property value against. Returns database entries with a text property value that does not include the provided string.
	DoesNotEqual   string `json:"does_not_equal,omitempty"`   // The string to compare the text property value against. Returns database entries with a text property value that does not match the provided string.
	EndsWith       string `json:"ends_with,omitempty"`        // The string to compare the text property value against. Returns database entries with a text property value that ends with the provided string.
	Equals         string `json:"equals,omitempty"`           // The string to compare the text property value against. Returns database entries with a text property value that matches the provided string.
	IsEmpty        bool   `json:"is_empty,omitempty"`         // Whether the text property value does not contain any data. Returns database entries with a text property value that is empty.
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`     // Whether the text property value contains any data. Returns database entries with a text property value that contains data.
	StartsWith     string `json:"starts_with,omitempty"`      // The string to compare the text property value against. Returns database entries with a text property value that starts with the provided string.
}

/*
Rollup

A rollup database property can evaluate to an array, date, or number value. The filter condition for the rollup property contains a rollup key and a corresponding object value that depends on the computed value type.
*/
type FilterRollup struct {
	Any    *Filter       `json:"any,omitempty"`    // The value to compare each rollup property value against. Can be a filter condition for any other type. Returns database entries where the rollup property value matches the provided criteria.
	Every  *Filter       `json:"every,omitempty"`  // The value to compare each rollup property value against. Can be a filter condition for any other type. Returns database entries where every rollup property value matches the provided criteria.
	None   *Filter       `json:"none,omitempty"`   // The value to compare each rollup property value against. Can be a filter condition for any other type. Returns database entries where no rollup property value matches the provided criteria.
	Date   *FilterDate   `json:"date,omitempty"`   // A date filter condition to compare the rollup value against. Returns database entries where the rollup value matches the provided condition.
	Number *FilterNumber `json:"number,omitempty"` // A number filter condition to compare the rollup value against. Returns database entries where the rollup value matches the provided condition.
}

// Select
type FilterSelect struct {
	Equals       string `json:"equals,omitempty"`         // The string to compare the select property value against. Returns database entries where the select property value matches the provided string.
	DoesNotEqual string `json:"does_not_equal,omitempty"` // The string to compare the select property value against. Returns database entries where the select property value does not match the provided string.
	IsEmpty      bool   `json:"is_empty,omitempty"`       // Whether the select property value does not contain data. Returns database entries where the select property value is empty.
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`   // Whether the select property value contains data. Returns database entries where the select property value is not empty.
}

// Status
type FilterStatus struct {
	Equals       string `json:"equals,omitempty"`         // The string to compare the status property value against. Returns database entries where the status property value matches the provided string.
	DoesNotEqual string `json:"does_not_equal,omitempty"` // The string to compare the status property value against. Returns database entries where the status property value does not match the provided string.
	IsEmpty      bool   `json:"is_empty,omitempty"`       // Whether the status property value does not contain data. Returns database entries where the status property value is empty.
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`   // Whether the status property value contains data. Returns database entries where the status property value is not empty.
}
