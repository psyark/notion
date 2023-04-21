package notion

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/post-database-query-filter

/*
When you query a database, you can send a filter object in the body of the request that limits the returned entries based on the specified criteria.

For example, the below query limits the response to entries where the "Task completed"  checkbox property value is true:
The filter object
*/
type Filter interface {
	isFilter()
	GetProperty() string
}

/*

Each filter object contains the following fields:
*/
type filterCommon struct {
	Property string `json:"property"` // The name of the property as it appears in the database, or the property ID.
}

func (c *filterCommon) GetProperty() string {
	return c.Property
}

// Checkbox
type CheckboxFilter struct {
	filterCommon
	Equals       *bool `json:"equals,omitempty"`         // Whether a checkbox property value matches the provided value exactly.  Returns or excludes all database entries with an exact value match.
	DoesNotEqual *bool `json:"does_not_equal,omitempty"` // Whether a checkbox property value differs from the provided value.   Returns or excludes all database entries with a difference in values.
}

func (_ *CheckboxFilter) isFilter() {}

/*
Date
For the `after`, `before`, `equals, on_or_before`, and `on_or_after` fields, if a date string with a time is provided, then the comparison is done with millisecond precision.

If no timezone is provided, then the timezone defaults to UTC.
A date filter condition can be used to limit date property value types and the timestamp property types created_time and last_edited_time.

The condition contains the below fields:
*/
type DateFilter struct {
	filterCommon
	After      ISO8601String `json:"after,omitempty"`        // The value to compare the date property value against.   Returns database entries where the date property value is after the provided date.
	Before     ISO8601String `json:"before,omitempty"`       // The value to compare the date property value against.  Returns database entries where the date property value is before the provided date.
	Equals     ISO8601String `json:"equals,omitempty"`       // The value to compare the date property value against.  Returns database entries where the date property value is the provided date.
	IsEmpty    bool          `json:"is_empty,omitempty"`     // The value to compare the date property value against.  Returns database entries where the date property value contains no data.
	IsNotEmpty bool          `json:"is_not_empty,omitempty"` // The value to compare the date property value against.  Returns database entries where the date property value is not empty.
	NextMonth  *struct{}     `json:"next_month,omitempty"`   // A filter that limits the results to database entries where the date property value is within the next month.
	NextWeek   *struct{}     `json:"next_week,omitempty"`    // A filter that limits the results to database entries where the date property value is within the next week.
	NextYear   *struct{}     `json:"next_year,omitempty"`    // A filter that limits the results to database entries where the date property value is within the next year.
	OnOrAfter  ISO8601String `json:"on_or_after,omitempty"`  // The value to compare the date property value against.  Returns database entries where the date property value is on or after the provided date.
	OnOrBefore ISO8601String `json:"on_or_before,omitempty"` // The value to compare the date property value against.   Returns database entries where the date property value is on or before the provided date.
	PastMonth  *struct{}     `json:"past_month,omitempty"`   // A filter that limits the results to database entries where the date property value is within the past month.
	PastWeek   *struct{}     `json:"past_week,omitempty"`    // A filter that limits the results to database entries where the date property value is within the past week.
	PastYear   *struct{}     `json:"past_year,omitempty"`    // A filter that limits the results to database entries where the date property value is within the past year.
	ThisWeek   *struct{}     `json:"this_week,omitempty"`    // A filter that limits the results to database entries where the date property value is this week.
}

func (_ *DateFilter) isFilter() {}

// Files
type FileFilter struct {
	filterCommon
	IsEmpty    bool `json:"is_empty,omitempty"`     // Whether the files property value does not contain any data.  Returns all database entries with an empty files property value.
	IsNotEmpty bool `json:"is_not_empty,omitempty"` // Whether the files property value contains data.   Returns all entries with a populated files property value.
}

func (_ *FileFilter) isFilter() {}

// Rich text
type RichTextFilter struct {
	filterCommon
	Contains       string `json:"contains,omitempty"`         // The string to compare the text property value against.  Returns database entries with a text property value that includes the provided string.
	DoesNotContain string `json:"does_not_contain,omitempty"` // The string to compare the text property value against.  Returns database entries with a text property value that does not include the provided string.
	DoesNotEqual   string `json:"does_not_equal,omitempty"`   // The string to compare the text property value against.  Returns database entries with a text property value that does not match the provided string.
	EndsWith       string `json:"ends_with,omitempty"`        // The string to compare the text property value against.  Returns database entries with a text property value that ends with the provided string.
	Equals         string `json:"equals,omitempty"`           // The string to compare the text property value against.  Returns database entries with a text property value that matches the provided string.
	IsEmpty        bool   `json:"is_empty,omitempty"`         // Whether the text property value does not contain any data.   Returns database entries with a text property value that is empty.
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`     // Whether the text property value contains any data.   Returns database entries with a text property value that contains data.
	StartsWith     string `json:"starts_with,omitempty"`      // The string to compare the text property value against.  Returns database entries with a text property value that starts with the provided string.
}

func (_ *RichTextFilter) isFilter() {}