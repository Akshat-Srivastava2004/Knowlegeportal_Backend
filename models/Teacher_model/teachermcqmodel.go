package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MCQ struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Question      string             `bson:"question"`
	Options       []string           `bson:"options"`
	CorrectAnswer string             `bson:"correct_answer"`
	Coursename    string             `bson:"coursename"`
}
