package main

import (
	"go_rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	routes.RegisterPetRoutes(router)
	router.Run("localhost:8000")
}
