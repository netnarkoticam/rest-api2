package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/netnarkoticam/rest-api2.git/internal/entity"
	service "github.com/netnarkoticam/rest-api2.git/internal/service/pgdb"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, user entity.User) (int, error)
}

func RegisterUser(service *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input entity.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := service.CreateUser(c.Request.Context(), input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetUser(service *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		user, err := service.GetUser(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(service *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var input entity.User
		if binderr := c.ShouldBindJSON(&input); binderr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		input.ID = id

		if binderr := service.UpdateUser(c.Request.Context(), input); binderr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": binderr.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user updated"})
	}
}

func DeleteUser(s *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err = s.DeleteUser(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	}
}
