package repository

import "github.com/geniuscynic/gin-xjjxmm-admin/model/entity"

func Login(userName *string) (*entity.User, error) {
	var user entity.User

	eor := GetDbContext().Where("username = ?", userName).First(&user).Error

	if eor != nil {
		return nil, eor
	}

	return &user, nil
}

func RegistUser(user *entity.User) (*entity.User, error) {
	result := GetDbContext().Create(&user)

	return user, result.Error
}
