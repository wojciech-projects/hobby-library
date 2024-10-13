package usecase

import (
	"crawler/src/domain"
	"testing"
)

// Mock #1

type emptyFavorite string

func (*emptyFavorite) FetchFavoriteMangas() (uuids []domain.Uuid) {
	return
}

func TestMangaCheckLatestUsecase(t *testing.T) {
	t.Run("no favorites means no events", func(t *testing.T) {

		fav := emptyFavorite("123")
		got, _ := MangaCheckLatestUsecase(nil, &fav, nil)

		if len(got) != 0 {
			t.Errorf("expected [] got %v", got)
		}
	})

	// TODO: add more tests
}
