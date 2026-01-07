package routes

import (
	"net/http"

	"example.com/event-mgmt/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. Try again later."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully!"})
}
