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
	roundTripper http.RoundTripper
	validator    func(wantBytes []byte, got any) error
}

type callOption func(*callOptions)

func WithRoundTripper(roundTripper http.RoundTripper) callOption {
	return func(co *callOptions) {
		co.roundTripper = roundTripper
	}
}

func WithValidator(validator func(data []byte, unmarshaller any) error) callOption {
	return func(co *callOptions) {
		co.validator = validator
	}
}

// Deprecated: use request
func call[U any, R any](ctx context.Context, accessToken string, method string, path string, params map[string]any, getResult func(unmarshaller *U) R, options ...callOption) (R, error) {
	var unmarshaller U
	var zero R

	co := &callOptions{roundTripper: http.DefaultTransport}
	for _, o := range options {
		o(co)
	}

	payload, err := json.Marshal(params)
	if err != nil {
		return zero, err
	}

	req, err := http.NewRequestWithContext(ctx, method, "https://api.notion.com"+path, bytes.NewBuffer(payload))
	if err != nil {
		return zero, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Notion-Version", APIVersion)

	switch method {
	case http.MethodPost, http.MethodPatch:
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := co.roundTripper.RoundTrip(req)
	if err != nil {
		return zero, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return zero, err
	}

	if res.StatusCode != http.StatusOK {
		errBody := Error{}
		if err := json.Unmarshal(resBody, &errBody); err != nil {
			return zero, fmt.Errorf("bad status: %v, %v", res.Status, string(resBody))
		} else {
			return zero, errBody
		}
	}

	if err := json.Unmarshal(resBody, &unmarshaller); err != nil {
		// if co.validateResult {
		// 	_ = os.Remove(fmt.Sprintf("testdata/pass/%s.json", fileName))
		// 	_ = os.Remove(fmt.Sprintf("testdata/fail/%s.got.json", fileName))
		// 	want := normalizeJSON(resBody)
		// 	if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.want.json", fileName), want, 0666); err != nil {
		// 		return zero, err
		// 	}
		// }
		return zero, err
	}

	if co.validator != nil {
		if err := co.validator(resBody, getResult(&unmarshaller)); err != nil {
			return zero, err
		}
	}

	return getResult(&unmarshaller), nil
}
