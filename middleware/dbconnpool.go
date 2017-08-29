package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/db"
)

// DbConnPool passes a pool object for every request
func DbConnPool() gin.HandlerFunc {
	dbConnPool := db.ConnPool("main")
	return func(c *gin.Context) {
		c.Set("DBCONNPOOL", dbConnPool)
		c.Next()
	}
}
