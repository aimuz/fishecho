package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PageIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "", map[string]interface{}{
		"title": "测试标题",
	})
}
