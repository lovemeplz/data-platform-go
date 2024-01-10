package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lovemeplz/data-platform-go/docs"
	"github.com/lovemeplz/data-platform-go/routers/api/v1/auth"
	"github.com/lovemeplz/data-platform-go/routers/api/v1/example"
	"github.com/lovemeplz/data-platform-go/routers/api/v1/log"
	"github.com/lovemeplz/data-platform-go/routers/api/v1/sys"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("api/v1")

	// 白名单
	{
		apiv1.POST("auth/login", auth.Login)
		apiv1.POST("auth/logout", auth.LogOut)

	}

	// apiv1.Use(jwt.JWT())
	// 用户相关
	{
		apiv1.GET("auth/userInfo", auth.GetUserInfo)
	}

	// 系统管理
	{
		// 角色管理
		apiv1.GET("/sys/role", sys.GetRole)
		apiv1.POST("/sys/role", sys.AddRole)
		apiv1.PUT("/sys/role/:id", sys.UpdateRole)
		apiv1.DELETE("/sys/role/:id", sys.DeleteRole)

		// 部门管理
		apiv1.GET("/sys/dept", sys.GetDept)
		apiv1.POST("/sys/dept", sys.AddDept)
		apiv1.PUT("/sys/dept/:id", sys.UpdateDept)
		apiv1.DELETE("/sys/dept/:id", sys.DeleteDept)

		// 用户管理
		apiv1.GET("/sys/user", sys.GetUser)
		apiv1.POST("/sys/user", sys.AddUser)
		apiv1.PUT("/sys/user/:id", sys.UpdateUser)
		apiv1.DELETE("/sys/user/:id", sys.DeleteUser)
	}

	// 业务日志
	{
		apiv1.GET("/log/bizlog", log.GetBizLog)
		apiv1.GET("/log/bizlog/export", log.ExportBizLog)
	}

	// 错误日志
	{
		//apiv1.GET("/log/errorlog", log.getErrorLogs)
		//apiv1.POST("/sys/errorlog", log.exportErrorLogs)
		//apiv1.DELETE("/sys/errorlog/:id", log.DeleteErrorLogs)
	}

	// 示例
	{
		apiv1.POST("/example/upload", example.Upload)
		apiv1.GET("/example/qrcode", example.Qrcode)
	}

	return r
}
