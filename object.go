package notion

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/psyark/notion/json"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/wI2L/jsondiff"
)

type PropertyMap map[string]Property // for test

type PropertyItemOrPropertyItemPaginationMap map[string]PropertyItemOrPropertyItemPagination

func (m *PropertyItemOrPropertyItemPaginationMap) UnmarshalJSON(data []byte) error {
	t := map[string]propertyItemOrPropertyItemPaginationUnmarshaler{}
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("unmarshaling PropertyValueMap: %w", err)
	}
	*m = PropertyItemOrPropertyItemPaginationMap{}
	for k, u := range t {
		(*m)[k] = u.value
	}
	return nil
}

type ISO8601String = string

type URLReference struct {
	URL string `json:"url"`
}

type PageReference struct {
	Id uuid.UUID `json:"id"`
}

// https://developers.notion.com/reference/errors
type Error struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("(%v) %v", e.Code, e.Message)
}

func String(rta []RichText) string {
	str := ""
	for _, rt := range rta {
		str += rt.PlainText
	}
	return str
}

func defined(fieldValue any) bool {
	v := reflect.ValueOf(fieldValue)

	switch v.Kind() {
	case reflect.Ptr, reflect.Slice:
		return !v.IsNil()
	case reflect.Struct:
		return false
	case reflect.Array: // UUID
		return !v.IsZero()
	case reflect.Bool:
		return !v.IsZero()
	case reflect.String:
		return !v.IsZero()
	default:
		return !v.IsZero()
		// panic(v.Kind())
	}
}

func checkUnmarshal[T any](t *testing.T, wantStr string) {
	var target T
	lo.Must0(json.Unmarshal([]byte(wantStr), &target))
	got := string(lo.Must(json.MarshalIndent(&target, "", "  ")))

	patch := lo.Must(jsondiff.CompareJSON([]byte(wantStr), []byte(got)))
	for _, diff := range patch {
		t.Error(diff)
	}

	assert.JSONEq(t, wantStr, got)
}
