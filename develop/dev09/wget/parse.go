package wget

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

type Link struct {
	Href *url.URL
}

// Parses given a html file and returns a slice of links
func ParseHTML(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)

	var parseNode func(node *html.Node)
	parseNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			u, err := parseHref(n.Attr)
			if err != nil {
				return
			}

			link := Link{
				Href: u,
			}

			links = append(links, link)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}

	parseNode(doc)

	return links, nil
}

// Extracts href attributes from link tags
func parseHref(attrs []html.Attribute) (*url.URL, error) {
	var href string

	for _, a := range attrs {
		if a.Key == "href" {
			href = a.Val
			break
		}
	}

	u, err := url.Parse(href)
	if err != nil {
		return nil, err
	}

	return u, nil
}
