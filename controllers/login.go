package controllers

import (
	"net/http"

	"github.com/almcr/crud-go/models"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		// get username and password with basic auth
		user_id, pwd, hasAuth := ginCtx.Request.BasicAuth()
		if !hasAuth {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "No Authorisation header provided"})
			return
		}

		// Check if user exist and credentials are correct
		if password, ok := models.AuthUsers[user_id]; !ok || pwd != password {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "wrong credentials"})
			return
		}

		ginCtx.JSON(http.StatusOK, models.AuthTokens[user_id])
	}
}
