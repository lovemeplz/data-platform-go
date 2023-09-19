package log

import (
	"github.com/lovemeplz/data-platform-go/models"
)

type Log struct {
	models.Model
	UserCode   string `json:"user_code"`
	UserName   string `json:"user_name"`
	Type       string `json:"type"`
	Time       int    `json:"time"`
	IP         string `json:"ip"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}
