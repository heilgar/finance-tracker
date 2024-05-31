package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserID       uint
	CategoryName string
	CategoryType string
}
