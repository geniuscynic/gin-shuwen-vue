package controller

import (
	"net/http"

	"github.com/geniuscynic/gin-xjjxmm-admin/model"
	"github.com/geniuscynic/gin-xjjxmm-admin/service"
	"github.com/gin-gonic/gin"
)


// AddRole
// @Tags role
// @Summary 添加角色
// @accept application/json
// @Produce application/json
// @Param data body service.AddRoleModel true "roleModel"
// @Success 200 {object} model.Response "返回空对象"
// @Router /role/add [post]
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

// ListRoles
// @Tags role
// @Summary 角色列表
// @accept application/json
// @Produce application/json
// @Success 200 {object} model.Response "返回空对象"
// @Router /role/list [get]
func ListRoles(ctx *gin.Context) {

	roleService := &service.RoleService{}
	result, error := roleService.List()

	if error != nil {

		ctx.JSON(http.StatusInternalServerError, model.CreateFailedResponse(200, error.Error()))

		return
	}

	ctx.JSON(http.StatusOK, model.CreateSuccessResponse(&result))

}
