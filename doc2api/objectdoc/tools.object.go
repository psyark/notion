package objectdoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

// symbolCoder はソースコードのトップレベルに置かれる、名前を持つシンボルの生成器です。
type symbolCoder interface {
	symbolCode(*builder) jen.Code
	name() string
}

var _ = []symbolCoder{
	&objectCommon{},
	&specificObject{},
	&abstractObject{},
	&abstractList{},
	&abstractMap{},
	&unmarshalTest{},
	alwaysString(""),
}

// objectCoder はオブジェクトを作成するためのCoderです
type objectCoder interface {
	symbolCoder
	getSpecifyingField(specifiedBy string) *fixedStringField
	addParent(*abstractObject)
}

var _ = []objectCoder{
	&specificObject{},
	&abstractObject{},
}

type objectCommon struct {
	name_   string
	comment string
	fields  []fieldCoder
	parents []*abstractObject
}

func (c *objectCommon) name() string {
	return c.name_
}

// TODO: これを廃止してvariantTypeみたいなプロパティ作る
func (c *objectCommon) getSpecifyingField(specifiedBy string) *fixedStringField {
	for _, f := range c.fields {
		if f, ok := f.(*fixedStringField); ok && f.name == specifiedBy {
			return f
		}
	}
	for _, p := range c.parents {
		if f := p.getSpecifyingField(specifiedBy); f != nil {
			return f
		}
	}
	return nil
}

func (c *objectCommon) addParent(parent *abstractObject) {
	c.parents = append(c.parents, parent)
}

func (c *objectCommon) getAncestors() []*abstractObject {
	ancestors := []*abstractObject{}
	for _, a := range c.parents {
		ancestors = append(ancestors, a)
		ancestors = append(ancestors, a.getAncestors()...)
	}
	return ancestors
}

func (c *objectCommon) fieldCodes() []jen.Code {
	fields := []jen.Code{}
	for _, p := range c.parents {
		if p.hasCommonField() {
			fields = append(fields, jen.Id(p.commonObjectName()))
		}
	}
	for _, f := range c.fields {
		fields = append(fields, f.fieldCode())
	}
	return fields
}

func (c *objectCommon) addFields(fields ...fieldCoder) *objectCommon {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *objectCommon) symbolCode(b *builder) jen.Code {
	code := &jen.Statement{}
	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	return code.Type().Id(c.name_).Struct(c.fieldCodes()...).Line()
}

func (c *objectCommon) fieldUnmarshalerCode() jen.Code {
	code := &jen.Statement{}

	interfaceFields := []*interfaceField{}
	for _, f := range c.fields {
		if f, ok := f.(*interfaceField); ok {
			interfaceFields = append(interfaceFields, f)
		}
	}

	if len(interfaceFields) != 0 {
		code.Comment("UnmarshalJSON assigns the appropriate implementation to interface field(s)").Line()
		code.Func().Params(jen.Id("o").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().BlockFunc(func(g *jen.Group) {
			g.Type().Id("Alias").Id(c.name())
			g.Id("t").Op(":=").Op("&").StructFunc(func(g *jen.Group) {
				g.Op("*").Id("Alias")
				for _, f := range interfaceFields {
					g.Id(strcase.UpperCamelCase(f.name)).Id(strcase.LowerCamelCase(f.typeName) + "Unmarshaler").Tag(map[string]string{"json": f.name})
				}
			}).Values(jen.Dict{
				jen.Id("Alias"): jen.Parens(jen.Op("*").Id("Alias")).Call(jen.Id("o")),
			})
			g.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
			)
			for _, f := range interfaceFields {
				fieldName := strcase.UpperCamelCase(f.name)
				g.Id("o").Dot(fieldName).Op("=").Id("t").Dot(fieldName).Dot("value")
			}
			g.Return().Nil()
		}).Line()
	}

	return code
}

// specificObject は"type"や"object"キーで区別される各オブジェクトです
// （例：type="external" である ExternalFile）
// 生成されるGoコードではstructポインタで表現されます
type specificObject struct {
	objectCommon

	// typeObject はこのspecificObjectが そのtype値と同名のフィールドに保持する固有データです
	// Every block object has a key corresponding to the value of type. Under the key is an object with type-specific block information.
	// TODO typeObjectがAbstractだった場合の対応（TemplateMentionData）
	// TODO fixedStringFieldを持たないオブジェクトがtypeObjectを使いたい場合（例：Filter）
	typeObject        objectCommon
	typeObjectMayNull bool
}

func (c *specificObject) addFields(fields ...fieldCoder) *specificObject {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *specificObject) symbolCode(b *builder) jen.Code {
	// typeObjectが使われているならtypeObjectへの参照を追加する
	if len(c.typeObject.fields) != 0 {
		typeField := c.getSpecifyingField("type")
		if typeField != nil {
			var valueOfTypeField *field
			for _, f := range c.fields {
				if f, ok := f.(*field); ok && f.name == typeField.value {
					valueOfTypeField = f
					break
				}
			}
			if valueOfTypeField == nil {
				if c.typeObjectMayNull {
					c.addFields(&field{name: typeField.value, typeCode: jen.Op("*").Id(c.name() + "Data")})
				} else {
					c.addFields(&field{name: typeField.value, typeCode: jen.Id(c.name() + "Data")})
				}
			}
		} else {
			panic(fmt.Sprintf("タイプが不明です: %v", c.name()))
		}
	}

	// struct本体
	code := &jen.Statement{}
	code.Add(c.objectCommon.symbolCode(b))

	// インターフェイスを実装
	for _, a := range c.getAncestors() {
		code.Func().Params(jen.Id("_").Op("*").Id(c.name())).Id("is" + a.name()).Params().Block().Line()
	}
	// 親のスペシャルメソッドを実装
	for _, p := range c.parents {
		for _, sm := range p.specialMethods {
			code.Add(sm.implementationCode(c))
		}
	}

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(c.fieldUnmarshalerCode())

	// type object
	if len(c.typeObject.fields) != 0 {
		c.typeObject.name_ = c.name() + "Data"
		code.Add(c.typeObject.symbolCode(b))
	}

	return code
}

// specialMethodCoder はabstractObjectに特別なメソッドを追加したい場合に使います
type specialMethodCoder interface {
	declarationCode() jen.Code
	implementationCode(*specificObject) jen.Code
}

// abstractObject は、共通の性質を持つspecificObject又はabstractObjectの総称です
// （例：File）
// 生成されるGoコードではinterface（共通するフィールドがある場合はstructも）で表現されます
type abstractObject struct {
	objectCommon
	specifiedBy    string // "type", "object" など、バリアントを識別するためのプロパティ // TODO variantTypeKeyとかにする？
	fieldsComment  string
	variants       []objectCoder
	specialMethods []specialMethodCoder
}

// addVariant は指定したobjectCoderをこのインターフェイスのバリアントとして登録し、code()に以下のことを行わせます
// - バリアントに対してインターフェイスメソッドを実装
// - JSONメッセージからこのインターフェイスの適切なバリアントを作成するUnmarshalerを作成
func (c *abstractObject) addVariant(variant objectCoder) *abstractObject {
	variant.addParent(c)
	c.variants = append(c.variants, variant)
	return c
}

func (c *abstractObject) addFields(fields ...fieldCoder) *abstractObject {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *abstractObject) hasCommonField() bool {
	if len(c.fields) != 0 {
		return true
	}
	for _, p := range c.parents {
		if p.hasCommonField() {
			return true
		}
	}
	return false
}

func (c *abstractObject) symbolCode(b *builder) jen.Code {
	code := jen.Line()

	if c.comment != "" {
		code.Comment(c.comment).Line()
	}

	// インターフェイス本体とisメソッド
	{
		methods := []jen.Code{}
		for _, p := range c.parents {
			methods = append(methods, jen.Id(p.name())) // 親インターフェイスの継承
		}
		methods = append(methods, jen.Id("is"+c.name()).Params()) // このインターフェイスのisメソッド
		for _, sm := range c.specialMethods {
			methods = append(methods, sm.declarationCode()) // 特殊メソッド
		}
		// methods = append(methods, jen.Id("Get"+strcase.UpperCamelCase(c.specifiedBy)).Params().String())
		// 共通フィールドのgetter宣言
		for _, f := range c.fields {
			methods = append(methods, jen.Id("Get"+f.goName()).Params().Add(f.getTypeCode()))
		}
		code.Type().Id(c.name()).Interface(methods...).Line()
	}

	// 共通フィールド
	if len(c.fieldCodes()) != 0 {
		copyOfC := *c
		copyOfC.name_ = c.commonObjectName()
		copyOfC.comment = c.fieldsComment
		code.Add(copyOfC.objectCommon.symbolCode(b))
		// 共通フィールドのgetter定義
		for _, f := range c.fields {
			code.Line().Func().Params(jen.Id("c").Op("*").Id(copyOfC.name_)).Id("Get" + f.goName()).Params().Add(f.getTypeCode()).Block(
				jen.Return().Id("c").Dot(f.goName()),
			)
		}
	}

	// variant Unmarshaler
	if c.specifiedBy != "" {
		code.Add(c.variantUnmarshaler())
	}

	return code
}

func (c *abstractObject) variantUnmarshaler() jen.Code {
	code := &jen.Statement{}
	cases := []jen.Code{}

	for _, variant := range c.variants {
		caseCode := jen.Case(jen.Lit(""))
		if sf := variant.getSpecifyingField(c.specifiedBy); sf != nil {
			caseCode = jen.Case(jen.Lit(sf.value))
		}

		switch variant := variant.(type) {
		case *specificObject:
			cases = append(cases, caseCode.Id("u").Dot("value").Op("=").Op("&").Id(variant.name()).Values())
		case *abstractObject:
			cases = append(cases,
				caseCode.Id("t").Op(":=").Op("&").Id(strcase.LowerCamelCase(variant.name())+"Unmarshaler").Values(),
				jen.If(jen.Err().Op(":=").Id("t").Dot("UnmarshalJSON").Call(jen.Id("data"))).Op(";").Err().Op("!=").Nil().Block(jen.Return().Err()),
				jen.Id("u").Dot("value").Op("=").Id("t").Dot("value"),
				jen.Return().Nil(),
			)
		default:
			panic(fmt.Sprintf("%#v", variant))
		}
	}

	cases = append(cases, jen.Default().Return(jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: data has unknown %s field: %%s", c.name(), c.specifiedBy)), jen.String().Call(jen.Id("data")))))

	if len(c.variants) != 0 {
		name := strcase.LowerCamelCase(c.name()) + "Unmarshaler"
		code.Line().Type().Id(name).Struct(
			jen.Id("value").Id(c.name()),
		)
		code.Line().Comment(fmt.Sprintf("UnmarshalJSON unmarshals a JSON message and sets the value field to the appropriate instance\naccording to the %q field of the message.", c.specifiedBy))
		code.Line().Func().Params(jen.Id("u").Op("*").Id(name)).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
			jen.If(jen.String().Call(jen.Id("data"))).Op("==").Lit("null").Block(
				jen.Id("u").Dot("value").Op("=").Nil(),
				jen.Return().Nil(),
			),
			jen.Switch().Id("get"+strcase.UpperCamelCase(c.specifiedBy)).Call(jen.Id("data")).Block(cases...),
			jen.Return().Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("u").Dot("value")),
		).Line()
		code.Line().Func().Params(jen.Id("u").Op("*").Id(name)).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
			jen.Return().Qual("encoding/json", "Marshal").Call(jen.Id("u").Dot("value")),
		).Line()
	}

	return code
}

func (c *abstractObject) commonObjectName() string {
	return c.name() + "Common"
}

type abstractList struct {
	name_      string // Deprecated: TODO 消す
	targetName string
}

func (c *abstractList) name() string { return c.name_ }
func (c *abstractList) symbolCode(b *builder) jen.Code {
	return jen.Line().Type().Id(c.name()).Index().Id(c.targetName).Line().Func().Params(jen.Id("a").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.Id("t").Op(":=").Index().Id(strcase.LowerCamelCase(c.targetName)+"Unmarshaler").Values(),
		jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("t")).Op(";").Err().Op("!=").Nil()).Block(
			jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
		),
		jen.Op("*").Id("a").Op("=").Make(jen.Index().Id(c.targetName), jen.Len(jen.Id("t"))),
		jen.For(jen.List(jen.Id("i"), jen.Id("u")).Op(":=").Range().Id("t")).Block(
			jen.Parens(jen.Op("*").Id("a")).Index(jen.Id("i")).Op("=").Id("u").Dot("value"),
		),
		jen.Return().Nil(),
	)
}

type abstractMap struct {
	name_      string // Deprecated: TODO 消す
	targetName string
}

func (c *abstractMap) name() string { return c.name_ }
func (c *abstractMap) symbolCode(b *builder) jen.Code {
	return jen.Line().Type().Id(c.name()).Map(jen.String()).Id(c.targetName).Line().Func().Params(jen.Id("m").Op("*").Id(c.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		jen.Id("t").Op(":=").Map(jen.String()).Id(strcase.LowerCamelCase(c.targetName)+"Unmarshaler").Values(),
		jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("t")).Op(";").Err().Op("!=").Nil()).Block(
			jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", c.name())), jen.Err()),
		),
		jen.Op("*").Id("m").Op("=").Id(c.name()).Values(),
		jen.For(jen.List(jen.Id("k"), jen.Id("u")).Op(":=").Range().Id("t")).Block(
			jen.Parens(jen.Op("*").Id("m")).Index(jen.Id("k")).Op("=").Id("u").Dot("value"),
		),
		jen.Return().Nil(),
	)
}

type unmarshalTest struct {
	targetName string
	jsonCodes  []string
}

func (c *unmarshalTest) name() string {
	return fmt.Sprintf("Test%s_unmarshal", c.targetName)
}
func (c *unmarshalTest) symbolCode(b *builder) jen.Code {
	var initCode, referCode jen.Code
	switch symbol := b.getSymbol(c.targetName).(type) {
	case *abstractObject:
		initCode = jen.Id("target").Op(":=").Op("&").Id(strcase.LowerCamelCase(c.targetName) + "Unmarshaler").Values()
		referCode = jen.Id("target").Dot("value")
	case *specificObject, *abstractMap:
		initCode = jen.Id("target").Op(":=").Op("&").Id(c.targetName).Values()
		referCode = jen.Id("target")
	default:
		panic(symbol)
	}

	return jen.Line().Func().Id(c.name()).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		initCode,
		jen.Id("tests").Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
			for _, t := range c.jsonCodes {
				g.Line().Lit(t)
			}
			g.Line()
		}),
		jen.For().List(jen.Id("_"), jen.Id("wantStr")).Op(":=").Range().Id("tests").Block(
			jen.Id("want").Op(":=").Index().Byte().Call(jen.Id("wantStr")),
			jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("want"), jen.Id("target"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Id("t").Dot("Fatal").Call(jen.Err()),
			),
			jen.List(jen.Id("got"), jen.Id("_")).Op(":=").Qual("encoding/json", "Marshal").Call(referCode),
			jen.If(jen.List(jen.Id("want"), jen.Id("got"), jen.Id("ok")).Op(":=").Id("compareJSON").Call(jen.Id("want"), jen.Id("got")).Op(";").Op("!").Id("ok")).Block(
				jen.Id("t").Dot("Fatal").Call(jen.Qual("fmt", "Errorf").Call(jen.Lit("mismatch:\nwant: %s\ngot : %s"), jen.Id("want"), jen.Id("got"))),
			),
		),
	)
}

type alwaysString string

func (c alwaysString) symbolCode(b *builder) jen.Code {
	code := jen.Type().Id(c.name()).String().Line()
	code.Func().Params(jen.Id("s").Id(c.name())).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return().List(jen.Index().Byte().Call(jen.Lit(fmt.Sprintf("%q", string(c)))), jen.Nil()),
	)
	return code
}

func (c alwaysString) name() string {
	return "always" + strcase.UpperCamelCase(string(c))
}
