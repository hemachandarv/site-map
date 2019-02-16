package sitemap

import (
	"encoding/xml"
	"net/http"
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
	h := host(urlName)
	link := []Link{}
	if h.err != nil {
		err = h.err
		return
	}
	resp, err := http.Get(urlName)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	clink, err := links(resp.Body)
	if err != nil {
		return
	}
	link = append(link, clink...)
	_ = link
	return
}
