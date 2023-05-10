package notion

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Code generated by notion.doc2api; DO NOT EDIT.
// https://developers.notion.com/reference/property-object

func TestPropertyMap_unmarshal(t *testing.T) {
	tests := []string{
		"{\"Task complete\": {\n  \"id\": \"BBla\",\n  \"name\": \"Task complete\",\n  \"type\": \"checkbox\",\n  \"checkbox\": {}\n}\n}",
		"{\"Created by\": {\n  \"id\": \"%5BJCR\",\n  \"name\": \"Created by\",\n  \"type\": \"created_by\",\n  \"created_by\": {}\n}\n}",
		"{\"Created time\": {\n  \"id\": \"XcAf\",\n  \"name\": \"Created time\",\n  \"type\": \"created_time\",\n  \"created_time\": {}\n}\n}",
		"{\"Task due date\": {\n  \"id\": \"AJP%7D\",\n  \"name\": \"Task due date\",\n  \"type\": \"date\",\n  \"date\": {}\n}\n}",
		"{\"Contact email\": {\n  \"id\": \"oZbC\",\n  \"name\": \"Contact email\",\n  \"type\": \"email\",\n  \"email\": {}\n}\n}",
		"{\"Product image\": {\n  \"id\": \"pb%3E%5B\",\n  \"name\": \"Product image\",\n  \"type\": \"files\",\n  \"files\": {}\n}\n}",
		"{\"Updated price\": {\n  \"id\": \"YU%7C%40\",\n  \"name\": \"Updated price\",\n  \"type\": \"formula\",\n  \"formula\": {\n    \"expression\": \"prop(\\\"Price\\\") * 2\"\n  }\n}\n}",
		"{\"Last edited time\": {\n  \"id\": \"jGdo\",\n  \"name\": \"Last edited time\",\n  \"type\": \"last_edited_time\",\n  \"last_edited_time\": {}\n}\n}",
		"{\"Store availability\": {\n  \"id\": \"flsb\",\n  \"name\": \"Store availability\",\n  \"type\": \"multi_select\",\n  \"multi_select\": {\n    \"options\": [\n      {\n        \"id\": \"5de29601-9c24-4b04-8629-0bca891c5120\",\n        \"name\": \"Duc Loi Market\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"385890b8-fe15-421b-b214-b02959b0f8d9\",\n        \"name\": \"Rainbow Grocery\",\n        \"color\": \"gray\"\n      },\n      {\n        \"id\": \"72ac0a6c-9e00-4e8c-80c5-720e4373e0b9\",\n        \"name\": \"Nijiya Market\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"9556a8f7-f4b0-4e11-b277-f0af1f8c9490\",\n        \"name\": \"Gus's Community Market\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}\n}",
		"{\"Price\":{\n  \"id\": \"%7B%5D_P\",\n  \"name\": \"Price\",\n  \"type\": \"number\",\n  \"number\": {\n    \"format\": \"dollar\"\n  }\n}\n}",
		"{\"Project owner\": {\n  \"id\": \"FlgQ\",\n  \"name\": \"Project owner\",\n  \"type\": \"people\",\n  \"people\": {}\n}\n}",
		"{\"Contact phone number\": {\n  \"id\": \"ULHa\",\n  \"name\": \"Contact phone number\",\n  \"type\": \"phone_number\",\n  \"phone_number\": {}\n}\n}",
		"{\"Project description\": {\n  \"id\": \"NZZ%3B\",\n  \"name\": \"Project description\",\n  \"type\": \"rich_text\",\n  \"rich_text\": {}\n}\n}",
		"{\"Estimated total project time\": {\n  \"id\": \"%5E%7Cy%3C\",\n  \"name\": \"Estimated total project time\",\n  \"type\": \"rollup\",\n  \"rollup\": {\n    \"rollup_property_name\": \"Days to complete\",\n    \"relation_property_name\": \"Tasks\",\n    \"rollup_property_id\": \"\\\\nyY\",\n    \"relation_property_id\": \"Y]<y\",\n    \"function\": \"sum\"\n  }\n}\n}",
		"{\"Food group\": {\n  \"id\": \"%40Q%5BM\",\n  \"name\": \"Food group\",\n  \"type\": \"select\",\n  \"select\": {\n    \"options\": [\n      {\n        \"id\": \"e28f74fc-83a7-4469-8435-27eb18f9f9de\",\n        \"name\": \"🥦Vegetable\",\n        \"color\": \"purple\"\n      },\n      {\n        \"id\": \"6132d771-b283-4cd9-ba44-b1ed30477c7f\",\n        \"name\": \"🍎Fruit\",\n        \"color\": \"red\"\n      },\n      {\n        \"id\": \"fc9ea861-820b-4f2b-bc32-44ed9eca873c\",\n        \"name\": \"💪Protein\",\n        \"color\": \"yellow\"\n      }\n    ]\n  }\n}\n}",
		"{\"Status\": {\n  \"id\": \"biOx\",\n  \"name\": \"Status\",\n  \"type\": \"status\",\n  \"status\": {\n    \"options\": [\n      {\n        \"id\": \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\",\n        \"name\": \"Not started\",\n        \"color\": \"default\"\n      },\n      {\n        \"id\": \"330aeafb-598c-4e1c-bc13-1148aa5963d3\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\"\n      },\n      {\n        \"id\": \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\",\n        \"name\": \"Done\",\n        \"color\": \"green\"\n      }\n    ],\n    \"groups\": [\n      {\n        \"id\": \"b9d42483-e576-4858-a26f-ed940a5f678f\",\n        \"name\": \"To-do\",\n        \"color\": \"gray\",\n        \"option_ids\": [\n          \"034ece9a-384d-4d1f-97f7-7f685b29ae9b\"\n        ]\n      },\n      {\n        \"id\": \"cf4952eb-1265-46ec-86ab-4bded4fa2e3b\",\n        \"name\": \"In progress\",\n        \"color\": \"blue\",\n        \"option_ids\": [\n          \"330aeafb-598c-4e1c-bc13-1148aa5963d3\"\n        ]\n      },\n      {\n        \"id\": \"4fa7348e-ae74-46d9-9585-e773caca6f40\",\n        \"name\": \"Complete\",\n        \"color\": \"green\",\n        \"option_ids\": [\n          \"497e64fb-01e2-41ef-ae2d-8a87a3bb51da\"\n        ]\n      }\n    ]\n  }\n}\n}",
		"{\"Project name\": {\n  \"id\": \"title\",\n  \"name\": \"Project name\",\n  \"type\": \"title\",\n  \"title\": {}\n}\n}",
		"{\"Project URL\": {\n  \"id\": \"BZKU\",\n  \"name\": \"Project URL\",\n  \"type\": \"url\",\n  \"url\": {}\n}\n}",
	}
	for _, wantStr := range tests {
		target := &PropertyMap{}
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
