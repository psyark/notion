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

func init() {
	if err := godotenv.Load("env.local"); err != nil {
		panic(err)
	}
	cli = NewClient(os.Getenv("API_ACCESS_TOKEN"))
}

func TestRetrievePage(t *testing.T) {
	ctx := context.Background()
	if _, err := cli.RetrievePage(ctx, uuid.MustParse("b05213d5c3af4de6924cc9b106ae93ec"), validateResult("RetrievePage")); err != nil {
		t.Fatal(err)
	}
}

func TestRetrievePagePropertyItem_a(t *testing.T) {
	ctx := context.Background()
	if _, err := cli.RetrievePagePropertyItem(ctx, uuid.MustParse("7e01d5af9d0e4d2584e4d5bfc39b65bf"), "title", validateResult("TestRetrievePagePropertyItem_a")); err != nil {
		t.Fatal(err)
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
			if _, err := cli.UpdatePage(ctx, uuid.MustParse("b8ff7c186ef2416cb9654daf0d7aa961"), params, validateResult("TestUpdatePage")); err != nil {
				x, _ := json.MarshalIndent(params, "", "  ")
				fmt.Println(string(x))
				t.Fatal(err)
			}
		})
	}
}
