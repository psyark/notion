package doc2api

import (
	"testing"
)

func TestObjects(t *testing.T) {
	if err := convertAll(); err != nil {
		t.Fatal(err)
	}
}
