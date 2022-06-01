package routes

import (
	"GoMicroServices/pkg/controller"

	"github.com/gorilla/mux"
)

var CourseRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/v1/", controller.Home)
	router.HandleFunc("/api/v1/courses", controller.GetCourse)                                          //need to display the courses in the DB
	router.HandleFunc("/api/v1/courses/add/{courseId}", controller.CreateCourse).Methods("PUT", "POST") //routes and attaching a CRUD request
	router.HandleFunc("/api/v1/courses/delete/{courseId}", controller.DeleteCourse).Methods("DELETE")   //routes and attaching a CRUD request
}
