package domain

type Uuid string

type Manga struct {
	Title        string
	VolumeCount  int
	AmazonUuid   Uuid
	ThumbnailUrl string
	IsFavorite   bool
}

type MangaUpdatedEvent struct {
	Title        string
	NewestVolume int
}

type MangaRepository interface {
	AddManga(manga Manga) (events []MangaUpdatedEvent)
	FetchFavoriteMangas() (mangas []Manga)
}

type Series struct {
	Title         string
	VolumeCount   int
	ThumbnailUrl  string
	RelatedSeries []Uuid
}

type MangaDownloader interface {
	fetchMangaByUuid(uuid Uuid) (series Series, err error)
}
