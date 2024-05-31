package account

import (
	"heilgar/finance-tracker/core"
)

type AccountController struct {
	*core.BaseController[Account]
}

func NewController(accountRepository AccountRepository) *AccountController {
	return &AccountController{
		BaseController: core.NewBaseController[Account](accountRepository, "accounts"),
	}
}
