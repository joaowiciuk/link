package main

import "golang.org/x/net/html"

type Algorithm int

const (
	// BreadthFirst expands the visited frontier in breadth first
	BreadthFirst Algorithm = iota

	// DepthFirst expands the visited frontier in depth first
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
	stack := make([]*html.Node, 0)
	visited := make(map[*html.Node]interface{})
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[node]; !ok {
			fn(node)
			visited[node] = nil
			for child := node.LastChild; child != nil; child = child.PrevSibling {
				stack = append(stack, child)
			}
		}
	}
}

func bfs(root *html.Node, fn func(*html.Node)) {
	queue := make([]*html.Node, 0)
	visited := make(map[*html.Node]interface{})
	fn(root)
	visited[root] = nil
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if _, ok := visited[child]; !ok {
				fn(child)
				visited[child] = nil
				queue = append(queue, child)
			}
		}
	}
}
