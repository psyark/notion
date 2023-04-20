package notion

import "reflect"

func (v TitlePropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Title)
}
func (v RichTextPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.RichText)
}
func (v NumberPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Number)
}
func (v SelectPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Select)
}
func (v StatusPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Status)
}
func (v MultiSelectPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.MultiSelect)
}
func (v DatePropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Date)
}
func (v FormulaPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Formula)
}
func (v RelationPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Relation)
}
func (v RollupPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Rollup)
}
func (v PeoplePropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.People)
}
func (v FilesPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Files)
}
func (v CheckboxPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Checkbox)
}
func (v UrlPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Url)
}
func (v EmailPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.Email)
}
func (v PhoneNumberPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.PhoneNumber)
}
func (v CreatedTimePropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.CreatedTime)
}
func (v CreatedByPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.CreatedBy)
}
func (v LastEditedTimePropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.LastEditedTime)
}
func (v LastEditedByPropertyValue) value() reflect.Value {
	return reflect.ValueOf(v.LastEditedBy)
}
func getTypeForBinding(p Property) string {
	switch p.(type) {
	case *TitleProperty:
		return "RichTextArray"
	case *RichTextProperty:
		return "RichTextArray"
	case *NumberProperty:
		return "nullv4.Float"
	case *SelectProperty:
		return "*Option"
	case *StatusProperty:
		return "Option"
	case *MultiSelectProperty:
		return "[]Option"
	case *DateProperty:
		return "*DatePropertyValueData"
	case *FormulaProperty:
		return "Formula"
	case *RelationProperty:
		return "[]PageReference"
	case *RollupProperty:
		return "Rollup"
	case *PeopleProperty:
		return "Users"
	case *FilesProperty:
		return "Files"
	case *CheckboxProperty:
		return "bool"
	case *UrlProperty:
		return "nullv4.String"
	case *EmailProperty:
		return "nullv4.String"
	case *PhoneNumberProperty:
		return "nullv4.String"
	case *CreatedTimeProperty:
		return "ISO8601String"
	case *CreatedByProperty:
		return "PartialUser"
	case *LastEditedTimeProperty:
		return "ISO8601String"
	case *LastEditedByProperty:
		return "User"
	}
	panic(p)
}
