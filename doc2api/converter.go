// Package doc2api は Notion API Reference の更新を検知し、Goコードへの適切な変換を
// 継続的に行うための一連の仕組みを提供します。
//
// Goコードへの変換ルールは、命令形のコードではなくデータ構造として objects.*.go に格納されます。
// このデータ構造には Notion API Referenceのローカルコピーも含まれるため、
// ドキュメントの更新に対してGoコードへの変換ルールが古いままになることを防ぎます。
package doc2api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// converter はNotion API ReferenceからGoコードへの変換ルールです。
type converter struct {
	url      string // ドキュメントのURL
	fileName string // 出力するファイル名
	matchers []elementMatcher
}

// convert は変換を実行します
func (c converter) convert() error {
	// URLの取得
	res, err := http.Get(c.url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// goqueryでのパース
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	ssrPropsBytes := []byte(doc.Find(`#ssr-props`).AttrOr("data-initial-props", ""))
	ssrProps := struct {
		Doc struct {
			Body string `json:"body"`
		} `json:"doc"`
	}{}
	if err := json.Unmarshal(ssrPropsBytes, &ssrProps); err != nil {
		return err
	}

	fmt.Println(c.fileName, c.url)

	lines := strings.Split(ssrProps.Doc.Body, "\n")
	odt := &objectDocTokenizer{lines, 0}

	for i := 0; ; i++ {
		remote, err := odt.next()
		if err != nil {
			return err
		}

		if len(c.matchers) < i+1 {
			return fmt.Errorf("matcherが足りません：element[%v]=%#v", i, remote)
		} else if err := c.matchers[i].match(remote); err != nil {
			return err
		} else {
			fmt.Println("OK")
		}
	}
}

var registeredConverters []converter

// registerConverter は後で実行するためにconverterを登録します
func registerConverter(c converter) {
	registeredConverters = append(registeredConverters, c)
}

// convertAll は登録された全てのconverterで変換を実行します
func convertAll() error {
	for _, c := range registeredConverters {
		if err := c.convert(); err != nil {
			return fmt.Errorf("convert: %s: %w", c.fileName, err)
		}
	}
	return nil
}

// elementMatcher はNotion API Referenceのローカルコピーと、Goコードへの変換ルールが一体となったデータです
//
// converter.convertは保有する各elementMatcherに対し、matchを呼び出します
// 各elementMatcherはローカルコピーとの比較を行い、失敗した場合はerrorを返し
// 成功した場合はコードを出力します。
type elementMatcher interface {
	match(remote objectDocElement) error
}

var _ = []elementMatcher{
	paragraphElementMatcher{},
	headingElementMatcher{},
	codeElementMatcher{},
}

type paragraphElementMatcher struct {
	local  *objectDocParagraphElement
	output func(*objectDocParagraphElement) error
}

func (m paragraphElementMatcher) match(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocParagraphElement); ok {
		if *remote != *m.local {
			return fmt.Errorf("mismatch: remote=%#v, local=%#v", remote, m.local)
		}
		return m.output(remote)
	}
	return fmt.Errorf("mismatch: remote is not objectDocParagraphElement (%#v)", remote)
}

type headingElementMatcher struct {
	local  *objectDocHeadingElement
	output func(*objectDocHeadingElement) error
}

func (m headingElementMatcher) match(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocHeadingElement); ok {
		if remote.Text != m.local.Text {
			return fmt.Errorf("mismatch: remote=%#v, local=%#v", remote, m.local)
		}
		return m.output(remote)
	}
	return fmt.Errorf("mismatch: remote is not objectDocHeadingElement (%#v)", remote)
}

type calloutElementMatcher struct {
	local  *objectDocCalloutElement
	output func(*objectDocCalloutElement) error
}

func (m calloutElementMatcher) match(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocCalloutElement); ok {
		if *remote != *m.local {
			return fmt.Errorf("mismatch: remote=%#v, local=%#v", remote, m.local)
		}
		return m.output(remote)
	}
	return fmt.Errorf("mismatch: remote is not objectDocCalloutElement (%#v)", remote)
}

type codeElementMatcher struct {
	local  *objectDocCodeElement
	output func(*objectDocCodeElement) error
}

func (m codeElementMatcher) match(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocCodeElement); ok {
		if len(remote.Codes) != len(m.local.Codes) {
			return fmt.Errorf("mismatch: len(remote.Codes)=%v, len(local.Codes)=%v", len(remote.Codes), len(m.local.Codes))
		}
		for i := range remote.Codes {
			if remote.Codes[i] != m.local.Codes[i] {
				return fmt.Errorf("mismatch: remote.Codes[%v]=%v, local.Codes[%v]=%v", i, remote.Codes[i], i, m.local.Codes[i])
			}
		}
		return m.output(remote)
	}
	return fmt.Errorf("mismatch: remote is not objectDocCodeElement (%#v)", remote)
}

type parametersElementMatcher struct {
	local  *objectDocParametersElement
	output func(*objectDocParametersElement) error
}

func (m parametersElementMatcher) match(remote objectDocElement) error {
	if remote, ok := remote.(*objectDocParametersElement); ok {
		if len(*remote) != len(*m.local) {
			return fmt.Errorf("mismatch: len(remote)=%v, len(local)=%v", len(*remote), len(*m.local))
		}
		for i := range *remote {
			if (*remote)[i] != (*m.local)[i] {
				return fmt.Errorf("mismatch: remote[%v]=%v, local[%v]=%v", i, (*remote)[i], i, (*m.local)[i])
			}
		}
		return m.output(remote)
	}
	return fmt.Errorf("mismatch: remote is not objectDocParametersElement (%#v)", remote)
}
