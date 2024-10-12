package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("examples/shima.html")
	if err != nil {
		log.Fatal(err)
	}

	length := parse(f)
	fmt.Println(length)
}
