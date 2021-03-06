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
	// todo make into a model
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
	// bind the model with the context body
	er := c.Bind(u)
	// panic!
	if er != nil {
		panic(err)
	}
	// crawl with the passed in data
	r := crawler.Crawl(u.URL, u.Keywords)
	fmt.Print(r)
	// return the links
	return c.JSON(http.StatusCreated, r)
}
