// Package doc2api は Notion API Reference の更新を検知し、Goコードへの適切な変換を
// 継続的に行うための一連の仕組みを提供します。
//
// Goコードへの変換ルールは、命令形のコードではなくデータ構造として objects.*.go に格納されます。
// このデータ構造には Notion API Referenceのローカルコピーも含まれるため、
// ドキュメントの更新に対してGoコードへの変換ルールが古いままになることを防ぎます。
package doc2api

import "fmt"

// converter はNotion API ReferenceからGoコードへの変換ルールです。
type converter struct {
	url      string // ドキュメントのURL
	fileName string // 出力するファイル名
	elements []docElement
}

// convert は変換を実行します
func (c converter) convert() error {
	fmt.Println(c.fileName, c.url)
	return nil
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

type docElement interface {
	equals(docElement) bool
}

type headingElement string

func (e headingElement) equals(e2 docElement) bool {
	if e2, ok := e2.(headingElement); ok {
		return string(e) == string(e2)
	}
	return false
}

type codeElement string

func (e codeElement) equals(e2 docElement) bool {
	if e2, ok := e2.(codeElement); ok {
		return string(e) == string(e2)
	}
	return false
}
