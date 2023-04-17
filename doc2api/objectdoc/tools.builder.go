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
	o.specifiedBy = specifiedBy
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

func (b *builder) addAlwaysStringIfNotExists(value string) {
	as := alwaysString(value)
	if _, ok := b.globalSymbols.Load(as.name()); ok {
		return
	}
	b.globalSymbols.Store(as.name(), as)
	b.localSymbols = append(b.localSymbols, as)
}

func (b *builder) getAbstractObject(name string) *abstractObject {
	if item, ok := b.globalSymbols.Load(name); ok {
		if item, ok := item.(*abstractObject); ok && item.name() == name {
			return item
		}
	}
	return nil
}

func (b *builder) getSpecificObject(name string) *specificObject {
	if item, ok := b.globalSymbols.Load(name); ok {
		if item, ok := item.(*specificObject); ok && item.name() == name {
			return item
		}
	}
	return nil
}
