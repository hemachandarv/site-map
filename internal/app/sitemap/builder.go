package sitemap

import "encoding/xml"

// Sitemap represents the XML sitemap data
// for a given URL.
type Sitemap struct {
	Header    string
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	Location  []Link
}

// Build accepts and url and returns a byte slice
// which represents the sitemap of the url.
func Build(url string) ([]byte, error) {
	return nil, nil
}
