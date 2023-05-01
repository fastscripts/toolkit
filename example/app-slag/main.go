package main

import (
	"log"

	"github.com/fastscripts/toolkit"
)

func main() {
	toSlug := "now is the time 123"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)

	if err != nil {
		log.Println(err)
	}

	log.Println(slugified)
}
