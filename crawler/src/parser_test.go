package main

import (
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	t.Run("shima kousaku", func(t *testing.T) {
		dat, err := os.Open("examples/shima.html")
		Check(err)

		got := parse(dat)
		want := Title{title: "社外取締役　島耕作", volumeCount: 5}

		assertTitle(t, want, got)
	})

	// t.Run("uchuu kyoudai", func(t *testing.T) {
	// 	dat, err := os.Open("examples/uchuu_kyoudai.html")
	// 	Check(err)

	// 	got := parse(dat)
	// 	want := Title{title: "宇宙兄弟", volumeCount: 5}

	// 	assertTitle(t, want, got)
	// })
}
