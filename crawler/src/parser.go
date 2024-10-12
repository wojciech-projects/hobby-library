package main

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func parse(reader io.Reader) (length int) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	return doc.Length()
}
