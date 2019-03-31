package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("a collection of any size", func(t *testing.T) {
		numbers := []int{420, 69, 177}
		got := Sum(numbers)
		want := 666
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}
