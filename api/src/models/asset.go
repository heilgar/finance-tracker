package models

import (
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	UserID        uint
	AccountID     uint
	AssetName     string
	Quantity      float64
	PurchasePrice float64
	CurrentPrice  float64
	PurchaseDate  string
}
