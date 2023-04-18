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
	path           string
	method         string
	requestId      string
	useCache       bool
	validateResult bool
	params         any
	result         any
}

type callOption func(*callOptions)

func requestId(requestId string) callOption {
	return func(co *callOptions) {
		co.requestId = requestId
	}
}

func useCache() callOption {
	return func(co *callOptions) {
		co.useCache = true
	}
}

func validateResult() callOption {
	return func(co *callOptions) {
		co.validateResult = true
	}
}

func (c *Client) call(ctx context.Context, options *callOptions) error {
	if options.useCache && options.requestId == "" {
		return fmt.Errorf("useCache requires requestId")
	}
	if options.validateResult && options.requestId == "" {
		return fmt.Errorf("validateResult requires requestId")
	}

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

	var res *http.Response

	var fileName string
	if options.requestId != "" {
		fileName, err = filenamify.FilenamifyV2(options.requestId)
		if err != nil {
			return err
		}
	}

	if options.useCache {
		if cache, err := os.Open(fmt.Sprintf("testout/cache/%s", fileName)); err == nil {
			defer func() {
				_ = cache.Close()
			}()

			res, err = http.ReadResponse(bufio.NewReader(cache), req)
			if err != nil {
				return err
			}

			defer func() {
				_ = res.Body.Close()
			}()
		}
	}

	if res == nil {
		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer func() {
			_ = res.Body.Close()
		}()

		if options.useCache {
			if dump, err := httputil.DumpResponse(res, true); err != nil {
				return err
			} else if err := os.WriteFile(fmt.Sprintf("testout/cache/%s", fileName), dump, 0666); err != nil {
				return err
			}
		}
	}

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
		if options.validateResult {
			_ = os.Remove(fmt.Sprintf("testout/pass/%s.json", fileName))
			_ = os.Remove(fmt.Sprintf("testout/fail/%s.got.json", fileName))
			want := normalizeJSON(resBody)
			if err := os.WriteFile(fmt.Sprintf("testout/fail/%s.want.json", fileName), want, 0666); err != nil {
				return err
			}
		}
		return err
	}

	if options.validateResult {
		got, err := json.Marshal(options.result)
		if err != nil {
			return err
		}

		got = normalizeJSON(got)
		want := normalizeJSON(resBody)

		if bytes.Equal(want, got) {
			_ = os.Remove(fmt.Sprintf("testout/fail/%s.want.json", fileName))
			_ = os.Remove(fmt.Sprintf("testout/fail/%s.got.json", fileName))
			return os.WriteFile(fmt.Sprintf("testout/pass/%s.json", fileName), want, 0666)
		} else {
			_ = os.Remove(fmt.Sprintf("testout/pass/%s.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testout/fail/%s.want.json", fileName), want, 0666); err != nil {
				return err
			}
			if err := os.WriteFile(fmt.Sprintf("testout/fail/%s.got.json", fileName), got, 0666); err != nil {
				return err
			}
			return fmt.Errorf("validation failed: %s", options.requestId)
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
