package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	livereload "github.com/mattn/echo-livereload"

	"github.com/imthaghost/makescraper/controllers"
)

func main() {
	// API
	// go app.StartAPI()
	// // WebAPP
	// go webapp.StartWebapp()
	// // This is used for closing the bot using various different termination signals.
	// // Ben showed me this fuckery
	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	e := echo.New()
	// logger
	e.Use(middleware.Logger())
	// Stream recovery
	e.Use(middleware.Recover())
	// Live reload
	e.Use(livereload.LiveReload())
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
	// Route => handler
	e.POST("/scrape", controllers.CreateURL)
	// Server
	e.Logger.Fatal(e.Start(":8000"))
}
