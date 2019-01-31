package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PageIndex(c echo.Context) error {
	page := c.Param("page")
	return c.Render(http.StatusOK, c.Param("page"), map[string]interface{}{
		"title": "测试标题",
	})
}
