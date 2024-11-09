package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("say greet with someone", func(t *testing.T) {
		buffer := bytes.Buffer{}

		// func Greet 會將內容準備好，存放在 Buffer(a.k.a Writer) 中
		Greet(&buffer, "Chris")

		want := "Hello, Chris"
		got := buffer.String()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
