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

// model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CourseProce int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper -file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
}

//controller - generally goes in a different file

// serve home route or controller
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is Home Route.</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses Route")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")

	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id & return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given Course-ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create One Course")

	w.Header().Set("Content-Type", "application/json")

	//what if :- Body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body can't be Empty")
	}

	//what about :- Body is {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id, string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return 
}
