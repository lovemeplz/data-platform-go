package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lovemeplz/data-platform-go/docs"
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

	{
		apiv1.GET("/sys/roles", sys.GetRoles)
		apiv1.POST("/sys/roles", sys.AddRoles)
		apiv1.PUT("/sys/roles/:id", sys.EditRoles)
		apiv1.DELETE("/sys/roles/:id", sys.DeleteRoles)
	}

	return r
}
