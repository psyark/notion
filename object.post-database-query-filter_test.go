// Code generated by notion.doc2api; DO NOT EDIT.

package notion

import "testing"

// https://developers.notion.com/reference/post-database-query-filter

func TestFilter_unmarshal(t *testing.T) {
	tests := []string{
		"{\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"equals\": true\n    }\n  }",
		"{\n    \"property\": \"Task completed\",\n    \"checkbox\": {\n      \"does_not_equal\": true\n    }\n  }",
		"{\n    \"property\": \"Due date\",\n    \"date\": {\n      \"on_or_after\": \"2023-02-08\"\n    }\n  }",
		"{\n    \"property\": \"Blueprint\",\n    \"files\": {\n      \"is_not_empty\": true\n    }\n  }",
		"{\n    \"property\": \"One month deadline\",\n    \"formula\": {\n      \"date\":{\n          \"after\": \"2021-05-10\"\n      }\n    }\n  }",
		"{\n    \"property\": \"Programming language\",\n    \"multi_select\": {\n      \"contains\": \"TypeScript\"\n    }\n  }",
		"{\n    \"property\": \"Estimated working days\",\n    \"number\": {\n      \"less_than_or_equal_to\": 5\n    }\n  }",
		"{\n    \"property\": \"Last edited by\",\n    \"people\": {\n      \"contains\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n    }\n  }",
		"{\n    \"property\": \"✔️ Task List\",\n    \"relation\": {\n      \"contains\": \"0c1f7cb2-8090-4f18-924e-d92965055e32\"\n    }\n  }",
		"{\n    \"property\": \"Description\",\n    \"rich_text\": {\n      \"contains\": \"cross-team\"\n    }\n  }",
		"{\n    \"property\": \"Related tasks\",\n    \"rollup\": {\n      \"any\": {\n        \"rich_text\": {\n          \"contains\": \"Migrate database\"\n        }\n      }\n    }\n  }",
		"{\n    \"property\": \"Parent project due date\",\n    \"rollup\": {\n      \"date\": {\n        \"on_or_before\": \"2023-02-08\"\n      }\n    }\n  }",
		"{\n    \"property\": \"Total estimated working days\",\n    \"rollup\": {\n      \"number\": {\n        \"does_not_equal\": 42\n      }\n    }\n  }",
		"{\n    \"property\": \"Frontend framework\",\n    \"select\": {\n      \"equals\": \"React\"\n    }\n  }",
		"{\n    \"property\": \"Project status\",\n    \"status\": {\n      \"equals\": \"Not started\"\n    }\n  }",
		"{\n    \"timestamp\": \"created_time\",\n    \"created_time\": {\n      \"on_or_before\": \"2022-10-13\"\n    }\n  }",
		"{\n    \"and\": [\n      {\n        \"property\": \"Complete\",\n        \"checkbox\": {\n          \"equals\": true\n        }\n      },\n      {\n        \"property\": \"Working days\",\n        \"number\": {\n          \"greater_than\": 10\n        }\n      }\n    ]\n  }",
		"{\n    \"or\": [\n      {\n        \"property\": \"Description\",\n        \"rich_text\": {\n          \"contains\": \"2023\"\n        }\n      },\n      {\n        \"and\": [\n          {\n            \"property\": \"Department\",\n            \"select\": {\n              \"equals\": \"Engineering\"\n            }\n          },\n          {\n            \"property\": \"Priority goal\",\n            \"checkbox\": {\n              \"equals\": true\n            }\n          }\n        ]\n      }\n    ]\n  }",
	}
	for _, wantStr := range tests {
		if err := checkUnmarshal[Filter](wantStr); err != nil {
			t.Error(err)
		}
	}
}
