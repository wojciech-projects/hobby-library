package main

import (
	"errors"
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

func Test_extractTitle(t *testing.T) {
	t.Run("empty string returns error", func(t *testing.T) {
		_, err := extractTitle("")

		got := err.Error()
		want := errors.New("title parse error").Error()

		if got != want {
			t.Errorf("got %q want %q\n", got, want)
		}
	})

	t.Run("title without count returns error", func(t *testing.T) {
		_, err := extractTitle("社外取締役　島耕作 Kindle Edition")

		got := err.Error()
		want := errors.New("title parse error").Error()

		if got != want {
			t.Errorf("got %q want %q\n", got, want)
		}
	})

	t.Run("title with count can be parsed", func(t *testing.T) {
		got, err := extractTitle("社外取締役　島耕作 (5 book series) Kindle Edition")

		want := Title{title: "社外取締役　島耕作", volumeCount: 5}

		if err != nil {
			t.Errorf("got unexpected error %q\n", err)
		}

		assertTitle(t, got, want)
	})
}

func Test_extractVolumeCount(t *testing.T) {
	t.Run("empty string returns error", func(t *testing.T) {
		_, err := extractVolumeCount("")

		got := err.Error()
		want := errors.New("volume parse error").Error()

		if got != want {
			t.Errorf("got %q want %q\n", got, want)
		}
	})

	t.Run("volume count can be parsed", func(t *testing.T) {
		got, err := extractVolumeCount("            44 Volumes | 1 Episode")
		want := 44

		if err != nil {
			t.Errorf("got unexpected error %q\n", err)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
