package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/lovemeplz/data-platform-go/models"
	"time"
)

type Role struct {
	models.Model
	Code           string `json:"code"`
	Name           string `json:"name"`
	AssignedPerson string `json:"assigned_person"`
	State          int    `json:"state"`
	Remark         string `json:"remark"`
	CreatedBy      string `json:"created_by"`
	ModifiedBy     string `json:"modified_by"`
}

func GetRoles(pageNum int, pageSize int, maps interface{}) (roles []Role) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&roles)
	return
}

func GetRolesTotal(maps interface{}) (count int) {
	models.Db.Model(&Role{}).Where(maps).Count(&count)

	return
}

func AddRoles(data map[string]interface{}) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Create(&Role{
		Code:           data["code"].(string),
		Name:           data["name"].(string),
		AssignedPerson: data["assigned_person"].(string),
		State:          data["state"].(int),
		Remark:         data["remark"].(string),
		CreatedBy:      data["created_by"].(string),
		ModifiedBy:     data["modified_by"].(string),
	})
	return true
}

func EditRoles(id int, data map[string]interface{}) bool {
	models.Db.AutoMigrate(&Role{})
	models.Db.Model(&Role{}).Where("id = ?", id).Updates(data)
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
