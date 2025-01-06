package main

import (
	"fmt"
	"go_rest_api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name       string
	Rollno     uint
	Course     string
	Department string
}

func main() {
	router := gin.Default()
	dsn := "gurpinder:password@123@tcp(localhost:3306)/go_rest_api1?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println(err.Error())
	}
	student := Student{}
	db.AutoMigrate(&student)
	db.Create(&Student{Name: "gurpinder", Rollno: 123, Course: "software development", Department: "Computer"})

	routes.RegisterUserRoutes(router)
	routes.RegisterPetRoutes(router)
	router.Run("localhost:8000")
}
