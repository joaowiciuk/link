package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestLinkFromNode(t *testing.T) {
	type args struct {
		node *html.Node
	}
	tests := []struct {
		name string
		args args
		want *Link
	}{
		{
			name: "test 1",
			args: args{
				node: &html.Node{
					Type: html.ElementNode,
					Data: "a",
					Attr: []html.Attribute{{Key: "href", Val: "/other-page"}},
					FirstChild: &html.Node{
						Type: html.TextNode,
						Data: "A link to another page",
					},
				},
			},
			want: &Link{
				Href: "/other-page",
				Text: "A link to another page",
			},
		},
		{
			name: "test 2",
			args: args{
				node: &html.Node{
					Type: html.ElementNode,
					Data: "a",
					Attr: []html.Attribute{{Key: "href", Val: "https://www.twitter.com/joncalhoun"}},
					FirstChild: &html.Node{
						Type: html.TextNode,
						Data: "Check me out on twitter",
					},
				},
			},
			want: &Link{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
		},
		{
			name: "test 3",
			args: args{
				node: func() *html.Node {
					s := `<a href="https://github.com/gophercises">
					  Gophercises is on <strong>Github</strong>!
					</a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			want: &Link{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github !",
			},
		},
		{
			name: "test 4",
			args: args{
				node: func() *html.Node {
					s := `<a href="#" class="btn btn-login">Login <i class="fa fa-sign-in" aria-hidden="true"></i></a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			want: &Link{
				Href: "#",
				Text: "Login",
			},
		},
		{
			name: "test 5",
			args: args{
				node: func() *html.Node {
					s := `<a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>`
					doc, _ := html.Parse(strings.NewReader(s))
					return doc.FirstChild.FirstChild.NextSibling.FirstChild
				}(),
			},
			want: &Link{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkFromNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkFromNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
