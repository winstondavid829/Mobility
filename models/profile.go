package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserFriends struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserID      string             `json:"user_id" validate:"required"`
	FriendsData `json:"friend_data"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	IsRemoved   bool      `json:"Isremoved"`
}

type FriendsData struct {
	ID          primitive.ObjectID `bson:"_id"`
	First_name  *string            `json:"first_name" validate:"required"`
	Last_name   *string            `json:"last_name" validate:"required"`
	Email       *string            `json:"email" validate:"email,required"`
	Phone       *string            `json:"phone" validate:"required"`
	User_id     string             `json:"user_id"`
	Gender      string             `json:"gender" validate:"required"`
	Bio         string             `json:"bio"`
	AccountType string             `json:"accountType" default:"public"`
	Website     string             `json:"website"`
	Birthday    string             `json:"birthday"`
}
