package router

import (
	"ginchat/docs"
	"ginchat/server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//静态资源
	r.Static("/asset", "./asset")
	r.LoadHTMLGlob("views/**/*")

	//首页
	r.GET("/", server.GetIndex)
	r.GET("/toRegister", server.ToRegister)
	r.GET("/toChat", server.ToChat)
	r.GET("/chat", server.Chat)
	r.POST("/searchFriends", server.SearchFriends)

	//用户模块
	r.POST("/user/getUserList", server.GetUserList)
	r.POST("/user/creatUser", server.CreateUser)
	r.POST("/user/deleteUser", server.DeleteUser)
	r.POST("/user/updateUser", server.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", server.FindUserByNameAndPwd)

	r.POST("user/find", server.FindByID)

	//发送消息
	r.GET("/user/sendMsg", server.SendMsg)
	r.GET("/user/sendUserMsg", server.SendUserMsg)

	//上传模块，文件
	r.POST("/attach/upload", server.Upload)

	//添加
	r.POST("/contact/addfriend", server.AddFriend)
	//r.POST("contact/joinGroup", server.JoinGroup)

	//创建群
	r.POST("/contact/createCommunity", server.CreateCommunity)

	//群列表
	r.POST("contact/loadCommunity", server.LoadCommunity)
	return r
}
