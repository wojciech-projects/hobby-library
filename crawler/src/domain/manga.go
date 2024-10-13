package domain

type Manga struct {
	Title        string
	VolumeCount  int
	AmazonUuid   string
	ThumbnailUrl string
}

type MangaEvent struct {
	Tag          string // "new" | "updated"
	Title        string
	NewestVolume int
}

type MangaRepository interface {
	AddManga(manga Manga) (events []MangaEvent)
	FetchMangaByUuid(uuid string) (manga Manga, ok bool)
}

type FavoritesRepository interface {
	FetchFavoriteMangas() (uuids []string)
}

type Series struct {
	Title         string
	VolumeCount   int
	ThumbnailUrl  string
	RelatedSeries []string
}

func (series *Series) ToManga(uuid string) (manga Manga) {
	manga.Title = series.Title
	manga.AmazonUuid = uuid
	manga.VolumeCount = series.VolumeCount
	manga.ThumbnailUrl = series.ThumbnailUrl
	return
}

type MangaDownloader interface {
	DownloadMangaByUuid(uuid string) (series Series, err error)
}
