package main

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Header struct {
	title       string
	volumeCount int
}

func extractTitle(title string) (header Header, err error) {
	regex := *regexp.MustCompile(`(.+) \((\d+) book series\).*`)
	results := regex.FindStringSubmatch(title)
	if len(results) < 3 {
		return header, errors.New("title parse error")
	}

	header.title = results[1]
	header.volumeCount, err = strconv.Atoi(results[2])

	if err != nil {
		return header, errors.New("title parse error")
	}
	return
}

func extractVolumeCount(text string) (volumeCount int, err error) {
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

func isHeaderType2(doc *goquery.Document) bool {
	return doc.Find("#collection-masthead__size").Length() == 1
}

func parseHeaderType1(doc *goquery.Document) (header Header, err error) {
	title := doc.Find("title").First().Text()

	return extractTitle(title)
}

func parseHeaderType2(doc *goquery.Document) (header Header, err error) {
	volumeCountText := doc.Find("#collection-masthead__size").First().Text()
	volumeCount, err := extractVolumeCount(volumeCountText)

	if err != nil {
		return header, err
	}

	parsedTitle, err := parseHeaderType1(doc)

	if err != nil {
		return header, err
	}

	header.title = parsedTitle.title
	header.volumeCount = volumeCount

	return
}

func parseHeader(doc *goquery.Document) (header Header, err error) {
	if isHeaderType2(doc) {
		return parseHeaderType2(doc)
	} else {
		return parseHeaderType1(doc)
	}
}
