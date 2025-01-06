package routes

import (
	"fmt"
	"go_rest_api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", middleWareFunc, middleWareFunc2, controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.POST("/", controllers.PostUsers)
	}
}

func middleWareFunc(ctx *gin.Context) {
	fmt.Println("adsfasdf")
	// ctx.Abort()
	ctx.Next()
}

func middleWareFunc2(ctx *gin.Context) {
	fmt.Println("mmidleware2 ")
}
