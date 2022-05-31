package models

import (
	"GoMicroServices/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Course struct {
	gorm.Model
	Id    string `gorm:""json:"id"`
	Title string `json: "Title"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Course{})
}

func (c *Course) CreateCourse() *Course {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllCourse() []Course {
	var Course []Course
	db.Find(&Course)
	return Course
}

func GetCourseByID(ID string) (*Course, *gorm.DB) {
	var getCourse Course
	db := db.Where("ID=?", ID).Find(&getCourse)
	return &getCourse, db
}

func DeleteCourse(ID string) Course {
	var course Course
	db.Where("ID=?", ID).Delete(course)
	return course
}
