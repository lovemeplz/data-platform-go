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

// GetRole
//
//	@Tags			角色管理
//	@Summary		获取角色列表
//	@Description	这是一段接口描述
//	@Produce		json
//	@Param			code	query		string	false	"角色编码"
//	@Param			name	query		string	false	"角色名称"
//	@Param			state	query		int		false	"角色状态"
//	@Success		200		{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/role [get]
func GetRole(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if c.Query("roleCode") != "" {
		maps["role_code"] = c.Query("roleCode")
	}
	if c.Query("roleName") != "" {
		maps["role_name"] = c.Query("roleName")
	}
	if c.Query("dataStatus") != "" {
		maps["data_status"] = c.Query("dataStatus")
	}

	data["list"] = sys.GetRole(util.GetPage(c), setting.AppSetting.PageSize, maps)
	// data["list"] = sys.GetRole(1, setting.AppSetting.PageSize, maps)
	data["total"] = sys.GetRoleTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"success": true,
		"data":    data,
	})
}

// AddRole
//
//	@Tags			角色管理
//	@Summary		新增角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/role [post]
func AddRole(c *gin.Context) {
	var role sys.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(role)
	code := e.InvalidParams

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.AddRole(role)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

// UpdateRole
//
//	@Tags			角色管理
//	@Summary		更新角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/role/:id [put]
func UpdateRole(c *gin.Context) {
	var role sys.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	validate := validator.New()
	err := validate.Struct(role)
	code := e.InvalidParams

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.UpdateRole(id, role)
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

// DeleteRole
//
//	@Tags			角色管理
//	@Summary		删除角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/role/:id [delete]
func DeleteRole(c *gin.Context) {
	code := e.SUCCESS
	id := c.Param("id")
	// id, _ := strconv.Atoi(ID)
	sys.DeleteRole(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
