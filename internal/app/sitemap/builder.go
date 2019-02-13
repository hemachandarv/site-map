package sitemap

import (
	"encoding/xml"
)

// Sitemap represents the XML sitemap data
// for a given URL.
type Sitemap struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	Location  []Link
}

// Build accepts and url and returns a byte slice
// which represents the sitemap of the url.
func Build(urlName string) (b []byte, err error) {
	h, err := host(urlName)
	_ = h
	return
}
