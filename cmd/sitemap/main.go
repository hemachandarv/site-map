package main

import (
	"flag"
	"fmt"

	"github.com/hemv/site-map/internal/app/sitemap"
)

func main() {
	urlName := flag.String("url", "", "url to build sitemap")
	flag.Parse()
	if *urlName == "" {
		panic("url flag not provided")
	}
	output, err := sitemap.Build(*urlName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("***\tSITEMAP - %s\t***\n%s\n", *urlName, output)
}
