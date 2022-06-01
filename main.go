package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Database opened on port 3306")

	Courses = make(map[string]CourseInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", Home)
	router.HandleFunc("/api/v1/courses/get", GetAllCourse).Methods("GET")                             //need to display the courses in the DB
	router.HandleFunc("/api/v1/courses/add/{courseid}/{title}", CourseHandler).Methods("POST", "PUT") //routes and attaching a CRUD request
	router.HandleFunc("/api/v1/courses/delete/{courseid}", CourseHandler).Methods("DELETE")           //routes and attaching a CRUD request

	fmt.Println("Listening to port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
