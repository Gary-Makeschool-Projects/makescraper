package main

import (
	"net/http"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"github.com/imthaghost/makescraper/controllers"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		description := "Simple website scraping API"
		return c.String(http.StatusOK, description)
	})
	e.POST("/scrape", controllers.CreateURL)
	// Server
	e.Logger.Fatal(e.Start(":8000"))
}
