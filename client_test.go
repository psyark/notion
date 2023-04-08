package notion

import (
	"context"
	"os"
	"testing"

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
	if _, err := cli.RetrievePage(ctx, "b05213d5c3af4de6924cc9b106ae93ec", validateResult("RetrievePage")); err != nil {
		t.Fatal(err)
	}
}
