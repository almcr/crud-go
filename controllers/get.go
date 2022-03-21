package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/almcr/crud-go/database"
	"github.com/almcr/crud-go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		var user models.UserData

		database.UsersCollection.FindOne(c, bson.M{"user_id": id}).Decode(&user)

		ctx.JSON(http.StatusOK, user)
	}
}
