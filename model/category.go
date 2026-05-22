package model

import (
	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"id,primarykey" json:"id" binding:"-"`
	CreatedAt LocalTime      `gorm:"created_at" json:"created_at" binding:"-"`
	UpdatedAt LocalTime      `gorm:"updated_at" json:"updated_at" binding:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" binding:"-"`
	Name      string         `gorm:"name" json:"name" binding:"required"`
	ParentID  *uint          `gorm:"parent_id" json:"parent_id" binding:"required,numeric"`
	Status    uint8          `gorm:"status" json:"status" binding:"numeric"`
	Sort      uint           `gorm:"sort" json:"sort" binding:"numeric"`
	Children  []Category     `gorm:"foreignKey:parent_id" json:"children" binding:"-"`
}
