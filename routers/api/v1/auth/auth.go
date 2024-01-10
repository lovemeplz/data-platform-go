package auth

import (
	"fmt"
	"github.com/lovemeplz/data-platform-go/models/auth"
	"github.com/lovemeplz/data-platform-go/pkg/logging"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/utils"
)

type Account struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {

	var account Account
	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	valid := validation.Validation{}
	a := Account{Username: account.Username, Password: account.Password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.InvalidParams
	fmt.Println("test::::", a)
	if ok {
		isExist := auth.CheckAuth(account.Username, account.Password)
		if isExist {
			token, err := utils.GenerateToken(account.Username, account.Password)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ErrorAuth
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
