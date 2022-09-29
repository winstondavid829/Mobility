package main

import (
	"context"
	"entertainment/configs"
	"entertainment/controllers"
	"entertainment/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	configs.IsProductionEnvironment(true)
	port := os.Getenv("PORT")
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.SetTrustedProxies([]string{"192.168.1.2", "*", "localhost"})
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.PostRoutes(router)

	// router.Use(middleware.Authentication())

	// API-1
	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	// API-2
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.POST("/v1/cloud-storage-bucket/insert", controllers.HandleFileUploadToBucket)
	// router.POST("/cloud-storage-bucket/getData", controllers.HandleFileDownloadFromBucket)

	router.Run(":" + port)

	// routes.Router(ch(router))
	srv := &http.Server{
		Handler: router,
		Addr:    "8000",
		// Good practice: enforce timeouts for servers you create!
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  55 * time.Second,
		WriteTimeout: 55 * time.Second,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	//This is for gracefully shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received request to terminate the server", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}

// func main() {
// 	configs.IsProductionEnvironment(true)
// 	controllers.AuthenticateSpotify()
// }
