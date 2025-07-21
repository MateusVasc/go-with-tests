package gowithtests

import "testing"

func TestHello(t *testing.T) {
	got := hello("Hiraeth")
	want := "Hello, Hiraeth"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
