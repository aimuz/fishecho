package router

import (
	"fishecho/controller"
	"github.com/labstack/echo"
)

func InitRouter(g *echo.Group) {
	// 首页
	g.GET("/", controller.PageIndex)
	g.GET("index.html", controller.PageIndex)
	g.GET("tags.html", controller.PageTags)
	g.GET("categories.html", controller.PageCategories)
	g.GET("archive.html", controller.PageArchive)
}
