package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoles(c *gin.Context) {
	//code := c.Query("code")
	//name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "3233",
		"data": "111",
	})
}

func AddRoles(c *gin.Context) {

}

func EditRoles(c *gin.Context) {

}

func DeleteRoles(c *gin.Context) {

}
