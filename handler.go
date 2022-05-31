package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type CourseInfo struct {
	Title string `json:Title`
}

var Courses map[string]CourseInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To The Available Courses")
}

func allCourseHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	if err != nil {
		panic(err.Error())
	}

	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v)
	}

	json.NewEncoder(w).Encode(Courses)
	GetRecords(db)
}

func courseHandler(w http.ResponseWriter, r *http.Request) {

	// if !validKey(r) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("401 - Invalid key"))
	// 	return
	// }

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	params := mux.Vars(r)
	if r.Header.Get("Content-Type") == "application/json" {
		if r.Method == "POST" {
			// read the string sent to the service
			var newCourse CourseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course " + "information " + "in JSON format"))
					return
				}
				// check if course exists; add only if
				// course does not exist
				if _, ok := Courses[params["courseid"]]; !ok {
					Courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
					InsertRecord(db, params["courseid"], params["title"])
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Duplicate course ID"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information " + "in JSON format"))
			}
		}

		if r.Method == "PUT" {
			var newCourse CourseInfo
			requestBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(requestBody, &newCourse)

				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course " + "information " + "in JSON format"))
					return
				}

				if _, ok := Courses[params["courseid"]]; !ok {
					Courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
					InsertRecord(db, params["courseid"], params["title"])
				} else {
					Courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("201 - Course Updated: " + params["courseid"]))
					EditRecord(db, params["courseid"], params["title"])
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " + "course information " + "in JSON format"))
			}
		}
	}

	if r.Method == "GET" {
		if _, ok := Courses[params["courseid"]]; ok {
			json.NewEncoder(w).Encode(
				Courses[params["courseid"]])
			GetRecords(db)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found from GET"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := Courses[params["courseid"]]; ok {
			delete(Courses, params["courseid"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - Course deleted: " + params["courseid"]))
			DeleteRecord(db, params["courseid"])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found"))
		}
	}
}
