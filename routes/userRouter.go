package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kiruiaaron/GO-JWT-AUTHENTICATION-WITH-GIN-GONIC/controllers"
	"github.com/kiruiaaron/GO-JWT-AUTHENTICATION-WITH-GIN-GONIC/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUsers())
}
