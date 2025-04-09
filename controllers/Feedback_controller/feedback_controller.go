package feedbackcontroller

import (
	"fmt"
	"net/http"

	studentcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Student_controller"
	"github.com/Akshat-Srivastava2004/educationportal/database"
	model "github.com/Akshat-Srivastava2004/educationportal/models/Feedback_model"
)

// Feedback handles feedback submission from users
func Feedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://blue-meadow-0b28d241e.6.azurestaticapps.net/")
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Step 1: Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Step 2: Get session
	session, err := studentcontroller.Store.Get(r, "Student-session")
	if err != nil {
		http.Error(w, "Session not found: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Step 3: Check if the session has expired by checking if "email" exists in the session
	email, ok := session.Values["email"].(string)
	if !ok || email == "" {
		http.Error(w, "Session expired. Please log in again.", http.StatusUnauthorized)
		return
	}

	// Step 4: Retrieve other session values
	username := session.Values["username"].(string)

	// Step 5: Create feedback struct
	var feedback model.Feedback
	feedback.Message = r.FormValue("message")
	feedback.Email = email
	feedback.Username = username

	// Step 6: Insert feedback into the database
	insertedID, err := database.Feedbackadd(feedback)
	if err != nil {
		http.Error(w, "Error inserting feedback: "+err.Error(), http.StatusInternalServerError)
		return
	}
	feedback.ID = insertedID
	fmt.Println("The inserted ID is ", insertedID)

	// Step 7: Redirect to feedback confirmation or a success page
	http.Redirect(w, r, "https://blue-meadow-0b28d241e.6.azurestaticapps.net/student_login.html", http.StatusSeeOther)
}
