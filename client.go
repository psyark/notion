package notion

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
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
	useCache       bool
	validateResult bool
}

type callOption func(*callOptions)

func requestId(requestId string) callOption {
	return func(co *callOptions) {
		co.requestId = requestId
	}
}

func useCacheDeprecated() callOption {
	fmt.Println("useCacheDeprecated")
	return func(co *callOptions) {
		co.useCache = true
	}
}

func validateResult() callOption {
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

	co := &callOptions{}
	for _, o := range options {
		o(co)
	}

	if co.useCache && co.requestId == "" {
		return zero, fmt.Errorf("useCache requires requestId")
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

	var res *http.Response

	var fileName string
	if co.requestId != "" {
		fileName, err = filenamify.FilenamifyV2(co.requestId)
		if err != nil {
			return zero, err
		}
	}

	if co.useCache {
		if cache, err := os.Open(fmt.Sprintf("testdata/cache/%s", fileName)); err == nil {
			defer func() {
				_ = cache.Close()
			}()

			res, err = http.ReadResponse(bufio.NewReader(cache), req)
			if err != nil {
				return zero, err
			}

			defer func() {
				_ = res.Body.Close()
			}()
		}
	}

	if res == nil {
		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return zero, err
		}

		defer func() {
			_ = res.Body.Close()
		}()

		if co.useCache {
			if dump, err := httputil.DumpResponse(res, true); err != nil {
				return zero, err
			} else if err := os.WriteFile(fmt.Sprintf("testdata/cache/%s", fileName), dump, 0666); err != nil {
				return zero, err
			}
		}
	}

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
