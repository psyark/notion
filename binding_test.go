package notion

import (
	"context"
	"fmt"
	"testing"

	nullv4 "gopkg.in/guregu/null.v4"
)

type TheDatabase struct {
	Date                   *DatePropertyValueData `notion:"gegF"`
	DualRelation           []PageReference        `notion:"Dopp"`
	Formula                Formula                `notion:"kutj"`
	DualRelation_40back_41 []PageReference        `notion:"vxwW"`
	Number                 nullv4.Float           `notion:"wSuU"`
	CreatedBy              PartialUser            `notion:"TB%5Dl"`
	Mail                   nullv4.String          `notion:"l_GI"`
	SingleRelation         []PageReference        `notion:"kOoD"`
	LastEditedBy           User                   `notion:"CA~Q"`
	ArrayRollup            Rollup                 `notion:"lxhZ"`
	Select                 *Option                `notion:"DaP%40"`
	Phone                  nullv4.String          `notion:"%7Cb%60H"`
	LastEditedTime         ISO8601String          `notion:"%7B%7Cmj"`
	Text                   RichTextArray          `notion:"Vl%40o"`
	Title                  RichTextArray          `notion:"title"`
	File                   Files                  `notion:"%7Dlj%7B"`
	Status                 Option                 `notion:"~_pB"`
	Checkbox               bool                   `notion:"%3Dh%3AT"`
	CreatedTime            ISO8601String          `notion:"Ldgn"`
	NumberRollup           Rollup                 `notion:"QdI%3C"`
	User                   Users                  `notion:"Ui%5B%3A"`
	URL                    nullv4.String          `notion:"nKu_"`
	MultiSelect            []Option               `notion:"qe%60%5E"`
}

func TestBinding(t *testing.T) {
	ctx := context.Background()
	db, err := cli.RetrieveDatabase(ctx, DATABASE, requestId(t.Name()), useCache())
	if err != nil {
		t.Fatal(err)
	}

	ToTaggedStruct(db)

	pagi, err := cli.QueryDatabase(ctx, DATABASE, &QueryDatabaseParams{}, requestId(t.Name()+"_Query"), useCache())
	if err != nil {
		t.Fatal(err)
	}

	for _, page := range pagi.Results {
		hoge := &TheDatabase{}
		if err := UnmarshalPage(page, hoge); err != nil {
			t.Fatal(err)
		}

		_, _ = fmt.Println(hoge.Title, hoge.Text, hoge.Number, hoge.Date, hoge.URL)
	}
}
