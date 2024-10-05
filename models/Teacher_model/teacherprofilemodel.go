package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeacherProfile struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProfilePhotoURL string             `json:"profilephoto" bson:"profilephoto"`
	Username        string             `json:"username" bson:"username"`
	Gender          string             `json:"gender" bson:"gender"`
	Email           string             `json:"email" bson:"email"`
	Fullname        string             `json:"fullname" bson:"fullname"`
	Phonenumber     int64              `json:"phonenumber" bson:"phonenumber"`
	Password        string             `json:"Password" bson:"Password"`
	CourseTeach     string             `json:"courseenrolled" bson:"courseenrolled"`
	Address         string             `json:"address" bson:"address"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// json format for response and bson format for mongodb storage
// Access token and refresh token generate  here

type Teachertokens struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Teacherid     primitive.ObjectID `json:"studentid" bson:"studentid"`
	Access_Token  string             `json:"accesstoken" bson:"accesstoken"`
	Refresh_Token string             `json:"refreshtoken" bson:"refreshtoken"`
	ExpireAt      time.Time          `json:"ExpireAt" bson:"ExpireAt"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAT" BSON:"updatedAT"`
}
