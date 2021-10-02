package service

import (
	"time"

	"github.com/geniuscynic/gin-xjjxmm-admin/repository"
	"github.com/geniuscynic/gin-xjjxmm-admin/util"
)

type AddRoleModel struct {
	Code string `json:"code"` // 角色code
	Name string `json:"name"` // 角色名
}

func (addRoleModel *AddRoleModel) Covert() *repository.Role {
	user := &repository.Role{
		Code: addRoleModel.Code,
		Name: addRoleModel.Name,
	}

	user.CreatedAt = util.Time(time.Now())
	user.UpdatedAt = util.Time(time.Now())

	return user
}

type RoleService struct{}

func (roleService *RoleService) Add(model *AddRoleModel) (*repository.Role, error) {

	role := model.Covert()

	role, eor := role.Add()

	if eor != nil {
		return nil, eor
	}

	return role, nil
}

func (roleService *RoleService) List() (*[]repository.Role, error) {

	roleResitory := &repository.Role{}

	role, eor := roleResitory.List()

	if eor != nil {
		return nil, eor
	}

	return role, nil
}
