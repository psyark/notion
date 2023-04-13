package notion

import (
	"context"
	"math/rand"
	"os"
	"testing"

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
				"Number": &NumberPropertyValue{Type: "number", Number: null.FloatFromPtr(nil)},
			},
		},
		{
			Properties: PropertyValueMap{
				"Number": &NumberPropertyValue{Type: "number", Number: null.FloatFrom(rand.Float64() * 1000)},
			},
		},
	}

	for _, params := range cases {
		if _, err := cli.UpdatePage(ctx, uuid.MustParse("b8ff7c186ef2416cb9654daf0d7aa961"), params, validateResult("TestUpdatePage")); err != nil {
			t.Fatal(err)
		}
	}
}
