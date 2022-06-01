package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type CourseInfoMySQL struct {
	ID    string
	Title string
}

func DeleteRecord(db *sql.DB, ID string) {
	query := fmt.Sprintf("DELETE FROM myCourse WHERE ID='%s'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Delete is successful")

}

func EditRecord(db *sql.DB, ID string, Title string) {
	query := fmt.Sprintf("UPDATE myCourse SET Title='%s' WHERE ID='%s'", Title, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertRecord(db *sql.DB, ID string, Title string) {
	query := fmt.Sprintf("INSERT INTO myCourse VALUES('%s', '%s')", ID, Title)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Insert is successful")
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * From my_db.myCourse")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var courseInfoMySQL CourseInfoMySQL
		err = results.Scan(&courseInfoMySQL.ID, &courseInfoMySQL.Title)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(courseInfoMySQL.ID, courseInfoMySQL.Title)
	}
}
