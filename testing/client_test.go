package testing

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
	. "github.com/psyark/notion"
	"github.com/samber/lo"
)

// TODO ボタンなど新しい要素のテスト

var client *Client

var (
	ROOT                    = uuid.MustParse("9c20de5e26af4959a26e390b537af4c8") // https://www.notion.so/Root-9c20de5e26af4959a26e390b537af4c8
	STANDALONE_PAGE         = uuid.MustParse("b05213d5c3af4de6924cc9b106ae93ec") // https://www.notion.so/Page-b05213d5c3af4de6924cc9b106ae93ec
	DATABASE                = uuid.MustParse("edd0404128004a83bd29deb729221ec7") // https://www.notion.so/edd0404128004a83bd29deb729221ec7
	DATABASE_PAGE_FOR_READ1 = uuid.MustParse("7e01d5af9d0e4d2584e4d5bfc39b65bf") // https://www.notion.so/ABCDEFG-7e01d5af9d0e4d2584e4d5bfc39b65bf
	DATABASE_PAGE_FOR_READ2 = uuid.MustParse("7e1105bc19a64a1381453cff0b488092") // https://www.notion.so/7e1105bc19a64a1381453cff0b488092
	DATABASE_PAGE_FOR_WRITE = uuid.MustParse("b8ff7c186ef2416cb9654daf0d7aa961") // https://www.notion.so/PageToUpdate-b8ff7c186ef2416cb9654daf0d7aa961
	currentPage             uuid.UUID
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	lo.Must0(godotenv.Load("env.local"))
	client = NewClient(os.Getenv("API_ACCESS_TOKEN"))

	params := SearchByTitleParams{}
	pagi := lo.Must(client.SearchByTitle(ctx, params))

	for _, pd := range pagi.Results {
		switch pd := pd.(type) {
		case *Page:
			if pd.Parent.PageId == ROOT {

			}
		case *Database:
			if pd.Parent.PageId == ROOT {
				fmt.Println(String(pd.Title))
			}
		}
	}

	m.Run()
}

func TestCreateDatabase(t *testing.T) {
	ctx := context.Background()

	params := CreateDatabaseParams{}
	params.Parent(Parent{PageId: ROOT})
	params.Title([]RichText{{Text: &RichTextText{Content: "テストデータベース"}}})
	params.Properties(map[string]PropertySchema{
		"タイトル": {Title: &struct{}{}},
		"テキスト": {RichText: &struct{}{}},
		"数値":   {Number: &PropertySchemaNumber{Format: "number_with_commas"}},
		"セレクト": {Select: &PropertySchemaSelect{Options: []PropertySchemaOption{{Name: "赤", Color: "red"}}}},
	})

	lo.Must(client.CreateDatabase(ctx, params, WithRoundTripper(useCache(t.Name())), WithValidator(compareJSON(t))))
}

func TestCreatePage(t *testing.T) {
	ctx := context.Background()

	params := CreatePageParams{}
	params.Parent(Parent{Type: "database_id", DatabaseId: DATABASE})
	params.Properties(map[string]PropertyValue{"title": {Title: []RichText{{Text: &RichTextText{Content: "test"}}}}})
	params.Cover(File{External: &FileExternal{Url: "https://picsum.photos/200"}})

	_, err := client.CreatePage(ctx, params, WithRoundTripper(useCache(t.Name())), WithValidator(compareJSON(t)))
	if err != nil {
		t.Error(err)
	}
}

func TestRetrievePage(t *testing.T) {
	ctx := context.Background()
	if _, err := client.RetrievePage(ctx, STANDALONE_PAGE, WithRoundTripper(useCache(t.Name())), WithValidator(compareJSON(t))); err != nil {
		t.Fatal(err)
	}
}

func TestRetrievePagePropertyItem(t *testing.T) {
	ctx := context.Background()
	if db, err := client.RetrieveDatabase(ctx, DATABASE, WithRoundTripper(useCache(t.Name())), WithValidator(compareJSON(t))); err != nil {
		t.Fatal(err)
	} else {
		for name, prop := range db.Properties {
			name, prop := name, prop
			for i, pageId := range []uuid.UUID{DATABASE_PAGE_FOR_READ1, DATABASE_PAGE_FOR_READ2} {
				testName := fmt.Sprintf("%s_%d", name, i+1)
				t.Run(testName, func(t *testing.T) {
					if _, err := client.RetrievePagePropertyItem(ctx, pageId, prop.Id, WithRoundTripper(useCache(t.Name())), WithValidator(compareJSON(t))); err != nil {
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
			if _, err := client.UpdatePageProperties(ctx, DATABASE_PAGE_FOR_WRITE, params, WithValidator(compareJSON(t)), WithRoundTripper(useCache(t.Name()))); err != nil {
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
			if pagi, err := client.QueryDatabase(ctx, DATABASE, params, WithRoundTripper(useCache(t.Name())), WithValidator(compareJSON(t))); err != nil {
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
	if pagi, err := client.RetrieveBlockChildren(ctx, STANDALONE_PAGE, WithValidator(compareJSON(t))); err != nil {
		t.Fatal(err)
	} else {
		for _, block := range pagi.Results {
			block := block
			t.Run(block.Id.String(), func(t *testing.T) {
				if _, err := client.DeleteBlock(ctx, block.Id, WithValidator(compareJSON(t))); err != nil {
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
		if _, err := client.AppendBlockChildren(ctx, STANDALONE_PAGE, params, WithValidator(compareJSON(t))); err != nil {
			t.Fatal(err)
		}
	})
}
