package objects

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

// CodeBuilder は
// builderは生成されるオブジェクトやフィールドの依存関係の構築を行います
// 実行時にオブジェクトが登録される順番が一定しなくても、出力されるソースコードは冪等に保たれます
type CodeBuilder struct {
	converter *Converter
	url       string
	fileName  string
	symbols   []CodeSymbol
}

func (b *CodeBuilder) AddConcreteObject(name string, comment string) *ConcreteObject {
	obj := &ConcreteObject{}
	obj.name_ = strings.TrimSpace(name)
	obj.comment = comment
	b.symbols = append(b.symbols, obj)
	b.converter.symbols.Store(name, obj)
	return obj
}

type AddAdaptiveObjectOption func(o *AdaptiveObject)

func DiscriminatorOmitEmpty() AddAdaptiveObjectOption {
	return func(o *AdaptiveObject) {
		for _, f := range o.fields {
			switch f := f.(type) {
			case *VariableField:
				if f.name == o.discriminatorKey {
					f.omitEmpty = true
				}
			}
		}
	}
}

func (b *CodeBuilder) AddAdaptiveObject(name string, discriminatorKey string, comment string, options ...AddAdaptiveObjectOption) *AdaptiveObject {
	o := &AdaptiveObject{}
	o.name_ = name
	o.discriminatorKey = discriminatorKey
	o.comment = comment
	if discriminatorKey != "" {
		o.AddFields(&VariableField{name: discriminatorKey, typeCode: jen.String()})
	}
	for _, option := range options {
		option(o)
	}
	b.symbols = append(b.symbols, o)
	b.converter.symbols.Store(name, o)
	return o
}

// AddUnionToGlobalIfNotExists は、指定された名前と識別キーを持つ Unionを定義し、返します。
// 二回目以降の呼び出しでは定義をスキップし、初回に定義されたものを返します。
// TODO identifierKey -> discriminatorKey
func (b *CodeBuilder) AddUnionToGlobalIfNotExists(name string, identifierKey string) *UnionObject {
	if u := b.converter.getUnionObject(name); u != nil {
		return u
	}

	u := &UnionObject{}
	u.name_ = strings.TrimSpace(name)
	u.identifierKey = identifierKey
	b.converter.globalBuilder.symbols = append(b.converter.globalBuilder.symbols, u)
	b.converter.symbols.Store(name, u)

	return u
}

func (b *CodeBuilder) AddUnmarshalTest(targetName string, jsonCode string) {
	ut := &UnmarshalTest{targetName: targetName} // UnmarshalTestを作る

	if exists := b.converter.getUnmarshalTest(ut.name()); exists != nil { // 同名のものが既にあるなら
		exists.jsonCodes = append(exists.jsonCodes, jsonCode) // JSONコードだけ追加
	} else { // 無ければ追加
		ut.jsonCodes = append(ut.jsonCodes, jsonCode)
		b.converter.symbols.Store(ut.name(), ut)
		b.converter.globalTestBuilder.symbols = append(b.converter.globalTestBuilder.symbols, ut)
	}
}

func (b *CodeBuilder) output(sortSymbols bool) {
	if sortSymbols {
		slices.SortFunc(b.symbols, func(a, b CodeSymbol) int {
			return strings.Compare(a.name(), b.name())
		})
	}

	filePath := fmt.Sprintf("../../%s", b.fileName)

	if len(b.symbols) == 0 {
		_ = os.Remove(filePath)
	} else {
		file := jen.NewFile("notion")
		file.HeaderComment("Code generated by notion.doc2api; DO NOT EDIT.")
		if b.url != "" {
			file.HeaderComment(b.url)
		}
		for _, s := range b.symbols {
			file.Line().Line().Add(s.code(b.converter))
		}
		lo.Must0(file.Save(filePath))
	}
}

// asField は、このドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (*CodeBuilder) NewField(p *Parameter, typeCode jen.Code, options ...fieldOption) *VariableField {
	f := &VariableField{
		name:     p.Property,
		typeCode: typeCode,
		comment:  p.Description,
	}
	for _, o := range options {
		o(f)
	}
	return f
}

// NewDiscriminatorField は、ドキュメントに書かれたパラメータを、渡されたタイプに従ってGoコードのフィールドに変換します
func (b *CodeBuilder) NewDiscriminatorField(p *Parameter) *discriminatorField {
	for _, value := range []string{p.ExampleValue, p.Type} {
		if value != "" {
			if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
				value = strings.TrimPrefix(strings.TrimSuffix(value, `"`), `"`)

				{
					as := DiscriminatorString(value)
					if _, ok := b.converter.symbols.Load(as.name()); !ok {
						b.converter.symbols.Store(as.name(), as)
						b.converter.globalBuilder.symbols = append(b.converter.globalBuilder.symbols, as)
					}
				}

				return &discriminatorField{
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

func DiscriminatorValue(value string) fieldOption {
	return func(f *VariableField) {
		f.discriminatorValue = value
	}
}
