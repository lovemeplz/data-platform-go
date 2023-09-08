package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/models/sys"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	util "github.com/lovemeplz/data-platform-go/utils"
	"net/http"
)

func GetRoles(c *gin.Context) {
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

func EditRoles(c *gin.Context) {

}

func DeleteRoles(c *gin.Context) {

}
