package main

import "fmt"

type ParseResult struct {
	length int
}

func parse(contents string) (result ParseResult) {
	result.length = len(contents)
	return
}

func main() {
	fmt.Println("Hello, world!")
}
