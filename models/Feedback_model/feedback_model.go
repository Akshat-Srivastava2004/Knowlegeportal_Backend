package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feedback struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Message  string             `json:"message" bson:"message"`
}
