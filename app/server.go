package app

import (
	"github.com/labstack/echo"

	"github.com/imthaghost/makescraper/controllers"
)

// StartAPI ...
func StartAPI() {
	// instantiate
	e := echo.New()
	// post request
	e.POST("/scrape", controllers.CreateURL)
	// Server
	e.Logger.Fatal(e.Start(":8000"))
}
