package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/lovemeplz/data-platform-go/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	pageNum, _ := com.StrTo(c.Query("pageNum")).Int()
	if pageNum > 0 {
		result = (pageNum - 1) * setting.PageSize
	}

	return result
}
