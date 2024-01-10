package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lovemeplz/data-platform-go/models/sys"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	util "github.com/lovemeplz/data-platform-go/utils"
	"net/http"
	"strconv"
)

// GetUser
//
//	@Tags			用户管理
//	@Summary		获取用户列表
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200		{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/user [get]
func GetUser(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if c.Query("code") != "" {
		maps["code"] = c.Query("code")
	}
	if c.Query("name") != "" {
		maps["name"] = c.Query("name")
	}
	if c.Query("state") != "" {
		maps["state"] = c.Query("state")
	}

	data["list"] = sys.GetUser(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = sys.GetUserTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"success": true,
		"data":    data,
	})
}

// AddUser
//
//	@Tags			用户管理
//	@Summary		新增用户
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/user [post]
func AddUser(c *gin.Context) {
	var user sys.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("user", user, &user)
		// UnderscoreToUpperCamelCase
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	code := e.InvalidParams

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.AddUser(user)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

// UpdateUser
//
//	@Tags			用户管理
//	@Summary		更新用户
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/user/:id [put]
func UpdateUser(c *gin.Context) {
	var user sys.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	validate := validator.New()
	err := validate.Struct(user)
	code := e.InvalidParams

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.UpdateUser(id, user)
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

// DeleteUser
//
//	@Tags			用户管理
//	@Summary		删除用户
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/user/:id [delete]
func DeleteUser(c *gin.Context) {
	//code := e.SUCCESS
	//ID := c.Param("id")
	//id, _ := strconv.Atoi(ID)
	//
	//sys.DeleteRole(id)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": "",
	//})
}
