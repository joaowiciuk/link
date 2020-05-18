package main

import (
	"golang.org/x/net/html"
)

//ParseLinks returns all links from html.Node of type html.DocumentNode
func ParseLinks(doc *html.Node) (links []*Link) {
	if doc.Type != html.DocumentNode {
		panic("input node must be of type DocumentNode")
	}
	links = make([]*Link, 0)
	fn := func(k *html.Node) {
		if k.Type == html.ElementNode && k.Data == "a" {
			link := LinkFromNode(k)
			links = append(links, link)
		}
	}
	VisitAll(BreadthFirst, doc, fn)
	return
}
