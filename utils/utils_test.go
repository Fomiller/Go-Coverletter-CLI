package utils

import (
	"testing"
)

func TestAppendIfMissing(t *testing.T) {
	list := []string{"apple", "orange", "banana", "kiwi", "grape"}
	xs := AppendIfMissing(list, "peach")
	if xs[5] != "peach" {
		t.Error("Expected", "peach", "Got", xs[5])
	}
}
