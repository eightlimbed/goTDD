package arrays

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {

	t.Run("a collection of one slice", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4})
		want := []int{6, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checker := func(got, want []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("a collection of 4 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8}, []int{9})
		want := []int{5, 11, 8, 0}
		checker(got, want, t)
	})

	t.Run("handle an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{}, []int{7, 8}, []int{9})
		want := []int{5, 0, 8, 0}
		checker(got, want, t)
	})
}
