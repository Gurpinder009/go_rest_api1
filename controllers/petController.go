package controllers

import (
	"go_rest_api/models"

	"github.com/gin-gonic/gin"
)

func GetPets(ctx *gin.Context) {

	users := []models.User{
		{Fname: "gurpinder", Lname: "singh", Email: "email", Phone: "12380928"},
		{Fname: "Randhir", Lname: "kajoor", Email: "email", Phone: "12380928"},
	}

	ctx.JSON(200, gin.H{
		"users": users,
	})
}

func PostPet(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "you are trying to post user",
	})
}

func GetPet(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{
		"message": "you are trying to get user with id " + id,
	})
}
