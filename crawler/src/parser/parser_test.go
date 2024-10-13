package parser

import (
	"crawler/src/domain"
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	t.Run("shima kousaku", func(t *testing.T) {
		dat, err := os.Open("examples/shima.html")
		Check(err)

		got, _ := Parse(dat)
		want := domain.Series{
			Title:         "社外取締役　島耕作",
			VolumeCount:   5,
			ThumbnailUrl:  "https://m.media-amazon.com/images/I/B1f12VWkmIL._SY300_.png",
			RelatedSeries: make([]string, 60),
		}

		assertSeries(t, want, got)
	})
	t.Run("uchuu kyoudai", func(t *testing.T) {
		dat, err := os.Open("examples/uchuu_kyoudai.html")
		Check(err)

		got, _ := Parse(dat)
		want := domain.Series{
			Title:         "宇宙兄弟",
			VolumeCount:   44,
			ThumbnailUrl:  "https://m.media-amazon.com/images/I/612Z9SegTKL._SY300_.jpg",
			RelatedSeries: make([]string, 60),
		}

		assertSeries(t, want, got)
	})
}

func assertSeries(t testing.TB, want, got domain.Series) {
	t.Helper()

	if got.Title != want.Title {
		t.Errorf("got %q want %q\n", got.Title, want.Title)
	}

	if got.VolumeCount != want.VolumeCount {
		t.Errorf("got %d want %d\n", got.VolumeCount, want.VolumeCount)
	}

	if got.ThumbnailUrl != want.ThumbnailUrl {
		t.Errorf("got %q want %q\n", got.ThumbnailUrl, want.ThumbnailUrl)
	}

	if len(got.RelatedSeries) != len(want.RelatedSeries) {
		t.Errorf("got %d want %d\n", len(got.RelatedSeries), len(want.RelatedSeries))
	}
}
