package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const APIVersion = "2022-06-28"

func NewClient(accessToken string) *Client {
	return &Client{accessToken: accessToken}
}

type Client struct {
	accessToken string
}

type callOptions struct {
	path           string
	method         string
	validateResult string
	params         any
	result         any
}

type callOption func(*callOptions)

func validateResult(testName string) callOption {
	return func(co *callOptions) {
		co.validateResult = testName
	}
}

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

	if res.StatusCode != http.StatusOK {
		errBody := Error{}
		if err := json.Unmarshal(resBody, &errBody); err != nil {
			return fmt.Errorf("bad status: %v, %v", res.Status, string(resBody))
		} else {
			return errBody
		}
	}

	if err := json.Unmarshal(resBody, options.result); err != nil {
		if options.validateResult != "" {
			_ = os.Remove(fmt.Sprintf("testout/%v.ok.json", options.validateResult))
			_ = os.Remove(fmt.Sprintf("testout/%v.want.json", options.validateResult))
			want := normalizeJSON(resBody)
			if err := os.WriteFile(fmt.Sprintf("testout/%v.want.json", options.validateResult), want, 0666); err != nil {
				return err
			}
		}
		return err
	}

	if options.validateResult != "" {
		got, err := json.Marshal(options.result)
		if err != nil {
			return err
		}

		got = normalizeJSON(got)
		want := normalizeJSON(resBody)

		if bytes.Equal(want, got) {
			_ = os.Remove(fmt.Sprintf("testout/%v.want.json", options.validateResult))
			_ = os.Remove(fmt.Sprintf("testout/%v.got.json", options.validateResult))
			return os.WriteFile(fmt.Sprintf("testout/%v.ok.json", options.validateResult), want, 0666)
		} else {
			_ = os.Remove(fmt.Sprintf("testout/%v.ok.json", options.validateResult))
			if err := os.WriteFile(fmt.Sprintf("testout/%v.want.json", options.validateResult), want, 0666); err != nil {
				return err
			}
			if err := os.WriteFile(fmt.Sprintf("testout/%v.got.json", options.validateResult), got, 0666); err != nil {
				return err
			}
			return fmt.Errorf("validation failed: %v", options.validateResult)
		}
	}

	return nil
}

func normalizeJSON(src []byte) []byte {
	tmp := map[string]interface{}{}
	json.Unmarshal(src, &tmp)
	out, _ := json.MarshalIndent(tmp, "", "  ")
	return out
}
