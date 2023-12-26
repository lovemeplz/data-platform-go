package log

import (
	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/models/log"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	util "github.com/lovemeplz/data-platform-go/utils"
	"net/http"
)

func GetBizLog(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if c.Query("userCode") != "" {
		maps["user_code"] = c.Query("userCode")
	}
	if c.Query("userName") != "" {
		maps["user_name"] = c.Query("userName")
	}
	if c.Query("ip") != "" {
		maps["ip"] = c.Query("ip")
	}

	data["list"] = log.GetBizLog(util.GetPage(c), setting.PageSize, maps)
	data["total"] = log.GetBizLogTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"success": true,
		"data":    data,
	})
}
