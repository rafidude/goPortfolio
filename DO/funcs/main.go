package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// function to fetch http url body
func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	//call fetch function
	body, err := fetch("https://www.example.com")
	if err != nil {
		log.Fatal(err)
	}
	//print body
	fmt.Println(body)
	a := `Say "hello" to 'Go'!`
	fmt.Println(a)
}
