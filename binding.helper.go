package notion

func getTypeForBinding(p Property) string {
	switch p.(type) {
	case *TitleProperty:
		return "RichTextList"
	case *RichTextProperty:
		return "RichTextList"
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
		return "UserList"
	case *FilesProperty:
		return "FileList"
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
