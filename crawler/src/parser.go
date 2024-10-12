package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Title struct {
	title       string
	volumeCount int
}

func parseTitle(title string) (parsedTitle Title, err error) {
	regex := *regexp.MustCompile(`(.+) \((\d+) book series\).*`)
	results := regex.FindStringSubmatch(title)
	if len(results) < 3 {
		return parsedTitle, errors.New("title parse error")
	}

	parsedTitle.title = results[1]
	parsedTitle.volumeCount, err = strconv.Atoi(results[2])

	if err != nil {
		return parsedTitle, errors.New("title parse error")
	}
	return
}

func parseVolumeCount(text string) (volumeCount int, err error) {
	regex := *regexp.MustCompile(`\s*(\d+) Volumes.+`)
	results := regex.FindStringSubmatch(text)
	if len(results) < 2 {
		return volumeCount, errors.New("volume parse error")
	}

	volumeCount, err = strconv.Atoi(results[1])

	if err != nil {
		return volumeCount, errors.New("title parse error")
	}
	return
}

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

	parsedTitle, err := parseTitle(title)
	if err != nil {
		log.Fatal(err)
	}

	return parsedTitle
}
