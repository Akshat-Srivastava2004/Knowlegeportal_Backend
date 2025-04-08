package teachercontroller

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/Akshat-Srivastava2004/educationportal/database"
	models "github.com/Akshat-Srivastava2004/educationportal/models/Teacher_model"
)

// File Upload Handler
func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://blue-meadow-0b28d241e.6.azurestaticapps.net/")
	// Parse form data
	r.ParseMultipartForm(10 << 20) // Limit file size to 10MB

	// Retrieve the file from the form
	file, _, err := r.FormFile("mcqFile")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving the file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get course name from the form
	coursename := r.FormValue("coursename")
	if coursename == "" {
		http.Error(w, "Course name is required", http.StatusBadRequest)
		return
	}

	// Create a temporary CSV file to store the uploaded file content
	tempFile, err := os.Create("uploaded_mcqs.csv")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating temp file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Write the uploaded file content to the temp file
	_, err = tempFile.ReadFrom(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing to temp file: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse CSV file and extract MCQs with the course name
	mcqs, err := parseCSV("uploaded_mcqs.csv", coursename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing CSV: %v", err), http.StatusInternalServerError)
		return
	}

	// Insert questions into MongoDB using the helper function
	err = InsertQuestions(mcqs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting MCQs into DB: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "MCQs for course %s successfully uploaded and stored in DB", coursename)
}

// Parse the CSV file and extract questions and options, assigning the course name to each MCQ.
func parseCSV(filePath, coursename string) ([]models.MCQ, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var mcqs []models.MCQ
	for _, line := range lines {
		if len(line) < 6 {
			continue // Ensure there are enough columns in the row
		}

		// Extract options from the CSV row
		options := []string{line[1], line[2], line[3], line[4]}

		// Create a new MCQ
		mcq := models.MCQ{
			Question:      line[0],
			Options:       options,
			CorrectAnswer: line[5],
			Coursename:    coursename, // Add the course name
		}

		// Append the MCQ to the slice
		mcqs = append(mcqs, mcq)
	}

	return mcqs, nil
}

// / InsertQuestions inserts multiple MCQs into the MongoDB collection
func InsertQuestions(mcqs []models.MCQ) error {
	for _, mcq := range mcqs {
		_, err := database.InsertMCQ(mcq)
		if err != nil {
			return err
		}
	}
	return nil
}
