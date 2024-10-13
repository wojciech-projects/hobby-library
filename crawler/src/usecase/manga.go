package usecase

import (
	"crawler/src/domain"
)

func MangaCheckLatestUsecase(
	repository domain.MangaRepository,
	favorites domain.FavoritesRepository,
	downloader domain.MangaDownloader,
) (resultEvents []domain.MangaUpdatedEvent, err error) {
	favoriteMangaUuids := favorites.FetchFavoriteMangas()

	for _, uuid := range favoriteMangaUuids {
		series, downloadErr := downloader.DownloadMangaByUuid(uuid)

		if downloadErr != nil {
			return resultEvents, downloadErr
		}

		manga, ok := repository.FetchMangaByUuid(uuid)

		if ok && series.VolumeCount <= manga.VolumeCount {
			continue
		}

		updatedManga := series.ToManga(uuid)
		events := repository.AddManga(updatedManga)
		resultEvents = append(resultEvents, events...)
	}
	return
}
