package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentProfile struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProfilePhotoURL string             `json:"profilephoto" bson:"profilephoto"`
	Username        string             `json:"username" bson:"username"`
	Email           string             `json:"email" bson:"email"`
	Fullname        string             `json:"fullname" bson:"fullname"`
	Phonenumber     int64              `json:"phonenumber" bson:"phonenumber"`
	Password        string             `json:"Password" bson:"Password"`
	Gender          string             `json:"gender" bson:"gender"`
	Address         string             `json:"address" bson:"address"`
	// Courseselected  string             `json:"courseselected" bson:"courseselected"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// json format for response and bson format for mongodb storage
// Access token and refresh token generate  here

type Studenttokens struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Studentid     primitive.ObjectID `json:"studentid" bson:"studentid"`
	Access_Token  string             `json:"accesstoken" bson:"accesstoken"`
	Refresh_Token string             `json:"refreshtoken" bson:"refreshtoken"`
	ExpireAt      time.Time          `json:"ExpireAt" bson:"ExpireAt"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAT" BSON:"updatedAT"`
}
