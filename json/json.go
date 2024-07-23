// Package json は、標準のjsonの以下のルールを置き換えます
//
// - スライス型の nil は [] にエンコードする
// - スライス型の nil は empty として扱う一方、長さ0のスライスは empty として扱わない
package json

import (
	jsoniter "github.com/json-iterator/go"
)

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.RegisterExtension(&customExtension{})
}

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
