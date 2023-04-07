package methoddoc

import (
	"encoding/json"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ParseMethod(url string) (*SSRProps, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	ssrProps := SSRProps{}
	if err := json.Unmarshal(ssrPropsBytes, &ssrProps); err != nil {
		return nil, err
	}

	return &ssrProps, nil
}
