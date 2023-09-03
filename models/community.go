package models

import (
	"gorm.io/gorm"
)

// 群
type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}
