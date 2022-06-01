package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type CourseInfo struct {
	ID    string `json:"id"`
	Title string `json:Title`
}

//Declaring Map existence globally
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

func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	db, err := OpenDataBase()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	PopulateData(db)
	json.NewEncoder(w).Encode(Courses)
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {

	db, err := OpenDataBase()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Read URL and take the data in(E.g {courseid}, this is the key) and places it into a map.
	//if the curl command is curl -H "Content-Type:application/json" -X POST http://localhost:5001/api/v1/courses/IOS401 -d "{\"id\":\"IOS401\",\"title\":\"Swift Programming\"}"
	//the key:value pair will be IOS401.
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Println(Courses)

	if r.Header.Get("Content-Type") == "application/json" {
		PopulateData(db)
		if r.Method == "POST" {
			// read the string sent to the service
			//Declaring an empty variable named newCourse of type CourseInfo
			var newCourse CourseInfo
			//Reading the all the json data
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				//converts JSON to object pointing to newCourse
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course " + "information " + "in JSON format"))
					return
				}
				// check if course exists; add only if
				// course does not exist
				//As Courses is a map of all the data in the DB,
				//params key:value pair is [courseId:IOS401]
				//Courses Key:value pair is [IOS401:{IOS401 Swift Programming} IOT301:{IOT301 TV}]
				//Soo it compares the Value in params and the key in Courses
				if _, ok := Courses[params["courseid"]]; !ok {
					query := fmt.Sprintf("INSERT INTO courses (ID, Title) VALUES('%s', '%s')", newCourse.ID, newCourse.Title)
					_, err := db.Query(query)
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("Insert is successful")
					Courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
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
			PopulateData(db)
			fmt.Println(Courses)
			fmt.Println(params)
			var newCourse CourseInfo
			requestBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(requestBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course " + "information " + "in JSON format"))
					return
				}

				//If there is a matching Key from Courses and Value from Param the it will update
				if k, ok := Courses[params["courseid"]]; !ok {
					query := fmt.Sprintf("INSERT INTO courses VALUES('%s', '%s')", newCourse.ID, newCourse.Title)
					_, err := db.Query(query)
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("Insert is successful")
					Courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
				} else {
					query := fmt.Sprintf("UPDATE courses SET Title='%s' WHERE ID='%s'", newCourse.Title, k.ID)
					_, err := db.Query(query)
					if err != nil {
						panic(err.Error())
					}
					Courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("201 - Course Updated: " + params["courseid"]))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " + "course information " + "in JSON format"))
			}
		}
	}

	if r.Method == "GET" {
		PopulateData(db)
		if _, ok := Courses[params["courseid"]]; ok {
			json.NewEncoder(w).Encode(Courses[params["courseid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found from GET"))
		}
	}

	if r.Method == "DELETE" {
		PopulateData(db)
		if key, ok := Courses[params["courseid"]]; ok {
			query := fmt.Sprintf("DELETE FROM courses WHERE ID='%s'", key.ID)
			_, err := db.Query(query)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("Delete is successful")
			delete(Courses, params["courseid"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - Course deleted: " + params["courseid"]))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found"))
		}
	}
}
