package main

import (
	"errors"
	"regexp"
	"strconv"
)

type Title struct {
	title       string
	volumeCount int
}

func extractTitle(title string) (parsedTitle Title, err error) {
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
