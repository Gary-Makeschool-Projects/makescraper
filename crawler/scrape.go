package crawler

import (
	"github.com/gocolly/colly"
)

// URLchecker checks to see if the url is valid if not returns the valid version of the url
func URLchecker(url string) string {
	// if the url is valid then return the url
	hi := "hi"
	// otherwise
	return hi

}

// Crawl is responsible for the sraping logic
func Crawl(url string, keywords string) []string {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}
	// Instantiate default collector
	c := colly.NewCollector(
	// todo alloed domains
	// colly.AllowedDomains(urls or something),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		knownUrls = append(knownUrls, link)
		// c.Visit(e.Request.AbsoluteURL(link))
	})
	// Start scraping on
	c.Visit(url)

	return knownUrls
}
