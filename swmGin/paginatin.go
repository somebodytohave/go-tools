package swmGin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// GetPage 保证了各接口的page处理是一致的
func GetPage(c *gin.Context) int {
	result := 0

	page, _ := cast.ToIntE(c.Param("page"))
	if page > 0 {
		result = (page - 1) * 20
	}
	return result
}
