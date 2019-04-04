package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lee")

	got := buffer.String()
	want := "Hello, Lee"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
