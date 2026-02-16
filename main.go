package main

import (
	"go-api-gin/config"
	"go-api-gin/handlers"
	"go-api-gin/models"
	"go-api-gin/repositories"
	"go-api-gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Student{})

	repo := repositories.NewStudentRepository(config.DB)
	service := services.NewStudentService(repo)
	handler := handlers.NewStudentHandler(service)

	r.GET("/students", handler.GetStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.POST("/students", handler.CreateStudent)
	r.PUT("/students/:id", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)

	r.Run(":8080")
}