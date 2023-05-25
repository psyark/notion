package notion

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Outer struct {
	Inner Nullable[float64] `json:"inner,omitempty"`
}

func TestNullable(t *testing.T) {
	cases := []string{
		`{"inner":null}`,
		`{"inner":123}`,
		`{}`,
	}
	for _, src := range cases {
		outer := Outer{}
		if err := json.Unmarshal([]byte(src), &outer); err != nil {
			t.Fatal(err)
		}

		data, err := json.Marshal(outer)
		if err != nil {
			t.Fatal(err)
		}

		if string(data) != src {
			t.Fatal(fmt.Errorf("%s != %s", string(data), src))
		}
	}
}
