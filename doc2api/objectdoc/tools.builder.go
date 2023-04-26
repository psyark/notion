package objectdoc

import (
	"strings"
	"sync"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
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

func (b *builder) addAbstractObject(name string, specifiedBy string, comment string) *abstractObject {
	obj := &abstractObject{}
	obj.name_ = strings.TrimSpace(name)
	obj.derivedIdentifierKey = specifiedBy
	obj.comment = comment
	b.localSymbols = append(b.localSymbols, obj)
	b.globalSymbols.Store(name, obj)
	return obj
}

type addDerivedOptions struct {
	derivedName            string
	addSpecificField       bool
	specificFieldMayNull   bool
	specificTypeIsAbstract bool
	specificTypeDIK        string
}
type addDerivedOption func(o *addDerivedOptions)

func withName(derivedName string) addDerivedOption {
	return func(o *addDerivedOptions) { o.derivedName = derivedName }
}
func addSpecificField(mayNull ...bool) addDerivedOption {
	return func(o *addDerivedOptions) {
		o.addSpecificField = true
		if len(mayNull) == 1 {
			o.specificFieldMayNull = mayNull[0]
		}
	}
}
func addAbstractSpecificField(derivedIdentifierKey string) addDerivedOption {
	return func(o *addDerivedOptions) {
		o.addSpecificField = true
		o.specificTypeIsAbstract = true
		o.specificTypeDIK = derivedIdentifierKey
	}
}

// addDerived はderivedIdentifierValueとparentNameから決まる名前で派生クラスを作成します
func (b *builder) addDerived(derivedIdentifierValue string, parentName string, comment string, options ...addDerivedOption) *concreteObject {
	opt := &addDerivedOptions{
		derivedName: strcase.UpperCamelCase(derivedIdentifierValue) + parentName,
	}
	for _, o := range options {
		o(opt)
	}

	parent := getSymbol[abstractObject](b, parentName)
	derived := &concreteObject{}
	derived.name_ = opt.derivedName
	derived.derivedIdentifierValue = derivedIdentifierValue
	derived.comment = comment

	if parent.derivedIdentifierKey != "" && derivedIdentifierValue != "" {
		derived.fields = append(derived.fields, &fixedStringField{name: parent.derivedIdentifierKey, value: derivedIdentifierValue})
	}

	// 親子関係を設定
	derived.setParent(parent)
	parent.derivedObjects = append(parent.derivedObjects, derived)

	b.localSymbols = append(b.localSymbols, derived)
	b.globalSymbols.Store(derived.name(), derived)

	// DIV固有のタイプとフィールドを作成
	if opt.addSpecificField {
		specifitFieldTypeName := opt.derivedName + "Data"
		if opt.specificTypeIsAbstract {
			derived.typeSpecificAbstract = b.addAbstractObject(specifitFieldTypeName, opt.specificTypeDIK, "")
			derived.addFields(&interfaceField{name: derivedIdentifierValue, typeName: specifitFieldTypeName})
		} else {
			derived.typeSpecificObject = b.addConcreteObject(specifitFieldTypeName, "")
			if opt.specificFieldMayNull {
				derived.addFields(&field{name: derivedIdentifierValue, typeCode: jen.Op("*").Id(specifitFieldTypeName)})
			} else {
				derived.addFields(&field{name: derivedIdentifierValue, typeCode: jen.Id(specifitFieldTypeName)})
			}
		}
	}

	return derived
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

// Deprecated:
func (b *builder) addAbstractList(targetName string, name string) *abstractList {
	o := &abstractList{}
	o.name_ = strings.TrimSpace(name)
	o.targetName = strings.TrimSpace(targetName)
	b.localSymbols = append(b.localSymbols, o)
	b.globalSymbols.Store(name, o)
	return o
}

// Deprecated:
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

func getSymbol[T abstractObject | concreteObject | unionObject](b *builder, name string) *T {
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
