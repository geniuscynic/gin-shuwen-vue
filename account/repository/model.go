package repository

import (
	"database/sql"

	"github.com/geniuscynic/gin-xjjxmm-admin/util"
)

type Model struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt util.Time    `gorm:"type:time" json:"created_at"`
	UpdatedAt util.Time    `gorm:"type:time" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
