package main

import "testing"

// Test suite for the Hello() function
func TestHello(t *testing.T) {

	// Since every subtest will compare 'got' with 'want', here is a helper
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	// Subtest #1
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lee", "")
		want := "Hello, Lee"
		assertCorrectMessage(t, got, want)
	})

	// Subtest #2
	t.Run("say 'Hello, World' when no name is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// Subtest #3
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Lee", "Spanish")
		want := "Hola, Lee"
		assertCorrectMessage(t, got, want)
	})

	// Subtest #4
	t.Run("in French", func(t *testing.T) {
		got := Hello("Lee", "French")
		want := "Bonjour, Lee"
		assertCorrectMessage(t, got, want)
	})

	// Subtest #5
	t.Run("in Swahili", func(t *testing.T) {
		got := Hello("Lee", "Swahili")
		want := "Hujambo, Lee"
		assertCorrectMessage(t, got, want)
	})
}
