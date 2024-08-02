package testing

import (
	"context"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/psyark/notion"
	"github.com/psyark/notion/binding"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed "testdata/tagged.txt"
	expectedTaggedStruct string

	//go:embed "testdata/records.json"
	expectedRecords string

	//go:embed "testdata/update_page_params.json"
	expectedUPP string
)

type TheDatabase struct {
	Select                 *notion.Option            `notion:"DaP%40"`
	DualRelation           []notion.PageReference    `notion:"Dopp"`
	SingleRelation         []notion.PageReference    `notion:"kOoD"`
	ArrayRollup            *notion.Rollup            `notion:"lxhZ"`
	DualRelation_40back_41 []notion.PageReference    `notion:"vxwW"`
	LastEditedTime         notion.ISO8601String      `notion:"%7B%7Cmj"`
	Checkbox               bool                      `notion:"%3Dh%3AT"`
	LastEditedBy           notion.User               `notion:"CA~Q"`
	Formula                *notion.Formula           `notion:"kutj"`
	Number                 *float64                  `notion:"wSuU"`
	Phone                  *string                   `notion:"%7Cb%60H"`
	Status                 *notion.Option            `notion:"~_pB"`
	CreatedTime            notion.ISO8601String      `notion:"Ldgn"`
	NumberRollup           *notion.Rollup            `notion:"QdI%3C"`
	CreatedBy              notion.User               `notion:"TB%5Dl"`
	Text                   []notion.RichText         `notion:"Vl%40o"`
	URL                    *string                   `notion:"nKu_"`
	MultiSelect            []notion.Option           `notion:"qe%60%5E"`
	File                   []notion.File             `notion:"%7Dlj%7B"`
	User                   []notion.User             `notion:"Ui%5B%3A"`
	Date                   *notion.PropertyValueDate `notion:"gegF"`
	Mail                   *string                   `notion:"l_GI"`
	Title                  []notion.RichText         `notion:"title"`
}

func TestBinding(t *testing.T) {
	ctx := context.Background()

	t.Run("ToTaggedStruct", func(t *testing.T) {
		db := lo.Must(client.RetrieveDatabase(ctx, DATABASE, useCache(t)))
		actualTaggedStruct := binding.ToTaggedStruct(db)
		assert.Equal(t, expectedTaggedStruct, actualTaggedStruct)
	})

	t.Run("UnmarshalPage", func(t *testing.T) {
		params := notion.QueryDatabaseParams{}
		params.Sorts([]notion.Sort{{Timestamp: "created_time", Direction: "ascending"}})
		pagi := lo.Must(client.QueryDatabase(ctx, DATABASE, params, useCache(t)))

		records := lo.Map(pagi.Results, func(page notion.Page, _ int) TheDatabase {
			record := TheDatabase{}
			lo.Must0(binding.UnmarshalPage(&page, &record))
			return record
		})

		actualRecords := string(lo.Must(json.MarshalIndent(records, "", "  ")))
		assert.JSONEq(t, expectedRecords, actualRecords)
	})

	t.Run("GetUpdatePageParams", func(t *testing.T) {
		page := lo.Must(client.RetrievePage(ctx, DATABASE_PAGE_FOR_READ1, useCache(t)))

		hoge := &TheDatabase{}
		lo.Must0(binding.UnmarshalPage(page, hoge))

		// hoge.Title = append(hoge.Title, &TextRichText{Text: Text{Content: "HOGE"}})
		hoge.Number = lo.ToPtr(*hoge.Number + 100)
		hoge.URL = lo.ToPtr(*hoge.URL + "/hoge")

		upp := lo.Must(binding.GetUpdatePageParams(hoge, page))
		actualUPP := string(lo.Must(json.MarshalIndent(upp, "", "  ")))
		assert.JSONEq(t, expectedUPP, actualUPP)
	})
}
