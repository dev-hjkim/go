package main

import "testing"

func TestHello(t *testing.T) {
	// got := Hello("Chris")
	// want := "Hello, Chris"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()	// 에러가 난 라인 알려준다
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// test case 1
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})


	// test case 2
	t.Run("say 'Hello, World' when an empty string is supplied0", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}