package server

import (
	"ginchat/models"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"text/template"
)

// GetIndex
// @Summary 首页
// @Tags 登录模块
// @Success 200 {string} welcome
// @Router / [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/login.html", "views/chat/head.html")
	if err != nil {
		panic(err)
		return
	}

	ind.Execute(c.Writer, "index")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "welcome",
	//})
}

// ToRegister
// @Summary 注册界面
// @Tags 登录模块
// @Success 200 {string} welcome
// @Router /toRegister [get]
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
		return
	}

	ind.Execute(c.Writer, "register")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "welcome",
	//})
}

// ToChat
// @Summary 个人界面
// @Tags 登录模块
// @Success 200 {string} welcome
// @Router /toChat [get]
func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
		"views/chat/foot.html",
		"views/chat/head.html",
		"views/chat/tabmenu.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html")
	if err != nil {
		panic(err)
		return
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("tochat>>>>>>>>>", user)
	ind.Execute(c.Writer, user)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "welcome",
	//})
}

func Chat(c *gin.Context) {
	utils.Chat(c.Writer, c.Request)
}
