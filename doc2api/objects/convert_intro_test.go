package objects_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestIntro(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/intro")

	var pagination *AdaptiveObject

	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The reference is your key to a comprehensive understanding of the Notion API."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Integrations use the API to access Notion's pages, databases, and users. Integrations can connect services to Notion and build interactive experiences for users within Notion. Using the navigation on the left, you'll find details for objects and endpoints used in the API."})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘You need an integration token to interact with the Notion API. You can find an integration token after you create an integration on the integration settings page. If this is your first look at the Notion API, we recommend beginning with the Getting started guide to learn how to create an integration.If you want to work on a specific integration, but can't access the token, confirm that you are an admin in the associated workspace. You can check inside the Notion UI via Settings & Members in the left sidebar. If you're not an admin in any of your workspaces, you can create a personal workspace for free."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Conventions"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The base URL to send all API requests is https://api.notion.com. HTTPS is required for all API requests."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The Notion API follows RESTful conventions when possible, with most operations performed via GET, POST, PATCH, and DELETE requests on page and database resources. Request and response bodies are encoded as JSON."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "JSON conventions"})
	c.ExpectBlock(&Block{Kind: "List", Text: "Top-level resources have an \"object\" property. This property can be used to determine the type of the resource (e.g. \"database\", \"user\", etc.)Top-level resources are addressable by a UUIDv4 \"id\" property. You may omit dashes from the ID when making requests to the API, e.g. when copying the ID from a Notion URL.Property names are in snake_case (not camelCase or kebab-case).Temporal values (dates and datetimes) are encoded in ISO 8601 strings. Datetimes will include the time value (2020-08-12T02:12:33.231Z) while dates will include only the date (2020-08-12)The Notion API does not support empty strings. To unset a string value for properties like a url Property value object, for example, use an explicit null instead of \"\"."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Code samples & SDKs"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Samples requests and responses are shown for each endpoint. Requests are shown using the Notion JavaScript SDK, and cURL. These samples make it easy to copy, paste, and modify as you build your integration."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Notion SDKs are open source projects that you can install to easily start building. You may also choose any other language or library that allows you to make HTTP requests."})

	c.ExpectBlock(&Block{
		Kind: "Heading",
		Text: "Pagination",
	}).Output(func(e *Block, b *CodeBuilder) {
		pagination = b.AddAdaptiveObject("Pagination", "type", e.Text)
		pagination.AddToUnion(b.AddUnionToGlobalIfNotExists("PropertyItemOrPropertyItemPagination", "object"))
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Endpoints that return lists of objects support cursor-based pagination requests. By default, Notion returns ten items per API call. If the number of items in a response from a support endpoint exceeds the default, then an integration can use pagination to request a specific set of the results and/or to limit the number of returned items.",
	}).Output(func(e *Block, b *CodeBuilder) {
		pagination.AddComment(e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Supported endpoints"})
	c.ExpectParameter(&Parameter{Description: "List all users", Type: "GET"})
	c.ExpectParameter(&Parameter{Description: "Retrieve block children", Type: "GET"})
	c.ExpectParameter(&Parameter{Description: "Retrieve a comment", Type: "GET"})
	c.ExpectParameter(&Parameter{Description: "Retrieve a page property item", Type: "GET"})
	c.ExpectParameter(&Parameter{Description: "Query a database", Type: "POST"})
	c.ExpectParameter(&Parameter{Description: "Search", Type: "POST"})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Responses"})

	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "If an endpoint supports pagination, then the response object contains the below fields.",
	}).Output(func(e *Block, b *CodeBuilder) {
		pagination.AddComment(e.Text)
	})

	c.ExpectParameter(&Parameter{
		Property:    "has_more",
		Type:        "boolean",
		Description: "Whether the response includes the end of the list. false if there are no more results. Otherwise, true.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		pagination.AddFields(b.NewField(e, jen.Bool()))
	})
	c.ExpectParameter(&Parameter{
		Property:    "next_cursor",
		Type:        "string",
		Description: "A string that can be used to retrieve the next page of results by passing the value as the start_cursor parameter to the same endpoint.  \n  \nOnly available when has_more is true.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		// RetrievePagePropertyItemでnullを確認
		pagination.AddFields(b.NewField(e, jen.Id("*").String()))
	})
	c.ExpectParameter(&Parameter{
		Property:    "object",
		Type:        `"list"`,
		Description: "The constant string \"list\".",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		pagination.AddFields(b.NewDiscriminatorField(e))
	})
	c.ExpectParameter(&Parameter{
		Property:    "results",
		Type:        "array of objects",
		Description: "The list, or partial list, of endpoint-specific results. Refer to a supported endpoint's individual documentation for details.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		pagination.AddFields(b.NewField(e, jen.Qual("encoding/json", "RawMessage")))
	})
	c.ExpectParameter(&Parameter{
		Property:    "type",
		Type:        "\"block\"  \n  \n\"comment\"  \n  \n\"database\"  \n  \n\"page\"  \n  \n\"page_or_database\"  \n  \n\"property_item\"  \n  \n\"user\"",
		Description: "A constant string that represents the type of the objects in results.",
	}).Output(func(e *Parameter, b *CodeBuilder) {
		for _, name := range strings.Split(e.Type, "  \n  \n") {
			name := strings.TrimPrefix(strings.TrimSuffix(name, `"`), `"`)
			if name == "property_item" {
				pagination.AddAdaptiveFieldWithType(name, "", jen.Id("PaginatedPropertyInfo"))
			} else {
				pagination.AddAdaptiveFieldWithEmptyStruct(name, "")
			}
		}
	})
	c.ExpectParameter(&Parameter{
		Property:    "{type}",
		Type:        "paginated list object",
		Description: "An object containing type-specific pagination information. For\u00a0property_items, the value corresponds to the\u00a0paginated page property type. For all other types, the value is an empty object.",
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Parameters for paginated requests"})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "🚧 Parameter location varies by endpointGET requests accept parameters in the query string.POST requests receive parameters in the request body."})

	c.ExpectParameter(&Parameter{
		Property:    "page_size",
		Type:        "number",
		Description: "The number of items from the full list to include in the response.  \n  \nDefault: 100  \nMaximum: 100  \n  \nThe response may contain fewer than the default number of results.",
	})
	c.ExpectParameter(&Parameter{
		Property:    "start_cursor",
		Type:        "string",
		Description: "A next_cursor value returned in a previous response. Treat this as an opaque value.  \n  \nDefaults to undefined, which returns results from the beginning of the list.",
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "How to send a paginated request"})
	c.ExpectBlock(&Block{Kind: "List", Text: "Send an initial request to the supported endpoint.Retrieve the next_cursor value from the response (only available when has_more is true).Send a follow up request to the endpoint that includes the next_cursor param in either the query string (for GET requests) or in the body params (POST requests)."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: request the next set of query results from a database"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl --location --request POST 'https://api.notion.com/v1/databases/<database_id>/query' \\\n--header 'Authorization: Bearer <secret_bot>' \\\n--header 'Content-Type: application/json' \\\n--data '{\n    \"start_cursor\": \"33e19cb9-751f-4993-b74d-234d67d0d534\"\n}'\n"})

	c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
		pagination.AddFields(UndocumentedRequestID(b))
	})
}
