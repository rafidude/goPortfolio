package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getLinks() {
	webPage := "https://golang.org"
	resp, err := http.Get(webPage)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
	}

	doc.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {
		link, _ := tag.Attr("href")
		linkText := tag.Text()
		fmt.Printf("%s %s\n", linkText, link)
	})
}

func main() {
	getLinks()
}
