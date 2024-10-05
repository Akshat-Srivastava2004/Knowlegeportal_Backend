package teachercontroller

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// Function to load environment variables
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return err
	}
	return nil
}

func TeacherMCq(w http.ResponseWriter, r *http.Request) {
	// Load environment variables
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading environment variables")
		return
	}

	// Initialize context
	ctx := context.Background()

	// Retrieve API key from environment variables
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Invalid API key")
		return
	}

	// Create a new GenAI client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	// Set up the generative model
	model := client.GenerativeModel("gemini-1.5-flash")

	// Customize model parameters
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	// Start a session for generating MCQs
	session := model.StartChat()
	sess, err := store.Get(r, "Teacher-session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}
	course := sess.Values["course"].(string)

	// Request MCQs from the generative model
	resp, err := session.SendMessage(ctx, genai.Text(fmt.Sprintf("Please give me 20 MCQ HARD QUESTION ON THE BASIS OF THIS course %s and GIVE ME THE QUESTIONS ONLY WITH 4 OPTIONS IN WHICH ONE IS TRUE and ASSUME THAT THE OTHER PERSON IS APPLYING FOR TEACHING IN BTECH COLLEGES SO ASK A QUESTION ACCORDING TO BTECH LEVEL AND ALSO GIVE ME A ANSWER OF EACH QUESTION ", course)))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// Prepare to parse the response into individual MCQs
	var mcqs []string

	// Iterate through each part of the response content
	for _, part := range resp.Candidates[0].Content.Parts {
		// Format the part (question) and replace ** or other markup with line breaks
		formattedMCQ := strings.ReplaceAll(fmt.Sprintf("%v", part), "**", "\n")
		mcqs = append(mcqs, formattedMCQ)
	}

	// Remove any extra whitespace or empty strings from MCQs
	for i, mcq := range mcqs {
		mcqs[i] = strings.TrimSpace(mcq)
	}

	templatePath := filepath.Join("template", "MCQ.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error loading template:", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template with the MCQs
	err = tmpl.Execute(w, mcqs)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
