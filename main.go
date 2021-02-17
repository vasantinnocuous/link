package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)


func processElement(index int, element *goquery.Selection) {
 
    href, exists := element.Attr("href")
    if exists {
        fmt.Println(href)
    }
}

func main() {
    // Make HTTP request
    response, err := http.Get("https://www.google.com")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Create a goquery document from the HTTP response
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    // Find all links and process them with the function
  
    document.Find("a").Each(processElement)
}