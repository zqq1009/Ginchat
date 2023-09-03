package utils

import (
	"fmt"
	"ginchat/models"
)

func SearchFriend(userId uint) []models.UserBasic {
	contacts := make([]models.Contact, 0)
	objIds := make([]uint64, 0)
	DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(">>>>>>>>>>>>", v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]models.UserBasic, 0)
	DB.Where("id in ?", objIds).Find(&users)
	return users
}

// AddFriend
// @Summary 添加好友
// @Tags 个人中心
// @param userId formData string false "userId"
// @param targetId formData string false "targetId"
// @Success 200 {string} json{"code","message"}
// @Router /contact/addfriend [post]
func AddFriend(userId uint, targetId uint) (int, string) {
	user := models.UserBasic{}
	if targetId != 0 {
		//查找ID
		user = FindUserByID(targetId)
		fmt.Println(targetId, " ", userId)
		//
		if user.Salt != "" {
			if userId == user.ID {
				return -1, "不能加自己"
			}
			contact0 := models.Contact{}
			DB.Where("owner_id =? and target_id =? and type=1", userId, targetId).Find(&contact0)
			if contact0.ID != 0 {
				return -1, "不能重复添加"
			}
			tx := DB.Begin() //事务一旦开始，不论什么异常最终都会 Rollback

			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			contact := models.Contact{}
			contact.OwnerId = userId
			contact.TargetId = targetId
			contact.Type = 1
			if err := DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}

			contact1 := models.Contact{}
			contact1.OwnerId = targetId
			contact1.TargetId = userId
			contact1.Type = 1
			if err := DB.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}

			tx.Commit()
			return 0, "添加好友成功"
		}
		return -1, "没有找到此用户"
	}
	return -1, "好友ID不能为空"
}

func SearchUserByGroupId(communityId uint) []uint {
	contacts := make([]models.Contact, 0)
	objIds := make([]uint, 0)
	DB.Where("target_id = ? and type=2", communityId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint(v.OwnerId))
	}
	return objIds
}
