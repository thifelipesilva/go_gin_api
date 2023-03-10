package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thifelipesilva/go_gin_api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ListAllStudents)
	r.GET("/students/:id", controllers.FindStudentWithID)
	r.GET("/students/registration/:registration", controllers.FindStudentWithRegistration)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.Run(":5000")
}
