package main

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func parse(reader io.Reader) (header Header) {
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		log.Fatal(err)
	}

	header, err = parseHeader(doc)

	if err != nil {
		log.Fatal(err)
	}

	return header
}
