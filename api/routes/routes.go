package routes

import (
	"heilgar/finance-tracker/config"
	"heilgar/finance-tracker/core"
	"heilgar/finance-tracker/src/account"

	"heilgar/finance-tracker/src/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Ensure the database is connected
	config.ConnectDatabase()

	// Initialize repositories with the connected database
	userRepository := user.NewRepository(config.DB)
	accountRepository := account.NewRepository(config.DB)

	// Initialize controllers with the repositories
	userController := user.NewController(userRepository)
	accountController := account.NewController(accountRepository)

	apiV1 := router.Group("/api/v1")
	{
		core.RegisterRoutesForGroup(apiV1, userController)
		core.RegisterCustomRouteForGroup(apiV1, "POST", userController.GetPath("/register"), userController.RegisterUser)
		core.RegisterRoutesForGroup(apiV1, accountController)
	}

	return router
}
