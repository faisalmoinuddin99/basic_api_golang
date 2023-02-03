package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for Course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("Welcome to Basic API Building")
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline.in</h1>"))
}

// get all courses data
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	// setting headers
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get  one course based on request course Id in go lang
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	// setting headers
	w.Header().Set("content-Type", "application/json")

	// gran id from request
	params := mux.Vars(r)
	fmt.Println(params) // exercise

	// loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given " + params["id"])

}
