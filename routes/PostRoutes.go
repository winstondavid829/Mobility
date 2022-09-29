package routes

import (
	controller "entertainment/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func PostRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/v1/post/add", controller.CreateUserPost())
	incomingRoutes.POST("/v1/post/getuserpost", controller.GetAllUserPosts())
	incomingRoutes.POST("/v1/post/getfriendsPosts", controller.ViewFriendsPost())
	incomingRoutes.POST("/v1/post/postlike", controller.LikePost())
	incomingRoutes.POST("/v1/music/get", controller.AuthenticateSpotify())
	// incomingRoutes.POST("/v1/storage/insert", controller.HandleFileUploadToBucket())

}
