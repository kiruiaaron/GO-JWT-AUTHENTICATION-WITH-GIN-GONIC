package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
	"github.com/kiruiaaron/GO-JWT-AUTHENTICATION-WITH-GIN-GONIC/database"
	"go.mongodb.org/mongo-driver/mongo"
)


var userCollection *mongo.Collection = database.OpenCollection(database.CLient, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func SignUp()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func (c *gin.Context)  {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return

		}
	}
}