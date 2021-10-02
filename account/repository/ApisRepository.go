package repository

type Apis struct {
	Model    `json:",inline"`
	Path string `gorm:"comment:api路径" json:"path"`             // 用户登录名
	Method string `gorm:"comment:方法" json:"method"`             // 用户登录密码
	Description string `gorm:"comment:描述" json:"description"` // 用户昵称
}



