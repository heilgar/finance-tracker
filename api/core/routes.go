package core

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutesForGroup(router *gin.RouterGroup, controller Controller) {
	// Create route handlers for CRUD operations
	router.POST(controller.GetBasePath(), controller.Create)
	router.GET(controller.GetBasePath()+"/:id", controller.Get)
	router.PUT(controller.GetBasePath()+"/:id", controller.Update)
	router.DELETE(controller.GetBasePath()+"/:id", controller.Delete)

	router.GET(controller.GetBasePath(), controller.Collection)
}

func RegisterCustomRouteForGroup(
	router *gin.RouterGroup,
	httpMethod string,
	path string,
	controllerMethod func(c *gin.Context)) {
	router.Handle(httpMethod, path, controllerMethod)
}
