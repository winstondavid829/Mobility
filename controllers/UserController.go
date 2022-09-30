package controllers

import (
	"context"
	"entertainment/auth"
	helper "entertainment/helpers"
	"entertainment/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/crypto/bcrypt"

	"entertainment/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "Entertainment", "user")
var userfriendsCollection *mongo.Collection = configs.GetCollection(configs.DB, "Entertainment", "userfriends")
var validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}

	return string(bytes)
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}

//CreateUser is the api used to tget a single user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

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

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while looking for email"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while looking for email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "User already exists"})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{"Status": true, "Result": resultInsertionNumber})

	}
}

//Login is the api used to tget a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		if auth.ValidateUserTokenInHeader(c.Request) == false {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
			return

		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "incorrect username or password"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, gin.H{"Status": true, "Result": foundUser})

	}
}

type UserSendEmail struct {
	First_name *string `json:"first_name"`
	Last_name  *string `json:"last_name"`
	Email      *string `json:"email" validate:"email,required"`
	Code       string  `json:"code" validate:"required"`
}

func SendUserOTP(email *UserSendEmail) {
	from := mail.NewEmail("Gymboo Customer Support", "judeharshan44@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail(*email.First_name, *email.Email)
	plainTextContent := "Dear " + *email.First_name + " \n" + "To verify your account please use the given code in your application\n" + "Code:" + email.Code
	message := mail.NewSingleEmailPlainText(from, subject, to, plainTextContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

//CreateUserprofile is the api used to tget a single user
func AddUserDetails() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserUpdateRequest

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

		count, err := userCollection.CountDocuments(ctx, bson.M{"user_id": user.User_id})
		defer cancel()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "Error occured while looking for user"})
			return
		}

		if count > 0 {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
			// return
			filter := bson.M{"user_id": user.User_id}

			// updateID := req.ID.Hex()
			update := bson.M{
				"$set": bson.M{
					"first_name":  user.First_name,
					"last_name":   user.Last_name,
					"email":       user.Email,
					"phone":       user.Phone,
					"gender":      user.Gender,
					"bio":         user.Bio,
					"accountType": user.AccountType,
					"website":     user.Website,
					"birthday":    user.Birthday,
				},
			}
			user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			resultInsertionNumber, insertErr := userCollection.UpdateOne(ctx, filter, update)
			if insertErr != nil {
				msg := fmt.Sprintf("User item was not created")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}

			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": resultInsertionNumber})

		}

	}
}

func AddFriendsToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserFriends

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

		countUser, err := userCollection.CountDocuments(ctx, bson.M{"user_id": user.FriendsData.User_id, "email": user.FriendsData.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "Error occured while looking for user exists"})
			return
		}

		count, err := userfriendsCollection.CountDocuments(ctx, bson.M{"user_id": user.User_id, "friend_data.$.user_id": user.FriendsData.User_id})

		fmt.Println(count, countUser)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while looking for user already a friend"})
			return
		}

		if countUser > 0 && count == 0 {
			user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.ID = primitive.NewObjectID()
			user.IsRemoved = false

			resultInsertionNumber, insertErr := userfriendsCollection.InsertOne(ctx, user)
			if insertErr != nil {
				msg := fmt.Sprintf("user cannot be added as friend")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}
			defer cancel()

			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": resultInsertionNumber})
		} else {
			msg := fmt.Sprintf("No user identified")
			c.JSON(http.StatusNotFound, gin.H{"Status": false, "Result": msg})
		}

	}
}

//CreateUserprofile is the api used to tget a single user
func GetUserFriends() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserGetRequest

		var friends []models.UserFriends

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
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while checking for email"})
			return
		}

		fmt.Println(count, user.User_id)

		if count > 0 {
			////////////////////////////////////////////////////////////// Sales Staging Pipelines //////////////////////////////////////////////////
			matchStageUserProfiles := bson.D{{"$match", bson.D{{"userid", user.User_id}, {"Isremoved", false}}}}

			result, getErr := userfriendsCollection.Aggregate(ctx, mongo.Pipeline{matchStageUserProfiles})
			if getErr != nil {
				msg := fmt.Sprintf("User profile was not obtained")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}
			if err = result.All(ctx, &friends); err != nil {
				msg := fmt.Sprintf("User profile was not obtained")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": friends})
		} else {
			msg := fmt.Sprintf("Cannot retrieve friends list")
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}

	}
}

//CreateUserprofile is the api used to tget a single user
func GetUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserGetRequest
		var userInfo []models.UserUpdateRequest

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
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while checking for user existance"})
			return
		}

		fmt.Println(count)

		if count > 0 {
			////////////////////////////////////////////////////////////// Sales Staging Pipelines //////////////////////////////////////////////////
			matchStageUserProfiles := bson.D{{"$match", bson.D{{"user_id", user.User_id}, {"email", user.Email}}}}

			result, getErr := userCollection.Aggregate(ctx, mongo.Pipeline{matchStageUserProfiles})
			if getErr = result.All(ctx, &userInfo); getErr != nil {
				msg := fmt.Sprintf("User profile was not obtained", getErr.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}

			defer cancel()

			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": userInfo})
		}

	}
}
