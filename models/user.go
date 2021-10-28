package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId       string `gorm:"primaryKey" json:"userId"`
	UserName     string `gorm:"type:varchar(225)" json:"userName"`
	UserPassword string `gorm:"type:varchar(225)" json:"userPassword"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
