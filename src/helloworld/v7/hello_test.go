package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello world to people", func(t *testing.T) {
		got := Hello("QingZhi", "")
		want := "Hello, World! QingZhi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world to people with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World! Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Cat", spanish)
		want := "Hola, Mundo! Cat"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Dog", french)
		want := "Bonjour, Monde! Dog"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
