package main

import "testing"

func TestMessage(t *testing.T) {
	want := "test message"
	got := message()
	if want != got {
		t.Errorf("message did not match, expected %s, got %s", want, got)
	}
}
