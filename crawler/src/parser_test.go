package main

import (
	"errors"
	"os"
	"testing"
)

func TestTitleParser(t *testing.T) {
	t.Run("empty string returns error", func(t *testing.T) {
		_, err := parseTitle("")

		got := err.Error()
		want := errors.New("title parse error").Error()

		if got != want {
			t.Errorf("got %q want %q\n", got, want)
		}
	})

	t.Run("title without count returns error", func(t *testing.T) {
		_, err := parseTitle("社外取締役　島耕作 Kindle Edition")

		got := err.Error()
		want := errors.New("title parse error").Error()

		if got != want {
			t.Errorf("got %q want %q\n", got, want)
		}
	})

	t.Run("title with count can be parsed", func(t *testing.T) {
		got, err := parseTitle("社外取締役　島耕作 (5 book series) Kindle Edition")

		want := Title{title: "社外取締役　島耕作", volumeCount: 5}

		if err != nil {
			t.Errorf("got unexpected error %q\n", err)
		}

		assertTitle(t, got, want)
	})
}

func assertTitle(t testing.TB, want, got Title) {
	t.Helper()

	if got.title != want.title {
		t.Errorf("got %q want %q\n", got.title, want.title)
	}

	if got.volumeCount != want.volumeCount {
		t.Errorf("got %d want %d\n", got.volumeCount, want.volumeCount)
	}
}

func TestParsing(t *testing.T) {
	dat, err := os.Open("examples/shima.html")
	Check(err)

	got := parse(dat)
	want := Title{title: "社外取締役　島耕作", volumeCount: 5}

	assertTitle(t, want, got)
}
