package request

import (
	"time"

	"github.com/geniuscynic/gin-xjjxmm-admin/model/entity"
	"github.com/geniuscynic/gin-xjjxmm-admin/util"
)

// User login structure
type LoginModel struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RegisterModel struct {
	Username string `json:"username"` // 用户登录名
	Password string `json:"password"` // 用户登录密码
	NickName string `json:"nickname"` // 用户昵称
}

func (registerModel *RegisterModel) Covert() *entity.User {
	user := &entity.User{
		Username: registerModel.Username,
		Password: registerModel.Password,
		NickName: registerModel.NickName,
	}

	user.CreatedAt = util.Time(time.Now())
	user.UpdatedAt = util.Time(time.Now())

	return user
}
