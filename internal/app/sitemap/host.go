package sitemap

import (
	"errors"
	"log"
	"net/url"
)

type hostData struct {
	host    string
	visited map[string]bool
	err     error
}

func host(urlName string) (h *hostData) {
	h = &hostData{}
	u, err := url.Parse(urlName)
	if err != nil {
		return
	}
	h.host = u.Hostname()
	if h.host == "" {
		h.err = errors.New("invalid hostname")
		return
	}
	log.Printf("Host: %s\n", h.host)
	return
}
