package controllers

import (
	"context"
	"entertainment/auth"
	"entertainment/configs"
	"entertainment/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userPostCollection *mongo.Collection = configs.GetCollection(configs.DB, "Entertainment", "userPosts")
var PostLikesCollection *mongo.Collection = configs.GetCollection(configs.DB, "Entertainment", "postLikes")

//CreateUser is the api used to tget a single user
func CreateUserPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var post models.UserPost

		defer cancel()
		if auth.ValidateUserTokenInHeader(c.Request) == false {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
			return

		}

		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}

		validationErr := validate.Struct(post)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"user_id": post.User_id})
		defer cancel()
		if err != nil {
			// log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while checking for the email"})
			return
		}

		if count > 0 {
			post.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			post.ID = primitive.NewObjectID()
			post.PostId = post.ID.Hex()
			post.IsDeleted = false
			post.LikeCount = 0

			resultInsertionNumber, insertErr := userPostCollection.InsertOne(ctx, post)
			if insertErr != nil {
				msg := fmt.Sprintf("User item was not created")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}
			defer cancel()

			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": resultInsertionNumber})
		} else {
			msg := fmt.Sprintf("Cannot create user post")
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}

	}
}

//CreateUserprofile is the api used to tget a single user
func GetAllUserPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserGetRequest
		var userPosts []models.UserPostResponse

		defer cancel()
		if auth.ValidateUserTokenInHeader(c.Request) == false {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
			return

		}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"user_id": user.User_id, "email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while checking for the email"})
			return
		}

		fmt.Println(count)

		if count > 0 {
			////////////////////////////////////////////////////////////// Sales Staging Pipelines //////////////////////////////////////////////////
			matchStageUserProfiles := bson.D{{"$match", bson.D{{"user_id", user.User_id}, {"email", user.Email}}}}
			inlinepipeRawProduct := []bson.D{{{"$match", bson.D{{"$expr", bson.D{{"$and", bson.A{
				bson.M{"$eq": bson.A{"$$rawProduct_OrgID", "$user_id"}},
			}}}}}}}}

			cloudlookupStageRawProduct := bson.D{{"$lookup", bson.D{{"from", "userPosts"},
				{"let", bson.D{{"rawProduct_OrgID", "$user_id"}}},
				{"pipeline", inlinepipeRawProduct},
				{"as", "posts"}}}}

			result, getErr := userCollection.Aggregate(ctx, mongo.Pipeline{matchStageUserProfiles, cloudlookupStageRawProduct})
			if getErr = result.All(ctx, &userPosts); getErr != nil {
				msg := fmt.Sprintf("User profile was not obtained", getErr.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}

			defer cancel()

			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": userPosts})
		}

	}
}

func ViewFriendsPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.FriendsPostGetRequest
		// var userPosts []models.UserFriendPostResponse
		var userPosts []bson.M
		defer cancel()
		if auth.ValidateUserTokenInHeader(c.Request) == false {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
			return

		}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": validationErr.Error()})
			return
		}

		////////////////////////////////////////////////////////////// Friends Staging Pipelines //////////////////////////////////////////////////
		matchStagefriendList := bson.D{{"$match", bson.D{{"user_id", user.User_id}}}}
		inlinepipeFriendList := []bson.D{{{"$match", bson.D{{"$expr", bson.D{{"$and", bson.A{
			bson.M{"$eq": bson.A{"$$rawProduct_OrgID", "$userid"}},
		}}}}}}}}

		cloudlookupStageFriendList := bson.D{{"$lookup", bson.D{{"from", "userfriends"},
			{"let", bson.D{{"rawProduct_OrgID", "$user_id"}}},
			{"pipeline", inlinepipeFriendList},
			{"as", "friends"}}}}

		unwindFriends := bson.D{{"$unwind", bson.D{{"path", "$friends"}}}}

		projectStageSalesOrder02 := bson.D{{"$project", bson.D{{"friends.friendsdata.first_name", 1}, {"friends.friendsdata.last_name", 1}, {"friends.friendsdata._id", 1}, {"friends.friendsdata.user_id", 1}}}}

		inlinepipeFriendPosts := []bson.D{{{"$match", bson.D{{"$expr", bson.D{{"$and", bson.A{
			bson.M{"$eq": bson.A{"$$rawProduct_OrgID", "$user_id"}},
		}}}}}}}}

		cloudlookupStageFriendPosts := bson.D{{"$lookup", bson.D{{"from", "userPosts"},
			{"let", bson.D{{"rawProduct_OrgID", "$friends.friendsdata.user_id"}}},
			{"pipeline", inlinepipeFriendPosts},
			{"as", "posts"}}}}

		sortPostsbyDate := bson.D{{"$sort", bson.D{{"posts.created_at", 1}}}}
		unwindPosts := bson.D{{"$unwind", bson.D{{"path", "$posts"}}}}

		inlinepipePostLikes := []bson.D{{{"$match", bson.D{{"Isliked", true}, {"$expr", bson.D{{"$and", bson.A{
			bson.M{"$eq": bson.A{"$$rawProduct_OrgID", "$postid"}},
		}}}}}}}}

		cloudlookupStagePostLikes := bson.D{{"$lookup", bson.D{{"from", "postLikes"},
			{"let", bson.D{{"rawProduct_OrgID", "$posts.postid"}}},
			{"pipeline", inlinepipePostLikes},
			{"as", "count"}}}}

		// CountLikes := bson.D{{"$project"}}

		result, getErr := userCollection.Aggregate(ctx, mongo.Pipeline{matchStagefriendList, cloudlookupStageFriendList, unwindFriends, projectStageSalesOrder02, cloudlookupStageFriendPosts, sortPostsbyDate, unwindPosts, cloudlookupStagePostLikes})
		if getErr = result.All(ctx, &userPosts); getErr != nil {
			msg := fmt.Sprintf("User profile was not obtained", getErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}

		////////////////////////////////////////////////////////////// FriendsPosts Staging Pipelines //////////////////////////////////////////////////
		// matchStagefriendPosts := bson.D{{"$match", bson.D{{"user_id", user.User_id}}}}

		// resultPostData, postErr := userPostCollection.Aggregate(ctx, mongo.Pipeline{matchStagefriendPosts})

		// countLikes, err := PostLikesCollection.CountDocuments(ctx, bson.M{"postid": user.PostId, "Isliked": true})

		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		// 	return
		// }

		//// Check whether the user liked the picture

		// Data := PostLikeResponse{
		// 	Count: int(countLikes),
		// }

		c.JSON(http.StatusOK, gin.H{"Status": true, "Result": userPosts})

	}
}

type PostLikeRequest struct {
	ID      primitive.ObjectID `json:"_id"`
	User_id string             `json:"user_id" validate:"required"`
	IsLiked bool               `json:"Isliked"`
	PostId  string             `json:"postId" validate:"required"`
}
type PostLikeResponse struct {
	Count int `json:"count"`
}

func LikePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user PostLikeRequest
		// var userPosts []models.UserFriendPostResponse
		// var userPosts []bson.M
		defer cancel()

		if auth.ValidateUserTokenInHeader(c.Request) == false {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
			return

		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": validationErr.Error()})
			return
		}
		/// Check Post Exist
		count, err := PostLikesCollection.CountDocuments(ctx, bson.M{"postid": user.PostId})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "Cannot find post exists"})
			return
		}

		fmt.Println(count)

		if count > 0 {
			filter := bson.M{"postid": user.PostId, "user_id": user.User_id}
			// updateID := req.ID.Hex()
			update := bson.M{
				"$set": bson.M{
					"user_id": user.User_id,
					"Isliked": user.IsLiked,
				},
			}

			resultInsertionNumber, insertErr := PostLikesCollection.UpdateOne(ctx, filter, update)
			if insertErr != nil {
				fmt.Println("2", insertErr)
				msg := fmt.Sprintf("Cannot insert user like")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}
			fmt.Println(resultInsertionNumber)
			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": resultInsertionNumber.UpsertedID})

		} else {
			fmt.Println("Insert Document")

			user.ID = primitive.NewObjectID()

			resultInsertionNumber, insertErr := PostLikesCollection.InsertOne(ctx, user)
			if insertErr != nil {
				fmt.Println("1", insertErr)
				msg := fmt.Sprintf("Cannot insert user like")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			fmt.Println(resultInsertionNumber)
			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": resultInsertionNumber})

		}

	}

}

// func LikePostCount() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		var user PostLikeRequest
// 		defer cancel()

// 	}
// }
