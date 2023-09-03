package models

import (
	"gorm.io/gorm"
)

// 群信息
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId int64
	Type    int
	Icon    string
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
