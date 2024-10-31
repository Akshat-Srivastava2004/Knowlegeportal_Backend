package teachercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Define the JWT Claims structure
type Claims struct {
	Teacherid    string `json:"teacherid"`
	Gender       string `json:"gender"`
	Phonenumber  int64  `json:"phonenumber"`
	Profilephoto string `json:"profilephoto"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Course       string `json:"course"`
	jwt.RegisteredClaims
}

// Struct for the login request
type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

var store = sessions.NewCookieStore([]byte("abcefghljfjkfkjnjkanjjabjgddghadjh"))

// Checkuser validates the teacher's login credentials and returns JWT tokens
func Checkuserserver(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// Ensure the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, akb := ioutil.ReadAll(r.Body)
	if akb != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Request Body:", string(body))
	// Unmarshal the JSON into the LoginRequest struct
	var loginRequest LoginRequest
	abf := json.Unmarshal(body, &loginRequest)
	if abf != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Now you can access the email and password from the struct
	email := loginRequest.Email
	password := loginRequest.Password

	fmt.Println("User entered email is :", email)
	fmt.Println("User entered password is :", password)

	// Create a context with a timeout for MongoDB operations
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	// Access the TeacherProfile collection in MongoDB
	collection := database.GetCollection("TeacherProfile")

	// Search for the user in the database using email
	var user model.TeacherProfile
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User not found in the database
			http.Error(w, `{"error": "User not found"}`, http.StatusUnauthorized)
			return
		}
		// Internal server error (database issue)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	// Compare the hashed password with the provided password
	password = strings.TrimSpace(password)
	user.Password = strings.TrimSpace(user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	// Password matched
	fmt.Println("Password matched")
	course := user.CourseTeach
	username := user.Username
	profilephoto := user.ProfilePhotoURL
	phonenumber := user.Phonenumber
	gender := user.Gender
	id := user.ID
	// Generate JWT access and refresh tokens
	accessToken, err := generateToken(id, gender, phonenumber, profilephoto, username, course, email, os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		fmt.Println("Error generating access token:", err)
		http.Error(w, `{"error": "Failed to generate access token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := generateToken(id, gender, phonenumber, profilephoto, username, course, email, os.Getenv("REFRESH_TOKEN_SECRET"), os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		fmt.Println("Error generating refresh token:", err)
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	// Send the tokens in the response
	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Helper function to generate JWT tokens
func generateToken(id primitive.ObjectID, gender string, phonenumber int64, profilephoto, username, course, email, secret, expiry string) (string, error) {
	// Convert expiry time from string to duration
	expiryDuration, err := time.ParseDuration(expiry)
	if err != nil {
		return "", err
	}

	// Define the JWT claims
	claims := &Claims{
		Teacherid:    id.Hex(),
		Phonenumber:  phonenumber,
		Gender:       gender,
		Profilephoto: profilephoto,
		Username:     username,
		Email:        email,
		Course:       course,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Checkuser validates the teacher's login credentials and returns JWT tokens
func Checkuser(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Access-Control-Allow-Origin", "https://knowlegeportal-production.up.railway.app/")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Ensure the method is POST

	// Now you can access the email and password from the struct
	email := r.FormValue("Email")
	password := r.FormValue("Password")

	fmt.Println("User entered email is :", email)
	fmt.Println("User entered password is :", password)

	// Create a context with a timeout for MongoDB operations
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	// Access the TeacherProfile collection in MongoDB
	collection := database.GetCollection("TeacherProfile")

	// Search for the user in the database using email
	var user model.TeacherProfile
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User not found in the database
			http.Error(w, `{"error": "User not found"}`, http.StatusUnauthorized)
			return
		}
		// Internal server error (database issue)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	// Compare the hashed password with the provided password
	password = strings.TrimSpace(password)
	user.Password = strings.TrimSpace(user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	// Password matched
	fmt.Println("Password matched")
	session, _ := store.Get(r, "Teacher-session")
	session.Values["teacherid"] = user.ID
	session.Values["email"] = email
	// session.Values["username"] = user.Username
	session.Values["course"] = user.CourseTeach
	fmt.Println("the value of course is ",session.Values["course"])
	session.Options = &sessions.Options{
		MaxAge:   3600, // 1 hour
		HttpOnly: true, // Only accessible via HTTP (not JavaScript)
	}

	// Save the session
	err = session.Save(r, w)
	if err!=nil{
		panic(err)
	}
	http.Redirect(w, r, "/resume.html", http.StatusSeeOther)
}
