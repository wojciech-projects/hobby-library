package usecase

import "crawler/src/domain"

type DummyAbstractRepository struct{}

func (d *DummyAbstractRepository) AddManga(manga domain.Manga) (events []domain.MangaEvent) {
	event := domain.MangaEvent{Tag: "new", Title: manga.Title, NewestVolume: manga.VolumeCount}
	return []domain.MangaEvent{event}
}

func (d *DummyAbstractRepository) FetchMangaByUuid(uuid string) (manga domain.Manga, ok bool) {
	return manga, false
}
