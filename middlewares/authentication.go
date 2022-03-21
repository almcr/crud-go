package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/almcr/crud-go/helper"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		fmt.Println("In Auth middleware")
		requestToken := ginCtx.Request.Header.Get("Authorization")
		if requestToken == "" {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "No token found in header"})
			ginCtx.Abort()
			return
		}

		parts := strings.Fields(requestToken)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Authorization header must be bearer *token*"})
			ginCtx.Abort()
			return
		}

		claims, err := helper.CheckToken(parts[1])
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
			ginCtx.Abort()
			return
		}

		ginCtx.Set("user_id", claims.Uid)

		ginCtx.Next()
	}
}
