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

	// ✅ Serve the homepage (index.html)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./template/index.html")
	})

	// ✅ Serve all HTML files like /about.html, /payment.html etc.
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./template"))))

	// ✅ Serve CSS, JS, IMG, etc.
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./template/css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./template/js"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./template/img"))))
	r.PathPrefix("/lib/").Handler(http.StripPrefix("/lib/", http.FileServer(http.Dir("./template/lib"))))
	r.PathPrefix("/scss/").Handler(http.StripPrefix("/scss/", http.FileServer(http.Dir("./template/scss"))))

	// ✅ Backend routes
	teacherroute.Router(r)
	courseroute.CourseRouter(r)
	studentroute.StudentRouter(r)
	feedbackroute.FeedbackRouter(r)

	// ✅ Set Render port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
