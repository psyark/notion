package notion

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/parent-object

func TestParent_unmarshal(t *testing.T) {
	target := &parentUnmarshaler{}
	tests := []string{
		"{\n  \"type\": \"database_id\",\n  \"database_id\": \"d9824bdc-8445-4327-be8b-5b47500af6ce\"\n}",
		"{\n  \"type\": \"page_id\",\n\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n}",
		"{\n\t\"type\": \"workspace\",\n\t\"workspace\": true\n}",
		"{\n\t\"type\": \"block_id\",\n\t\"block_id\": \"7d50a184-5bbe-4d90-8f29-6bec57ed817b\"\n}",
	}
	for _, wantStr := range tests {
		want := []byte(wantStr)
		if err := json.Unmarshal(want, target); err != nil {
			t.Fatal(err)
		}
		got, _ := json.Marshal(target.value)
		if want, got, ok := compareJSON(want, got); !ok {
			t.Fatal(fmt.Errorf("mismatch:\nwant: %s\ngot : %s", want, got))
		}
	}
}
