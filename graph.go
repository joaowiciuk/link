package main

import "golang.org/x/net/html"

//VisitAll travels the node three applying fn to each node
func VisitAll(root *html.Node, fn func(*html.Node)) {
	visited := make(map[*html.Node]interface{})
	queue := make([]*html.Node, 0)
	queue = append(queue, root)
	fn(root)
	visited[root] = nil
	for len(queue) > 0 {
		front := queue[0]
		// expand the visit frontier in breadth
		for inBreadth := front.FirstChild; inBreadth != nil; inBreadth = inBreadth.NextSibling {
			if _, ok := visited[inBreadth]; !ok {
				fn(inBreadth)
				visited[inBreadth] = nil
				queue = append(queue, inBreadth)
			}
			// expand the visit frontier in depth
			for inDepth := inBreadth.FirstChild; inDepth != nil; inDepth = inDepth.FirstChild {
				if _, ok := visited[inDepth]; !ok {
					fn(inDepth)
					visited[inDepth] = nil
					queue = append(queue, inDepth)
				}
			}
		}
		queue = queue[1:]
	}
}
