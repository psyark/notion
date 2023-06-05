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
					if _, err := cli.RetrievePagePropertyItem(ctx, pageId, prop.Id, requestId("RetrievePagePropertyItem_"+testName), useCache(), validateResult()); err != nil {
						t.Fatal(err)
					}
				})
			}
		}
	}
}

func TestUpdatePage(t *testing.T) {
	ctx := context.Background()

	configs := []func(p UpdatePagePropertiesParams){
		func(p UpdatePagePropertiesParams) {},
		func(p UpdatePagePropertiesParams) {
			p.Properties(map[string]PropertyValue{
				"Number":   {Type: "number", Number: nil},
				"Date":     {Type: "date"},
				"Checkbox": {Type: "checkbox", Checkbox: false},
			})
		},
		func(p UpdatePagePropertiesParams) {
			num := rand.Float64() * 1000
			p.Properties(map[string]PropertyValue{
				"Number":   {Type: "number", Number: &num},
				"Date":     {Type: "date", Date: &PropertyValueDate{Start: time.Now().Format(time.RFC3339)}},
				"Checkbox": {Type: "checkbox", Checkbox: true},
			})
		},
	}

	for i, config := range configs {
		config := config
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			params := UpdatePagePropertiesParams{}
			config(params)
			if _, err := cli.UpdatePageProperties(ctx, DATABASE_PAGE_FOR_WRITE, params, requestId(t.Name()), useCache(), validateResult()); err != nil {
				x, _ := json.MarshalIndent(params, "", "  ")
				fmt.Println(string(x))
				t.Fatal(err)
			}
		})
	}
}

func TestQueryDatabase(t *testing.T) {
	filters := []Filter{
		{
			Property: "URL",
			RichText: &FilterRichText{Equals: "http://example.com"},
		},
	}

	ctx := context.Background()
	params := QueryDatabaseParams{}
	for i, filter := range filters {
		filter := filter
		t.Run(fmt.Sprintf("%s_%d", filter.Property, i), func(t *testing.T) {
			params.Filter(filter)
			if pagi, err := cli.QueryDatabase(ctx, DATABASE, params, requestId(t.Name()), useCache(), validateResult()); err != nil {
				t.Fatal(err)
			} else {
				fmt.Println(len(pagi.Results))
			}
		})
	}
}

func TestRetrieveBlockChildren(t *testing.T) {
	t.Skip()

	ctx := context.Background()
	if pagi, err := cli.RetrieveBlockChildren(ctx, STANDALONE_PAGE, requestId(t.Name()), validateResult()); err != nil {
		t.Fatal(err)
	} else {
		blocks, err := pagi.Blocks()
		if err != nil {
			t.Fatal(err)
		}

		for _, block := range blocks {
			block := block
			t.Run(block.Id.String(), func(t *testing.T) {
				if _, err := cli.DeleteBlock(ctx, block.Id, requestId(t.Name()), validateResult()); err != nil {
					t.Fatal(err)
				}
			})
		}
	}
	t.Run("AppendBlockChildren", func(t *testing.T) {
		params := AppendBlockChildrenParams{}
		params.Children([]Block{
			{Type: "breadcrumb"},
			{Heading1: &BlockHeading{RichText: []RichText{{Text: &RichTextText{Content: "Heading 1"}}}}},
			{Heading2: &BlockHeading{RichText: []RichText{{Text: &RichTextText{Content: "Heading 2"}}}}},
			{Heading3: &BlockHeading{RichText: []RichText{{Text: &RichTextText{Content: "Heading 3"}}}}},
			{ToDo: &BlockToDo{RichText: []RichText{{Text: &RichTextText{Content: "To Do"}}}}},
			{Callout: &BlockCallout{
				RichText: []RichText{{Text: &RichTextText{Content: "Callout"}}},
				Icon:     &Emoji{Emoji: "🍣"},
				Children: []Block{
					{Image: &File{External: &FileExternal{Url: "https://placehold.jp/640x640.png"}}},
				},
			}},
			{SyncedBlock: &BlockSyncedBlock{
				Children: []Block{
					{Paragraph: &BlockParagraph{RichText: []RichText{{Text: &RichTextText{Content: "synced"}}}}},
				},
			}},
		})
		if _, err := cli.AppendBlockChildren(ctx, STANDALONE_PAGE, params, requestId(t.Name()), validateResult()); err != nil {
			t.Fatal(err)
		}
	})
}
