package controller

import (
	"time"
)

func newPageData() map[string]interface{} {
	return map[string]interface{}{
		"now": time.Now(), // 当前时间
	}
}

// 获取主题设置
func getOptions() map[string]interface{} {
	return map[string]interface{}{
		"siteURL":   "https://aimuz.me",
		"staticURL": "https://static.aimuz.me",
	}
}

// 获取菜单列表
func getMenu(pageType string) []map[string]interface{} {
	menus := []map[string]interface{}{
		map[string]interface{}{
			"name":     "首页",
			"icon":     "fa-home",
			"class":    "menu-item-home",
			"pageType": "index",
		},
		map[string]interface{}{
			"name":     "标签",
			"icon":     "fa-home",
			"class":    "menu-item-home",
			"pageType": "tags",
		},
	}

	for k := range menus {
		if menus[k]["pageType"] == pageType {
			menus[k]["class"] = menus[k]["class"].(string) + " menu-item-active"
		}
	}

	return menus
}
