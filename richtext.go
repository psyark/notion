package notion

import (
	"strings"

	"github.com/samber/lo"
)

type rtOption func(rt *RichText)

func NewRichText(text string, options ...rtOption) RichText {
	rt := RichText{Text: &RichTextText{Content: text}, PlainText: text}
	for _, option := range options {
		option(&rt)
	}
	return rt
}

func NewRichTextArray(text string, options ...rtOption) RichTextArray {
	return RichTextArray{NewRichText(text, options...)}
}

func (r RichText) String() string {
	return r.PlainText
}

type RichTextArray []RichText

func (r RichTextArray) String() string {
	return strings.Join(lo.Map(r, func(r RichText, _ int) string { return r.String() }), "")
}

func WithAnnotation(annot Annotations) rtOption {
	return func(rta *RichText) {
		rta.Annotations = annot
	}
}
