package hello

import "testing"

func TestString(t *testing.T) {
	result := "Hello World!"
	if String() != result {
		t.Error("Expected " + result + " got something else.")
	}
}
