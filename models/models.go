package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
)

var Db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"createdOn"`
	ModifiedOn int `json:"modifiedOn"`
}

func Setup() {
	DatabaseSetting := setting.DatabaseSetting

	dbType := DatabaseSetting.Type
	dbName := DatabaseSetting.Name
	user := DatabaseSetting.User
	password := DatabaseSetting.Password
	host := DatabaseSetting.Host
	tablePrefix := DatabaseSetting.TablePrefix

	Db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Db.Close()
}
