package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, _ := http.Get("https://finance.yahoo.com/quote/ORCL")
	// Load the HTML document
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// Find the review items
	doc.Find("#quote-header-info").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the name.
		name := s.Find("h1").Text()
		fmt.Println(name)
	})
}

func main() {
	ExampleScrape()
}
