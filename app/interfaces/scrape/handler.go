package scrape

import "github.com/PuerkitoBio/goquery"

type ScrapeHandler interface {
	Find(word string) string
	FindTransportation() string
	GetDoc(url string) *goquery.Document
}
