package main

import (
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	dat, err := os.Open("examples/shima.html")
	Check(err)

	got := parse(dat)
	want := 1

	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}
