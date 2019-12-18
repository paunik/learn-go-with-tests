package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to named people", func(t *testing.T) {
		got := Hello("Test", "English")
		want := "Hello, Test"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Servus, Klaus' when zGerman", func(t *testing.T) {
		got := Hello("Klaus", "German")
		want := "Servus, Klaus"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Serbian", func(t *testing.T) {
		got := Hello("Milivoje", "Serbian")
		want := "Pomaze Bog, Milivoje"
		assertCorrectMessage(t, got, want)
	})
}
