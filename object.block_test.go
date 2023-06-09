package notion

import "testing"

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/block

func TestBlock_unmarshal(t *testing.T) {
	tests := []string{
		"{\n\t\"object\": \"block\",\n\t\"id\": \"c02fc1d3-db8b-45c5-a222-27595b15aea7\",\n\t\"parent\": {\n\t\t\"type\": \"page_id\",\n\t\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n\t},\n\t\"created_time\": \"2022-03-01T19:05:00.000Z\",\n\t\"last_edited_time\": \"2022-07-06T19:41:00.000Z\",\n\t\"created_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"last_edited_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"has_children\": false,\n\t\"archived\": false,\n\t\"type\": \"heading_2\",\n\t\"heading_2\": {\n\t\t\"rich_text\": [\n\t\t\t{\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"content\": \"Lacinato kale\",\n\t\t\t\t\t\"link\": null\n\t\t\t\t},\n\t\t\t\t\"annotations\": {\n\t\t\t\t\t\"bold\": false,\n\t\t\t\t\t\"italic\": false,\n\t\t\t\t\t\"strikethrough\": false,\n\t\t\t\t\t\"underline\": false,\n\t\t\t\t\t\"code\": false,\n\t\t\t\t\t\"color\": \"green\"\n\t\t\t\t},\n\t\t\t\t\"plain_text\": \"Lacinato kale\",\n\t\t\t\t\"href\": null\n\t\t\t}\n\t\t],\n\t\t\"color\": \"default\",\n    \"is_toggleable\": false\n\t}\n}\n",
		"{\n  \"object\": \"block\",\n  \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n  \"parent\": {\n    \"type\": \"page_id\",\n    \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n  },\n  \"created_time\": \"2022-12-15T01:14:00.000Z\",\n  \"last_edited_time\": \"2022-12-16T21:23:00.000Z\",\n  \"created_by\": {\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n  },\n  \"last_edited_by\": {\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n  },\n  \"has_children\": false,\n  \"archived\": false,\n  \"type\": \"pdf\",\n  \"pdf\": {\n    \"caption\": [],\n    \"type\": \"external\",\n    \"external\": {\n      \"url\": \"https://www.yourwebsite.dev/files/NewFile.pdf\"\n    }\n  }\n}",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[Block](wantStr); err != nil {
			t.Error(err)
		}
	}
}
