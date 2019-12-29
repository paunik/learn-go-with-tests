package arrayandslices

// Sum sums slice of integers
func Sum(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}

	return
}

// SumAll sums slice of slices
func SumAll(numberOfSums ...[]int) []int {
	sums := make([]int, len(numberOfSums))
	for i, numbers := range numberOfSums {
		sums[i] = Sum(numbers)
	}

	return sums
}

// SumAllTails sums all tails of slices
func SumAllTails(numberOfSums ...[]int) []int {
	sums := make([]int, len(numberOfSums))
	for i, numbers := range numberOfSums {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}

	return sums
}
