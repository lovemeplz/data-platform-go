package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/models/sys"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	util "github.com/lovemeplz/data-platform-go/utils"
	"net/http"
	"strconv"
)

// GetRoles
// @Title       测试
// @Tags        系统管理
// @Summary     获取角色列表
// @Description 这是一段接口描述
// @Produce     json
// @Success     200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router      /api/v1/sys/roles [get]
func GetRoles(c *gin.Context) {

	//id := com.StrTo(c.Param("id")).MustInt()
	//
	//valid := validation.Validation{}
	//valid.Min(id, 1, "id").Message("ID必须大于0")
	//
	//code := e.INVALID_PARAMS
	//var data interface{}
	//if !valid.HasErrors() {
	//	if models.ExistArticleByID(id) {
	//		data = models.GetArticle(id)
	//		code = e.SUCCESS
	//	} else {
	//		code = e.ERROR_NOT_EXIST_ARTICLE
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
	//	}
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	//code := c.Query("code")
	//name := c.Query("name")

	data["lists"] = sys.GetRoles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = sys.GetRolesTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddRoles
// @Title       测试
// @Tags 系统管理
// @Summary 新增角色
// @Description 这是一段接口描述
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/sys/roles [post]
func AddRoles(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})

	Code := c.Query("code")
	Name := c.Query("name")
	AssignedPerson := c.Query("assigned_person")
	State := c.Query("state")
	Remark := c.Query("remark")
	CreatedBy := c.Query("created_by")
	ModifiedBy := c.Query("modified_by")

	data["code"] = Code
	data["name"] = Name
	data["state"] = State
	data["assigned_person"] = AssignedPerson
	data["remark"] = Remark
	data["created_by"] = CreatedBy
	data["modified_by"] = ModifiedBy

	sys.AddRoles(data)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditRoles
// @Tags 系统管理
// @Summary 编辑角色
// @Description 这是一段接口描述
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/sys/roles/:id [put]
func EditRoles(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})

	ID := c.Param("id")
	Name := c.Query("name")
	AssignedPerson := c.Query("assigned_person")

	data["name"] = Name
	data["assigned_person"] = AssignedPerson

	id, _ := strconv.Atoi(ID)
	sys.EditRoles(id, data)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteRoles
// @Tags 系统管理
// @Summary 删除角色
// @Description 这是一段接口描述
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/sys/roles/:id [delete]
func DeleteRoles(c *gin.Context) {
	code := e.SUCCESS
	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)

	sys.DeleteRoles(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
