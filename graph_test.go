package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisitAll(t *testing.T) {
	type args struct {
		algorithm Algorithm
		root      *html.Node
	}
	tests := []struct {
		name string
		args args
		got  []string
		want []string
	}{
		{
			name: "Breadth First Search 1",
			args: args{
				algorithm: BreadthFirst,
				root: func() *html.Node {
					s := `<a href="https://github.com/gophercises">Gophercises is on <strong>Github</strong>!</a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			got:  make([]string, 0),
			want: []string{"a", "Gophercises is on", "strong", "!", "Github"},
		},
		{
			name: "Breadth First Search 2",
			args: args{
				algorithm: BreadthFirst,
				root: func() *html.Node {
					s := `<a>Login <i>(click to <b>login</b>)</i><i>#</i></a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			got:  make([]string, 0),
			want: []string{"a", "Login", "i", "i", "(click to", "b", ")", "#", "login"},
		},
		{
			name: "Depth First Search 1",
			args: args{
				algorithm: DepthFirst,
				root: func() *html.Node {
					s := `<a href="https://github.com/gophercises">Gophercises is on <strong>Github</strong>!</a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			got:  make([]string, 0),
			want: []string{"a", "Gophercises is on", "strong", "Github", "!"},
		},
		{
			name: "Depth First Search 2",
			args: args{
				algorithm: DepthFirst,
				root: func() *html.Node {
					s := `<a>Login <i>(click to <b>login</b>)</i><i>#</i></a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			got:  make([]string, 0),
			want: []string{"a", "Login", "i", "(click to", "b", "login", ")", "i", "#"},
		},
		{
			name: "Depth First Search 3",
			args: args{
				algorithm: DepthFirst,
				root: func() *html.Node {
					s := `<div><a href="https://www.twitter.com/joncalhoun">Check me out on twitter<i class="fa fa-twitter" aria-hidden="true"></i></a><a href="https://github.com/gophercises">Gophercises is on <strong>Github</strong>!</a></div>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			got:  make([]string, 0),
			want: []string{"div", "a", "Check me out on twitter", "i", "a", "Gophercises is on", "strong", "Github", "!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(node *html.Node) {
				tt.got = append(tt.got, strings.TrimFunc(node.Data, Trimmer))
			}
			VisitAll(tt.args.algorithm, tt.args.root, fn)
			if len(tt.got) != len(tt.want) {
				t.Fatalf("got %v, want %v", tt.got, tt.want)
			}
			for i := range tt.got {
				if tt.got[i] != tt.want[i] {
					t.Fatalf("got %v, want %v", tt.got, tt.want)
				}
			}
		})
	}
}
