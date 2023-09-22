package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lovemeplz/data-platform-go/models/sys"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"net/http"
	"strconv"
)

// GetRoles
//
//	@Tags			部门管理
//	@Summary		获取部门
//	@Description	这是一段接口描述
//	@Produce		json
//	@Router			/api/v1/sys/roles [get]

func GetDept(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	data["list"] = sys.GetDept(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"success": true,
		"data":    data,
	})
}

// AddDept
//
//	@Tags			部门管理
//	@Summary		新增部门
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/dept [post]
func AddDept(c *gin.Context) {
	var dept sys.Dept
	if err := c.BindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(dept)
	code := e.INVALID_PARAMS

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.AddDept(dept)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

// UpdateDept
//
//	@Tags			部门管理
//	@Summary		更新部门
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/dept/:id [put]
func UpdateDept(c *gin.Context) {
	var dept sys.Dept
	if err := c.BindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	validate := validator.New()
	err := validate.Struct(dept)
	code := e.INVALID_PARAMS

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	sys.UpdateDept(id, dept)
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

// DeleteDept
//
//	@Tags			部门管理
//	@Summary		删除部门
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/dept/:id [delete]
func DeleteDept(c *gin.Context) {
	code := e.SUCCESS
	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	sys.DeleteDept(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
