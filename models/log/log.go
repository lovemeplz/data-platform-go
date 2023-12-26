package log

import (
	"github.com/jinzhu/gorm"
	"github.com/lovemeplz/data-platform-go/models"
	"time"
)

type BizLog struct {
	models.Model
	UserCode    string `json:"userCode"`
	UserName    string `json:"userName"`
	Level       string `json:"level"`
	IP          string `json:"ip"`
	Path        string `json:"path"`
	OperateInfo string `json:"operateInfo"`
}

// TableName 会将表名重写
func (BizLog) TableName() string {
	return "dpg_biz_log"
}

func GetBizLog(pageNum int, pageSize int, maps interface{}) (bizlogs []BizLog) {
	models.Db.AutoMigrate(&BizLog{})
	models.Db.Table("dpg_biz_log").Where(maps).Offset(pageNum).Limit(pageSize).Find(&bizlogs)
	return
}

func GetBizLogTotal(maps interface{}) (count int) {
	models.Db.Model(&BizLog{}).Where(maps).Count(&count)
	return
}

func (bizlog *BizLog) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (bizlog *BizLog) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
