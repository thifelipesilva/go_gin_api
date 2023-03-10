package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thifelipesilva/go_gin_api/controllers"
	"github.com/thifelipesilva/go_gin_api/database"
	"github.com/thifelipesilva/go_gin_api/models"
)

var ID int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func StudentCreateMock() {
	student := models.Student{
		Name:         "Test Name",
		Registration: "123456",
		Shift:        "Noite",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func StudentDeleteMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestListAllStudentsHandler(t *testing.T) {
	database.ConecctionWithDB()
	StudentCreateMock()
	defer StudentDeleteMock()
	r := SetupRoutesTest()
	r.GET("/students", controllers.ListAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindStudentRegistrationHandler(t *testing.T) {
	database.ConecctionWithDB()
	StudentCreateMock()
	defer StudentDeleteMock()
	r := SetupRoutesTest()
	r.GET("/students/registration/:registration", controllers.FindStudentWithRegistration)
	req, _ := http.NewRequest("GET", "/students/registration/123456", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindStudentWithIDHandler(t *testing.T) {
	database.ConecctionWithDB()
	StudentCreateMock()
	defer StudentDeleteMock()
	r := SetupRoutesTest()
	r.GET("/students/:id", controllers.FindStudentWithID)
	pathReq := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathReq, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)
	assert.Equal(t, "Test Name", studentMock.Name, "Names should to equals")
	assert.Equal(t, "123456", studentMock.Registration, "Registration should to equals")
	assert.Equal(t, "Noite", studentMock.Shift, "Shift should to equals")
}

func TestDeleteStudentHandler(t *testing.T) {
	database.ConecctionWithDB()
	StudentCreateMock()
	r := SetupRoutesTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	pathReq := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathReq, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateStudentHandler(t *testing.T) {
	database.ConecctionWithDB()
	StudentCreateMock()
	defer StudentDeleteMock()
	r := SetupRoutesTest()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	pathReq := "/students/" + strconv.Itoa(ID)
	student := models.Student{
		Name:         "Test2 Name",
		Registration: "654321",
		Shift:        "Tarde",
	}
	studentToJson, _ := json.Marshal(student)
	req, _ := http.NewRequest("PATCH", pathReq, bytes.NewBuffer(studentToJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var mockDataStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockDataStudent)
	assert.Equal(t, "Test2 Name", mockDataStudent.Name)
	assert.Equal(t, "654321", mockDataStudent.Registration)
	assert.Equal(t, "Tarde", mockDataStudent.Shift)
}
