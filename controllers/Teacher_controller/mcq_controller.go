package teachercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/Akshat-Srivastava2004/educationportal/database"
	models "github.com/Akshat-Srivastava2004/educationportal/models/Teacher_model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Akshat-Srivastava2004/educationportal/middleware"
	"github.com/golang-jwt/jwt/v5"
)

// FetchQuestionsHandler - Fetch MCQs for the course stored in session
func FetchQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session
	w.Header().Set("Access-Control-Allow-Origin", "https://blue-meadow-0b28d241e.6.azurestaticapps.net")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true") // If you're sending cookies or auth headers
	claims, ok := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)
   if !ok || claims == nil {
	http.Error(w, "Unauthorized: invalid token context", http.StatusUnauthorized)
	return
}
	// email := claims["email"].(string)
	course := claims["course"].(string)

	// Fetch MCQs from the database based on the course name
	mcqs, err := FetchMCQsByCourse(course)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching MCQs: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println("the mcq questions is ", mcqs)

	// Render the questions as JSON to the frontend
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mcqs)
	if err != nil {
		http.Error(w, "Unable to encode MCQs to JSON", http.StatusInternalServerError)
		return
	}
}

// FetchMCQsByCourse - Helper function to get MCQs by course name
func FetchMCQsByCourse(courseName string) ([]models.MCQ, error) {
	var mcqs []models.MCQ
	collection := database.GetCollection("MCQCollection")

	// Set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query to find MCQs for the given course name
	filter := bson.M{"coursename": courseName}
	cursor, err := collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var mcq models.MCQ
		if err := cursor.Decode(&mcq); err != nil {
			return nil, err
		}
		mcqs = append(mcqs, mcq)
	}
	return mcqs, nil
}

// EvaluateAnswersHandler - Evaluate the user's quiz answers
func EvaluateAnswersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	w.Header().Set("Access-Control-Allow-Origin", "https://blue-meadow-0b28d241e.6.azurestaticapps.net")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true") // If you're sending cookies or auth headers
	claims, ok := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized: invalid token context", http.StatusUnauthorized)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Get the session to retrieve the course

	// email := claims["email"].(string)
	course := claims["course"].(string)
	email :=claims["email"].(string)

	// Fetch the correct MCQs from the database based on the course name
	mcqs, err := FetchMCQsByCourse(course)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching MCQs: %v", err), http.StatusInternalServerError)
		return
	}

	// Compare the user's answers with the correct answers
	var correctCount int
	for i, mcq := range mcqs {
		// Fetch the user's answer from the form data (assuming inputs are named as answer0, answer1, etc.)
		userAnswer := r.FormValue("answer" + strconv.Itoa(i))

		// Check if the user's answer matches the correct answer
		if userAnswer == mcq.CorrectAnswer {
			correctCount++
		}
	}

	// Calculate the score
	totalQuestions := len(mcqs)
	score := (float64(correctCount) / float64(totalQuestions)) * 100

	// // Render the result as JSON or send a response to the client
	// result := map[string]interface{}{
	// 	"total_questions": totalQuestions,
	// 	"correct_answers": correctCount,
	// 	"score":           score,
	// }
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")
	from := os.Getenv("EMAIL_HOST_USER")
	smtpPassword := os.Getenv("EMAIL_HOST_PASSWORD") // Renamed to avoid conflict with the user's password
	to := []string{email}                            // smtp.SendMail expects a slice of strings

	var subject string
	var body string

	// Check score and set email content
	if score > 5 { // Changed to check if score is greater than 20%
		subject = "Subject: Congratulations!\n"
		body = `Congratulations! You have successfully passed the Final evaluation round  and are now part of our teaching community. We're thrilled to have you onboard!

Now you can access our Dashboard Portal using your existing login ID and password. This platform is designed to support you in guiding and inspiring students. Here is the link to get started: http://portal-production-2862.up.railway.app.

As a valued teacher, you'll have the chance to share your expertise, make a meaningful impact, and connect with students eager to learn. Our portal is tailored to empower you in creating engaging and impactful learning experiences.

Thank you for joining us in shaping the future, one student at a time!`
	} else {
		subject = "Subject: Application Status\n"
		body = "We regret to inform you that your application will not proceed further."
		// No need to send email if score is 20% or below
		w.Write([]byte("Your score is below 20%. Email notification not sent."))
		return
	}

	message := []byte(subject + "\n" + body)

	// SMTP authentication
	auth := smtp.PlainAuth("", from, smtpPassword, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
		return
	}
	fmt.Println("Email sent successfully to:", email)
	response := map[string]interface{}{
		"message":         "Email sent successfully",
		"email":           email,
		// "score":           score,
		"total_questions": totalQuestions,
		"correct_answers": correctCount,
		"status":          "pass",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	// Redirect to the result page with the score passed as a query parameter
	// http.Redirect(w, r, "https://blue-meadow-0b28d241e.6.azurestaticapps.net/result.html", http.StatusSeeOther)
}
