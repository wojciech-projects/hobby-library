package domain

type Uuid string

type Manga struct {
	Title        string
	VolumeCount  int
	AmazonUuid   Uuid
	ThumbnailUrl string
}

type MangaUpdatedEvent struct {
	Title        string
	NewestVolume int
}

type MangaRepository interface {
	AddManga(manga Manga) (events []MangaUpdatedEvent)
	FetchMangaByUuid(uuid Uuid) (manga Manga, ok bool)
}

type FavoritesRepository interface {
	FetchFavoriteMangas() (uuids []Uuid)
}

type Series struct {
	Title         string
	VolumeCount   int
	ThumbnailUrl  string
	RelatedSeries []Uuid
}

func (series *Series) ToManga(uuid Uuid) (manga Manga) {
	manga.Title = series.Title
	manga.AmazonUuid = uuid
	manga.VolumeCount = series.VolumeCount
	manga.ThumbnailUrl = series.ThumbnailUrl
	return
}

type MangaDownloader interface {
	DownloadMangaByUuid(uuid Uuid) (series Series, err error)
}
