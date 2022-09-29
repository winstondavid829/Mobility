package routes

import (
	controller "entertainment/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/v1/users/signup", controller.SignUp())
	incomingRoutes.POST("/v1/users/login", controller.Login())
	incomingRoutes.POST("/v1/userprofile/addInfo", controller.AddUserDetails())
	incomingRoutes.POST("/v1/friends/add", controller.AddFriendsToUser())
	incomingRoutes.POST("/v1/friends/getall", controller.GetUserFriends())
	incomingRoutes.POST("/v1/users/getDetails", controller.GetUserDetails())
	incomingRoutes.POST("/v1/users/bmiCalculation", controller.BMI_Calculation())
	incomingRoutes.POST("/v1/createtoken", controller.GetRSA256Token())

}
