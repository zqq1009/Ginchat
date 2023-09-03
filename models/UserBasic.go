package models

import (
	"gorm.io/gorm"
	"time"
)

// 用户信息
type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time `gorm:"default:null"`
	HeartbeatTime time.Time `gorm:"default:null"`
	LoginOutTime  time.Time `gorm:"default:null"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
