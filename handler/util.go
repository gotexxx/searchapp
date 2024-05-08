package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, components templ.Component) error {
	return components.Render(c.Request().Context(), c.Response())
}
