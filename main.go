package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()

	models.GetUserList()

	//
	r := router.Router()
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
