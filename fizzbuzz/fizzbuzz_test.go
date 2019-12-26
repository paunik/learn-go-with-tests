package fizzbuzz

import "testing"

func TestFizz(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		got, _ := FizzBuzz(1)

		want := "1"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("one two", func(t *testing.T) {
		got, _ := FizzBuzz(2)

		want := "12"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("one two three", func(t *testing.T) {
		got, _ := FizzBuzz(3)

		want := "12Fizz"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("one until 10", func(t *testing.T) {
		got, _ := FizzBuzz(10)

		want := "12Fizz4BuzzFizz78FizzBuzz"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("one until 15", func(t *testing.T) {
		got, _ := FizzBuzz(15)

		want := "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("fail minus one", func(t *testing.T) {
		_, got := FizzBuzz(-1)

		want := ErrInvalidInput

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FizzBuzz(100)
	}
}
