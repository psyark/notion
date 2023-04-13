package notion

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
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
	if v, err := cli.RetrievePagePropertyItem(ctx, uuid.MustParse("7e01d5af9d0e4d2584e4d5bfc39b65bf"), "title", validateResult("RetrievePage")); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(v)
	}
}
