package main

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Series struct {
	title        string
	volumeCount  int
	thumbnailUrl string
}

func parse(reader io.Reader) (series Series) {
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		log.Fatal(err)
	}

	header, err := parseHeader(doc)

	if err != nil {
		log.Fatal(err)
	}

	thumbnailUrl, err := parseThumbnail(doc)

	if err != nil {
		log.Fatal(err)
	}

	series.title = header.title
	series.volumeCount = header.volumeCount
	series.thumbnailUrl = thumbnailUrl

	return series
}
