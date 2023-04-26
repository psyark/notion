package notion

import (
	"encoding/json"
	"fmt"
	nullv4 "gopkg.in/guregu/null.v4"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-item-object

/*
A property_item object describes the identifier, type, and value of a page property. It's returned from the Retrieve a page property item
*/
type PropertyItem interface {
	PropertyItemOrPropertyItemPagination
	isPropertyItem()
	GetObject() alwaysPropertyItem
	GetId() string
}

/*

Each page property item object contains the following keys. In addition, it will contain a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
*/
type PropertyItemCommon struct {
	Object alwaysPropertyItem `json:"object"` // Always "property_item".
	Id     string             `json:"id"`     // Underlying identifier for the property. This identifier is guaranteed to remain constant when the property name changes. It may be a UUID, but is often a short random string.  The id may be used in place of name when creating or updating pages.
}

func (c *PropertyItemCommon) GetObject() alwaysPropertyItem {
	return c.Object
}
func (c *PropertyItemCommon) GetId() string {
	return c.Id
}

type propertyItemUnmarshaler struct {
	value PropertyItem
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *propertyItemUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	t := struct {
		Title          json.RawMessage `json:"title"`
		RichText       json.RawMessage `json:"rich_text"`
		Number         json.RawMessage `json:"number"`
		Select         json.RawMessage `json:"select"`
		Status         json.RawMessage `json:"status"`
		MultiSelect    json.RawMessage `json:"multi_select"`
		Date           json.RawMessage `json:"date"`
		Formula        json.RawMessage `json:"formula"`
		Relation       json.RawMessage `json:"relation"`
		Rollup         json.RawMessage `json:"rollup"`
		People         json.RawMessage `json:"people"`
		Files          json.RawMessage `json:"files"`
		Checkbox       json.RawMessage `json:"checkbox"`
		Url            json.RawMessage `json:"url"`
		Email          json.RawMessage `json:"email"`
		PhoneNumber    json.RawMessage `json:"phone_number"`
		CreatedTime    json.RawMessage `json:"created_time"`
		CreatedBy      json.RawMessage `json:"created_by"`
		LastEditedTime json.RawMessage `json:"last_edited_time"`
		LastEditedBy   json.RawMessage `json:"last_edited_by"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	switch {
	case t.Title != nil:
		u.value = &TitlePropertyItem{}
	case t.RichText != nil:
		u.value = &RichTextPropertyItem{}
	case t.Number != nil:
		u.value = &NumberPropertyItem{}
	case t.Select != nil:
		u.value = &SelectPropertyItem{}
	case t.Status != nil:
		u.value = &StatusPropertyItem{}
	case t.MultiSelect != nil:
		u.value = &MultiSelectPropertyItem{}
	case t.Date != nil:
		u.value = &DatePropertyItem{}
	case t.Formula != nil:
		u.value = &FormulaPropertyItem{}
	case t.Relation != nil:
		u.value = &RelationPropertyItem{}
	case t.Rollup != nil:
		u.value = &RollupPropertyItem{}
	case t.People != nil:
		u.value = &PeoplePropertyItem{}
	case t.Files != nil:
		u.value = &FilesPropertyItem{}
	case t.Checkbox != nil:
		u.value = &CheckboxPropertyItem{}
	case t.Url != nil:
		u.value = &UrlPropertyItem{}
	case t.Email != nil:
		u.value = &EmailPropertyItem{}
	case t.PhoneNumber != nil:
		u.value = &PhoneNumberPropertyItem{}
	case t.CreatedTime != nil:
		u.value = &CreatedTimePropertyItem{}
	case t.CreatedBy != nil:
		u.value = &CreatedByPropertyItem{}
	case t.LastEditedTime != nil:
		u.value = &LastEditedTimePropertyItem{}
	case t.LastEditedBy != nil:
		u.value = &LastEditedByPropertyItem{}
	default:
		return fmt.Errorf("unmarshal PropertyItem: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *propertyItemUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type PropertyItemList []PropertyItem

func (a *PropertyItemList) UnmarshalJSON(data []byte) error {
	t := []propertyItemUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("unmarshaling PropertyItemList: %w", err)
	}
	*a = make([]PropertyItem, len(t))
	for i, u := range t {
		(*a)[i] = u.value
	}
	return nil
}

type PropertyItemMap map[string]PropertyItem

func (m *PropertyItemMap) UnmarshalJSON(data []byte) error {
	t := map[string]propertyItemUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("unmarshaling PropertyItemMap: %w", err)
	}
	*m = PropertyItemMap{}
	for k, u := range t {
		(*m)[k] = u.value
	}
	return nil
}

/*

The title, rich_text, relation and people property items of are returned as a paginated list object of individual property_item objects in the results. An abridged set of the the properties found in the list object are found below, see the Pagination documentation for additional information.
*/
type PaginatedPropertyInfo interface {
	isPaginatedPropertyInfo()
	GetId() string
	GetNextUrl() *string
}
type PaginatedPropertyInfoCommon struct {
	Id      string  `json:"id"`
	NextUrl *string `json:"next_url"` // The URL the user can request to get the next page of results.
}

func (c *PaginatedPropertyInfoCommon) GetId() string {
	return c.Id
}
func (c *PaginatedPropertyInfoCommon) GetNextUrl() *string {
	return c.NextUrl
}

type paginatedPropertyInfoUnmarshaler struct {
	value PaginatedPropertyInfo
}

/*
UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance
according to the "type" field of the message.
*/
func (u *paginatedPropertyInfoUnmarshaler) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.value = nil
		return nil
	}
	t := struct {
		Title    json.RawMessage `json:"title"`
		RichText json.RawMessage `json:"rich_text"`
		Relation json.RawMessage `json:"relation"`
		People   json.RawMessage `json:"people"`
		Rollup   json.RawMessage `json:"rollup"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	switch {
	case t.Title != nil:
		u.value = &TitlePaginatedPropertyInfo{}
	case t.RichText != nil:
		u.value = &RichTextPaginatedPropertyInfo{}
	case t.Relation != nil:
		u.value = &RelationPaginatedPropertyInfo{}
	case t.People != nil:
		u.value = &PeoplePaginatedPropertyInfo{}
	case t.Rollup != nil:
		u.value = &RollupPaginatedPropertyInfo{}
	default:
		return fmt.Errorf("unmarshal PaginatedPropertyInfo: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *paginatedPropertyInfoUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

type TitlePaginatedPropertyInfo struct {
	PaginatedPropertyInfoCommon
	Type  alwaysTitle `json:"type"`
	Title struct{}    `json:"title"`
}

func (_ *TitlePaginatedPropertyInfo) isPaginatedPropertyInfo() {}

type RichTextPaginatedPropertyInfo struct {
	PaginatedPropertyInfoCommon
	Type     alwaysRichText `json:"type"`
	RichText struct{}       `json:"rich_text"`
}

func (_ *RichTextPaginatedPropertyInfo) isPaginatedPropertyInfo() {}

type RelationPaginatedPropertyInfo struct {
	PaginatedPropertyInfoCommon
	Type     alwaysRelation `json:"type"`
	Relation struct{}       `json:"relation"`
}

func (_ *RelationPaginatedPropertyInfo) isPaginatedPropertyInfo() {}

type PeoplePaginatedPropertyInfo struct {
	PaginatedPropertyInfoCommon
	Type   alwaysPeople `json:"type"`
	People struct{}     `json:"people"`
}

func (_ *PeoplePaginatedPropertyInfo) isPaginatedPropertyInfo() {}

// undocumented
type RollupPaginatedPropertyInfo struct {
	PaginatedPropertyInfoCommon
	Type   alwaysRollup `json:"type"`
	Rollup Rollup       `json:"rollup"`
}

func (_ *RollupPaginatedPropertyInfo) isPaginatedPropertyInfo() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *RollupPaginatedPropertyInfo) UnmarshalJSON(data []byte) error {
	type Alias RollupPaginatedPropertyInfo
	t := &struct {
		*Alias
		Rollup rollupUnmarshaler `json:"rollup"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling RollupPaginatedPropertyInfo: %w", err)
	}
	o.Rollup = t.Rollup.value
	return nil
}

// Title property values
type TitlePropertyItem struct {
	PropertyItemCommon
	Type  alwaysTitle `json:"type"`
	Title RichText    `json:"title"` //  Title property value objects contain an array of rich text objects within the title property.
}

func (_ *TitlePropertyItem) isPropertyItem()                         {}
func (_ *TitlePropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *TitlePropertyItem) UnmarshalJSON(data []byte) error {
	type Alias TitlePropertyItem
	t := &struct {
		*Alias
		Title richTextUnmarshaler `json:"title"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling TitlePropertyItem: %w", err)
	}
	o.Title = t.Title.value
	return nil
}

// Rich Text property values
type RichTextPropertyItem struct {
	PropertyItemCommon
	Type     alwaysRichText `json:"type"`
	RichText RichText       `json:"rich_text"` //  Rich Text property value objects contain an array of rich text objects within the rich_text property.
}

func (_ *RichTextPropertyItem) isPropertyItem()                         {}
func (_ *RichTextPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *RichTextPropertyItem) UnmarshalJSON(data []byte) error {
	type Alias RichTextPropertyItem
	t := &struct {
		*Alias
		RichText richTextUnmarshaler `json:"rich_text"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling RichTextPropertyItem: %w", err)
	}
	o.RichText = t.RichText.value
	return nil
}

// Number property values
type NumberPropertyItem struct {
	PropertyItemCommon
	Type   alwaysNumber `json:"type"`
	Number nullv4.Float `json:"number"` //  Number property value objects contain a number within the number property.
}

func (_ *NumberPropertyItem) isPropertyItem()                         {}
func (_ *NumberPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Select property values
type SelectPropertyItem struct {
	PropertyItemCommon
	Type   alwaysSelect `json:"type"`
	Select Option       `json:"select"`
}

func (_ *SelectPropertyItem) isPropertyItem()                         {}
func (_ *SelectPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// undocumented
type StatusPropertyItem struct {
	PropertyItemCommon
	Type   alwaysStatus `json:"type"`
	Status Option       `json:"status"`
}

func (_ *StatusPropertyItem) isPropertyItem()                         {}
func (_ *StatusPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Multi-select property values
type MultiSelectPropertyItem struct {
	PropertyItemCommon
	Type        alwaysMultiSelect `json:"type"`
	MultiSelect []Option          `json:"multi_select"` //  Multi-select property value objects contain an array of multi-select option values within the multi_select property.
}

func (_ *MultiSelectPropertyItem) isPropertyItem()                         {}
func (_ *MultiSelectPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Date property values
Date property value objects contain the following data within the date property:
*/
type DatePropertyItem struct {
	PropertyItemCommon
	Type alwaysDate            `json:"type"`
	Date *DatePropertyItemData `json:"date"`
}

func (_ *DatePropertyItem) isPropertyItem()                         {}
func (_ *DatePropertyItem) isPropertyItemOrPropertyItemPagination() {}

type DatePropertyItemData struct {
	Start    ISO8601String  `json:"start"`     // An ISO 8601 format date, with optional time.
	End      *ISO8601String `json:"end"`       // An ISO 8601 formatted date, with optional time. Represents the end of a date range.  If null, this property's date value is not a range.
	TimeZone nullv4.String  `json:"time_zone"` // Time zone information for start and end. Possible values are extracted from the IANA database and they are based on the time zones from Moment.js.  When time zone is provided, start and end should not have any UTC offset. In addition, when time zone  is provided, start and end cannot be dates without time information.  If null, time zone information will be contained in UTC offsets in start and end.
}

// Formula property values
type FormulaPropertyItem struct {
	PropertyItemCommon
	Type    alwaysFormula `json:"type"`
	Formula Formula       `json:"formula"` //  Formula property value objects represent the result of evaluating a formula described in the  database's properties. These objects contain a type key and a key corresponding with the value of type. The value is an object containing type-specific data. The type-specific data are described in the sections below.
}

func (_ *FormulaPropertyItem) isPropertyItem()                         {}
func (_ *FormulaPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *FormulaPropertyItem) UnmarshalJSON(data []byte) error {
	type Alias FormulaPropertyItem
	t := &struct {
		*Alias
		Formula formulaUnmarshaler `json:"formula"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling FormulaPropertyItem: %w", err)
	}
	o.Formula = t.Formula.value
	return nil
}

// Relation property values
type RelationPropertyItem struct {
	PropertyItemCommon
	Type     alwaysRelation `json:"type"`
	Relation PageReference  `json:"relation"` //  Relation property value objects contain an array of relation property items with page references within the relation property. A page reference is an object with an id property which is a string value (UUIDv4) corresponding to a page ID in another database.
}

func (_ *RelationPropertyItem) isPropertyItem()                         {}
func (_ *RelationPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// Rollup property values
type RollupPropertyItem struct {
	PropertyItemCommon
	Type   alwaysRollup `json:"type"`
	Rollup Rollup       `json:"rollup"` //  Rollup property value objects represent the result of evaluating a rollup described in the  database's properties. The property is returned as a list object of type property_item with a list of relation items used to computed the rollup under results.   A rollup property item is also returned under the property_type key that describes the rollup aggregation and computed result.   In order to avoid timeouts, if the rollup has a with a large number of aggregations or properties the endpoint returns a next_cursor value that is used to determinate the aggregation value so far for the subset of relations that have been paginated through.   Once has_more is false, then the final rollup value is returned.  See the Pagination documentation for more information on pagination in the Notion API.   Computing the values of following aggregations are not supported. Instead the endpoint returns a list of property_item objects for the rollup: * show_unique (Show unique values) * unique (Count unique values) * median(Median)
}

func (_ *RollupPropertyItem) isPropertyItem()                         {}
func (_ *RollupPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *RollupPropertyItem) UnmarshalJSON(data []byte) error {
	type Alias RollupPropertyItem
	t := &struct {
		*Alias
		Rollup rollupUnmarshaler `json:"rollup"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling RollupPropertyItem: %w", err)
	}
	o.Rollup = t.Rollup.value
	return nil
}

type Rollup interface {
	isRollup()
	GetFunction() string
}
type RollupCommon struct {
	Function string `json:"function"` // Describes the aggregation used.  Possible values include: count,  count_values,  empty,  not_empty,  unique,  show_unique,  percent_empty,  percent_not_empty,  sum,  average,  median,  min,  max,  range,  earliest_date,  latest_date,  date_range,  checked,  unchecked,  percent_checked,  percent_unchecked,  count_per_group,  percent_per_group,  show_original
}

func (c *RollupCommon) GetFunction() string {
	return c.Function
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
	t := struct {
		Number json.RawMessage `json:"number"`
		Date   json.RawMessage `json:"date"`
		Array  json.RawMessage `json:"array"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	switch {
	case t.Number != nil:
		u.value = &NumberRollup{}
	case t.Date != nil:
		u.value = &DateRollup{}
	case t.Array != nil:
		u.value = &ArrayRollup{}
	default:
		return fmt.Errorf("unmarshal Rollup: %s", string(data))
	}
	return json.Unmarshal(data, u.value)
}

func (u *rollupUnmarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

/*
Number rollup property values
Number rollup property values contain a number within the number property.
*/
type NumberRollup struct {
	RollupCommon
	Type   alwaysNumber `json:"type"`
	Number nullv4.Float `json:"number"`
}

func (_ *NumberRollup) isRollup() {}

/*
Date rollup property values
Date rollup property values contain a date property value within the date property.
*/
type DateRollup struct {
	RollupCommon
	Type alwaysDate       `json:"type"`
	Date DatePropertyItem `json:"date"`
}

func (_ *DateRollup) isRollup() {}

/*
Array rollup property values
Array rollup property values contain an array of property_item objects within the results property.
*/
type ArrayRollup struct {
	RollupCommon
	Type  alwaysArray       `json:"type"`
	Array PropertyValueList `json:"array"`
}

func (_ *ArrayRollup) isRollup() {}

/*
People property values
People property value objects contain an array of user objects within the people property.
*/
type PeoplePropertyItem struct {
	PropertyItemCommon
	Type   alwaysPeople `json:"type"`
	People User         `json:"people"`
}

func (_ *PeoplePropertyItem) isPropertyItem()                         {}
func (_ *PeoplePropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *PeoplePropertyItem) UnmarshalJSON(data []byte) error {
	type Alias PeoplePropertyItem
	t := &struct {
		*Alias
		People userUnmarshaler `json:"people"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling PeoplePropertyItem: %w", err)
	}
	o.People = t.People.value
	return nil
}

/*
Files property values
File property value objects contain an array of file references within the files property. A file reference is an object with a File Object and name property, with a string value corresponding to a filename of the original file upload (i.e. "Whole_Earth_Catalog.jpg").
*/
type FilesPropertyItem struct {
	PropertyItemCommon
	Type  alwaysFiles `json:"type"`
	Files FileList    `json:"files"`
}

func (_ *FilesPropertyItem) isPropertyItem()                         {}
func (_ *FilesPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Checkbox property values
Checkbox property value objects contain a boolean within the checkbox property.
*/
type CheckboxPropertyItem struct {
	PropertyItemCommon
	Type     alwaysCheckbox `json:"type"`
	Checkbox bool           `json:"checkbox"`
}

func (_ *CheckboxPropertyItem) isPropertyItem()                         {}
func (_ *CheckboxPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
URL property values
URL property value objects contain a non-empty string within the url property. The string describes a web address (i.e. "http://worrydream.com/EarlyHistoryOfSmalltalk/").
*/
type UrlPropertyItem struct {
	PropertyItemCommon
	Type alwaysUrl     `json:"type"`
	Url  nullv4.String `json:"url"`
}

func (_ *UrlPropertyItem) isPropertyItem()                         {}
func (_ *UrlPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Email property values
Email property value objects contain a string within the email property. The string describes an email address (i.e. "hello@example.org").
*/
type EmailPropertyItem struct {
	PropertyItemCommon
	Type  alwaysEmail   `json:"type"`
	Email nullv4.String `json:"email"`
}

func (_ *EmailPropertyItem) isPropertyItem()                         {}
func (_ *EmailPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Phone number property values
Phone number property value objects contain a string within the phone_number property. No structure is enforced.
*/
type PhoneNumberPropertyItem struct {
	PropertyItemCommon
	Type        alwaysPhoneNumber `json:"type"`
	PhoneNumber nullv4.String     `json:"phone_number"`
}

func (_ *PhoneNumberPropertyItem) isPropertyItem()                         {}
func (_ *PhoneNumberPropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Created time property values
Created time property value objects contain a string within the created_time property. The string contains the date and time when this page was created. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z").
*/
type CreatedTimePropertyItem struct {
	PropertyItemCommon
	Type        alwaysCreatedTime `json:"type"`
	CreatedTime ISO8601String     `json:"created_time"`
}

func (_ *CreatedTimePropertyItem) isPropertyItem()                         {}
func (_ *CreatedTimePropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Created by property values
Created by property value objects contain a user object within the created_by property. The user object describes the user who created this page.
*/
type CreatedByPropertyItem struct {
	PropertyItemCommon
	Type      alwaysCreatedBy `json:"type"`
	CreatedBy User            `json:"created_by"`
}

func (_ *CreatedByPropertyItem) isPropertyItem()                         {}
func (_ *CreatedByPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *CreatedByPropertyItem) UnmarshalJSON(data []byte) error {
	type Alias CreatedByPropertyItem
	t := &struct {
		*Alias
		CreatedBy userUnmarshaler `json:"created_by"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling CreatedByPropertyItem: %w", err)
	}
	o.CreatedBy = t.CreatedBy.value
	return nil
}

/*
Last edited time property values
Last edited time property value objects contain a string within the last_edited_time property. The string contains the date and time when this page was last updated. It is formatted as an ISO 8601 date time string (i.e. "2020-03-17T19:10:04.968Z").
*/
type LastEditedTimePropertyItem struct {
	PropertyItemCommon
	Type           alwaysLastEditedTime `json:"type"`
	LastEditedTime ISO8601String        `json:"last_edited_time"`
}

func (_ *LastEditedTimePropertyItem) isPropertyItem()                         {}
func (_ *LastEditedTimePropertyItem) isPropertyItemOrPropertyItemPagination() {}

/*
Last edited by property values
Last edited by property value objects contain a user object within the last_edited_by property. The user object describes the user who last updated this page.
*/
type LastEditedByPropertyItem struct {
	PropertyItemCommon
	Type         alwaysLastEditedBy `json:"type"`
	LastEditedBy User               `json:"last_edited_by"`
}

func (_ *LastEditedByPropertyItem) isPropertyItem()                         {}
func (_ *LastEditedByPropertyItem) isPropertyItemOrPropertyItemPagination() {}

// UnmarshalJSON assigns the appropriate implementation to interface field(s)
func (o *LastEditedByPropertyItem) UnmarshalJSON(data []byte) error {
	type Alias LastEditedByPropertyItem
	t := &struct {
		*Alias
		LastEditedBy userUnmarshaler `json:"last_edited_by"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshaling LastEditedByPropertyItem: %w", err)
	}
	o.LastEditedBy = t.LastEditedBy.value
	return nil
}
