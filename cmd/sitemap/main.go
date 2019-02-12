package main

import (
	"fmt"

	"github.com/hemv/site-map/internal/app/sitemap"
)

func main() {
	url := ""
	output, err := sitemap.Build(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}
