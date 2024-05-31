package user

import (
	"gorm.io/gorm"
	"heilgar/finance-tracker/core"
)

type UserRepository interface {
	core.Repository[User]
}

type userRepository struct {
	core.Repository[User]
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Repository: core.NewBaseRepository[User](db),
	}
}
