package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		correctMessage(t, got, want)
	})
	t.Run("Empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		correctMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jose", "ES")
		want := "Hola, Jose"

		correctMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Jacques", "FR")
		want := "Bonjour, Jacques"

		correctMessage(t, got, want)
	})
}

func correctMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
