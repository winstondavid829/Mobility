package controllers

import (
	"context"
	"entertainment/configs"
	"entertainment/models"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var BMICollection *mongo.Collection = configs.GetCollection(configs.DB, "Entertainment", "UserBMI")

func BMI_Calculation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var bmi models.BMI

		defer cancel()
		// if auth.ValidateUserTokenInHeader(c.Request) == false {
		// 	c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
		// 	return

		// }

		if err := c.BindJSON(&bmi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}

		validationErr := validate.Struct(bmi)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"user_id": bmi.UserID})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": "error occured while checking for the user"})
			return
		}
		if count > 0 {
			bmi.CreatedDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			bmi.ID = primitive.NewObjectID()

			height := get(bmi.Height)
			weight := get(bmi.Weight)

			bmi.BMI = weight / math.Pow(height/100, 2)

			resultInsertionNumber, insertErr := BMICollection.InsertOne(ctx, bmi)
			if insertErr != nil {
				msg := fmt.Sprintf("Cannot Insert User BMI Information")
				c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
				return
			}
			defer cancel()

			fmt.Println(resultInsertionNumber)

			c.JSON(http.StatusOK, gin.H{"Status": true, "Result": weight / math.Pow(height/100, 2)})
		} else {
			msg := fmt.Sprintf("Failed to create user BMI info")
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}

	}
}

func get(stri string) float64 {

	s, err := strconv.ParseFloat(stri, 64)

	if err != nil {
		fmt.Println("Cannot Parse string to float")

	}
	return s
}
