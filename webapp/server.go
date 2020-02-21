package webapp

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	livereload "github.com/mattn/echo-livereload"
)

// StartWebapp ...
func StartWebapp() {
	g := echo.New()
	// logger
	g.Use(middleware.Logger())
	// Stream recovery
	g.Use(middleware.Recover())
	// Live reload
	g.Use(livereload.LiveReload())
	//CORS
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Route => handler
	// Static file handler
	g.Static("/", "assets")
	// html handler
	g.File("/", "ui/index.html")
	// Server
	g.Logger.Fatal(g.Start(":5000"))
}
