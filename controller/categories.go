package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PageCategories(c echo.Context) error {
	data := newPageData()

	data["Title"] = "Aimuz 的小站"
	data["TitleFormat"] = "xxx %s xxx"
	data["Keywords"] = ""
	data["description"] = ""
	data["options"] = getOptions()
	data["menus"] = []map[string]interface{}{
		map[string]interface{}{
			"name":      "标签",
			"url":       "tags.html",
			"iconClass": "fa-home fa-fw",
		},
	}
	return c.Render(http.StatusOK, "categories.html", data)
}
