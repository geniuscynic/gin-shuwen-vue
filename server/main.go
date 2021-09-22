package main

import (
	"github.com/geniuscynic/gin-xjjxmm-admin/repository"
	"github.com/geniuscynic/gin-xjjxmm-admin/router"
)

func main() {

	dbContext := repository.GetDbContext()

	if dbContext != nil {
		repository.InitDbAutoMigrate()

		db, _ := dbContext.DB()
		defer db.Close()
	}

	router.Run()

	//r := gin.Default()
	//r.GET("/ping", controller.Ping)

	//r.POST("api/user/register", controller.Register)

	//r.Run()
}
