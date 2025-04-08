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
	fmt.Println("Welcome to Education portal")

	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./template/index.html")
	})

	// Static files
	fs := http.FileServer(http.Dir("./template"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// Routes
	teacherroute.Router(r)
	courseroute.CourseRouter(r)
	studentroute.StudentRouter(r)
	feedbackroute.FeedbackRouter(r)

	// PORT and Bind Fix
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe("0.0.0.0:"+port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
