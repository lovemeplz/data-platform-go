package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/lovemeplz/data-platform-go/models"
	"time"
)

// Role model info
// @Description Code 角色编码
// @Description Name 角色名称

type Role struct {
	models.Model
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	AssignedUser string `json:"assigned_user"`
	State        *int   `json:"state" validate:"required"`
	Remark       string `json:"remark"`
	CreatedBy    string `json:"created_by"`
	ModifiedBy   string `json:"modified_by"`
}

func GetRoles(pageNum int, pageSize int, maps interface{}) (roles []Role) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&roles)
	return
}

func GetRolesTotal(maps interface{}) (count int) {
	models.Db.Model(&Role{}).Where(maps).Count(&count)
	return
}

func AddRoles(role Role) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Create(&role)
	return true
}

func EditRoles(id int, role Role) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Model(&Role{}).Where("id = ?", id).Updates(&role)
	return true
}

func DeleteRoles(id int) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Where("id = ?", id).Delete(&Role{})
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
