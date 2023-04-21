package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gopkg.in/guregu/null.v4"
)

var cli *Client

var (
	STANDALONE_PAGE         = uuid.MustParse("b05213d5c3af4de6924cc9b106ae93ec") // https://www.notion.so/Page-b05213d5c3af4de6924cc9b106ae93ec
	DATABASE                = uuid.MustParse("edd0404128004a83bd29deb729221ec7") // https://www.notion.so/edd0404128004a83bd29deb729221ec7
	DATABASE_PAGE_FOR_READ1 = uuid.MustParse("7e01d5af9d0e4d2584e4d5bfc39b65bf") // https://www.notion.so/ABCDEFG-7e01d5af9d0e4d2584e4d5bfc39b65bf
	DATABASE_PAGE_FOR_READ2 = uuid.MustParse("7e1105bc19a64a1381453cff0b488092") // https://www.notion.so/7e1105bc19a64a1381453cff0b488092
	DATABASE_PAGE_FOR_WRITE = uuid.MustParse("b8ff7c186ef2416cb9654daf0d7aa961") // https://www.notion.so/PageToUpdate-b8ff7c186ef2416cb9654daf0d7aa961
)

func init() {
	if err := godotenv.Load("env.local"); err != nil {
		panic(err)
	}
	cli = NewClient(os.Getenv("API_ACCESS_TOKEN"))
}

func TestRetrievePage(t *testing.T) {
	ctx := context.Background()
	if _, err := cli.RetrievePage(ctx, STANDALONE_PAGE, requestId("RetrievePage"), useCache(), validateResult()); err != nil {
		t.Fatal(err)
	}
}

func TestRetrievePagePropertyItem(t *testing.T) {
	ctx := context.Background()
	if db, err := cli.RetrieveDatabase(ctx, DATABASE, requestId("RetrieveDatabase"), useCache(), validateResult()); err != nil {
		t.Fatal(err)
	} else {
		for name, prop := range db.Properties {
			name, prop := name, prop
			for i, pageId := range []uuid.UUID{DATABASE_PAGE_FOR_READ1, DATABASE_PAGE_FOR_READ2} {
				testName := fmt.Sprintf("%s_%d", name, i+1)
				t.Run(testName, func(t *testing.T) {
					if _, err := cli.RetrievePagePropertyItem(ctx, pageId, prop.GetId(), requestId("RetrievePagePropertyItem_"+testName), useCache(), validateResult()); err != nil {
						t.Fatal(err)
					}
				})
			}
		}
	}
}

func TestUpdatePage(t *testing.T) {
	ctx := context.Background()

	cases := []*UpdatePageParams{
		{},
		{
			Properties: PropertyValueMap{
				"Number":   &NumberPropertyValue{Number: null.FloatFromPtr(nil)},
				"Date":     &DatePropertyValue{},
				"Checkbox": &CheckboxPropertyValue{Checkbox: false},
			},
		},
		{
			Properties: PropertyValueMap{
				"Number":   &NumberPropertyValue{Number: null.FloatFrom(rand.Float64() * 1000)},
				"Date":     &DatePropertyValue{Date: &DatePropertyValueData{Start: time.Now().Format(time.RFC3339)}},
				"Checkbox": &CheckboxPropertyValue{Checkbox: true},
			},
		},
	}

	for i, params := range cases {
		params := params
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			if _, err := cli.UpdatePage(ctx, DATABASE_PAGE_FOR_WRITE, params, requestId("UpdatePage"), useCache(), validateResult()); err != nil {
				x, _ := json.MarshalIndent(params, "", "  ")
				fmt.Println(string(x))
				t.Fatal(err)
			}
		})
	}
}

func TestQueryDatabase(t *testing.T) {
	filters := []Filter{
		&RichTextFilter{
			filterCommon: filterCommon{Property: "URL"},
			RichText:     RichTextFilterData{Equals: "http://example.com"},
		},
	}

	ctx := context.Background()
	params := &QueryDatabaseParams{
		Filter: RichTextFilter{
			filterCommon: filterCommon{
				Property: "URL",
			},
			RichText: RichTextFilterData{
				Equals: "http://example.com",
			},
		},
	}
	for _, filter := range filters {
		filter := filter
		t.Run(fmt.Sprintf("%T", filter), func(t *testing.T) {
			params.Filter = filter
			if pagi, err := cli.QueryDatabase(ctx, DATABASE, params, requestId(t.Name()), useCache(), validateResult()); err != nil {
				t.Fatal(err)
			} else {
				fmt.Println(len(pagi.Results))
			}
		})
	}
}
