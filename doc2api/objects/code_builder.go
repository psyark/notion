package objects

import (
	"fmt"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

// CodeBuilder は
// builderは生成されるオブジェクトやフィールドの依存関係の構築を行います
// 実行時にオブジェクトが登録される順番が一定しなくても、出力されるソースコードは冪等に保たれます
type CodeBuilder struct {
	converter    *Converter
	url          string
	fileName     string
	localSymbols []CodeSymbol
	testSymbols  []CodeSymbol
}

func (b *CodeBuilder) AddConcreteObject(name string, comment string) *ConcreteObject {
	obj := &ConcreteObject{}
	obj.name_ = strings.TrimSpace(name)
	obj.comment = comment
	b.localSymbols = append(b.localSymbols, obj)
	b.converter.symbols.Store(name, obj)
	return obj
}

func (b *CodeBuilder) addAdaptiveObject(name string, discriminatorKey string, comment string) *AdaptiveObject {
	ao := &AdaptiveObject{}
	ao.name_ = name
	ao.discriminatorKey = discriminatorKey
	ao.comment = comment
	if discriminatorKey != "" {
		ao.addFields(&VariableField{
			name:     discriminatorKey,
			typeCode: jen.String(),
		})
	}
	b.localSymbols = append(b.localSymbols, ao)
	b.converter.symbols.Store(name, ao)
	return ao
}

// AddUnionToGlobalIfNotExists は、指定された名前と識別キーを持つ Unionを定義し、返します。
// 二回目以降の呼び出しでは定義をスキップし、初回に定義されたものを返します。
// TODO identifierKey -> discriminatorKey
func (b *CodeBuilder) AddUnionToGlobalIfNotExists(name string, identifierKey string) *UnionObject {
	if u := b.GetUnionObject(name); u != nil {
		return u
	}

	u := &UnionObject{}
	u.name_ = strings.TrimSpace(name)
	u.identifierKey = identifierKey
	b.converter.globalBuilder.localSymbols = append(b.converter.globalBuilder.localSymbols, u)
	b.converter.symbols.Store(name, u)

	return u
}

func (b *CodeBuilder) addAlwaysStringIfNotExists(value string) {
	as := AlwaysString(value)
	if _, ok := b.converter.symbols.Load(as.name()); ok {
		return
	}
	b.converter.symbols.Store(as.name(), as)
	b.localSymbols = append(b.localSymbols, as)
}

func getSymbol[T CodeSymbol](name string, b *CodeBuilder) T {
	if item, ok := b.converter.symbols.Load(name); ok {
		if item, ok := item.(T); ok {
			return item
		}
	}

	var zero T
	return zero
}

func (b *CodeBuilder) GetConcreteObject(name string) *ConcreteObject {
	return getSymbol[*ConcreteObject](name, b)
}
func (b *CodeBuilder) GetAdaptiveObject(name string) *AdaptiveObject {
	return getSymbol[*AdaptiveObject](name, b)
}
func (b *CodeBuilder) GetUnionObject(name string) *UnionObject {
	return getSymbol[*UnionObject](name, b)
}
func (b *CodeBuilder) GetUnmarshalTest(name string) *UnmarshalTest {
	return getSymbol[*UnmarshalTest](name, b)
}

func (b *CodeBuilder) AddUnmarshalTest(targetName string, jsonCode string) {
	ut := &UnmarshalTest{targetName: targetName} // UnmarshalTestを作る

	if exists := b.GetUnmarshalTest(ut.name()); exists != nil { // 同名のものが既にあるなら
		exists.jsonCodes = append(exists.jsonCodes, jsonCode) // JSONコードだけ追加
	} else { // 無ければ追加
		ut.jsonCodes = append(ut.jsonCodes, jsonCode)
		b.converter.symbols.Store(ut.name(), ut)
		b.testSymbols = append(b.testSymbols, ut)
	}
}

func (b *CodeBuilder) output() {
	filePath := fmt.Sprintf("../../%s", b.fileName)

	if len(b.localSymbols) == 0 {
		_ = os.Remove(filePath)
	} else {
		file := jen.NewFile("notion")
		file.HeaderComment("Code generated by notion.doc2api; DO NOT EDIT.")
		if b.url != "" {
			file.Comment(b.url)
		}
		for _, s := range b.localSymbols {
			file.Line().Line().Add(s.code())
		}
		lo.Must0(file.Save(filePath))
	}

	// テストコード出力
	if len(b.testSymbols) == 0 {
		_ = os.Remove(strings.Replace(filePath, ".go", "_test.go", 1))
	} else {
		file := jen.NewFile("notion")
		file.HeaderComment("Code generated by notion.doc2api; DO NOT EDIT.")
		if b.url != "" {
			file.Comment(b.url)
		}
		for _, s := range b.testSymbols {
			file.Line().Line().Add(s.code())
		}
		lo.Must0(file.Save(strings.Replace(filePath, ".go", "_test.go", 1)))
	}
}

// asField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (b *CodeBuilder) NewField(p *Parameter, typeCode jen.Code, options ...fieldOption) *VariableField {
	f := &VariableField{
		name:     p.Property,
		typeCode: typeCode,
		comment:  p.Description,
		builder:  b,
	}
	for _, o := range options {
		o(f)
	}
	return f
}

// NewFieldStringField は、ドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (b *CodeBuilder) NewFieldStringField(p *Parameter) *fixedStringField {
	for _, value := range []string{p.ExampleValue, p.Type} {
		if value != "" {
			if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
				value = strings.TrimPrefix(strings.TrimSuffix(value, `"`), `"`)

				// TODO
				// b.global.addAlwaysStringIfNotExists(value)

				return &fixedStringField{
					name:    p.Property,
					value:   value,
					comment: p.Description,
				}
			}
			panic(value)
		}
	}
	panic(fmt.Errorf("パラメータを fixedStringField に変換できません: %v", p))
}

type fieldOption func(f *VariableField)

var OmitEmpty fieldOption = func(f *VariableField) {
	f.omitEmpty = true
}

var DiscriminatorNotEmpty fieldOption = func(f *VariableField) {
	f.discriminatorNotEmpty = true
}
