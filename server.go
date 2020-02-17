package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	livereload "github.com/mattn/echo-livereload"

	"github.com/imthaghost/makescraper/controllers"
)

func main() {
	e := echo.New()
	// Middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Route => handler
	// Static file handler
	e.Static("/", "assets")
	// html handler
	e.File("/", "ui/index.html")
	// e.GET("/", func(c echo.Context) error {
	// 	description := "Simple website scraping API"
	// 	return c.String(http.StatusOK, description)
	// })

	e.POST("/scrape", controllers.CreateURL)
	e.Use(livereload.LiveReload())
	// Server
	e.Logger.Fatal(e.Start(":8000"))
}
