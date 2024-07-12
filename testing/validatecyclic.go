package testing

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/flytam/filenamify"
	"github.com/samber/lo"
	"github.com/wI2L/jsondiff"
)

func checkRoundTrip(t *testing.T) func([]byte, any) error {
	fileName := lo.Must(filenamify.FilenamifyV2(t.Name()))

	return func(resBody []byte, unmarshaller any) error {
		got, err := json.Marshal(unmarshaller)
		if err != nil {
			return err
		}

		if diff, err := jsondiff.CompareJSON(resBody, got); err != nil {
			return err
		} else if diff != nil {
			_ = os.Remove(fmt.Sprintf("testdata/pass/%s.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.want.json", fileName), resBody, 0666); err != nil {
				return err
			}
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.got.json", fileName), got, 0666); err != nil {
				return err
			}
			return fmt.Errorf("validation failed:\n%s", diff.String())
		} else {
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.want.json", fileName))
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.got.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/pass/%s.json", fileName), resBody, 0666); err != nil {
				return err
			}
		}

		return nil
	}
}
