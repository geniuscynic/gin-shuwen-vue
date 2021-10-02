package service

import (
	"errors"
	"time"

	"github.com/geniuscynic/gin-xjjxmm-admin/repository"
	"github.com/geniuscynic/gin-xjjxmm-admin/util"
)

type LoginModel struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RegisterModel struct {
	Username string `json:"username"` // 用户登录名
	Password string `json:"password"` // 用户登录密码
	NickName string `json:"nickname"` // 用户昵称

}

type AddUserModel struct {
	Username string   `json:"username"` // 用户登录名
	Password string   `json:"password"` // 用户登录密码
	NickName string   `json:"nickname"` // 用户昵称
	Roles    []string `json:"roles"`    //所属角色

}

type EditUserModel struct {
	ID uint `json:"id"` // 用户id
	Username string   `json:"username"` // 用户登录名
	Password string   `json:"password"` // 用户登录密码
	NickName string   `json:"nickname"` // 用户昵称
	Roles    []string `json:"roles"`    //所属角色

}


func (registerModel *RegisterModel) Covert() *repository.User {
	user := &repository.User{
		Username: registerModel.Username,
		Password: registerModel.Password,
		NickName: registerModel.NickName,
	}

	user.CreatedAt = util.Time(time.Now())
	user.UpdatedAt = util.Time(time.Now())

	return user
}

func (model *AddUserModel) Covert() *repository.User {
	user := &repository.User{
		Username: model.Username,
		Password: model.Password,
		NickName: model.NickName,
	}

	user.CreatedAt = util.Time(time.Now())
	user.UpdatedAt = util.Time(time.Now())

	return user
}

func (model *EditUserModel) Covert() *repository.User {
	user := &repository.User{
		Username: model.Username,
		Password: model.Password,
		NickName: model.NickName,
	}

user.ID = model.ID
	user.CreatedAt = util.Time(time.Now())
	user.UpdatedAt = util.Time(time.Now())

	return user
}

type UserService struct{}

func (userService *UserService) Login(model *LoginModel) (*repository.User, error) {

	user := &repository.User{}
	user, eor := user.FindByName(&model.Username)

	if user == nil {
		return nil, eor
	}

	if user.Password != model.Password {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

func (userService *UserService) Regist(model *RegisterModel) (*repository.User, error) {
	user := model.Covert()

	user, eor := user.Add()

	if eor != nil {
		return nil, eor
	}

	return user, nil
}

func (userService *UserService) AddUser(model *AddUserModel) error {

	userDao := model.Covert()
	userDao, eor := userDao.Add()
	if eor != nil {
		return nil
	}

	//user := &repository.User{}
	userRole := &repository.UserRole{}

	var userRoles []*repository.UserRole

	for _, role := range model.Roles {
		temp := &repository.UserRole{
			UserId:   userDao.ID,
			RoleCode: role,
		}

		userRoles = append(userRoles, temp)

		//temp.RoleCode = "aa"
	}

	//println(userRoles)

	_, eor = userRole.Add(&userRoles)

	return nil
}


func (userService *UserService) EditUser(model *EditUserModel) error  {
	userDao := model.Covert()
	userDao, eor := userDao.Update()
	if eor != nil {
		return nil
	}

	//user := &repository.User{}
	userRole := &repository.UserRole{}

	var userRoles []*repository.UserRole

	for _, role := range model.Roles {
		temp := &repository.UserRole{
			UserId:   userDao.ID,
			RoleCode: role,
		}

		userRoles = append(userRoles, temp)

		//temp.RoleCode = "aa"
	}

	//println(userRoles)
	eor = userRole.Delete(&model.ID)
	_, eor = userRole.Add(&userRoles)

	return nil
}


func (userService *UserService) List()  (*[]*repository.User, error){
	userRepository := &repository.User{}

	users, eor := userRepository.List()

	if eor != nil {
		return nil, eor
	}

	return users, nil
}


func (userService *UserService) GetUsers(id *uint) (user *repository.User, err error) {
	userRepository := &repository.User{}

	users, eor := userRepository.FindById(id)

	if eor != nil {
		return nil, eor
	}

	return users, nil
}