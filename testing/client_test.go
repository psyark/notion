package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
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
	// TODO 環境変数に移動
	ROOT                    = uuid.MustParse("9c20de5e26af4959a26e390b537af4c8") // https://www.notion.so/Root-9c20de5e26af4959a26e390b537af4c8
	STANDALONE_PAGE         = uuid.MustParse("b05213d5c3af4de6924cc9b106ae93ec") // https://www.notion.so/Page-b05213d5c3af4de6924cc9b106ae93ec
	DATABASE                = uuid.MustParse("edd0404128004a83bd29deb729221ec7") // https://www.notion.so/edd0404128004a83bd29deb729221ec7
	DATABASE_PAGE_FOR_READ1 = uuid.MustParse("7e01d5af9d0e4d2584e4d5bfc39b65bf") // https://www.notion.so/ABCDEFG-7e01d5af9d0e4d2584e4d5bfc39b65bf
	DATABASE_PAGE_FOR_READ2 = uuid.MustParse("7e1105bc19a64a1381453cff0b488092") // https://www.notion.so/7e1105bc19a64a1381453cff0b488092
	DATABASE_PAGE_FOR_WRITE = uuid.MustParse("b8ff7c186ef2416cb9654daf0d7aa961") // https://www.notion.so/PageToUpdate-b8ff7c186ef2416cb9654daf0d7aa961
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
				params := UpdatePagePropertiesParams{}
				params.InTrash(true)
				lo.Must(client.UpdatePageProperties(ctx, pd.Id, params))
			}
		}
	}

	m.Run()
}

func TestClient(t *testing.T) {
	ctx := context.Background()

	var generatedPage *Page

	t.Run("CreatePage", func(t *testing.T) {
		params := CreatePageParams{}
		params.Parent(Parent{PageId: ROOT})
		params.Icon(Emoji{Emoji: "✨"})
		params.Cover(File{External: &FileExternal{Url: "https://picsum.photos/200"}})
		params.Properties(map[string]PropertyValue{
			"title": {Title: NewRichTextArray(fmt.Sprintf("生成されたページ (%s)", time.Now().Format(time.RFC3339)))},
		})
		generatedPage = lo.Must(client.CreatePage(ctx, params))
	})

	t.Run("RetrievePage", func(t *testing.T) {
		t.Parallel()
		lo.Must(client.RetrievePage(ctx, generatedPage.Id, useCache(t), compareJSON(t)))
	})

	var generatedDatabase *Database

	t.Run("CreateDatabase", func(t *testing.T) {
		params := CreateDatabaseParams{}
		params.Parent(Parent{PageId: generatedPage.Id})
		params.Title(NewRichTextArray("生成されたデータベース"))
		params.Properties(map[string]PropertySchema{
			"タイトル":     {Title: &struct{}{}},
			"テキスト":     {RichText: &struct{}{}},
			"数値":       {Number: &PropertySchemaNumber{Format: "number_with_commas"}},
			"セレクト":     {Select: &PropertySchemaSelect{Options: []PropertySchemaOption{{Name: "赤", Color: "red"}}}},
			"マルチセレクト":  {MultiSelect: &PropertySchemaMultiSelect{Options: []PropertySchemaOption{{Name: "赤", Color: "red"}}}},
			"日付":       {Date: &struct{}{}},
			"ユーザー":     {People: &struct{}{}},
			"ファイル":     {Files: &struct{}{}},
			"チェックボックス": {Checkbox: &struct{}{}},
			"URL":      {Url: &struct{}{}},
			"メール":      {Email: &struct{}{}},
			"電話":       {Email: &struct{}{}},
			"数式":       {Formula: &PropertySchemaFormula{}},
			"リレーション":   {Relation: &PropertySchemaRelation{SingleProperty: &struct{}{}, DatabaseId: DATABASE}},
			"ロールアップ":   {Rollup: &PropertySchemaRollup{RollupPropertyId: "title", RelationPropertyName: "リレーション", Function: "show_original"}},
			"作成日時":     {CreatedTime: &struct{}{}},
			"作成者":      {CreatedBy: &struct{}{}},
			"最終更新日時":   {LastEditedTime: &struct{}{}},
			"最終更新者":    {LastEditedBy: &struct{}{}},
			"ボタン":      {Button: &struct{}{}},
			"ID":       {UniqueId: &PropertySchemaUniqueId{Prefix: lo.ToPtr("OK")}},
		})

		generatedDatabase = lo.Must(client.CreateDatabase(ctx, params, compareJSON(t)))

		t.Run("CreatePage", func(t *testing.T) {
			params := CreatePageParams{}
			params.Cover(File{External: &FileExternal{Url: "https://picsum.photos/200"}})
			params.Icon(Emoji{Emoji: "🍣"})
			params.Parent(Parent{DatabaseId: generatedDatabase.Id})
			params.Properties(PropertyValueMap{
				"タイトル":     {Title: NewRichTextArray("生成されたエントリー")},
				"テキスト":     {RichText: NewRichTextArray("これは生成されたエントリーです")},
				"数値":       {Number: lo.ToPtr(math.Pi)},
				"セレクト":     {Select: &Option{Name: "赤"}},
				"マルチセレクト":  {MultiSelect: []Option{{Name: "赤"}}},
				"日付":       {Date: &PropertyValueDate{Start: "2024-07-27"}}, // TODO 値を入れないのはどうする？
				"ユーザー":     {People: []User{}},                              // TODO 値をいれる
				"チェックボックス": {Checkbox: true},
				"URL":      {Url: lo.ToPtr("https://picsum.photos/200")},
				"メール":      {Email: lo.ToPtr("me@example.com")},
				"電話":       {Email: lo.ToPtr("117")},
				"リレーション":   {Relation: []PageReference{}},
			})
			lo.Must(client.CreatePage(ctx, params, compareJSON(t)))
		})
	})

	t.Run("RetrieveDatabase", func(t *testing.T) {
		t.Parallel()
		lo.Must(client.RetrieveDatabase(ctx, generatedDatabase.Id, compareJSON(t)))
	})

	t.Run("AppendBlockChildren", func(t *testing.T) {
		t.Parallel()
		params := AppendBlockChildrenParams{}
		params.Children([]Block{
			{Divider: &struct{}{}},

			{Bookmark: &BlockBookmark{Url: "http://example.com"}},
			{Breadcrumb: &struct{}{}},
			{BulletedListItem: &BlockBulletedListItem{RichText: NewRichTextArray("箇条書きリスト")}},
			{Callout: &BlockCallout{
				RichText: NewRichTextArray("コールアウト"),
				Icon:     &Emoji{Emoji: "🍣"},
			}},
			{Code: &BlockCode{
				RichText: NewRichTextArray(`// hoge`),
				Language: "go",
			}},
			{Image: &File{External: &FileExternal{Url: "https://placehold.jp/640x640.png"}}},
			{Heading1: &BlockHeading{RichText: NewRichTextArray("Heading 1")}},
			{Heading2: &BlockHeading{RichText: NewRichTextArray("Heading 2")}},
			{Heading3: &BlockHeading{RichText: NewRichTextArray("Heading 3")}},
			{ToDo: &BlockToDo{RichText: NewRichTextArray("To Do")}},
			{SyncedBlock: &BlockSyncedBlock{
				Children: []Block{
					{Paragraph: &BlockParagraph{RichText: NewRichTextArray("synced")}},
				},
			}},
		})
		lo.Must(client.AppendBlockChildren(ctx, generatedPage.Id, params, compareJSON(t)))
	})
}

func TestRetrievePagePropertyItem(t *testing.T) {
	ctx := context.Background()
	if db, err := client.RetrieveDatabase(ctx, DATABASE, useCache(t), compareJSON(t)); err != nil {
		t.Fatal(err)
	} else {
		for name, prop := range db.Properties {
			name, prop := name, prop
			for i, pageId := range []uuid.UUID{DATABASE_PAGE_FOR_READ1, DATABASE_PAGE_FOR_READ2} {
				testName := fmt.Sprintf("%s_%d", name, i+1)
				t.Run(testName, func(t *testing.T) {
					if _, err := client.RetrievePagePropertyItem(ctx, pageId, prop.Id, useCache(t), compareJSON(t)); err != nil {
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
			if _, err := client.UpdatePageProperties(ctx, DATABASE_PAGE_FOR_WRITE, params, compareJSON(t), useCache(t)); err != nil {
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
			if pagi, err := client.QueryDatabase(ctx, DATABASE, params, useCache(t), compareJSON(t)); err != nil {
				t.Fatal(err)
			} else {
				fmt.Println(len(pagi.Results))
			}
		})
	}
}

func TestRetrieveBlockChildren(t *testing.T) {
	ctx := context.Background()

	pagi := lo.Must(client.RetrieveBlockChildren(ctx, STANDALONE_PAGE, compareJSON(t)))
	for _, block := range pagi.Results {
		t.Run(block.Id.String(), func(t *testing.T) {
			lo.Must(client.DeleteBlock(ctx, block.Id, compareJSON(t)))
		})
	}
}
