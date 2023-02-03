package middleware

import (
	"net/http"
	"strings"
	"todo-api/config"
	"todo-api/server"

	"github.com/gin-gonic/gin"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"singup",
	}
)

func shouldCheckToken(router string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(router, p) {
			return false
		}

	}
	return true
}

func CheckAuthMiddleware(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !shouldCheckToken(ctx.Request.URL.Path) {
			ctx.Next()
			return
		}

		_, err := config.Token(s, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err})
			return
		}

		ctx.Next()
	}
}
