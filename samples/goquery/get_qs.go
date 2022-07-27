package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	webPage := "https://stackoverflow.com/questions/tagged/raku"
	resp, err := http.Get(webPage)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".questions").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3").Text()
		fmt.Println(i+1, title)
	})
}
