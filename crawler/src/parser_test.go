package main

import (
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	t.Run("shima kousaku", func(t *testing.T) {
		dat, err := os.Open("examples/shima.html")
		Check(err)

		got := parse(dat)
		want := Series{
			title:         "社外取締役　島耕作",
			volumeCount:   5,
			thumbnailUrl:  "https://m.media-amazon.com/images/I/B1f12VWkmIL._SY300_.png",
			relatedSeries: make([]Uuid, 60),
		}

		assertSeries(t, want, got)
	})
	t.Run("uchuu kyoudai", func(t *testing.T) {
		dat, err := os.Open("examples/uchuu_kyoudai.html")
		Check(err)

		got := parse(dat)
		want := Series{
			title:         "宇宙兄弟",
			volumeCount:   44,
			thumbnailUrl:  "https://m.media-amazon.com/images/I/612Z9SegTKL._SY300_.jpg",
			relatedSeries: make([]Uuid, 60),
		}

		assertSeries(t, want, got)
	})
}

func assertSeries(t testing.TB, want, got Series) {
	t.Helper()

	if got.title != want.title {
		t.Errorf("got %q want %q\n", got.title, want.title)
	}

	if got.volumeCount != want.volumeCount {
		t.Errorf("got %d want %d\n", got.volumeCount, want.volumeCount)
	}

	if got.thumbnailUrl != want.thumbnailUrl {
		t.Errorf("got %q want %q\n", got.thumbnailUrl, want.thumbnailUrl)
	}

	if len(got.relatedSeries) != len(want.relatedSeries) {
		t.Errorf("got %d want %d\n", len(got.relatedSeries), len(want.relatedSeries))
	}
}
