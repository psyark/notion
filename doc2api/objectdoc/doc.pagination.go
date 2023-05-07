package objectdoc

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

func init() {
	registerTranslator(
		"https://developers.notion.com/reference/intro",
		func(c *comparator, b *builder) /*  */ {
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The reference is your key to a comprehensive understanding of the Notion API.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Integrations use the API to access Notion's pages, databases, and users. Integrations can connect services to Notion and build interactive experiences for users within Notion. Using the navigation on the left, you'll find details for objects and endpoints used in the API.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "You need an integration token to interact with the Notion API. You can find an integration token after you create an integration on the integration settings page. If this is your first look at the Notion API, we recommend beginning with the Getting started guide to learn how to create an integration. \n\nIf you want to work on a specific integration, but can't access the token, confirm that you are an admin in the associated workspace. You can check inside the Notion UI via Settings & Members in the left sidebar. If you're not an admin in any of your workspaces, you can create a personal workspace for free.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Conventions */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Conventions",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The base URL to send all API requests is https://api.notion.com. HTTPS is required for all API requests.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "The Notion API follows RESTful conventions when possible, with most operations performed via GET, POST, PATCH, and DELETE requests on page and database resources. Request and response bodies are encoded as JSON.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* JSON conventions */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "JSON conventions",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Top-level resources have an \"object\" property. This property can be used to determine the type of the resource (e.g. \"database\", \"user\", etc.)Top-level resources are addressable by a UUIDv4 \"id\" property. You may omit dashes from the ID when making requests to the API, e.g. when copying the ID from a Notion URL.Property names are in snake_case (not camelCase or kebab-case).Temporal values (dates and datetimes) are encoded in ISO 8601 strings. Datetimes will include the time value (2020-08-12T02:12:33.231Z) while dates will include only the date (2020-08-12)The Notion API does not support empty strings. To unset a string value for properties like a url Property value object, for example, use an explicit null instead of \"\".",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Code samples & SDKs */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Code samples & SDKs",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Samples requests and responses are shown for each endpoint. Requests are shown using the Notion JavaScript SDK, and cURL. These samples make it easy to copy, paste, and modify as you build your integration.",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Notion SDKs are open source projects that you can install to easily start building. You may also choose any other language or library that allows you to make HTTP requests.",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Pagination */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Pagination",
			}, func(e blockElement) {
				b.addAbstractObject("Pagination", "type", e.Text)
			})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "Endpoints that return lists of objects support cursor-based pagination requests. By default, Notion returns ten items per API call. If the number of items in a response from a support endpoint exceeds the default, then an integration can use pagination to request a specific set of the results and/or to limit the number of returned items.",
			}, func(e blockElement) {
				getSymbol[abstractObject](b, "Pagination").addComment(e.Text)
			})
		},
		func(c *comparator, b *builder) /* Supported endpoints */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Supported endpoints",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description: "List all users",
				Type:        "GET",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description: "Retrieve block children",
				Type:        "GET",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description: "Retrieve a comment",
				Type:        "GET",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description: "Retrieve a page property item",
				Type:        "GET",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description: "Query a database",
				Type:        "POST",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description: "Search",
				Type:        "POST",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* Responses */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Responses",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Paragraph",
				Text: "If an endpoint supports pagination, then the response object contains the below fields.",
			}, func(e blockElement) {
				getSymbol[abstractObject](b, "Pagination").addComment(e.Text)
			})
			c.nextMustParameter(parameterElement{
				Property:    "has_more",
				Type:        "boolean",
				Description: "Whether the response includes the end of the list. false if there are no more results. Otherwise, true.",
			}, func(e parameterElement) {
				getSymbol[abstractObject](b, "Pagination").addFields(e.asField(jen.Bool()))
			})
			c.nextMustParameter(parameterElement{
				Property:    "next_cursor",
				Type:        "string",
				Description: "A string that can be used to retrieve the next page of results by passing the value as the start_cursor parameter to the same endpoint.\n\nOnly available when has_more is true.",
			}, func(e parameterElement) {
				// RetrievePagePropertyItemでnullを確認
				getSymbol[abstractObject](b, "Pagination").addFields(e.asField(jen.Id("*").String()))
			})
			c.nextMustParameter(parameterElement{
				Property:    "object",
				Type:        `"list"`,
				Description: "The constant string \"list\".",
			}, func(e parameterElement) {
				getSymbol[abstractObject](b, "Pagination").addFields(e.asFixedStringField(b))
			})
			c.nextMustParameter(parameterElement{
				Description:  "The list, or partial list, of endpoint-specific results. Refer to a supported endpoint's individual documentation for details.",
				ExampleValue: "",
				Property:     "results",
				Type:         "array of objects",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "A constant string that represents the type of the objects in results.",
				ExampleValue: "",
				Property:     "type",
				Type:         "\"block\"\n\n\"comment\"\n\n\"database\"\n\n\"page\"\n\n\"page_or_database\"\n\n\"property_item\"\n\n\"user\"",
			}, func(e parameterElement) {
				for _, name := range strings.Split(e.Type, "\n\n") {
					name := strings.TrimPrefix(strings.TrimSuffix(name, `"`), `"`)

					typeSpecificField := &field{name: name, typeCode: jen.Struct()}
					if name == "property_item" {
						typeSpecificField.typeCode = jen.Id("PaginatedPropertyInfo")
					}

					var resultsField fieldCoder = &field{name: "results", typeCode: jen.Id(strcase.UpperCamelCase(name) + "List")}
					switch name {
					case "database", "page", "block", "user", "property_item":
						resultsField = &field{name: "results", typeCode: jen.Index().Id(strcase.UpperCamelCase(name))}
					}

					b.addDerived(name, "Pagination", "").addFields(
						typeSpecificField,
						resultsField,
					)
				}
				b.addUnionToGlobalIfNotExists("PropertyItemOrPropertyItemPagination", "object")
				getSymbol[concreteObject](b, "PropertyItemPagination").addToUnion(getSymbol[unionObject](b, "PropertyItemOrPropertyItemPagination"))
			})
			c.nextMustParameter(parameterElement{
				Property:    "{type}",
				Type:        "paginated list object",
				Description: "An object containing type-specific pagination information. For\u00a0property_items, the value corresponds to the\u00a0paginated page property type. For all other types, the value is an empty object.",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* Parameters for paginated requests */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Parameters for paginated requests",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "Blockquote",
				Text: "GET requests accept parameters in the query string. \n\nPOST requests receive parameters in the request body.",
			}, func(e blockElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "The number of items from the full list to include in the response.\n\nDefault: 100\nMaximum: 100\n\nThe response may contain fewer than the default number of results.",
				ExampleValue: "",
				Property:     "page_size",
				Type:         "number",
			}, func(e parameterElement) {})
			c.nextMustParameter(parameterElement{
				Description:  "A next_cursor value returned in a previous response. Treat this as an opaque value.\n\nDefaults to undefined, which returns results from the beginning of the list.",
				ExampleValue: "",
				Property:     "start_cursor",
				Type:         "string",
			}, func(e parameterElement) {})
		},
		func(c *comparator, b *builder) /* How to send a paginated request */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "How to send a paginated request",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "List",
				Text: "Send an initial request to the supported endpoint.Retrieve the next_cursor value from the response (only available when has_more is true).Send a follow up request to the endpoint that includes the next_cursor param in either the query string (for GET requests) or in the body params (POST requests).",
			}, func(e blockElement) {})
		},
		func(c *comparator, b *builder) /* Example: request the next set of query results from a database */ {
			c.nextMustBlock(blockElement{
				Kind: "Heading",
				Text: "Example: request the next set of query results from a database",
			}, func(e blockElement) {})
			c.nextMustBlock(blockElement{
				Kind: "FencedCodeBlock",
				Text: "curl --location --request POST 'https://api.notion.com/v1/databases/<database_id>/query' \\\n--header 'Authorization: Bearer <secret_bot>' \\\n--header 'Content-Type: application/json' \\\n--data '{\n    \"start_cursor\": \"33e19cb9-751f-4993-b74d-234d67d0d534\"\n}'",
			}, func(e blockElement) {})
		},
	)
}
