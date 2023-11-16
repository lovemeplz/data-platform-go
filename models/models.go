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

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.Cfg.GetString("database.TYPE")
	dbName = setting.Cfg.GetString("database.NAME")
	user = setting.Cfg.GetString("database.USER")
	password = setting.Cfg.GetString("database.PASSWORD")
	host = setting.Cfg.GetString("database.HOST")
	tablePrefix = setting.Cfg.GetString("database.TABLE_PREFIX")

	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
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
