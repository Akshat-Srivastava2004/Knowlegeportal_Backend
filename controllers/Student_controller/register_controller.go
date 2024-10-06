package studentcontroller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Akshat-Srivastava2004/educationportal/cloudinary"
	"github.com/Akshat-Srivastava2004/educationportal/database"
	"github.com/Akshat-Srivastava2004/educationportal/helper"
	model "github.com/Akshat-Srivastava2004/educationportal/models/Student_model"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles user creation
func CreateUserstudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Step 1: Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	Phonestr := r.FormValue("Phonenumber")
	// Step 2: Retrieve user profile data from the form field "user_data"
	var userstudent model.StudentProfile

	userstudent.Fullname = r.FormValue("Fullname")
	fmt.Print("the user fullname is ", userstudent.Fullname) // Replace with actual field names
	userstudent.Username = r.FormValue("Username")
	userstudent.Email = r.FormValue("Email")
	userstudent.Password = r.FormValue("Password") // Make sure this field matches your form name
	userstudent.Gender = r.FormValue("Gender")
	userstudent.Address = r.FormValue("Address")

	// Step 3: Convert the phone number to an integer
	Phonenumber, err := strconv.ParseInt(Phonestr, 10, 64) // Use int64 to handle larger numbers
	if err != nil {
		// Handle error, e.g., invalid number
		http.Error(w, "Invalid phone number format", http.StatusBadRequest)
		return
	}

	// Step 4: Store the phone number as an integer in the user struct
	userstudent.Phonenumber = Phonenumber // If Phonenumber is int64, cast it to int

	// Step 3: Hash the user's password
	hashedPassword, err := helper.HashPassword(userstudent.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	userstudent.Password = hashedPassword // Store the hashed password

	// Step 4: Retrieve the profile photo file from the form data
	file, handler, err := r.FormFile("ProfilePhotoURL")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	fmt.Println("The profile photo is ", file)
	defer file.Close()

	// Step 5: Save the file locally (you might want to handle this better in production)
	localFilePath := "./uploads/" + handler.Filename
	os.MkdirAll("./uploads", os.ModePerm) // Ensure the directory exists

	localFile, err := os.Create(localFilePath)
	if err != nil {
		http.Error(w, "Error creating local file", http.StatusInternalServerError)
		return
	}
	fmt.Println("The local file path is ", localFile)
	defer localFile.Close()

	_, err = io.Copy(localFile, file)
	if err != nil {
		http.Error(w, "Error saving file locally", http.StatusInternalServerError)
		return
	}

	// Step 6: Upload the file to Cloudinary
	cld := cloudinary.InitCloudinary()
	if cld == nil {
		http.Error(w, "Cloudinary initialization error", http.StatusInternalServerError)
		return
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), localFilePath, uploader.UploadParams{Folder: "profile_photos"})
	if err != nil {
		http.Error(w, "Error uploading file to Cloudinary", http.StatusInternalServerError)
		return
	}
	fmt.Println("The upload result is ", uploadResult)

	// Step 7: Delete the local file after successful upload
	err = os.Remove(localFilePath)
	if err != nil {
		fmt.Println("Warning: Failed to delete local file:", err)
	} else {
		fmt.Println("Local file deleted successfully")
	}

	// Step 8: Update user profile with Cloudinary URL
	userstudent.ProfilePhotoURL = uploadResult.SecureURL

	fmt.Println("The photo URL is ", uploadResult.SecureURL)
	// Step 9: Insert user into the database
	insertedID := database.Insertstudent(userstudent)
	userstudent.ID = insertedID
	fmt.Println("The inserted ID is ", insertedID)

	// Step 10: Respond with the created user data, including the Cloudinary URL
	// Step 10: Redirect to teacher_login.html after successful user creation
	http.Redirect(w, r, "https://storied-ganache-31318f.netlify.app/template/student_login.html", http.StatusSeeOther)

}

var Store = sessions.NewCookieStore([]byte("abcefghljfjkfkjnjkanjjadddwdbjgddghadjh"))

func Checkuserstudent(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Get email and password from the form
	email := r.FormValue("Email")
	password := r.FormValue("Password")

	fmt.Println("The user entered email is:", email)
	fmt.Println("The user entered password is:", password)

	// Create a context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	collection := database.GetCollection("StudentProfile")

	// Find the user in the database
	var user model.StudentProfile
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Email not found in the database
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		} else {
			// Other errors (e.g., database connection issues)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("Stored hashed password:", user.Password)
	fmt.Println("User provided password:", password)

	// Trim leading/trailing spaces
	password = strings.TrimSpace(password)
	user.Password = strings.TrimSpace(user.Password)

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else {
		fmt.Println("Password matched")
	}

	// smtpHost := os.Getenv("EMAIL_HOST")
	// smtpPort := os.Getenv("EMAIL_PORT")
	// from := os.Getenv("EMAIL_HOST_USER")
	// smtpPassword := os.Getenv("EMAIL_HOST_PASSWORD") // Renamed to avoid conflict with the user's password

	// to := []string{email} // smtp.SendMail expects a slice of strings

	// subject := "Subject: Login Notification\n"
	// body := "You have successfully logged in!"
	// message := []byte(subject + "\n" + body)

	// auth := smtp.PlainAuth("", from, smtpPassword, smtpHost)

	// // Send the email
	// err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// if err != nil {
	// 	fmt.Println("Failed to send email:", err)
	// 	http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println("Email sent successfully to:", email)

	// Create a new session and store the user's details in the session
	session, err := Store.Get(r, "Student-session")
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}
	// session.Values["studendid"] = user.ID
	session.Values["email"] = user.Email
	session.Values["username"] = user.Username
	session.Values["fullname"] = user.Fullname
	session.Values["profilephoto"] = user.ProfilePhotoURL
	session.Options = &sessions.Options{
		MaxAge:   3600, // 1 hour
		HttpOnly: true, // Only accessible via HTTP (not JavaScript)
	}

	err = session.Save(r, w) // Save the session
	if err != nil {
		fmt.Println("Error saving session:", err) // More context
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
	// Redirect to the index page after login
	http.Redirect(w, r, "/index.html", http.StatusSeeOther)
}
