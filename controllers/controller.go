package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thifelipesilva/go_gin_api/database"
	"github.com/thifelipesilva/go_gin_api/models"
)

func ListAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func FindStudentWithID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(200, student)
}

func FindStudentWithRegistration(c *gin.Context) {
	var student models.Student
	registration := c.Param("registration")
	database.DB.Where(&models.Student{Registration: registration}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)

}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"messsage": "deleted sucess",
	})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}
