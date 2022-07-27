package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}

	text := doc.Find("h1,p").Text()

	re := regexp.MustCompile("\\s{2,}")
	fmt.Println(re.ReplaceAllString(text, "\n"))

	words := doc.Find("li").Map(func(i int, sel *goquery.Selection) string {
		return fmt.Sprintf("%d: %s", i+1, sel.Text())
	})
	for _, word := range words {
		fmt.Println(word)
	}
}
