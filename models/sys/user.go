package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/lovemeplz/data-platform-go/models"
	"time"
)

type User struct {
	models.Model
	DeptCode    string `json:"dept_code" validate:"required"`
	UserCode    string `json:"user_code"`
	UserName    string `json:"user_name"`
	Gender      uint
	State       int `json:"state" validate:"required"`
	Email       string
	AccountFrom string
	Remark      string `json:"remark"`

	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

//type SysUser struct {
//	global.GVA_MODEL
//	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
//	Username    string         `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
//	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
//	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
//	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
//	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
//	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
//	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
//	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
//	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
//	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
//	Phone       string         `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
//	Email       string         `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
//	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
//}

// TableName 会将表名重写
func (User) TableName() string {
	return "dpg_sys_user"
}

func GetUser(pageNum int, pageSize int, maps interface{}) (user []User) {
	models.Db.AutoMigrate(&User{})
	models.Db.Table("dpg_sys_user").Where(maps).Offset(pageNum).Limit(pageSize).Find(&user)
	return
}

func GetUserTotal(maps interface{}) (count int) {
	models.Db.Model(&User{}).Where(maps).Count(&count)
	return
}

func AddUser(user User) bool {
	models.Db.AutoMigrate(&User{})
	models.Db.Create(&user)
	return true
}

func UpdateUser(id int, user User) bool {
	models.Db.AutoMigrate(&User{})
	models.Db.Model(&User{}).Where("id = ?", id).Updates(&user)
	return true
}

func DeleteUser(id int) bool {
	models.Db.AutoMigrate(&User{})
	models.Db.Where("id = ?", id).Delete(&User{})
	return true
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
