package routes

import (
	"github.com/gin-gonic/gin"
	controller"github.com/kiruiaaron/GO-JWT-AUTHENTICATION-WITH-GIN-GONIC/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.SignUp())
	incomingRoutes.POST("users/login", controller.Login())
}