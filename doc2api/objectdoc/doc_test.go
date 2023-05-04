package objectdoc

import (
	"testing"
)

func TestConvertAll(t *testing.T) {
	if err := translateAll(); err != nil {
		t.Fatal(err)
	}
	if err := convertAll(); err != nil {
		t.Fatal(err)
	}
}
