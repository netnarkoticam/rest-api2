package handlesrs

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netnarkoticam/rest-api2.git/internal/entity"
)

type UserRepo interface {
	RegisterUser (ctx context.Context, user entity.User) (int, error)
}

func RegisterUser (repo UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input entity.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := repo.RegisterUser(c.Request.Context(),input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
        c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}