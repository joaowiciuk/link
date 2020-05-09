package main

import (
	"net/http"
	"os"
	"unicode"

	"golang.org/x/net/html"
)

//Trimmer helper function to trim spaces and tabs from text
func Trimmer(r rune) bool {
	return unicode.IsControl(r) || unicode.IsSpace(r)
}

//DocFromFile loads html.DocumentNode from local file
func DocFromFile(path string) (doc *html.Node) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	doc, err = html.Parse(file)
	if err != nil {
		panic(err)
	}
	return
}

//DocFromURL loads html.DocumentNode from external URL
func DocFromURL(url string) (doc *html.Node) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	doc, err = html.Parse(response.Body)
	if err != nil {
		panic(err)
	}
	return
}
