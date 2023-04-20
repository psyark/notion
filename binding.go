package notion

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"

	"github.com/dave/jennifer/jen"
)

var safeName func(string) string

func init() {
	nonAscii := regexp.MustCompile(`_|\W`)
	safeName = func(s string) string {
		return nonAscii.ReplaceAllStringFunc(s, func(s string) string {
			if s == "_" {
				return "__"
			}
			return fmt.Sprintf("_%d", s[0])
		})
	}
}

// UnmarshalPage は渡されたpageのプロパティ・カバー・アイコンをdstに格納します
// dstは適切にタグ付けされたstructへのポインタである必要があります
func UnmarshalPage(page *Page, dst any) error {
	t := reflect.TypeOf(dst)
	v := reflect.ValueOf(dst)
	if t.Kind() != reflect.Pointer {
		return fmt.Errorf("dst must be a pointer to a tagged struct")
	}

	t = t.Elem()
	v = v.Elem()
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("dst must be a pointer to a tagged struct")
	}

	setProperty := func(fv reflect.Value, prop PropertyValue) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
		}()
		fv.Set(prop.get())
		return err
	}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.Tag.Get("notion") != "" {
			propId := sf.Tag.Get("notion") // TODO カバーやアイコンの考慮
			prop := page.Properties.Get(propId)
			if prop == nil {
				return fmt.Errorf("タグ %q に相当するプロパティがありません", propId)
			}

			if err := setProperty(v.Field(i), prop); err != nil {
				return fmt.Errorf("notion.UnmarshalPage(%v): %w", sf.Name, err)
			}
		}
	}

	return nil
}

// GetUpdatePageParams は渡されたsrcと現在のpageを比較し、
// プロパティ・カバー・アイコンを更新するためのUpdatePageParams、または更新が不要な場合のnilを返します
// srcは適切にタグ付けされたstruct（またはそのポインタ）である必要があります
func GetUpdatePageParams(src any, page *Page) (*UpdatePageParams, error) {
	t := reflect.TypeOf(src)
	v := reflect.ValueOf(src)
	if t.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("src must be a pointer to a tagged struct")
	}

	t = t.Elem()
	v = v.Elem()
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("src must be a pointer to a tagged struct")
	}

	var params *UpdatePageParams
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.Tag.Get("notion") != "" {
			propId := sf.Tag.Get("notion") // TODO カバーやアイコンの考慮
			prop := page.Properties.Get(propId)
			if prop == nil {
				return nil, fmt.Errorf("タグ %q に相当するプロパティがありません", propId)
			}

			json1, _ := json.Marshal(prop.get().Interface())
			json2, _ := json.Marshal(v.Field(i).Interface())
			if string(json1) != string(json2) {
				if params == nil {
					params = &UpdatePageParams{Properties: PropertyValueMap{}}
				}
				pvu := propertyValueUnmarshaler{}
				data, _ := json.Marshal(prop)
				_ = pvu.UnmarshalJSON(data)
				pvu.value.set(v.Field(i))

				params.Properties[propId] = pvu.value
			}
		}
	}

	return params, nil
}

// GetCreatePageParams は渡されたsrcからdbの定義に従って
// プロパティ・カバー・アイコンが設定されたCreatePageParamsを返します
// srcは適切にタグ付けされたstruct（またはそのポインタ）である必要があります
func GetCreatePageParams(src any, db *Database) *CreatePageParams {
	return nil
}

func ToTaggedStruct(db *Database) string {
	fields := []jen.Code{}
	for _, prop := range db.Properties {
		name := safeName(prop.GetName())
		// data, _ := json.Marshal(prop)
		fields = append(fields, jen.Id(name).Op(getTypeForBinding(prop)).Tag(map[string]string{"notion": prop.GetId()}))
	}

	code := jen.Type().Id(safeName(db.Title.String())).Struct(fields...)
	return (&jen.Statement{code}).GoString()
}
