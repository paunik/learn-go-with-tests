package arrayandslices

import "testing"

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("Initial test sum", func(t *testing.T) {
		var numbers []int
		numbers = []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15

		assertCorrectMessage(t, sum, expected)
	})
}
