package main

import (
	"crawler/src/usecase"
	"fmt"
	"log"
)

func main() {
	shima := "B0C493GKLX"
	uchu := "B09478443K"
	favorites := []string{shima, uchu}

	repo := usecase.DummyAbstractRepository{}
	downloader := usecase.HttpMangaDownloader{}
	events, err := usecase.MangaCheckLatestUsecase(favorites, &repo, &downloader)

	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Printf("%s [%s] newest = %d\n", event.Title, event.Tag, event.NewestVolume)
	}
}
