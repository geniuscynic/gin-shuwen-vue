package controller

import (
	"net/http"

	"github.com/geniuscynic/gin-xjjxmm-admin/model"
	"github.com/geniuscynic/gin-xjjxmm-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags account
// @Summary 登入
// @accept application/json
// @Produce application/json
// @Param data body service.LoginModel true "登入"
// @Success 200 {object} model.Response "返回空对象"
// @Router /account/login [post]
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

// @Tags account
// @Summary 注册
// @accept application/json
// @Produce application/json
// @Param data body service.RegisterModel true "注册"
// @Success 200 {object} model.Response "返回空对象"
// @Router /account/regist [post]
func Regist(ctx *gin.Context) {
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

// @Tags account
// @Summary 注册
// @accept application/json
// @Produce application/json
// @Param data body service.RegisterModel true "注册"
// @Success 200 {object} model.Response "返回空对象"
// @Router /account/add [post]
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

// @Tags account
// @Summary 添加角色
// @accept application/json
// @Produce application/json
// @Param data body service.AddRoleModel true "roleModel"
// @Success 200 {object} model.Response "返回空对象"
// @Router /account/role/add [post]
func AddRole(ctx *gin.Context) {
	var role service.AddRoleModel

	if ctx.ShouldBindJSON(&role) != nil {
		//log.Println("name", login.Username)
		//log.Println("password", login.Password)

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(500, "模型无法解析"))

		return
	}

	roleService := &service.RoleService{}
	result, error := roleService.Add(&role)

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&result))

}

// @Tags account
// @Summary 角色列表
// @accept application/json
// @Produce application/json
// @Success 200 {object} model.Response "返回空对象"
// @Router /account/role/list [get]
func ListRoles(ctx *gin.Context) {

	roleService := &service.RoleService{}
	result, error := roleService.List()

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&result))

}
