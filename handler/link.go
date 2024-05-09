package handler

import (
	"github.com/gotexxx/searchapp/model"
	"github.com/gotexxx/searchapp/view/link"
	"github.com/labstack/echo/v4"
)

type LinkHandler struct{}

func (h LinkHandler) HandleLinkShow(c echo.Context) error {
	l := model.Link{
		Url: "test.com",
	}

	return render(c, link.Show(l))
}
