package studentenrollcontroller

import (
	"fmt"
	"net/http"

	studentcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Student_controller"
	"github.com/Akshat-Srivastava2004/educationportal/database"
	model "github.com/Akshat-Srivastava2004/educationportal/models/Studentenrolled_model"
)

func Enrolledstudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Initialize and populate the studentenrollment struct with form data
	var studentenrollment model.StudentEnrolled
	studentenrollment.Coursename = r.FormValue("coursename")
	coursename := r.FormValue("coursename") // Course name for enrollment

	// Retrieve the session
	session, err := studentcontroller.Store.Get(r, "Student-session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Check if the session is valid and retrieve email
	email, ok := session.Values["email"].(string)
	if !ok || email == "" {
		http.Error(w, "Session expired. Please log in again.", http.StatusUnauthorized)
		return
	}

	// Retrieve fullname and update studentenrollment struct with session details
	fullname := session.Values["fullname"].(string)
	studentenrollment.Fullname = fullname
	studentenrollment.Email = email

	// Insert enrolled student into the database
	insertedID := database.Enrolledstudent(studentenrollment)
	studentenrollment.ID = insertedID

	// Call the helper function to update course enrollment (increment studentenrolled by 1)
	updateResult, err := database.UpdateStudentEnrollment(coursename)
	if err != nil {
		http.Error(w, "Failed to update course enrollment", http.StatusInternalServerError)
		return
	}

	// Log the results for debugging purposes
	fmt.Println("Inserted enrollment ID:", insertedID)
	fmt.Printf("Course enrollment updated successfully for course '%s'. Update result: %+v\n", coursename, updateResult)

	profileUpdateResult, err := database.AddSelectedCourse(email, coursename)
	if err != nil {
		// Log the error for debugging
		fmt.Printf("Failed to update student profile for email '%s': %v\n", email, err)
		http.Error(w, "Failed to update student profile", http.StatusInternalServerError)
		return
	}

	// Log successful update
	fmt.Printf("Student profile updated successfully for student '%s'. Update result: %+v\n", email, profileUpdateResult)

	// Redirect the user back to the form page after successful enrollment
	http.Redirect(w, r, "/index.html", http.StatusSeeOther)
}
