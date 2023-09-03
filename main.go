package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()

	//
	r := router.Router()
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
