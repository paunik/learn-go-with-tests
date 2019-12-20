package arrayandslices

func Sum(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}

	return
}
