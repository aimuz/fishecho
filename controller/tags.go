package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PageTags(c echo.Context) error {
	data := newPageData()

	data["Title"] = "Aimuz 的小站"
	data["TitleFormat"] = "xxx %s xxx"
	data["Keywords"] = ""
	data["description"] = ""
	data["options"] = getOptions()
	data["menus"] = getMenu("tags")
	data["tagTotal"] = 10
	data["tags"] = []map[string]interface{}{
		map[string]interface{}{
			"name":  "C",
			"total": 10,
			"url":   "",
			"class": "tag-size-10",
		},
		map[string]interface{}{
			"name":  "Nginx",
			"total": 20,
			"url":   "",
			"class": "tag-size-15",
		},
		map[string]interface{}{
			"name":  "CPP",
			"total": 2,
			"url":   "",
			"class": "tag-size-5",
		},
		map[string]interface{}{
			"name":  "Golang",
			"total": 3,
			"url":   "",
			"class": "tag-size-5",
		},
		map[string]interface{}{
			"name":  "Go",
			"total": 10,
			"url":   "",
			"class": "tag-size-10",
		},
	}
	return c.Render(http.StatusOK, "tags.html", data)
}
