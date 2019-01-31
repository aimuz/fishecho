package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PageArchive(c echo.Context) error {
	data := newPageData()

	data["Title"] = "Aimuz 的小站"
	data["TitleFormat"] = "xxx %s xxx"
	data["Keywords"] = ""
	data["description"] = ""
	data["options"] = getOptions()
	data["menus"] = []map[string]interface{}{}
	return c.Render(http.StatusOK, "archive.html", data)
}
