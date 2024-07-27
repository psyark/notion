package testing

import (
	"context"
	_ "embed"
	"encoding/json"
	"testing"

	. "github.com/psyark/notion"
	"github.com/psyark/notion/binding"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed "testdata/tagged.txt"
	tagged string
	//go:embed "testdata/update_page_params.json"
	updatePageParams string
)

type TheDatabase struct {
	Select                 *Option            `notion:"DaP%40"`
	DualRelation           []PageReference    `notion:"Dopp"`
	SingleRelation         []PageReference    `notion:"kOoD"`
	ArrayRollup            *Rollup            `notion:"lxhZ"`
	DualRelation_40back_41 []PageReference    `notion:"vxwW"`
	LastEditedTime         ISO8601String      `notion:"%7B%7Cmj"`
	Checkbox               bool               `notion:"%3Dh%3AT"`
	LastEditedBy           User               `notion:"CA~Q"`
	Formula                *Formula           `notion:"kutj"`
	Number                 *float64           `notion:"wSuU"`
	Phone                  *string            `notion:"%7Cb%60H"`
	Status                 *Option            `notion:"~_pB"`
	CreatedTime            ISO8601String      `notion:"Ldgn"`
	NumberRollup           *Rollup            `notion:"QdI%3C"`
	CreatedBy              User               `notion:"TB%5Dl"`
	Text                   []RichText         `notion:"Vl%40o"`
	URL                    *string            `notion:"nKu_"`
	MultiSelect            []Option           `notion:"qe%60%5E"`
	File                   []File             `notion:"%7Dlj%7B"`
	User                   []User             `notion:"Ui%5B%3A"`
	Date                   *PropertyValueDate `notion:"gegF"`
	Mail                   *string            `notion:"l_GI"`
	Title                  []RichText         `notion:"title"`
}

func TestBinding(t *testing.T) {
	ctx := context.Background()

	t.Run("ToTaggedStruct", func(t *testing.T) {
		db := lo.Must(client.RetrieveDatabase(ctx, DATABASE, WithRoundTripper(useCache(t.Name()))))
		ts := binding.ToTaggedStruct(db)

		assert.Equal(t, tagged, ts)
	})

	t.Run("UnmarshalPage", func(t *testing.T) {
		pagi := lo.Must(client.QueryDatabase(ctx, DATABASE, QueryDatabaseParams{}, WithRoundTripper(useCache(t.Name()))))
		for _, page := range pagi.Results {
			hoge := &TheDatabase{}
			lo.Must0(binding.UnmarshalPage(&page, hoge))
			lo.Must(json.MarshalIndent(hoge, "", "  "))
		}
	})

	t.Run("GetUpdatePageParams", func(t *testing.T) {
		page := lo.Must(client.RetrievePage(ctx, DATABASE_PAGE_FOR_READ1, WithRoundTripper(useCache(t.Name()))))

		hoge := &TheDatabase{}
		lo.Must0(binding.UnmarshalPage(page, hoge))

		// hoge.Title = append(hoge.Title, &TextRichText{Text: Text{Content: "HOGE"}})
		hoge.Number = lo.ToPtr(*hoge.Number + 100)
		hoge.URL = lo.ToPtr(*hoge.URL + "/hoge")

		params := lo.Must(binding.GetUpdatePageParams(hoge, page))
		data := lo.Must(json.MarshalIndent(params, "", "  "))

		assert.JSONEq(t, updatePageParams, string(data))
	})
}
