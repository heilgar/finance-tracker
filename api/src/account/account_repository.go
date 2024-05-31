package account

import (
	"gorm.io/gorm"
	"heilgar/finance-tracker/core"
)

type AccountRepository interface {
	core.Repository[Account]
}

type accountRepository struct {
	core.Repository[Account]
}

func NewRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		Repository: core.NewBaseRepository[Account](db),
	}
}
