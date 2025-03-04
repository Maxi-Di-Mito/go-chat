package main

import (
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	if resp, err := http.Get("https://dolarhoy.com/cotizaciondolarbolsa"); err == nil {
		defer resp.Body.Close()

		if rootNode, err := html.Parse(resp.Body); err == nil {
		}
	}
}

func findLinks(n *html.Node) (links []string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c)...)
	}
	return links
}
