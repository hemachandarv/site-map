package sitemap

import (
	"errors"
	"log"
	"net/url"
)

type hostData struct {
	host    string
	visited map[string]bool
}

var h = hostData{}

func host(urlName string) (name string, err error) {
	u, err := url.Parse(urlName)
	if err != nil {
		return
	}
	h.host = u.Hostname()
	if h.host == "" {
		err = errors.New("invalid hostname")
		return
	}
	log.Printf("Host: %s\n", h.host)
	return
}
