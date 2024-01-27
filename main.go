package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

func readHTMLContent(url string) (string, error) {
	var htmlContent string

	c := colly.NewCollector()

	c.OnHTML("html", func(e *colly.HTMLElement) {
		htmlContent = e.Text
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	url := os.Getenv("URL")
	htmlContent, err := readHTMLContent(url)

	if err != nil {
		log.Fatalf("Error reading HTML content: %v", err)
	}

	fmt.Printf("HTML Content of %s:\n%s\n", url, htmlContent)
}
