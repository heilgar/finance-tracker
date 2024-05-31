package models

import (
	"gorm.io/gorm"
)

type CryptoTrade struct {
	gorm.Model
	UserID       uint
	AccountID    uint
	AssetID      uint
	TradeType    string
	Quantity     float64
	PricePerUnit float64
	TradeDate    string
}
