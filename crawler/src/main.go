package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("examples/shima.html")
	if err != nil {
		log.Fatal(err)
	}

	parse(f)
}
