package router

import (
	"github.com/geniuscynic/gin-xjjxmm-admin/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("/account")
	{
		group.POST("/login", controller.Login)
		group.POST("/regist", controller.Regist)
		group.POST("/add", controller.AddUser)

		group.POST("/role/add", controller.AddRole)
		group.GET("/role/list", controller.ListRoles)
	}

}
