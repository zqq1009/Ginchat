package utils

import (
	"fmt"
	utils "ginchat/asset"
	"ginchat/models"
	"gorm.io/gorm"
	"time"
)

func GetUserList() []*models.UserBasic {
	data := make([]*models.UserBasic, 10)
	DB.Find(&data)
	//for _, v := range data {
	//	//fmt.Println(v)
	//}

	return data
}

// 通过用户名查询用户
func FindUserByName(name string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("name = ?", name).First(&user)
	return user
}

// 通过电话号码查询用户
func FindUserByPhone(phone string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("phone = ?", phone).First(&user)
	return user
}

// 通过邮箱查询用户
func FindUserByEmail(email string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("email = ?", email).First(&user)
	return user
}

// 通过ID查询用户
func FindUserByID(id uint) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("id = ?", id).First(&user)
	return user
}

// 登录用户校验
func FindUserByNameAndPwd(name string, password string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("name = ? and pass_word = ?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	DB.Model(&user).Where("id = ?", user.ID).Update("Identity", temp)
	return user
}

func CreatUser(user models.UserBasic) *gorm.DB {
	return DB.Create(&user)
}

func DeleteUser(user models.UserBasic) *gorm.DB {
	return DB.Delete(&user)
}

func UpdateUser(user models.UserBasic) *gorm.DB {
	return DB.Model(&user).Updates(models.UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})
}
