package main

import (
	"strings"

	"golang.org/x/net/html"
)

//Link holds information about a html anchor element
type Link struct {
	Href string
	Text string
}

//NewLink returns a pointer to Link with given href and text
func NewLink(href, text string) *Link {
	return &Link{
		Href: href,
		Text: text,
	}
}

//LinkFromNode build a Link from html.*Node and returns a pointer to it
func LinkFromNode(node *html.Node) *Link {
	texts := make([]string, 0)
	var href string
	fn := func(k *html.Node) {
		if k.Type == html.TextNode {
			texts = append(texts, strings.TrimFunc(k.Data, Trimmer))
		} else if k.Type == html.ElementNode && k.Data == "a" && href == "" {
			for _, attr := range k.Attr {
				if attr.Key == "href" {
					href = strings.TrimFunc(attr.Val, Trimmer)
				}
			}
		}
	}
	VisitAll(DepthFirst, node, fn)
	return &Link{
		Href: href,
		Text: strings.TrimFunc(strings.Join(texts, " "), Trimmer),
	}
}
