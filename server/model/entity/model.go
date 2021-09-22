package entity

import (
	"database/sql"

	"github.com/geniuscynic/gin-xjjxmm-admin/util"
)

type Model struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt util.Time    `json:"created_at"`
	UpdatedAt util.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
