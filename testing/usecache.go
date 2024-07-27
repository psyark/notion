package testing

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"

	"github.com/flytam/filenamify"
	"github.com/samber/lo"
)

type cache struct {
	filePath string
}

// TODO 引数を *testing.T にする
func useCache(t *testing.T) *cache {
	fileName := lo.Must(filenamify.FilenamifyV2(t.Name()))
	return &cache{filePath: fmt.Sprintf("testdata/cache/%s", fileName)}
}

func (c *cache) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	if cache, err := os.Open(c.filePath); err == nil {
		defer cache.Close()

		res, err = http.ReadResponse(bufio.NewReader(cache), req)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
	} else {
		res, err = http.DefaultTransport.RoundTrip(req)
		if err != nil {
			return nil, err
		}
	}

	if dump, err := httputil.DumpResponse(res, true); err != nil {
		return nil, err
	} else if err := os.WriteFile(c.filePath, dump, 0666); err != nil {
		return nil, err
	}

	return res, err
}
