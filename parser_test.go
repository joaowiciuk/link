package main

import (
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestParseLinks(t *testing.T) {
	type args struct {
		doc *html.Node
	}
	tests := []struct {
		name      string
		args      args
		wantLinks []*Link
	}{
		{
			name: "Example 1",
			args: args{
				doc: DocFromFile("./testdata/ex1.html"),
			},
			wantLinks: []*Link{NewLink("/other-page", "A link to another page")},
		},
		{
			name: "Example 2",
			args: args{
				doc: DocFromFile("./testdata/ex2.html"),
			},
			wantLinks: []*Link{
				NewLink("https://www.twitter.com/joncalhoun", "Check me out on twitter"),
				NewLink("https://github.com/gophercises", "Gophercises is on Github !"),
			},
		},
		{
			name: "Example 3",
			args: args{
				doc: DocFromFile("./testdata/ex3.html"),
			},
			wantLinks: []*Link{
				NewLink("#", "Login"),
				NewLink("/lost", "Lost? Need help?"),
				NewLink("https://twitter.com/marcusolsson", "@marcusolsson"),
			},
		},
		{
			name: "Example 4",
			args: args{
				doc: DocFromFile("./testdata/ex4.html"),
			},
			wantLinks: []*Link{
				NewLink("/dog-cat", "dog cat"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLinks := ParseLinks(tt.args.doc); !reflect.DeepEqual(gotLinks, tt.wantLinks) {
				t.Errorf("ParseLinks() = %v, want %v", gotLinks, tt.wantLinks)
			}
		})
	}
}
