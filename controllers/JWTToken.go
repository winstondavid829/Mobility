package controllers

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetRSA256Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		///
		privateKey := os.Getenv("PRIVATE_KEY")
		Issuer := os.Getenv("ISSUER")
		Audience := os.Getenv("AUDIENCE")

		token, err := CreateRSA256Token(Audience, Issuer, privateKey)
		if err != nil {
			msg := fmt.Sprintf("Failed creating JWT Token")
			c.JSON(http.StatusInternalServerError, gin.H{"Status": false, "Result": msg})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Status": true, "Result": token})

	}
}

func CreateRSA256Token(audience, issuer, privatekey string) (string, error) {
	var err error
	//Setting Environmental Data at start up
	//data.IsProductionEnvironment(false)
	var privatePem = privatekey
	block, _ := pem.Decode([]byte(privatePem))
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	atClaims := jwt.MapClaims{}
	atClaims["iss"] = issuer
	atClaims["aud"] = audience
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	token, err := at.SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}
