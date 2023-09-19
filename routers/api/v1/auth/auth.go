package auth

import (
	"fmt"
	"github.com/lovemeplz/data-platform-go/pkg/logging"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/models/auth"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/utils"
)

type authUp struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := authUp{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	fmt.Println("a:::::", a)
	if ok {
		isExist := auth.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func LogOut(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetUserInfo(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["name"] = "Serati Ma"
	data["avatar"] = "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png"
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
