package handler

import (
	"searchApp/view/link"

	"github.com/labstack/echo/v4"
)

type linkHandler struct{}

func (h linkHandler) HandleLinkShow(c echo.Context) error {
	return render(c, link.Show())
}
