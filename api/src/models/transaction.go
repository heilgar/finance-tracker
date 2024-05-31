package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint
	AccountID       uint
	CategoryID      uint
	TransactionType string
	Amount          float64
	TransactionDate string
	Description     string
}
