package arrayandslices

func Sum(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}

	return
}

func SumAll(numberOfSums ...[]int) []int {
	sums := make([]int, len(numberOfSums))
	for i, numbers := range numberOfSums {
		sums[i] = Sum(numbers)
	}

	return sums
}

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
