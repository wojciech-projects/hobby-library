package main

import (
	"crawler/src/usecase"
	"log"
)

func main() {
	favorites := []string{}
	repo := usecase.DummyAbstractRepository{}
	downloader := usecase.HttpMangaDownloader{}
	_, err := usecase.MangaCheckLatestUsecase(favorites, &repo, &downloader)

	if err != nil {
		log.Fatal(err)
	}
}
