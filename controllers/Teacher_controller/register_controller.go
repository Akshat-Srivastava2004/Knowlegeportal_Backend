package teachercontroller

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
	model "github.com/Akshat-Srivastava2004/educationportal/models/Teacher_model"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://knowlegeportal-production.up.railway.app/")
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
	var user model.TeacherProfile

	user.Fullname = r.FormValue("Fullname") // Replace with actual field names
	user.Username = r.FormValue("Username")
	user.Email = r.FormValue("Email")
	user.Password = r.FormValue("Password") // Make sure this field matches your form name
	user.Gender = r.FormValue("Gender")
	user.Address = r.FormValue("Address")
	user.CourseTeach = r.FormValue("CourseTeach")

	// Step 3: Convert the phone number to an integer
	Phonenumber, err := strconv.ParseInt(Phonestr, 10, 64) // Use int64 to handle larger numbers
	if err != nil {
		// Handle error, e.g., invalid number
		http.Error(w, "Invalid phone number format", http.StatusBadRequest)
		return
	}

	// Step 4: Store the phone number as an integer in the user struct
	user.Phonenumber = Phonenumber // If Phonenumber is int64, cast it to int

	// Step 3: Hash the user's password
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword // Store the hashed password

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
	user.ProfilePhotoURL = uploadResult.SecureURL

	fmt.Println("The photo URL is ", uploadResult.SecureURL)
	// Step 9: Insert user into the database
	insertedID := database.Insertuser(user)
	user.ID = insertedID
	fmt.Println("The inserted ID is ", insertedID)

	// Step 10: Respond with the created user data, including the Cloudinary URL
	// Step 10: Redirect to teacher_login.html after successful user creation
	http.Redirect(w, r, "/teacher_login.html", http.StatusSeeOther)

}

var store = sessions.NewCookieStore([]byte("abcefghljfjkfkjnjkanjjabjgddghadjh"))

func Checkuser(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Access-Control-Allow-Origin", "https://knowlegeportal-production.up.railway.app/")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Get email and password from the form
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	fmt.Println("the user enter password is ", email)
	fmt.Println("the user enter password is ", password)
	// Create a context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	collection := database.GetCollection("TeacherProfile")

	// Find the user in the database
	var user model.TeacherProfile
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

	password = strings.TrimSpace(password)
	user.Password = strings.TrimSpace(user.Password)
	hashedPwdBytes := user.Password
	fmt.Println("the string hash password is ", hashedPwdBytes)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwdBytes), []byte(password))
	if err != nil {

		fmt.Println("Password comparison failed:", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else {
		fmt.Println("password matched")
	}
	session, _ := store.Get(r, "Teacher-session")
	session.Values["email"] = user.Email
	session.Values["username"] = user.Username
	session.Values["fullname"] = user.Fullname
	session.Values["course"] = user.CourseTeach
	session.Options = &sessions.Options{
		MaxAge:   3600, // 10 seconds
		HttpOnly: true, // Only accessible via HTTP (not JavaScript)
	}

	session.Save(r, w) // Save the session
	http.Redirect(w, r, "/resume.html", http.StatusSeeOther)
}
