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
    "github.com/rs/cors"
)

func main() {
    fmt.Println("Welcome to Education portal")

    r := mux.NewRouter()

    // Register routes
    teacherroute.Router(r)
    courseroute.CourseRouter(r)
    studentroute.StudentRouter(r)
    feedbackroute.FeedbackRouter(r)

    // Determine port for HTTP service
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000" // Default port if not specified
    }

    // Set up CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Update this to your frontend's URL in production
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    })

    handler := c.Handler(r)

    fmt.Println("Server is starting on port:", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}
