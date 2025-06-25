package middleware

import (
	"net/http"
	"strings"

	token2 "library_api/pkg/token"

	"github.com/gin-gonic/gin"
)

func AdminCheck(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required and must be Bearer token"})
		ctx.Abort()
		return
	}

	token := parts[1]

	user, err := token2.ValidateAccessToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if user.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
		ctx.Abort()
		return
	}

	ctx.Set("users", user)

	ctx.Next()
}
