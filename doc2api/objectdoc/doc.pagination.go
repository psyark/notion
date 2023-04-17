package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

func init() {
	registerConverter(converter{
		url: "https://developers.notion.com/reference/intro",
		localCopy: []objectDocElement{
			&objectDocParagraphElement{
				Text:   "The reference is your key to a comprehensive understanding of the Notion API.\n\nIntegrations use the API to access Notion's pages, databases, and users. Integrations can connect services to Notion and build interactive experiences for users within Notion. Using the navigation on the left, you'll find details for objects and endpoints used in the API.",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocCalloutElement{
				Body:   "You need an integration token to interact with the Notion API. You can find an integration token after you create an integration on the integration settings page. If this is your first look at the Notion API, we recommend beginning with the [Getting started guide](doc:getting-started) to learn how to create an integration. \n\nIf you want to work on a specific integration, but can't access the token, confirm that you are an admin in the associated workspace. You can check inside the Notion UI via `Settings & Members` in the left sidebar. If you're not an admin in any of your workspaces, you can create a personal workspace for free.",
				Title:  "",
				Type:   "info",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Conventions",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\nThe base URL to send all API requests is https://api.notion.com. HTTPS is required for all API requests.\n\nThe Notion API follows RESTful conventions when possible, with most operations performed via GET, POST, PATCH, and DELETE requests on page and database resources. Request and response bodies are encoded as JSON.\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "JSON conventions",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "* Top-level resources have an \"object\" property. This property can be used to determine the type of the resource (e.g. \"database\", \"user\", etc.)\n* Top-level resources are addressable by a UUIDv4 \"id\" property. You may omit dashes from the ID when making requests to the API, e.g. when copying the ID from a Notion URL.\n* Property names are in snake_case (not camelCase or kebab-case).\n* Temporal values (dates and datetimes) are encoded in ISO 8601 strings. Datetimes will include the time value (2020-08-12T02:12:33.231Z) while dates will include only the date (2020-08-12)\n* The Notion API does not support empty strings. To unset a string value for properties like a url Property value object, for example, use an explicit null instead of \"\".\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Code samples & SDKs",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\nSamples requests and responses are shown for each endpoint. Requests are shown using the Notion JavaScript SDK, and cURL. These samples make it easy to copy, paste, and modify as you build your integration. \n\nNotion SDKs are open source projects that you can install to easily start building. You may also choose any other language or library that allows you to make HTTP requests.\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text: "Pagination ",
				output: func(e *objectDocHeadingElement, b *builder) error {
					b.addAbstractObjectToGlobalIfNotExists("Pagination", "type")
					return nil
				},
			},
			&objectDocParagraphElement{
				Text: "\nEndpoints that return lists of objects support cursor-based pagination requests. By default, Notion returns ten items per API call. If the number of items in a response from a support endpoint exceeds the default, then an integration can use pagination to request a specific set of the results and/or to limit the number of returned items.\n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("Pagination").comment += e.Text
					return nil
				},
			},
			&objectDocHeadingElement{
				Text:   "Supported endpoints",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParametersElement{
				{output: func(e *objectDocParameter, b *builder) error { return nil }},
				{output: func(e *objectDocParameter, b *builder) error { return nil }},
				{output: func(e *objectDocParameter, b *builder) error { return nil }},
				{output: func(e *objectDocParameter, b *builder) error { return nil }},
				{output: func(e *objectDocParameter, b *builder) error { return nil }},
				{output: func(e *objectDocParameter, b *builder) error { return nil }},
			},
			&objectDocHeadingElement{
				Text:   "Responses",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text: "\nIf an endpoint supports pagination, then the response object contains the below fields. \n",
				output: func(e *objectDocParagraphElement, b *builder) error {
					b.getAbstractObject("Pagination").fieldsComment += e.Text
					return nil
				},
			},
			&objectDocParametersElement{{
				Field:       "has_more",
				Type:        "boolean",
				Description: "Whether the response includes the end of the list. false if there are no more results. Otherwise, true.",
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("Pagination").addFields(e.asField(jen.Bool()))
					return nil
				},
			}, {
				Field:       "next_cursor",
				Type:        "string",
				Description: "A string that can be used to retrieve the next page of results by passing the value as the start_cursor parameter to the same endpoint.\n\nOnly available when has_more is true.",
				output: func(e *objectDocParameter, b *builder) error {
					// RetrievePagePropertyItemでnullを確認
					b.getAbstractObject("Pagination").addFields(e.asField(jen.Id("*").String()))
					return nil
				},
			}, {
				Field:       "object",
				Type:        `"list"`,
				Description: `The constant string "list".`,
				output: func(e *objectDocParameter, b *builder) error {
					b.getAbstractObject("Pagination").addFields(e.asFixedStringField())
					return nil
				},
			}, {
				Description: "The list, or partial list, of endpoint-specific results. Refer to a supported endpoint's individual documentation for details.",
				Field:       "results",
				Type:        "array of objects",
				output: func(e *objectDocParameter, b *builder) error {
					return nil // 各バリアントで定義
				},
			}, {
				Description: "A constant string that represents the type of the objects in results.",
				Field:       "type",
				Type:        "\"block\"\n\n\"comment\"\n\n\"database\"\n\n\"page\"\n\n\"page_or_database\"\n\n\"property_item\"\n\n\"user\"",
				output: func(e *objectDocParameter, b *builder) error {
					// 各種バリアントを作成
					for _, name := range strings.Split(e.Type, "\n\n") {
						name := strings.TrimPrefix(strings.TrimSuffix(name, `"`), `"`)
						b.getAbstractObject("Pagination").addVariant(
							b.addSpecificObject(strcase.UpperCamelCase(name)+"Pagination", "").addFields(
								&fixedStringField{name: "type", value: name},
							),
						)
					}
					b.addAbstractObjectToGlobalIfNotExists("PropertyItemOrPropertyItemPagination", "object").addVariant(
						b.getSpecificObject("PropertyItemPagination"),
					)
					return nil
				},
			}, {
				Description: "An object containing type-specific pagination information. For\u00a0property_items, the value corresponds to the\u00a0paginated page property type. For all other types, the value is an empty object.",
				Field:       "{type}",
				Type:        "paginated list object",
				output: func(e *objectDocParameter, b *builder) error {
					// TODO 各バリアントを定義
					b.getSpecificObject("PagePagination").addFields(
						&field{name: "page", typeCode: jen.Struct(), comment: e.Description},
						&field{name: "results", typeCode: jen.Index().Id("Page")},
					)
					b.getSpecificObject("PropertyItemPagination").addFields(
						&interfaceField{name: "property_item", typeName: "PaginatedPropertyInfo", comment: e.Description},
						&field{name: "results", typeCode: jen.Id("PropertyItems")},
					)
					return nil
				},
			}},
			&objectDocHeadingElement{
				Text:   "Parameters for paginated requests",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocCalloutElement{
				Body:   "`GET` requests accept parameters in the query string. \n\n`POST` requests receive parameters in the request body.",
				Title:  "Parameter location varies by endpoint",
				Type:   "warning",
				output: func(e *objectDocCalloutElement, b *builder) error { return nil },
			},
			&objectDocParametersElement{{
				Description: "The number of items from the full list to include in the response.\n\nDefault: 100\nMaximum: 100\n\nThe response may contain fewer than the default number of results.",
				Type:        "number",
				output:      func(e *objectDocParameter, b *builder) error { return nil },
			}, {
				Description: "A next_cursor value returned in a previous response. Treat this as an opaque value.\n\nDefaults to undefined, which returns results from the beginning of the list.",
				Type:        "string",
				output:      func(e *objectDocParameter, b *builder) error { return nil },
			}},
			&objectDocHeadingElement{
				Text:   "How to send a paginated request ",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocParagraphElement{
				Text:   "\n1. Send an initial request to the supported endpoint.\n2. Retrieve the next_cursor value from the response (only available when has_more is true). \n3. Send a follow up request to the endpoint that includes the next_cursor param in either the query string (for GET requests) or in the body params (POST requests).\n",
				output: func(e *objectDocParagraphElement, b *builder) error { return nil },
			},
			&objectDocHeadingElement{
				Text:   "Example: request the next set of query results from a database ",
				output: func(e *objectDocHeadingElement, b *builder) error { return nil },
			},
			&objectDocCodeElement{Codes: []*objectDocCodeElementCode{{
				Code:     "curl --location --request POST 'https://api.notion.com/v1/databases/<database_id>/query' \\\n--header 'Authorization: Bearer <secret_bot>' \\\n--header 'Content-Type: application/json' \\\n--data '{\n    \"start_cursor\": \"33e19cb9-751f-4993-b74d-234d67d0d534\"\n}'",
				Language: "curl",
				Name:     "",
				output:   func(e *objectDocCodeElementCode, b *builder) error { return nil },
			}}},
		},
	})
}
