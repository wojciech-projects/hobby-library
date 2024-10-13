package usecase

import (
	"crawler/src/domain"
)

func MangaCheckLatestUsecase(repository domain.MangaRepository) {
	favoriteMangas := repository.FetchFavoriteMangas()

	for manga := range favoriteMangas {

	}
}
