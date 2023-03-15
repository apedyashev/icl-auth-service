package controllers

import (
	"icl-auth/database"
	"icl-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email, "username": user.Username})
}

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

func GetUserByEmail(context *gin.Context) {
	email := context.Param("email")

	var user models.User
	record := database.Instance.Where("email = ?", email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	user.Password = ""

	context.JSON(http.StatusOK, user)
}
