package main

import (
	"fmt"
	"log"
	"net/http"

	"GoMicroServices/pkg/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.CourseRoutes(r)
	http.Handle("/", r)

	fmt.Println("Listening to port 5001")
	log.Fatal(http.ListenAndServe(":5001", r))

}
