package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Courses = make(map[string]CourseInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", Home)
	router.HandleFunc("/api/v1/courses", AllCourses)
	router.HandleFunc("/api/v1/courses/{courseid}", Course).Methods("GET", "PUT", "POST", "DELETE") //routes and attaching a CRUD request

	fmt.Println("Listening to port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
