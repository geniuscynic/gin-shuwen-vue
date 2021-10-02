package repository

type User struct {
	Model    `json:",inline"`
	Username string `gorm:"comment:用户登录名" json:"user_name"`             // 用户登录名
	Password string `gorm:"comment:用户登录密码" json:"password"`             // 用户登录密码
	NickName string `gorm:"default:系统用户;comment:用户昵称" json:"nick_name"` // 用户昵称
}

func (user *User) FindByName(userName *string) (*User, error) {
	//var user entity.User

	eor := GetDbContext().Where("username = ?", userName).First(&user).Error

	if eor != nil {
		return nil, eor
	}

	return user, nil
}

func (user *User) Add() (*User, error) {
	result := GetDbContext().Create(&user)

	return user, result.Error
}

func (user *User) Update() (*User, error) {
	result := GetDbContext().Save(&user)

	return user, result.Error
}

func (user *User) List() (*[]*User, error) {
	var users []*User

	result := GetDbContext().Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (user *User) FindById(id *uint) (*User, error) {
	result := GetDbContext().First(user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
