package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

func CreateCommunity(community Community) (int, string) {
	if len(community.Name) == 0 {
		return -1, "群名称不能为空"
	}

	if community.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := utils.DB.Create(&community).Error; err != nil {
		fmt.Println(err)
		return -1, "建群失败"
	}

	return 0, "创建成功"
}

func LoadCommunity(ownerId uint) ([]*Community, string) {
	data := make([]*Community, 10)
	utils.DB.Where("owner_id = ?", ownerId).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}

	//utils.DB.Where()
	return data, "查询成功"
}

//
//// 添加群功能
//func JoinGroup(name string) (int, string) {
//	group := GroupBasic{}
//	//添加自己
//	if targetId == userId {
//		return -1, "不能添加自己为好友"
//	}
//	contact0 := Contact{}
//	utils.DB.Where("owner_id = ? and target_id = ? and type = 2", name).Find(&group)
//	if contact0.ID != 0 {
//		return -1, "不能重复添加好友"
//	}
//	if targetId != 0 {
//		user = FindUserByID(targetId)
//		if user.Salt != "" {
//			//添加好友是相互的,开启事务
//			tx := utils.DB.Begin()
//			//事务一旦开始，不论期间什么异常，最终都会Rollback
//			defer func() {
//				if r := recover(); r != nil {
//					tx.Rollback()
//				}
//			}()
//
//			contact := Contact{}
//
//			contact.Type = 1
//			if err := utils.DB.Create(&contact).Error; err != nil {
//				tx.Rollback()
//				return -1, "添加失败"
//			}
//
//			contact1 := Contact{}
//			contact1.OwnerId = targetId
//			contact1.TargetId = userId
//			contact1.Type = 1
//			if err := utils.DB.Create(&contact1).Error; err != nil {
//				tx.Rollback()
//				return -1, "添加失败"
//			}
//
//			tx.Commit()
//			return 0, "添加成功"
//		}
//		return -1, "没有找到此好友"
//	}
//	return -1, "输入好友ID为空"
//}
