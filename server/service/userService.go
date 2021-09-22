package service

import (
	"errors"

	"github.com/geniuscynic/gin-xjjxmm-admin/model/entity"
	"github.com/geniuscynic/gin-xjjxmm-admin/model/request"
	"github.com/geniuscynic/gin-xjjxmm-admin/repository"
)

func Login(model *request.LoginModel) (*entity.User, error) {

	user, eor := repository.Login(&model.Username)

	if user == nil {
		return nil, eor
	}

	if user.Password != model.Password {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

func RegistUser(model *request.RegisterModel) (*entity.User, error) {
	user := model.Covert()

	user, eor := repository.RegistUser(user)

	if eor != nil {
		return nil, eor
	}

	return user, nil
}
