package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lovemeplz/data-platform-go/models/sys"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	util "github.com/lovemeplz/data-platform-go/utils"
	"net/http"
	"strconv"
)

// GetRoles
//
//	@Title			测试
//	@Tags			系统管理
//	@Summary		获取角色列表
//	@Description	这是一段接口描述
//	@Produce		json
//	@Header			200		{string}	Token	"token"
//	@Param			code	query		string	false	"角色编码"
//	@Param			name	query		string	false	"角色名称"
//	@Param			state	query		int		false	"角色状态"
//	@Router			/api/v1/sys/roles [get]

func GetRoles(c *gin.Context) {
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

	data["list"] = sys.GetRoles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = sys.GetRolesTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"success": true,
		"data":    data,
	})
}

// AddRoles
//
//	@Title			测试
//	@Tags			系统管理
//	@Summary		新增角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/roles [post]
func AddRoles(c *gin.Context) {
	var role sys.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(role)
	code := e.INVALID_PARAMS

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.AddRoles(role)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

// EditRoles
//
//	@Tags			系统管理
//	@Summary		编辑角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/roles/:id [put]
func EditRoles(c *gin.Context) {
	var role sys.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	validate := validator.New()
	err := validate.Struct(role)
	code := e.INVALID_PARAMS

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.EditRoles(id, role)
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

// DeleteRoles
//
//	@Tags			系统管理
//	@Summary		删除角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/roles/:id [delete]
func DeleteRoles(c *gin.Context) {
	code := e.SUCCESS
	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	sys.DeleteRoles(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
