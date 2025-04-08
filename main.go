package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	courseroute "github.com/Akshat-Srivastava2004/educationportal/routes/Course_route"
	feedbackroute "github.com/Akshat-Srivastava2004/educationportal/routes/Feedback_route"
	studentroute "github.com/Akshat-Srivastava2004/educationportal/routes/Student_route"
	teacherroute "github.com/Akshat-Srivastava2004/educationportal/routes/Teacher_route"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Education portal ")
	r := mux.NewRouter()

	// Register routes
	teacherroute.Router(r)
	courseroute.CourseRouter(r)
	studentroute.StudentRouter(r)
	feedbackroute.FeedbackRouter(r)

	// Serve static files
	fs := http.FileServer(http.Dir("./template"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// Get the port from the environment (Render provides this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
