package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people when an empty strinh is given for both language", func(t *testing.T) {
		got := hello("Hiraeth", "")
		want := "Hello, Hiraeth"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when an empty string is given for both name and language", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to people in spanish", func(t *testing.T) {
		got := hello("Hiraeth", "Spanish")
		want := "Hola, Hiraeth"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to people in french", func(t *testing.T) {
		got := hello("Hiraeth", "French")
		want := "Bonjour, Hiraeth"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to people with an invalid language", func(t *testing.T) {
		got := hello("Hiraeth", "Japanese")
		want := "Hello, Hiraeth"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
