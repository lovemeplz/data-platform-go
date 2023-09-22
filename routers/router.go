package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lovemeplz/data-platform-go/docs"
	"github.com/lovemeplz/data-platform-go/routers/api/v1/auth"
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
	// apiv1.Use(jwt.JWT())

	// 用户、登入登出相关
	{
		apiv1.POST("auth/login", auth.Login)
		apiv1.POST("auth/logout", auth.LogOut)
		apiv1.GET("auth/userInfo", auth.GetUserInfo)
	}

	// 系统管理
	{
		// 角色管理
		apiv1.GET("/sys/roles", sys.GetRoles)
		apiv1.POST("/sys/roles", sys.AddRoles)
		apiv1.PUT("/sys/roles/:id", sys.EditRoles)
		apiv1.DELETE("/sys/roles/:id", sys.DeleteRoles)

		// 部门管理
		apiv1.GET("/sys/dept", sys.GetDept)
		apiv1.POST("/sys/dept", sys.AddDept)
		apiv1.PUT("/sys/dept/:id", sys.UpdateDept)
		apiv1.DELETE("/sys/dept/:id", sys.DeleteDept)
	}

	// 错误日志
	{
		//apiv1.GET("/log/errorlog", log.getErrorLogs)
		//apiv1.POST("/sys/errorlog", log.exportErrorLogs)
		//apiv1.DELETE("/sys/errorlog/:id", log.DeleteErrorLogs)
	}

	return r
}
