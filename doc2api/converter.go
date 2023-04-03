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
			return fmt.Errorf("matcherが足りません：%v", i)
		} else if c.matchers[i].match(remote) {
			// OK
		} else {
			return fmt.Errorf("mismatch: remote=%#v", remote)
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
			return err
		}
	}
	return nil
}

// elementMatcher はNotion API Referenceの要素と、Goコードへの変換ルールが一体となったデータです
type elementMatcher interface {
	match(remote objectDocElement) bool
}

type paragraphElementMatcher struct {
	element objectDocParagraphElement
	output  func(objectDocParagraphElement) error
}

func (m paragraphElementMatcher) match(remote objectDocElement) bool {
	return remote == m.element
}

type headingElementMatcher struct {
	element objectDocHeadingElement
	output  func(objectDocHeadingElement) error
}

func (m headingElementMatcher) match(remote objectDocElement) bool {
	return remote == m.element
}

type codeElementMatcher struct {
	element objectDocCodeElement
	output  func(objectDocCodeElement) error
}

func (m codeElementMatcher) match(remote objectDocElement) bool {
	if remote, ok := remote.(objectDocCodeElement); ok {
		if len(remote.Codes) != len(m.element.Codes) {
			return false
		}
		for i := range remote.Codes {
			if remote.Codes[i] != m.element.Codes[i] {
				return false
			}
		}
		return true
	}
	return false
}
