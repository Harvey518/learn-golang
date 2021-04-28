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

	// 当传入人名时，返回Hello, world
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Jerry", "")
		want := "Hello, Jerry"
		assertCorrectMessage(t, got, want)
	})

	// 当传入人名为空时，name=world
	t.Run("say hello world when an enmpty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	// 添加第二个参数指定问候语，如果不能识别这种语言则默认为英语
	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Jerry", "Spanish")
		want := "Hola, Jerry"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Frech", func(t *testing.T) {
		got := Hello("Jerry", "Frech")
		want := "Bonjour, Jerry"
		assertCorrectMessage(t, got, want)
	})

	t.Run("unknow language", func(t *testing.T) {
		got := Hello("Jerry", "Chinese")
		want := "Hello, Jerry"
		assertCorrectMessage(t, got, want)
	})
}