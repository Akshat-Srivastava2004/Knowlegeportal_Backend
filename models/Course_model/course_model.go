package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseStudent struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Coursename      string             `json:"coursename" bson:"coursename"`
	Description     string             `json:"description" bson:"description"`
	Category        string             `json:"category" bson:"category"`
	Instructor      string             `json:"instructor" bson:"instructor"`
	Courseduration  int64              `json:"courseduration" bson:"courseduration"`
	Studentenrolled int64              `json:"studentenrolled" bson:"studentenrolled"`
	Prerequist      string             `json:"prerequist" bson:"preqequist"`
	Courseprice     int64              `json:"courseprice" bson:"courseprice"`
	Rating          string             `json:"rating" bson:"rating"`
	Language        string             `json:"language" bson:"language"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
}
