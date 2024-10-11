package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save the user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Succesfully created the user"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate the user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Succesfully logged in"})
}
