package main

import "testing"

func TestHello(t *testing.T){
	// got := Hello("Chris")
	// want := "Hello, Chris"
	
	// if got != want {
	// 	t.Errorf("want '%q' got '%q'", want, got)
	// }

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Jerry")
		want := "Hello, Jerry"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an enmpty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}