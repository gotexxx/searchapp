package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	linkHandler := handler.linkHandler{}
	app.GET("/link", linkHandler.HandleLinkShow)

	app.Start(":8080")
}
