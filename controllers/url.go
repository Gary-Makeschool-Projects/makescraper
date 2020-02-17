package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	"github.com/imthaghost/makescraper/crawler"
)

//Site is data about the site we will scrape
type Site struct {
	URL      string `json:"url" form:"url" query:"url"`
	Keywords string `json:"keywords" form:"keywords" query:"keywords"`
}

//CreateURL mimics creating a user
func CreateURL(c echo.Context) (err error) {
	// Read the Body content
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(Site)
	// lets start crawling
	r := crawler.Crawl(u.URL)
	er := c.Bind(r)
	fmt.Print(er)
	fmt.Print(u)
	return c.JSON(http.StatusCreated, r)
}
