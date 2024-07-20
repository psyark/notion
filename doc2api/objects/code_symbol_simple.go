package objects

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

type symbolWithFields interface {
	name() string
	AddFields(fields ...fieldRenderer)
}

var _ symbolWithFields = &SimpleObject{}

// SimpleObject は単純なオブジェクトに使われるGoコードを生成します
type SimpleObject struct {
	nameImpl
	comment           string
	fields            []fieldRenderer
	genericConstraint jen.Code
}

func (o *SimpleObject) isGeneric() bool {
	return o.genericConstraint != nil
}

func (o *SimpleObject) AddComment(comment string) {
	if o.comment != "" {
		o.comment += "\n\n"
	}
	o.comment += strings.TrimSuffix(strings.TrimPrefix(comment, "\n"), "\n")
}

func (o *SimpleObject) AddFields(fields ...fieldRenderer) {
	o.fields = append(o.fields, fields...)
}

func (o *SimpleObject) getDiscriminatorValues(discriminator string) []string {
	for _, f := range o.fields {
		if f, ok := f.(*DiscriminatorField); ok && f.name == discriminator {
			return []string{f.value}
		}
	}
	return nil
}

func (o *SimpleObject) code(c *Converter) jen.Code {
	code := &jen.Statement{}
	if o.comment != "" {
		code.Comment(o.comment).Line()
	}

	code.Type().Add(o.typeCode(true)).StructFunc(func(g *jen.Group) {
		for _, f := range o.fields {
			g.Add(f.renderField())
		}
	}).Line()

	// フィールドにインターフェイスを含むならUnmarshalJSONで前処理を行う
	code.Add(o.fieldUnmarshalerCode(c))

	for _, entry := range c.unionMemberRegistry {
		if entry.member.name() == o.name() {
			code.Func().Params(o.typeCode(false)).Id("is" + entry.union.name()).Params().Block().Line()
		}
	}

	return code
}
func (o *SimpleObject) typeCode(constraint bool) jen.Code {
	code := jen.Id(o.name_)
	if o.genericConstraint != nil {
		code.IndexFunc(func(g *jen.Group) {
			if constraint {
				g.Id("T").Add(o.genericConstraint)
			} else {
				g.Id("T")
			}
		})
	}
	return code
}

func (o *SimpleObject) fieldUnmarshalerCode(c *Converter) jen.Code {
	code := &jen.Statement{}

	unionFields := []*VariableField{}
	for _, f := range o.fields {
		if f, ok := f.(*VariableField); ok && f.getUnionInterface(c) != nil {
			unionFields = append(unionFields, f)
		}
	}

	if len(unionFields) != 0 {
		code.Comment("UnmarshalJSON assigns the appropriate implementation to interface field(s)").Line()
		code.Func().Params(jen.Id("o").Op("*").Id(o.name())).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().BlockFunc(func(g *jen.Group) {
			g.Type().Id("Alias").Id(o.name())
			g.Id("t").Op(":=").Op("&").StructFunc(func(g *jen.Group) {
				g.Op("*").Id("Alias")
				for _, f := range unionFields {
					g.Id(strcase.UpperCamelCase(f.name)).Id(f.getUnionInterface(c).memberUnmarshalerName()).Tag(map[string]string{"json": f.name})
				}
			}).Values(jen.Dict{
				jen.Id("Alias"): jen.Parens(jen.Op("*").Id("Alias")).Call(jen.Id("o")),
			})
			g.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Id("t"))).Op(";").Err().Op("!=").Nil().Block(
				jen.Return().Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("unmarshaling %s: %%w", o.name())), jen.Err()),
			)
			for _, f := range unionFields {
				fieldName := strcase.UpperCamelCase(f.name)
				g.Id("o").Dot(fieldName).Op("=").Id("t").Dot(fieldName).Dot("value")
			}
			g.Return().Nil()
		}).Line()
	}

	return code
}
