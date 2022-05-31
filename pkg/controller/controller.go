package controller

import (
	"GoMicroServices/pkg/models"
	"GoMicroServices/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var NewCourse models.Course

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To The Available Courses")
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	NewCourse := models.GetAllCourse()
	res, _ := json.Marshal(NewCourse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	courseDetails, _ := models.GetCourseByID(courseId)

	res, _ := json.Marshal(courseDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	CreateCourse := &models.Course{}
	utils.ParseBody(r, CreateCourse)
	c := CreateCourse.CreateCourse()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	course := models.DeleteCourse(courseId)
	res, _ := json.Marshal(course)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var updateCourse = &models.Course{}
	utils.ParseBody(r, updateCourse)
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	courseDetails, db := models.GetCourseByID(courseId)
	if updateCourse.Id != "" {
		courseDetails.Id = updateCourse.Id
	}
	if updateCourse.Title != "" {
		courseDetails.Title = updateCourse.Title
	}
	db.Save(&courseDetails)
	res, _ := json.Marshal(courseDetails)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
