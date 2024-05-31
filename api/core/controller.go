package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Controller interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Collection(c *gin.Context)

	GetBasePath() string
	GetPath(path string) string
}

type BaseController[T any] struct {
	repository Repository[T]
	basePath   string
}

func NewBaseController[T any](repository Repository[T], basePath string) *BaseController[T] {
	return &BaseController[T]{repository: repository, basePath: basePath}
}

func (ctrl *BaseController[T]) Repository() Repository[T] {
	return ctrl.repository
}

func (ctrl *BaseController[T]) GetBasePath() string {
	return ctrl.basePath
}

func (ctrl *BaseController[T]) GetPath(path string) string {
	return ctrl.basePath + path
}

func (ctrl *BaseController[T]) Get(c *gin.Context) {
	// Get entity ID from request
	id := c.Param("id")

	// Convert ID to uint
	entityID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Retrieve entity from repository
	entity, err := ctrl.repository.FindByID(uint(entityID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Return entity as JSON response
	c.JSON(http.StatusOK, entity)
}

func (ctrl *BaseController[T]) Collection(c *gin.Context) {
	// Retrieve all entities from repository
	entities, err := ctrl.repository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve entities"})
		return
	}

	// Return entities as JSON response
	c.JSON(http.StatusOK, entities)
}

func (ctrl *BaseController[T]) Create(c *gin.Context) {
	// Parse request body into entity struct
	var entity T
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Create entity in repository
	if err := ctrl.repository.Create(&entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entity"})
		return
	}

	// Return created entity as JSON response
	c.JSON(http.StatusCreated, entity)
}

func (ctrl *BaseController[T]) Update(c *gin.Context) {
	// Get entity ID from request
	id := c.Param("id")

	// Convert ID to uint
	entityID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Retrieve entity from repository
	entity, err := ctrl.repository.FindByID(uint(entityID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Parse request body into updated entity struct
	var updatedEntity T
	if err := c.BindJSON(&updatedEntity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update entity in repository
	entity = &updatedEntity
	if err := ctrl.repository.Update(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update entity"})
		return
	}

	// Return updated entity as JSON response
	c.JSON(http.StatusOK, entity)
}

func (ctrl *BaseController[T]) Delete(c *gin.Context) {
	// Get entity ID from request
	id := c.Param("id")

	// Convert ID to uint
	entityID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Delete entity in repository
	if err := ctrl.repository.Delete(uint(entityID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entity"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Entity deleted successfully"})
}
