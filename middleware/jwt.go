package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/config"
)

// Jwt ...
func Jwt() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "restricted",
		Key:           []byte(config.JWTKey()),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		Authenticator: authentication,
		Authorizator:  authorization,
		Unauthorized:  unauthorized,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func authentication(userId string, password string, c *gin.Context) (string, bool) {
	// TODO: implement
	if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		return userId, true
	}

	return userId, false
}

func authorization(userId string, c *gin.Context) bool {
	// TODO: implement
	if userId == "admin" {
		return true
	}

	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
