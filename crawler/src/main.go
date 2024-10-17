package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func fetchLatestUpdates(favorites []string) []Manga {
	downloader := HttpMangaDownloader{}
	mangas := make([]Manga, len(favorites))

	ch := make(chan Manga)

	for _, uuid := range favorites {
		go CheckLatestVolume(uuid, &downloader, ch)
	}

	for i := range len(favorites) {
		manga := <-ch
		mangas[i] = manga
		fmt.Printf("%s | %s -> %d\n", manga.AmazonUuid, manga.Title, manga.VolumeCount)
	}

	return mangas
}

func handleLatestVolumes(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got request\n")

	content, _ := io.ReadAll(r.Body)
	body := string(content)

	var favorites []string
	_ = json.Unmarshal([]byte(body), &favorites)

	mangas := fetchLatestUpdates(favorites)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mangas)

	fmt.Printf("Request handled\n")
}

func main() {
	http.HandleFunc("/latest_volumes", handleLatestVolumes)
	fmt.Println("Waiting for requests...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
