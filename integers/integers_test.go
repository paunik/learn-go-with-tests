package integers

import (
	"fmt"
	"testing"
)

func assertCorrectResult(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}
func TestAdder(t *testing.T) {
	t.Run("first case", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectResult(t, sum, expected)
	})

	t.Run("2 + 6 case ", func(t *testing.T) {
		sum := Add(2, 6)
		expected := 8

		assertCorrectResult(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
