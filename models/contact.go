package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint   //谁的关系
	TargetId uint   //对应的谁
	Type     int    //对应的类型  0 1  2
	Desc     string //描述信息

}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(">>>>>>>>>>>>", v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}

// 添加好友功能
func AddFriend(userId uint, targetId uint) (int, string) {
	user := UserBasic{}
	//添加自己
	if targetId == userId {
		return -1, "不能添加自己为好友"
	}
	contact0 := Contact{}
	utils.DB.Where("owner_id = ? and target_id = ? and type = 1", userId, targetId).Find(&contact0)
	if contact0.ID != 0 {
		return -1, "不能重复添加好友"
	}
	if targetId != 0 {
		user = FindUserByID(targetId)
		if user.Salt != "" {
			//添加好友是相互的,开启事务
			tx := utils.DB.Begin()
			//事务一旦开始，不论期间什么异常，最终都会Rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			contact := Contact{}
			contact.OwnerId = userId
			contact.TargetId = targetId
			contact.Type = 1
			if err := utils.DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加失败"
			}

			contact1 := Contact{}
			contact1.OwnerId = targetId
			contact1.TargetId = userId
			contact1.Type = 1
			if err := utils.DB.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "添加失败"
			}

			tx.Commit()
			return 0, "添加成功"
		}
		return -1, "没有找到此好友"
	}
	return -1, "输入好友ID为空"
}
