package auth

import (
	"github.com/lovemeplz/data-platform-go/models"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return "dpg_sys_user"
}

func CheckAuth(username, password string) bool {
	var auth Auth
	models.Db.Select("id").Where("user_name = ? AND password = ?", username, password).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
