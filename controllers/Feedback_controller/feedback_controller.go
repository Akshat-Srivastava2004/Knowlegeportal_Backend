package feedbackcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Akshat-Srivastava2004/educationportal/database"
	"github.com/Akshat-Srivastava2004/educationportal/middleware"
	model "github.com/Akshat-Srivastava2004/educationportal/models/Feedback_model"
	"github.com/golang-jwt/jwt/v5"
)

// Feedback handles feedback submission from users
func Feedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://blue-meadow-0b28d241e.6.azurestaticapps.net")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true") // If you're sending cookies or auth headers
	claims, ok := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized: invalid token context", http.StatusUnauthorized)
		return
	}

	// Step 1: Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
    
	// Step 2: Get session
	email :=claims["email"].(string)
	username:=claims["username"].(string)
    

	fmt.Println("the email in  the feedback is ",email)
	fmt.Println("the username in the feedback is ",username)
	// Step 3: Check if the session has expired by checking if "email" exists in the session
    
	// Step 5: Create feedback struct
	var feedback model.Feedback
	 message:=r.FormValue("message")
	fmt.Println("the message in the messageboxis ",message)
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
	// http.Redirect(w, r, "https://blue-meadow-0b28d241e.6.azurestaticapps.net/student_login.html", http.StatusSeeOther)
    response :=map[string]string{
		"email":email,
		"username":username,
		"message":message,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
