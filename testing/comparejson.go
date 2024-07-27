package testing

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/flytam/filenamify"
	"github.com/psyark/notion"
	"github.com/samber/lo"
	"github.com/wI2L/jsondiff"
)

func compareJSON(t *testing.T) notion.CallOption {
	fileName := lo.Must(filenamify.FilenamifyV2(t.Name()))

	return notion.WithValidator(func(wantBytes []byte, got any) error {
		gotBytes, err := json.Marshal(got)
		if err != nil {
			return err
		}

		if diff, err := jsondiff.CompareJSON(gotBytes, wantBytes); err != nil {
			return err
		} else if diff != nil {
			_ = os.Remove(fmt.Sprintf("testdata/pass/%s.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.want.json", fileName), wantBytes, 0666); err != nil {
				return err
			}
			if err := os.WriteFile(fmt.Sprintf("testdata/fail/%s.got.json", fileName), gotBytes, 0666); err != nil {
				return err
			}
			return fmt.Errorf("validation failed:\n%s", diff.String())
		} else {
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.want.json", fileName))
			_ = os.Remove(fmt.Sprintf("testdata/fail/%s.got.json", fileName))
			if err := os.WriteFile(fmt.Sprintf("testdata/pass/%s.json", fileName), wantBytes, 0666); err != nil {
				return err
			}
		}

		return nil
	})
}
