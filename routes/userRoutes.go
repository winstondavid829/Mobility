package routes

import (
	controller "entertainment/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/v1/users/signup", controller.SignUp())
	incomingRoutes.POST("/v1/users/login", controller.Login())
}
