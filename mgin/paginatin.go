package mgin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 修改全局 limit 大小
const DefaultLimit = 20

// GetPage 保证了各接口的page处理是一致的
func GetPage(c *gin.Context) int {
	result := 0

	page, _ := cast.ToIntE(c.Param("page"))
	if page > 0 {
		result = (page - 1) * DefaultLimit
	}
	return result
}

// GetPageLimit 自定义单个请求 limit 大小
func GetPageLimit(c *gin.Context, limit int) int {
	result := 0
	page, _ := cast.ToIntE(c.Param("page"))
	if page > 0 {
		result = (page - 1) * limit
	}
	return result
}
