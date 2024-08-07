package endpoints

// ssrProps は endpoints ドキュメントページから取得できるJSON構造体です
type ssrProps struct {
	APIBaseURL               string         `json:"apiBaseUrl"`
	BaseURL                  string         `json:"baseUrl"`
	Config                   map[string]any `json:"config"`
	Context                  map[string]any `json:"context"`                  // map[project:map[appearance:map[body:map[style:n...
	Doc                      ssrPropsDoc    `json:"doc"`                      // map[__v:57 _id:609176570b6bf20019821ce5 api:map...
	GlossaryTerms            []any          `json:"glossaryTerms"`            // [map[_id:608fece52e9bb4000fea89a2 definition:TO...
	HideTOC                  bool           `json:"hideTOC"`                  // false
	IsDetachedProductionSite bool           `json:"isDetachedProductionSite"` // false
	Lang                     string         `json:"lang"`                     // en
	LangFull                 string         `json:"langFull"`                 // English
	LoginURL                 string         `json:"loginUrl"`                 // https://dash.readme.com/to/notion-group
	Meta                     map[string]any `json:"meta"`                     // map[hidden:false title:Query a database type:re...
	OasDefinition            map[string]any `json:"oasDefinition"`            // map[_id:606ecc2cd9e93b0044cf6e47:609176570b6bf2...
	OasPublicURL             any            `json:"oasPublicUrl"`             // @notionapi/v1#4rkc1lkz7dw9us
	Oauth                    bool           `json:"oauth"`                    // false
	Rdmd                     map[string]any `json:"rdmd"`
	RdmdOpts                 map[string]any `json:"rdmdOpts"`       // map[compatibilityMode:false correctnewlines:fal...
	ReqURL                   string         `json:"reqUrl"`         // /reference/post-database-query
	Search                   map[string]any `json:"search"`         // map[UrlManager:map[defaults:map[child:<nil> lan...
	Sidebars                 map[string]any `json:"sidebars"`       // map[docs:[map[__v:0 _id:6038057d9c4b200067ba3ca...
	SuggestedEdits           bool           `json:"suggestedEdits"` // true
	Variables                map[string]any `json:"variables"`      // map[defaults:[map[_id:60c02cf4f26fa80064c30a79 ...
	Version                  map[string]any `json:"version"`        // map[__v:2 _id:6038057d9c4b200067ba3c9f categori...
}

type ssrPropsDoc struct {
	API                   ssrPropsAPI    `json:"api"`                   // map[apiSetting:606ecc2cd9e93b0044cf6e47 auth:re...
	Title                 string         `json:"title"`                 // Query a database
	V                     float64        `json:"__v"`                   // 57
	ID                    string         `json:"id,omitempty"`          // 614943b3de71ea001c546257
	ID2                   string         `json:"_id"`                   // 609176570b6bf20019821ce5
	Body                  string         `json:"body"`                  // Gets a list of [Pages](ref:page) contained in t...
	Category              map[string]any `json:"category"`              // 6091386ce2ca9200479fb438
	CreatedAt             string         `json:"createdAt"`             // 2021-05-04T16:29:11.027Z
	Deprecated            bool           `json:"deprecated"`            // false
	Excerpt               string         `json:"excerpt"`               //
	Hidden                bool           `json:"hidden"`                // false
	Icon                  any            `json:"icon,omitempty"`        //
	IsApi                 bool           `json:"isApi,omitempty"`       // true
	IsReference           bool           `json:"isReference"`           // true
	LinkExternal          bool           `json:"link_external"`         // false
	LinkURL               string         `json:"link_url"`              //
	Metadata              map[string]any `json:"metadata"`              // map[description: image:[] title:]
	Next                  map[string]any `json:"next"`                  // map[description: pages:[]]
	Order                 float64        `json:"order"`                 // 1
	ParentDoc             any            `json:"parentDoc"`             // <nil>
	PendingAlgoliaPublish bool           `json:"pendingAlgoliaPublish"` // false
	PreviousSlug          string         `json:"previousSlug"`          // post-databases-query
	Project               string         `json:"project"`               // 6038057d9c4b200067ba3c9a
	Slug                  string         `json:"slug"`                  // post-database-query
	SlugUpdatedAt         string         `json:"slugUpdatedAt"`         // 2021-05-10T00:46:29.470Z
	Swagger               map[string]any `json:"swagger,omitempty"`     // map[path:/v1/databases/{database_id}/query]
	SyncUnique            string         `json:"sync_unique"`           //
	Type                  string         `json:"type"`                  // endpoint
	UpdatedAt             string         `json:"updatedAt"`             // 2021-12-23T16:56:23.254Z
	Updates               []any          `json:"updates"`               // []
	User                  string         `json:"user"`                  // 60917de732252800631fcd43
	Version               any            `json:"version"`               // 6038057d9c4b200067ba3c9f
	Algolia               any            `json:"algolia"`
	ReusableContent       any            `json:"reusableContent"`
	Revision              any            `json:"revision"`
	Tutorials             []any          `json:"tutorials"`
	LastUpdatedHash       string         `json:"lastUpdatedHash"`
}

type ssrPropsAPI struct {
	Method     string          `json:"method"`               // post
	Params     []ssrPropsParam `json:"params"`               // [map[_id:609176570b6bf20019821ce8 default: desc...
	URL        string          `json:"url"`                  // /v1/databases/{database_id}/query
	Results    map[string]any  `json:"results"`              // map[codes:[map[code:{"object": "list","resu...
	APISetting string          `json:"apiSetting,omitempty"` // 606ecc2cd9e93b0044cf6e47
	Auth       string          `json:"auth"`                 // required
	Examples   map[string]any  `json:"examples,omitempty"`   // map[codes:[map[code:const { Client } = require(...
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
	ID         string `json:"_id"`        // 609176570b6bf20019821ce8
}
