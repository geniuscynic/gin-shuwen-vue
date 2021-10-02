package repository

type Role struct {
	Model `json:",inline"`
	Code  string `gorm:"comment:角色编码;size:90"` // 角色code
	Name  string `gorm:"comment:角色名;size:90"`  // 角色名
}

func (role *Role) Add() (*Role, error) {

	eor := GetDbContext().Create(&role).Error

	if eor != nil {
		return nil, eor
	}

	return role, nil
}

func (role *Role) List() (*[]Role, error) {
	var roles []Role

	result := GetDbContext().Find(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return &roles, nil
}
