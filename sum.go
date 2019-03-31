package arrays

// Sum returns the sum of numbers in a slice
func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
