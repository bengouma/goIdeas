package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func processSite(index int, element *goquery.Selection) {
	div, exists := element.Attr("class")
	if exists {
		fmt.Println(div)
	}
}

func main() {
	response, err := http.Get("https://bgoumalatsos.com")
	if err != nil {
		fmt.Println("Error fetching site")
	}
	defer response.Body.Close()
	
	// response is an http.Response object and needs to be converted into a goquery document
	// so it can be manipulated and scraped
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Error creating document for goquery to parse")
	}
	
	doc.Find("div").Each(processSite)
}
