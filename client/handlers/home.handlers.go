package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hi!, I'm Robot jajajaj"})
	}
}
