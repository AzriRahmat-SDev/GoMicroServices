package main

import (
	"database/sql"
)

func OpenDataBase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "azri:password@tcp(127.0.0.1:3306)/my_db?charset=utf8&parseTime&loc=Local")
	return db, err
}

func PopulateData(db *sql.DB) {
	results, err := db.Query("Select * From courses")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var course CourseInfo
		err = results.Scan(&course.ID, &course.Title) //at the end of this line, course will print all the data in the row
		if err != nil {
			panic(err.Error())
		}
		//Store information globally in main.go to be used later on
		Courses[course.ID] = course
	}

}
