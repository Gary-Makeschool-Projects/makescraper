package crawler

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// Crawl is responsible for the sraping logic
func Crawl(url string, keywords string) {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}
	// // Instantiate default collector
	// c := colly.NewCollector(
	// 	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// 	colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	// 	colly.Async(true),
	// )

	// // On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	knownUrls = append(knownUrls, e.Text)
	// 	// Print link
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// 	// Visit link found on page
	// 	// Only those links are visited which are in AllowedDomains
	// 	// c.Visit(e.Request.AbsoluteURL(link))
	// })

	// // Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// // Start scraping on https://hackerspaces.org
	// c.Visit("https://hackerspaces.org/")
	// return knownUrls
	c := colly.NewCollector(
		colly.AllowedDomains("emojipedia.org"),
	)

	// Callback for when a scraped page contains an article element
	c.OnHTML("article", func(e *colly.HTMLElement) {
		isEmojiPage := false
		knownUrls = append(knownUrls, e.Text)
		// Extract meta tags from the document
		metaTags := e.DOM.ParentsUntil("~").Find("meta")
		metaTags.Each(func(_ int, s *goquery.Selection) {
			// Search for og:type meta tags
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				// Emoji pages have "article" as their og:type
				isEmojiPage = strings.EqualFold(content, "article")
			}
		})

		if isEmojiPage {
			// Find the emoji page title
			fmt.Println("Emoji: ", e.DOM.Find("h1").Text())
			// Grab all the text from the emoji's description
			fmt.Println(
				"Description: ",
				e.DOM.Find(".description").Find("p").Text())
		}
	})

	// Callback for links on scraped pages
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the linked URL from the anchor tag
		link := e.Attr("href")
		// Have our crawler visit the linked URL
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://emojipedia.org")
}
