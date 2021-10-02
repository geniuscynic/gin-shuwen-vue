package main

import (
	//"github.com/geniuscynic/gin-xjjxmm-admin/docs"
	"github.com/geniuscynic/gin-xjjxmm-admin/repository"
	"github.com/geniuscynic/gin-xjjxmm-admin/router"
	"github.com/geniuscynic/gin-xjjxmm-admin/util/middleware"
	"github.com/gin-gonic/gin"

	_ "github.com/geniuscynic/gin-xjjxmm-admin/docs" // 这里需要引入本地已生成文档
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Docker监控服务
// @version 1.0
// @description docker监控服务后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8889
// @BasePath /api
func main() {
	r := gin.Default()

	r.Use(middleware.Cors())
	//docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	dbContext := repository.GetDbContext()

	if dbContext != nil {
		repository.InitDbAutoMigrate()

		db, _ := dbContext.DB()
		defer db.Close()
	}

	router.Run(r)

	//r := gin.Default()
	//r.GET("/ping", controller.Ping)

	//r.POST("api/user/register", controller.Register)

	//r.Run()
}
