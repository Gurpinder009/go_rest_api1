package routes

import (
	"go_rest_api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPetRoutes(router *gin.Engine) {
	petRoutes := router.Group("/pets")
	{
		petRoutes.GET("/", controllers.GetPets)
		petRoutes.GET("/:id", controllers.GetPet)
		petRoutes.POST("/", controllers.PostPet)
	}
}
