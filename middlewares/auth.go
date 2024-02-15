package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pouryapb/go-tutorials/rest-api/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized."})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
