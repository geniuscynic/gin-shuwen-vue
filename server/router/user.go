package router

import (
	"github.com/geniuscynic/gin-xjjxmm-admin/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("/api/user")
	{
		group.POST("/login", controller.Login)
		group.POST("/regist", controller.Regist)
	}

}
