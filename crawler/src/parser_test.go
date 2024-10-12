package main

import (
	"errors"
	"os"
	"testing"
)

func assertTitle(t testing.TB, want, got Title) {
	t.Helper()

	if got.title != want.title {
		t.Errorf("got %q want %q\n", got.title, want.title)
	}

	if got.volumeCount != want.volumeCount {
		t.Errorf("got %d want %d\n", got.volumeCount, want.volumeCount)
	}
}

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

func TestVolumeCountParser(t *testing.T) {
	t.Run("empty string returns error", func(t *testing.T) {
		_, err := parseVolumeCount("")

		got := err.Error()
		want := errors.New("volume parse error").Error()

		if got != want {
			t.Errorf("got %q want %q\n", got, want)
		}
	})

	t.Run("volume count can be parsed", func(t *testing.T) {
		got, err := parseVolumeCount("            44 Volumes | 1 Episode")
		want := 44

		if err != nil {
			t.Errorf("got unexpected error %q\n", err)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestParsing(t *testing.T) {
	t.Run("shima kousaku", func(t *testing.T) {
		dat, err := os.Open("examples/shima.html")
		Check(err)

		got := parse(dat)
		want := Title{title: "社外取締役　島耕作", volumeCount: 5}

		assertTitle(t, want, got)
	})

	// t.Run("uchuu kyoudai", func(t *testing.T) {
	// 	dat, err := os.Open("examples/uchuu_kyoudai.html")
	// 	Check(err)

	// 	got := parse(dat)
	// 	want := Title{title: "宇宙兄弟", volumeCount: 5}

	// 	assertTitle(t, want, got)
	// })
}
