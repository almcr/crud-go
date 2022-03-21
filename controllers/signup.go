package controllers

import (
	"net/http"

	"github.com/almcr/crud-go/helper"
	"github.com/almcr/crud-go/models"
	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		// get username and password with basic auth
		user_id, pwd, hasAuth := ginCtx.Request.BasicAuth()

		if !hasAuth {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "No Authorisation header provided"})
			return
		}

		if _, ok := models.AuthUsers[user_id]; ok {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
			return
		}

		models.AuthUsers[user_id] = pwd
		token, refresh, _ := helper.GenerateTokens(user_id)
		models.AuthTokens[user_id] = models.TokenPair{token, refresh}

		ginCtx.JSON(http.StatusOK, gin.H{
			"Token":   token,
			"Refresh": refresh,
		})
	}
}
