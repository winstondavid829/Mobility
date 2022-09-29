package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the model that governs all notes objects retrived or inserted into the DB
type UserPost struct {
	ID          primitive.ObjectID `bson:"_id"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
	User_id     string             `json:"user_id" validate:"required"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	PostId      string             `json:"postId"`
	Post        string             `json:"post" validate:"required"`
	LikeCount   int                `json:"likecount"`
	IsDeleted   bool               `json:"Isdeleted"`
	Likes       []PostLike         `json:"likes"`
}

type PostLike struct {
	User_id string `json:"user_id" validate:"required"`
	IsLiked bool   `json:"Isliked" validate:"required"`
}

type UserPostResponse struct {
	ID      primitive.ObjectID `json:"_id"`
	User_id string             `json:"user_id"`
	Posts   []UserPost         `json:"posts"`
}

type FriendsPostGetRequest struct {
	User_id string `json:"user_id" validate:"required"`
}

type UserFriendPostResponse struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id" `
	Friend Friends            `json:"friends"`
	Posts  []UserPost         `json:"posts"`
}
type Friends struct {
	FriendsData PostFriendsInfo `json:"friendsdata"`
}
type PostFriendsInfo struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	First_name string             `json:"first_name" bson:"first_name" validate:"required"`
	Last_name  string             `json:"last_name" bson:"last_name" validate:"required"`
	User_id    string             `json:"user_id" bson:"user_id" validate:"required"`
}
