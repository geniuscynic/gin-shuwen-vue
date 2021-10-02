package router

import (
	"github.com/geniuscynic/gin-xjjxmm-admin/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	//router := gin.Default()

	group := router.Group("/api")
	{
		InitUserRoute(group)
		InitRoleRoute(group)
	}
	//InitUserRoute(router.Group(""))

	s := &http.Server{
		Addr:           ":8889",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func InitRoleRoute(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("/role")
	{

		group.POST("/add", controller.AddRole)
		group.GET("/list", controller.ListRoles)
	}

}

func InitUserRoute(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("/user")
	{
		group.POST("/login", controller.Login)
		group.POST("/register", controller.Register)
		group.POST("/add", controller.AddUser)
		group.POST("/edit", controller.EditUser)
		group.GET("", controller.ListUsers)
		group.GET("/:id", controller.GetUsers)
	}

}