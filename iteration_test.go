package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("x", 5)
	want := "xxxxx"
	if got != want {
		t.Errorf("got '%s' want'%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A", 20)
	}
}

func ExampleRepeat() {
	result := Repeat("x", 20)
	fmt.Println(result)
	// Output: xxxxxxxxxxxxxxxxxxxx
}
