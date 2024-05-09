package main

import (
	"context"

	"github.com/gotexxx/searchapp/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	linkHandler := handler.LinkHandler{} // Use LinkHandler struct
	app.Use(withLink)
	app.GET("/link", linkHandler.HandleLinkShow)
	app.Start(":8080")
}

func withLink(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "link", "example.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return (next(c))
	}
}
