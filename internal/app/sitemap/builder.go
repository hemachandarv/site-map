package sitemap

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

// Sitemap represents the XML sitemap data
// for a given URL.
type Sitemap struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	URL       []Link   `xml:"url"`
}

// Build accepts and url and returns a byte slice
// which represents the sitemap of the url.
func Build(urlName string) ([]byte, error) {
	h := host(urlName)
	if h.err != nil {
		return nil, h.err
	}
	links := get(urlName, h)
	s := &Sitemap{
		Namespace: xmlns,
		URL:       links,
	}
	buf := bytes.NewBufferString("")
	enc := xml.NewEncoder(buf)
	enc.Indent("", "  ")
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func get(url string, h *hostData) (links []Link) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	h.seen[url] = true
	clinks, _ := parse(resp.Body)
	for _, l := range clinks {
		u := h.link(l)
		if h.visit(u) {
			tlinks := get(u, h)
			links = append(links, tlinks...)
		}
	}
	return
}
