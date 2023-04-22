package notion

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/user

func TestUser_unmarshal(t *testing.T) {
	target := &userUnmarshaler{}
	tests := []string{
		"{\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\",\n    \"name\": \"Test Integration\",\n    \"avatar_url\": null,\n    \"type\": \"bot\",\n    \"bot\": {\n        \"owner\": {\n        \"type\": \"workspace\",\n        \"workspace\": true\n        },\n \"workspace_name\": \"Ada Lovelace’s Notion\"\n    }\n}",
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

func TestBotDataOwner_unmarshal(t *testing.T) {
	target := &BotDataOwner{}
	tests := []string{
		"{\n    \"type\": \"workspace\",\n    \"workspace\": true\n}",
	}
	for _, wantStr := range tests {
		want := []byte(wantStr)
		if err := json.Unmarshal(want, target); err != nil {
			t.Fatal(err)
		}
		got, _ := json.Marshal(target)
		if want, got, ok := compareJSON(want, got); !ok {
			t.Fatal(fmt.Errorf("mismatch:\nwant: %s\ngot : %s", want, got))
		}
	}
}