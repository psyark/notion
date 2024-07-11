package notion

import "testing"

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/user

func TestUser_unmarshal(t *testing.T) {
	tests := []string{
		"{     \"object\": \"user\",     \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\",     \"name\": \"Test Integration\",     \"avatar_url\": null,     \"type\": \"bot\",     \"bot\": {         \"owner\": {         \"type\": \"workspace\",         \"workspace\": true         },  \"workspace_name\": \"Ada Lovelace’s Notion\"     } }",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[User](wantStr); err != nil {
			t.Error(err)
		}
	}
}

func TestBotUserDataOwner_unmarshal(t *testing.T) {
	tests := []string{
		"{     \"type\": \"workspace\",     \"workspace\": true }",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[BotUserDataOwner](wantStr); err != nil {
			t.Error(err)
		}
	}
}
