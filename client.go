package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/flytam/filenamify"
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

	requestId      string
	validateResult bool
}

type callOption func(*callOptions)

func RequestId(requestId string) callOption {
	return func(co *callOptions) {
		co.requestId = requestId
	}
}

func ValidateResult() callOption {
	return func(co *callOptions) {
		co.validateResult = true
	}
}

func WithRoundTripper(roundTripper http.RoundTripper) callOption {
	return func(co *callOptions) {
		co.roundTripper = roundTripper
	}
}

// Deprecated: use request
func call[U any, R any](ctx context.Context, accessToken string, method string, path string, params map[string]any, getResult func(unmarshaller *U) R, options ...callOption) (R, error) {
	var unmarshaller U
	var zero R

	co := &callOptions{
		roundTripper: http.DefaultTransport,
	}
	for _, o := range options {
		o(co)
	}

	if co.validateResult && co.requestId == "" {
		return zero, fmt.Errorf("validateResult requires requestId")
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

	var fileName string
	if co.requestId != "" {
		fileName, err = filenamify.FilenamifyV2(co.requestId)
		if err != nil {
			return zero, err
		}
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
		if co.validateResult {
			_ = os.Remove(fmt.Sprintf("testdata/pass/%s.json", fileName))
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.got.json", fileName))
			want := normalizeJSON(resBody)
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.want.json", fileName), want, 0666); err != nil {
				return zero, err
			}
		}
		return zero, err
	}

	// TODO ここにファイルシステムのパスを含む処理を書かず、オプションとして渡すようにする
	if co.validateResult {
		got, err := json.Marshal(unmarshaller)
		if err != nil {
			return zero, err
		}

		if want, got, ok := compareJSON(resBody, got); ok {
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.want.json", fileName))
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.got.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/pass/%s.json", fileName), want, 0666); err != nil {
				return zero, err
			}
		} else {
			_ = os.Remove(fmt.Sprintf("testdata/pass/%s.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.want.json", fileName), want, 0666); err != nil {
				return zero, err
			}
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.got.json", fileName), got, 0666); err != nil {
				return zero, err
			}
			return zero, fmt.Errorf("validation failed: %s", co.requestId)
		}
	}

	return getResult(&unmarshaller), nil
}
