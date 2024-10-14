package helloworld_test

import (
	"go-tdd/helloworld"
	"testing"
)

func TestHello(t *testing.T) {
	got := helloworld.Hello("Yuki")
	want := "Hello, Yuki!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
