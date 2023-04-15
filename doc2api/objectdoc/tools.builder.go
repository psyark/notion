package objectdoc

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// builderは生成されるオブジェクトやフィールドの依存関係の構築を行います
// 実行時にオブジェクトが登録される順番が一定しなくても、出力されるソースコードは冪等に保たれます
type builder struct {
	fileName string
	url      string
	coders   []symbolCoder

	// Deprecated:
	global *builder // TODO ローカル/グローバルビルダーを作るのではなく、唯一のビルダーを作る
}

func (b *builder) addSpecificObject(name string, comment string) *specificObject {
	o := &specificObject{}
	o.name_ = strings.TrimSpace(name)
	o.comment = comment
	b.coders = append(b.coders, o)
	return o
}

func (b *builder) addAbstractObject(name string, specifiedBy string, comment string) *abstractObject {
	o := &abstractObject{}
	o.name_ = strings.TrimSpace(name)
	o.specifiedBy = specifiedBy
	o.comment = comment
	b.coders = append(b.coders, o)
	return o
}

func (b *builder) addAbstractObjectToGlobalIfNotExists(name string, specifiedBy string) *abstractObject {
	if o := b.global.getAbstractObject(name); o != nil {
		return o
	}
	return b.global.addAbstractObject(name, specifiedBy, "")
}

func (b *builder) addAlwaysStringIfNotExists(value string) {
	for _, c := range b.coders {
		if c, ok := c.(alwaysString); ok && string(c) == value {
			return
		}
	}
	b.coders = append(b.coders, alwaysString(value))
}

func (b *builder) statement() jen.Statement {
	s := jen.Statement{}
	for _, item := range b.coders {
		s = append(s, jen.Line().Line(), item.symbolCode())
	}
	return s
}

func (b *builder) getAbstractObject(name string) *abstractObject {
	for _, item := range b.coders {
		if item, ok := item.(*abstractObject); ok && item.name() == name {
			return item
		}
	}
	if b.global != nil {
		return b.global.getAbstractObject(name)
	}
	return nil
}

func (b *builder) getSpecificObject(name string) *specificObject {
	for _, item := range b.coders {
		if item, ok := item.(*specificObject); ok && item.name() == name {
			return item
		}
	}
	if b.global != nil {
		return b.global.getSpecificObject(name)
	}
	return nil
}

type alwaysString string

func (c alwaysString) symbolCode() jen.Code {
	code := jen.Type().Id(c.name()).String().Line()
	code.Func().Params(jen.Id("s").Id(c.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().List(jen.Index().Byte().Call(jen.Lit(fmt.Sprintf("%q", string(c)))), jen.Nil()),
	)
	return code
}

func (c alwaysString) name() string {
	return "always" + strcase.UpperCamelCase(string(c))
}
