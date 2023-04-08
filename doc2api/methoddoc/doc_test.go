package methoddoc

import "testing"

func TestConvertAll(t *testing.T) {
	if err := convertAll(); err != nil {
		t.Fatal(err)
	}
}
