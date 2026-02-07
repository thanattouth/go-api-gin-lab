package main

import (
	"example.com/student-api/config"
	"example.com/student-api/handlers"
	"example.com/student-api/repositories"
	"example.com/student-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. เปิด DB
	db := config.InitDB()

	// 2. ต่อ Layer (Dependency Injection)
	repo := &repositories.StudentRepository{DB: db}
	service := &services.StudentService{Repo: repo}
	handler := &handlers.StudentHandler{Service: service}

	// 3. ตั้งค่า Gin
	r := gin.Default()

	// 4. กำหนด Routes
	r.GET("/students", handler.GetStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.POST("/students", handler.CreateStudent)
	r.PUT("/students/:id", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)

	// 5. เริ่ม Server
	r.Run(":8080")
}
