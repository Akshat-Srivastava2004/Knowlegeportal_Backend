package main

import (
	"fmt"
	"log"
	"net/http"

	courseroute "github.com/Akshat-Srivastava2004/educationportal/routes/Course_route"
	feedbackroute "github.com/Akshat-Srivastava2004/educationportal/routes/Feedback_route"
	studentroute "github.com/Akshat-Srivastava2004/educationportal/routes/Student_route"
	teacherroute "github.com/Akshat-Srivastava2004/educationportal/routes/Teacher_route"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Education portal ")
	r := mux.NewRouter()
	// Define the session store globally

	// Register teacher routes
	teacherroute.Router(r)

	// Register student routes
	courseroute.CourseRouter(r)
	studentroute.StudentRouter(r)
	feedbackroute.FeedbackRouter(r)
	fmt.Println("Server getting started ...")

	fs := http.FileServer(http.Dir("./template")) // Make sure this folder has index.html and other static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listening at port 8000 ...")
}
