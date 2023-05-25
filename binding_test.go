package notion

import (
	"context"
	"encoding/json"
	"fmt"

	nullv4 "gopkg.in/guregu/null.v4"
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
	Number                 nullv4.Float       `notion:"wSuU"`
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

func Example_binding() {
	ctx := context.Background()
	db, err := cli.RetrieveDatabase(ctx, DATABASE, requestId("ExampleBinding"), useCache())
	if err != nil {
		panic(err)
	}

	ts := ToTaggedStruct(db)
	fmt.Println(ts)

	pagi, err := cli.QueryDatabase(ctx, DATABASE, QueryDatabaseParams{}, requestId("ExampleBinding_Query"), useCache())
	if err != nil {
		panic(err)
	}

	pages, err := pagi.Pages()
	if err != nil {
		panic(err)
	}
	for _, page := range pages {
		hoge := &TheDatabase{}
		if err := UnmarshalPage(&page, hoge); err != nil {
			panic(err)
		}

		// _, _ = fmt.Println(hoge.Title, hoge.Text, hoge.Number, hoge.Date, hoge.URL)
	}

	{
		page, err := cli.RetrievePage(ctx, DATABASE_PAGE_FOR_READ1, requestId("ExampleBinding_Page"), useCache())
		if err != nil {
			panic(err)
		}

		hoge := &TheDatabase{}
		if err := UnmarshalPage(page, hoge); err != nil {
			panic(err)
		}

		// hoge.Title = append(hoge.Title, &TextRichText{Text: Text{Content: "HOGE"}})
		hoge.Number = nullv4.FloatFrom(hoge.Number.Float64 + 100)
		newUrl := *hoge.URL + "/hoge"
		hoge.URL = &newUrl
		if params, err := GetUpdatePageParams(hoge, page); err != nil {
			panic(err)
		} else {
			data, _ := json.MarshalIndent(params, "", "  ")
			fmt.Println(string(data))
		}
	}

	// Output:
	// type Database struct {
	// 	ArrayRollup            *Rollup            `notion:"lxhZ"`
	// 	Checkbox               bool               `notion:"%3Dh%3AT"`
	// 	CreatedBy              User               `notion:"TB%5Dl"`
	// 	CreatedTime            ISO8601String      `notion:"Ldgn"`
	// 	Date                   *PropertyValueDate `notion:"gegF"`
	// 	DualRelation           []PageReference    `notion:"Dopp"`
	// 	DualRelation_40back_41 []PageReference    `notion:"vxwW"`
	// 	File                   []File             `notion:"%7Dlj%7B"`
	// 	Formula                *Formula           `notion:"kutj"`
	// 	LastEditedBy           User               `notion:"CA~Q"`
	// 	LastEditedTime         ISO8601String      `notion:"%7B%7Cmj"`
	// 	Mail                   *string            `notion:"l_GI"`
	// 	MultiSelect            []Option           `notion:"qe%60%5E"`
	// 	Number                 nullv4.Float       `notion:"wSuU"`
	// 	NumberRollup           *Rollup            `notion:"QdI%3C"`
	// 	Phone                  *string            `notion:"%7Cb%60H"`
	// 	Select                 *Option            `notion:"DaP%40"`
	// 	SingleRelation         []PageReference    `notion:"kOoD"`
	// 	Status                 *Option            `notion:"~_pB"`
	// 	Text                   []RichText         `notion:"Vl%40o"`
	// 	Title                  []RichText         `notion:"title"`
	// 	URL                    *string            `notion:"nKu_"`
	// 	User                   []User             `notion:"Ui%5B%3A"`
	// }
	// {
	//   "properties": {
	//     "nKu_": {
	//       "type": "url",
	//       "url": "http://example.com/hoge"
	//     },
	//     "wSuU": {
	//       "number": 103.1415926535898,
	//       "type": "number"
	//     }
	//   }
	// }
}
