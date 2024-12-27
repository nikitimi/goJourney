package deferrer

import "testing"

func TestDeferrer(t *testing.T) {
	result := "End of deferrer"
	if Deferrer("") != result {
		t.Error("Expecting to return " + result)
	}
}
