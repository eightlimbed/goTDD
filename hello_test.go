package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lee")
	want := "Hello, Lee"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
