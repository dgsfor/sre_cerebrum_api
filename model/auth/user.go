package auth

import (
	"gorm.io/gorm"
	"ssopa/model"
	"ssopa/util"
)

type SsoPaUsers struct {
	model.BaseModel
	UserName string `json:"user_name" gorm:"type:varchar(20);default:not null"` // 用户名
	Password string `json:"password" gorm:"default:not null"`                   // 密码
	Email    string `json:"email" gorm:"default:not null"`                      // 邮箱
	Status   bool   `json:"status" gorm:"default:false"`                        // 是否禁用
}

// 钩子，在注册的时候，会存入加密后的密码
func (u *SsoPaUsers) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = util.HashAndSalt([]byte(u.Password))
	return nil
}

func (SsoPaUsers) TableName() string {
	return "sso_pa_users"
}
