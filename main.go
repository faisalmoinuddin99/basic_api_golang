package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("Welcome to Basic API Building")
}

// services - file

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

	// grab id from request
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

// Add a course service in go lang
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	// setting headers
	w.Header().Set("content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please fill  course name" )
	return
	}

	// generate unique id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))

	// append course into fake db courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}