package controller

import (
	"net/http"

	"github.com/geniuscynic/gin-xjjxmm-admin/model/request"
	"github.com/geniuscynic/gin-xjjxmm-admin/model/response"
	"github.com/geniuscynic/gin-xjjxmm-admin/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var login request.LoginModel

	if ctx.ShouldBindJSON(&login) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, response.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	user, error := service.Login(&login)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, response.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, response.CreateSuccessResponse(&user))

}

func Regist(ctx *gin.Context) {
	var user request.RegisterModel

	if ctx.ShouldBindJSON(&user) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, response.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	result, error := service.RegistUser(&user)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, response.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, response.CreateSuccessResponse(&result))

}
