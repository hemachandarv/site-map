package sitemap

import (
	"errors"
	"net/url"
	"strings"
)

type hostData struct {
	base string
	seen map[string]bool
	err  error
}

func host(urlName string) (h *hostData) {
	h = &hostData{
		seen: make(map[string]bool),
	}
	u, err := url.Parse(urlName)
	h.err = err
	if h.err != nil {
		return
	}
	baseURL := &url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
	}
	h.base = baseURL.String()
	if h.base == "" {
		h.err = errors.New("invalid hostname")
		return
	}
	return
}

func (h *hostData) link(l Link) (url string) {
	switch {
	case strings.HasPrefix(l.Href, "/"):
		url = h.base + l.Href
	case strings.HasPrefix(l.Href, "http"):
		url = l.Href
	}
	return url
}

func (h *hostData) visit(url string) bool {
	switch {
	case !strings.HasPrefix(url, h.base), url == "", h.seen[url]:
		return false
	default:
		return true
	}
}
