package main

import (
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestParsing(t *testing.T) {
	dat, err := os.ReadFile("shima.html")
	check(err)

	got := parse(string(dat))
	want := 481376

	if got.length != want {
		t.Errorf("got %d want %d\n", got.length, want)
	}
}
