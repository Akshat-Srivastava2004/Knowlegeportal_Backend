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

	// Add this health check route (Render scans for these)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})

	// Register your routes
	teacherroute.Router(r)
	courseroute.CourseRouter(r)
	studentroute.StudentRouter(r)
	feedbackroute.FeedbackRouter(r)

	// Serve static files
	fs := http.FileServer(http.Dir("./template"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// ⬇️ This is CRUCIAL: use the port Render gives you
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // fallback for local dev
	}

	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
