package doc2api

import (
	"encoding/json"
	"fmt"
)

// objectDocElement はNotion API Reference のObjects以下 (e.g. https://developers.notion.com/reference/block) で
// 出現するエレメントのインターフェイスです
type objectDocElement interface {
	isObjectDocElement()
}

var _ = []objectDocElement{
	&objectDocHeadingElement{},
	&objectDocParagraphElement{},
	&objectDocCodeElement{},
	&objectDocCalloutElement{},
	&objectDocParametersElement{},
	&objectDocAPIHeaderElement{},
}

// objectDocHeadingElement は見出しのobjectDocElementです
type objectDocHeadingElement struct {
	Text string
}

// objectDocParagraphElement は段落のobjectDocElementです
type objectDocParagraphElement struct {
	Text string
}

type objectDocCodeElement struct {
	Codes []objectDocCodeElementCode `json:"codes"`
}

type objectDocCodeElementCode struct {
	Name     string `json:"string"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type objectDocCalloutElement struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type objectDocParametersElement []objectDocParameter

func (t *objectDocParametersElement) UnmarshalJSON(data []byte) error {
	raw := struct {
		Data map[string]string `json:"data"`
		Cols int               `json:"cols"`
		Rows int               `json:"rows"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mapSlice := make([]map[string]string, raw.Rows)
	for r := range mapSlice {
		m := map[string]string{}
		for c := 0; c < raw.Cols; c++ {
			m[raw.Data[fmt.Sprintf("h-%d", c)]] = stripMarkdown(raw.Data[fmt.Sprintf("%d-%d", r, c)])
		}
		mapSlice[r] = m
	}

	data, err := json.Marshal(mapSlice)
	if err != nil {
		return err
	}

	type Alias objectDocParametersElement
	return json.Unmarshal(data, (*Alias)(t))
}

type objectDocAPIHeaderElement struct {
	Title string `json:"title"`
}

func (t *objectDocHeadingElement) isObjectDocElement()    {}
func (t *objectDocParagraphElement) isObjectDocElement()  {}
func (t *objectDocCodeElement) isObjectDocElement()       {}
func (t *objectDocCalloutElement) isObjectDocElement()    {}
func (t *objectDocParametersElement) isObjectDocElement() {}
func (t *objectDocAPIHeaderElement) isObjectDocElement()  {}
