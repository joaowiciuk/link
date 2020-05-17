package main

import "golang.org/x/net/html"

type Algorithm int

const (
	// expand the visit frontier in breadth
	BreadthFirst Algorithm = iota

	// expand the visit frontier in depth
	DepthFirst
)

//VisitAll travels the node three applying fn to each node
func VisitAll(algorithm Algorithm, root *html.Node, fn func(*html.Node)) {
	if algorithm == BreadthFirst {
		bfs(root, fn)
		return
	}
	dfs(root, fn)
}

func dfs(root *html.Node, fn func(*html.Node)) {
	visited := make(map[*html.Node]interface{})
	queue := make([]*html.Node, 0)
	queue = append(queue, root)
	fn(root)
	visited[root] = nil
	for len(queue) > 0 {
		front := queue[0]
		for inBreadth := front.FirstChild; inBreadth != nil; inBreadth = inBreadth.NextSibling {
			if _, ok := visited[inBreadth]; !ok {
				fn(inBreadth)
				visited[inBreadth] = nil
				queue = append(queue, inBreadth)
			}
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

func bfs(root *html.Node, fn func(*html.Node)) {
	visited := make(map[*html.Node]interface{})
	queue := make([]*html.Node, 0)
	queue = append(queue, root)
	fn(root)
	visited[root] = nil
	for len(queue) > 0 {
		front := queue[0]
		for inDepth := front.FirstChild; inDepth != nil; inDepth = inDepth.FirstChild {
			if _, ok := visited[inDepth]; !ok {
				fn(inDepth)
				visited[inDepth] = nil
				queue = append(queue, inDepth)
			}
			for inBreadth := front.FirstChild; inBreadth != nil; inBreadth = inBreadth.NextSibling {
				if _, ok := visited[inBreadth]; !ok {
					fn(inBreadth)
					visited[inBreadth] = nil
					queue = append(queue, inBreadth)
				}
			}
		}
		queue = queue[1:]
	}
}
