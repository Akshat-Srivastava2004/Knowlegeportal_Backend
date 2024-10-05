package contactcontroller

import (
	"html/template"
	"net/http"
	"path/filepath"

	studentcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Student_controller"
)

// ContactPageHandler renders the feedback form with the username and email pre-populated
func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := studentcontroller.Store.Get(r, "Student-session")
	if err != nil {
		http.Error(w, "Session not found: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Get the session data (email and username)
	email, emailOk := session.Values["email"].(string)
	username, usernameOk := session.Values["username"].(string)

	// Check if the session is valid
	if !emailOk || !usernameOk || email == "" || username == "" {
		http.Error(w, "Session expired. Please log in again.", http.StatusUnauthorized)
		return
	}

	// Data to be passed to the template
	data := struct {
		Email    string
		Username string
	}{
		Email:    email,
		Username: username,
	}

	// Parse and execute the feedback/contact form template
	templatePath := filepath.Join("template", "contact.html") // Adjust the path according to your project structure
	tmpl := template.Must(template.ParseFiles(templatePath))
	tmpl.Execute(w, data)
}

// ContactPageHandler renders the feedback form with the username and email pre-populated
func StudentEnrollmentPageHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := studentcontroller.Store.Get(r, "Student-session")
	if err != nil {
		http.Error(w, "Session not found: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Get the session data (email and username)
	email, emailOk := session.Values["email"].(string)
	username, usernameOk := session.Values["username"].(string)

	// Check if the session is valid
	if !emailOk || !usernameOk || email == "" || username == "" {
		http.Error(w, "Session expired. Please log in again.", http.StatusUnauthorized)
		return
	}

	// Data to be passed to the template
	data := struct {
		Email    string
		Username string
	}{
		Email:    email,
		Username: username,
	}

	// Parse and execute the feedback/contact form template
	templatePath := filepath.Join("template", "studentenrollment.html") // Adjust the path according to your project structure
	tmpl := template.Must(template.ParseFiles(templatePath))
	tmpl.Execute(w, data)
}
