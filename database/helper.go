package database

import (
	"context"
	"fmt"
	"log"
	"time"

	courseModel "github.com/Akshat-Srivastava2004/educationportal/models/Course_model"
	feedbackModel "github.com/Akshat-Srivastava2004/educationportal/models/Feedback_model"
	studentModel "github.com/Akshat-Srivastava2004/educationportal/models/Student_model"
	studentenrolledModel "github.com/Akshat-Srivastava2004/educationportal/models/Studentenrolled_model"
	teacherModel "github.com/Akshat-Srivastava2004/educationportal/models/Teacher_model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func Insertuser(user teacherModel.TeacherProfile) primitive.ObjectID {
	Collection := GetCollection("TeacherProfile")
	inserteduser, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("the inserted  user is ", inserteduser.InsertedID)
	return inserteduser.InsertedID.(primitive.ObjectID)
}

func Insertusertoken(token teacherModel.Teachertokens) primitive.ObjectID {
	Collection := GetCollection("Teachertokens")
	insertedtoken, err := Collection.InsertOne(context.Background(), token)
	if err != nil {
		panic(err)
	}
	fmt.Println("the inserted token is ", insertedtoken.InsertedID)
	return insertedtoken.InsertedID.(primitive.ObjectID)
}
func Insertstudent(userstudent studentModel.StudentProfile) primitive.ObjectID {
	Collection := GetCollection("StudentProfile")
	insertedstudent, err := Collection.InsertOne(context.Background(), userstudent)
	if err != nil {
		panic(err)
	}
	fmt.Println("the inserted user student is ", insertedstudent.InsertedID)
	return insertedstudent.InsertedID.(primitive.ObjectID)
}
func Insertstudenttoken(userstudenttoken studentModel.Studenttokens) primitive.ObjectID {
	Collection := GetCollection("Studenttokens")
	insertedstudenttoken, err := Collection.InsertOne(context.Background(), userstudenttoken)
	if err != nil {
		panic(err)
	}
	fmt.Println("the inserted user student is ", insertedstudenttoken.InsertedID)
	return insertedstudenttoken.InsertedID.(primitive.ObjectID)
}
func Insertcourse(studentcourse courseModel.CourseStudent) primitive.ObjectID {
	Collection := GetCollection("CourseStudent")
	inseretedcourse, err := Collection.InsertOne(context.Background(), studentcourse)
	if err != nil {
		panic(err)
	}
	fmt.Println("the inserted course student is ", inseretedcourse.InsertedID)
	return inseretedcourse.InsertedID.(primitive.ObjectID)
}
func Enrolledstudent(studentenroll studentenrolledModel.StudentEnrolled) primitive.ObjectID {
	Collection := GetCollection("StudentEnrollment")
	inseretedstudentintocourse, err := Collection.InsertOne(context.Background(), studentenroll)
	if err != nil {
		panic(err)
	}
	fmt.Println("the inserted enrolled student is ", inseretedstudentintocourse.InsertedID)
	return inseretedstudentintocourse.InsertedID.(primitive.ObjectID)
}
func UpdateStudentEnrollment(courseName string) (*mongo.UpdateResult, error) {
	// Get the reference to the collection where courses are stored
	collection := GetCollection("CourseStudent")

	// Log the courseName for debugging
	fmt.Printf("Updating course enrollment for course: %s\n", courseName)

	// Define the filter to find the course by its name
	filter := bson.M{"coursename": courseName}

	// Log the filter for debugging
	fmt.Printf("Filter used in update: %+v\n", filter)

	// Define the update to increment the studentenrolled field
	update := bson.M{
		"$inc": bson.M{
			"studentenrolled": 1, // Increment studentenrolled by 1
		},
	}

	// Log the update operation for debugging
	fmt.Printf("Update operation: %+v\n", update)

	// Perform the update operation
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err // Return the error if update fails
	}

	// Log the update result for debugging
	fmt.Printf("Update result: %+v\n", updateResult)

	// Return the result of the update operation
	return updateResult, nil
}
// func AddSelectedCourse(email string, courseName string) (*mongo.UpdateResult, error) {
// 	// Get the reference to the student profile collection
// 	collection := GetCollection("StudentProfile")

// 	// Define the filter to find the student by their email
// 	filter := bson.M{"email": email}

// 	// Check if the student profile exists
// 	var studentProfile studentModel.StudentProfile
// 	err := collection.FindOne(context.Background(), filter).Decode(&studentProfile)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// If no document is found, return a clear error message
// 			fmt.Printf("No student profile found for email: %s\n", email)
// 			return nil, fmt.Errorf("no student profile found for email: %s", email)
// 		}
// 		// Return the error if there is an issue decoding the document
// 		fmt.Printf("Error finding student profile: %v\n", err)
// 		return nil, err
// 	}

// 	// Check if courseselected is null or not an array and initialize it as an empty array
// 	if studentProfile.Courseselected == "" {
// 		fmt.Printf("Initializing 'courseselected' as an empty array for email: %s\n", email)

// 		// Initialize courseselected as an empty array
// 		update := bson.M{
// 			"$set": bson.M{
// 				"courseselected": []string{},
// 			},
// 		}

// 		_, err := collection.UpdateOne(context.Background(), filter, update)
// 		if err != nil {
// 			fmt.Printf("Error initializing 'courseselected' as array: %v\n", err)
// 			return nil, err
// 		}
// 	}

// 	// Define the update operation to add the course to the Courseselected array
// 	update := bson.M{
// 		"$addToSet": bson.M{
// 			"courseselected": courseName, // Adds the course if it's not already in the array
// 		},
// 	}

// 	// Log the update operation for debugging
// 	fmt.Printf("Update operation: %+v\n", update)

// 	// Perform the update operation
// 	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		// Log the error if update fails
// 		fmt.Printf("Error updating student profile: %v\n", err)
// 		return nil, err // Return the error if update fails
// 	}

// 	// Log the result of the update operation
// 	fmt.Printf("Update result: %+v\n", updateResult)

// 	// Return the result of the update operation
// 	return updateResult, nil
// }

// InsertMCQ inserts a single MCQ document into the specified MongoDB collection.
func InsertMCQ(mcq teacherModel.MCQ) (primitive.ObjectID, error) {
	// Get the MongoDB collection where the MCQs will be stored.
	collection := GetCollection("MCQCollection")

	// Set a timeout context for the database operation.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the MCQ document into the MongoDB collection.
	insertedMCQ, err := collection.InsertOne(ctx, mcq)
	if err != nil {
		log.Printf("Error inserting MCQ: %v", err)
		return primitive.ObjectID{}, err // Return zero ObjectID and the error
	}

	// Log and return the inserted MCQ's ObjectID.
	insertedID := insertedMCQ.InsertedID.(primitive.ObjectID)
	fmt.Println("Inserted MCQ with ID: ", insertedID)

	return insertedID, nil // Return the inserted ID and no error
}

// Feedbackadd inserts feedback into the database and returns the inserted ID
func Feedbackadd(feedback feedbackModel.Feedback) (primitive.ObjectID, error) {
	Collection := GetCollection("Feedback")

	// Insert feedback into the collection
	insertedFeedback, err := Collection.InsertOne(context.Background(), feedback)
	if err != nil {
		// Log the error and return it
		log.Printf("Error inserting feedback: %v", err)
		return primitive.ObjectID{}, err
	}

	// Successfully inserted feedback, return the inserted ID
	fmt.Println("The inserted feedback ID is ", insertedFeedback.InsertedID)
	return insertedFeedback.InsertedID.(primitive.ObjectID), nil
}
