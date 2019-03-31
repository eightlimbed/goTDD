package arrays

// Sum returns the sum of numbers in a slice
func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

// SumAll returns the sum of a slice of slices
func SumAll(data ...[]int) (sums []int) {
	for _, numbers := range data {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumAllTails returns the sum of a slice of slices, without the first slice
func SumAllTails(data ...[]int) (sums []int) {
	for _, numbers := range data {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return
}
