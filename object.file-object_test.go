// Code generated by notion.doc2api; DO NOT EDIT.

package notion

import "testing"

// https://developers.notion.com/reference/file-object

func TestFile_unmarshal(t *testing.T) {
	tests := []string{
		"{\n  \"type\": \"file\",\n  \"file\": {\n    \"url\": \"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/7b8b0713-dbd4-4962-b38b-955b6c49a573/My_test_image.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20221024%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20221024T205211Z&X-Amz-Expires=3600&X-Amz-Signature=208aa971577ff05e75e68354e8a9488697288ff3fb3879c2d599433a7625bf90&X-Amz-SignedHeaders=host&x-id=GetObject\",\n    \"expiry_time\": \"2022-10-24T22:49:22.765Z\"\n  }\n}\n",
		"{\n  \"type\": \"external\",\n  \"external\": {\n    \"url\": \"https://images.unsplash.com/photo-1525310072745-f49212b5ac6d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1065&q=80\"\n  }\n}\n",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[File](wantStr); err != nil {
			t.Error(err)
		}
	}
}

func TestPagination_unmarshal(t *testing.T) {
	tests := []string{
		"{\n    \"object\": \"list\",\n    \"results\": [\n        {\n            \"object\": \"block\",\n            \"id\": \"47a920e4-346c-4df8-ae78-905ce10adcb8\",\n            \"parent\": {\n                \"type\": \"page_id\",\n                \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n            },\n            \"created_time\": \"2022-12-15T00:18:00.000Z\",\n            \"last_edited_time\": \"2022-12-15T00:18:00.000Z\",\n            \"created_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"last_edited_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"has_children\": false,\n          \t\"archived\": false,\n            \"in_trash\": false,\n            \"type\": \"paragraph\",\n            \"paragraph\": {\n                \"rich_text\": [],\n                \"color\": \"default\"\n            }\n        },\n        {\n            \"object\": \"block\",\n            \"id\": \"3c29dedf-00a5-4915-b137-120c61f5e5d8\",\n            \"parent\": {\n                \"type\": \"page_id\",\n                \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n            },\n            \"created_time\": \"2022-12-15T00:18:00.000Z\",\n            \"last_edited_time\": \"2022-12-15T00:18:00.000Z\",\n            \"created_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"last_edited_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"has_children\": false,\n          \t\"archived\": false,\n            \"in_trash\": false,\n            \"type\": \"file\",\n            \"file\": {\n                \"caption\": [],\n                \"type\": \"file\",\n                \"file\": {\n                    \"url\": \"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/fa6c03f0-e608-45d0-9327-4cd7a5e56e71/TestFile.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20221215%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20221215T002012Z&X-Amz-Expires=3600&X-Amz-Signature=bf13ca59f618077852298cb92aedc4dd1becdc961c31d73cbc030ef93f2853c4&X-Amz-SignedHeaders=host&x-id=GetObject\",\n                    \"expiry_time\": \"2022-12-15T01:20:12.928Z\"\n                }\n            }\n        } \n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"block\",\n    \"block\": {}\n}\n",
		"{\n  \"object\": \"list\",\n  \"results\": [\n    {\n      \"object\": \"block\",\n      \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n      \"parent\": {\n        \"type\": \"page_id\",\n        \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n      },\n      \"created_time\": \"2022-12-15T01:14:00.000Z\",\n      \"last_edited_time\": \"2022-12-15T01:14:00.000Z\",\n      \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"has_children\": false,\n      \"archived\": false,\n      \"in_trash\": false,\n      \"type\": \"pdf\",\n      \"pdf\": {\n        \"caption\": [],\n        \"type\": \"external\",\n        \"external\": {\n          \"url\": \"https://www.yourwebsite.dev/files/TestFile.pdf\"\n        }\n      }\n    }\n  ],\n  \"next_cursor\": null,\n  \"has_more\": false,\n  \"type\": \"block\",\n  \"block\": {}\n}\n",
		"{\n  \"object\": \"list\",\n  \"results\": [\n    {\n      \"object\": \"block\",\n      \"id\": \"47a920e4-346c-4df8-ae78-905ce10adcb8\",\n      \"parent\": {\n        \"type\": \"page_id\",\n        \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n      },\n      \"created_time\": \"2022-12-15T00:18:00.000Z\",\n      \"last_edited_time\": \"2022-12-15T00:18:00.000Z\",\n      \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n      },\n      \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n      },\n      \"has_children\": false,\n      \"archived\": false,\n      \"in_trash\": false,\n      \"type\": \"paragraph\",\n      \"paragraph\": {\n        \"rich_text\": [],\n        \"color\": \"default\"\n      }\n    },\n    {\n      \"object\": \"block\",\n      \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n      \"parent\": {\n        \"type\": \"page_id\",\n        \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n      },\n      \"created_time\": \"2022-12-15T01:14:00.000Z\",\n      \"last_edited_time\": \"2022-12-15T01:14:00.000Z\",\n      \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"has_children\": false,\n      \"archived\": false,\n      \"in_trash\": false,\n      \"type\": \"pdf\",\n      \"pdf\": {\n        \"caption\": [],\n        \"type\": \"external\",\n        \"external\": {\n          \"url\": \"https://www.yourwebsite.dev/files/TestFile.pdf\"\n        }\n      }\n    }\n  ],\n  \"next_cursor\": null,\n  \"has_more\": false,\n  \"type\": \"block\",\n  \"block\": {}\n}\n",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[Pagination](wantStr); err != nil {
			t.Error(err)
		}
	}
}

func TestBlock_unmarshal_(t *testing.T) {
	tests := []string{
		"{\n  \"object\": \"block\",\n  \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n  \"parent\": {\n    \"type\": \"page_id\",\n    \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n  },\n  \"created_time\": \"2022-12-15T01:14:00.000Z\",\n  \"last_edited_time\": \"2022-12-16T21:23:00.000Z\",\n  \"created_by\": {\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n  },\n  \"last_edited_by\": {\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n  },\n  \"has_children\": false,\n  \"archived\": false,\n  \"in_trash\": false,\n  \"type\": \"pdf\",\n  \"pdf\": {\n    \"caption\": [],\n    \"type\": \"external\",\n    \"external\": {\n      \"url\": \"https://www.yourwebsite.dev/files/NewFile.pdf\"\n    }\n  }\n}\n",
		"{\n\t\"object\": \"block\",\n\t\"id\": \"c02fc1d3-db8b-45c5-a222-27595b15aea7\",\n\t\"parent\": {\n\t\t\"type\": \"page_id\",\n\t\t\"page_id\": \"59833787-2cf9-4fdf-8782-e53db20768a5\"\n\t},\n\t\"created_time\": \"2022-03-01T19:05:00.000Z\",\n\t\"last_edited_time\": \"2022-07-06T19:41:00.000Z\",\n\t\"created_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"last_edited_by\": {\n\t\t\"object\": \"user\",\n\t\t\"id\": \"ee5f0f84-409a-440f-983a-a5315961c6e4\"\n\t},\n\t\"has_children\": false,\n\t\"archived\": false,\n  \"in_trash\": false,\n\t\"type\": \"heading_2\",\n\t\"heading_2\": {\n\t\t\"rich_text\": [\n\t\t\t{\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"content\": \"Lacinato kale\",\n\t\t\t\t\t\"link\": null\n\t\t\t\t},\n\t\t\t\t\"annotations\": {\n\t\t\t\t\t\"bold\": false,\n\t\t\t\t\t\"italic\": false,\n\t\t\t\t\t\"strikethrough\": false,\n\t\t\t\t\t\"underline\": false,\n\t\t\t\t\t\"code\": false,\n\t\t\t\t\t\"color\": \"green\"\n\t\t\t\t},\n\t\t\t\t\"plain_text\": \"Lacinato kale\",\n\t\t\t\t\"href\": null\n\t\t\t}\n\t\t],\n\t\t\"color\": \"default\",\n    \"is_toggleable\": false\n\t}\n}\n",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[Block](wantStr); err != nil {
			t.Error(err)
		}
	}
}
