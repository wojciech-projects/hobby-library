package main

import (
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func parse(reader io.Reader) Title {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	// volumeCountText := doc.Find("#collection-masthead__size").First().Text()
	// fmt.Println(volumeCountText)
	// volumeCount, err := parseVolumeCount(volumeCountText)

	title := doc.Find("title").First().Text()
	fmt.Println(title)

	parsedTitle, err := extractTitle(title)
	if err != nil {
		log.Fatal(err)
	}

	return parsedTitle
}
