package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"heilgar/finance-tracker/core"
	"net/http"
)

type UserController struct {
	*core.BaseController[User]
}

func NewController(repository core.Repository[User]) *UserController {
	return &UserController{
		BaseController: core.NewBaseController(repository, "users"),
	}
}

func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var input struct {
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := User{
		PasswordHash: string(passwordHash),
		Email:        input.Email,
	}

	if err := ctrl.Repository().Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
