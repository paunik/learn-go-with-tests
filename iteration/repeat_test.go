package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Repeat same char 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
