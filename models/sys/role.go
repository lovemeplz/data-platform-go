package sys

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lovemeplz/data-platform-go/models"
	"strings"
	"time"
)

type Role struct {
	models.Model
	RoleCode      string `json:"roleCode" validate:"required"`
	RoleName      string `json:"roleName" validate:"required"`
	AssignedUsers string `json:"assignedUsers"`
	DataStatus    *int   `json:"dataStatus" validate:"required"`
	Remark        string `json:"remark"`
	CreatedBy     string `json:"createdBy"`
	ModifiedBy    string `json:"modifiedBy"`
}

type RoleReq struct {
	models.Model
	ID int `gorm:"primary_key" json:"id"`
}

// TableName 会将表名重写
func (Role) TableName() string {
	return "dpg_sys_role"
}

func GetRole(pageNum int, pageSize int, maps interface{}) (roles []Role) {
	models.Db.AutoMigrate(&Role{})
	models.Db.Table("dpg_sys_role").Where(maps).Offset(pageNum).Limit(pageSize).Find(&roles)
	return
}

func GetRoleTotal(maps interface{}) (count int) {
	// models.Db.Model(&Role{}).Where(maps).Count(&count)
	fmt.Println("test::::")

	return 11
}

func AddRole(role Role) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Create(&role)
	return true
}

func UpdateRole(id int, role Role) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Model(&Role{}).Where("id = ?", id).Updates(&role)
	return true
}

func DeleteRole(id string) bool {
	models.Db.AutoMigrate(&Role{})
	roles := strings.Split(id, ",")
	models.Db.Delete(&Role{}, roles)

	//var roles = []Role{{ID: 15}, {ID: 14}}
	//models.Db.Delete(&roles)
	return true
}

func (role *Role) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (role *Role) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func Test() (roles []Role) {
	//models.Db.AutoMigrate(&Role{})
	//models.Db.Table("dpg_sys_role").Offset(10).Limit(0).Find(&roles)
	return
}
