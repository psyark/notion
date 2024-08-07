package objects_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/psyark/notion/doc2api/objects"
)

func TestFile(t *testing.T) {
	t.Parallel()

	c := converter.FetchDocument("https://developers.notion.com/reference/file-object")

	var file *UnionStruct

	c.ExpectBlock(&Block{
		Kind: "Blockquote",
		Text: "📘The Notion API does not yet support uploading files to Notion.",
	}).Output(func(e *Block, b *CodeBuilder) {
		file = b.AddUnionStruct("File", "type", e.Text)
		file.AddFields(b.NewField(&Parameter{Property: "name", Description: UNDOCUMENTED}, jen.String(), OmitEmpty))

		union := converter.RegisterUnionInterface("FileOrEmoji", "type")
		converter.RegisterUnionMember(union, file, "")
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "File objects contain data about a file that is uploaded to Notion, or data about an external file that is linked to in Notion.",
	}).Output(func(e *Block, b *CodeBuilder) {
		file.AddComment(e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"file\",\n  \"file\": {\n    \"url\": \"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/7b8b0713-dbd4-4962-b38b-955b6c49a573/My_test_image.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20221024%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20221024T205211Z&X-Amz-Expires=3600&X-Amz-Signature=208aa971577ff05e75e68354e8a9488697288ff3fb3879c2d599433a7625bf90&X-Amz-SignedHeaders=host&x-id=GetObject\",\n    \"expiry_time\": \"2022-10-24T22:49:22.765Z\"\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("File", e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"type\": \"external\",\n  \"external\": {\n    \"url\": \"https://images.unsplash.com/photo-1525310072745-f49212b5ac6d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1065&q=80\"\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("File", e.Text)
	})
	c.ExpectBlock(&Block{
		Kind: "Paragraph",
		Text: "Page, embed, image, video, file, pdf, and bookmark block types all contain file objects. Icon and cover page object values also contain file objects.",
	}).Output(func(e *Block, b *CodeBuilder) {
		file.AddComment(e.Text)
	})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Each file object includes the following fields:"})
	c.ExpectParameter(&Parameter{
		Property:     "type",
		Type:         `string (enum)`,
		Description:  "The type of the file object. Possible type values are:\u00a0\"external\",\u00a0\"file\".",
		ExampleValue: `"external"`,
	})
	c.ExpectParameter(&Parameter{
		Property:     `external \| file`,
		Type:         "object",
		Description:  "An object containing type-specific configuration. The key of the object is external for external files, and file for Notion-hosted files.  \n  \nRefer to the type sections below for details on type-specific values.",
		ExampleValue: "Refer to to the type-specific sections below for examples.",
	})

	{

		var fileFile *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "Notion-hosted files",
		}).Output(func(e *Block, b *CodeBuilder) {
			fileFile = file.AddPayloadField("file", e.Text, WithPayloadObject(b))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: `All Notion-hosted files have a type of "file". The corresponding file specific object contains the following fields:`,
		}).Output(func(e *Block, b *CodeBuilder) {
			fileFile.AddComment(e.Text)
		})
		c.ExpectParameter(&Parameter{
			Property:     "url",
			Type:         "string",
			Description:  "An authenticated S3 URL to the file.  \n  \nThe URL is valid for one hour. If the link expires, then you can send an API request to get an updated URL.",
			ExampleValue: `"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/9bc6c6e0-32b8-4d55-8c12-3ae931f43a01/brocolli.jpeg?..."`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			fileFile.AddFields(b.NewField(e, jen.String()))
		})
		c.ExpectParameter(&Parameter{
			Property:     "expiry_time",
			Type:         `string (ISO 8601 date time)`,
			Description:  `The date and time when the link expires, formatted as an ISO 8601 date time string.`,
			ExampleValue: `"2020-03-17T19:10:04.968Z"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			fileFile.AddFields(b.NewField(e, jen.Id("ISO8601String")))
		})
		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "You can retrieve links to Notion-hosted files via the Retrieve block children endpoint.",
		}).Output(func(e *Block, b *CodeBuilder) {
			fileFile.AddComment(e.Text)
		})
	}

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: Retrieve a URL to a Notion-hosted file using GET /children"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The following example passes the ID of the page that includes the desired file as the block_id path param."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Request"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl 'https://api.notion.com/v1/blocks/13d6da822f9343fa8ec14c89b8184d5a/children?page_size=100' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Notion-Version: <<latestNotionVersion>>\"\n"})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Response"})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n    \"object\": \"list\",\n    \"results\": [\n        {\n            \"object\": \"block\",\n            \"id\": \"47a920e4-346c-4df8-ae78-905ce10adcb8\",\n            \"parent\": {\n                \"type\": \"page_id\",\n                \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n            },\n            \"created_time\": \"2022-12-15T00:18:00.000Z\",\n            \"last_edited_time\": \"2022-12-15T00:18:00.000Z\",\n            \"created_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"last_edited_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"has_children\": false,\n          \t\"archived\": false,\n            \"in_trash\": false,\n            \"type\": \"paragraph\",\n            \"paragraph\": {\n                \"rich_text\": [],\n                \"color\": \"default\"\n            }\n        },\n        {\n            \"object\": \"block\",\n            \"id\": \"3c29dedf-00a5-4915-b137-120c61f5e5d8\",\n            \"parent\": {\n                \"type\": \"page_id\",\n                \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n            },\n            \"created_time\": \"2022-12-15T00:18:00.000Z\",\n            \"last_edited_time\": \"2022-12-15T00:18:00.000Z\",\n            \"created_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"last_edited_by\": {\n                \"object\": \"user\",\n                \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n            },\n            \"has_children\": false,\n          \t\"archived\": false,\n            \"in_trash\": false,\n            \"type\": \"file\",\n            \"file\": {\n                \"caption\": [],\n                \"type\": \"file\",\n                \"file\": {\n                    \"url\": \"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/fa6c03f0-e608-45d0-9327-4cd7a5e56e71/TestFile.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20221215%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20221215T002012Z&X-Amz-Expires=3600&X-Amz-Signature=bf13ca59f618077852298cb92aedc4dd1becdc961c31d73cbc030ef93f2853c4&X-Amz-SignedHeaders=host&x-id=GetObject\",\n                    \"expiry_time\": \"2022-12-15T01:20:12.928Z\"\n                }\n            }\n        },\n    ],\n    \"next_cursor\": null,\n    \"has_more\": false,\n    \"type\": \"block\",\n    \"block\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		// ドキュメントの間違い（コンマが余計）
		e.Text = strings.Replace(
			e.Text,
			"},\n    ],\n    \"next_cursor\": null,",
			"} \n    ],\n    \"next_cursor\": null,",
			1,
		)
		converter.AddUnmarshalTest("Pagination", e.Text, "Block")
	})

	{
		var fileExternal *SimpleObject

		c.ExpectBlock(&Block{
			Kind: "Heading",
			Text: "External files",
		}).Output(func(e *Block, b *CodeBuilder) {
			fileExternal = file.AddPayloadField("external", e.Text, WithPayloadObject(b))
		})

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "An external file is any URL linked to in Notion that isn’t hosted by Notion. All external files have a type of \"external\". The corresponding file specific object contains the following fields:",
		}).Output(func(e *Block, b *CodeBuilder) {
			fileExternal.AddComment(e.Text)
		})

		c.ExpectParameter(&Parameter{
			Property:     "url",
			Type:         "string",
			Description:  "A link to the externally hosted content.",
			ExampleValue: `"https://website.domain/files/doc.txt"`,
		}).Output(func(e *Parameter, b *CodeBuilder) {
			fileExternal.AddFields(b.NewField(e, jen.String()))
		})

		c.ExpectBlock(&Block{
			Kind: "Paragraph",
			Text: "The Notion API supports adding, retrieving, and updating links to external files.",
		}).Output(func(e *Block, b *CodeBuilder) {
			fileExternal.AddComment(e.Text)
		})
	}

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: Add a URL to an external file using PATCH /children"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Use the Append block children endpoint to add external files to Notion. Pass a  block type object  in the body that that details information about the external file."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The following example request embeds a PDF in a Notion page. It passes the ID of the target page as the block_id path param and information about the file to append in the request body."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Request"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl -X PATCH 'https://api.notion.com/v1/blocks/13d6da822f9343fa8ec14c89b8184d5a/children' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  --data '{\n  \"children\": [\n    {\n      \"object\": \"block\",\n      \"type\": \"pdf\",\n      \"pdf\": {\n        \"type\": \"external\",\n        \"external\": {\n          \"url\": \"https://www.yourwebsite.dev/files/TestFile.pdf\"\n        }\n      }\n    }\n  ]\n}'\n"})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Response"})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"object\": \"list\",\n  \"results\": [\n    {\n      \"object\": \"block\",\n      \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n      \"parent\": {\n        \"type\": \"page_id\",\n        \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n      },\n      \"created_time\": \"2022-12-15T01:14:00.000Z\",\n      \"last_edited_time\": \"2022-12-15T01:14:00.000Z\",\n      \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"has_children\": false,\n      \"archived\": false,\n      \"in_trash\": false,\n      \"type\": \"pdf\",\n      \"pdf\": {\n        \"caption\": [],\n        \"type\": \"external\",\n        \"external\": {\n          \"url\": \"https://www.yourwebsite.dev/files/TestFile.pdf\"\n        }\n      }\n    }\n  ],\n  \"next_cursor\": null,\n  \"has_more\": false,\n  \"type\": \"block\",\n  \"block\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Pagination", e.Text, "Block")
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: Retrieve a link to an external file using GET /children"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Use the Retrieve block children endpoint on the file’s parent block in order to retrieve the file object. The file itself is contained in its own block of type file."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The following example passes the ID of the page that includes the desired file as the block_id path param."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Request"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl 'https://api.notion.com/v1/blocks/af1459f2-d2c5-4ca6-9f05-8038e6eb167f/children?page_size=100' \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Notion-Version: 2022-06-28\"\n"})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Response"})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"object\": \"list\",\n  \"results\": [\n    {\n      \"object\": \"block\",\n      \"id\": \"47a920e4-346c-4df8-ae78-905ce10adcb8\",\n      \"parent\": {\n        \"type\": \"page_id\",\n        \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n      },\n      \"created_time\": \"2022-12-15T00:18:00.000Z\",\n      \"last_edited_time\": \"2022-12-15T00:18:00.000Z\",\n      \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n      },\n      \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"c2f20311-9e54-4d11-8c79-7398424ae41e\"\n      },\n      \"has_children\": false,\n      \"archived\": false,\n      \"in_trash\": false,\n      \"type\": \"paragraph\",\n      \"paragraph\": {\n        \"rich_text\": [],\n        \"color\": \"default\"\n      }\n    },\n    {\n      \"object\": \"block\",\n      \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n      \"parent\": {\n        \"type\": \"page_id\",\n        \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n      },\n      \"created_time\": \"2022-12-15T01:14:00.000Z\",\n      \"last_edited_time\": \"2022-12-15T01:14:00.000Z\",\n      \"created_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"last_edited_by\": {\n        \"object\": \"user\",\n        \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n      },\n      \"has_children\": false,\n      \"archived\": false,\n      \"in_trash\": false,\n      \"type\": \"pdf\",\n      \"pdf\": {\n        \"caption\": [],\n        \"type\": \"external\",\n        \"external\": {\n          \"url\": \"https://www.yourwebsite.dev/files/TestFile.pdf\"\n        }\n      }\n    }\n  ],\n  \"next_cursor\": null,\n  \"has_more\": false,\n  \"type\": \"block\",\n  \"block\": {}\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Pagination", e.Text, "Block")
	})

	c.ExpectBlock(&Block{Kind: "Heading", Text: "Example: Update a link to an external file using PATCH /blocks/{block_id}"})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "Use the Update a block endpoint to update a link to an external file in Notion."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The body of the request depends on the file’s Notion block type. Common file block types include image, pdf, and video."})
	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘If you don’t know the file’s Notion block type, then send a request to the Retrieve block children endpoint using the parent block ID and find the corresponding child block."})
	c.ExpectBlock(&Block{Kind: "Paragraph", Text: "The following example updates the external file in a pdf block type. It passes the ID of the block as the block_id path param and information about the new file in the request body."})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Request"})
	c.ExpectBlock(&Block{Kind: "FencedCodeBlock", Text: "curl https://api.notion.com/v1/blocks/af1459f2-d2c5-4ca6-9f05-8038e6eb167f \\\n  -H 'Authorization: Bearer '\"$NOTION_API_KEY\"'' \\\n  -H \"Content-Type: application/json\" \\\n  -H \"Notion-Version: 2022-06-28\" \\\n  -X PATCH \\\n  --data '{\n\t  \"pdf\": { \n\t\t\t\"external\": {\n\t\t\t\t\"url\": \"https://www.yourwebsite.dev/files/NewFile.pdf\"\n\t\t\t}\n\t  }\n\t}'\n"})
	c.ExpectBlock(&Block{Kind: "Heading", Text: "Response"})

	c.ExpectBlock(&Block{
		Kind: "FencedCodeBlock",
		Text: "{\n  \"object\": \"block\",\n  \"id\": \"af1459f2-d2c5-4ca6-9f05-8038e6eb167f\",\n  \"parent\": {\n    \"type\": \"page_id\",\n    \"page_id\": \"13d6da82-2f93-43fa-8ec1-4c89b8184d5a\"\n  },\n  \"created_time\": \"2022-12-15T01:14:00.000Z\",\n  \"last_edited_time\": \"2022-12-16T21:23:00.000Z\",\n  \"created_by\": {\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n  },\n  \"last_edited_by\": {\n    \"object\": \"user\",\n    \"id\": \"9188c6a5-7381-452f-b3dc-d4865aa89bdf\"\n  },\n  \"has_children\": false,\n  \"archived\": false,\n  \"in_trash\": false,\n  \"type\": \"pdf\",\n  \"pdf\": {\n    \"caption\": [],\n    \"type\": \"external\",\n    \"external\": {\n      \"url\": \"https://www.yourwebsite.dev/files/NewFile.pdf\"\n    }\n  }\n}\n",
	}).Output(func(e *Block, b *CodeBuilder) {
		converter.AddUnmarshalTest("Block", e.Text)
	})

	c.ExpectBlock(&Block{Kind: "Blockquote", Text: "📘To modify page or database property values that are made from file objects, like icon, cover, or files page property values, use the update page or update database endpoints."})

	c.RequestBuilderForUndocumented(func(b *CodeBuilder) {
		file.AddFields(b.NewField(&Parameter{Property: "caption", Description: UNDOCUMENTED}, jen.Id("RichTextArray"), OmitEmpty))
	})
}
