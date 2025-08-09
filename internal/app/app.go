//nolint:mnd // Стандартный HTTP код
package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
