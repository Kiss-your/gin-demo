package middleware

import (
	"gin-demo/common"
	"gin-demo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)

		if user.ID == 0 {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
