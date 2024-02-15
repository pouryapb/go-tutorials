package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pouryapb/go-tutorials/rest-api/models"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully."})
}
