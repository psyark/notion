package objectdoc

import (
	"strings"
	"sync"
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

func (b *builder) addSpecificObject(name string, comment string) *specificObject {
	o := &specificObject{}
	o.name_ = strings.TrimSpace(name)
	o.comment = comment
	b.localSymbols = append(b.localSymbols, o)
	b.globalSymbols.Store(name, o)
	return o
}

func (b *builder) addAbstractObject(name string, specifiedBy string, comment string) *abstractObject {
	o := &abstractObject{}
	o.name_ = strings.TrimSpace(name)
	o.derivedIdentifierKey = specifiedBy
	o.comment = comment
	b.localSymbols = append(b.localSymbols, o)
	b.globalSymbols.Store(name, o)
	return o
}

func (b *builder) addAbstractObjectToGlobalIfNotExists(name string, specifiedBy string) *abstractObject {
	if o := b.getAbstractObject(name); o != nil {
		return o
	}
	return b.global.addAbstractObject(name, specifiedBy, "")
}

func (b *builder) addAbstractList(targetName string, name string) *abstractList {
	o := &abstractList{}
	o.name_ = strings.TrimSpace(name)
	o.targetName = strings.TrimSpace(targetName)
	b.localSymbols = append(b.localSymbols, o)
	b.globalSymbols.Store(name, o)
	return o
}

func (b *builder) addAbstractMap(targetName string, name string) *abstractMap {
	o := &abstractMap{}
	o.name_ = strings.TrimSpace(name)
	o.targetName = strings.TrimSpace(targetName)
	b.localSymbols = append(b.localSymbols, o)
	b.globalSymbols.Store(name, o)
	return o
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

func (b *builder) getAbstractObject(name string) *abstractObject {
	if item, ok := b.globalSymbols.Load(name); ok {
		if item, ok := item.(*abstractObject); ok {
			return item
		}
	}
	return nil
}

func (b *builder) getSpecificObject(name string) *specificObject {
	if item, ok := b.globalSymbols.Load(name); ok {
		if item, ok := item.(*specificObject); ok {
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
