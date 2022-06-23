package main

import "testing"

func TestHello(t *testing.T) {
	correctingMessages := func(t testing.TB, want, get string) {
		// t.Helper()
		if want != get {
			t.Errorf("Expected %s Got %s ", want, get)
		}
	}
	t.Run("defualt for empyt string 'World'", func(t *testing.T) {
		want := "Hola, World"
		get := Hello(" ", "spanish")
		if want != get {
			correctingMessages(t, want, get)
		}
	})
	t.Run("Say hello to the people ", func(t *testing.T) {
		want := "Hello, Alazar"
		got := Hello("Alazar", "English")
		if want != got {
			correctingMessages(t, want, got)
		}
	})
}
