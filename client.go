package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const APIVersion = "2022-06-28"

func NewClient(accessToken string) *Client {
	return &Client{accessToken: accessToken}
}

type Client struct {
	accessToken string
}

type callOptions struct {
	path       string
	method     string
	bodyWriter io.Writer
	params     any
	result     any
}

type callOption func(*callOptions)

func (c *Client) call(ctx context.Context, options *callOptions) error {
	payload, err := json.Marshal(options.params)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, options.method, "https://api.notion.com"+options.path, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("Notion-Version", APIVersion)

	switch options.method {
	case http.MethodPost, http.MethodPatch:
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if options.bodyWriter != nil {
		if _, err := options.bodyWriter.Write(resBody); err != nil {
			return err
		}
	}

	if res.StatusCode != http.StatusOK {
		errBody := Error{}
		if err := json.Unmarshal(resBody, &errBody); err != nil {
			return fmt.Errorf("bad status: %v, %v", res.Status, string(resBody))
		} else {
			return errBody
		}
	}

	return json.Unmarshal(resBody, &options.result)
}
