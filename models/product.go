package models

import (
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type Product struct {
	ProductId    string `gorm:"primaryKey" json:"productId"`
	ProductName  string `gorm:"type:varchar(225)" json:"productName" binding:"required"`
	ProductImage string `gorm:"type:varchar(225)" json:"productImage" binding:"required"`
	CreatedBy    string
	UpdatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type ReqProduct struct {
	ProductId    string                `form:"productId"`
	ProductName  string                `form:"productName" binding:"required"`
	ProductImage *multipart.FileHeader `form:"productImage" binding:"required"`
}
