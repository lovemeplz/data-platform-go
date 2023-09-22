package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/lovemeplz/data-platform-go/models"
	"time"
)

type Dept struct {
	models.Model
	DeptCode       string `json:"dept_code" validate:"required"`
	DeptName       string `json:"dept_name"`
	DeptParentCode string `json:"dept_parent_code"`
	OrderNo        uint   `json:"order_no"`
	State          int    `json:"state" validate:"required"`
	Remark         string `json:"remark"`
	CreatedBy      string `json:"created_by"`
	ModifiedBy     string `json:"modified_by"`
}

// TableName 会将表名重写
func (Dept) TableName() string {
	return "dpg_sys_dept"
}

func GetDept(maps interface{}) (dept []Dept) {
	models.Db.Table("dpg_sys_dept").Find(&dept)
	return
}

func AddDept(dept Dept) bool {
	models.Db.AutoMigrate(&Dept{})
	models.Db.Create(&dept)
	return true
}

func UpdateDept(id int, dept Dept) bool {
	models.Db.AutoMigrate(&Dept{})
	models.Db.Model(&Dept{}).Where("id = ?", id).Updates(&dept)
	return true
}

func DeleteDept(id int) bool {
	models.Db.AutoMigrate(&Dept{})
	models.Db.Where("id = ?", id).Delete(&Dept{})
	return true
}

func (dept *Dept) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (dept *Dept) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
