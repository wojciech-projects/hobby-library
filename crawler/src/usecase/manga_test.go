package usecase

import (
	"crawler/src/domain"
	"testing"
)

func TestMangaCheckLatestUsecase(t *testing.T) {
	t.Run("no favorites means no events", func(t *testing.T) {

		got, _ := MangaCheckLatestUsecase([]domain.Uuid{}, nil, nil)

		if len(got) != 0 {
			t.Errorf("expected [] got %v", got)
		}
	})

	// TODO: add more tests
}
