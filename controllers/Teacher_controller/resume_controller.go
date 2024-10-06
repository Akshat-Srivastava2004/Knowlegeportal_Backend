package teachercontroller

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/google/generative-ai-go/genai"
// 	"google.golang.org/api/option"
// )

// func uploadToGemini(ctx context.Context, client *genai.Client, path, mimeType string) string {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	options := genai.UploadFileOptions{
// 		DisplayName: path,
// 		MIMEType:    mimeType,
// 	}
// 	fileData, err := client.UploadFile(ctx, "", file, &options)
// 	if err != nil {
// 		log.Fatalf("Error uploading file: %v", err)
// 	}

// 	log.Printf("Uploaded file %s as: %s", fileData.DisplayName, fileData.URI)
// 	return fileData.URI
// }

// func Resume() {
// 	ctx := context.Background()

// 	apiKey, ok := os.LookupEnv("AIzaSyAqV8jIWIUj-b-Ug9yTPAMu3SrQSxWSNDM")
// 	if !ok {
// 		log.Fatalln("Environment variable GEMINI_API_KEY not set")
// 	}

// 	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
// 	if err != nil {
// 		log.Fatalf("Error creating client: %v", err)
// 	}
// 	defer client.Close()

// 	model := client.GenerativeModel("gemini-1.5-flash")

// 	model.SetTemperature(1)
// 	model.SetTopK(64)
// 	model.SetTopP(0.95)
// 	model.SetMaxOutputTokens(8192)
// 	model.ResponseMIMEType = "text/plain"

// 	// model.SafetySettings = Adjust safety settings
// 	// See https://ai.google.dev/gemini-api/docs/safety-settings

// 	// TODO Make these files available on the local file system
// 	// You may need to update the file paths
// 	fileURIs := []string{
// 		uploadToGemini(ctx, client, "Unknown File", "application/octet-stream"),
// 		uploadToGemini(ctx, client, "latestresumehain.pdf", "application/pdf"),
// 		uploadToGemini(ctx, client, "latestresumehain.pdf", "application/pdf"),
// 		uploadToGemini(ctx, client, "latestresumehain.pdf", "application/pdf"),
// 		uploadToGemini(ctx, client, "latestresumehain.pdf", "application/pdf"),
// 		uploadToGemini(ctx, client, "latestresumehain.pdf", "application/pdf"),
// 	}

// 	session := model.StartChat()
// 	session.History = []*genai.Content{
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.Text("LISTEN VERY CAREFULLY I WILL GIVE YOU A RESUME OF A TEACHER BECAUSE TEACHER IS APPLYING FOR A TEACHING JOB OF SOME COURSE FOR EXAMPLE IF USER CHOOSE THE COURSE\nAND SUBMIT THEIR RESUME SO YOU HAVE TO GIVE ME A MARKS OUT OF 100 IS THAT TEACHER IS SUITABLE AUR MUCH TALENTED TO TEACH THAT PARTICULAR COURSE TO STUDENT  .JUST GIVE ME A MARKS OUT OF 100  NOTHING MORE  I WANT ONLY MARKS NO EXPLANATION ."),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.FileData{URI: fileURIs[0]},
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.Text("LISTEN VERY CAREFULLY I WILL GIVE YOU A RESUME OF A TEACHER BECAUSE TEACHER IS APPLYING FOR A TEACHING JOB OF SOME COURSE FOR EXAMPLE IF USER CHOOSE THE COURSE\nAND SUBMIT THEIR RESUME SO YOU HAVE TO GIVE ME A MARKS OUT OF 100 IS THAT TEACHER IS SUITABLE AUR MUCH TALENTED TO TEACH THAT PARTICULAR COURSE TO STUDENT .JUST GIVE ME A MARKS OUT OF 100 NOTHING MORE I WANT ONLY MARKS NO"),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("Okay, I understand. Please provide the teacher's resume and the course they are applying for. I will then give you a score out of 100 based on their suitability for that specific course. \n"),
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.FileData{URI: fileURIs[1]},
// 				genai.Text("FOR WEB DEVLOPMENT \n"),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("85 \n"),
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.FileData{URI: fileURIs[2]},
// 				genai.Text("FOR DSA TEACHING\n"),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("65 \n"),
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.FileData{URI: fileURIs[3]},
// 				genai.Text("FOR APP DEVEOPMENT "),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("70 \n"),
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.Text("Analyze the submitted teacher's resume and evaluate their qualifications to teach a specific subject. Consider the following factors:\n\nDoes the teacher have prior teaching experience in the subject?\nHas the teacher completed relevant projects in the subject area, and how complex were they?\nDoes the teacher show a deep understanding of the subject?\nHas the teacher published any research or papers in the field?\nAre there any relevant certifications or credentials for the subject?\nWhat is the impact of the teacher's work on student outcomes (if available)?\nBased on this analysis, provide a score out of 100 to determine the teacher’s suitability for teaching the subject. Return only the score"),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("Please provide the subject the teacher is applying to teach. I need to know the specific subject to analyze the resume and provide an accurate score. \n"),
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.FileData{URI: fileURIs[4]},
// 				genai.Text("for web development "),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("75 \n"),
// 			},
// 		},
// 		{
// 			Role: "user",
// 			Parts: []genai.Part{
// 				genai.FileData{URI: fileURIs[5]},
// 				genai.Text("for dsa"),
// 			},
// 		},
// 		{
// 			Role: "model",
// 			Parts: []genai.Part{
// 				genai.Text("50 \n"),
// 			},
// 		},
// 	}

// 	resp, err := session.SendMessage(ctx, genai.Text("INSERT_INPUT_HERE"))
// 	if err != nil {
// 		log.Fatalf("Error sending message: %v", err)
// 	}

// 	for _, part := range resp.Candidates[0].Content.Parts {
// 		fmt.Printf("%v\n", part)
// 	}
// }
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func TeacherDashboard(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := store.Get(r, "Teacher-session")
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
	fullname := session.Values["fullname"].(string)
	course := session.Values["course"].(string)

	// Display user dashboard or send JSON response with session details
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome, %s! Your email is %s and your full name is %s and your course is %s", username, email, fullname, course)
}

// Function to handle the resume upload and scoring
func UploadResumeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://knowlegeportal-production.up.railway.app/")
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Max 10MB file size
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	// Get the file and other form data
	file, handler, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "Error Retrieving the File", http.StatusBadRequest)
		return
	}
	defer file.Close()

	name := r.FormValue("name")
	email := r.FormValue("email")

	// Log basic info
	log.Printf("Uploaded File: %v, Name: %s, Email: %s\n", handler.Filename, name, email)

	// Open Gemini Client
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		http.Error(w, "GEMINI_API_KEY not set", http.StatusInternalServerError)
		return
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		http.Error(w, "Error creating Gemini client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Upload the file to Gemini
	options := genai.UploadFileOptions{
		DisplayName: handler.Filename,
		MIMEType:    "application/pdf",
	}

	uploadedFile, err := client.UploadFile(ctx, "", file, &options)
	if err != nil {
		http.Error(w, "Error uploading file to Gemini", http.StatusInternalServerError)
		return
	}

	// Start a chat session and send resume for evaluation
	model := client.GenerativeModel("gemini-1.5-flash")
	session := model.StartChat()
	sess, err := store.Get(r, "Teacher-session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}
	email = sess.Values["email"].(string)
	course := sess.Values["course"].(string)

	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.Text(fmt.Sprintf("Please evaluate this resume for teaching %s and provide a score out of 100 only .", course)),
				genai.FileData{URI: uploadedFile.URI},
			},
		},
	}

	resp, err := session.SendMessage(ctx, genai.Text(fmt.Sprintf("Please evaluate the resume for the course %s and provide score out of 100 dont give me a reason just give me a marks ", course)))

	if err != nil {
		http.Error(w, "Error processing resume", http.StatusInternalServerError)
		return
	}

	// Extract and return the score from the response
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		score := resp.Candidates[0].Content.Parts[0]
		scoreStr := fmt.Sprintf("%v", score)    // Convert the score to a string
		scoreStr = strings.TrimSpace(scoreStr)  // Remove any surrounding whitespace or newlines
		scoreInt, err := strconv.Atoi(scoreStr) // Convert the cleaned string to an integer
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Resume Evaluation Score (Integer): %d", scoreInt)

		// Email setup variables
		smtpHost := os.Getenv("EMAIL_HOST")
		smtpPort := os.Getenv("EMAIL_PORT")
		from := os.Getenv("EMAIL_HOST_USER")
		smtpPassword := os.Getenv("EMAIL_HOST_PASSWORD") // Renamed to avoid conflict with the user's password
		to := []string{email}                            // smtp.SendMail expects a slice of strings

		var subject string
		var body string

		// Check score and set email content
		if scoreInt > 20 {
			subject = "Subject: Congratulations!\n"
			body = "Congratulations! You have cleared the first round of the evaluation."
		} else {
			subject = "Subject: Application Status\n"
			body = "We regret to inform you that your application will not proceed further."
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
	} else {
		http.Error(w, "No score received", http.StatusInternalServerError)
	}
	fmt.Println("Email sent successfully to:", email)
	http.Redirect(w, r, "/Teachertest.html", http.StatusSeeOther)
}
func init() {
	// Load .env file at the start of the application
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
