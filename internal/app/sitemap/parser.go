package sitemap

import "io"

// Link represents a HTML link "<a href='..'>...</a>"
// and stores the href attribute.
type Link struct {
	Href string `xml:"url>loc"`
	Text string `xml:"-"`
}

func parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
