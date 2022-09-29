package controllers

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket(c *gin.Context) {
	bucket := "entertainment-app" //your bucket name

	var err error

	ctx := appengine.NewContext(c.Request)

	// if auth.ValidateUserTokenInHeader(c.Request) == false {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
	// 	return

	// }

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("bucket-access.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	myBucket_Base := os.Getenv("Bucket_Base_URL")

	u, err := url.Parse(myBucket_Base + "/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}

	finalPath := u.Host + u.EscapedPath()

	c.JSON(http.StatusOK, gin.H{
		"message": "file uploaded successfully",
		"Path":    finalPath,
	})
}

// HandleFileDownloadfromBucket uploads file to bucket
// func HandleFileDownloadFromBucket(c *gin.Context) {
// 	bucket := "entertainment-app" //your bucket name

// 	var err error

// 	ctx := appengine.NewContext(c.Request)

// 	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("bucket-access.json"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 			"error":   true,
// 		})
// 		return
// 	}

// 	// f, uploadedFile, err := c.Request.FormFile("file")
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{
// 	// 		"message": err.Error(),
// 	// 		"error":   true,
// 	// 	})
// 	// 	return
// 	// }

// 	// defer f.Close()

// 	sw, err := storageClient.Bucket(bucket).Object("Capture.PNG").NewReader(ctx)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 			"error":   true,
// 		})
// 		return

// 	}

// 	defer sw.Close()

// 	// if _, err := io.Copy(os.Stdout, sw); err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{
// 	// 		"message": err.Error(),
// 	// 		"error":   true,
// 	// 	})
// 	// 	return
// 	// }

// 	slurp, err := ioutil.ReadAll(sw)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 			"error":   true,
// 		})
// 		return

// 	}

// 	fmt.Println(slurp)
// 	pixels := make([]byte, 100*100) // slice of your gray pixels, size of 100x100

// 	img := image.NewGray(image.Rect(0, 0, 100, 100))
// 	img.Pix = pixels

// 	// if err := sw.Close(); err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{
// 	// 		"message": err.Error(),
// 	// 		"error":   true,
// 	// 	})
// 	// 	return
// 	// }

// 	// u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{
// 	// 		"message": err.Error(),
// 	// 		"Error":   true,
// 	// 	})
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":  "file uploaded successfully",
// 		"pathname": img,
// 	})
// }
