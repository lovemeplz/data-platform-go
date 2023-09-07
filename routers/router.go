package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/routers/api/v1/sys"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("api/v1")

	{
		apiv1.GET("/sys/roles", sys.GetRoles)
		apiv1.POST("/sys/roles", sys.AddRoles)
		apiv1.PUT("/sys/roles/:id", sys.EditRoles)
		apiv1.DELETE("/sys/roles/:id", sys.DeleteRoles)
	}

	return r
}
