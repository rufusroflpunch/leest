package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rufusroflpunch/wicky/lists"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	api := e.Group("/api")
	api.GET("/lists", lists.GetLists)
	e.Logger.Fatal(e.Start(":8080"))
}
