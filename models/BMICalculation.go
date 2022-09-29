package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BMI struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Height       string             `json:"height" bson:"height" validate:"required"`
	Weight       string             `json:"Weight" bson:"Weight" validate:"required"`
	BMI          float64            `json:"BMI" bson:"BMI"`
	UserID       string             `json:"user_id" bson:"user_id" validate:"required"`
	CreatedDate  time.Time          `json:"CreatedDate" bson:"CreatedDate"`
	ModifiedDate time.Time          `json:"ModifiedDate" bson:"ModifiedDate"`
}
