package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	AccountName string  `gorm:"type:varchar(100);not null"`
	AccountType string  `gorm:"type:varchar(10);not null;check:account_type in ('Bank', 'Debit', 'Credit', 'Savings', 'Cash', 'Crypto')"`
	Balance     float64 `gorm:"type:decimal(15,2);default:0.00"`
	Currency    string  `gorm:"type:varchar(10);default:'USD'"`
}
