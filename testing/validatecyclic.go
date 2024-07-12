package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/flytam/filenamify"
	"github.com/samber/lo"
)

func checkRoundTrip(t *testing.T) func([]byte, any) error {
	fileName := lo.Must(filenamify.FilenamifyV2(t.Name()))

	return func(resBody []byte, unmarshaller any) error {
		got, err := json.Marshal(unmarshaller)
		if err != nil {
			return err
		}

		if want, got, ok := compareJSON(resBody, got); ok {
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.want.json", fileName))
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.got.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/pass/%s.json", fileName), want, 0666); err != nil {
				return err
			}
		} else {
			_ = os.Remove(fmt.Sprintf("testdata/pass/%s.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.want.json", fileName), want, 0666); err != nil {
				return err
			}
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.got.json", fileName), got, 0666); err != nil {
				return err
			}
			return fmt.Errorf("validation failed")
		}

		return nil
	}
}

func normalizeJSON(src []byte) []byte {
	tmp := map[string]any{}
	json.Unmarshal(src, &tmp)
	out, _ := json.MarshalIndent(tmp, "", "  ")
	return out
}

func compareJSON(data1 []byte, data2 []byte) (data1N []byte, data2N []byte, ok bool) {
	data1N = normalizeJSON(data1)
	data2N = normalizeJSON(data2)
	ok = bytes.Equal(data1N, data2N)
	return
}
