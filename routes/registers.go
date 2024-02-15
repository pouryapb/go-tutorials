package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pouryapb/go-tutorials/rest-api/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered for event."})
}

func cancelRegisteration(ctx *gin.Context) {}
