package notion

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/emoji-object

func TestEmoji_unmarshal(t *testing.T) {
	tests := []string{
		"{\n  \"type\": \"emoji\",\n  \"emoji\": \"😻\"\n}",
	}
	for _, wantStr := range tests {
		target := &Emoji{}
		want := []byte(wantStr)
		if err := json.Unmarshal(want, target); err != nil {
			t.Fatal(fmt.Errorf("%w : %s", err, wantStr))
		}
		got, _ := json.Marshal(target)
		if want, got, ok := compareJSON(want, got); !ok {
			t.Fatal(fmt.Errorf("mismatch:\nwant: %s\ngot : %s", want, got))
		}
	}
}
