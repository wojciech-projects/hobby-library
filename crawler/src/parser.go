package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Uuid string

type Series struct {
	title         string
	volumeCount   int
	thumbnailUrl  string
	relatedSeries []Uuid
}

func parse(reader io.Reader) (series Series, err error) {
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return
	}

	header, err := parseHeader(doc)

	if err != nil {
		return
	}

	thumbnailUrl, err := parseThumbnail(doc)

	if err != nil {
		return
	}

	relatedUuids, err := parseRelatedSeriesUuids(doc)

	if err != nil {
		return
	}

	series.title = header.title
	series.volumeCount = header.volumeCount
	series.thumbnailUrl = thumbnailUrl
	series.relatedSeries = relatedUuids

	return series, nil
}
