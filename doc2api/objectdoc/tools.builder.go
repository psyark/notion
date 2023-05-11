package objectdoc

import (
	"strings"
	"sync"

	"github.com/dave/jennifer/jen"
)

// builderは生成されるオブジェクトやフィールドの依存関係の構築を行います
// 実行時にオブジェクトが登録される順番が一定しなくても、出力されるソースコードは冪等に保たれます
type builder struct {
	fileName      string
	url           string
	globalSymbols *sync.Map
	localSymbols  []symbolCoder
	testSymbols   []symbolCoder

	// Deprecated:
	global *builder // TODO ローカル/グローバルビルダーを作るのではなく、唯一のビルダーを作る
}

func (b *builder) addConcreteObject(name string, comment string) *concreteObject {
	obj := &concreteObject{}
	obj.name_ = strings.TrimSpace(name)
	obj.comment = comment
	b.localSymbols = append(b.localSymbols, obj)
	b.globalSymbols.Store(name, obj)
	return obj
}

func (b *builder) addAdaptiveObject(name string, discriminatorKey string, comment string) *adaptiveObject {
	ao := &adaptiveObject{}
	ao.name_ = name
	ao.discriminatorKey = discriminatorKey
	ao.comment = comment
	if discriminatorKey != "" {
		ao.addFields(&field{
			name:     discriminatorKey,
			typeCode: jen.String(),
		})
	}
	b.localSymbols = append(b.localSymbols, ao)
	b.globalSymbols.Store(name, ao)
	return ao
}

func (b *builder) addUnionToGlobalIfNotExists(name string, identifierKey string) *unionObject {
	if u := getSymbol[unionObject](b, name); u != nil {
		return u
	}

	u := &unionObject{}
	u.name_ = strings.TrimSpace(name)
	u.identifierKey = identifierKey
	b.global.localSymbols = append(b.global.localSymbols, u)
	b.globalSymbols.Store(name, u)

	return u
}

func (b *builder) addAlwaysStringIfNotExists(value string) {
	as := alwaysString(value)
	if _, ok := b.globalSymbols.Load(as.name()); ok {
		return
	}
	b.globalSymbols.Store(as.name(), as)
	b.localSymbols = append(b.localSymbols, as)
}

func (b *builder) getSymbol(name string) symbolCoder {
	if item, ok := b.globalSymbols.Load(name); ok {
		if item, ok := item.(symbolCoder); ok {
			return item
		}
	}
	return nil
}

func getSymbol[T concreteObject | unionObject | adaptiveObject](b *builder, name string) *T {
	if item, ok := b.globalSymbols.Load(name); ok {
		if item, ok := item.(*T); ok {
			return item
		}
	}
	return nil
}

func (b *builder) addUnmarshalTest(targetName string, jsonCode string) {
	ut := &unmarshalTest{targetName: targetName}
	if item, ok := b.globalSymbols.Load(ut.name()); ok {
		if item, ok := item.(*unmarshalTest); ok {
			item.jsonCodes = append(item.jsonCodes, jsonCode)
			return
		}
	}

	ut.jsonCodes = append(ut.jsonCodes, jsonCode)
	b.globalSymbols.Store(ut.name(), ut)
	b.testSymbols = append(b.testSymbols, ut)
}
