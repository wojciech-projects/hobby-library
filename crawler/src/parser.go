package main

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func isType2(doc *goquery.Document) bool {
	return doc.Find("#collection-masthead__size").Length() == 1
}

func parseType2Doc(doc *goquery.Document) (title Title) {
	volumeCountText := doc.Find("#collection-masthead__size").First().Text()
	volumeCount, err := extractVolumeCount(volumeCountText)

	if err != nil {
		log.Fatal(err)
	}

	parsedTitle, err := parseTitle(doc)

	if err != nil {
		log.Fatal(err)
	}

	title.title = parsedTitle.title
	title.volumeCount = volumeCount

	return
}

func parseTitle(doc *goquery.Document) (title Title, err error) {
	titleText := doc.Find("title").First().Text()

	return extractTitle(titleText)
}

func parse(reader io.Reader) (title Title) {
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		log.Fatal(err)
	}

	if isType2(doc) {
		return parseType2Doc(doc)
	}

	title, err = parseTitle(doc)

	if err != nil {
		log.Fatal(err)
	}

	return title
}
