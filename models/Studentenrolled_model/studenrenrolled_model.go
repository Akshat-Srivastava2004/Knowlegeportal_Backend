package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentEnrolled struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Fullname   string             `json:"name" bson:"name"`
	Email      string             `json:"email" bson:"email"`
	Coursename string             `json:"coursename" bson:"coursename"`
	Payment    bool               `json:"payment" bson:"payment"`
	Studentid  string             `json:"studentid" bson:"studentid"`
	CreatedAt  time.Time          `json:"createdat" bson:"createdat"`
	UpdatedAt  time.Time          `json:"updatedtat" bson:"updatedat"`
}
