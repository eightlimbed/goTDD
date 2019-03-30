package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(5, 6)
	expected := 11
	if sum != expected {
		t.Errorf("got '%d' want '%d'", sum, expected)
	}
}

// ExampleAdd will be available in godoc
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
