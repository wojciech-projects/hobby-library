package main

import (
	"crawler/src/domain"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func parse(reader io.Reader) (series domain.Series, err error) {
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

	series.Title = header.title
	series.VolumeCount = header.volumeCount
	series.ThumbnailUrl = thumbnailUrl
	series.RelatedSeries = relatedUuids

	return series, nil
}
