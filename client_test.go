package notion

import (
	"context"
	"encoding/json"
	"fmt"
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
	if p, err := cli.RetrievePage(ctx, "b05213d5c3af4de6924cc9b106ae93ec", checkMarshaller(true)); err != nil {
		t.Fatal(err)
	} else {
		data, _ := json.MarshalIndent(p, "", "  ")
		fmt.Println(string(data))
	}
}
