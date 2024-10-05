package studentcontroller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func StudentDashboard(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := Store.Get(r, "Student-session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Check if the session has expired by checking if "email" exists in the session
	email, ok := session.Values["email"].(string)
	if !ok || email == "" {
		http.Error(w, "Session expired. Please log in again.", http.StatusUnauthorized)
		return
	}

	// Retrieve other session values
	username := session.Values["username"].(string)
	// fullname := session.Values["fullname"].(string)

	// Step 7: Pass the username and email to the frontend using a template
	data := struct {
		Email    string
		Username string
	}{
		Email:    email,
		Username: username,
	}

	// Render the "contact.html" template and inject the session data
	templatePath := filepath.Join("template", "contact.html")
	tmpl := template.Must(template.ParseFiles(templatePath))
	tmpl.Execute(w, data)
	// // Display user dashboard or send JSON response with session details
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Welcome, %s! Your email is %s and your full name is %s.", username, email, fullname)
}
