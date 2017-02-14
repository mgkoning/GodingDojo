package main

import "testing"

func TestGetGreeting(t *testing.T) {
	var greeting = getGreeting()
	var expected = "Hello, world!"
	if greeting != expected {
		t.Errorf("Expected '%v' but got '%v'\n", expected, greeting)
	}
}
