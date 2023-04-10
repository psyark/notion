package methoddoc

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type ssrProps struct {
	Doc ssrPropsDoc `json:"doc"` // map[__v:57 _id:609176570b6bf20019821ce5 api:map...
	// BaseURL                  string         `json:"baseUrl"`
	// APIBaseURL               string         `json:"apiBaseUrl"`
	// Config                   any            `json:"config"`
	// Context                  map[string]any `json:"context"`                  // map[project:map[appearance:map[body:map[style:n...
	// GlossaryTerms            []any          `json:"glossaryTerms"`            // [map[_id:608fece52e9bb4000fea89a2 definition:TO...
	// TideTOC                  bool           `json:"hideTOC"`                  // false
	// IsDetachedProductionSite bool           `json:"isDetachedProductionSite"` // false
	// Lang                     string         `json:"lang"`                     // en
	// LangFull                 string         `json:"langFull"`                 // English
	// LoginURL                 string         `json:"loginUrl"`                 // https://dash.readme.com/to/notion-group
	// Meta                     map[string]any `json:"meta"`                     // map[hidden:false title:Query a database type:re...
	// OasDefinition            map[string]any `json:"oasDefinition"`            // map[_id:606ecc2cd9e93b0044cf6e47:609176570b6bf2...
	// OasPublicURL             any            `json:"oasPublicUrl"`             // @notionapi/v1#4rkc1lkz7dw9us
	// Oauth                    bool           `json:"oauth"`                    // false
	// RdmdOpts                 map[string]any `json:"rdmdOpts"`                 // map[compatibilityMode:false correctnewlines:fal...
	// ReqURL                   string         `json:"reqUrl"`                   // /reference/post-database-query
	// Search                   map[string]any `json:"search"`                   // map[UrlManager:map[defaults:map[child:<nil> lan...
	// Sidebars                 any            `json:"sidebars"`                 // map[docs:[map[__v:0 _id:6038057d9c4b200067ba3ca...
	// SuggestedEdits           bool           `json:"suggestedEdits"`           // true
	// Variables                map[string]any `json:"variables"`                // map[defaults:[map[_id:60c02cf4f26fa80064c30a79 ...
	// Version                  map[string]any `json:"version"`                  // map[__v:2 _id:6038057d9c4b200067ba3c9f categori...
}

type ssrPropsDoc struct {
	API   ssrPropsAPI `json:"api"`   // map[apiSetting:606ecc2cd9e93b0044cf6e47 auth:re...
	Title string      `json:"title"` // Query a database
	// V                     float64        `json:"__v"`                   // 57
	// ID                    string         `json:"id,omitempty"`          // 614943b3de71ea001c546257
	// ID2                   string         `json:"_id"`                   // 609176570b6bf20019821ce5
	// Body                  string         `json:"body"`                  // Gets a list of [Pages](ref:page) contained in t...
	// Category              string         `json:"category"`              // 6091386ce2ca9200479fb438
	// Children              []any          `json:"children"`              // [map[__v:0 _id:6098885974ae4300418f9a18 api:map...
	// ChildrenPages         []any          `json:"childrenPages"`         // [map[__v:0 _id:6098885974ae4300418f9a18 api:map...
	// CreatedAt             string         `json:"createdAt"`             // 2021-05-04T16:29:11.027Z
	// Deprecated            bool           `json:"deprecated"`            // false
	// Excerpt               string         `json:"excerpt"`               //
	// Hidden                bool           `json:"hidden"`                // false
	// Icon                  any            `json:"icon,omitempty"`        //
	// IsApi                 bool           `json:"isApi,omitempty"`       // true
	// IsReference           bool           `json:"isReference"`           // true
	// LinkExternal          bool           `json:"link_external"`         // false
	// LinkURL               string         `json:"link_url"`              //
	// Metadata              map[string]any `json:"metadata"`              // map[description: image:[] title:]
	// Next                  map[string]any `json:"next"`                  // map[description: pages:[]]
	// Order                 float64        `json:"order"`                 // 1
	// ParentDoc             any            `json:"parentDoc"`             // <nil>
	// PendingAlgoliaPublish bool           `json:"pendingAlgoliaPublish"` // false
	// PreviousSlug          string         `json:"previousSlug"`          // post-databases-query
	// Project               string         `json:"project"`               // 6038057d9c4b200067ba3c9a
	// Slug                  string         `json:"slug"`                  // post-database-query
	// SlugUpdatedAt         string         `json:"slugUpdatedAt"`         // 2021-05-10T00:46:29.470Z
	// Swagger               map[string]any `json:"swagger,omitempty"`     // map[path:/v1/databases/{database_id}/query]
	// SyncUnique            string         `json:"sync_unique"`           //
	// Type                  string         `json:"type"`                  // endpoint
	// UpdatedAt             string         `json:"updatedAt"`             // 2021-12-23T16:56:23.254Z
	// Updates               []any          `json:"updates"`               // []
	// User                  string         `json:"user"`                  // 60917de732252800631fcd43
	// Version               any            `json:"version"`               // 6038057d9c4b200067ba3c9f
}

type ssrPropsAPI struct {
	Method string          `json:"method"` // post
	Params []ssrPropsParam `json:"params"` // [map[_id:609176570b6bf20019821ce8 default: desc...
	URL    string          `json:"url"`    // /v1/databases/{database_id}/query
	// Results    map[string]any  `json:"results"`              // map[codes:[map[code:{"object": "list","resu...
	// APISetting string          `json:"apiSetting,omitempty"` // 606ecc2cd9e93b0044cf6e47
	// Auth       string          `json:"auth"`                 // required
	// Examples   map[string]any  `json:"examples,omitempty"`   // map[codes:[map[code:const { Client } = require(...
}

func (api *ssrPropsAPI) filterParams(in string) []ssrPropsParam {
	params := []ssrPropsParam{}
	for _, param := range api.Params {
		if param.In == in {
			params = append(params, param)
		}
	}
	return params
}

type ssrPropsParam struct {
	Default    string `json:"default"`    //
	Desc       string `json:"desc"`       // When supplied, limits which pages are returned ...
	EnumValues string `json:"enumValues"` //
	In         string `json:"in"`         // body
	Name       string `json:"name"`       // filter
	Ref        string `json:"ref"`        //
	Required   bool   `json:"required"`   // false
	Type       string `json:"type"`       // json
	typeCode   jen.Code
	// ID         string `json:"_id"`        // 609176570b6bf20019821ce8
}

func (p ssrPropsParam) compare(p2 ssrPropsParam) error {
	s1 := &jen.Statement{p.Code()}
	s2 := &jen.Statement{p2.Code()}
	if s1.GoString() != s2.GoString() {
		return fmt.Errorf("mismatch: \n%#v\n%#v", s1, s2)
	}
	return nil
}

func (p ssrPropsParam) Code() jen.Code {
	dict := jen.Dict{}
	if p.Default != "" {
		dict[jen.Id("Default")] = jen.Lit(p.Default)
	}
	if p.Desc != "" {
		dict[jen.Id("Desc")] = jen.Lit(p.Desc)
	}
	if p.EnumValues != "" {
		dict[jen.Id("EnumValues")] = jen.Lit(p.EnumValues)
	}
	if p.In != "" {
		dict[jen.Id("In")] = jen.Lit(p.In)
	}
	if p.Name != "" {
		dict[jen.Id("Name")] = jen.Lit(p.Name)
	}
	if p.Ref != "" {
		dict[jen.Id("Ref")] = jen.Lit(p.Ref)
	}
	if p.Type != "" {
		dict[jen.Id("Type")] = jen.Lit(p.Type)
	}
	if p.Required {
		dict[jen.Id("Required")] = jen.True()
	}
	return jen.Id("ssrPropsParam").Values(dict)
}
