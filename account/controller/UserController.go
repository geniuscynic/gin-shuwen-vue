package controller

import (
	"net/http"
	"strconv"

	"github.com/geniuscynic/gin-xjjxmm-admin/model"
	"github.com/geniuscynic/gin-xjjxmm-admin/service"
	"github.com/gin-gonic/gin"
)

// Login
// @Tags user
// @Summary 登入
// @accept application/json
// @Produce application/json
// @Param data body service.LoginModel true "登入"
// @Success 200 {object} model.Response "返回空对象"
// @Router /user/login [post]
func Login(ctx *gin.Context) {
	var login service.LoginModel

	if ctx.ShouldBindJSON(&login) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	userService := &service.UserService{}

	user, error := userService.Login(&login)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&user))

}

//Register
// @Tags user
// @Summary 注册
// @accept application/json
// @Produce application/json
// @Param data body service.RegisterModel true "注册"
// @Success 200 {object} model.Response "返回空对象"
// @Router /user/register [post]
func Register(ctx *gin.Context) {
	var user service.RegisterModel

	if ctx.ShouldBindJSON(&user) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	userService := &service.UserService{}
	result, error := userService.Regist(&user)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&result))

}

// AddUser
// @Tags user
// @Summary 添加用户
// @accept application/json
// @Produce application/json
// @Param data body service.AddUserModel true "添加用户"
// @Success 200 {object} model.Response "返回空对象"
// @Router /user/add [post]
func AddUser(ctx *gin.Context) {
	var user service.AddUserModel

	if ctx.ShouldBindJSON(&user) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	userService := &service.UserService{}
	error := userService.AddUser(&user)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(""))

}

// EditUser
// @Tags user
// @Summary 添加用户
// @accept application/json
// @Produce application/json
// @Param data body service.EditUserModel true "修改用户"
// @Success 200 {object} model.Response "返回空对象"
// @Router /user/add [post]
func EditUser(ctx *gin.Context)  {
	var user service.EditUserModel

	if ctx.ShouldBindJSON(&user) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	userService := &service.UserService{}
	error := userService.EditUser(&user)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(""))
}

// ListUsers
// @Tags user
// @Summary 添加用户
// @accept application/json
// @Produce application/json
// @Success 200 {object} model.Response "返回空对象"
// @Router /user [get]
func ListUsers(ctx *gin.Context)  {
	userService := &service.UserService{}
	result, error := userService.List()

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&result))
}

// GetUsers
// @Tags user
// @Summary 添加用户
// @accept application/json
// @Produce application/json
// @Param id path int true "用户id"
// @Success 200 {object} model.Response "返回空对象"
// @Router /user/{id} [get]
func GetUsers(ctx *gin.Context)  {
	id := ctx.Param("id")

	uid, error := strconv.Atoi(id)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(500, "id无效"))
		return
	}

	uuid := uint(uid)

	userService := &service.UserService{}
	result, error := userService.GetUsers(&uuid)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&result))
}