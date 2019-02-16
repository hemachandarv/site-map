package sitemap

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Link represents a HTML link "<a href='..'>...</a>"
// and stores the href attribute.
type Link struct {
	Href string `xml:"url>loc"`
	Text string `xml:"-"`
}

func links(r io.Reader) (link []Link, err error) {
	n, err := html.Parse(r)
	if err != nil {
		return
	}
	link = build(anchors(n))
	log.Printf("%v\n", link)
	return
}

func anchors(n *html.Node) (a []*html.Node) {
	if n.Type == html.ElementNode && n.DataAtom == atom.A {
		a = append(a, n)
	}
	for next := n.FirstChild; next != nil; next = next.NextSibling {
		a = append(a, anchors(next)...)
	}
	return
}

func build(a []*html.Node) (link []Link) {
	for _, n := range a {
		l := Link{}
		t := ""
		for _, attr := range n.Attr {
			if attr.Key == atom.Href.String() {
				l.Href = attr.Val
				break
			}
		}
		for next := n.FirstChild; next != nil; next = next.NextSibling {
			if next.Type == html.TextNode {
				t += next.Data
			}
		}
		l.Text = strings.Join(strings.Fields(t), " ")
		link = append(link, l)
	}
	return
}
