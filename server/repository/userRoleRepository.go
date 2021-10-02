package repository

type UserRole struct {
	UserId   uint   `gorm:"comment:用户id" json:"user_id"`     // 用户id
	RoleCode string `gorm:"comment:角色code" json:"role_code"` // 角色code
}

func (userRole *UserRole) Add(userRoles *[]UserRole) (*[]UserRole, error) {
	result := GetDbContext().Create(&userRoles)

	return userRoles, result.Error
}
